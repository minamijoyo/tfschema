package command

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mitchellh/cli"
)

// newMockCLI returns a new mock CLI for testing.
func newMockCLI(ui cli.Ui, args []string) *cli.CLI {
	commands := InitCommands(ui)

	return &cli.CLI{
		Commands: commands,
		Args:     args,
	}
}

// skipUnlessAcceptanceTestEnabled skips acceptance tests unless TEST_ACC is set to 1.
func skipUnlessAcceptanceTestEnabled(t *testing.T) {
	t.Helper()
	if os.Getenv("TEST_ACC") != "1" {
		t.Skip("skip acceptance tests")
	}
}

// setupTestAcc is a common setup helper for acceptance tests.
func setupTestAcc(t *testing.T, providerName string, providerVersion string) {
	t.Helper()

	// generate source
	source := fmt.Sprintf(`provider "%s" { version = "%s" }`, providerName, providerVersion)

	// create a workDir
	workDir, err := setupTestWorkDir(source)
	if err != nil {
		t.Fatalf("failed to setup work dir: %s", err)
	}

	// cd workDir
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %s", err)
	}
	if err := os.Chdir(workDir); err != nil {
		t.Fatalf("failed to change dir to %s on setup: %s", workDir, err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(oldDir); err != nil {
			t.Fatalf("failed to change dir to %s on cleanup: %s", oldDir, err)
		}
		os.RemoveAll(workDir)
	})

	// terraform init
	cmd := exec.Command("terraform", "init")
	cmd.Dir = workDir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to run terraform init: %s, out: %s", err, out)
	}

	// check if the workDir was initizalied.
	if _, err := os.Stat(".terraform"); os.IsNotExist(err) {
		t.Fatalf("failed to find .terraform directory: %s", err)
	}
}

// setupTestWorkDir creates temporary working directory with a given source for testing.
func setupTestWorkDir(source string) (string, error) {
	workDir, err := os.MkdirTemp("", "workDir")
	if err != nil {
		return "", fmt.Errorf("failed to create work dir: %s", err)
	}

	if err := os.WriteFile(filepath.Join(workDir, "main.tf"), []byte(source), 0600); err != nil {
		os.RemoveAll(workDir)
		return "", fmt.Errorf("failed to create main.tf: %s", err)
	}

	return workDir, nil
}
