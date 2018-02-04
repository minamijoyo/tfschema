package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/posener/complete"
)

type DataShowCommand struct {
	Meta
	format string
}

func (c *DataShowCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("data show", flag.ContinueOnError)
	cmdFlags.StringVar(&c.format, "format", "table", "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if len(cmdFlags.Args()) != 1 {
		c.Ui.Error("The data show command expects DATA_SOURCE")
		c.Ui.Error(c.Help())
		return 1
	}

	dataSource := cmdFlags.Args()[0]
	providerName, err := detectProviderName(dataSource)
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

	block, err := client.GetDataSourceSchema(dataSource)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	var out string
	switch c.format {
	case "table":
		out, err = block.FormatTable()
	case "json":
		out, err = block.FormatJSON()
	default:
		c.Ui.Error(fmt.Sprintf("Unknown output format: %s", c.format))
		c.Ui.Error(c.Help())
		return 1
	}

	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Output(out)

	return 0
}

func (c *DataShowCommand) AutocompleteArgs() complete.Predictor {
	return c.completePredictDataSource()
}

func (c *DataShowCommand) AutocompleteFlags() complete.Flags {
	return nil
}

func (c *DataShowCommand) Help() string {
	helpText := `
Usage: tfschema data show [options] DATA_SOURCE

Options:

  -format=type    Set output format to table or json (default: table)
`
	return strings.TrimSpace(helpText)
}

func (c *DataShowCommand) Synopsis() string {
	return "Show a type definition of data source"
}
