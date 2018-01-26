package command

import (
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

type DataListCommand struct {
	Meta
}

func (c *DataListCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The data list command expects PROVIDER")
		c.Ui.Error(c.Help())
		return 1
	}

	providerName := args[0]

	client, err := tfschema.NewClient(providerName)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	defer client.Kill()

	dataSources := client.DataSources()
	c.Ui.Output(strings.Join(dataSources, "\n"))

	return 0
}

func (c *DataListCommand) Help() string {
	helpText := `
Usage: tfschema data list PROVIDER
`
	return strings.TrimSpace(helpText)
}

func (c *DataListCommand) Synopsis() string {
	return "List data sources"
}
