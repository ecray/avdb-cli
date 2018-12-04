package group

import (
	"fmt"
	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
	"strings"
	"time"
)

type Group struct {
	Name      string      `json:"group"`
	Data      interface{} `json:"data"`
	Hosts     []string    `json:"hosts"`
	CreatedAt time.Time   `json:"-"`
	DeletedAt time.Time   `json:"-"`
	UpdatedAt time.Time   `json:"-"`
}

var groupGetCmd = cli.Command{
	Name:      "get",
	Usage:     "get group",
	ArgsUsage: "<group/name>",
	Action: func(c *cli.Context) error {
		group := c.Args().First()
		//host := c.Args().Get(0)
		if group != "all" {
			return groupGet(c)
		} else {
			return groupGetAll(c)
		}
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "query, q",
			Usage: "Query group. Example: avdb-cli get group all -q colo=aws1",
		},
	},
}

func groupGet(c *cli.Context) error {
	name := c.Args().First()

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/groups/%s", conn.Server, name)

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

func groupGetAll(c *cli.Context) error {
	query := strings.Split(c.String("query"), ",")

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/groups", conn.Server)

	resp, err := conn.DoQueryRequest("GET", uri, "", query)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = formatOutput(resp)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
