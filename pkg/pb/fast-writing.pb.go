// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: fast-writing.proto

package pb

import (
	models "fast-writing/pkg/pb/models"
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

type CreateContentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents *models.Contents `protobuf:"bytes,1,opt,name=contents,proto3" json:"contents,omitempty"`
	UserId   *models.UserId   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateContentsRequest) Reset() {
	*x = CreateContentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fast_writing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentsRequest) ProtoMessage() {}

func (x *CreateContentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fast_writing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentsRequest.ProtoReflect.Descriptor instead.
func (*CreateContentsRequest) Descriptor() ([]byte, []int) {
	return file_fast_writing_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContentsRequest) GetContents() *models.Contents {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *CreateContentsRequest) GetUserId() *models.UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

type CreateContentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created  bool             `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Message  string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Contents *models.Contents `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *CreateContentsResponse) Reset() {
	*x = CreateContentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fast_writing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentsResponse) ProtoMessage() {}

func (x *CreateContentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fast_writing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentsResponse.ProtoReflect.Descriptor instead.
func (*CreateContentsResponse) Descriptor() ([]byte, []int) {
	return file_fast_writing_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContentsResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *CreateContentsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateContentsResponse) GetContents() *models.Contents {
	if x != nil {
		return x.Contents
	}
	return nil
}

type DeleteContentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentsId *models.ContentsId `protobuf:"bytes,1,opt,name=contentsId,proto3" json:"contentsId,omitempty"`
}

func (x *DeleteContentsRequest) Reset() {
	*x = DeleteContentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fast_writing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteContentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContentsRequest) ProtoMessage() {}

func (x *DeleteContentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fast_writing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContentsRequest.ProtoReflect.Descriptor instead.
func (*DeleteContentsRequest) Descriptor() ([]byte, []int) {
	return file_fast_writing_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteContentsRequest) GetContentsId() *models.ContentsId {
	if x != nil {
		return x.ContentsId
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool   `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fast_writing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fast_writing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_fast_writing_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_fast_writing_proto protoreflect.FileDescriptor

var file_fast_writing_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x61, 0x73, 0x74, 0x2d, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x61, 0x73, 0x74, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7a,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x94, 0x01,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x00, 0x32, 0x99, 0x03, 0x0a, 0x0e, 0x57, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x14,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x14,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x66,
	0x61, 0x73, 0x74, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e,
	0x66, 0x61, 0x73, 0x74, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x16, 0x5a, 0x14, 0x66, 0x61, 0x73, 0x74, 0x2d, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fast_writing_proto_rawDescOnce sync.Once
	file_fast_writing_proto_rawDescData = file_fast_writing_proto_rawDesc
)

func file_fast_writing_proto_rawDescGZIP() []byte {
	file_fast_writing_proto_rawDescOnce.Do(func() {
		file_fast_writing_proto_rawDescData = protoimpl.X.CompressGZIP(file_fast_writing_proto_rawDescData)
	})
	return file_fast_writing_proto_rawDescData
}

var file_fast_writing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fast_writing_proto_goTypes = []interface{}{
	(*CreateContentsRequest)(nil),          // 0: fastwriting.CreateContentsRequest
	(*CreateContentsResponse)(nil),         // 1: fastwriting.CreateContentsResponse
	(*DeleteContentsRequest)(nil),          // 2: fastwriting.DeleteContentsRequest
	(*DeleteResponse)(nil),                 // 3: fastwriting.DeleteResponse
	(*models.Contents)(nil),                // 4: models.Contents
	(*models.UserId)(nil),                  // 5: models.UserId
	(*models.ContentsId)(nil),              // 6: models.ContentsId
	(*models.User)(nil),                    // 7: models.User
	(*models.ContentsQueryParams)(nil),     // 8: models.ContentsQueryParams
	(*models.UserContentsQueryParams)(nil), // 9: models.UserContentsQueryParams
	(*models.ContentsList)(nil),            // 10: models.ContentsList
}
var file_fast_writing_proto_depIdxs = []int32{
	4,  // 0: fastwriting.CreateContentsRequest.contents:type_name -> models.Contents
	5,  // 1: fastwriting.CreateContentsRequest.userId:type_name -> models.UserId
	4,  // 2: fastwriting.CreateContentsResponse.contents:type_name -> models.Contents
	6,  // 3: fastwriting.DeleteContentsRequest.contentsId:type_name -> models.ContentsId
	5,  // 4: fastwriting.UserService.GetUser:input_type -> models.UserId
	7,  // 5: fastwriting.UserService.CreateUser:input_type -> models.User
	7,  // 6: fastwriting.UserService.UpdateUser:input_type -> models.User
	6,  // 7: fastwriting.WritingService.GetContents:input_type -> models.ContentsId
	8,  // 8: fastwriting.WritingService.GetContentsList:input_type -> models.ContentsQueryParams
	9,  // 9: fastwriting.WritingService.GetUserContentsList:input_type -> models.UserContentsQueryParams
	0,  // 10: fastwriting.WritingService.CreateUserContents:input_type -> fastwriting.CreateContentsRequest
	2,  // 11: fastwriting.WritingService.DeleteUserContents:input_type -> fastwriting.DeleteContentsRequest
	7,  // 12: fastwriting.UserService.GetUser:output_type -> models.User
	5,  // 13: fastwriting.UserService.CreateUser:output_type -> models.UserId
	5,  // 14: fastwriting.UserService.UpdateUser:output_type -> models.UserId
	4,  // 15: fastwriting.WritingService.GetContents:output_type -> models.Contents
	10, // 16: fastwriting.WritingService.GetContentsList:output_type -> models.ContentsList
	10, // 17: fastwriting.WritingService.GetUserContentsList:output_type -> models.ContentsList
	1,  // 18: fastwriting.WritingService.CreateUserContents:output_type -> fastwriting.CreateContentsResponse
	3,  // 19: fastwriting.WritingService.DeleteUserContents:output_type -> fastwriting.DeleteResponse
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_fast_writing_proto_init() }
func file_fast_writing_proto_init() {
	if File_fast_writing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fast_writing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentsRequest); i {
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
		file_fast_writing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentsResponse); i {
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
		file_fast_writing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteContentsRequest); i {
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
		file_fast_writing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_fast_writing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_fast_writing_proto_goTypes,
		DependencyIndexes: file_fast_writing_proto_depIdxs,
		MessageInfos:      file_fast_writing_proto_msgTypes,
	}.Build()
	File_fast_writing_proto = out.File
	file_fast_writing_proto_rawDesc = nil
	file_fast_writing_proto_goTypes = nil
	file_fast_writing_proto_depIdxs = nil
}
