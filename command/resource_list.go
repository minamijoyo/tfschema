package command

import (
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

type ResourceListCommand struct {
	Meta
}

func (c *ResourceListCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The resource list command expects PROVIDER")
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

	res := client.Resources()

	resourceTypes := []string{}
	for _, r := range res {
		resourceTypes = append(resourceTypes, r.Name)
	}

	c.Ui.Output(strings.Join(resourceTypes, "\n"))

	return 0
}

func (c *ResourceListCommand) Help() string {
	helpText := `
Usage: tfschema resource list PROVIDER
`
	return strings.TrimSpace(helpText)
}

func (c *ResourceListCommand) Synopsis() string {
	return "List resource types"
}
