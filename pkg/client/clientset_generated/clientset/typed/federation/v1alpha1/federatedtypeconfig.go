/*
Copyright 2018 The Kubernetes Authors.

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
package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation/v1alpha1"
	scheme "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedTypeConfigsGetter has a method to return a FederatedTypeConfigInterface.
// A group's client should implement this interface.
type FederatedTypeConfigsGetter interface {
	FederatedTypeConfigs() FederatedTypeConfigInterface
}

// FederatedTypeConfigInterface has methods to work with FederatedTypeConfig resources.
type FederatedTypeConfigInterface interface {
	Create(*v1alpha1.FederatedTypeConfig) (*v1alpha1.FederatedTypeConfig, error)
	Update(*v1alpha1.FederatedTypeConfig) (*v1alpha1.FederatedTypeConfig, error)
	UpdateStatus(*v1alpha1.FederatedTypeConfig) (*v1alpha1.FederatedTypeConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FederatedTypeConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.FederatedTypeConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedTypeConfig, err error)
	FederatedTypeConfigExpansion
}

// federatedTypeConfigs implements FederatedTypeConfigInterface
type federatedTypeConfigs struct {
	client rest.Interface
}

// newFederatedTypeConfigs returns a FederatedTypeConfigs
func newFederatedTypeConfigs(c *FederationV1alpha1Client) *federatedTypeConfigs {
	return &federatedTypeConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the federatedTypeConfig, and returns the corresponding federatedTypeConfig object, and an error if there is any.
func (c *federatedTypeConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedTypeConfig, err error) {
	result = &v1alpha1.FederatedTypeConfig{}
	err = c.client.Get().
		Resource("federatedtypeconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedTypeConfigs that match those selectors.
func (c *federatedTypeConfigs) List(opts v1.ListOptions) (result *v1alpha1.FederatedTypeConfigList, err error) {
	result = &v1alpha1.FederatedTypeConfigList{}
	err = c.client.Get().
		Resource("federatedtypeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedTypeConfigs.
func (c *federatedTypeConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("federatedtypeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedTypeConfig and creates it.  Returns the server's representation of the federatedTypeConfig, and an error, if there is any.
func (c *federatedTypeConfigs) Create(federatedTypeConfig *v1alpha1.FederatedTypeConfig) (result *v1alpha1.FederatedTypeConfig, err error) {
	result = &v1alpha1.FederatedTypeConfig{}
	err = c.client.Post().
		Resource("federatedtypeconfigs").
		Body(federatedTypeConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedTypeConfig and updates it. Returns the server's representation of the federatedTypeConfig, and an error, if there is any.
func (c *federatedTypeConfigs) Update(federatedTypeConfig *v1alpha1.FederatedTypeConfig) (result *v1alpha1.FederatedTypeConfig, err error) {
	result = &v1alpha1.FederatedTypeConfig{}
	err = c.client.Put().
		Resource("federatedtypeconfigs").
		Name(federatedTypeConfig.Name).
		Body(federatedTypeConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedTypeConfigs) UpdateStatus(federatedTypeConfig *v1alpha1.FederatedTypeConfig) (result *v1alpha1.FederatedTypeConfig, err error) {
	result = &v1alpha1.FederatedTypeConfig{}
	err = c.client.Put().
		Resource("federatedtypeconfigs").
		Name(federatedTypeConfig.Name).
		SubResource("status").
		Body(federatedTypeConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedTypeConfig and deletes it. Returns an error if one occurs.
func (c *federatedTypeConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("federatedtypeconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedTypeConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("federatedtypeconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedTypeConfig.
func (c *federatedTypeConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedTypeConfig, err error) {
	result = &v1alpha1.FederatedTypeConfig{}
	err = c.client.Patch(pt).
		Resource("federatedtypeconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
