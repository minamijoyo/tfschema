package command

import (
	"strings"

	"github.com/pkg/browser"
	"github.com/posener/complete"
)

type ResourceBrowseCommand struct {
	Meta
}

func (c *ResourceBrowseCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("The resource browse command expects RESOURCE_TYPE")
		c.UI.Error(c.Help())
		return 1
	}

	resourceType := args[0]
	url, err := buildResourceDocURL(resourceType)
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

func (c *ResourceBrowseCommand) AutocompleteArgs() complete.Predictor {
	return c.completePredictResourceType()
}

func (c *ResourceBrowseCommand) AutocompleteFlags() complete.Flags {
	return nil
}

func (c *ResourceBrowseCommand) Help() string {
	helpText := `
Usage: tfschema resource browse RESOURCE_TYPE
`
	return strings.TrimSpace(helpText)
}

func (c *ResourceBrowseCommand) Synopsis() string {
	return "Browse a documentation of resource"
}
