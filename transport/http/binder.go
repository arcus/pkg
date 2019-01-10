package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	gendescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/labstack/echo"
)

// protoBinder implements the echo.Binder interface to handle decoding
// query parameters in GET and DELETE requests and JSON or form encoded
// bodies in all other requests into protobuf messages. This relies on
// the jsonpb library.
type ProtoBinder struct{}

func (b *ProtoBinder) Bind(i interface{}, c echo.Context) error {
	// Ignore if not a proto message.
	m, ok := i.(proto.Message)
	if !ok {
		d := &echo.DefaultBinder{}
		return d.Bind(i, c)
	}

	req := c.Request()
	if req.Method == http.MethodGet || req.Method == http.MethodDelete {
		return paramsToProto(c.QueryParams(), m.(descriptor.Message))
	}

	ctype := req.Header.Get(echo.HeaderContentType)
	switch {
	case strings.HasPrefix(ctype, echo.MIMEApplicationJSON):
		if err := jsonpb.Unmarshal(req.Body, m); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	case strings.HasPrefix(ctype, echo.MIMEApplicationForm):
		data, err := c.FormParams()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return paramsToProto(data, m.(descriptor.Message))
	default:
		d := &echo.DefaultBinder{}
		return d.Bind(i, c)
	}

	return nil
}

// paramsToProto decodes query or form parameter values into a protobuf
// descriptor message by sniffing the types of the message fields and
// coercing the values into that type before encoding the values as JSON
// and decoding them into the message using the jsonpb package.
func paramsToProto(params url.Values, msg descriptor.Message) error {
	// Use generic map for setting fields.
	m := make(map[string]interface{})

	// Get the message descriptor.
	_, md := descriptor.ForMessage(msg)
	// Loop through fields
	for _, f := range md.GetField() {
		key := f.GetName()
		// Not set in params.
		if _, ok := params[key]; !ok {
			continue
		}

		var (
			err error
			val interface{}
		)

		switch f.GetType() {
		case gendescriptor.FieldDescriptorProto_TYPE_DOUBLE:
			val, err = strconv.ParseFloat(params.Get(key), 64)
			if err != nil {
				return fmt.Errorf("%s: expected double: %s", key, err)
			}

		case gendescriptor.FieldDescriptorProto_TYPE_FLOAT:
			val, err = strconv.ParseFloat(params.Get(key), 32)
			if err != nil {
				return fmt.Errorf("%s: expected float: %s", key, err)
			}

		case gendescriptor.FieldDescriptorProto_TYPE_INT64, gendescriptor.FieldDescriptorProto_TYPE_SFIXED64, gendescriptor.FieldDescriptorProto_TYPE_SINT64:
			val, err = strconv.ParseInt(params.Get(key), 0, 64)
			if err != nil {
				return fmt.Errorf("%s: expected int64 : %s", key, err)
			}

		case gendescriptor.FieldDescriptorProto_TYPE_UINT64, gendescriptor.FieldDescriptorProto_TYPE_FIXED64:
			val, err = strconv.ParseUint(params.Get(key), 0, 64)
			if err != nil {
				return fmt.Errorf("%s: expected uint64 : %s", key, err)
			}

		case gendescriptor.FieldDescriptorProto_TYPE_INT32, gendescriptor.FieldDescriptorProto_TYPE_SFIXED32, gendescriptor.FieldDescriptorProto_TYPE_SINT32:
			val, err = strconv.ParseInt(params.Get(key), 0, 32)
			if err != nil {
				return fmt.Errorf("%s: expected int32: %s", key, err)
			}

		case gendescriptor.FieldDescriptorProto_TYPE_UINT32, gendescriptor.FieldDescriptorProto_TYPE_FIXED32:
			val, err = strconv.ParseUint(params.Get(key), 0, 32)
			if err != nil {
				return fmt.Errorf("%s: expected uint32 : %s", key, err)
			}

		case gendescriptor.FieldDescriptorProto_TYPE_BOOL:
			val, err = strconv.ParseBool(params.Get(key))
			if err != nil {
				return fmt.Errorf("%s: expected bool: %s", key, err)
			}

		case gendescriptor.FieldDescriptorProto_TYPE_STRING:
			val = params.Get(key)

		case gendescriptor.FieldDescriptorProto_TYPE_BYTES:
			// This is assumed to be base64 encoded which the jsobpb library will handle.
			val = params.Get(key)

		case gendescriptor.FieldDescriptorProto_TYPE_ENUM:
			// Enum passed as string, jsonpb handles converting this to the
			// internal integer. The target value is assumed to be upper-cased with underscores
			// per the protobuf style guide.
			val = strings.Replace(strings.ToUpper(params.Get(key)), "-", "_", -1)

		case gendescriptor.FieldDescriptorProto_TYPE_MESSAGE:
			// Handle common message types.
			switch f.GetTypeName() {
			case ".google.protobuf.Duration", ".google.protobuf.Timestamp":
				val = params.Get(key)

			// Handles struct, any, and empty WKTs, and any custom message type.
			default:
				val = json.RawMessage(params.Get(key))
			}

		default:
			return fmt.Errorf("unsupported type: %s", f.GetTypeName())
		}

		m[key] = val
	}

	// Marshal as JSON and unpack into message.
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return jsonpb.Unmarshal(bytes.NewBuffer(b), msg)
}
