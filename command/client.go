package command

import (
	"os"

	"github.com/minamijoyo/tfschema/tfschema"
)

// NewDefaultClient creates a new Client instance.
func NewDefaultClient(providerName string) (tfschema.Client, error) {
	rootDir := os.Getenv("TFSCHEMA_ROOT_DIR")
	if rootDir == "" {
		rootDir = "."
	}

	options := tfschema.Option{RootDir: rootDir}

	return tfschema.NewClient(providerName, options)
}
