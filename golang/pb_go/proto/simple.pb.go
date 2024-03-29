// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.13.0
// source: proto/simple.proto

package simple

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request
type SimpleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SimpleRequest) Reset() {
	*x = SimpleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRequest) ProtoMessage() {}

func (x *SimpleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRequest.ProtoReflect.Descriptor instead.
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return file_proto_simple_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// Response
type SimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SimpleResponse) Reset() {
	*x = SimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleResponse) ProtoMessage() {}

func (x *SimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleResponse.ProtoReflect.Descriptor instead.
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return file_proto_simple_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_proto_simple_proto protoreflect.FileDescriptor

var file_proto_simple_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x24, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x3e, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x6d, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x3b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_simple_proto_rawDescOnce sync.Once
	file_proto_simple_proto_rawDescData = file_proto_simple_proto_rawDesc
)

func file_proto_simple_proto_rawDescGZIP() []byte {
	file_proto_simple_proto_rawDescOnce.Do(func() {
		file_proto_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_simple_proto_rawDescData)
	})
	return file_proto_simple_proto_rawDescData
}

var file_proto_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_simple_proto_goTypes = []interface{}{
	(*SimpleRequest)(nil),  // 0: main.SimpleRequest
	(*SimpleResponse)(nil), // 1: main.SimpleResponse
}
var file_proto_simple_proto_depIdxs = []int32{
	0, // 0: main.Simple.GetLine:input_type -> main.SimpleRequest
	1, // 1: main.Simple.GetLine:output_type -> main.SimpleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_simple_proto_init() }
func file_proto_simple_proto_init() {
	if File_proto_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRequest); i {
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
		file_proto_simple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleResponse); i {
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
			RawDescriptor: file_proto_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_simple_proto_goTypes,
		DependencyIndexes: file_proto_simple_proto_depIdxs,
		MessageInfos:      file_proto_simple_proto_msgTypes,
	}.Build()
	File_proto_simple_proto = out.File
	file_proto_simple_proto_rawDesc = nil
	file_proto_simple_proto_goTypes = nil
	file_proto_simple_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleClient interface {
	GetLine(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type simpleClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleClient(cc grpc.ClientConnInterface) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) GetLine(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/main.Simple/GetLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
type SimpleServer interface {
	GetLine(context.Context, *SimpleRequest) (*SimpleResponse, error)
}

// UnimplementedSimpleServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleServer struct {
}

func (*UnimplementedSimpleServer) GetLine(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLine not implemented")
}

func RegisterSimpleServer(s *grpc.Server, srv SimpleServer) {
	s.RegisterService(&_Simple_serviceDesc, srv)
}

func _Simple_GetLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).GetLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Simple/GetLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).GetLine(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Simple_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLine",
			Handler:    _Simple_GetLine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/simple.proto",
}
