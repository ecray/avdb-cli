package group

import (
	"encoding/json"
	"fmt"
	"github.com/ecray/avdb-cli/util"
	"github.com/urfave/cli"
	"strings"
)

var groupAddCmd = cli.Command{
	Name:      "add",
	Usage:     "add group",
	ArgsUsage: "<group/name>",
	Action:    groupAdd,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "data, d",
			Usage: "Add group. Example: avdb-cli add group infdcpdns01 -d '{\"colo\":\"las1\"}'",
			Value: "{}",
		},
		cli.StringFlag{
			Name:  "hosts",
			Usage: "Add group hosts. Ex: avdb-cli add group las1-adm --hosts infdcpdns01,lvopsdcadm01",
		},
	},
}

func groupAdd(c *cli.Context) error {
	name := c.Args().First()
	data := c.String("data")
	var h []string
	// convert hosts to []string, then marshal
	if len(c.String("hosts")) == 0 {
		// leave h null slice
	} else {
		h = strings.Split(c.String("hosts"), ",")
	}

	hosts, err := json.Marshal(h)
	if err != nil {
		return cli.NewExitError("Failed to marshal object", 1)
	}
	payload := fmt.Sprintf("{\"data\":%v,\"hosts\":%+v}", data, string(hosts))
	//fmt.Println("String Payload: ", string(payload))

	conn, err := util.NewConnection(c)
	uri := fmt.Sprintf("%s/groups/%s", conn.Server, name)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

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
