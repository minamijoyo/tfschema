package command

import (
	"encoding/json"
	"strings"

	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/posener/complete"
)

type DataShowCommand struct {
	Meta
}

func (c *DataShowCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The data show command expects DATA_SOURCE")
		c.Ui.Error(c.Help())
		return 1
	}

	dataSource := args[0]
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

	res, err := client.GetDataSourceSchema(dataSource)
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

func (c *DataShowCommand) AutocompleteArgs() complete.Predictor {
	return completePredictSequence{
		complete.PredictNothing,
		c.completePredictDataSource(),
	}
}

func (c *DataShowCommand) AutocompleteFlags() complete.Flags {
	return nil
}

func (c *DataShowCommand) Help() string {
	helpText := `
Usage: tfschema data show DATA_SOURCE
`
	return strings.TrimSpace(helpText)
}

func (c *DataShowCommand) Synopsis() string {
	return "Show a type definition of data source"
}
