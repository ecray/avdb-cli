package host

import (
	"github.com/urfave/cli"
)

// Host commands
var Command = cli.Command{
	Name:  "host",
	Usage: "manage hosts",
	Subcommands: []cli.Command{
		hostAddCmd,
		hostGetCmd,
		hostUpdateCmd,
		hostDeleteCmd,
	},
}
