/*
Copyright AppsCode Inc. and Contributors

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
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-ec-api/apis/deployment/v1alpha1"
	scheme "kubeform.dev/provider-ec-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ElasticsearchKeystoresGetter has a method to return a ElasticsearchKeystoreInterface.
// A group's client should implement this interface.
type ElasticsearchKeystoresGetter interface {
	ElasticsearchKeystores(namespace string) ElasticsearchKeystoreInterface
}

// ElasticsearchKeystoreInterface has methods to work with ElasticsearchKeystore resources.
type ElasticsearchKeystoreInterface interface {
	Create(ctx context.Context, elasticsearchKeystore *v1alpha1.ElasticsearchKeystore, opts v1.CreateOptions) (*v1alpha1.ElasticsearchKeystore, error)
	Update(ctx context.Context, elasticsearchKeystore *v1alpha1.ElasticsearchKeystore, opts v1.UpdateOptions) (*v1alpha1.ElasticsearchKeystore, error)
	UpdateStatus(ctx context.Context, elasticsearchKeystore *v1alpha1.ElasticsearchKeystore, opts v1.UpdateOptions) (*v1alpha1.ElasticsearchKeystore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ElasticsearchKeystore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ElasticsearchKeystoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ElasticsearchKeystore, err error)
	ElasticsearchKeystoreExpansion
}

// elasticsearchKeystores implements ElasticsearchKeystoreInterface
type elasticsearchKeystores struct {
	client rest.Interface
	ns     string
}

// newElasticsearchKeystores returns a ElasticsearchKeystores
func newElasticsearchKeystores(c *DeploymentV1alpha1Client, namespace string) *elasticsearchKeystores {
	return &elasticsearchKeystores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the elasticsearchKeystore, and returns the corresponding elasticsearchKeystore object, and an error if there is any.
func (c *elasticsearchKeystores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ElasticsearchKeystore, err error) {
	result = &v1alpha1.ElasticsearchKeystore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElasticsearchKeystores that match those selectors.
func (c *elasticsearchKeystores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ElasticsearchKeystoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElasticsearchKeystoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elasticsearchKeystores.
func (c *elasticsearchKeystores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a elasticsearchKeystore and creates it.  Returns the server's representation of the elasticsearchKeystore, and an error, if there is any.
func (c *elasticsearchKeystores) Create(ctx context.Context, elasticsearchKeystore *v1alpha1.ElasticsearchKeystore, opts v1.CreateOptions) (result *v1alpha1.ElasticsearchKeystore, err error) {
	result = &v1alpha1.ElasticsearchKeystore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticsearchKeystore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a elasticsearchKeystore and updates it. Returns the server's representation of the elasticsearchKeystore, and an error, if there is any.
func (c *elasticsearchKeystores) Update(ctx context.Context, elasticsearchKeystore *v1alpha1.ElasticsearchKeystore, opts v1.UpdateOptions) (result *v1alpha1.ElasticsearchKeystore, err error) {
	result = &v1alpha1.ElasticsearchKeystore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		Name(elasticsearchKeystore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticsearchKeystore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *elasticsearchKeystores) UpdateStatus(ctx context.Context, elasticsearchKeystore *v1alpha1.ElasticsearchKeystore, opts v1.UpdateOptions) (result *v1alpha1.ElasticsearchKeystore, err error) {
	result = &v1alpha1.ElasticsearchKeystore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		Name(elasticsearchKeystore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticsearchKeystore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the elasticsearchKeystore and deletes it. Returns an error if one occurs.
func (c *elasticsearchKeystores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elasticsearchKeystores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched elasticsearchKeystore.
func (c *elasticsearchKeystores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ElasticsearchKeystore, err error) {
	result = &v1alpha1.ElasticsearchKeystore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elasticsearchkeystores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
