package main

import (
	"fmt"
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
	setLogFilter()
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

func getLogLevel() string {
	defaultLevel := "WARN"

	envLevel := os.Getenv("TFSCHEMA_LOG")
	if envLevel == "" {
		return defaultLevel
	}

	return envLevel
}

func setLogFilter() {
	levels := []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"}
	minLevel := getLogLevel()

	filter := &logutils.LevelFilter{
		Levels:   levels,
		MinLevel: logutils.LogLevel(minLevel),
		Writer:   os.Stderr,
	}

	log.SetOutput(filter)
}
