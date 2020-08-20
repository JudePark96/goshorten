//
//CreateURL will take in a long url, generate a 6 character code to associate with the URL
//and store it in a redis datastore for XX amount of time. (length of time undetermined) and
//return a shortened URL code m
//
//GetURL will take in a code and return/redirect users to the long URL, if still active.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: url_service.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongUrl string `protobuf:"bytes,1,opt,name=LongUrl,proto3" json:"LongUrl,omitempty"`
	TTL     int64  `protobuf:"varint,2,opt,name=TTL,proto3" json:"TTL,omitempty"`
	Stats   string `protobuf:"bytes,3,opt,name=Stats,proto3" json:"Stats,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_url_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_url_service_proto_rawDescGZIP(), []int{0}
}

func (x *URL) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

func (x *URL) GetTTL() int64 {
	if x != nil {
		return x.TTL
	}
	return 0
}

func (x *URL) GetStats() string {
	if x != nil {
		return x.Stats
	}
	return ""
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_url_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_url_service_proto_rawDescGZIP(), []int{1}
}

func (x *Code) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats string `protobuf:"bytes,1,opt,name=Stats,proto3" json:"Stats,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_url_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_url_service_proto_rawDescGZIP(), []int{2}
}

func (x *Stats) GetStats() string {
	if x != nil {
		return x.Stats
	}
	return ""
}

var File_url_service_proto protoreflect.FileDescriptor

var file_url_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f,
	0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4c, 0x6f, 0x6e,
	0x67, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x54, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x54, 0x54, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x32, 0x57, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x52,
	0x4c, 0x12, 0x04, 0x2e, 0x55, 0x52, 0x4c, 0x1a, 0x05, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x05, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x1a,
	0x04, 0x2e, 0x55, 0x52, 0x4c, 0x12, 0x19, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x05, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x06, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_url_service_proto_rawDescOnce sync.Once
	file_url_service_proto_rawDescData = file_url_service_proto_rawDesc
)

func file_url_service_proto_rawDescGZIP() []byte {
	file_url_service_proto_rawDescOnce.Do(func() {
		file_url_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_service_proto_rawDescData)
	})
	return file_url_service_proto_rawDescData
}

var file_url_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_url_service_proto_goTypes = []interface{}{
	(*URL)(nil),   // 0: URL
	(*Code)(nil),  // 1: Code
	(*Stats)(nil), // 2: Stats
}
var file_url_service_proto_depIdxs = []int32{
	0, // 0: Shortener.CreateURL:input_type -> URL
	1, // 1: Shortener.GetURL:input_type -> Code
	1, // 2: Shortener.GetStats:input_type -> Code
	1, // 3: Shortener.CreateURL:output_type -> Code
	0, // 4: Shortener.GetURL:output_type -> URL
	2, // 5: Shortener.GetStats:output_type -> Stats
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_service_proto_init() }
func file_url_service_proto_init() {
	if File_url_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
		file_url_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_url_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
			RawDescriptor: file_url_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_service_proto_goTypes,
		DependencyIndexes: file_url_service_proto_depIdxs,
		MessageInfos:      file_url_service_proto_msgTypes,
	}.Build()
	File_url_service_proto = out.File
	file_url_service_proto_rawDesc = nil
	file_url_service_proto_goTypes = nil
	file_url_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShortenerClient interface {
	CreateURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Code, error)
	GetURL(ctx context.Context, in *Code, opts ...grpc.CallOption) (*URL, error)
	GetStats(ctx context.Context, in *Code, opts ...grpc.CallOption) (*Stats, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) CreateURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Code, error) {
	out := new(Code)
	err := c.cc.Invoke(ctx, "/Shortener/CreateURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetURL(ctx context.Context, in *Code, opts ...grpc.CallOption) (*URL, error) {
	out := new(URL)
	err := c.cc.Invoke(ctx, "/Shortener/GetURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetStats(ctx context.Context, in *Code, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/Shortener/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
type ShortenerServer interface {
	CreateURL(context.Context, *URL) (*Code, error)
	GetURL(context.Context, *Code) (*URL, error)
	GetStats(context.Context, *Code) (*Stats, error)
}

// UnimplementedShortenerServer can be embedded to have forward compatible implementations.
type UnimplementedShortenerServer struct {
}

func (*UnimplementedShortenerServer) CreateURL(context.Context, *URL) (*Code, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateURL not implemented")
}
func (*UnimplementedShortenerServer) GetURL(context.Context, *Code) (*URL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURL not implemented")
}
func (*UnimplementedShortenerServer) GetStats(context.Context, *Code) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}

func RegisterShortenerServer(s *grpc.Server, srv ShortenerServer) {
	s.RegisterService(&_Shortener_serviceDesc, srv)
}

func _Shortener_CreateURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).CreateURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortener/CreateURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).CreateURL(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortener/GetURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetURL(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortener/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetStats(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shortener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateURL",
			Handler:    _Shortener_CreateURL_Handler,
		},
		{
			MethodName: "GetURL",
			Handler:    _Shortener_GetURL_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Shortener_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "url_service.proto",
}
