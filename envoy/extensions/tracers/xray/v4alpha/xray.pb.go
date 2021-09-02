// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/tracers/xray/v4alpha/xray.proto

package envoy_extensions_tracers_xray_v4alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v4alpha "github.com/malmanzor/go-control-plane/envoy/config/core/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// [#extension: envoy.tracers.xray]
type XRayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The UDP endpoint of the X-Ray Daemon where the spans will be sent.
	// If this value is not set, the default value of 127.0.0.1:2000 will be used.
	DaemonEndpoint *v4alpha.SocketAddress `protobuf:"bytes,1,opt,name=daemon_endpoint,json=daemonEndpoint,proto3" json:"daemon_endpoint,omitempty"`
	// The name of the X-Ray segment.
	SegmentName string `protobuf:"bytes,2,opt,name=segment_name,json=segmentName,proto3" json:"segment_name,omitempty"`
	// The location of a local custom sampling rules JSON file.
	// For an example of the sampling rules see:
	// `X-Ray SDK documentation
	// <https://docs.aws.amazon.com/xray/latest/devguide/xray-sdk-go-configuration.html#xray-sdk-go-configuration-sampling>`_
	SamplingRuleManifest *v4alpha.DataSource `protobuf:"bytes,3,opt,name=sampling_rule_manifest,json=samplingRuleManifest,proto3" json:"sampling_rule_manifest,omitempty"`
	// Optional custom fields to be added to each trace segment.
	// see: `X-Ray Segment Document documentation
	// <https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html>`__
	SegmentFields *XRayConfig_SegmentFields `protobuf:"bytes,4,opt,name=segment_fields,json=segmentFields,proto3" json:"segment_fields,omitempty"`
}

func (x *XRayConfig) Reset() {
	*x = XRayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XRayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XRayConfig) ProtoMessage() {}

func (x *XRayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XRayConfig.ProtoReflect.Descriptor instead.
func (*XRayConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescGZIP(), []int{0}
}

func (x *XRayConfig) GetDaemonEndpoint() *v4alpha.SocketAddress {
	if x != nil {
		return x.DaemonEndpoint
	}
	return nil
}

func (x *XRayConfig) GetSegmentName() string {
	if x != nil {
		return x.SegmentName
	}
	return ""
}

func (x *XRayConfig) GetSamplingRuleManifest() *v4alpha.DataSource {
	if x != nil {
		return x.SamplingRuleManifest
	}
	return nil
}

func (x *XRayConfig) GetSegmentFields() *XRayConfig_SegmentFields {
	if x != nil {
		return x.SegmentFields
	}
	return nil
}

type XRayConfig_SegmentFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of AWS resource, e.g. "AWS::AppMesh::Proxy".
	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	// AWS resource metadata dictionary.
	// See: `X-Ray Segment Document documentation <https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html#api-segmentdocuments-aws>`__
	Aws *_struct.Struct `protobuf:"bytes,2,opt,name=aws,proto3" json:"aws,omitempty"`
}

func (x *XRayConfig_SegmentFields) Reset() {
	*x = XRayConfig_SegmentFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XRayConfig_SegmentFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XRayConfig_SegmentFields) ProtoMessage() {}

func (x *XRayConfig_SegmentFields) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XRayConfig_SegmentFields.ProtoReflect.Descriptor instead.
func (*XRayConfig_SegmentFields) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescGZIP(), []int{0, 0}
}

func (x *XRayConfig_SegmentFields) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *XRayConfig_SegmentFields) GetAws() *_struct.Struct {
	if x != nil {
		return x.Aws
	}
	return nil
}

var File_envoy_extensions_tracers_xray_v4alpha_xray_proto protoreflect.FileDescriptor

var file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2f,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x85, 0x04, 0x0a, 0x0a, 0x58, 0x52, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x51, 0x0a, 0x0f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x5b, 0x0a, 0x16, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x14, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x52, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x0e,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e,
	0x78, 0x72, 0x61, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x58, 0x52, 0x61,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x1a, 0x89, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x29,
	0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x03, 0x61, 0x77, 0x73, 0x3a, 0x35, 0x9a, 0xc5, 0x88, 0x1e, 0x30,
	0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x58, 0x52, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x3a, 0x27, 0x9a, 0xc5, 0x88, 0x1e, 0x22, 0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x58,
	0x52, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x4a, 0x0a, 0x33, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x73, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x09, 0x58, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescOnce sync.Once
	file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescData = file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDesc
)

func file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescGZIP() []byte {
	file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescData)
	})
	return file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDescData
}

var file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_tracers_xray_v4alpha_xray_proto_goTypes = []interface{}{
	(*XRayConfig)(nil),               // 0: envoy.extensions.tracers.xray.v4alpha.XRayConfig
	(*XRayConfig_SegmentFields)(nil), // 1: envoy.extensions.tracers.xray.v4alpha.XRayConfig.SegmentFields
	(*v4alpha.SocketAddress)(nil),    // 2: envoy.config.core.v4alpha.SocketAddress
	(*v4alpha.DataSource)(nil),       // 3: envoy.config.core.v4alpha.DataSource
	(*_struct.Struct)(nil),           // 4: google.protobuf.Struct
}
var file_envoy_extensions_tracers_xray_v4alpha_xray_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.tracers.xray.v4alpha.XRayConfig.daemon_endpoint:type_name -> envoy.config.core.v4alpha.SocketAddress
	3, // 1: envoy.extensions.tracers.xray.v4alpha.XRayConfig.sampling_rule_manifest:type_name -> envoy.config.core.v4alpha.DataSource
	1, // 2: envoy.extensions.tracers.xray.v4alpha.XRayConfig.segment_fields:type_name -> envoy.extensions.tracers.xray.v4alpha.XRayConfig.SegmentFields
	4, // 3: envoy.extensions.tracers.xray.v4alpha.XRayConfig.SegmentFields.aws:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_extensions_tracers_xray_v4alpha_xray_proto_init() }
func file_envoy_extensions_tracers_xray_v4alpha_xray_proto_init() {
	if File_envoy_extensions_tracers_xray_v4alpha_xray_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XRayConfig); i {
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
		file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XRayConfig_SegmentFields); i {
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
			RawDescriptor: file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_tracers_xray_v4alpha_xray_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_tracers_xray_v4alpha_xray_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_tracers_xray_v4alpha_xray_proto_msgTypes,
	}.Build()
	File_envoy_extensions_tracers_xray_v4alpha_xray_proto = out.File
	file_envoy_extensions_tracers_xray_v4alpha_xray_proto_rawDesc = nil
	file_envoy_extensions_tracers_xray_v4alpha_xray_proto_goTypes = nil
	file_envoy_extensions_tracers_xray_v4alpha_xray_proto_depIdxs = nil
}
