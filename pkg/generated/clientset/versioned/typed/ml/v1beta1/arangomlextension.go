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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/arangodb/kube-arangodb/pkg/apis/ml/v1beta1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoMLExtensionsGetter has a method to return a ArangoMLExtensionInterface.
// A group's client should implement this interface.
type ArangoMLExtensionsGetter interface {
	ArangoMLExtensions(namespace string) ArangoMLExtensionInterface
}

// ArangoMLExtensionInterface has methods to work with ArangoMLExtension resources.
type ArangoMLExtensionInterface interface {
	Create(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.CreateOptions) (*v1beta1.ArangoMLExtension, error)
	Update(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.UpdateOptions) (*v1beta1.ArangoMLExtension, error)
	UpdateStatus(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.UpdateOptions) (*v1beta1.ArangoMLExtension, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ArangoMLExtension, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ArangoMLExtensionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArangoMLExtension, err error)
	ArangoMLExtensionExpansion
}

// arangoMLExtensions implements ArangoMLExtensionInterface
type arangoMLExtensions struct {
	client rest.Interface
	ns     string
}

// newArangoMLExtensions returns a ArangoMLExtensions
func newArangoMLExtensions(c *MlV1beta1Client, namespace string) *arangoMLExtensions {
	return &arangoMLExtensions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoMLExtension, and returns the corresponding arangoMLExtension object, and an error if there is any.
func (c *arangoMLExtensions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ArangoMLExtension, err error) {
	result = &v1beta1.ArangoMLExtension{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangomlextensions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoMLExtensions that match those selectors.
func (c *arangoMLExtensions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ArangoMLExtensionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ArangoMLExtensionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangomlextensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoMLExtensions.
func (c *arangoMLExtensions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangomlextensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a arangoMLExtension and creates it.  Returns the server's representation of the arangoMLExtension, and an error, if there is any.
func (c *arangoMLExtensions) Create(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.CreateOptions) (result *v1beta1.ArangoMLExtension, err error) {
	result = &v1beta1.ArangoMLExtension{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangomlextensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoMLExtension).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a arangoMLExtension and updates it. Returns the server's representation of the arangoMLExtension, and an error, if there is any.
func (c *arangoMLExtensions) Update(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.UpdateOptions) (result *v1beta1.ArangoMLExtension, err error) {
	result = &v1beta1.ArangoMLExtension{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangomlextensions").
		Name(arangoMLExtension.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoMLExtension).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *arangoMLExtensions) UpdateStatus(ctx context.Context, arangoMLExtension *v1beta1.ArangoMLExtension, opts v1.UpdateOptions) (result *v1beta1.ArangoMLExtension, err error) {
	result = &v1beta1.ArangoMLExtension{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangomlextensions").
		Name(arangoMLExtension.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoMLExtension).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the arangoMLExtension and deletes it. Returns an error if one occurs.
func (c *arangoMLExtensions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangomlextensions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoMLExtensions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangomlextensions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched arangoMLExtension.
func (c *arangoMLExtensions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArangoMLExtension, err error) {
	result = &v1beta1.ArangoMLExtension{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangomlextensions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
