package tfschema

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/hashicorp/terraform/config/configschema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mitchellh/go-homedir"
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
	pluginMeta, err := findPlugin("provider", providerName)
	if err != nil {
		return nil, err
	}

	pluginClient := plugin.Client(*pluginMeta)
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

func findPlugin(pluginType string, pluginName string) (*discovery.PluginMeta, error) {
	dirs, err := pluginDirs()
	if err != nil {
		return nil, err
	}

	pluginMetaSet := discovery.FindPlugins(pluginType, dirs)

	for plugin := range pluginMetaSet {
		if plugin.Name == pluginName {
			return &plugin, nil
		}
	}

	return nil, fmt.Errorf("Failed to find plugin: %s", pluginName)
}

func pluginDirs() ([]string, error) {
	// build dirs for finding plugin
	// This is almost the same as Terraform, but not exactly the same.
	dirs := []string{}

	// current directory
	dirs = append(dirs, ".")

	// same directory as this executable
	exePath, err := os.Executable()
	if err != nil {
		return []string{}, fmt.Errorf("Failed to get executable path: %s", err)
	}
	dirs = append(dirs, filepath.Dir(exePath))

	// user vendor directory
	arch := runtime.GOOS + "_" + runtime.GOARCH
	vendorDir := filepath.Join("terraform.d", "plugins", arch)
	dirs = append(dirs, vendorDir)

	// auto installed directory
	// This does not take into account overriding the data directory.
	autoInstalledDir := filepath.Join(".terraform", "plugins", arch)
	dirs = append(dirs, autoInstalledDir)

	// global plugin directory
	homeDir, err := homedir.Dir()
	if err != nil {
		return []string{}, fmt.Errorf("Failed to get home dir: %s", err)
	}
	configDir := filepath.Join(homeDir, ".terraform.d", "plugins")
	dirs = append(dirs, configDir)
	dirs = append(dirs, filepath.Join(configDir, arch))

	// GOPATH
	// This is not included in the Terraform, but for convenience.
	gopath := build.Default.GOPATH
	dirs = append(dirs, filepath.Join(gopath, "bin"))

	log.Printf("[DEBUG] plugin dirs: %#v", dirs)
	return dirs, nil
}

func (c *Client) GetProviderSchema() (*configschema.Block, error) {
	req := &terraform.ProviderSchemaRequest{
		ResourceTypes: []string{},
		DataSources:   []string{},
	}

	res, err := c.provider.GetSchema(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", err)
	}

	return res.Provider, nil
}

func (c *Client) GetResourceTypeSchema(resourceType string) (*configschema.Block, error) {
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

	return res.ResourceTypes[resourceType], nil
}

func (c *Client) GetDataSourceSchema(dataSource string) (*configschema.Block, error) {
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

	return res.DataSources[dataSource], nil
}

func (c *Client) Resources() []terraform.ResourceType {
	return c.provider.Resources()
}

func (c *Client) DataSources() []terraform.DataSource {
	return c.provider.DataSources()
}

func (c *Client) Kill() {
	v := reflect.ValueOf(c.pluginClient).MethodByName("Kill")
	v.Call([]reflect.Value{})
}
