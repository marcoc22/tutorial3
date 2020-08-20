// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: books_info.proto

package booksapp

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=Title,json=title,proto3" json:"Title,omitempty"`
	Edition   string `protobuf:"bytes,3,opt,name=Edition,json=edition,proto3" json:"Edition,omitempty"`
	Copyright string `protobuf:"bytes,4,opt,name=Copyright,json=copyright,proto3" json:"Copyright,omitempty"`
	Language  string `protobuf:"bytes,5,opt,name=Language,json=language,proto3" json:"Language,omitempty"`
	Pages     string `protobuf:"bytes,6,opt,name=Pages,json=pages,proto3" json:"Pages,omitempty"`
	Author    string `protobuf:"bytes,7,opt,name=Author,json=author,proto3" json:"Author,omitempty"`
	Publisher string `protobuf:"bytes,8,opt,name=Publisher,json=publisher,proto3" json:"Publisher,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_books_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_info_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetEdition() string {
	if x != nil {
		return x.Edition
	}
	return ""
}

func (x *Book) GetCopyright() string {
	if x != nil {
		return x.Copyright
	}
	return ""
}

func (x *Book) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Book) GetPages() string {
	if x != nil {
		return x.Pages
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

type BookID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BookID) Reset() {
	*x = BookID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookID) ProtoMessage() {}

func (x *BookID) ProtoReflect() protoreflect.Message {
	mi := &file_books_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookID.ProtoReflect.Descriptor instead.
func (*BookID) Descriptor() ([]byte, []int) {
	return file_books_info_proto_rawDescGZIP(), []int{1}
}

func (x *BookID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_books_info_proto protoreflect.FileDescriptor

var file_books_info_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x22, 0xcc, 0x01, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x06, 0x42,
	0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xc2, 0x01, 0x0a, 0x08,
	0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x2c, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x1a, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x2e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x10,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44,
	0x1a, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_info_proto_rawDescOnce sync.Once
	file_books_info_proto_rawDescData = file_books_info_proto_rawDesc
)

func file_books_info_proto_rawDescGZIP() []byte {
	file_books_info_proto_rawDescOnce.Do(func() {
		file_books_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_info_proto_rawDescData)
	})
	return file_books_info_proto_rawDescData
}

var file_books_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_books_info_proto_goTypes = []interface{}{
	(*Book)(nil),   // 0: booksapp.Book
	(*BookID)(nil), // 1: booksapp.BookID
}
var file_books_info_proto_depIdxs = []int32{
	0, // 0: booksapp.BookInfo.addBook:input_type -> booksapp.Book
	1, // 1: booksapp.BookInfo.getBook:input_type -> booksapp.BookID
	0, // 2: booksapp.BookInfo.updateBook:input_type -> booksapp.Book
	1, // 3: booksapp.BookInfo.deleteBook:input_type -> booksapp.BookID
	1, // 4: booksapp.BookInfo.addBook:output_type -> booksapp.BookID
	0, // 5: booksapp.BookInfo.getBook:output_type -> booksapp.Book
	0, // 6: booksapp.BookInfo.updateBook:output_type -> booksapp.Book
	0, // 7: booksapp.BookInfo.deleteBook:output_type -> booksapp.Book
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_books_info_proto_init() }
func file_books_info_proto_init() {
	if File_books_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_books_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookID); i {
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
			RawDescriptor: file_books_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_info_proto_goTypes,
		DependencyIndexes: file_books_info_proto_depIdxs,
		MessageInfos:      file_books_info_proto_msgTypes,
	}.Build()
	File_books_info_proto = out.File
	file_books_info_proto_rawDesc = nil
	file_books_info_proto_goTypes = nil
	file_books_info_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookInfoClient is the client API for BookInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookInfoClient interface {
	AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookID, error)
	GetBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error)
}

type bookInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewBookInfoClient(cc grpc.ClientConnInterface) BookInfoClient {
	return &bookInfoClient{cc}
}

func (c *bookInfoClient) AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookID, error) {
	out := new(BookID)
	err := c.cc.Invoke(ctx, "/booksapp.BookInfo/addBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookInfoClient) GetBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/booksapp.BookInfo/getBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookInfoClient) UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/booksapp.BookInfo/updateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookInfoClient) DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/booksapp.BookInfo/deleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookInfoServer is the server API for BookInfo service.
type BookInfoServer interface {
	AddBook(context.Context, *Book) (*BookID, error)
	GetBook(context.Context, *BookID) (*Book, error)
	UpdateBook(context.Context, *Book) (*Book, error)
	DeleteBook(context.Context, *BookID) (*Book, error)
}

// UnimplementedBookInfoServer can be embedded to have forward compatible implementations.
type UnimplementedBookInfoServer struct {
}

func (*UnimplementedBookInfoServer) AddBook(context.Context, *Book) (*BookID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (*UnimplementedBookInfoServer) GetBook(context.Context, *BookID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookInfoServer) UpdateBook(context.Context, *Book) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (*UnimplementedBookInfoServer) DeleteBook(context.Context, *BookID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}

func RegisterBookInfoServer(s *grpc.Server, srv BookInfoServer) {
	s.RegisterService(&_BookInfo_serviceDesc, srv)
}

func _BookInfo_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookInfoServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booksapp.BookInfo/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookInfoServer).AddBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookInfo_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookInfoServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booksapp.BookInfo/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookInfoServer).GetBook(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookInfo_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookInfoServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booksapp.BookInfo/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookInfoServer).UpdateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookInfo_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookInfoServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booksapp.BookInfo/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookInfoServer).DeleteBook(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "booksapp.BookInfo",
	HandlerType: (*BookInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addBook",
			Handler:    _BookInfo_AddBook_Handler,
		},
		{
			MethodName: "getBook",
			Handler:    _BookInfo_GetBook_Handler,
		},
		{
			MethodName: "updateBook",
			Handler:    _BookInfo_UpdateBook_Handler,
		},
		{
			MethodName: "deleteBook",
			Handler:    _BookInfo_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books_info.proto",
}