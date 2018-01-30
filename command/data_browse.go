package command

import (
	"strings"

	"github.com/pkg/browser"
	"github.com/posener/complete"
)

type DataBrowseCommand struct {
	Meta
}

func (c *DataBrowseCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The data browse command expects DATA_SOURCE")
		c.Ui.Error(c.Help())
		return 1
	}

	dataSource := args[0]
	url, err := buildDataDocURL(dataSource)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	err = browser.OpenURL(url)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func (c *DataBrowseCommand) AutocompleteArgs() complete.Predictor {
	return c.completePredictDataSource()
}

func (c *DataBrowseCommand) AutocompleteFlags() complete.Flags {
	return nil
}

func (c *DataBrowseCommand) Help() string {
	helpText := `
Usage: tfschema data browse DATA_SOURCE
`
	return strings.TrimSpace(helpText)
}

func (c *DataBrowseCommand) Synopsis() string {
	return "Browse a documentation of data source"
}
