// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: loginAndRegister.proto

package loginAndRegister

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

// The request message containing the user's name.
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginAndRegister_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loginAndRegister_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_loginAndRegister_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// The response message containing the greetings
type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int32  `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginAndRegister_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_loginAndRegister_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_loginAndRegister_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *LoginReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The request message containing the user's name.
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail *UserDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginAndRegister_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loginAndRegister_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_loginAndRegister_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetDetail() *UserDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

// The response message containing the greetings
type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int32 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginAndRegister_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_loginAndRegister_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_loginAndRegister_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterReply) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Age      int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Id       int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginAndRegister_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_loginAndRegister_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_loginAndRegister_proto_rawDescGZIP(), []int{4}
}

func (x *UserDetail) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserDetail) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserDetail) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_loginAndRegister_proto protoreflect.FileDescriptor

var file_loginAndRegister_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41,
	0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x64,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x32, 0xa1, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e,
	0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e,
	0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x26, 0x0a, 0x10, 0x69, 0x6f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x5a, 0x12, 0x2e, 0x3b,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loginAndRegister_proto_rawDescOnce sync.Once
	file_loginAndRegister_proto_rawDescData = file_loginAndRegister_proto_rawDesc
)

func file_loginAndRegister_proto_rawDescGZIP() []byte {
	file_loginAndRegister_proto_rawDescOnce.Do(func() {
		file_loginAndRegister_proto_rawDescData = protoimpl.X.CompressGZIP(file_loginAndRegister_proto_rawDescData)
	})
	return file_loginAndRegister_proto_rawDescData
}

var file_loginAndRegister_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_loginAndRegister_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),    // 0: loginAndRegister.LoginRequest
	(*LoginReply)(nil),      // 1: loginAndRegister.LoginReply
	(*RegisterRequest)(nil), // 2: loginAndRegister.RegisterRequest
	(*RegisterReply)(nil),   // 3: loginAndRegister.RegisterReply
	(*UserDetail)(nil),      // 4: loginAndRegister.userDetail
}
var file_loginAndRegister_proto_depIdxs = []int32{
	4, // 0: loginAndRegister.RegisterRequest.detail:type_name -> loginAndRegister.userDetail
	0, // 1: loginAndRegister.Test.Login:input_type -> loginAndRegister.LoginRequest
	2, // 2: loginAndRegister.Test.Register:input_type -> loginAndRegister.RegisterRequest
	1, // 3: loginAndRegister.Test.Login:output_type -> loginAndRegister.LoginReply
	3, // 4: loginAndRegister.Test.Register:output_type -> loginAndRegister.RegisterReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_loginAndRegister_proto_init() }
func file_loginAndRegister_proto_init() {
	if File_loginAndRegister_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loginAndRegister_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_loginAndRegister_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_loginAndRegister_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_loginAndRegister_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_loginAndRegister_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
			RawDescriptor: file_loginAndRegister_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loginAndRegister_proto_goTypes,
		DependencyIndexes: file_loginAndRegister_proto_depIdxs,
		MessageInfos:      file_loginAndRegister_proto_msgTypes,
	}.Build()
	File_loginAndRegister_proto = out.File
	file_loginAndRegister_proto_rawDesc = nil
	file_loginAndRegister_proto_goTypes = nil
	file_loginAndRegister_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	// Sends a greeting
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/loginAndRegister.Test/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/loginAndRegister.Test/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	// Sends a greeting
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedTestServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loginAndRegister.Test/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loginAndRegister.Test/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loginAndRegister.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Test_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Test_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loginAndRegister.proto",
}
