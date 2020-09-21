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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/eventing/pkg/apis/sources/v1beta1"
)

// ApiServerSourceLister helps list ApiServerSources.
type ApiServerSourceLister interface {
	// List lists all ApiServerSources in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ApiServerSource, err error)
	// ApiServerSources returns an object that can list and get ApiServerSources.
	ApiServerSources(namespace string) ApiServerSourceNamespaceLister
	ApiServerSourceListerExpansion
}

// apiServerSourceLister implements the ApiServerSourceLister interface.
type apiServerSourceLister struct {
	indexer cache.Indexer
}

// NewApiServerSourceLister returns a new ApiServerSourceLister.
func NewApiServerSourceLister(indexer cache.Indexer) ApiServerSourceLister {
	return &apiServerSourceLister{indexer: indexer}
}

// List lists all ApiServerSources in the indexer.
func (s *apiServerSourceLister) List(selector labels.Selector) (ret []*v1beta1.ApiServerSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ApiServerSource))
	})
	return ret, err
}

// ApiServerSources returns an object that can list and get ApiServerSources.
func (s *apiServerSourceLister) ApiServerSources(namespace string) ApiServerSourceNamespaceLister {
	return apiServerSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiServerSourceNamespaceLister helps list and get ApiServerSources.
type ApiServerSourceNamespaceLister interface {
	// List lists all ApiServerSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.ApiServerSource, err error)
	// Get retrieves the ApiServerSource from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.ApiServerSource, error)
	ApiServerSourceNamespaceListerExpansion
}

// apiServerSourceNamespaceLister implements the ApiServerSourceNamespaceLister
// interface.
type apiServerSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiServerSources in the indexer for a given namespace.
func (s apiServerSourceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ApiServerSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ApiServerSource))
	})
	return ret, err
}

// Get retrieves the ApiServerSource from the indexer for a given namespace and name.
func (s apiServerSourceNamespaceLister) Get(name string) (*v1beta1.ApiServerSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("apiserversource"), name)
	}
	return obj.(*v1beta1.ApiServerSource), nil
}
