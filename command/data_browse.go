package command

import (
	"strings"

	"github.com/pkg/browser"
	"github.com/posener/complete"
)

// DataBrowseCommand is a command which browses a documentation of data source.
type DataBrowseCommand struct {
	Meta
}

// Run runs the procedure of this command.
func (c *DataBrowseCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("The data browse command expects DATA_SOURCE")
		c.UI.Error(c.Help())
		return 1
	}

	dataSource := args[0]
	url, err := buildDataDocURL(dataSource)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	err = browser.OpenURL(url)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	return 0
}

// AutocompleteArgs returns the argument predictor.
func (c *DataBrowseCommand) AutocompleteArgs() complete.Predictor {
	return c.completePredictDataSource()
}

// AutocompleteFlags returns a mapping of supported flags and options.
func (c *DataBrowseCommand) AutocompleteFlags() complete.Flags {
	return nil
}

// Help returns long-form help text.
func (c *DataBrowseCommand) Help() string {
	helpText := `
Usage: tfschema data browse DATA_SOURCE
`
	return strings.TrimSpace(helpText)
}

// Synopsis returns one-line help text.
func (c *DataBrowseCommand) Synopsis() string {
	return "Browse a documentation of data source"
}
