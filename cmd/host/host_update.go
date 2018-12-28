package host

import (
	"fmt"
	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
)

var hostUpdateCmd = cli.Command{
	Name:      "update",
	Usage:     "update host",
	ArgsUsage: "<host/name>",
	Action:    hostUpdate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "data, d",
			Usage: "Update host. Example: avdb-cli update host infdcpdns01 -d '{\"colo\":\"las1\"}'",
		},
	},
}

func hostUpdate(c *cli.Context) error {
	name := c.Args().First()
	data := c.String("data")

	conn, err := util.NewConnection(c)
	uri := fmt.Sprintf("%s/api/v1/hosts/%s", conn.Server, name)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	resp, err := conn.DoRequest("PUT", uri, data)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = formatOutput(resp)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
