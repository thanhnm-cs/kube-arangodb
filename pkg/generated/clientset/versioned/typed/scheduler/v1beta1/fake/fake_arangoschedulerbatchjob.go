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

	v1beta1 "github.com/arangodb/kube-arangodb/pkg/apis/scheduler/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoSchedulerBatchJobs implements ArangoSchedulerBatchJobInterface
type FakeArangoSchedulerBatchJobs struct {
	Fake *FakeSchedulerV1beta1
	ns   string
}

var arangoschedulerbatchjobsResource = v1beta1.SchemeGroupVersion.WithResource("arangoschedulerbatchjobs")

var arangoschedulerbatchjobsKind = v1beta1.SchemeGroupVersion.WithKind("ArangoSchedulerBatchJob")

// Get takes name of the arangoSchedulerBatchJob, and returns the corresponding arangoSchedulerBatchJob object, and an error if there is any.
func (c *FakeArangoSchedulerBatchJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ArangoSchedulerBatchJob, err error) {
	emptyResult := &v1beta1.ArangoSchedulerBatchJob{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(arangoschedulerbatchjobsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoSchedulerBatchJob), err
}

// List takes label and field selectors, and returns the list of ArangoSchedulerBatchJobs that match those selectors.
func (c *FakeArangoSchedulerBatchJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ArangoSchedulerBatchJobList, err error) {
	emptyResult := &v1beta1.ArangoSchedulerBatchJobList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(arangoschedulerbatchjobsResource, arangoschedulerbatchjobsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ArangoSchedulerBatchJobList{ListMeta: obj.(*v1beta1.ArangoSchedulerBatchJobList).ListMeta}
	for _, item := range obj.(*v1beta1.ArangoSchedulerBatchJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoSchedulerBatchJobs.
func (c *FakeArangoSchedulerBatchJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(arangoschedulerbatchjobsResource, c.ns, opts))

}

// Create takes the representation of a arangoSchedulerBatchJob and creates it.  Returns the server's representation of the arangoSchedulerBatchJob, and an error, if there is any.
func (c *FakeArangoSchedulerBatchJobs) Create(ctx context.Context, arangoSchedulerBatchJob *v1beta1.ArangoSchedulerBatchJob, opts v1.CreateOptions) (result *v1beta1.ArangoSchedulerBatchJob, err error) {
	emptyResult := &v1beta1.ArangoSchedulerBatchJob{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(arangoschedulerbatchjobsResource, c.ns, arangoSchedulerBatchJob, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoSchedulerBatchJob), err
}

// Update takes the representation of a arangoSchedulerBatchJob and updates it. Returns the server's representation of the arangoSchedulerBatchJob, and an error, if there is any.
func (c *FakeArangoSchedulerBatchJobs) Update(ctx context.Context, arangoSchedulerBatchJob *v1beta1.ArangoSchedulerBatchJob, opts v1.UpdateOptions) (result *v1beta1.ArangoSchedulerBatchJob, err error) {
	emptyResult := &v1beta1.ArangoSchedulerBatchJob{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(arangoschedulerbatchjobsResource, c.ns, arangoSchedulerBatchJob, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoSchedulerBatchJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoSchedulerBatchJobs) UpdateStatus(ctx context.Context, arangoSchedulerBatchJob *v1beta1.ArangoSchedulerBatchJob, opts v1.UpdateOptions) (result *v1beta1.ArangoSchedulerBatchJob, err error) {
	emptyResult := &v1beta1.ArangoSchedulerBatchJob{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(arangoschedulerbatchjobsResource, "status", c.ns, arangoSchedulerBatchJob, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoSchedulerBatchJob), err
}

// Delete takes name of the arangoSchedulerBatchJob and deletes it. Returns an error if one occurs.
func (c *FakeArangoSchedulerBatchJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(arangoschedulerbatchjobsResource, c.ns, name, opts), &v1beta1.ArangoSchedulerBatchJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoSchedulerBatchJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(arangoschedulerbatchjobsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ArangoSchedulerBatchJobList{})
	return err
}

// Patch applies the patch and returns the patched arangoSchedulerBatchJob.
func (c *FakeArangoSchedulerBatchJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArangoSchedulerBatchJob, err error) {
	emptyResult := &v1beta1.ArangoSchedulerBatchJob{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(arangoschedulerbatchjobsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ArangoSchedulerBatchJob), err
}
