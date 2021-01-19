/*
Copyright 2021 The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/zalando-incubator/es-operator/pkg/apis/zalando.org/v1"
	scheme "github.com/zalando-incubator/es-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ElasticsearchDataSetsGetter has a method to return a ElasticsearchDataSetInterface.
// A group's client should implement this interface.
type ElasticsearchDataSetsGetter interface {
	ElasticsearchDataSets(namespace string) ElasticsearchDataSetInterface
}

// ElasticsearchDataSetInterface has methods to work with ElasticsearchDataSet resources.
type ElasticsearchDataSetInterface interface {
	Create(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.CreateOptions) (*v1.ElasticsearchDataSet, error)
	Update(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (*v1.ElasticsearchDataSet, error)
	UpdateStatus(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (*v1.ElasticsearchDataSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ElasticsearchDataSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ElasticsearchDataSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ElasticsearchDataSet, err error)
	ElasticsearchDataSetExpansion
}

// elasticsearchDataSets implements ElasticsearchDataSetInterface
type elasticsearchDataSets struct {
	client rest.Interface
	ns     string
}

// newElasticsearchDataSets returns a ElasticsearchDataSets
func newElasticsearchDataSets(c *ZalandoV1Client, namespace string) *elasticsearchDataSets {
	return &elasticsearchDataSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the elasticsearchDataSet, and returns the corresponding elasticsearchDataSet object, and an error if there is any.
func (c *elasticsearchDataSets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ElasticsearchDataSet, err error) {
	result = &v1.ElasticsearchDataSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElasticsearchDataSets that match those selectors.
func (c *elasticsearchDataSets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ElasticsearchDataSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ElasticsearchDataSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elasticsearchDataSets.
func (c *elasticsearchDataSets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a elasticsearchDataSet and creates it.  Returns the server's representation of the elasticsearchDataSet, and an error, if there is any.
func (c *elasticsearchDataSets) Create(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.CreateOptions) (result *v1.ElasticsearchDataSet, err error) {
	result = &v1.ElasticsearchDataSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticsearchDataSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a elasticsearchDataSet and updates it. Returns the server's representation of the elasticsearchDataSet, and an error, if there is any.
func (c *elasticsearchDataSets) Update(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (result *v1.ElasticsearchDataSet, err error) {
	result = &v1.ElasticsearchDataSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		Name(elasticsearchDataSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticsearchDataSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *elasticsearchDataSets) UpdateStatus(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (result *v1.ElasticsearchDataSet, err error) {
	result = &v1.ElasticsearchDataSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		Name(elasticsearchDataSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticsearchDataSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the elasticsearchDataSet and deletes it. Returns an error if one occurs.
func (c *elasticsearchDataSets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elasticsearchDataSets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched elasticsearchDataSet.
func (c *elasticsearchDataSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ElasticsearchDataSet, err error) {
	result = &v1.ElasticsearchDataSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elasticsearchdatasets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
