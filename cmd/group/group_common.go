package group

import (
	"encoding/json"
	"os"
)

type GroupInfo struct {
	Group string                 `json:"group"`
	Data  map[string]interface{} `json:"data"`
	Hosts []string               `json:"hosts"`
}

func formatOutput(resp []byte) error {
	var group *GroupInfo
	err := json.Unmarshal(resp, &group)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(group)

	return nil
}

func formatOutputAll(resp []byte) error {
	var groups []*GroupInfo
	err := json.Unmarshal(resp, &groups)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(groups)

	return nil
}
