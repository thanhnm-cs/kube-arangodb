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
	clientset "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	analyticsv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/analytics/v1alpha1"
	fakeanalyticsv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/analytics/v1alpha1/fake"
	appsv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/apps/v1"
	fakeappsv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/apps/v1/fake"
	backupv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/backup/v1"
	fakebackupv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/backup/v1/fake"
	databasev1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/deployment/v1"
	fakedatabasev1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/deployment/v1/fake"
	databasev2alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/deployment/v2alpha1"
	fakedatabasev2alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/deployment/v2alpha1/fake"
	mlv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/ml/v1alpha1"
	fakemlv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/ml/v1alpha1/fake"
	mlv1beta1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/ml/v1beta1"
	fakemlv1beta1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/ml/v1beta1/fake"
	networkingv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/networking/v1alpha1"
	fakenetworkingv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/networking/v1alpha1/fake"
	replicationv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/replication/v1"
	fakereplicationv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/replication/v1/fake"
	replicationv2alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/replication/v2alpha1"
	fakereplicationv2alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/replication/v2alpha1/fake"
	schedulerv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/scheduler/v1alpha1"
	fakeschedulerv1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/scheduler/v1alpha1/fake"
	schedulerv1beta1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/scheduler/v1beta1"
	fakeschedulerv1beta1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/scheduler/v1beta1/fake"
	storagev1alpha "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/storage/v1alpha"
	fakestoragev1alpha "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/storage/v1alpha/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any field management, validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
//
// DEPRECATED: NewClientset replaces this with support for field management, which significantly improves
// server side apply testing. NewClientset is only available when apply configurations are generated (e.g.
// via --with-applyconfig).
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// AnalyticsV1alpha1 retrieves the AnalyticsV1alpha1Client
func (c *Clientset) AnalyticsV1alpha1() analyticsv1alpha1.AnalyticsV1alpha1Interface {
	return &fakeanalyticsv1alpha1.FakeAnalyticsV1alpha1{Fake: &c.Fake}
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return &fakeappsv1.FakeAppsV1{Fake: &c.Fake}
}

// BackupV1 retrieves the BackupV1Client
func (c *Clientset) BackupV1() backupv1.BackupV1Interface {
	return &fakebackupv1.FakeBackupV1{Fake: &c.Fake}
}

// DatabaseV1 retrieves the DatabaseV1Client
func (c *Clientset) DatabaseV1() databasev1.DatabaseV1Interface {
	return &fakedatabasev1.FakeDatabaseV1{Fake: &c.Fake}
}

// DatabaseV2alpha1 retrieves the DatabaseV2alpha1Client
func (c *Clientset) DatabaseV2alpha1() databasev2alpha1.DatabaseV2alpha1Interface {
	return &fakedatabasev2alpha1.FakeDatabaseV2alpha1{Fake: &c.Fake}
}

// MlV1alpha1 retrieves the MlV1alpha1Client
func (c *Clientset) MlV1alpha1() mlv1alpha1.MlV1alpha1Interface {
	return &fakemlv1alpha1.FakeMlV1alpha1{Fake: &c.Fake}
}

// MlV1beta1 retrieves the MlV1beta1Client
func (c *Clientset) MlV1beta1() mlv1beta1.MlV1beta1Interface {
	return &fakemlv1beta1.FakeMlV1beta1{Fake: &c.Fake}
}

// NetworkingV1alpha1 retrieves the NetworkingV1alpha1Client
func (c *Clientset) NetworkingV1alpha1() networkingv1alpha1.NetworkingV1alpha1Interface {
	return &fakenetworkingv1alpha1.FakeNetworkingV1alpha1{Fake: &c.Fake}
}

// ReplicationV1 retrieves the ReplicationV1Client
func (c *Clientset) ReplicationV1() replicationv1.ReplicationV1Interface {
	return &fakereplicationv1.FakeReplicationV1{Fake: &c.Fake}
}

// ReplicationV2alpha1 retrieves the ReplicationV2alpha1Client
func (c *Clientset) ReplicationV2alpha1() replicationv2alpha1.ReplicationV2alpha1Interface {
	return &fakereplicationv2alpha1.FakeReplicationV2alpha1{Fake: &c.Fake}
}

// SchedulerV1alpha1 retrieves the SchedulerV1alpha1Client
func (c *Clientset) SchedulerV1alpha1() schedulerv1alpha1.SchedulerV1alpha1Interface {
	return &fakeschedulerv1alpha1.FakeSchedulerV1alpha1{Fake: &c.Fake}
}

// SchedulerV1beta1 retrieves the SchedulerV1beta1Client
func (c *Clientset) SchedulerV1beta1() schedulerv1beta1.SchedulerV1beta1Interface {
	return &fakeschedulerv1beta1.FakeSchedulerV1beta1{Fake: &c.Fake}
}

// StorageV1alpha retrieves the StorageV1alphaClient
func (c *Clientset) StorageV1alpha() storagev1alpha.StorageV1alphaInterface {
	return &fakestoragev1alpha.FakeStorageV1alpha{Fake: &c.Fake}
}
