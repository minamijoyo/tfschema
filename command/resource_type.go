package command

import (
	"fmt"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

type ResourceTypeCommand struct {
	Meta
}

func (c *ResourceTypeCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The resource type command expects NAME")
		c.Ui.Error(c.Help())
		return 1
	}

	resourceType := args[0]
	providerName, err := detectProviderName(resourceType)
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

	err = client.GetSchema(resourceType)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func detectProviderName(resourceType string) (string, error) {
	s := strings.SplitN(resourceType, "_", 2)
	if len(s) < 2 {
		return "", fmt.Errorf("Failed to detect a provider name from the resource type: %s", resourceType)
	}
	return s[0], nil
}

func (c *ResourceTypeCommand) Help() string {
	helpText := `
Usage: tfschema resource type NAME
`
	return strings.TrimSpace(helpText)
}

func (c *ResourceTypeCommand) Synopsis() string {
	return "Get schema for a resource type"
}
