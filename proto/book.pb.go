// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: proto/book.proto

package __

import (
	context "context"
	"gRPC_Mongo_HTTP_TestTask/server"
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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tag      string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[0]
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
	return file_proto_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type CreateBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookReq) Reset() {
	*x = CreateBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookReq) ProtoMessage() {}

func (x *CreateBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookReq.ProtoReflect.Descriptor instead.
func (*CreateBookReq) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookReq) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type CreateBookRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookRes) Reset() {
	*x = CreateBookRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRes) ProtoMessage() {}

func (x *CreateBookRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRes.ProtoReflect.Descriptor instead.
func (*CreateBookRes) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookRes) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type ReadBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadBookReq) Reset() {
	*x = ReadBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBookReq) ProtoMessage() {}

func (x *ReadBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBookReq.ProtoReflect.Descriptor instead.
func (*ReadBookReq) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{3}
}

func (x *ReadBookReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadBookRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Book `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *ReadBookRes) Reset() {
	*x = ReadBookRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBookRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBookRes) ProtoMessage() {}

func (x *ReadBookRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBookRes.ProtoReflect.Descriptor instead.
func (*ReadBookRes) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{4}
}

func (x *ReadBookRes) GetBlog() *Book {
	if x != nil {
		return x.Blog
	}
	return nil
}

type UpdateBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Book `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBookReq) Reset() {
	*x = UpdateBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookReq) ProtoMessage() {}

func (x *UpdateBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookReq.ProtoReflect.Descriptor instead.
func (*UpdateBookReq) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBookReq) GetBlog() *Book {
	if x != nil {
		return x.Blog
	}
	return nil
}

type UpdateBookRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Book `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBookRes) Reset() {
	*x = UpdateBookRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRes) ProtoMessage() {}

func (x *UpdateBookRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRes.ProtoReflect.Descriptor instead.
func (*UpdateBookRes) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBookRes) GetBlog() *Book {
	if x != nil {
		return x.Blog
	}
	return nil
}

type DeleteBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookReq) Reset() {
	*x = DeleteBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookReq) ProtoMessage() {}

func (x *DeleteBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookReq.ProtoReflect.Descriptor instead.
func (*DeleteBookReq) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBookRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteBookRes) Reset() {
	*x = DeleteBookRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRes) ProtoMessage() {}

func (x *DeleteBookRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRes.ProtoReflect.Descriptor instead.
func (*DeleteBookRes) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBookRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBookReq) Reset() {
	*x = ListBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookReq) ProtoMessage() {}

func (x *ListBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookReq.ProtoReflect.Descriptor instead.
func (*ListBookReq) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{9}
}

type ListBookRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Book `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *ListBookRes) Reset() {
	*x = ListBookRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookRes) ProtoMessage() {}

func (x *ListBookRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookRes.ProtoReflect.Descriptor instead.
func (*ListBookRes) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{10}
}

func (x *ListBookRes) GetBlog() *Book {
	if x != nil {
		return x.Blog
	}
	return nil
}

var File_proto_book_proto protoreflect.FileDescriptor

var file_proto_book_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x59, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x22, 0x2f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x2f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62,
	0x6c, 0x6f, 0x67, 0x22, 0x2f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04,
	0x62, 0x6c, 0x6f, 0x67, 0x22, 0x2f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x22, 0x2d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x32,
	0x9b, 0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x13, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x30, 0x01, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_book_proto_rawDescOnce sync.Once
	file_proto_book_proto_rawDescData = file_proto_book_proto_rawDesc
)

func file_proto_book_proto_rawDescGZIP() []byte {
	file_proto_book_proto_rawDescOnce.Do(func() {
		file_proto_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_book_proto_rawDescData)
	})
	return file_proto_book_proto_rawDescData
}

var file_proto_book_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_book_proto_goTypes = []interface{}{
	(*Book)(nil),          // 0: book.Book
	(*CreateBookReq)(nil), // 1: book.CreateBookReq
	(*CreateBookRes)(nil), // 2: book.CreateBookRes
	(*ReadBookReq)(nil),   // 3: book.ReadBookReq
	(*ReadBookRes)(nil),   // 4: book.ReadBookRes
	(*UpdateBookReq)(nil), // 5: book.UpdateBookReq
	(*UpdateBookRes)(nil), // 6: book.UpdateBookRes
	(*DeleteBookReq)(nil), // 7: book.DeleteBookReq
	(*DeleteBookRes)(nil), // 8: book.DeleteBookRes
	(*ListBookReq)(nil),   // 9: book.ListBookReq
	(*ListBookRes)(nil),   // 10: book.ListBookRes
}
var file_proto_book_proto_depIdxs = []int32{
	0,  // 0: book.CreateBookReq.book:type_name -> book.Book
	0,  // 1: book.CreateBookRes.book:type_name -> book.Book
	0,  // 2: book.ReadBookRes.blog:type_name -> book.Book
	0,  // 3: book.UpdateBookReq.blog:type_name -> book.Book
	0,  // 4: book.UpdateBookRes.blog:type_name -> book.Book
	0,  // 5: book.ListBookRes.blog:type_name -> book.Book
	1,  // 6: book.BookService.CreateBook:input_type -> book.CreateBookReq
	3,  // 7: book.BookService.ReadBook:input_type -> book.ReadBookReq
	5,  // 8: book.BookService.UpdateBook:input_type -> book.UpdateBookReq
	7,  // 9: book.BookService.DeleteBook:input_type -> book.DeleteBookReq
	9,  // 10: book.BookService.ListBook:input_type -> book.ListBookReq
	2,  // 11: book.BookService.CreateBook:output_type -> book.CreateBookRes
	4,  // 12: book.BookService.ReadBook:output_type -> book.ReadBookRes
	6,  // 13: book.BookService.UpdateBook:output_type -> book.UpdateBookRes
	8,  // 14: book.BookService.DeleteBook:output_type -> book.DeleteBookRes
	10, // 15: book.BookService.ListBook:output_type -> book.ListBookRes
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_book_proto_init() }
func file_proto_book_proto_init() {
	if File_proto_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookReq); i {
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
		file_proto_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRes); i {
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
		file_proto_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBookReq); i {
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
		file_proto_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBookRes); i {
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
		file_proto_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookReq); i {
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
		file_proto_book_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRes); i {
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
		file_proto_book_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookReq); i {
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
		file_proto_book_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRes); i {
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
		file_proto_book_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookReq); i {
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
		file_proto_book_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookRes); i {
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
			RawDescriptor: file_proto_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_book_proto_goTypes,
		DependencyIndexes: file_proto_book_proto_depIdxs,
		MessageInfos:      file_proto_book_proto_msgTypes,
	}.Build()
	File_proto_book_proto = out.File
	file_proto_book_proto_rawDesc = nil
	file_proto_book_proto_goTypes = nil
	file_proto_book_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	CreateBook(ctx context.Context, in *CreateBookReq, opts ...grpc.CallOption) (*CreateBookRes, error)
	ReadBook(ctx context.Context, in *ReadBookReq, opts ...grpc.CallOption) (*ReadBookRes, error)
	UpdateBook(ctx context.Context, in *UpdateBookReq, opts ...grpc.CallOption) (*UpdateBookRes, error)
	DeleteBook(ctx context.Context, in *DeleteBookReq, opts ...grpc.CallOption) (*DeleteBookRes, error)
	ListBook(ctx context.Context, in *ListBookReq, opts ...grpc.CallOption) (BookService_ListBookClient, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookReq, opts ...grpc.CallOption) (*CreateBookRes, error) {
	out := new(CreateBookRes)
	err := c.cc.Invoke(ctx, "/book.BookService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ReadBook(ctx context.Context, in *ReadBookReq, opts ...grpc.CallOption) (*ReadBookRes, error) {
	out := new(ReadBookRes)
	err := c.cc.Invoke(ctx, "/book.BookService/ReadBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *UpdateBookReq, opts ...grpc.CallOption) (*UpdateBookRes, error) {
	out := new(UpdateBookRes)
	err := c.cc.Invoke(ctx, "/book.BookService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *DeleteBookReq, opts ...grpc.CallOption) (*DeleteBookRes, error) {
	out := new(DeleteBookRes)
	err := c.cc.Invoke(ctx, "/book.BookService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ListBook(ctx context.Context, in *ListBookReq, opts ...grpc.CallOption) (BookService_ListBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BookService_serviceDesc.Streams[0], "/book.BookService/ListBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceListBookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_ListBookClient interface {
	Recv() (*ListBookRes, error)
	grpc.ClientStream
}

type bookServiceListBookClient struct {
	grpc.ClientStream
}

func (x *bookServiceListBookClient) Recv() (*ListBookRes, error) {
	m := new(ListBookRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	CreateBook(context.Context, *CreateBookReq) (*CreateBookRes, error)
	ReadBook(context.Context, *ReadBookReq) (*ReadBookRes, error)
	UpdateBook(context.Context, *UpdateBookReq) (*UpdateBookRes, error)
	DeleteBook(context.Context, *DeleteBookReq) (*DeleteBookRes, error)
	ListBook(*ListBookReq, BookService_ListBookServer) error
}

// UnimplementedBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (*UnimplementedBookServiceServer) CreateBook(context.Context, *CreateBookReq) (*CreateBookRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (*UnimplementedBookServiceServer) ReadBook(context.Context, *ReadBookReq) (*ReadBookRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBook not implemented")
}
func (*UnimplementedBookServiceServer) UpdateBook(context.Context, *UpdateBookReq) (*UpdateBookRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (*UnimplementedBookServiceServer) DeleteBook(context.Context, *DeleteBookReq) (*DeleteBookRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (*UnimplementedBookServiceServer) ListBook(*ListBookReq, BookService_ListBookServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBook not implemented")
}

func RegisterBookServiceServer(s *grpc.Server, srv *server.BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ReadBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ReadBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/ReadBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ReadBook(ctx, req.(*ReadBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*UpdateBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*DeleteBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ListBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBookReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).ListBook(m, &bookServiceListBookServer{stream})
}

type BookService_ListBookServer interface {
	Send(*ListBookRes) error
	grpc.ServerStream
}

type bookServiceListBookServer struct {
	grpc.ServerStream
}

func (x *bookServiceListBookServer) Send(m *ListBookRes) error {
	return x.ServerStream.SendMsg(m)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "ReadBook",
			Handler:    _BookService_ReadBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBook",
			Handler:       _BookService_ListBook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/book.proto",
}
