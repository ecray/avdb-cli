package host

import (
	"fmt"
	"github.com/urfave/cli"
	"github.marqeta.com/ecray/avdb-cli/util"
)

var hostAddCmd = cli.Command{
	Name:      "add",
	Usage:     "add host",
	ArgsUsage: "<host/name>",
	Action:    hostAdd,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "data, d",
			Usage: "Add host. Example: avdb-cli add host infdcpdns01 -d '{\"colo\":\"las1\"}'",
		},
	},
}

func hostAdd(c *cli.Context) error {
	name := c.Args().First()
	data := c.String("data")

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/hosts/%s", conn.Server, name)

	resp, err := conn.DoRequest("POST", uri, data)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = formatOutput(resp)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
