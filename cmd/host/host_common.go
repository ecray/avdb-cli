package host

import (
	"encoding/json"
	"os"
)

type Hosts interface{}

type HostInfo struct {
	Host string                 `json:"host"`
	Data map[string]interface{} `json:"data"`
}

func formatOutput(resp []byte) error {
	var hosts *Hosts
	err := json.Unmarshal(resp, &hosts)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(hosts)

	return nil
}
