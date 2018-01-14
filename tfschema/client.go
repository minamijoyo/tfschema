package tfschema

import (
	"fmt"
	"go/build"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/terraform"
)

type Client struct {
	provider terraform.ResourceProvider
}

func NewClient(providerName string) (*Client, error) {
	provider, err := newProvider(providerName)
	if err != nil {
		return nil, err
	}

	return &Client{
		provider: provider,
	}, nil
}

func newProvider(name string) (terraform.ResourceProvider, error) {
	pluginMeta := findPlugin("provider", name)

	client := plugin.Client(pluginMeta)
	rpcClient, err := client.Client()
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize plugin: %s", err)
	}

	raw, err := rpcClient.Dispense(plugin.ProviderPluginName)
	if err != nil {
		return nil, fmt.Errorf("Failed to dispense plugin: %s", err)
	}
	provider := raw.(terraform.ResourceProvider)

	return provider, nil
}

func findPlugin(pluginType string, pluginName string) discovery.PluginMeta {
	pluginMetaSet := discovery.FindPlugins(pluginType, pluginDirs())

	plugins := make(map[string]discovery.PluginMeta)
	for plugin := range pluginMetaSet {
		name := plugin.Name
		plugins[name] = plugin
	}

	return plugins[pluginName]
}

func pluginDirs() []string {
	gopath := build.Default.GOPATH
	pluginDirs := []string{gopath + "/bin"}
	return pluginDirs
}

func (c *Client) GetSchema(resourceType string) error {
	req := &terraform.ProviderSchemaRequest{
		ResourceTypes: []string{resourceType},
		DataSources:   []string{},
	}

	res, err := c.provider.GetSchema(req)
	if err != nil {
		return fmt.Errorf("Faild to get schema from provider: %s", err)
	}

	spew.Dump(res.ResourceTypes)

	return nil
}
