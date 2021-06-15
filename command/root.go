package command

import "github.com/mitchellh/cli"

func InitCommands(ui cli.Ui) map[string]cli.CommandFactory {
	meta := Meta{
		UI: ui,
	}

	commands := map[string]cli.CommandFactory{
		"provider show": func() (cli.Command, error) {
			return &ProviderShowCommand{
				Meta: meta,
			}, nil
		},
		"provider browse": func() (cli.Command, error) {
			return &ProviderBrowseCommand{
				Meta: meta,
			}, nil
		},
		"resource list": func() (cli.Command, error) {
			return &ResourceListCommand{
				Meta: meta,
			}, nil
		},
		"resource show": func() (cli.Command, error) {
			return &ResourceShowCommand{
				Meta: meta,
			}, nil
		},
		"resource browse": func() (cli.Command, error) {
			return &ResourceBrowseCommand{
				Meta: meta,
			}, nil
		},
		"data list": func() (cli.Command, error) {
			return &DataListCommand{
				Meta: meta,
			}, nil
		},
		"data show": func() (cli.Command, error) {
			return &DataShowCommand{
				Meta: meta,
			}, nil
		},
		"data browse": func() (cli.Command, error) {
			return &DataBrowseCommand{
				Meta: meta,
			}, nil
		},
	}

	return commands
}
