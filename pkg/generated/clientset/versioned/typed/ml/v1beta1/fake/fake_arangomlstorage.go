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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/arangodb/kube-arangodb/pkg/apis/ml/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoMLStorages implements ArangoMLStorageInterface
type FakeArangoMLStorages struct {
	Fake *FakeMlV1beta1
	ns   string
}

var arangomlstoragesResource = v1beta1.SchemeGroupVersion.WithResource("arangomlstorages")

var arangomlstoragesKind = v1beta1.SchemeGroupVersion.WithKind("ArangoMLStorage")

// Get takes name of the arangoMLStorage, and returns the corresponding arangoMLStorage object, and an error if there is any.
func (c *FakeArangoMLStorages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ArangoMLStorage, err error) {
	emptyResult := &v1beta1.ArangoMLStorage{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(arangomlstoragesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoMLStorage), err
}

// List takes label and field selectors, and returns the list of ArangoMLStorages that match those selectors.
func (c *FakeArangoMLStorages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ArangoMLStorageList, err error) {
	emptyResult := &v1beta1.ArangoMLStorageList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(arangomlstoragesResource, arangomlstoragesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ArangoMLStorageList{ListMeta: obj.(*v1beta1.ArangoMLStorageList).ListMeta}
	for _, item := range obj.(*v1beta1.ArangoMLStorageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoMLStorages.
func (c *FakeArangoMLStorages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(arangomlstoragesResource, c.ns, opts))

}

// Create takes the representation of a arangoMLStorage and creates it.  Returns the server's representation of the arangoMLStorage, and an error, if there is any.
func (c *FakeArangoMLStorages) Create(ctx context.Context, arangoMLStorage *v1beta1.ArangoMLStorage, opts v1.CreateOptions) (result *v1beta1.ArangoMLStorage, err error) {
	emptyResult := &v1beta1.ArangoMLStorage{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(arangomlstoragesResource, c.ns, arangoMLStorage, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoMLStorage), err
}

// Update takes the representation of a arangoMLStorage and updates it. Returns the server's representation of the arangoMLStorage, and an error, if there is any.
func (c *FakeArangoMLStorages) Update(ctx context.Context, arangoMLStorage *v1beta1.ArangoMLStorage, opts v1.UpdateOptions) (result *v1beta1.ArangoMLStorage, err error) {
	emptyResult := &v1beta1.ArangoMLStorage{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(arangomlstoragesResource, c.ns, arangoMLStorage, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoMLStorage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoMLStorages) UpdateStatus(ctx context.Context, arangoMLStorage *v1beta1.ArangoMLStorage, opts v1.UpdateOptions) (result *v1beta1.ArangoMLStorage, err error) {
	emptyResult := &v1beta1.ArangoMLStorage{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(arangomlstoragesResource, "status", c.ns, arangoMLStorage, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoMLStorage), err
}

// Delete takes name of the arangoMLStorage and deletes it. Returns an error if one occurs.
func (c *FakeArangoMLStorages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(arangomlstoragesResource, c.ns, name, opts), &v1beta1.ArangoMLStorage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoMLStorages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(arangomlstoragesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ArangoMLStorageList{})
	return err
}

// Patch applies the patch and returns the patched arangoMLStorage.
func (c *FakeArangoMLStorages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArangoMLStorage, err error) {
	emptyResult := &v1beta1.ArangoMLStorage{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(arangomlstoragesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoMLStorage), err
}
