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

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	sourcesv1alpha1 "knative.dev/eventing-ceph/pkg/apis/sources/v1alpha1"
	scheme "knative.dev/eventing-ceph/pkg/client/clientset/versioned/scheme"
)

// CephSourcesGetter has a method to return a CephSourceInterface.
// A group's client should implement this interface.
type CephSourcesGetter interface {
	CephSources(namespace string) CephSourceInterface
}

// CephSourceInterface has methods to work with CephSource resources.
type CephSourceInterface interface {
	Create(ctx context.Context, cephSource *sourcesv1alpha1.CephSource, opts v1.CreateOptions) (*sourcesv1alpha1.CephSource, error)
	Update(ctx context.Context, cephSource *sourcesv1alpha1.CephSource, opts v1.UpdateOptions) (*sourcesv1alpha1.CephSource, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, cephSource *sourcesv1alpha1.CephSource, opts v1.UpdateOptions) (*sourcesv1alpha1.CephSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*sourcesv1alpha1.CephSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*sourcesv1alpha1.CephSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *sourcesv1alpha1.CephSource, err error)
	CephSourceExpansion
}

// cephSources implements CephSourceInterface
type cephSources struct {
	*gentype.ClientWithList[*sourcesv1alpha1.CephSource, *sourcesv1alpha1.CephSourceList]
}

// newCephSources returns a CephSources
func newCephSources(c *SourcesV1alpha1Client, namespace string) *cephSources {
	return &cephSources{
		gentype.NewClientWithList[*sourcesv1alpha1.CephSource, *sourcesv1alpha1.CephSourceList](
			"cephsources",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *sourcesv1alpha1.CephSource { return &sourcesv1alpha1.CephSource{} },
			func() *sourcesv1alpha1.CephSourceList { return &sourcesv1alpha1.CephSourceList{} },
		),
	}
}
