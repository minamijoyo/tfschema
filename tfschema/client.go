package tfschema

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/providers"
	"github.com/mitchellh/go-homedir"
)

// Client represents a tfschema Client.
type Client struct {
	// provider is a provider interface of Terraform.
	provider providers.Interface
	// pluginClient is a pointer to plugin client instance.
	// The type of pluginClient is
	// *github.com/hashicorp/terraform/vendor/github.com/hashicorp/go-plugin.Client.
	// But, we cannot import the vendor version of go-plugin using terraform.
	// So, we store this as interface{}, and use it by reflection.
	pluginClient interface{}
}

// NewClient creates a new Client instance.
func NewClient(providerName string) (*Client, error) {
	// find a provider plugin
	pluginMeta, err := findPlugin("provider", providerName)
	if err != nil {
		return nil, err
	}

	// initialize a plugin client.
	pluginClient := plugin.Client(*pluginMeta)
	gRPCClient, err := pluginClient.Client()
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize plugin: %s", err)
	}

	// create a new resource provider.
	raw, err := gRPCClient.Dispense(plugin.ProviderPluginName)
	if err != nil {
		return nil, fmt.Errorf("Failed to dispense plugin: %s", err)
	}
	provider := raw.(*plugin.GRPCProvider)

	return &Client{
		provider:     provider,
		pluginClient: pluginClient,
	}, nil
}

// findPlugin finds a plugin with the name specified in the arguments.
func findPlugin(pluginType string, pluginName string) (*discovery.PluginMeta, error) {
	dirs, err := pluginDirs()
	if err != nil {
		return nil, err
	}

	pluginMetaSet := discovery.FindPlugins(pluginType, dirs).WithName(pluginName)

	// if pluginMetaSet doesn't have any pluginMeta, pluginMetaSet.Newest() will call panic.
	// so check it here.
	if pluginMetaSet.Count() > 0 {
		ret := pluginMetaSet.Newest()
		return &ret, nil
	}

	return nil, fmt.Errorf("Failed to find plugin: %s. Plugin binary was not found in any of the following directories: [%s]", pluginName, strings.Join(dirs, ", "))
}

// pluginDirs returns a list of directories to find plugin.
// This is almost the same as Terraform, but not exactly the same.
func pluginDirs() ([]string, error) {
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

// GetProviderSchema returns a type definiton of provider schema.
func (c *Client) GetProviderSchema() (*Block, error) {
	res := c.provider.GetSchema()
	if res.Diagnostics.HasErrors() {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", res.Diagnostics.Err())
	}

	b := NewBlock(res.Provider.Block)
	return b, nil
}

// GetResourceTypeSchema returns a type definiton of resource type.
func (c *Client) GetResourceTypeSchema(resourceType string) (*Block, error) {
	res := c.provider.GetSchema()
	if res.Diagnostics.HasErrors() {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", res.Diagnostics.Err())
	}

	schema, ok := res.ResourceTypes[resourceType]
	if !ok {
		return nil, fmt.Errorf("Failed to find resource type: %s", resourceType)
	}

	b := NewBlock(schema.Block)
	return b, nil
}

// GetDataSourceSchema returns a type definiton of data source.
func (c *Client) GetDataSourceSchema(dataSource string) (*Block, error) {
	res := c.provider.GetSchema()
	if res.Diagnostics.HasErrors() {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", res.Diagnostics.Err())
	}

	schema, ok := res.DataSources[dataSource]
	if !ok {
		return nil, fmt.Errorf("Failed to find data source: %s", dataSource)
	}

	b := NewBlock(schema.Block)
	return b, nil
}

// ResourceTypes returns a list of resource types.
func (c *Client) ResourceTypes() ([]string, error) {
	res := c.provider.GetSchema()
	if res.Diagnostics.HasErrors() {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", res.Diagnostics.Err())
	}

	keys := make([]string, 0, len(res.ResourceTypes))
	for k := range res.ResourceTypes {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys, nil
}

// DataSources returns a list of data sources.
func (c *Client) DataSources() ([]string, error) {
	res := c.provider.GetSchema()
	if res.Diagnostics.HasErrors() {
		return nil, fmt.Errorf("Failed to get schema from provider: %s", res.Diagnostics.Err())
	}

	keys := make([]string, 0, len(res.DataSources))
	for k := range res.DataSources {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys, nil
}

// Kill kills a process of the plugin.
func (c *Client) Kill() {
	// We cannot import the vendor version of go-plugin using terraform.
	// So, we call (*go-plugin.Client).Kill() by reflection here.
	v := reflect.ValueOf(c.pluginClient).MethodByName("Kill")
	v.Call([]reflect.Value{})
}
