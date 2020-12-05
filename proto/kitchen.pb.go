// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: kitchen.proto

package proto

import (
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

type TicketStatus int32

const (
	TicketStatus_TICKET_UNKNOWN         TicketStatus = 0
	TicketStatus_TICKET_CREATED         TicketStatus = 1
	TicketStatus_TICKET_APPROVED        TicketStatus = 2
	TicketStatus_TICKET_CREATE_REJECTED TicketStatus = 3
)

// Enum value maps for TicketStatus.
var (
	TicketStatus_name = map[int32]string{
		0: "TICKET_UNKNOWN",
		1: "TICKET_CREATED",
		2: "TICKET_APPROVED",
		3: "TICKET_CREATE_REJECTED",
	}
	TicketStatus_value = map[string]int32{
		"TICKET_UNKNOWN":         0,
		"TICKET_CREATED":         1,
		"TICKET_APPROVED":        2,
		"TICKET_CREATE_REJECTED": 3,
	}
)

func (x TicketStatus) Enum() *TicketStatus {
	p := new(TicketStatus)
	*p = x
	return p
}

func (x TicketStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TicketStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kitchen_proto_enumTypes[0].Descriptor()
}

func (TicketStatus) Type() protoreflect.EnumType {
	return &file_kitchen_proto_enumTypes[0]
}

func (x TicketStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TicketStatus.Descriptor instead.
func (TicketStatus) EnumDescriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{0}
}

type TicketEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId int64        `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status  TicketStatus `protobuf:"varint,3,opt,name=status,proto3,enum=proto.TicketStatus" json:"status,omitempty"`
}

func (x *TicketEntity) Reset() {
	*x = TicketEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketEntity) ProtoMessage() {}

func (x *TicketEntity) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketEntity.ProtoReflect.Descriptor instead.
func (*TicketEntity) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{0}
}

func (x *TicketEntity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TicketEntity) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *TicketEntity) GetStatus() TicketStatus {
	if x != nil {
		return x.Status
	}
	return TicketStatus_TICKET_UNKNOWN
}

type CreateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CreateTicketRequest) Reset() {
	*x = CreateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketRequest) ProtoMessage() {}

func (x *CreateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketRequest.ProtoReflect.Descriptor instead.
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTicketRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type CreateTicketReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *TicketEntity `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *CreateTicketReply) Reset() {
	*x = CreateTicketReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketReply) ProtoMessage() {}

func (x *CreateTicketReply) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketReply.ProtoReflect.Descriptor instead.
func (*CreateTicketReply) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTicketReply) GetTicket() *TicketEntity {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type ApproveTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId int64 `protobuf:"varint,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
}

func (x *ApproveTicketRequest) Reset() {
	*x = ApproveTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveTicketRequest) ProtoMessage() {}

func (x *ApproveTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveTicketRequest.ProtoReflect.Descriptor instead.
func (*ApproveTicketRequest) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{3}
}

func (x *ApproveTicketRequest) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type ApproveTicketReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApproveTicketReply) Reset() {
	*x = ApproveTicketReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveTicketReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveTicketReply) ProtoMessage() {}

func (x *ApproveTicketReply) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveTicketReply.ProtoReflect.Descriptor instead.
func (*ApproveTicketReply) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{4}
}

type RejectTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId int64 `protobuf:"varint,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
}

func (x *RejectTicketRequest) Reset() {
	*x = RejectTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectTicketRequest) ProtoMessage() {}

func (x *RejectTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectTicketRequest.ProtoReflect.Descriptor instead.
func (*RejectTicketRequest) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{5}
}

func (x *RejectTicketRequest) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type RejectTicketReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RejectTicketReply) Reset() {
	*x = RejectTicketReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectTicketReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectTicketReply) ProtoMessage() {}

func (x *RejectTicketReply) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectTicketReply.ProtoReflect.Descriptor instead.
func (*RejectTicketReply) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{6}
}

var File_kitchen_proto protoreflect.FileDescriptor

var file_kitchen_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x40, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x22, 0x33, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x0a,
	0x13, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49,
	0x64, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x67, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x49,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x32,
	0xe4, 0x01, 0x0a, 0x07, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x12, 0x46, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0c, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x6e, 0x34, 0x34, 0x30, 0x2f, 0x73, 0x61,
	0x67, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kitchen_proto_rawDescOnce sync.Once
	file_kitchen_proto_rawDescData = file_kitchen_proto_rawDesc
)

func file_kitchen_proto_rawDescGZIP() []byte {
	file_kitchen_proto_rawDescOnce.Do(func() {
		file_kitchen_proto_rawDescData = protoimpl.X.CompressGZIP(file_kitchen_proto_rawDescData)
	})
	return file_kitchen_proto_rawDescData
}

var file_kitchen_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kitchen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_kitchen_proto_goTypes = []interface{}{
	(TicketStatus)(0),            // 0: proto.TicketStatus
	(*TicketEntity)(nil),         // 1: proto.TicketEntity
	(*CreateTicketRequest)(nil),  // 2: proto.CreateTicketRequest
	(*CreateTicketReply)(nil),    // 3: proto.CreateTicketReply
	(*ApproveTicketRequest)(nil), // 4: proto.ApproveTicketRequest
	(*ApproveTicketReply)(nil),   // 5: proto.ApproveTicketReply
	(*RejectTicketRequest)(nil),  // 6: proto.RejectTicketRequest
	(*RejectTicketReply)(nil),    // 7: proto.RejectTicketReply
}
var file_kitchen_proto_depIdxs = []int32{
	0, // 0: proto.TicketEntity.status:type_name -> proto.TicketStatus
	1, // 1: proto.CreateTicketReply.ticket:type_name -> proto.TicketEntity
	2, // 2: proto.Kitchen.CreateTicket:input_type -> proto.CreateTicketRequest
	4, // 3: proto.Kitchen.ApproveTicket:input_type -> proto.ApproveTicketRequest
	6, // 4: proto.Kitchen.RejectTicket:input_type -> proto.RejectTicketRequest
	3, // 5: proto.Kitchen.CreateTicket:output_type -> proto.CreateTicketReply
	5, // 6: proto.Kitchen.ApproveTicket:output_type -> proto.ApproveTicketReply
	7, // 7: proto.Kitchen.RejectTicket:output_type -> proto.RejectTicketReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kitchen_proto_init() }
func file_kitchen_proto_init() {
	if File_kitchen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kitchen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketEntity); i {
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
		file_kitchen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketRequest); i {
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
		file_kitchen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketReply); i {
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
		file_kitchen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveTicketRequest); i {
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
		file_kitchen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveTicketReply); i {
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
		file_kitchen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectTicketRequest); i {
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
		file_kitchen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectTicketReply); i {
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
			RawDescriptor: file_kitchen_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kitchen_proto_goTypes,
		DependencyIndexes: file_kitchen_proto_depIdxs,
		EnumInfos:         file_kitchen_proto_enumTypes,
		MessageInfos:      file_kitchen_proto_msgTypes,
	}.Build()
	File_kitchen_proto = out.File
	file_kitchen_proto_rawDesc = nil
	file_kitchen_proto_goTypes = nil
	file_kitchen_proto_depIdxs = nil
}
