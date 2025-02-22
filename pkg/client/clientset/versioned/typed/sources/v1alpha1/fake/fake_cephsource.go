/*
Copyright 2020 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-ceph/pkg/apis/sources/v1alpha1"
)

// FakeCephSources implements CephSourceInterface
type FakeCephSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var cephsourcesResource = v1alpha1.SchemeGroupVersion.WithResource("cephsources")

var cephsourcesKind = v1alpha1.SchemeGroupVersion.WithKind("CephSource")

// Get takes name of the cephSource, and returns the corresponding cephSource object, and an error if there is any.
func (c *FakeCephSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CephSource, err error) {
	emptyResult := &v1alpha1.CephSource{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(cephsourcesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CephSource), err
}

// List takes label and field selectors, and returns the list of CephSources that match those selectors.
func (c *FakeCephSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CephSourceList, err error) {
	emptyResult := &v1alpha1.CephSourceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(cephsourcesResource, cephsourcesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CephSourceList{ListMeta: obj.(*v1alpha1.CephSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CephSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephSources.
func (c *FakeCephSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(cephsourcesResource, c.ns, opts))

}

// Create takes the representation of a cephSource and creates it.  Returns the server's representation of the cephSource, and an error, if there is any.
func (c *FakeCephSources) Create(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.CreateOptions) (result *v1alpha1.CephSource, err error) {
	emptyResult := &v1alpha1.CephSource{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(cephsourcesResource, c.ns, cephSource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CephSource), err
}

// Update takes the representation of a cephSource and updates it. Returns the server's representation of the cephSource, and an error, if there is any.
func (c *FakeCephSources) Update(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.UpdateOptions) (result *v1alpha1.CephSource, err error) {
	emptyResult := &v1alpha1.CephSource{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(cephsourcesResource, c.ns, cephSource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CephSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCephSources) UpdateStatus(ctx context.Context, cephSource *v1alpha1.CephSource, opts v1.UpdateOptions) (result *v1alpha1.CephSource, err error) {
	emptyResult := &v1alpha1.CephSource{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(cephsourcesResource, "status", c.ns, cephSource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CephSource), err
}

// Delete takes name of the cephSource and deletes it. Returns an error if one occurs.
func (c *FakeCephSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cephsourcesResource, c.ns, name, opts), &v1alpha1.CephSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(cephsourcesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CephSourceList{})
	return err
}

// Patch applies the patch and returns the patched cephSource.
func (c *FakeCephSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CephSource, err error) {
	emptyResult := &v1alpha1.CephSource{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(cephsourcesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CephSource), err
}
