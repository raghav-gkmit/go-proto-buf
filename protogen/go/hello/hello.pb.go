// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/hello.proto

package hello

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age           int32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_proto_hello_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_proto_hello_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_hello_proto protoreflect.FileDescriptor

const file_proto_hello_proto_rawDesc = "" +
	"\n" +
	"\x11proto/hello.proto\x12\x05hello\"4\n" +
	"\fHelloRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x10\n" +
	"\x03age\x18\x02 \x01(\x05R\x03age\")\n" +
	"\rHelloResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2E\n" +
	"\fHelloService\x125\n" +
	"\bSayHello\x12\x13.hello.HelloRequest\x1a\x14.hello.HelloResponseB8Z6github.com/raghav-gkmit/go-proto-buf/protogen/go/hellob\x06proto3"

var (
	file_proto_hello_proto_rawDescOnce sync.Once
	file_proto_hello_proto_rawDescData []byte
)

func file_proto_hello_proto_rawDescGZIP() []byte {
	file_proto_hello_proto_rawDescOnce.Do(func() {
		file_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_hello_proto_rawDesc), len(file_proto_hello_proto_rawDesc)))
	})
	return file_proto_hello_proto_rawDescData
}

var file_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_hello_proto_goTypes = []any{
	(*HelloRequest)(nil),  // 0: hello.HelloRequest
	(*HelloResponse)(nil), // 1: hello.HelloResponse
}
var file_proto_hello_proto_depIdxs = []int32{
	0, // 0: hello.HelloService.SayHello:input_type -> hello.HelloRequest
	1, // 1: hello.HelloService.SayHello:output_type -> hello.HelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_hello_proto_init() }
func file_proto_hello_proto_init() {
	if File_proto_hello_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_hello_proto_rawDesc), len(file_proto_hello_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_proto_goTypes,
		DependencyIndexes: file_proto_hello_proto_depIdxs,
		MessageInfos:      file_proto_hello_proto_msgTypes,
	}.Build()
	File_proto_hello_proto = out.File
	file_proto_hello_proto_goTypes = nil
	file_proto_hello_proto_depIdxs = nil
}
