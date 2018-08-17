package command

import (
	"flag"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
)

// ProviderShowCommand is a command which shows a type definition of provider.
type ProviderShowCommand struct {
	Meta
	format string
}

// Run runs the procedure of this command.
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

	out, err := formatBlock(block, c.format)

	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(out)

	return 0
}

// Help returns long-form help text.
func (c *ProviderShowCommand) Help() string {
	helpText := `
Usage: tfschema provider show [options] PROVIDER

Options:

  -format=type    Set output format to table or json (default: table)
`
	return strings.TrimSpace(helpText)
}

// Synopsis returns one-line help text.
func (c *ProviderShowCommand) Synopsis() string {
	return "Show a type definition of provider"
}
