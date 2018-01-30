package command

import (
	"strings"

	"github.com/pkg/browser"
)

type ProviderBrowseCommand struct {
	Meta
}

func (c *ProviderBrowseCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("The provider browse command expects PROVIDER")
		c.Ui.Error(c.Help())
		return 1
	}

	providerName := args[0]
	url, err := buildProviderDocURL(providerName)
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

func (c *ProviderBrowseCommand) Help() string {
	helpText := `
Usage: tfschema provider browse PROVIDER
`
	return strings.TrimSpace(helpText)
}

func (c *ProviderBrowseCommand) Synopsis() string {
	return "Browse a documentation of provider"
}
