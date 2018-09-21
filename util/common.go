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

// DRY work to be done
func (conn *Connection) DoQueryRequest(method, uri, data string, query []string) ([]byte, error) {
	var req *http.Request
	var err error

	req, err = http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	if len(query[0]) != 0 {
		q := req.URL.Query()
		for _, s := range query {
			tmp := strings.Split(s, "=")
			q.Add(tmp[0], tmp[1])
		}
		req.URL.RawQuery = q.Encode()
	}
	return conn.requestHandler(req)
}

func (conn *Connection) DoRequest(method, uri, data string) ([]byte, error) {
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
	return conn.requestHandler(req)
}

func (conn *Connection) requestHandler(req *http.Request) ([]byte, error) {
	client := &http.Client{}

	//req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-Token", conn.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// error handling
	if resp.StatusCode >= 400 {
		rc := errors.New(fmt.Sprintf("%s - %d", http.StatusText(resp.StatusCode), resp.StatusCode))
		return nil, rc
	}

	// convert to []byte to return
	rq, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rq, nil
}
