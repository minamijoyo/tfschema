package tfschema

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	plugin "github.com/hashicorp/go-plugin"
	tfplugin "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/mitchellh/go-homedir"
)

// Client represents a set of methods required to get resource type definitons
// from Terraform providers.
// Terraform v0.12+ has a different provider interface from v0.11.
// This is a compatibility layer for Terraform v0.11/v0.12+.
type Client interface {
	// GetProviderSchema returns a type definiton of provider schema.
	GetProviderSchema() (*Block, error)

	// GetResourceTypeSchema returns a type definiton of resource type.
	GetResourceTypeSchema(resourceType string) (*Block, error)

	// GetDataSourceSchema returns a type definiton of data source.
	GetDataSourceSchema(dataSource string) (*Block, error)

	// ResourceTypes returns a list of resource types.
	ResourceTypes() ([]string, error)

	// DataSources returns a list of data sources.
	DataSources() ([]string, error)

	// Close closes a connection and kills a process of the plugin.
	Close()
}

// NewClient creates a new Client instance.
func NewClient(providerName string) (Client, error) {
	// find a provider plugin
	pluginMeta, err := findPlugin("provider", providerName)
	if err != nil {
		return nil, err
	}

	// create a plugin client config
	config := newClientConfig(pluginMeta)

	// initialize a plugin client.
	pluginClient := plugin.NewClient(config)
	client, err := pluginClient.Client()
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize plugin: %s", err)
	}

	// create a new resource provider.
	raw, err := client.Dispense(tfplugin.ProviderPluginName)
	if err != nil {
		return nil, fmt.Errorf("Failed to dispense plugin: %s", err)
	}

	switch provider := raw.(type) {
	// For Terraform v0.11
	case *tfplugin.ResourceProvider:
		return &NetRPCClient{
			provider:     provider,
			pluginClient: pluginClient,
		}, nil

	// For Terraform v0.12+
	case *tfplugin.GRPCProvider:
		// To clean up the plugin process, we need to explicitly store references.
		provider.PluginClient = pluginClient

		return &GRPCClient{
			provider: provider,
		}, nil

	default:
		return nil, fmt.Errorf("Failed to type cast. unknown provider type: %+v", raw)
	}
}

// newClientConfig returns a plugin client config.
func newClientConfig(pluginMeta *discovery.PluginMeta) *plugin.ClientConfig {
	// create a default plugin client config.
	config := tfplugin.ClientConfig(*pluginMeta)
	// override config to dual protocols suport
	config.AllowedProtocols = []plugin.Protocol{
		plugin.ProtocolNetRPC,
		plugin.ProtocolGRPC,
	}
	config.VersionedPlugins = map[int]plugin.PluginSet{
		4: {
			"provider":    &tfplugin.ResourceProviderPlugin{},
			"provisioner": &tfplugin.ResourceProvisionerPlugin{},
		},
		5: {
			"provider":    &tfplugin.GRPCProviderPlugin{},
			"provisioner": &tfplugin.GRPCProvisionerPlugin{},
		},
	}

	return config
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
