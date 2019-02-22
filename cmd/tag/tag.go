package tag

import (
	"github.com/urfave/cli"
)

// Tag commands
var Command = cli.Command{
	Name:  "tag",
	Usage: "manage tags",
	Subcommands: []cli.Command{
		tagAddCmd,
		tagGetCmd,
		tagDeleteCmd,
	},
}
