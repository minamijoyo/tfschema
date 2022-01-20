package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"syscall"

	"github.com/hashicorp/logutils"
	"github.com/minamijoyo/tfschema/command"
	"github.com/mitchellh/cli"
	"github.com/mitchellh/panicwrap"
)

// Version is a version number.
var version = "0.7.2"

// ui is a user interface which is a global variable for mocking.
var ui cli.Ui

func init() {
	ui = &cli.BasicUi{
		Writer: os.Stdout,
	}
}

func main() {
	// abuse panicwrap to discard noisy debug log from go-plugin
	var wrapConfig panicwrap.WrapConfig
	if !panicwrap.Wrapped(&wrapConfig) {
		doneCh := make(chan struct{})
		outR, outW := io.Pipe()
		go copyOutput(outR, doneCh)

		wrapConfig.Handler = panicHandler
		wrapConfig.Writer = logOutput()
		wrapConfig.Stdout = outW
		wrapConfig.IgnoreSignals = []os.Signal{os.Interrupt}
		wrapConfig.ForwardSignals = []os.Signal{syscall.SIGTERM}

		exitStatus, err := panicwrap.Wrap(&wrapConfig)
		if err != nil {
			panic(err)
		}

		if exitStatus >= 0 {
			outW.Close()
			<-doneCh
			os.Exit(exitStatus)
		}
	}

	os.Exit(wrappedMain())
}

func wrappedMain() int {
	log.SetOutput(logOutput())
	log.Printf("[INFO] CLI args: %#v", os.Args)

	commands := command.InitCommands(ui)

	args := os.Args[1:]

	c := &cli.CLI{
		Name:                  "tfschema",
		Version:               version,
		Args:                  args,
		Commands:              commands,
		HelpWriter:            os.Stdout,
		Autocomplete:          true,
		AutocompleteInstall:   "install-autocomplete",
		AutocompleteUninstall: "uninstall-autocomplete",
	}

	exitStatus, err := c.Run()
	if err != nil {
		ui.Error(fmt.Sprintf("Failed to execute CLI: %s", err))
	}

	return exitStatus
}

func panicHandler(output string) {
	ui.Error(fmt.Sprintf("The child panicked:\n\n%s\n", output))
	os.Exit(1)
}

func logOutput() io.Writer {
	levels := []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"}
	minLevel := os.Getenv("TFSCHEMA_LOG")

	// default log writer is null device.
	writer := ioutil.Discard
	if minLevel != "" {
		writer = os.Stderr
	}

	filter := &logutils.LevelFilter{
		Levels:   levels,
		MinLevel: logutils.LogLevel(minLevel),
		Writer:   writer,
	}

	return filter
}

func copyOutput(r io.Reader, doneCh chan<- struct{}) {
	defer close(doneCh)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// nolint: errcheck
		// We should check for errors here, but haven't done yet.
		io.Copy(os.Stdout, r)
	}()

	wg.Wait()
}
