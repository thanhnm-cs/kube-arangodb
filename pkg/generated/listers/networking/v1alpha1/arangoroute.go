//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/networking/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoRouteLister helps list ArangoRoutes.
// All objects returned here must be treated as read-only.
type ArangoRouteLister interface {
	// List lists all ArangoRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArangoRoute, err error)
	// ArangoRoutes returns an object that can list and get ArangoRoutes.
	ArangoRoutes(namespace string) ArangoRouteNamespaceLister
	ArangoRouteListerExpansion
}

// arangoRouteLister implements the ArangoRouteLister interface.
type arangoRouteLister struct {
	indexer cache.Indexer
}

// NewArangoRouteLister returns a new ArangoRouteLister.
func NewArangoRouteLister(indexer cache.Indexer) ArangoRouteLister {
	return &arangoRouteLister{indexer: indexer}
}

// List lists all ArangoRoutes in the indexer.
func (s *arangoRouteLister) List(selector labels.Selector) (ret []*v1alpha1.ArangoRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArangoRoute))
	})
	return ret, err
}

// ArangoRoutes returns an object that can list and get ArangoRoutes.
func (s *arangoRouteLister) ArangoRoutes(namespace string) ArangoRouteNamespaceLister {
	return arangoRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoRouteNamespaceLister helps list and get ArangoRoutes.
// All objects returned here must be treated as read-only.
type ArangoRouteNamespaceLister interface {
	// List lists all ArangoRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArangoRoute, err error)
	// Get retrieves the ArangoRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ArangoRoute, error)
	ArangoRouteNamespaceListerExpansion
}

// arangoRouteNamespaceLister implements the ArangoRouteNamespaceLister
// interface.
type arangoRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoRoutes in the indexer for a given namespace.
func (s arangoRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ArangoRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArangoRoute))
	})
	return ret, err
}

// Get retrieves the ArangoRoute from the indexer for a given namespace and name.
func (s arangoRouteNamespaceLister) Get(name string) (*v1alpha1.ArangoRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("arangoroute"), name)
	}
	return obj.(*v1alpha1.ArangoRoute), nil
}
