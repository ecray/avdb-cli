package host

import (
	"fmt"
	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
	"strings"
)

var hostGetCmd = cli.Command{
	Name:      "get",
	Usage:     "get host",
	ArgsUsage: "<host/name>",
	Action: func(c *cli.Context) error {
		host := c.Args().First()
		//host := c.Args().Get(0)
		if host != "all" {
			return hostGet(c)
		} else {
			return hostGetAll(c)
		}
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "query, q",
			Usage: "Query hosts. Example: avdb-cli get host all -q colo=aws1",
		},
	},
}

func hostGet(c *cli.Context) error {
	name := c.Args().First()

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/api/v1/hosts/%s", conn.Server, name)

	resp, err := conn.DoRequest("GET", uri, "")
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = formatOutput(resp)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

func hostGetAll(c *cli.Context) error {
	query := strings.Split(c.String("query"), ",")

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/api/v1/hosts", conn.Server)

	resp, err := conn.DoQueryRequest("GET", uri, "", query)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	formatOutput(resp)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
