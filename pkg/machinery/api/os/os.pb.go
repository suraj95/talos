// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: os/os.proto

package os

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	common "github.com/talos-systems/talos/pkg/machinery/api/common"
	machine "github.com/talos-systems/talos/pkg/machinery/api/machine"
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

var File_os_os_proto protoreflect.FileDescriptor

var file_os_os_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x73, 0x2f, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x6f,
	0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf9, 0x02, 0x0a, 0x09, 0x4f,
	0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x05, 0x44, 0x6d, 0x65, 0x73, 0x67, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2e, 0x44, 0x6d, 0x65, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x12,
	0x39, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1a, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x03, 0x88, 0x02, 0x01, 0x42, 0x4a, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x42, 0x05, 0x4f, 0x73, 0x41, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_os_os_proto_goTypes = []interface{}{
	(*machine.ContainersRequest)(nil),  // 0: machine.ContainersRequest
	(*machine.DmesgRequest)(nil),       // 1: machine.DmesgRequest
	(*empty.Empty)(nil),                // 2: google.protobuf.Empty
	(*machine.RestartRequest)(nil),     // 3: machine.RestartRequest
	(*machine.StatsRequest)(nil),       // 4: machine.StatsRequest
	(*machine.ContainersResponse)(nil), // 5: machine.ContainersResponse
	(*common.Data)(nil),                // 6: common.Data
	(*machine.MemoryResponse)(nil),     // 7: machine.MemoryResponse
	(*machine.ProcessesResponse)(nil),  // 8: machine.ProcessesResponse
	(*machine.RestartResponse)(nil),    // 9: machine.RestartResponse
	(*machine.StatsResponse)(nil),      // 10: machine.StatsResponse
}

var file_os_os_proto_depIdxs = []int32{
	0,  // 0: os.OSService.Containers:input_type -> machine.ContainersRequest
	1,  // 1: os.OSService.Dmesg:input_type -> machine.DmesgRequest
	2,  // 2: os.OSService.Memory:input_type -> google.protobuf.Empty
	2,  // 3: os.OSService.Processes:input_type -> google.protobuf.Empty
	3,  // 4: os.OSService.Restart:input_type -> machine.RestartRequest
	4,  // 5: os.OSService.Stats:input_type -> machine.StatsRequest
	5,  // 6: os.OSService.Containers:output_type -> machine.ContainersResponse
	6,  // 7: os.OSService.Dmesg:output_type -> common.Data
	7,  // 8: os.OSService.Memory:output_type -> machine.MemoryResponse
	8,  // 9: os.OSService.Processes:output_type -> machine.ProcessesResponse
	9,  // 10: os.OSService.Restart:output_type -> machine.RestartResponse
	10, // 11: os.OSService.Stats:output_type -> machine.StatsResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_os_os_proto_init() }
func file_os_os_proto_init() {
	if File_os_os_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_os_os_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_os_os_proto_goTypes,
		DependencyIndexes: file_os_os_proto_depIdxs,
	}.Build()
	File_os_os_proto = out.File
	file_os_os_proto_rawDesc = nil
	file_os_os_proto_goTypes = nil
	file_os_os_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConnInterface
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OSServiceClient is the client API for OSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type OSServiceClient interface {
	Containers(ctx context.Context, in *machine.ContainersRequest, opts ...grpc.CallOption) (*machine.ContainersResponse, error)
	Dmesg(ctx context.Context, in *machine.DmesgRequest, opts ...grpc.CallOption) (OSService_DmesgClient, error)
	Memory(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*machine.MemoryResponse, error)
	Processes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*machine.ProcessesResponse, error)
	Restart(ctx context.Context, in *machine.RestartRequest, opts ...grpc.CallOption) (*machine.RestartResponse, error)
	Stats(ctx context.Context, in *machine.StatsRequest, opts ...grpc.CallOption) (*machine.StatsResponse, error)
}

type oSServiceClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewOSServiceClient(cc grpc.ClientConnInterface) OSServiceClient {
	return &oSServiceClient{cc}
}

func (c *oSServiceClient) Containers(ctx context.Context, in *machine.ContainersRequest, opts ...grpc.CallOption) (*machine.ContainersResponse, error) {
	out := new(machine.ContainersResponse)
	err := c.cc.Invoke(ctx, "/os.OSService/Containers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSServiceClient) Dmesg(ctx context.Context, in *machine.DmesgRequest, opts ...grpc.CallOption) (OSService_DmesgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OSService_serviceDesc.Streams[0], "/os.OSService/Dmesg", opts...)
	if err != nil {
		return nil, err
	}
	x := &oSServiceDmesgClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OSService_DmesgClient interface {
	Recv() (*common.Data, error)
	grpc.ClientStream
}

type oSServiceDmesgClient struct {
	grpc.ClientStream
}

func (x *oSServiceDmesgClient) Recv() (*common.Data, error) {
	m := new(common.Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oSServiceClient) Memory(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*machine.MemoryResponse, error) {
	out := new(machine.MemoryResponse)
	err := c.cc.Invoke(ctx, "/os.OSService/Memory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSServiceClient) Processes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*machine.ProcessesResponse, error) {
	out := new(machine.ProcessesResponse)
	err := c.cc.Invoke(ctx, "/os.OSService/Processes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSServiceClient) Restart(ctx context.Context, in *machine.RestartRequest, opts ...grpc.CallOption) (*machine.RestartResponse, error) {
	out := new(machine.RestartResponse)
	err := c.cc.Invoke(ctx, "/os.OSService/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSServiceClient) Stats(ctx context.Context, in *machine.StatsRequest, opts ...grpc.CallOption) (*machine.StatsResponse, error) {
	out := new(machine.StatsResponse)
	err := c.cc.Invoke(ctx, "/os.OSService/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OSServiceServer is the server API for OSService service.
//
// Deprecated: Do not use.
type OSServiceServer interface {
	Containers(context.Context, *machine.ContainersRequest) (*machine.ContainersResponse, error)
	Dmesg(*machine.DmesgRequest, OSService_DmesgServer) error
	Memory(context.Context, *empty.Empty) (*machine.MemoryResponse, error)
	Processes(context.Context, *empty.Empty) (*machine.ProcessesResponse, error)
	Restart(context.Context, *machine.RestartRequest) (*machine.RestartResponse, error)
	Stats(context.Context, *machine.StatsRequest) (*machine.StatsResponse, error)
}

// UnimplementedOSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOSServiceServer struct {
}

func (*UnimplementedOSServiceServer) Containers(context.Context, *machine.ContainersRequest) (*machine.ContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Containers not implemented")
}

func (*UnimplementedOSServiceServer) Dmesg(*machine.DmesgRequest, OSService_DmesgServer) error {
	return status.Errorf(codes.Unimplemented, "method Dmesg not implemented")
}

func (*UnimplementedOSServiceServer) Memory(context.Context, *empty.Empty) (*machine.MemoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Memory not implemented")
}

func (*UnimplementedOSServiceServer) Processes(context.Context, *empty.Empty) (*machine.ProcessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Processes not implemented")
}

func (*UnimplementedOSServiceServer) Restart(context.Context, *machine.RestartRequest) (*machine.RestartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}

func (*UnimplementedOSServiceServer) Stats(context.Context, *machine.StatsRequest) (*machine.StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}

// Deprecated: Do not use.
func RegisterOSServiceServer(s *grpc.Server, srv OSServiceServer) {
	s.RegisterService(&_OSService_serviceDesc, srv)
}

func _OSService_Containers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(machine.ContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServiceServer).Containers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/os.OSService/Containers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServiceServer).Containers(ctx, req.(*machine.ContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSService_Dmesg_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(machine.DmesgRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OSServiceServer).Dmesg(m, &oSServiceDmesgServer{stream})
}

type OSService_DmesgServer interface {
	Send(*common.Data) error
	grpc.ServerStream
}

type oSServiceDmesgServer struct {
	grpc.ServerStream
}

func (x *oSServiceDmesgServer) Send(m *common.Data) error {
	return x.ServerStream.SendMsg(m)
}

func _OSService_Memory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServiceServer).Memory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/os.OSService/Memory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServiceServer).Memory(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSService_Processes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServiceServer).Processes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/os.OSService/Processes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServiceServer).Processes(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSService_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(machine.RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServiceServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/os.OSService/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServiceServer).Restart(ctx, req.(*machine.RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSService_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(machine.StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServiceServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/os.OSService/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServiceServer).Stats(ctx, req.(*machine.StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "os.OSService",
	HandlerType: (*OSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Containers",
			Handler:    _OSService_Containers_Handler,
		},
		{
			MethodName: "Memory",
			Handler:    _OSService_Memory_Handler,
		},
		{
			MethodName: "Processes",
			Handler:    _OSService_Processes_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _OSService_Restart_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _OSService_Stats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Dmesg",
			Handler:       _OSService_Dmesg_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "os/os.proto",
}
