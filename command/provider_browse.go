package command

import (
	"strings"

	"github.com/pkg/browser"
)

// ProviderBrowseCommand is a command which browses a documentation of provider.
type ProviderBrowseCommand struct {
	Meta
}

// Run runs the procedure of this command.
func (c *ProviderBrowseCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("The provider browse command expects PROVIDER")
		c.UI.Error(c.Help())
		return 1
	}

	providerName := args[0]
	url, err := buildProviderDocURL(providerName)
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

// Help returns long-form help text.
func (c *ProviderBrowseCommand) Help() string {
	helpText := `
Usage: tfschema provider browse PROVIDER
`
	return strings.TrimSpace(helpText)
}

// Synopsis returns one-line help text.
func (c *ProviderBrowseCommand) Synopsis() string {
	return "Browse a documentation of provider"
}
