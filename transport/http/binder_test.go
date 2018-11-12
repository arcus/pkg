package http

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/arcus/pkg/transport/http/internal/test"
	"github.com/golang/protobuf/ptypes"
	durationpb "github.com/golang/protobuf/ptypes/duration"
	emptypb "github.com/golang/protobuf/ptypes/empty"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

func TestParamsToProto(t *testing.T) {
	ts := ptypes.TimestampNow()
	any, _ := ptypes.MarshalAny(&durationpb.Duration{
		Seconds: 18000,
	})

	e := test.All{
		Int32:   32,
		Int64:   64,
		String_: "string",
		Bool:    true,
		Duration: &durationpb.Duration{
			Seconds: 18000,
		},
		Timestamp: ts,
		Any:       any,
		Empty:     &emptypb.Empty{},
		Struct: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"foo": &structpb.Value{
					Kind: &structpb.Value_NumberValue{
						NumberValue: 1,
					},
				},
			},
		},
		Enum: test.Enum_BAR_BAZ,
		Message: &test.Message{
			Message: "hello",
		},
		Map: map[string]float32{
			"foo": 1.3,
			"bar": 2.1,
		},
	}

	p := url.Values{
		"int32":     {"32"},
		"int64":     {"64"},
		"string":    {"string"},
		"bool":      {"true"},
		"duration":  {"5h"},
		"timestamp": {ptypes.TimestampString(ts)},
		"any":       {`{"@type": "type.googleapis.com/google.protobuf.Duration", "value": "5h"}`},
		"empty":     {`{}`},
		"struct":    {`{"foo": 1}`},
		"enum":      {"bar-baz"},
		"message":   {`{"message": "hello"}`},
		"map":       {`{"foo": 1.3, "bar": 2.1}`},
	}

	var a test.All
	if err := paramsToProto(p, &a); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(e, a) {
		t.Errorf("expected %s, got %s", &e, &a)
	}
}
