// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: permission/permission.proto

package sso

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddNewPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string           `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Settings map[string]string `protobuf:"bytes,2,rep,name=settings" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IsActual *uint32           `protobuf:"varint,3,opt,name=isActual" json:"isActual,omitempty"`
}

func (x *AddNewPermissionRequest) Reset() {
	*x = AddNewPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewPermissionRequest) ProtoMessage() {}

func (x *AddNewPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewPermissionRequest.ProtoReflect.Descriptor instead.
func (*AddNewPermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{0}
}

func (x *AddNewPermissionRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *AddNewPermissionRequest) GetSettings() map[string]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *AddNewPermissionRequest) GetIsActual() uint32 {
	if x != nil && x.IsActual != nil {
		return *x.IsActual
	}
	return 0
}

type PermissionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuids    []string          `protobuf:"bytes,1,rep,name=uuids" json:"uuids,omitempty"`
	Names    []string          `protobuf:"bytes,2,rep,name=names" json:"names,omitempty"`
	Settings map[string]string `protobuf:"bytes,3,rep,name=settings" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IsActual *uint32           `protobuf:"varint,4,opt,name=isActual" json:"isActual,omitempty"`
}

func (x *PermissionFilter) Reset() {
	*x = PermissionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionFilter) ProtoMessage() {}

func (x *PermissionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionFilter.ProtoReflect.Descriptor instead.
func (*PermissionFilter) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionFilter) GetUuids() []string {
	if x != nil {
		return x.Uuids
	}
	return nil
}

func (x *PermissionFilter) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *PermissionFilter) GetSettings() map[string]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *PermissionFilter) GetIsActual() uint32 {
	if x != nil && x.IsActual != nil {
		return *x.IsActual
	}
	return 0
}

type RemovePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuids   []string          `protobuf:"bytes,1,rep,name=uuids" json:"uuids,omitempty"`
	Filters *PermissionFilter `protobuf:"bytes,2,opt,name=filters" json:"filters,omitempty"`
}

func (x *RemovePermissionRequest) Reset() {
	*x = RemovePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionRequest) ProtoMessage() {}

func (x *RemovePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionRequest.ProtoReflect.Descriptor instead.
func (*RemovePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{2}
}

func (x *RemovePermissionRequest) GetUuids() []string {
	if x != nil {
		return x.Uuids
	}
	return nil
}

func (x *RemovePermissionRequest) GetFilters() *PermissionFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type RemovePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *string                       `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Items  []*PermissionListItemResponse `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *RemovePermissionResponse) Reset() {
	*x = RemovePermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionResponse) ProtoMessage() {}

func (x *RemovePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionResponse.ProtoReflect.Descriptor instead.
func (*RemovePermissionResponse) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{3}
}

func (x *RemovePermissionResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *RemovePermissionResponse) GetItems() []*PermissionListItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]*AddNewPermissionRequest `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *UpdatePermissionRequest) Reset() {
	*x = UpdatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionRequest) ProtoMessage() {}

func (x *UpdatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePermissionRequest) GetData() map[string]*AddNewPermissionRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdatePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *string                                `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Items  map[string]*PermissionListItemResponse `protobuf:"bytes,2,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *UpdatePermissionResponse) Reset() {
	*x = UpdatePermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionResponse) ProtoMessage() {}

func (x *UpdatePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionResponse.ProtoReflect.Descriptor instead.
func (*UpdatePermissionResponse) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePermissionResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *UpdatePermissionResponse) GetItems() map[string]*PermissionListItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type PermissionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartFrom *uint64           `protobuf:"varint,1,req,name=startFrom" json:"startFrom,omitempty"`
	PageSize  *uint64           `protobuf:"varint,2,req,name=pageSize" json:"pageSize,omitempty"`
	Filter    *PermissionFilter `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
}

func (x *PermissionListRequest) Reset() {
	*x = PermissionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionListRequest) ProtoMessage() {}

func (x *PermissionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionListRequest.ProtoReflect.Descriptor instead.
func (*PermissionListRequest) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{6}
}

func (x *PermissionListRequest) GetStartFrom() uint64 {
	if x != nil && x.StartFrom != nil {
		return *x.StartFrom
	}
	return 0
}

func (x *PermissionListRequest) GetPageSize() uint64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *PermissionListRequest) GetFilter() *PermissionFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type PermissionListItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      *string `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	Name      *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Settings  *string `protobuf:"bytes,3,req,name=settings" json:"settings,omitempty"`
	IsActual  *uint32 `protobuf:"varint,4,req,name=isActual" json:"isActual,omitempty"`
	CreatedAt *string `protobuf:"bytes,5,opt,name=createdAt" json:"createdAt,omitempty"`
}

func (x *PermissionListItemResponse) Reset() {
	*x = PermissionListItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionListItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionListItemResponse) ProtoMessage() {}

func (x *PermissionListItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionListItemResponse.ProtoReflect.Descriptor instead.
func (*PermissionListItemResponse) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{7}
}

func (x *PermissionListItemResponse) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *PermissionListItemResponse) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PermissionListItemResponse) GetSettings() string {
	if x != nil && x.Settings != nil {
		return *x.Settings
	}
	return ""
}

func (x *PermissionListItemResponse) GetIsActual() uint32 {
	if x != nil && x.IsActual != nil {
		return *x.IsActual
	}
	return 0
}

func (x *PermissionListItemResponse) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

type PermissionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount *uint64                       `protobuf:"varint,1,req,name=totalCount" json:"totalCount,omitempty"`
	Items      []*PermissionListItemResponse `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *PermissionListResponse) Reset() {
	*x = PermissionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_permission_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionListResponse) ProtoMessage() {}

func (x *PermissionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_permission_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionListResponse.ProtoReflect.Descriptor instead.
func (*PermissionListResponse) Descriptor() ([]byte, []int) {
	return file_permission_permission_proto_rawDescGZIP(), []int{8}
}

func (x *PermissionListResponse) GetTotalCount() uint64 {
	if x != nil && x.TotalCount != nil {
		return *x.TotalCount
	}
	return 0
}

func (x *PermissionListResponse) GetItems() []*PermissionListItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_permission_permission_proto protoreflect.FileDescriptor

var file_permission_permission_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x75, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x75, 0x61,
	0x6c, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xdf,
	0x01, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x75, 0x75, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x46, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x67, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75,
	0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x75, 0x69, 0x64,
	0x73, 0x12, 0x36, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x70, 0x0a, 0x18, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x5c, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xdb, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x45, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0x60, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x9a, 0x01, 0x0a, 0x1a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x76, 0x0a,
	0x16, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xf4, 0x02, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5d, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f,
	0x70, 0x61, 0x6e, 0x74, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x73, 0x6f,
}

var (
	file_permission_permission_proto_rawDescOnce sync.Once
	file_permission_permission_proto_rawDescData = file_permission_permission_proto_rawDesc
)

func file_permission_permission_proto_rawDescGZIP() []byte {
	file_permission_permission_proto_rawDescOnce.Do(func() {
		file_permission_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_permission_proto_rawDescData)
	})
	return file_permission_permission_proto_rawDescData
}

var file_permission_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_permission_permission_proto_goTypes = []interface{}{
	(*AddNewPermissionRequest)(nil),    // 0: permission.AddNewPermissionRequest
	(*PermissionFilter)(nil),           // 1: permission.PermissionFilter
	(*RemovePermissionRequest)(nil),    // 2: permission.RemovePermissionRequest
	(*RemovePermissionResponse)(nil),   // 3: permission.RemovePermissionResponse
	(*UpdatePermissionRequest)(nil),    // 4: permission.UpdatePermissionRequest
	(*UpdatePermissionResponse)(nil),   // 5: permission.UpdatePermissionResponse
	(*PermissionListRequest)(nil),      // 6: permission.PermissionListRequest
	(*PermissionListItemResponse)(nil), // 7: permission.PermissionListItemResponse
	(*PermissionListResponse)(nil),     // 8: permission.PermissionListResponse
	nil,                                // 9: permission.AddNewPermissionRequest.SettingsEntry
	nil,                                // 10: permission.PermissionFilter.SettingsEntry
	nil,                                // 11: permission.UpdatePermissionRequest.DataEntry
	nil,                                // 12: permission.UpdatePermissionResponse.ItemsEntry
	(*emptypb.Empty)(nil),              // 13: google.protobuf.Empty
}
var file_permission_permission_proto_depIdxs = []int32{
	9,  // 0: permission.AddNewPermissionRequest.settings:type_name -> permission.AddNewPermissionRequest.SettingsEntry
	10, // 1: permission.PermissionFilter.settings:type_name -> permission.PermissionFilter.SettingsEntry
	1,  // 2: permission.RemovePermissionRequest.filters:type_name -> permission.PermissionFilter
	7,  // 3: permission.RemovePermissionResponse.items:type_name -> permission.PermissionListItemResponse
	11, // 4: permission.UpdatePermissionRequest.data:type_name -> permission.UpdatePermissionRequest.DataEntry
	12, // 5: permission.UpdatePermissionResponse.items:type_name -> permission.UpdatePermissionResponse.ItemsEntry
	1,  // 6: permission.PermissionListRequest.filter:type_name -> permission.PermissionFilter
	7,  // 7: permission.PermissionListResponse.items:type_name -> permission.PermissionListItemResponse
	0,  // 8: permission.UpdatePermissionRequest.DataEntry.value:type_name -> permission.AddNewPermissionRequest
	7,  // 9: permission.UpdatePermissionResponse.ItemsEntry.value:type_name -> permission.PermissionListItemResponse
	0,  // 10: permission.Permission.AddNewPermission:input_type -> permission.AddNewPermissionRequest
	2,  // 11: permission.Permission.RemovePermission:input_type -> permission.RemovePermissionRequest
	4,  // 12: permission.Permission.UpdatePermission:input_type -> permission.UpdatePermissionRequest
	6,  // 13: permission.Permission.PermissionList:input_type -> permission.PermissionListRequest
	13, // 14: permission.Permission.AddNewPermission:output_type -> google.protobuf.Empty
	3,  // 15: permission.Permission.RemovePermission:output_type -> permission.RemovePermissionResponse
	5,  // 16: permission.Permission.UpdatePermission:output_type -> permission.UpdatePermissionResponse
	8,  // 17: permission.Permission.PermissionList:output_type -> permission.PermissionListResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_permission_permission_proto_init() }
func file_permission_permission_proto_init() {
	if File_permission_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_permission_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewPermissionRequest); i {
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
		file_permission_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionFilter); i {
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
		file_permission_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePermissionRequest); i {
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
		file_permission_permission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePermissionResponse); i {
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
		file_permission_permission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermissionRequest); i {
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
		file_permission_permission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermissionResponse); i {
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
		file_permission_permission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionListRequest); i {
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
		file_permission_permission_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionListItemResponse); i {
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
		file_permission_permission_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionListResponse); i {
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
			RawDescriptor: file_permission_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_permission_permission_proto_goTypes,
		DependencyIndexes: file_permission_permission_proto_depIdxs,
		MessageInfos:      file_permission_permission_proto_msgTypes,
	}.Build()
	File_permission_permission_proto = out.File
	file_permission_permission_proto_rawDesc = nil
	file_permission_permission_proto_goTypes = nil
	file_permission_permission_proto_depIdxs = nil
}