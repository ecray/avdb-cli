package group

import (
	"github.com/urfave/cli"
)

// Group commands
var Command = cli.Command{
	Name:  "group",
	Usage: "manage groups",
	Subcommands: []cli.Command{
		groupAddCmd,
		groupGetCmd,
		groupUpdateCmd,
		groupDeleteCmd,
	},
}
