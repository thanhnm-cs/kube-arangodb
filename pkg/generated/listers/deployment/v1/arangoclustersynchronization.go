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

package v1

import (
	v1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ArangoClusterSynchronizationLister helps list ArangoClusterSynchronizations.
// All objects returned here must be treated as read-only.
type ArangoClusterSynchronizationLister interface {
	// List lists all ArangoClusterSynchronizations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoClusterSynchronization, err error)
	// ArangoClusterSynchronizations returns an object that can list and get ArangoClusterSynchronizations.
	ArangoClusterSynchronizations(namespace string) ArangoClusterSynchronizationNamespaceLister
	ArangoClusterSynchronizationListerExpansion
}

// arangoClusterSynchronizationLister implements the ArangoClusterSynchronizationLister interface.
type arangoClusterSynchronizationLister struct {
	listers.ResourceIndexer[*v1.ArangoClusterSynchronization]
}

// NewArangoClusterSynchronizationLister returns a new ArangoClusterSynchronizationLister.
func NewArangoClusterSynchronizationLister(indexer cache.Indexer) ArangoClusterSynchronizationLister {
	return &arangoClusterSynchronizationLister{listers.New[*v1.ArangoClusterSynchronization](indexer, v1.Resource("arangoclustersynchronization"))}
}

// ArangoClusterSynchronizations returns an object that can list and get ArangoClusterSynchronizations.
func (s *arangoClusterSynchronizationLister) ArangoClusterSynchronizations(namespace string) ArangoClusterSynchronizationNamespaceLister {
	return arangoClusterSynchronizationNamespaceLister{listers.NewNamespaced[*v1.ArangoClusterSynchronization](s.ResourceIndexer, namespace)}
}

// ArangoClusterSynchronizationNamespaceLister helps list and get ArangoClusterSynchronizations.
// All objects returned here must be treated as read-only.
type ArangoClusterSynchronizationNamespaceLister interface {
	// List lists all ArangoClusterSynchronizations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoClusterSynchronization, err error)
	// Get retrieves the ArangoClusterSynchronization from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ArangoClusterSynchronization, error)
	ArangoClusterSynchronizationNamespaceListerExpansion
}

// arangoClusterSynchronizationNamespaceLister implements the ArangoClusterSynchronizationNamespaceLister
// interface.
type arangoClusterSynchronizationNamespaceLister struct {
	listers.ResourceIndexer[*v1.ArangoClusterSynchronization]
}
