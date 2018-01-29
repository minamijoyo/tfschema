package command

import (
	"encoding/json"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/posener/complete"
)

type ResourceShowCommand struct {
	Meta
}

func (c *ResourceShowCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The resource show command expects RESOURCE_TYPE")
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

	res, err := client.GetResourceTypeSchema(resourceType)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	bytes, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Output(string(bytes))

	return 0
}

func (c *ResourceShowCommand) AutocompleteArgs() complete.Predictor {
	return completePredictSequence{
		complete.PredictNothing,
		c.completePredictResourceType(),
	}
}

func (c *ResourceShowCommand) AutocompleteFlags() complete.Flags {
	return nil
}

func (c *ResourceShowCommand) Help() string {
	helpText := `
Usage: tfschema resource show RESOURCE_TYPE
`
	return strings.TrimSpace(helpText)
}

func (c *ResourceShowCommand) Synopsis() string {
	return "Show a type definition of resource"
}
