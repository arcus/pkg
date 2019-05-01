package opa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNoDecision = Error("opa: no decision")
)

type Client struct {
	Addr string
	HTTP *http.Client
}

func (c *Client) Query(ctx context.Context, policy string, input, result interface{}) error {
	path := strings.Replace(strings.Replace(policy, ".", "/", -1), "-", "_", -1)
	url := fmt.Sprintf("%s/v1/data/%s", c.Addr, path)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(map[string]interface{}{
		"input": input,
	}); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/json")

	rep, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer rep.Body.Close()

	b, _ := ioutil.ReadAll(rep.Body)
	if rep.StatusCode != 200 {
		return fmt.Errorf("opa: %s", string(b))
	}

	if string(b) == "{}" {
		return ErrNoDecision
	}

	return json.Unmarshal(b, &jsonResult{
		Result: result,
	})
}

type jsonResult struct {
	Result interface{}
}

func NewClient(addr string) *Client {
	return &Client{
		Addr: addr,
		HTTP: http.DefaultClient,
	}
}
