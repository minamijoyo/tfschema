package tfschema

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/hashicorp/terraform/terraform"
)

// NetRPCClient implements Client interface.
// This implementaion is for Terraform v0.11.
type NetRPCClient struct {
	// provider is a resource provider of Terraform.
	provider terraform.ResourceProvider
	// pluginClient is a pointer to plugin client instance.
	// The type of pluginClient is
	// *github.com/hashicorp/terraform/vendor/github.com/hashicorp/go-plugin.Client.
	// But, we cannot import the vendor version of go-plugin using terraform.
	// So, we store this as interface{}, and use it by reflection.
	pluginClient interface{}
}

// GetProviderSchema returns a type definiton of provider schema.
func (c *NetRPCClient) GetProviderSchema() (*Block, error) {
	req := &terraform.ProviderSchemaRequest{
		ResourceTypes: []string{},
		DataSources:   []string{},
	}

	res, err := c.provider.GetSchema(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", err)
	}

	b := NewBlock(res.Provider)
	return b, nil
}

// GetResourceTypeSchema returns a type definiton of resource type.
func (c *NetRPCClient) GetResourceTypeSchema(resourceType string) (*Block, error) {
	req := &terraform.ProviderSchemaRequest{
		ResourceTypes: []string{resourceType},
		DataSources:   []string{},
	}

	res, err := c.provider.GetSchema(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", err)
	}

	if res.ResourceTypes[resourceType] == nil {
		return nil, fmt.Errorf("Failed to find resource type: %s", resourceType)
	}

	b := NewBlock(res.ResourceTypes[resourceType])
	return b, nil
}

// GetDataSourceSchema returns a type definiton of data source.
func (c *NetRPCClient) GetDataSourceSchema(dataSource string) (*Block, error) {
	req := &terraform.ProviderSchemaRequest{
		ResourceTypes: []string{},
		DataSources:   []string{dataSource},
	}

	res, err := c.provider.GetSchema(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", err)
	}

	if res.DataSources[dataSource] == nil {
		return nil, fmt.Errorf("Failed to find data source: %s", dataSource)
	}

	b := NewBlock(res.DataSources[dataSource])
	return b, nil
}

// ResourceTypes returns a list of resource types.
func (c *NetRPCClient) ResourceTypes() ([]string, error) {
	res := c.provider.Resources()

	keys := make([]string, 0, len(res))
	for _, r := range res {
		keys = append(keys, r.Name)
	}

	sort.Strings(keys)
	return keys, nil
}

// DataSources returns a list of data sources.
func (c *NetRPCClient) DataSources() ([]string, error) {
	res := c.provider.DataSources()

	keys := make([]string, 0, len(res))
	for _, r := range res {
		keys = append(keys, r.Name)
	}

	sort.Strings(keys)
	return keys, nil
}

// Close kills a process of the plugin.
func (c *NetRPCClient) Close() {
	// We cannot import the vendor version of go-plugin using terraform.
	// So, we call (*go-plugin.Client).Kill() by reflection here.
	v := reflect.ValueOf(c.pluginClient).MethodByName("Kill")
	v.Call([]reflect.Value{})
}
