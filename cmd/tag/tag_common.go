package tag

import (
	"encoding/json"
	"os"
)

type Tags interface{}

type TagInfo struct {
	Name string `json:"tag"`
	Host string `json:"host"`
}

func formatOutput(resp []byte) error {
	var tags *Tags
	err := json.Unmarshal(resp, &tags)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(tags)

	return nil
}
