package main

import (
	"fmt"
	"go/build"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	// find provider plugins
	gopath := build.Default.GOPATH
	pluginDirs := []string{gopath + "/bin"}
	pluginMetaSet := discovery.FindPlugins("provider", pluginDirs)
	spew.Dump(pluginMetaSet)

	plugins := make(map[string]discovery.PluginMeta)
	for plugin := range pluginMetaSet {
		name := plugin.Name
		plugins[name] = plugin
	}

	// initialize aws plugin
	client := plugin.Client(plugins["aws"])
	rpcClient, err := client.Client()
	if err != nil {
		panic(err)
	}

	raw, err := rpcClient.Dispense(plugin.ProviderPluginName)
	if err != nil {
		panic(err)
	}
	provider := raw.(terraform.ResourceProvider)

	// invoke GetSchema
	req := &terraform.ProviderSchemaRequest{
		ResourceTypes: []string{"aws_security_group"},
		DataSources:   []string{},
	}
	res, err := provider.GetSchema(req)

	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	spew.Dump(res)
}
