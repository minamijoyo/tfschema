package command

import (
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

type DataShowCommand struct {
	Meta
}

func (c *DataShowCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The data show command expects DATA_SOURCE")
		c.Ui.Error(c.Help())
		return 1
	}

	dataSource := args[0]
	providerName, err := detectProviderName(dataSource)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	client, err := tfschema.NewClient(providerName)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	defer client.Kill()

	output, err := client.GetDataSourceSchema(dataSource)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Output(output)

	return 0
}

func (c *DataShowCommand) Help() string {
	helpText := `
Usage: tfschema data show DATA_SOURCE
`
	return strings.TrimSpace(helpText)
}

func (c *DataShowCommand) Synopsis() string {
	return "Show a type definition of data source"
}
