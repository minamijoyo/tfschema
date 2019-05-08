package tfschema

import (
	"fmt"
	"sort"

	"github.com/hashicorp/terraform/providers"
)

// GRPCClient implements Client interface.
// This implementaion is for Terraform v0.12+.
type GRPCClient struct {
	// provider is a provider interface of Terraform.
	provider providers.Interface
}

// getSchema is a helper function to get a schema from provider.
func (c *GRPCClient) getSchema() (providers.GetSchemaResponse, error) {
	res := c.provider.GetSchema()
	if res.Diagnostics.HasErrors() {
		return res, fmt.Errorf("Failed to get schema from provider: %s", res.Diagnostics.Err())
	}

	return res, nil
}

// GetProviderSchema returns a type definiton of provider schema.
func (c *GRPCClient) GetProviderSchema() (*Block, error) {
	res, err := c.getSchema()
	if err != nil {
		return nil, err
	}

	b := NewBlock(res.Provider.Block)
	return b, nil
}

// GetResourceTypeSchema returns a type definiton of resource type.
func (c *GRPCClient) GetResourceTypeSchema(resourceType string) (*Block, error) {
	res, err := c.getSchema()
	if err != nil {
		return nil, err
	}

	schema, ok := res.ResourceTypes[resourceType]
	if !ok {
		return nil, fmt.Errorf("Failed to find resource type: %s", resourceType)
	}

	b := NewBlock(schema.Block)
	return b, nil
}

// GetDataSourceSchema returns a type definiton of data source.
func (c *GRPCClient) GetDataSourceSchema(dataSource string) (*Block, error) {
	res, err := c.getSchema()
	if err != nil {
		return nil, err
	}

	schema, ok := res.DataSources[dataSource]
	if !ok {
		return nil, fmt.Errorf("Failed to find data source: %s", dataSource)
	}

	b := NewBlock(schema.Block)
	return b, nil
}

// ResourceTypes returns a list of resource types.
func (c *GRPCClient) ResourceTypes() ([]string, error) {
	res, err := c.getSchema()
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(res.ResourceTypes))
	for k := range res.ResourceTypes {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys, nil
}

// DataSources returns a list of data sources.
func (c *GRPCClient) DataSources() ([]string, error) {
	res, err := c.getSchema()
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(res.DataSources))
	for k := range res.DataSources {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys, nil
}

// Close closes a connection and kills a process of the plugin.
func (c *GRPCClient) Close() {
	c.provider.Close()
}
