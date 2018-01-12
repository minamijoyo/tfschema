package main

import (
	"github.com/minamijoyo/tfschema/command"
	"github.com/mitchellh/cli"
)

var Ui cli.Ui
var Commands map[string]cli.CommandFactory

func initCommands() {
	meta := command.Meta{
		Ui: Ui,
	}

	Commands = map[string]cli.CommandFactory{
		"resource show": func() (cli.Command, error) {
			return &command.ResourceShowCommand{
				Meta: meta,
			}, nil
		},
	}
}
