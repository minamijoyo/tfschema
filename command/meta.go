package command

import (
	"fmt"
	"strings"

	"github.com/minamijoyo/tfschema/formatter"
	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/mitchellh/cli"
)

const docBaseURL = "https://registry.terraform.io/providers/hashicorp/"
const latestDocs = "/latest/docs"

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
	// https://registry.terraform.io/providers/hashicorp/aws/latest/docs
	url := docBaseURL + providerName + latestDocs
	return url, nil
}

func buildResourceDocURL(resourceType string) (string, error) {
	s := strings.SplitN(resourceType, "_", 2)
	if len(s) < 2 {
		return "", fmt.Errorf("Failed to build a resource doc URL: %s", resourceType)
	}

	// build a doc URL like this
	// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group
	url := docBaseURL + s[0] + latestDocs + "/resources/" + s[1]
	return url, nil
}

func buildDataDocURL(dataSource string) (string, error) {
	s := strings.SplitN(dataSource, "_", 2)
	if len(s) < 2 {
		return "", fmt.Errorf("Failed to build a data source doc URL: %s", dataSource)
	}

	// build a doc URL like this
	// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/security_group
	url := docBaseURL + s[0] + latestDocs + "/data-sources/" + s[1]
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
