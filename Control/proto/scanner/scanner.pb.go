// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: scanner.proto

package scanner

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

type ResourceRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port        int64  `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	ServiceType string `protobuf:"bytes,2,opt,name=ServiceType,proto3" json:"ServiceType,omitempty"`
}

func (x *ResourceRegister) Reset() {
	*x = ResourceRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRegister) ProtoMessage() {}

func (x *ResourceRegister) ProtoReflect() protoreflect.Message {
	mi := &file_scanner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRegister.ProtoReflect.Descriptor instead.
func (*ResourceRegister) Descriptor() ([]byte, []int) {
	return file_scanner_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceRegister) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ResourceRegister) GetServiceType() string {
	if x != nil {
		return x.ServiceType
	}
	return ""
}

type ResourceRegisterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ResourceRegisterResult) Reset() {
	*x = ResourceRegisterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRegisterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRegisterResult) ProtoMessage() {}

func (x *ResourceRegisterResult) ProtoReflect() protoreflect.Message {
	mi := &file_scanner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRegisterResult.ProtoReflect.Descriptor instead.
func (*ResourceRegisterResult) Descriptor() ([]byte, []int) {
	return file_scanner_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceRegisterResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_scanner_proto protoreflect.FileDescriptor

var file_scanner_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x30, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x63, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x48, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_scanner_proto_rawDescOnce sync.Once
	file_scanner_proto_rawDescData = file_scanner_proto_rawDesc
)

func file_scanner_proto_rawDescGZIP() []byte {
	file_scanner_proto_rawDescOnce.Do(func() {
		file_scanner_proto_rawDescData = protoimpl.X.CompressGZIP(file_scanner_proto_rawDescData)
	})
	return file_scanner_proto_rawDescData
}

var file_scanner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scanner_proto_goTypes = []interface{}{
	(*ResourceRegister)(nil),       // 0: scanner.ResourceRegister
	(*ResourceRegisterResult)(nil), // 1: scanner.ResourceRegisterResult
}
var file_scanner_proto_depIdxs = []int32{
	0, // 0: scanner.ResourceRegisterService.register:input_type -> scanner.ResourceRegister
	1, // 1: scanner.ResourceRegisterService.register:output_type -> scanner.ResourceRegisterResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scanner_proto_init() }
func file_scanner_proto_init() {
	if File_scanner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scanner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRegister); i {
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
		file_scanner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRegisterResult); i {
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
			RawDescriptor: file_scanner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanner_proto_goTypes,
		DependencyIndexes: file_scanner_proto_depIdxs,
		MessageInfos:      file_scanner_proto_msgTypes,
	}.Build()
	File_scanner_proto = out.File
	file_scanner_proto_rawDesc = nil
	file_scanner_proto_goTypes = nil
	file_scanner_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResourceRegisterServiceClient is the client API for ResourceRegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceRegisterServiceClient interface {
	Register(ctx context.Context, in *ResourceRegister, opts ...grpc.CallOption) (*ResourceRegisterResult, error)
}

type resourceRegisterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceRegisterServiceClient(cc grpc.ClientConnInterface) ResourceRegisterServiceClient {
	return &resourceRegisterServiceClient{cc}
}

func (c *resourceRegisterServiceClient) Register(ctx context.Context, in *ResourceRegister, opts ...grpc.CallOption) (*ResourceRegisterResult, error) {
	out := new(ResourceRegisterResult)
	err := c.cc.Invoke(ctx, "/scanner.ResourceRegisterService/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceRegisterServiceServer is the server API for ResourceRegisterService service.
type ResourceRegisterServiceServer interface {
	Register(context.Context, *ResourceRegister) (*ResourceRegisterResult, error)
}

// UnimplementedResourceRegisterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResourceRegisterServiceServer struct {
}

func (*UnimplementedResourceRegisterServiceServer) Register(context.Context, *ResourceRegister) (*ResourceRegisterResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterResourceRegisterServiceServer(s *grpc.Server, srv ResourceRegisterServiceServer) {
	s.RegisterService(&_ResourceRegisterService_serviceDesc, srv)
}

func _ResourceRegisterService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceRegisterServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanner.ResourceRegisterService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceRegisterServiceServer).Register(ctx, req.(*ResourceRegister))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceRegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scanner.ResourceRegisterService",
	HandlerType: (*ResourceRegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _ResourceRegisterService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanner.proto",
}
