// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/aggregate/v3alpha/cluster.proto

package envoy_config_cluster_aggregate_v3alpha

import (
	fmt "fmt"
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

type ClusterConfig struct {
	Clusters             []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa65a3002ca6d323, []int{0}
}

func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfig.Unmarshal(m, b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
}
func (m *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(m, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return xxx_messageInfo_ClusterConfig.Size(m)
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterConfig)(nil), "envoy.config.cluster.aggregate.v3alpha.ClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/aggregate/v3alpha/cluster.proto", fileDescriptor_fa65a3002ca6d323)
}

var fileDescriptor_fa65a3002ca6d323 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d,
	0xd2, 0x4f, 0x4c, 0x4f, 0x2f, 0x4a, 0x4d, 0x4f, 0x2c, 0x49, 0xd5, 0x2f, 0x33, 0x4e, 0xcc, 0x29,
	0xc8, 0x48, 0x84, 0xc9, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xa9, 0x81, 0x75, 0xe9, 0x41,
	0x74, 0xe9, 0xc1, 0xe4, 0xe0, 0xba, 0xf4, 0xa0, 0xba, 0xa4, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53,
	0xc0, 0xe6, 0x40, 0x19, 0x10, 0x03, 0x94, 0x4c, 0xb9, 0x78, 0x9d, 0x21, 0xba, 0x9c, 0xc1, 0x66,
	0x08, 0xa9, 0x70, 0x71, 0x40, 0x8d, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x74, 0xe2, 0xf8,
	0xe5, 0xc4, 0x3a, 0x89, 0x91, 0x89, 0x83, 0x31, 0x08, 0x2e, 0xe3, 0xe4, 0xc6, 0x65, 0x92, 0x99,
	0xaf, 0x07, 0xb6, 0xbc, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x8f, 0x38, 0x77, 0x38, 0xf1, 0x40, 0x2d,
	0x0b, 0x00, 0x59, 0x1e, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0x85, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0x50, 0x1c, 0xf2, 0xfe, 0x00, 0x00, 0x00,
}
