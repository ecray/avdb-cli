package host

import (
	"encoding/json"
	"os"
)

type HostInfo struct {
	Host string                 `json:"host"`
	Data map[string]interface{} `json:"data"`
}

func formatOutput(resp []byte) error {
	var host *HostInfo
	err := json.Unmarshal(resp, &host)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(host)

	return nil
}

func formatOutputAll(resp []byte) error {
	var hosts []*HostInfo
	err := json.Unmarshal(resp, &hosts)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(hosts)

	return nil
}
