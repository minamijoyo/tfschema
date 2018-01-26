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

var Ui cli.Ui

func init() {
	Ui = &cli.BasicUi{
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

	commands := initCommands()

	args := os.Args[1:]

	c := &cli.CLI{
		Name:       "tfschema",
		Args:       args,
		Commands:   commands,
		HelpWriter: os.Stdout,
	}

	exitStatus, err := c.Run()
	if err != nil {
		Ui.Error(fmt.Sprintf("Failed to execute CLI: %s", err))
	}

	return exitStatus
}

func panicHandler(output string) {
	Ui.Error(fmt.Sprintf("The child panicked:\n\n%s\n", output))
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
		io.Copy(os.Stdout, r)
	}()

	wg.Wait()
}

func initCommands() map[string]cli.CommandFactory {
	meta := command.Meta{
		Ui: Ui,
	}

	commands := map[string]cli.CommandFactory{
		"resource show": func() (cli.Command, error) {
			return &command.ResourceShowCommand{
				Meta: meta,
			}, nil
		},
		"resource list": func() (cli.Command, error) {
			return &command.ResourceListCommand{
				Meta: meta,
			}, nil
		},
		"data list": func() (cli.Command, error) {
			return &command.DataListCommand{
				Meta: meta,
			}, nil
		},
	}

	return commands
}
