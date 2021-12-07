// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: probuffile.proto

package protobuff

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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probuffile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_probuffile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_probuffile_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probuffile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_probuffile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_probuffile_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type InputNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *InputNum) Reset() {
	*x = InputNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probuffile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputNum) ProtoMessage() {}

func (x *InputNum) ProtoReflect() protoreflect.Message {
	mi := &file_probuffile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputNum.ProtoReflect.Descriptor instead.
func (*InputNum) Descriptor() ([]byte, []int) {
	return file_probuffile_proto_rawDescGZIP(), []int{2}
}

func (x *InputNum) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AddOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddOutput) Reset() {
	*x = AddOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probuffile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOutput) ProtoMessage() {}

func (x *AddOutput) ProtoReflect() protoreflect.Message {
	mi := &file_probuffile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOutput.ProtoReflect.Descriptor instead.
func (*AddOutput) Descriptor() ([]byte, []int) {
	return file_probuffile_proto_rawDescGZIP(), []int{3}
}

func (x *AddOutput) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_probuffile_proto protoreflect.FileDescriptor

var file_probuffile_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x1b, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1c, 0x0a, 0x08, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x1d, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x73, 0x75, 0x6d, 0x32, 0x7d, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x12, 0x34, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x0b, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x0e,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x1a, 0x0f,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_probuffile_proto_rawDescOnce sync.Once
	file_probuffile_proto_rawDescData = file_probuffile_proto_rawDesc
)

func file_probuffile_proto_rawDescGZIP() []byte {
	file_probuffile_proto_rawDescOnce.Do(func() {
		file_probuffile_proto_rawDescData = protoimpl.X.CompressGZIP(file_probuffile_proto_rawDescData)
	})
	return file_probuffile_proto_rawDescData
}

var file_probuffile_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_probuffile_proto_goTypes = []interface{}{
	(*Input)(nil),     // 0: demo.Input
	(*Output)(nil),    // 1: demo.Output
	(*InputNum)(nil),  // 2: demo.InputNum
	(*AddOutput)(nil), // 3: demo.AddOutput
}
var file_probuffile_proto_depIdxs = []int32{
	0, // 0: demo.Streaming.ServerSideStreaming:input_type -> demo.Input
	2, // 1: demo.Streaming.ClientSideStreaming:input_type -> demo.InputNum
	1, // 2: demo.Streaming.ServerSideStreaming:output_type -> demo.Output
	3, // 3: demo.Streaming.ClientSideStreaming:output_type -> demo.AddOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_probuffile_proto_init() }
func file_probuffile_proto_init() {
	if File_probuffile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_probuffile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_probuffile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_probuffile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputNum); i {
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
		file_probuffile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOutput); i {
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
			RawDescriptor: file_probuffile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_probuffile_proto_goTypes,
		DependencyIndexes: file_probuffile_proto_depIdxs,
		MessageInfos:      file_probuffile_proto_msgTypes,
	}.Build()
	File_probuffile_proto = out.File
	file_probuffile_proto_rawDesc = nil
	file_probuffile_proto_goTypes = nil
	file_probuffile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamingClient is the client API for Streaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamingClient interface {
	ServerSideStreaming(ctx context.Context, in *Input, opts ...grpc.CallOption) (Streaming_ServerSideStreamingClient, error)
	ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (Streaming_ClientSideStreamingClient, error)
}

type streamingClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingClient(cc grpc.ClientConnInterface) StreamingClient {
	return &streamingClient{cc}
}

func (c *streamingClient) ServerSideStreaming(ctx context.Context, in *Input, opts ...grpc.CallOption) (Streaming_ServerSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streaming_serviceDesc.Streams[0], "/demo.Streaming/ServerSideStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServerSideStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streaming_ServerSideStreamingClient interface {
	Recv() (*Output, error)
	grpc.ClientStream
}

type streamingServerSideStreamingClient struct {
	grpc.ClientStream
}

func (x *streamingServerSideStreamingClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingClient) ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (Streaming_ClientSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streaming_serviceDesc.Streams[1], "/demo.Streaming/ClientSideStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingClientSideStreamingClient{stream}
	return x, nil
}

type Streaming_ClientSideStreamingClient interface {
	Send(*InputNum) error
	CloseAndRecv() (*AddOutput, error)
	grpc.ClientStream
}

type streamingClientSideStreamingClient struct {
	grpc.ClientStream
}

func (x *streamingClientSideStreamingClient) Send(m *InputNum) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingClientSideStreamingClient) CloseAndRecv() (*AddOutput, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServer is the server API for Streaming service.
type StreamingServer interface {
	ServerSideStreaming(*Input, Streaming_ServerSideStreamingServer) error
	ClientSideStreaming(Streaming_ClientSideStreamingServer) error
}

// UnimplementedStreamingServer can be embedded to have forward compatible implementations.
type UnimplementedStreamingServer struct {
}

func (*UnimplementedStreamingServer) ServerSideStreaming(*Input, Streaming_ServerSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStreaming not implemented")
}
func (*UnimplementedStreamingServer) ClientSideStreaming(Streaming_ClientSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSideStreaming not implemented")
}

func RegisterStreamingServer(s *grpc.Server, srv StreamingServer) {
	s.RegisterService(&_Streaming_serviceDesc, srv)
}

func _Streaming_ServerSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Input)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServer).ServerSideStreaming(m, &streamingServerSideStreamingServer{stream})
}

type Streaming_ServerSideStreamingServer interface {
	Send(*Output) error
	grpc.ServerStream
}

type streamingServerSideStreamingServer struct {
	grpc.ServerStream
}

func (x *streamingServerSideStreamingServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

func _Streaming_ClientSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServer).ClientSideStreaming(&streamingClientSideStreamingServer{stream})
}

type Streaming_ClientSideStreamingServer interface {
	SendAndClose(*AddOutput) error
	Recv() (*InputNum, error)
	grpc.ServerStream
}

type streamingClientSideStreamingServer struct {
	grpc.ServerStream
}

func (x *streamingClientSideStreamingServer) SendAndClose(m *AddOutput) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingClientSideStreamingServer) Recv() (*InputNum, error) {
	m := new(InputNum)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Streaming_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.Streaming",
	HandlerType: (*StreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStreaming",
			Handler:       _Streaming_ServerSideStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSideStreaming",
			Handler:       _Streaming_ClientSideStreaming_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "probuffile.proto",
}