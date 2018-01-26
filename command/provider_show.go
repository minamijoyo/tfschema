package command

import (
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

type ProviderShowCommand struct {
	Meta
}

func (c *ProviderShowCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The provider show command expects PROVIDER")
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

	output, err := client.GetProviderSchema()
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Output(output)

	return 0
}

func (c *ProviderShowCommand) Help() string {
	helpText := `
Usage: tfschema provider show PROVIDER
`
	return strings.TrimSpace(helpText)
}

func (c *ProviderShowCommand) Synopsis() string {
	return "Show a type definition of provider"
}
