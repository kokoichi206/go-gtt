package cmd

import (
	"github.com/kokoichi206/go-gtt/cmd/generate"
	"github.com/kokoichi206/go-gtt/cmd/update"
	"github.com/urfave/cli/v2"
)

// All commands.
func NewCommands() []*cli.Command {
	return []*cli.Command{
		generate.Command(),
		update.Command(),
	}
}
