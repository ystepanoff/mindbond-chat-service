// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: proto/translator.proto

package pb

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

type TranslateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	FromLang string `protobuf:"bytes,2,opt,name=fromLang,proto3" json:"fromLang,omitempty"`
	ToLang   string `protobuf:"bytes,3,opt,name=toLang,proto3" json:"toLang,omitempty"`
}

func (x *TranslateRequest) Reset() {
	*x = TranslateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_translator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateRequest) ProtoMessage() {}

func (x *TranslateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_translator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateRequest.ProtoReflect.Descriptor instead.
func (*TranslateRequest) Descriptor() ([]byte, []int) {
	return file_proto_translator_proto_rawDescGZIP(), []int{0}
}

func (x *TranslateRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TranslateRequest) GetFromLang() string {
	if x != nil {
		return x.FromLang
	}
	return ""
}

func (x *TranslateRequest) GetToLang() string {
	if x != nil {
		return x.ToLang
	}
	return ""
}

type TranslateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error       string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Translation string `protobuf:"bytes,3,opt,name=translation,proto3" json:"translation,omitempty"`
}

func (x *TranslateResponse) Reset() {
	*x = TranslateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_translator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateResponse) ProtoMessage() {}

func (x *TranslateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_translator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateResponse.ProtoReflect.Descriptor instead.
func (*TranslateResponse) Descriptor() ([]byte, []int) {
	return file_proto_translator_proto_rawDescGZIP(), []int{1}
}

func (x *TranslateResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TranslateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *TranslateResponse) GetTranslation() string {
	if x != nil {
		return x.Translation
	}
	return ""
}

var File_proto_translator_proto protoreflect.FileDescriptor

var file_proto_translator_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x22, 0x60, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x4c, 0x61, 0x6e, 0x67, 0x22, 0x63, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x5f, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_translator_proto_rawDescOnce sync.Once
	file_proto_translator_proto_rawDescData = file_proto_translator_proto_rawDesc
)

func file_proto_translator_proto_rawDescGZIP() []byte {
	file_proto_translator_proto_rawDescOnce.Do(func() {
		file_proto_translator_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_translator_proto_rawDescData)
	})
	return file_proto_translator_proto_rawDescData
}

var file_proto_translator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_translator_proto_goTypes = []interface{}{
	(*TranslateRequest)(nil),  // 0: translator.TranslateRequest
	(*TranslateResponse)(nil), // 1: translator.TranslateResponse
}
var file_proto_translator_proto_depIdxs = []int32{
	0, // 0: translator.TranslatorService.Translate:input_type -> translator.TranslateRequest
	1, // 1: translator.TranslatorService.Translate:output_type -> translator.TranslateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_translator_proto_init() }
func file_proto_translator_proto_init() {
	if File_proto_translator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_translator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateRequest); i {
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
		file_proto_translator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateResponse); i {
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
			RawDescriptor: file_proto_translator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_translator_proto_goTypes,
		DependencyIndexes: file_proto_translator_proto_depIdxs,
		MessageInfos:      file_proto_translator_proto_msgTypes,
	}.Build()
	File_proto_translator_proto = out.File
	file_proto_translator_proto_rawDesc = nil
	file_proto_translator_proto_goTypes = nil
	file_proto_translator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TranslatorServiceClient is the client API for TranslatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TranslatorServiceClient interface {
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
}

type translatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslatorServiceClient(cc grpc.ClientConnInterface) TranslatorServiceClient {
	return &translatorServiceClient{cc}
}

func (c *translatorServiceClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	out := new(TranslateResponse)
	err := c.cc.Invoke(ctx, "/translator.TranslatorService/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServiceServer is the server API for TranslatorService service.
type TranslatorServiceServer interface {
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
}

// UnimplementedTranslatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTranslatorServiceServer struct {
}

func (*UnimplementedTranslatorServiceServer) Translate(context.Context, *TranslateRequest) (*TranslateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}

func RegisterTranslatorServiceServer(s *grpc.Server, srv TranslatorServiceServer) {
	s.RegisterService(&_TranslatorService_serviceDesc, srv)
}

func _TranslatorService_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServiceServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.TranslatorService/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServiceServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TranslatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "translator.TranslatorService",
	HandlerType: (*TranslatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _TranslatorService_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/translator.proto",
}
