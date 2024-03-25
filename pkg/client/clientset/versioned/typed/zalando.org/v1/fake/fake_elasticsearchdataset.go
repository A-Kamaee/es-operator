/*
Copyright 2024 The Kubernetes Authors.

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

package fake

import (
	"context"

	v1 "github.com/zalando-incubator/es-operator/pkg/apis/zalando.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeElasticsearchDataSets implements ElasticsearchDataSetInterface
type FakeElasticsearchDataSets struct {
	Fake *FakeZalandoV1
	ns   string
}

var elasticsearchdatasetsResource = v1.SchemeGroupVersion.WithResource("elasticsearchdatasets")

var elasticsearchdatasetsKind = v1.SchemeGroupVersion.WithKind("ElasticsearchDataSet")

// Get takes name of the elasticsearchDataSet, and returns the corresponding elasticsearchDataSet object, and an error if there is any.
func (c *FakeElasticsearchDataSets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ElasticsearchDataSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elasticsearchdatasetsResource, c.ns, name), &v1.ElasticsearchDataSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ElasticsearchDataSet), err
}

// List takes label and field selectors, and returns the list of ElasticsearchDataSets that match those selectors.
func (c *FakeElasticsearchDataSets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ElasticsearchDataSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elasticsearchdatasetsResource, elasticsearchdatasetsKind, c.ns, opts), &v1.ElasticsearchDataSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ElasticsearchDataSetList{ListMeta: obj.(*v1.ElasticsearchDataSetList).ListMeta}
	for _, item := range obj.(*v1.ElasticsearchDataSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticsearchDataSets.
func (c *FakeElasticsearchDataSets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elasticsearchdatasetsResource, c.ns, opts))

}

// Create takes the representation of a elasticsearchDataSet and creates it.  Returns the server's representation of the elasticsearchDataSet, and an error, if there is any.
func (c *FakeElasticsearchDataSets) Create(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.CreateOptions) (result *v1.ElasticsearchDataSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elasticsearchdatasetsResource, c.ns, elasticsearchDataSet), &v1.ElasticsearchDataSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ElasticsearchDataSet), err
}

// Update takes the representation of a elasticsearchDataSet and updates it. Returns the server's representation of the elasticsearchDataSet, and an error, if there is any.
func (c *FakeElasticsearchDataSets) Update(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (result *v1.ElasticsearchDataSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elasticsearchdatasetsResource, c.ns, elasticsearchDataSet), &v1.ElasticsearchDataSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ElasticsearchDataSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticsearchDataSets) UpdateStatus(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (*v1.ElasticsearchDataSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elasticsearchdatasetsResource, "status", c.ns, elasticsearchDataSet), &v1.ElasticsearchDataSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ElasticsearchDataSet), err
}

// Delete takes name of the elasticsearchDataSet and deletes it. Returns an error if one occurs.
func (c *FakeElasticsearchDataSets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(elasticsearchdatasetsResource, c.ns, name, opts), &v1.ElasticsearchDataSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticsearchDataSets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elasticsearchdatasetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ElasticsearchDataSetList{})
	return err
}

// Patch applies the patch and returns the patched elasticsearchDataSet.
func (c *FakeElasticsearchDataSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ElasticsearchDataSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elasticsearchdatasetsResource, c.ns, name, pt, data, subresources...), &v1.ElasticsearchDataSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ElasticsearchDataSet), err
}
