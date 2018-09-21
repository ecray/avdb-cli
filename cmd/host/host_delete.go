package host

import (
	"fmt"
	"github.com/urfave/cli"
	"github.marqeta.com/ecray/avdb-cli/util"
)

var hostDeleteCmd = cli.Command{
	Name:      "delete",
	Usage:     "delete host",
	ArgsUsage: "<host/name>",
	Action:    hostDelete,
}

func hostDelete(c *cli.Context) error {
	name := c.Args().First()

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/hosts/%s", conn.Server, name)

	_, err = conn.DoRequest("DELETE", uri, "")
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
