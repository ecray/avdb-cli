package group

import (
	"encoding/json"
	"os"
)

type Groups interface{}

type GroupInfo struct {
	Group string                 `json:"group"`
	Data  map[string]interface{} `json:"data,omitempty"`
	Hosts []*string              `json:"hosts,omitempty"`
}

func formatOutput(resp []byte) error {
	var groups *Groups

	err := json.Unmarshal(resp, &groups)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(groups)

	return nil
}
