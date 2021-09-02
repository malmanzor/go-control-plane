// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/filters/http/alternate_protocols_cache/v3/alternate_protocols_cache.proto

package envoy_extensions_filters_http_alternate_protocols_cache_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/malmanzor/go-control-plane/envoy/config/core/v3"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the alternate protocols cache HTTP filter.
// [#extension: envoy.filters.http.alternate_protocols_cache]
// TODO(RyanTheOptimist): Move content from source/docs/http3_upstream.md to
// docs/root/intro/arch_overview/upstream/connection_pooling.rst when unhiding the proto.
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [#not-implemented-hide:]
	// If set, causes the use of the alternate protocols cache, which is responsible for
	// parsing and caching HTTP Alt-Svc headers. This enables the use of HTTP/3 for upstream
	// servers that advertise supporting it.
	// TODO(RyanTheOptimist): Make this field required when HTTP/3 is enabled via auto_http.
	AlternateProtocolsCacheOptions *v3.AlternateProtocolsCacheOptions `protobuf:"bytes,1,opt,name=alternate_protocols_cache_options,json=alternateProtocolsCacheOptions,proto3" json:"alternate_protocols_cache_options,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetAlternateProtocolsCacheOptions() *v3.AlternateProtocolsCacheOptions {
	if x != nil {
		return x.AlternateProtocolsCacheOptions
	}
	return nil
}

var File_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDesc = []byte{
	0x0a, 0x5a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x6c, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x7f, 0x0a,
	0x21, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x1e,
	0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x72,
	0x0a, 0x48, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x6c,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x1c, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescData = file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDesc
)

func file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDescData
}

var file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_goTypes = []interface{}{
	(*FilterConfig)(nil),                      // 0: envoy.extensions.filters.http.alternate_protocols_cache.v3.FilterConfig
	(*v3.AlternateProtocolsCacheOptions)(nil), // 1: envoy.config.core.v3.AlternateProtocolsCacheOptions
}
var file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.alternate_protocols_cache.v3.FilterConfig.alternate_protocols_cache_options:type_name -> envoy.config.core.v3.AlternateProtocolsCacheOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_init()
}
func file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_init() {
	if File_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto = out.File
	file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_rawDesc = nil
	file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_goTypes = nil
	file_envoy_extensions_filters_http_alternate_protocols_cache_v3_alternate_protocols_cache_proto_depIdxs = nil
}
