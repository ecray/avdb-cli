package group

import (
	"fmt"
	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
)

var groupDeleteCmd = cli.Command{
	Name:      "delete",
	Usage:     "delete group",
	ArgsUsage: "<group/name>",
	Action:    groupDelete,
}

func groupDelete(c *cli.Context) error {
	name := c.Args().First()

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/api/v1/groups/%s", conn.Server, name)

	_, err = conn.DoRequest("DELETE", uri, "")
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
