package opa

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestClient_Query(t *testing.T) {
	opaURL := os.Getenv("OPA_URL")
	if opaURL == "" {
		t.Skip("OPA_URL not set")
	}

	tests := []struct {
		Policy string
		Input  interface{}
		Result func() interface{}
		Cmp    func(interface{}) error
	}{
		{
			Policy: "arcus.pkg.opa.example.allow",
			Input: map[string]interface{}{
				"user": "joe",
			},
			Result: func() interface{} {
				var b bool
				return &b
			},
			Cmp: func(v interface{}) error {
				r := v.(*bool)
				if !*r {
					return errors.New("expected true")
				}
				return nil
			},
		},
		{
			Policy: "arcus.pkg.opa.example.allow",
			Input: map[string]interface{}{
				"role": "admin",
			},
			Result: func() interface{} {
				var b bool
				return &b
			},
			Cmp: func(v interface{}) error {
				r := v.(*bool)
				if !*r {
					return errors.New("expected true")
				}
				return nil
			},
		},
		{
			Policy: "arcus.pkg.opa.example.allow",
			Input:  nil,
			Result: func() interface{} {
				var b bool
				return &b
			},
			Cmp: func(v interface{}) error {
				r := v.(*bool)
				if *r {
					return errors.New("expected false")
				}
				return nil
			},
		},
		{
			Policy: "arcus.pkg.opa.example.get_perms",
			Input: map[string]interface{}{
				"role": "admin",
			},
			Result: func() interface{} {
				var a []string
				return &a
			},
			Cmp: func(v interface{}) error {
				r := *(v.(*[]string))
				e := []string{"a", "b"}
				if !reflect.DeepEqual(r, e) {
					return fmt.Errorf("expected %v, got %v", e, r)
				}
				return nil
			},
		},
		{
			Policy: "arcus.pkg.opa.example.get_perms",
			Input: map[string]interface{}{
				"user": "joe",
			},
			Result: func() interface{} {
				var a []string
				return &a
			},
			Cmp: func(v interface{}) error {
				r := *(v.(*[]string))
				e := []string{"x", "y", "z"}
				if !reflect.DeepEqual(r, e) {
					return fmt.Errorf("expected %v, got %v", e, r)
				}
				return nil
			},
		},
	}

	c := NewClient(opaURL)
	ctx := context.Background()

	for i, x := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := x.Result()

			if err := c.Query(
				ctx,
				x.Policy,
				x.Input,
				result,
			); err != nil {
				t.Fatal(err)
			}

			if err := x.Cmp(result); err != nil {
				t.Error(err)
			}
		})
	}
}
