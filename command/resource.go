package command

import (
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
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

	client, err := tfschema.NewClient("aws")
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	err = client.GetSchema(resourceType)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

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
