package main

/*
Rewrite of avdb python client

avdb-cli host get all
avdb-cli host get infdcpdn01
avdb-cli host get all -q colo=las1
avdb-cli host add infdcpdns01 -d $(jo colo=las1)
avdb-cli host update infdcpdns01 -d $(jo colo=aws)
avdb-cli host delete infdcpdns01
*/

import (
	"fmt"
	"github.com/ecray/avdb-cli/cmd/group"
	"github.com/ecray/avdb-cli/cmd/host"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "avdb-cli"
	app.Usage = "manage AVDB"
	app.Version = "0.0.1"
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
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command %q not found!\n", command)
	}
	app.Run(os.Args)
}
