//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
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

package v1alpha

import (
	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/backup/v1alpha"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoBackupPolicyLister helps list ArangoBackupPolicies.
type ArangoBackupPolicyLister interface {
	// List lists all ArangoBackupPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha.ArangoBackupPolicy, err error)
	// ArangoBackupPolicies returns an object that can list and get ArangoBackupPolicies.
	ArangoBackupPolicies(namespace string) ArangoBackupPolicyNamespaceLister
	ArangoBackupPolicyListerExpansion
}

// arangoBackupPolicyLister implements the ArangoBackupPolicyLister interface.
type arangoBackupPolicyLister struct {
	indexer cache.Indexer
}

// NewArangoBackupPolicyLister returns a new ArangoBackupPolicyLister.
func NewArangoBackupPolicyLister(indexer cache.Indexer) ArangoBackupPolicyLister {
	return &arangoBackupPolicyLister{indexer: indexer}
}

// List lists all ArangoBackupPolicies in the indexer.
func (s *arangoBackupPolicyLister) List(selector labels.Selector) (ret []*v1alpha.ArangoBackupPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha.ArangoBackupPolicy))
	})
	return ret, err
}

// ArangoBackupPolicies returns an object that can list and get ArangoBackupPolicies.
func (s *arangoBackupPolicyLister) ArangoBackupPolicies(namespace string) ArangoBackupPolicyNamespaceLister {
	return arangoBackupPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoBackupPolicyNamespaceLister helps list and get ArangoBackupPolicies.
type ArangoBackupPolicyNamespaceLister interface {
	// List lists all ArangoBackupPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha.ArangoBackupPolicy, err error)
	// Get retrieves the ArangoBackupPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha.ArangoBackupPolicy, error)
	ArangoBackupPolicyNamespaceListerExpansion
}

// arangoBackupPolicyNamespaceLister implements the ArangoBackupPolicyNamespaceLister
// interface.
type arangoBackupPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoBackupPolicies in the indexer for a given namespace.
func (s arangoBackupPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha.ArangoBackupPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha.ArangoBackupPolicy))
	})
	return ret, err
}

// Get retrieves the ArangoBackupPolicy from the indexer for a given namespace and name.
func (s arangoBackupPolicyNamespaceLister) Get(name string) (*v1alpha.ArangoBackupPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha.Resource("arangobackuppolicy"), name)
	}
	return obj.(*v1alpha.ArangoBackupPolicy), nil
}
