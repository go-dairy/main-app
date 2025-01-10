// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: dairy/dairy.proto

package godairyv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordUuid    string                 `protobuf:"bytes,1,opt,name=record_uuid,json=recordUuid,proto3" json:"record_uuid,omitempty"` //ID of database record
	Datetime      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`                       //time and date of new task
	Task          string                 `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`                               //message of task
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_dairy_dairy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetRecordUuid() string {
	if x != nil {
		return x.RecordUuid
	}
	return ""
}

func (x *Task) GetDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

func (x *Task) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type TaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordUuid    string                 `protobuf:"bytes,1,opt,name=record_uuid,json=recordUuid,proto3" json:"record_uuid,omitempty"` //ID of database record
	Datetime      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`                       //time and date of new task
	Task          string                 `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`                               //message of task
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	mi := &file_dairy_dairy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{1}
}

func (x *TaskRequest) GetRecordUuid() string {
	if x != nil {
		return x.RecordUuid
	}
	return ""
}

func (x *TaskRequest) GetDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

func (x *TaskRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type TaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"` //planning tasks
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	mi := &file_dairy_dairy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{2}
}

func (x *TaskResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type NewTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordUuid    string                 `protobuf:"bytes,1,opt,name=record_uuid,json=recordUuid,proto3" json:"record_uuid,omitempty"` //ID of database record
	Datetime      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`                       //time and date of new task
	Task          string                 `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`                               //message of task
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewTaskRequest) Reset() {
	*x = NewTaskRequest{}
	mi := &file_dairy_dairy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTaskRequest) ProtoMessage() {}

func (x *NewTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTaskRequest.ProtoReflect.Descriptor instead.
func (*NewTaskRequest) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{3}
}

func (x *NewTaskRequest) GetRecordUuid() string {
	if x != nil {
		return x.RecordUuid
	}
	return ""
}

func (x *NewTaskRequest) GetDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

func (x *NewTaskRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type NewTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordUuid    string                 `protobuf:"bytes,1,opt,name=record_uuid,json=recordUuid,proto3" json:"record_uuid,omitempty"` //ID of database record
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewTaskResponse) Reset() {
	*x = NewTaskResponse{}
	mi := &file_dairy_dairy_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTaskResponse) ProtoMessage() {}

func (x *NewTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTaskResponse.ProtoReflect.Descriptor instead.
func (*NewTaskResponse) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{4}
}

func (x *NewTaskResponse) GetRecordUuid() string {
	if x != nil {
		return x.RecordUuid
	}
	return ""
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordUuid    string                 `protobuf:"bytes,1,opt,name=record_uuid,json=recordUuid,proto3" json:"record_uuid,omitempty"` //ID of database record
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_dairy_dairy_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskRequest) GetRecordUuid() string {
	if x != nil {
		return x.RecordUuid
	}
	return ""
}

type DateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DateUuid      string                 `protobuf:"bytes,1,opt,name=date_uuid,json=dateUuid,proto3" json:"date_uuid,omitempty"` //ID of database record
	Date          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`                         //selected date
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateRequest) Reset() {
	*x = DateRequest{}
	mi := &file_dairy_dairy_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRequest) ProtoMessage() {}

func (x *DateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRequest.ProtoReflect.Descriptor instead.
func (*DateRequest) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{6}
}

func (x *DateRequest) GetDateUuid() string {
	if x != nil {
		return x.DateUuid
	}
	return ""
}

func (x *DateRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type DateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DateUuid      string                 `protobuf:"bytes,1,opt,name=date_uuid,json=dateUuid,proto3" json:"date_uuid,omitempty"` //ID of database record
	Date          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`                         //selected date
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateResponse) Reset() {
	*x = DateResponse{}
	mi := &file_dairy_dairy_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateResponse) ProtoMessage() {}

func (x *DateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dairy_dairy_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateResponse.ProtoReflect.Descriptor instead.
func (*DateResponse) Descriptor() ([]byte, []int) {
	return file_dairy_dairy_proto_rawDescGZIP(), []int{7}
}

func (x *DateResponse) GetDateUuid() string {
	if x != nil {
		return x.DateUuid
	}
	return ""
}

func (x *DateResponse) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_dairy_dairy_proto protoreflect.FileDescriptor

var file_dairy_dairy_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2f, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x61, 0x69, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x7a, 0x0a, 0x0b, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x75, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x31, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x7d, 0x0a, 0x0e, 0x4e, 0x65,
	0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x75, 0x69, 0x64, 0x12, 0x36, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x32, 0x0a, 0x0f, 0x4e, 0x65, 0x77,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55,
	0x75, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x5b, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0xd4, 0x03, 0x0a,
	0x07, 0x47, 0x6f, 0x64, 0x61, 0x69, 0x72, 0x79, 0x12, 0x4a, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x4e, 0x65, 0x77, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64, 0x61, 0x69,
	0x72, 0x79, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x12, 0x44, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x4f, 0x66, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x69, 0x72,
	0x79, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x12, 0x4c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x12, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x12, 0x13, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x7d, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x6f, 0x64, 0x61, 0x69, 0x72, 0x79, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x3b, 0x67, 0x6f, 0x64, 0x61, 0x69, 0x72, 0x79, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dairy_dairy_proto_rawDescOnce sync.Once
	file_dairy_dairy_proto_rawDescData = file_dairy_dairy_proto_rawDesc
)

func file_dairy_dairy_proto_rawDescGZIP() []byte {
	file_dairy_dairy_proto_rawDescOnce.Do(func() {
		file_dairy_dairy_proto_rawDescData = protoimpl.X.CompressGZIP(file_dairy_dairy_proto_rawDescData)
	})
	return file_dairy_dairy_proto_rawDescData
}

var file_dairy_dairy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dairy_dairy_proto_goTypes = []any{
	(*Task)(nil),                  // 0: dairy.Task
	(*TaskRequest)(nil),           // 1: dairy.TaskRequest
	(*TaskResponse)(nil),          // 2: dairy.TaskResponse
	(*NewTaskRequest)(nil),        // 3: dairy.NewTaskRequest
	(*NewTaskResponse)(nil),       // 4: dairy.NewTaskResponse
	(*DeleteTaskRequest)(nil),     // 5: dairy.DeleteTaskRequest
	(*DateRequest)(nil),           // 6: dairy.DateRequest
	(*DateResponse)(nil),          // 7: dairy.DateResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_dairy_dairy_proto_depIdxs = []int32{
	8,  // 0: dairy.Task.datetime:type_name -> google.protobuf.Timestamp
	8,  // 1: dairy.TaskRequest.datetime:type_name -> google.protobuf.Timestamp
	0,  // 2: dairy.TaskResponse.tasks:type_name -> dairy.Task
	8,  // 3: dairy.NewTaskRequest.datetime:type_name -> google.protobuf.Timestamp
	8,  // 4: dairy.DateRequest.date:type_name -> google.protobuf.Timestamp
	8,  // 5: dairy.DateResponse.date:type_name -> google.protobuf.Timestamp
	3,  // 6: dairy.Godairy.NewTask:input_type -> dairy.NewTaskRequest
	6,  // 7: dairy.Godairy.NewDate:input_type -> dairy.DateRequest
	6,  // 8: dairy.Godairy.Date:input_type -> dairy.DateRequest
	6,  // 9: dairy.Godairy.TasksOfDate:input_type -> dairy.DateRequest
	1,  // 10: dairy.Godairy.Task:input_type -> dairy.TaskRequest
	5,  // 11: dairy.Godairy.DeleteTask:input_type -> dairy.DeleteTaskRequest
	4,  // 12: dairy.Godairy.NewTask:output_type -> dairy.NewTaskResponse
	7,  // 13: dairy.Godairy.NewDate:output_type -> dairy.DateResponse
	7,  // 14: dairy.Godairy.Date:output_type -> dairy.DateResponse
	2,  // 15: dairy.Godairy.TasksOfDate:output_type -> dairy.TaskResponse
	2,  // 16: dairy.Godairy.Task:output_type -> dairy.TaskResponse
	9,  // 17: dairy.Godairy.DeleteTask:output_type -> google.protobuf.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_dairy_dairy_proto_init() }
func file_dairy_dairy_proto_init() {
	if File_dairy_dairy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dairy_dairy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dairy_dairy_proto_goTypes,
		DependencyIndexes: file_dairy_dairy_proto_depIdxs,
		MessageInfos:      file_dairy_dairy_proto_msgTypes,
	}.Build()
	File_dairy_dairy_proto = out.File
	file_dairy_dairy_proto_rawDesc = nil
	file_dairy_dairy_proto_goTypes = nil
	file_dairy_dairy_proto_depIdxs = nil
}
