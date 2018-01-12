package command

import (
	"fmt"
	"go/build"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/terraform"
)

type ResourceShowCommand struct {
	Meta
}

func (c *ResourceShowCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The resource show command expects RESOURCE_TYPE.")
		c.Ui.Error(c.Help())
		return 1
	}

	resourceType := args[0]

	// find provider plugins
	gopath := build.Default.GOPATH
	pluginDirs := []string{gopath + "/bin"}
	pluginMetaSet := discovery.FindPlugins("provider", pluginDirs)

	plugins := make(map[string]discovery.PluginMeta)
	for plugin := range pluginMetaSet {
		name := plugin.Name
		plugins[name] = plugin
	}

	// initialize aws plugin
	client := plugin.Client(plugins["aws"])
	rpcClient, err := client.Client()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to initialize plugin: %s", err))
		return 1
	}

	raw, err := rpcClient.Dispense(plugin.ProviderPluginName)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to dispense plugin: %s", err))
		return 1
	}
	provider := raw.(terraform.ResourceProvider)

	// invoke GetSchema
	req := &terraform.ProviderSchemaRequest{
		ResourceTypes: []string{resourceType},
		DataSources:   []string{},
	}
	res, err := provider.GetSchema(req)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to get schema from provider: %s", err))
		return 1
	}

	spew.Dump(res)

	return 0
}

func (c *ResourceShowCommand) Help() string {
	helpText := `
Usage: tfschema resource show RESOURCE_TYPE
`
	return strings.TrimSpace(helpText)
}

func (c *ResourceShowCommand) Synopsis() string {
	return "Show a resource type in the schema"
}
