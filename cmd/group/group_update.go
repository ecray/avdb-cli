package group

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
)

var groupUpdateCmd = cli.Command{
	Name:      "update",
	Usage:     "update group",
	ArgsUsage: "<group/name>",
	Action:    groupUpdate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "data, d",
			Usage: "Update group. Example: avdb-cli update group foodtrucks -d '{\"colo\":\"las1\"}'",
			Value: "{}",
		},
		cli.StringFlag{
			Name:  "hosts",
			Usage: "Update/Delete group hosts. Ex: avdb-cli update group las1-adm --hosts lvopsdcadm01,-infdcpdns01",
		},
	},
}

func groupUpdate(c *cli.Context) error {
	name := c.Args().First()
	data := c.String("data")
	// convert hosts to []string
	h := strings.Split(c.String("hosts"), ",")

	// check if there is anything to do
	if data == "{}" && len(h[0]) == 0 {
		return cli.NewExitError("Nothing to update", 0)
	}

	// marshal hosts
	hosts, err := json.Marshal(h)
	if err != nil {
		return cli.NewExitError("Failed to marshal object", 1)
	}
	payload := fmt.Sprintf("{\"data\":%v,\"hosts\":%+v}", data, string(hosts))

	conn, err := util.NewConnection(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	uri := fmt.Sprintf("%s/api/v1/groups/%s", conn.Server, name)

	resp, err := conn.DoRequest("PUT", uri, payload)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = formatOutput(resp)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
