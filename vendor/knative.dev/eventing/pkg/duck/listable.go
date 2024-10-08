/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package duck

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/apis/duck"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/tracker"
)

// ListableTracker is a tracker capable of tracking any object that implements the apis.Listable interface.
type ListableTracker interface {
	// TrackInNamespace returns a function that can be used to watch arbitrary apis.Listable resources in the same
	// namespace as obj. Any change will cause a callback for obj.
	TrackInNamespace(ctx context.Context, obj metav1.Object) Track
	// TrackInNamespaceKReference returns a function that can be used to watch arbitrary apis.Listable resources
	// in the same namespace as obj. Any change will cause a callback for obj.
	TrackInNamespaceKReference(ctx context.Context, obj metav1.Object) TrackKReference
	// Track returns a function that can be used to watch arbitrary apis.Listable resources in the
	// provided namespace. Any change will cause a callback for obj.
	Track(ctx context.Context, obj metav1.Object, namespace string) Track
	// TrackKReference returns a function that can be used to watch arbitrary apis.Listable resources
	// in the provided namespace. Any change will cause a callback for obj.
	TrackKReference(ctx context.Context, obj metav1.Object, namespace string) TrackKReference
	// ListerFor returns the lister for the object reference. It returns an error if the lister does not exist.
	ListerFor(ref corev1.ObjectReference) (cache.GenericLister, error)
	// InformerFor returns the informer for the object reference. It returns an error if the informer does not exist.
	InformerFor(ref corev1.ObjectReference) (cache.SharedIndexInformer, error)
	// ListerFor returns the lister for the KReference. It returns an error if the lister does not exist.
	ListerForKReference(ref duckv1.KReference) (cache.GenericLister, error)
	// InformerFor returns the informer for the KReference. It returns an error if the informer does not exist.
	InformerForKReference(ref duckv1.KReference) (cache.SharedIndexInformer, error)
}

type Track func(corev1.ObjectReference) error
type TrackKReference func(duckv1.KReference) error

// NewListableTrackerFromTracker creates a new ListableTracker, backed by a TypedInformerFactory.
func NewListableTrackerFromTracker(ctx context.Context, getter func(context.Context) duck.InformerFactory, t tracker.Interface) ListableTracker {
	return &listableTracker{
		informerFactory: getter(ctx),
		tracker:         t,
		concrete:        map[schema.GroupVersionResource]informerListerPair{},
	}
}

type listableTracker struct {
	informerFactory duck.InformerFactory
	tracker         tracker.Interface

	concrete     map[schema.GroupVersionResource]informerListerPair
	concreteLock sync.RWMutex
}

type informerListerPair struct {
	informer cache.SharedIndexInformer
	lister   cache.GenericLister
}

// ensureTracking ensures that there is an informer watching and sending events to tracker for the
// concrete GVK. It also ensures that there is the corresponding lister for that informer.
func (t *listableTracker) ensureTracking(ctx context.Context, ref corev1.ObjectReference) error {
	if equality.Semantic.DeepEqual(ref, &corev1.ObjectReference{}) {
		return errors.New("cannot track empty object ref")
	}
	gvk := ref.GroupVersionKind()
	gvr, _ := meta.UnsafeGuessKindToResource(gvk)

	t.concreteLock.RLock()
	_, present := t.concrete[gvr]
	t.concreteLock.RUnlock()
	if present {
		// There is already an informer running for this GVR, we don't need or want to make
		// a second one.
		return nil
	}
	// Tracking has not been setup for this GVR.
	t.concreteLock.Lock()
	defer t.concreteLock.Unlock()
	// Now that we have acquired the write lock, make sure no one has added tracking handlers.
	if _, present = t.concrete[gvr]; present {
		return nil
	}
	informer, lister, err := t.informerFactory.Get(ctx, gvr)
	if err != nil {
		return err
	}
	informer.AddEventHandler(controller.HandleAll(
		// Call the tracker's OnChanged method, but we've seen the objects coming through
		// this path missing TypeMeta, so ensure it is properly populated.
		controller.EnsureTypeMeta(
			t.tracker.OnChanged,
			gvk,
		),
	))
	t.concrete[gvr] = informerListerPair{informer: informer, lister: lister}
	return nil
}

// TrackInNamespace satisfies the ListableTracker interface.
func (t *listableTracker) TrackInNamespace(ctx context.Context, obj metav1.Object) Track {
	return func(ref corev1.ObjectReference) error {
		// This is often used by Trigger and Subscription, both of which pass in refs that do not
		// specify the namespace.
		ref.Namespace = obj.GetNamespace()
		if err := t.ensureTracking(ctx, ref); err != nil {
			return err
		}

		return t.tracker.TrackReference(tracker.Reference{
			APIVersion: ref.APIVersion,
			Kind:       ref.Kind,
			Namespace:  obj.GetNamespace(),
			Name:       ref.Name,
		}, obj)
	}
}

// TrackInNamespaceKReference satisfies the ListableTracker interface.
func (t *listableTracker) TrackInNamespaceKReference(ctx context.Context, obj metav1.Object) TrackKReference {
	return func(ref duckv1.KReference) error {
		// This is often used by Trigger and Subscription, both of which pass in refs that do not
		// specify the namespace.
		ref.Namespace = obj.GetNamespace()
		coreRef := corev1.ObjectReference{APIVersion: ref.APIVersion, Kind: ref.Kind, Name: ref.Name, Namespace: ref.Namespace}
		if err := t.ensureTracking(ctx, coreRef); err != nil {
			return err
		}

		return t.tracker.TrackReference(tracker.Reference{
			APIVersion: ref.APIVersion,
			Kind:       ref.Kind,
			Namespace:  obj.GetNamespace(),
			Name:       ref.Name,
		}, obj)
	}
}

// Track satisfies the ListableTracker interface.
func (t *listableTracker) Track(ctx context.Context, obj metav1.Object, namespace string) Track {
	return func(ref corev1.ObjectReference) error {
		// This is often used by Trigger and Subscription, both of which pass in refs that do not
		// specify the namespace.
		ref.Namespace = namespace
		if err := t.ensureTracking(ctx, ref); err != nil {
			return err
		}

		return t.tracker.TrackReference(tracker.Reference{
			APIVersion: ref.APIVersion,
			Kind:       ref.Kind,
			Namespace:  namespace,
			Name:       ref.Name,
		}, obj)
	}
}

// TrackKReference satisfies the ListableTracker interface.
func (t *listableTracker) TrackKReference(ctx context.Context, obj metav1.Object, namespace string) TrackKReference {
	return func(ref duckv1.KReference) error {
		// This is often used by Trigger and Subscription, both of which pass in refs that do not
		// specify the namespace.
		ref.Namespace = namespace
		coreRef := corev1.ObjectReference{APIVersion: ref.APIVersion, Kind: ref.Kind, Name: ref.Name, Namespace: ref.Namespace}
		if err := t.ensureTracking(ctx, coreRef); err != nil {
			return err
		}

		return t.tracker.TrackReference(tracker.Reference{
			APIVersion: ref.APIVersion,
			Kind:       ref.Kind,
			Namespace:  namespace,
			Name:       ref.Name,
		}, obj)
	}
}

// ListerForKReference satisfies the ListableTracker interface.
func (t *listableTracker) ListerForKReference(ref duckv1.KReference) (cache.GenericLister, error) {
	return t.ListerFor(corev1.ObjectReference{APIVersion: ref.APIVersion, Kind: ref.Kind, Name: ref.Name, Namespace: ref.Namespace})
}

// ListerFor satisfies the ListableTracker interface.
func (t *listableTracker) ListerFor(ref corev1.ObjectReference) (cache.GenericLister, error) {
	if equality.Semantic.DeepEqual(ref, &corev1.ObjectReference{}) {
		return nil, errors.New("cannot get lister for empty object ref")
	}
	gvk := ref.GroupVersionKind()
	gvr, _ := meta.UnsafeGuessKindToResource(gvk)

	t.concreteLock.RLock()
	defer t.concreteLock.RUnlock()
	informerListerPair, present := t.concrete[gvr]
	if !present {
		return nil, fmt.Errorf("no lister available for GVR %s", gvr.String())
	}
	return informerListerPair.lister, nil
}

// InformerForKReference satisfies the ListableTracker interface.
func (t *listableTracker) InformerForKReference(ref duckv1.KReference) (cache.SharedIndexInformer, error) {
	return t.InformerFor(corev1.ObjectReference{APIVersion: ref.APIVersion, Kind: ref.Kind, Name: ref.Name, Namespace: ref.Namespace})
}

// InformerFor satisfies the ListableTracker interface.
func (t *listableTracker) InformerFor(ref corev1.ObjectReference) (cache.SharedIndexInformer, error) {
	if equality.Semantic.DeepEqual(ref, &corev1.ObjectReference{}) {
		return nil, errors.New("cannot get informer for empty object ref")
	}
	gvk := ref.GroupVersionKind()
	gvr, _ := meta.UnsafeGuessKindToResource(gvk)

	t.concreteLock.RLock()
	defer t.concreteLock.RUnlock()
	informerListerPair, present := t.concrete[gvr]
	if !present {
		return nil, fmt.Errorf("no informer available for GVR %s", gvr.String())
	}
	return informerListerPair.informer, nil
}
