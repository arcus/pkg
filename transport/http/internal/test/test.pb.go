// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Enum int32

const (
	Enum_FOO     Enum = 0
	Enum_BAR_BAZ Enum = 1
)

var Enum_name = map[int32]string{
	0: "FOO",
	1: "BAR_BAZ",
}

var Enum_value = map[string]int32{
	"FOO":     0,
	"BAR_BAZ": 1,
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

type Message struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type All struct {
	Int32                int32                `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64                int64                `protobuf:"varint,2,opt,name=int64,proto3" json:"int64,omitempty"`
	Uint32               uint32               `protobuf:"varint,3,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64               uint64               `protobuf:"varint,4,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Bool                 bool                 `protobuf:"varint,7,opt,name=bool,proto3" json:"bool,omitempty"`
	String_              string               `protobuf:"bytes,10,opt,name=string,proto3" json:"string,omitempty"`
	Bytes                []byte               `protobuf:"bytes,11,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Map                  map[string]float32   `protobuf:"bytes,12,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	Any                  *any.Any             `protobuf:"bytes,20,opt,name=any,proto3" json:"any,omitempty"`
	Empty                *empty.Empty         `protobuf:"bytes,21,opt,name=empty,proto3" json:"empty,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,22,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Struct               *_struct.Struct      `protobuf:"bytes,23,opt,name=struct,proto3" json:"struct,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,24,opt,name=duration,proto3" json:"duration,omitempty"`
	Enum                 Enum                 `protobuf:"varint,30,opt,name=enum,proto3,enum=test.Enum" json:"enum,omitempty"`
	Message              *Message             `protobuf:"bytes,31,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}
func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (m *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(m, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetInt32() int32 {
	if m != nil {
		return m.Int32
	}
	return 0
}

func (m *All) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *All) GetUint32() uint32 {
	if m != nil {
		return m.Uint32
	}
	return 0
}

func (m *All) GetUint64() uint64 {
	if m != nil {
		return m.Uint64
	}
	return 0
}

func (m *All) GetBool() bool {
	if m != nil {
		return m.Bool
	}
	return false
}

func (m *All) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

func (m *All) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *All) GetMap() map[string]float32 {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *All) GetAny() *any.Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (m *All) GetEmpty() *empty.Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

func (m *All) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *All) GetStruct() *_struct.Struct {
	if m != nil {
		return m.Struct
	}
	return nil
}

func (m *All) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *All) GetEnum() Enum {
	if m != nil {
		return m.Enum
	}
	return Enum_FOO
}

func (m *All) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "test.Message")
	proto.RegisterType((*All)(nil), "test.All")
	proto.RegisterMapType((map[string]float32)(nil), "test.All.MapEntry")
	proto.RegisterEnum("test.Enum", Enum_name, Enum_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xbb, 0x91, 0x1c, 0x39, 0xe3, 0xa4, 0x98, 0xc1, 0x75, 0x36, 0x6a, 0x70, 0x96, 0xb4,
	0xb4, 0x4b, 0x29, 0x32, 0x38, 0xa9, 0x09, 0xbd, 0x29, 0xd4, 0xbd, 0x85, 0xc0, 0xb6, 0xa7, 0x5e,
	0x8a, 0x9c, 0x6e, 0x8d, 0xa9, 0xb4, 0x12, 0xd2, 0xaa, 0xa0, 0xd7, 0xeb, 0x93, 0x95, 0xfd, 0x23,
	0x07, 0xec, 0xdb, 0x7c, 0xf3, 0xfd, 0x46, 0xb3, 0x9a, 0x19, 0x00, 0x2d, 0x1b, 0x9d, 0x54, 0x75,
	0xa9, 0x4b, 0x0c, 0x4d, 0x1c, 0x5f, 0x6c, 0xca, 0x72, 0x93, 0xcb, 0xb9, 0xcd, 0xad, 0xdb, 0xdf,
	0xf3, 0x4c, 0x75, 0x0e, 0x88, 0x5f, 0xef, 0x5b, 0xb2, 0xa8, 0x74, 0x6f, 0x5e, 0xed, 0x9b, 0x7a,
	0x5b, 0xc8, 0x46, 0x67, 0x45, 0xe5, 0x81, 0xcb, 0x7d, 0xa0, 0xd1, 0x75, 0xfb, 0xe4, 0x9b, 0xc7,
	0xb3, 0x7d, 0xf7, 0x57, 0x5b, 0x67, 0x7a, 0x5b, 0x2a, 0xe7, 0x5f, 0xbf, 0x81, 0xe8, 0x41, 0x36,
	0x4d, 0xb6, 0x91, 0x48, 0x21, 0x2a, 0x5c, 0x48, 0x09, 0x23, 0xfc, 0x44, 0xf4, 0xf2, 0xfa, 0x5f,
	0x08, 0x41, 0x9a, 0xe7, 0x38, 0x81, 0xc1, 0x56, 0xe9, 0x9b, 0x85, 0xf5, 0x07, 0xc2, 0x09, 0x9f,
	0x5d, 0xde, 0xd2, 0x23, 0x46, 0x78, 0x20, 0x9c, 0xc0, 0x29, 0x1c, 0xb7, 0x0e, 0x0e, 0x18, 0xe1,
	0x67, 0xc2, 0xab, 0x3e, 0xbf, 0xbc, 0xa5, 0x21, 0x23, 0x3c, 0x14, 0x5e, 0x21, 0x42, 0xb8, 0x2e,
	0xcb, 0x9c, 0x46, 0x8c, 0xf0, 0xa1, 0xb0, 0xb1, 0x61, 0x1b, 0x5d, 0x6f, 0xd5, 0x86, 0x82, 0x7d,
	0x90, 0x57, 0xa6, 0xe3, 0xba, 0xd3, 0xb2, 0xa1, 0x23, 0x46, 0xf8, 0xa9, 0x70, 0x02, 0xdf, 0x42,
	0x50, 0x64, 0x15, 0x3d, 0x65, 0x01, 0x1f, 0x2d, 0x30, 0xb1, 0x1b, 0x48, 0xf3, 0x3c, 0x79, 0xc8,
	0xaa, 0x95, 0xd2, 0x75, 0x27, 0x8c, 0x8d, 0xef, 0x20, 0xc8, 0x54, 0x47, 0x27, 0x8c, 0xf0, 0xd1,
	0x62, 0x92, 0xb8, 0xf1, 0x24, 0xfd, 0x78, 0x92, 0x54, 0x75, 0xc2, 0x00, 0xf8, 0x11, 0x06, 0x76,
	0x0d, 0xf4, 0x95, 0x25, 0xa7, 0x07, 0xe4, 0xca, 0xb8, 0xc2, 0x41, 0x78, 0x07, 0x27, 0xbb, 0xbd,
	0xd0, 0xa9, 0xad, 0x88, 0x0f, 0x2a, 0xbe, 0xf7, 0x84, 0x78, 0x86, 0x71, 0x6e, 0xff, 0xb1, 0x7d,
	0xd2, 0xf4, 0xdc, 0x96, 0x9d, 0x1f, 0x94, 0x7d, 0xb3, 0xb6, 0xf0, 0x18, 0x7e, 0x82, 0x61, 0xbf,
	0x43, 0x4a, 0x6d, 0xc9, 0xc5, 0x41, 0xc9, 0x17, 0x0f, 0x88, 0x1d, 0x8a, 0x33, 0x08, 0xa5, 0x6a,
	0x0b, 0x3a, 0x63, 0x84, 0xbf, 0x5c, 0x80, 0x1b, 0xcf, 0x4a, 0xb5, 0x85, 0xb0, 0x79, 0x7c, 0xff,
	0xbc, 0xfd, 0x2b, 0xfb, 0xd5, 0x33, 0x87, 0xf8, 0xeb, 0xd8, 0x1d, 0x43, 0xbc, 0x84, 0x61, 0x3f,
	0x51, 0x1c, 0x43, 0xf0, 0x47, 0x76, 0xfe, 0x5c, 0x4c, 0x68, 0x56, 0xf3, 0x37, 0xcb, 0x5b, 0x69,
	0x8f, 0xe1, 0x48, 0x38, 0xf1, 0xf9, 0xe8, 0x8e, 0x7c, 0xb8, 0x84, 0xd0, 0xb4, 0xc3, 0x08, 0x82,
	0xaf, 0x8f, 0x8f, 0xe3, 0x17, 0x38, 0x82, 0xe8, 0x3e, 0x15, 0x3f, 0xef, 0xd3, 0x1f, 0x63, 0xb2,
	0x3e, 0xb6, 0x6f, 0xbf, 0xf9, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x48, 0x4b, 0x00, 0xa8, 0x39, 0x03,
	0x00, 0x00,
}
