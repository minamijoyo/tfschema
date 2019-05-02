package command

import (
	"flag"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/posener/complete"
)

// ResourceShowCommand is a command which shows a type definition of resource.
type ResourceShowCommand struct {
	Meta
	format string
}

// Run runs the procedure of this command.
func (c *ResourceShowCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("resource show", flag.ContinueOnError)
	cmdFlags.StringVar(&c.format, "format", "table", "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if len(cmdFlags.Args()) != 1 {
		c.UI.Error("The resource show command expects RESOURCE_TYPE")
		c.UI.Error(c.Help())
		return 1
	}

	resourceType := cmdFlags.Args()[0]
	providerName, err := detectProviderName(resourceType)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	client, err := tfschema.NewClient(providerName)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	defer client.Close()

	block, err := client.GetResourceTypeSchema(resourceType)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	out, err := formatBlock(block, c.format)

	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(out)

	return 0
}

// AutocompleteArgs returns the argument predictor.
func (c *ResourceShowCommand) AutocompleteArgs() complete.Predictor {
	return c.completePredictResourceType()
}

// AutocompleteFlags returns a mapping of supported flags and options.
func (c *ResourceShowCommand) AutocompleteFlags() complete.Flags {
	return nil
}

// Help returns long-form help text.
func (c *ResourceShowCommand) Help() string {
	helpText := `
Usage: tfschema resource show [options] RESOURCE_TYPE

Options:

  -format=type    Set output format to table or json (default: table)
`
	return strings.TrimSpace(helpText)
}

// Synopsis returns one-line help text.
func (c *ResourceShowCommand) Synopsis() string {
	return "Show a type definition of resource"
}
