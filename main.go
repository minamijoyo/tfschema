package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/logutils"
	"github.com/mitchellh/cli"
)

func init() {
	Ui = &cli.BasicUi{
		Writer: os.Stdout,
	}
}

func main() {
	setLogFilter(os.Getenv("TFSCHEMA_LOG"))
	log.Printf("[INFO] CLI args: %#v", os.Args)

	if Commands == nil {
		initCommands()
	}

	args := os.Args[1:]

	c := &cli.CLI{
		Name:       "tfschema",
		Args:       args,
		Commands:   Commands,
		HelpWriter: os.Stdout,
	}

	exitStatus, err := c.Run()
	if err != nil {
		Ui.Error(fmt.Sprintf("Failed to execute CLI: %s", err))
	}

	os.Exit(exitStatus)
}

func setLogFilter(minLevel string) {
	levels := []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR"}

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

	log.SetOutput(filter)
}
