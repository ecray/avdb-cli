package tag

import (
	"fmt"
	"strings"

	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
)

var tagGetCmd = cli.Command{
	Name:      "get",
	Usage:     "get tag",
	ArgsUsage: "<tag/name>",
	Action: func(c *cli.Context) error {
		host := c.Args().First()
		//host := c.Args().Get(0)
		if host != "all" {
			return tagGet(c)
		} else {
			return tagGetAll(c)
		}
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "query, q",
			Usage: "Query tags. Example: avdb-cli tag get all -q host=web01",
		},
	},
}

func tagGet(c *cli.Context) error {
	name := c.Args().First()

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/api/v1/tags/%s", conn.Server, name)

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

func tagGetAll(c *cli.Context) error {
	query := strings.Split(c.String("query"), ",")

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/api/v1/tags", conn.Server)

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
