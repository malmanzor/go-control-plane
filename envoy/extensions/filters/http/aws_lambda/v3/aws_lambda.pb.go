// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto

package envoy_extensions_filters_http_aws_lambda_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Config_InvocationMode int32

const (
	Config_SYNCHRONOUS  Config_InvocationMode = 0
	Config_ASYNCHRONOUS Config_InvocationMode = 1
)

var Config_InvocationMode_name = map[int32]string{
	0: "SYNCHRONOUS",
	1: "ASYNCHRONOUS",
}

var Config_InvocationMode_value = map[string]int32{
	"SYNCHRONOUS":  0,
	"ASYNCHRONOUS": 1,
}

func (x Config_InvocationMode) String() string {
	return proto.EnumName(Config_InvocationMode_name, int32(x))
}

func (Config_InvocationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a22f80ea45c04127, []int{0, 0}
}

type Config struct {
	Arn                  string                `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	PayloadPassthrough   bool                  `protobuf:"varint,2,opt,name=payload_passthrough,json=payloadPassthrough,proto3" json:"payload_passthrough,omitempty"`
	InvocationMode       Config_InvocationMode `protobuf:"varint,3,opt,name=invocation_mode,json=invocationMode,proto3,enum=envoy.extensions.filters.http.aws_lambda.v3.Config_InvocationMode" json:"invocation_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_a22f80ea45c04127, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetArn() string {
	if m != nil {
		return m.Arn
	}
	return ""
}

func (m *Config) GetPayloadPassthrough() bool {
	if m != nil {
		return m.PayloadPassthrough
	}
	return false
}

func (m *Config) GetInvocationMode() Config_InvocationMode {
	if m != nil {
		return m.InvocationMode
	}
	return Config_SYNCHRONOUS
}

type PerRouteConfig struct {
	InvokeConfig         *Config  `protobuf:"bytes,1,opt,name=invoke_config,json=invokeConfig,proto3" json:"invoke_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a22f80ea45c04127, []int{1}
}

func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerRouteConfig.Unmarshal(m, b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return xxx_messageInfo_PerRouteConfig.Size(m)
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

func (m *PerRouteConfig) GetInvokeConfig() *Config {
	if m != nil {
		return m.InvokeConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.http.aws_lambda.v3.Config_InvocationMode", Config_InvocationMode_name, Config_InvocationMode_value)
	proto.RegisterType((*Config)(nil), "envoy.extensions.filters.http.aws_lambda.v3.Config")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.extensions.filters.http.aws_lambda.v3.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto", fileDescriptor_a22f80ea45c04127)
}

var fileDescriptor_a22f80ea45c04127 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x8e, 0xd3, 0x30,
	0x18, 0xc7, 0x71, 0x0e, 0xca, 0xe1, 0x3b, 0x72, 0x51, 0x18, 0x28, 0x27, 0x81, 0x4a, 0xa7, 0x4a,
	0x48, 0xb6, 0x94, 0xb0, 0x5c, 0x75, 0x4b, 0x73, 0x0b, 0x48, 0x70, 0x8d, 0x52, 0x55, 0x82, 0x29,
	0x72, 0x1b, 0xb7, 0xb1, 0x48, 0xed, 0xc8, 0x76, 0xd2, 0x76, 0x43, 0x4c, 0x3c, 0x03, 0xef, 0xc0,
	0x0b, 0xc0, 0x8c, 0xc4, 0xca, 0xeb, 0x74, 0x42, 0x89, 0x03, 0x6d, 0xd5, 0xa9, 0xb7, 0xf9, 0xf3,
	0xff, 0xfb, 0xff, 0xfd, 0xf3, 0xa7, 0x0f, 0x5e, 0x53, 0x5e, 0x8a, 0x35, 0xa6, 0x2b, 0x4d, 0xb9,
	0x62, 0x82, 0x2b, 0x3c, 0x63, 0x99, 0xa6, 0x52, 0xe1, 0x54, 0xeb, 0x1c, 0x93, 0xa5, 0x8a, 0x33,
	0xb2, 0x98, 0x24, 0x04, 0x97, 0xfe, 0x4e, 0x85, 0x72, 0x29, 0xb4, 0x70, 0x5f, 0xd5, 0x6e, 0xb4,
	0x75, 0xa3, 0xc6, 0x8d, 0x2a, 0x37, 0xda, 0xe9, 0x2f, 0xfd, 0xcb, 0xe7, 0x45, 0x92, 0x13, 0x4c,
	0x38, 0x17, 0x9a, 0xe8, 0xfa, 0x29, 0xa5, 0x89, 0x2e, 0x94, 0xc9, 0xba, 0x7c, 0x79, 0x20, 0x97,
	0x54, 0x56, 0xa1, 0x8c, 0xcf, 0x9b, 0x96, 0xa7, 0x25, 0xc9, 0x58, 0x42, 0x34, 0xc5, 0xff, 0x0e,
	0x46, 0xe8, 0xfe, 0xb4, 0x60, 0xeb, 0x46, 0xf0, 0x19, 0x9b, 0xbb, 0xcf, 0xe0, 0x09, 0x91, 0xbc,
	0x0d, 0x3a, 0xa0, 0xf7, 0x28, 0x78, 0xb8, 0x09, 0xee, 0x4b, 0xcb, 0x01, 0x51, 0x75, 0xe7, 0x62,
	0xf8, 0x24, 0x27, 0xeb, 0x4c, 0x90, 0x24, 0xce, 0x89, 0x52, 0x3a, 0x95, 0xa2, 0x98, 0xa7, 0x6d,
	0xab, 0x03, 0x7a, 0xa7, 0x91, 0xdb, 0x48, 0xe1, 0x56, 0x71, 0x0b, 0x78, 0xc1, 0x78, 0x29, 0xa6,
	0x35, 0x4f, 0xbc, 0x10, 0x09, 0x6d, 0x9f, 0x74, 0x40, 0xcf, 0xf6, 0x02, 0x74, 0xc4, 0xc7, 0x91,
	0x21, 0x43, 0x6f, 0xff, 0x47, 0xbd, 0x17, 0x09, 0x0d, 0x4e, 0x37, 0xc1, 0x83, 0x2f, 0xa0, 0x82,
	0xb3, 0xd9, 0x9e, 0xd2, 0xf5, 0xa1, 0xbd, 0xdf, 0xeb, 0x5e, 0xc0, 0xb3, 0xd1, 0xc7, 0xdb, 0x9b,
	0x37, 0xd1, 0xf0, 0x76, 0x38, 0x1e, 0x39, 0xf7, 0x5c, 0x07, 0x9e, 0x0f, 0x76, 0x6f, 0x40, 0xff,
	0xea, 0xdb, 0xaf, 0xaf, 0x2f, 0x5e, 0x43, 0xcf, 0x80, 0x4d, 0xcd, 0x93, 0x06, 0xea, 0x90, 0xc9,
	0x23, 0x59, 0x9e, 0x92, 0x06, 0xac, 0xfb, 0x1d, 0x40, 0x3b, 0xa4, 0x32, 0x12, 0x85, 0xa6, 0xcd,
	0x14, 0x3f, 0xc0, 0xc7, 0x15, 0xd4, 0x27, 0x1a, 0x9b, 0xa4, 0x7a, 0x9e, 0x67, 0x9e, 0x7f, 0x87,
	0x7f, 0x47, 0xe7, 0x26, 0xc9, 0x54, 0xfd, 0x41, 0xc5, 0x79, 0x0d, 0xfb, 0xc7, 0x70, 0xee, 0xc3,
	0x05, 0xe3, 0x1f, 0x9f, 0x7f, 0xff, 0x69, 0x59, 0x8e, 0x05, 0xaf, 0x98, 0x30, 0x44, 0xb9, 0x14,
	0xab, 0xf5, 0x31, 0x70, 0x81, 0x3d, 0x58, 0xaa, 0x77, 0x75, 0x15, 0x56, 0x2b, 0x14, 0x82, 0x49,
	0xab, 0xde, 0x25, 0xff, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0xaf, 0x9e, 0x68, 0x13, 0x03,
	0x00, 0x00,
}