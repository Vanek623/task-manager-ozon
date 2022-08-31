// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: storage/api.storage.proto

package storage

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description *string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *TaskAddRequest) Reset() {
	*x = TaskAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskAddRequest) ProtoMessage() {}

func (x *TaskAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskAddRequest.ProtoReflect.Descriptor instead.
func (*TaskAddRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{0}
}

func (x *TaskAddRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskAddRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type TaskAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TaskAddResponse) Reset() {
	*x = TaskAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskAddResponse) ProtoMessage() {}

func (x *TaskAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskAddResponse.ProtoReflect.Descriptor instead.
func (*TaskAddResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{1}
}

func (x *TaskAddResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type TaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{2}
}

func (x *TaskListRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TaskListRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TaskListResponse_Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{3}
}

func (x *TaskListResponse) GetTasks() []*TaskListResponse_Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TaskUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *TaskUpdateRequest) Reset() {
	*x = TaskUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskUpdateRequest) ProtoMessage() {}

func (x *TaskUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskUpdateRequest.ProtoReflect.Descriptor instead.
func (*TaskUpdateRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{4}
}

func (x *TaskUpdateRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TaskUpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskUpdateRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type TaskUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskUpdateResponse) Reset() {
	*x = TaskUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskUpdateResponse) ProtoMessage() {}

func (x *TaskUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskUpdateResponse.ProtoReflect.Descriptor instead.
func (*TaskUpdateResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{5}
}

type TaskDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TaskDeleteRequest) Reset() {
	*x = TaskDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDeleteRequest) ProtoMessage() {}

func (x *TaskDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDeleteRequest.ProtoReflect.Descriptor instead.
func (*TaskDeleteRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{6}
}

func (x *TaskDeleteRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type TaskDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskDeleteResponse) Reset() {
	*x = TaskDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDeleteResponse) ProtoMessage() {}

func (x *TaskDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDeleteResponse.ProtoReflect.Descriptor instead.
func (*TaskDeleteResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{7}
}

type TaskGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TaskGetRequest) Reset() {
	*x = TaskGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskGetRequest) ProtoMessage() {}

func (x *TaskGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskGetRequest.ProtoReflect.Descriptor instead.
func (*TaskGetRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{8}
}

func (x *TaskGetRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type TaskGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description *string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Edited      int64   `protobuf:"varint,3,opt,name=edited,proto3" json:"edited,omitempty"`
	Created     int64   `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *TaskGetResponse) Reset() {
	*x = TaskGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskGetResponse) ProtoMessage() {}

func (x *TaskGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskGetResponse.ProtoReflect.Descriptor instead.
func (*TaskGetResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{9}
}

func (x *TaskGetResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskGetResponse) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *TaskGetResponse) GetEdited() int64 {
	if x != nil {
		return x.Edited
	}
	return 0
}

func (x *TaskGetResponse) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

type TaskListResponse_Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *TaskListResponse_Task) Reset() {
	*x = TaskListResponse_Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_storage_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse_Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse_Task) ProtoMessage() {}

func (x *TaskListResponse_Task) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_storage_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse_Task.ProtoReflect.Descriptor instead.
func (*TaskListResponse_Task) Descriptor() ([]byte, []int) {
	return file_storage_api_storage_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TaskListResponse_Task) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TaskListResponse_Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_storage_api_storage_proto protoreflect.FileDescriptor

var file_storage_api_storage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b, 0x36, 0x32, 0x33, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x5d, 0x0a, 0x0e, 0x54,
	0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x54, 0x61,
	0x73, 0x6b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3f, 0x0a,
	0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x9d,
	0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x45, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61,
	0x6e, 0x65, 0x6b, 0x36, 0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x1a, 0x2c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x70,
	0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x14, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x14, 0x0a, 0x12, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x90, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xe9, 0x05, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x12, 0x3e,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b, 0x36,
	0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b, 0x36,
	0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x8f, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b, 0x36,
	0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x40, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b,
	0x36, 0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x95, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x41, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61,
	0x6e, 0x65, 0x6b, 0x36, 0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b, 0x36, 0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x95, 0x01, 0x0a, 0x0a,
	0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x41, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b, 0x36, 0x32, 0x33, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b, 0x36, 0x32,
	0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x12,
	0x3e, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b,
	0x36, 0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3f, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x61, 0x6e, 0x65, 0x6b,
	0x36, 0x32, 0x33, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x56, 0x61, 0x6e, 0x65, 0x6b, 0x36, 0x32, 0x33, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_storage_api_storage_proto_rawDescOnce sync.Once
	file_storage_api_storage_proto_rawDescData = file_storage_api_storage_proto_rawDesc
)

func file_storage_api_storage_proto_rawDescGZIP() []byte {
	file_storage_api_storage_proto_rawDescOnce.Do(func() {
		file_storage_api_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_api_storage_proto_rawDescData)
	})
	return file_storage_api_storage_proto_rawDescData
}

var file_storage_api_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_storage_api_storage_proto_goTypes = []interface{}{
	(*TaskAddRequest)(nil),        // 0: ozon.dev.vanek623.task_manager_bot.api.storage.TaskAddRequest
	(*TaskAddResponse)(nil),       // 1: ozon.dev.vanek623.task_manager_bot.api.storage.TaskAddResponse
	(*TaskListRequest)(nil),       // 2: ozon.dev.vanek623.task_manager_bot.api.storage.TaskListRequest
	(*TaskListResponse)(nil),      // 3: ozon.dev.vanek623.task_manager_bot.api.storage.TaskListResponse
	(*TaskUpdateRequest)(nil),     // 4: ozon.dev.vanek623.task_manager_bot.api.storage.TaskUpdateRequest
	(*TaskUpdateResponse)(nil),    // 5: ozon.dev.vanek623.task_manager_bot.api.storage.TaskUpdateResponse
	(*TaskDeleteRequest)(nil),     // 6: ozon.dev.vanek623.task_manager_bot.api.storage.TaskDeleteRequest
	(*TaskDeleteResponse)(nil),    // 7: ozon.dev.vanek623.task_manager_bot.api.storage.TaskDeleteResponse
	(*TaskGetRequest)(nil),        // 8: ozon.dev.vanek623.task_manager_bot.api.storage.TaskGetRequest
	(*TaskGetResponse)(nil),       // 9: ozon.dev.vanek623.task_manager_bot.api.storage.TaskGetResponse
	(*TaskListResponse_Task)(nil), // 10: ozon.dev.vanek623.task_manager_bot.api.storage.TaskListResponse.Task
}
var file_storage_api_storage_proto_depIdxs = []int32{
	10, // 0: ozon.dev.vanek623.task_manager_bot.api.storage.TaskListResponse.tasks:type_name -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskListResponse.Task
	0,  // 1: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskAdd:input_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskAddRequest
	2,  // 2: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskList:input_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskListRequest
	4,  // 3: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskUpdate:input_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskUpdateRequest
	6,  // 4: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskDelete:input_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskDeleteRequest
	8,  // 5: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskGet:input_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskGetRequest
	1,  // 6: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskAdd:output_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskAddResponse
	3,  // 7: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskList:output_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskListResponse
	5,  // 8: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskUpdate:output_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskUpdateResponse
	7,  // 9: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskDelete:output_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskDeleteResponse
	9,  // 10: ozon.dev.vanek623.task_manager_bot.api.storage.Storage.TaskGet:output_type -> ozon.dev.vanek623.task_manager_bot.api.storage.TaskGetResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_storage_api_storage_proto_init() }
func file_storage_api_storage_proto_init() {
	if File_storage_api_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_api_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskAddRequest); i {
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
		file_storage_api_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskAddResponse); i {
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
		file_storage_api_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListRequest); i {
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
		file_storage_api_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResponse); i {
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
		file_storage_api_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskUpdateRequest); i {
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
		file_storage_api_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskUpdateResponse); i {
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
		file_storage_api_storage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDeleteRequest); i {
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
		file_storage_api_storage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDeleteResponse); i {
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
		file_storage_api_storage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskGetRequest); i {
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
		file_storage_api_storage_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskGetResponse); i {
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
		file_storage_api_storage_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResponse_Task); i {
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
	file_storage_api_storage_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_storage_api_storage_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_storage_api_storage_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_api_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_api_storage_proto_goTypes,
		DependencyIndexes: file_storage_api_storage_proto_depIdxs,
		MessageInfos:      file_storage_api_storage_proto_msgTypes,
	}.Build()
	File_storage_api_storage_proto = out.File
	file_storage_api_storage_proto_rawDesc = nil
	file_storage_api_storage_proto_goTypes = nil
	file_storage_api_storage_proto_depIdxs = nil
}
