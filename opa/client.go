package opa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/arcus/pkg/service"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

const (
	ErrNoDecision = Error("opa: no decision")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type jsonResult struct {
	Result interface{}
}

type authzResult struct {
	Allow  bool
	Claims interface{}
}

// newServiceInput derives an authorization request input from the service context and request itself.
func newServiceInput(ctx *service.Context, req proto.Message) (interface{}, error) {
	mpb := &jsonpb.Marshaler{
		OrigName: true,
	}

	reqjs, err := mpb.MarshalToString(req)
	if err != nil {
		return nil, err
	}

	input := map[string]interface{}{
		"service": map[string]interface{}{
			"name":    ctx.ServiceName,
			"version": ctx.ServiceVersion,
		},
		"principal": map[string]interface{}{
			"id":   ctx.PrincipalId,
			"role": ctx.PrincipalRole,
		},
		"operation": map[string]interface{}{
			"id":    proto.MessageName(req),
			"input": json.RawMessage(reqjs),
		},
	}

	return input, nil
}

type Client struct {
	Addr string
	HTTP *http.Client
}

func (c *Client) Authorized(ctx *service.Context, policy string, req proto.Message, claims interface{}) (bool, error) {
	input, err := newServiceInput(ctx, req)
	if err != nil {
		return false, err
	}

	result := authzResult{
		Claims: claims,
	}

	err = c.Query(ctx, policy, input, &result)
	if err != nil {
		return false, err
	}

	return result.Allow, nil
}

func (c *Client) Query(ctx context.Context, policy string, input interface{}, result interface{}) error {
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

func NewClient(addr string) *Client {
	return &Client{
		Addr: addr,
		HTTP: http.DefaultClient,
	}
}
