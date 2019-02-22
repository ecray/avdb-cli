package main

import (
	"fmt"
	"os"

	"github.com/ecray/avdb-cli/cmd/group"
	"github.com/ecray/avdb-cli/cmd/host"
	"github.com/ecray/avdb-cli/cmd/tag"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "avdb-cli"
	app.Usage = "manage AVDB"
	app.Version = "0.2.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Eric Raymond",
			Email: "ecraymond@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "t, token",
			Usage:  "server auth token",
			EnvVar: "AVDB_TOKEN",
		},
		cli.StringFlag{
			Name:   "s, server",
			Usage:  "server address",
			Value:  "http://127.0.0.1:3000",
			EnvVar: "AVDB_SERVER",
		},
	}
	app.Commands = []cli.Command{
		host.Command,
		group.Command,
		tag.Command,
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command %q not found!\n", command)
	}
	app.Run(os.Args)
}
