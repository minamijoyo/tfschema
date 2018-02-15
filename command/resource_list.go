package command

import (
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

// ResourceListCommand is a command which lists resource types.
type ResourceListCommand struct {
	Meta
}

// Run runs the procedure of this command.
func (c *ResourceListCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("The resource list command expects PROVIDER")
		c.UI.Error(c.Help())
		return 1
	}

	providerName := args[0]

	client, err := tfschema.NewClient(providerName)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	defer client.Kill()

	res := client.Resources()

	resourceTypes := []string{}
	for _, r := range res {
		resourceTypes = append(resourceTypes, r.Name)
	}

	c.UI.Output(strings.Join(resourceTypes, "\n"))

	return 0
}

// Help returns long-form help text.
func (c *ResourceListCommand) Help() string {
	helpText := `
Usage: tfschema resource list PROVIDER
`
	return strings.TrimSpace(helpText)
}

// Synopsis returns one-line help text.
func (c *ResourceListCommand) Synopsis() string {
	return "List resource types"
}
