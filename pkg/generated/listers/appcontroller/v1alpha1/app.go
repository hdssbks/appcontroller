/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "appcontroller/pkg/apis/appcontroller/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppLister helps list Apps.
// All objects returned here must be treated as read-only.
type AppLister interface {
	// List lists all Apps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.App, err error)
	// Apps returns an object that can list and get Apps.
	Apps(namespace string) AppNamespaceLister
	AppListerExpansion
}

// appLister implements the AppLister interface.
type appLister struct {
	indexer cache.Indexer
}

// NewAppLister returns a new AppLister.
func NewAppLister(indexer cache.Indexer) AppLister {
	return &appLister{indexer: indexer}
}

// List lists all Apps in the indexer.
func (s *appLister) List(selector labels.Selector) (ret []*v1alpha1.App, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.App))
	})
	return ret, err
}

// Apps returns an object that can list and get Apps.
func (s *appLister) Apps(namespace string) AppNamespaceLister {
	return appNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppNamespaceLister helps list and get Apps.
// All objects returned here must be treated as read-only.
type AppNamespaceLister interface {
	// List lists all Apps in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.App, err error)
	// Get retrieves the App from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.App, error)
	AppNamespaceListerExpansion
}

// appNamespaceLister implements the AppNamespaceLister
// interface.
type appNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Apps in the indexer for a given namespace.
func (s appNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.App, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.App))
	})
	return ret, err
}

// Get retrieves the App from the indexer for a given namespace and name.
func (s appNamespaceLister) Get(name string) (*v1alpha1.App, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("app"), name)
	}
	return obj.(*v1alpha1.App), nil
}
