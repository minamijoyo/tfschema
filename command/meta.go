package command

import (
	"fmt"
	"strings"

	"github.com/minamijoyo/tfschema/formatter"
	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/mitchellh/cli"
)

const docBaseURL = "https://www.terraform.io/docs/providers/"

// Meta are the meta-options that are available on all or most commands.
type Meta struct {
	// UI is a user interface representing input and output.
	UI cli.Ui
}

func detectProviderName(name string) (string, error) {
	s := strings.SplitN(name, "_", 2)
	if len(s) < 2 {
		return "", fmt.Errorf("Failed to detect a provider name: %s", name)
	}
	return s[0], nil
}

func buildProviderDocURL(providerName string) (string, error) {
	// build a doc URL like this
	// https://www.terraform.io/docs/providers/aws/index.html
	url := docBaseURL + providerName + "/index.html"
	return url, nil
}

func buildResourceDocURL(resourceType string) (string, error) {
	s := strings.SplitN(resourceType, "_", 2)
	if len(s) < 2 {
		return "", fmt.Errorf("Failed to build a resource doc URL: %s", resourceType)
	}

	// build a doc URL like this
	// https://www.terraform.io/docs/providers/aws/r/security_group.html
	url := docBaseURL + s[0] + "/r/" + s[1] + ".html"
	return url, nil
}

func buildDataDocURL(dataSource string) (string, error) {
	s := strings.SplitN(dataSource, "_", 2)
	if len(s) < 2 {
		return "", fmt.Errorf("Failed to build a data source doc URL: %s", dataSource)
	}

	// build a doc URL like this
	// https://www.terraform.io/docs/providers/aws/d/security_group.html
	url := docBaseURL + s[0] + "/d/" + s[1] + ".html"
	return url, nil
}

// formatBlock is a helper function for formatting tfschema.Block.
func formatBlock(b *tfschema.Block, format string) (string, error) {
	f, err := formatter.NewBlockFormatter(b, format)
	if err != nil {
		return "", err
	}
	return f.Format()
}
