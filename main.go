package main

import (
	"log"
	"os"

	"github.com/minamijoyo/tfschema/command"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("tfschema", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"resource show": func() (cli.Command, error) {
			return &command.ResourceShowCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
