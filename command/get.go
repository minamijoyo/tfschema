package command

import (
	"fmt"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

type GetCommand struct {
	Meta
}

func (c *GetCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The get command expects RESOURCE_TYPE.")
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

func (c *GetCommand) Help() string {
	helpText := `
Usage: tfschema get RESOURCE_TYPE
`
	return strings.TrimSpace(helpText)
}

func (c *GetCommand) Synopsis() string {
	return "get schema for a resource type"
}
