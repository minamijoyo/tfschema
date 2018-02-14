package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

type ProviderShowCommand struct {
	Meta
	format string
}

func (c *ProviderShowCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("provider show", flag.ContinueOnError)
	cmdFlags.StringVar(&c.format, "format", "table", "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if len(cmdFlags.Args()) != 1 {
		c.UI.Error("The provider show command expects PROVIDER")
		c.UI.Error(c.Help())
		return 1
	}

	providerName := cmdFlags.Args()[0]

	client, err := tfschema.NewClient(providerName)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	defer client.Kill()

	block, err := client.GetProviderSchema()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	var out string
	switch c.format {
	case "table":
		out, err = block.FormatTable()
	case "json":
		out, err = block.FormatJSON()
	default:
		c.UI.Error(fmt.Sprintf("Unknown output format: %s", c.format))
		c.UI.Error(c.Help())
		return 1
	}

	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(out)

	return 0
}

func (c *ProviderShowCommand) Help() string {
	helpText := `
Usage: tfschema provider show [options] PROVIDER

Options:

  -format=type    Set output format to table or json (default: table)
`
	return strings.TrimSpace(helpText)
}

func (c *ProviderShowCommand) Synopsis() string {
	return "Show a type definition of provider"
}
