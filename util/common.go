package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	_ "os"
	"strings"
)

type Connection struct {
	Server string
	ApiKey string
}

func NewConnection(c *cli.Context) (Connection, error) {
	var (
		conn   Connection
		token  = c.GlobalString("token")
		server = c.GlobalString("server")
	)

	server = strings.TrimRight(server, "/")
	conn.Server, conn.ApiKey = server, token

	// token required for post, put, delete operations
	if len(token) == 0 {
		return conn, fmt.Errorf("Missing token.")
	}

	return conn, nil
}

//func DoRequest(method, uri, data string) ([]byte, error) {
func (conn *Connection) DoRequest(method, uri, data, query string) ([]byte, error) {
	client := &http.Client{}

	var req *http.Request
	var err error
	if method == "POST" || method == "PUT" {
		req, err = http.NewRequest(method, uri, bytes.NewBuffer([]byte(data)))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, uri, nil)
		if err != nil {
			return nil, err
		}
	}

	if query != "" {
		qmap := make(map[string]string)
		// needs more for loop - bruce dickinson
		split := strings.Split(query, "=")
		qmap[split[0]] = split[1]

		q := req.URL.Query()
		for k, v := range qmap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	//req.SetBasicAuth("USER", "PASSWD")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-Token", conn.ApiKey)
	//req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("HTTP Code: %d", resp.StatusCode))
	} else if resp.StatusCode > 300 {
		return nil, errors.New(fmt.Sprintf("Something is fucked: ", resp.StatusCode))
	}
	// May need handler for 204 with DELETE

	// convert to []byte to return
	rq, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rq, nil
}
