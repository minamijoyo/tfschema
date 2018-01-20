package tfschema

import (
	"fmt"
	"go/build"
	"reflect"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/terraform"
	"github.com/k0kubun/pp"
)

type Client struct {
	provider terraform.ResourceProvider
	// The type of pluginClient is
	// *github.com/hashicorp/terraform/vendor/github.com/hashicorp/go-plugin.Client.
	// But, we cannot import the vendor version of go-plugin using terraform.
	// So, we store this as interface{}, and use it by reflection.
	pluginClient interface{}
}

func NewClient(providerName string) (*Client, error) {
	pluginMeta := findPlugin("provider", providerName)

	pluginClient := plugin.Client(pluginMeta)
	rpcClient, err := pluginClient.Client()
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize plugin: %s", err)
	}

	raw, err := rpcClient.Dispense(plugin.ProviderPluginName)
	if err != nil {
		return nil, fmt.Errorf("Failed to dispense plugin: %s", err)
	}
	provider := raw.(terraform.ResourceProvider)

	return &Client{
		provider:     provider,
		pluginClient: pluginClient,
	}, nil
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

	pp.Println(res.ResourceTypes)

	return nil
}

func (c *Client) Kill() {
	v := reflect.ValueOf(c.pluginClient).MethodByName("Kill")
	v.Call([]reflect.Value{})
}
