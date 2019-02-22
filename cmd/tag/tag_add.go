package tag

import (
	"fmt"

	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
)

var tagAddCmd = cli.Command{
	Name:      "add",
	Usage:     "avdb-cli tag add WEB_CLUSTER web01",
	ArgsUsage: "<tag/name>",
	Action:    tagAdd,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "Add tag. Example: avdb-cli tag add WEB_CLUSTER --host web01",
		},
	},
}

func tagAdd(c *cli.Context) error {
	// Parse args
	name := c.Args().First()
	host := c.Args().Get(1)
	if host == "" {
		return cli.NewExitError("Missing host argument", 1)
	}

	payload := fmt.Sprintf("{\"host\":\"%+v\"}", host)

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/api/v1/tags/%s", conn.Server, name)

	resp, err := conn.DoRequest("POST", uri, payload)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = formatOutput(resp)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
