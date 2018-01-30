package command

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

const docBaseURL = "https://www.terraform.io/docs/providers/"

type Meta struct {
	Ui cli.Ui
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
