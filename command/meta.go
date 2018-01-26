package command

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

type Meta struct {
	Ui cli.Ui
}

func detectProviderName(name string) (string, error) {
	s := strings.SplitN(name, "_", 2)
	if len(s) < 2 {
		return "", fmt.Errorf("Failed to detect a provider name from the argument: %s", name)
	}
	return s[0], nil
}
