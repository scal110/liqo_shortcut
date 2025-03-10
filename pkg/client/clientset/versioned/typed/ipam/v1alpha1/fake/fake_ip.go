// Copyright 2019-2025 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/liqotech/liqo/apis/ipam/v1alpha1"
)

// FakeIPs implements IPInterface
type FakeIPs struct {
	Fake *FakeIpamV1alpha1
	ns   string
}

var ipsResource = v1alpha1.SchemeGroupVersion.WithResource("ips")

var ipsKind = v1alpha1.SchemeGroupVersion.WithKind("IP")

// Get takes name of the iP, and returns the corresponding iP object, and an error if there is any.
func (c *FakeIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(ipsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// List takes label and field selectors, and returns the list of IPs that match those selectors.
func (c *FakeIPs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IPList, err error) {
	emptyResult := &v1alpha1.IPList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(ipsResource, ipsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IPList{ListMeta: obj.(*v1alpha1.IPList).ListMeta}
	for _, item := range obj.(*v1alpha1.IPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPs.
func (c *FakeIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(ipsResource, c.ns, opts))

}

// Create takes the representation of a iP and creates it.  Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Create(ctx context.Context, iP *v1alpha1.IP, opts v1.CreateOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(ipsResource, c.ns, iP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// Update takes the representation of a iP and updates it. Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Update(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(ipsResource, c.ns, iP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIPs) UpdateStatus(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(ipsResource, "status", c.ns, iP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// Delete takes name of the iP and deletes it. Returns an error if one occurs.
func (c *FakeIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ipsResource, c.ns, name, opts), &v1alpha1.IP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(ipsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IPList{})
	return err
}

// Patch applies the patch and returns the patched iP.
func (c *FakeIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(ipsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}
