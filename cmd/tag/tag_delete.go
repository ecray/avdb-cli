package tag

import (
	"fmt"

	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
)

var tagDeleteCmd = cli.Command{
	Name:      "delete",
	Usage:     "avdb-cli tag delete WEB_CLUSTER web01",
	ArgsUsage: "<tag/name>",
	Action:    tagDelete,
}

func tagDelete(c *cli.Context) error {
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

	_, err = conn.DoRequest("DELETE", uri, payload)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
