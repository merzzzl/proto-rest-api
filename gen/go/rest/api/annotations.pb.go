// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: rest/api/annotations.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MethodRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP method used to bind this RPC.
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// The path pattern associated with this RPC.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// The name of the request field whose value is mapped to the HTTP body, or `*` for mapping all fields not captured by the path pattern to the HTTP body.
	Request string `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	// The name of the response field whose value is mapped to the HTTP body of response. Other response fields are ignored. When not set, the response message will be used as HTTP body of response.
	Response string `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	// The HTTP status code used for successful responses. Defaults to 200.
	SuccessCode int32 `protobuf:"varint,5,opt,name=success_code,json=successCode,proto3" json:"success_code,omitempty"`
}

func (x *MethodRule) Reset() {
	*x = MethodRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rest_api_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodRule) ProtoMessage() {}

func (x *MethodRule) ProtoReflect() protoreflect.Message {
	mi := &file_rest_api_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodRule.ProtoReflect.Descriptor instead.
func (*MethodRule) Descriptor() ([]byte, []int) {
	return file_rest_api_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *MethodRule) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MethodRule) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MethodRule) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *MethodRule) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *MethodRule) GetSuccessCode() int32 {
	if x != nil {
		return x.SuccessCode
	}
	return 0
}

type ServiceRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base path for all HTTP bindings in this service.
	BasePath string `protobuf:"bytes,1,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
}

func (x *ServiceRule) Reset() {
	*x = ServiceRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rest_api_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRule) ProtoMessage() {}

func (x *ServiceRule) ProtoReflect() protoreflect.Message {
	mi := &file_rest_api_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRule.ProtoReflect.Descriptor instead.
func (*ServiceRule) Descriptor() ([]byte, []int) {
	return file_rest_api_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceRule) GetBasePath() string {
	if x != nil {
		return x.BasePath
	}
	return ""
}

var file_rest_api_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodRule)(nil),
		Field:         70001,
		Name:          "rest.api.method",
		Tag:           "bytes,70001,opt,name=method",
		Filename:      "rest/api/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceRule)(nil),
		Field:         70002,
		Name:          "rest.api.service",
		Tag:           "bytes,70002,opt,name=service",
		Filename:      "rest/api/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional rest.api.MethodRule method = 70001;
	E_Method = &file_rest_api_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional rest.api.ServiceRule service = 70002;
	E_Service = &file_rest_api_annotations_proto_extTypes[1]
)

var File_rest_api_annotations_proto protoreflect.FileDescriptor

var file_rest_api_annotations_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65,
	0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2a, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x3a, 0x4e, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xf1, 0xa2, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x52, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf2, 0xa2, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x72, 0x7a, 0x7a,
	0x7a, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rest_api_annotations_proto_rawDescOnce sync.Once
	file_rest_api_annotations_proto_rawDescData = file_rest_api_annotations_proto_rawDesc
)

func file_rest_api_annotations_proto_rawDescGZIP() []byte {
	file_rest_api_annotations_proto_rawDescOnce.Do(func() {
		file_rest_api_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_rest_api_annotations_proto_rawDescData)
	})
	return file_rest_api_annotations_proto_rawDescData
}

var file_rest_api_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rest_api_annotations_proto_goTypes = []interface{}{
	(*MethodRule)(nil),                  // 0: rest.api.MethodRule
	(*ServiceRule)(nil),                 // 1: rest.api.ServiceRule
	(*descriptorpb.MethodOptions)(nil),  // 2: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
}
var file_rest_api_annotations_proto_depIdxs = []int32{
	2, // 0: rest.api.method:extendee -> google.protobuf.MethodOptions
	3, // 1: rest.api.service:extendee -> google.protobuf.ServiceOptions
	0, // 2: rest.api.method:type_name -> rest.api.MethodRule
	1, // 3: rest.api.service:type_name -> rest.api.ServiceRule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rest_api_annotations_proto_init() }
func file_rest_api_annotations_proto_init() {
	if File_rest_api_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rest_api_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodRule); i {
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
		file_rest_api_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRule); i {
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
			RawDescriptor: file_rest_api_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_rest_api_annotations_proto_goTypes,
		DependencyIndexes: file_rest_api_annotations_proto_depIdxs,
		MessageInfos:      file_rest_api_annotations_proto_msgTypes,
		ExtensionInfos:    file_rest_api_annotations_proto_extTypes,
	}.Build()
	File_rest_api_annotations_proto = out.File
	file_rest_api_annotations_proto_rawDesc = nil
	file_rest_api_annotations_proto_goTypes = nil
	file_rest_api_annotations_proto_depIdxs = nil
}
