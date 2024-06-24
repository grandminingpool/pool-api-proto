// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: pool_miners.proto

package poolMinersProto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type MinerAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *MinerAddressRequest) Reset() {
	*x = MinerAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_miners_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinerAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinerAddressRequest) ProtoMessage() {}

func (x *MinerAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pool_miners_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinerAddressRequest.ProtoReflect.Descriptor instead.
func (*MinerAddressRequest) Descriptor() ([]byte, []int) {
	return file_pool_miners_proto_rawDescGZIP(), []int{0}
}

func (x *MinerAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type MinerAddressesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *MinerAddressesRequest) Reset() {
	*x = MinerAddressesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_miners_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinerAddressesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinerAddressesRequest) ProtoMessage() {}

func (x *MinerAddressesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pool_miners_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinerAddressesRequest.ProtoReflect.Descriptor instead.
func (*MinerAddressesRequest) Descriptor() ([]byte, []int) {
	return file_pool_miners_proto_rawDescGZIP(), []int{1}
}

func (x *MinerAddressesRequest) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type ValidateAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ValidateAddressResponse) Reset() {
	*x = ValidateAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_miners_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAddressResponse) ProtoMessage() {}

func (x *ValidateAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pool_miners_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAddressResponse.ProtoReflect.Descriptor instead.
func (*ValidateAddressResponse) Descriptor() ([]byte, []int) {
	return file_pool_miners_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateAddressResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type MinerWorker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker      string                 `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
	SlaveRegion string                 `protobuf:"bytes,2,opt,name=slave_region,json=slaveRegion,proto3" json:"slave_region,omitempty"`
	Solo        bool                   `protobuf:"varint,3,opt,name=solo,proto3" json:"solo,omitempty"`
	Hashrate    []byte                 `protobuf:"bytes,4,opt,name=hashrate,proto3" json:"hashrate,omitempty"`
	ConnectedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=connected_at,json=connectedAt,proto3" json:"connected_at,omitempty"`
}

func (x *MinerWorker) Reset() {
	*x = MinerWorker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_miners_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinerWorker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinerWorker) ProtoMessage() {}

func (x *MinerWorker) ProtoReflect() protoreflect.Message {
	mi := &file_pool_miners_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinerWorker.ProtoReflect.Descriptor instead.
func (*MinerWorker) Descriptor() ([]byte, []int) {
	return file_pool_miners_proto_rawDescGZIP(), []int{3}
}

func (x *MinerWorker) GetWorker() string {
	if x != nil {
		return x.Worker
	}
	return ""
}

func (x *MinerWorker) GetSlaveRegion() string {
	if x != nil {
		return x.SlaveRegion
	}
	return ""
}

func (x *MinerWorker) GetSolo() bool {
	if x != nil {
		return x.Solo
	}
	return false
}

func (x *MinerWorker) GetHashrate() []byte {
	if x != nil {
		return x.Hashrate
	}
	return nil
}

func (x *MinerWorker) GetConnectedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ConnectedAt
	}
	return nil
}

type MinerWorkers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workers []*MinerWorker `protobuf:"bytes,1,rep,name=workers,proto3" json:"workers,omitempty"`
}

func (x *MinerWorkers) Reset() {
	*x = MinerWorkers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_miners_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinerWorkers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinerWorkers) ProtoMessage() {}

func (x *MinerWorkers) ProtoReflect() protoreflect.Message {
	mi := &file_pool_miners_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinerWorkers.ProtoReflect.Descriptor instead.
func (*MinerWorkers) Descriptor() ([]byte, []int) {
	return file_pool_miners_proto_rawDescGZIP(), []int{4}
}

func (x *MinerWorkers) GetWorkers() []*MinerWorker {
	if x != nil {
		return x.Workers
	}
	return nil
}

type MinersWorkers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workers map[string]*MinerWorkers `protobuf:"bytes,1,rep,name=workers,proto3" json:"workers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MinersWorkers) Reset() {
	*x = MinersWorkers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_miners_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinersWorkers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinersWorkers) ProtoMessage() {}

func (x *MinersWorkers) ProtoReflect() protoreflect.Message {
	mi := &file_pool_miners_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinersWorkers.ProtoReflect.Descriptor instead.
func (*MinersWorkers) Descriptor() ([]byte, []int) {
	return file_pool_miners_proto_rawDescGZIP(), []int{5}
}

func (x *MinersWorkers) GetWorkers() map[string]*MinerWorkers {
	if x != nil {
		return x.Workers
	}
	return nil
}

var File_pool_miners_proto protoreflect.FileDescriptor

var file_pool_miners_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2f, 0x0a, 0x13, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x35, 0x0a, 0x15, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x17, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x4d,
	0x69, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x6c, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x6f, 0x6c, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73,
	0x68, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x68, 0x61, 0x73,
	0x68, 0x72, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x4d, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x1a, 0x55, 0x0a,
	0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x69, 0x6e,
	0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x8c, 0x02, 0x0a, 0x11, 0x50, 0x6f, 0x6f, 0x6c, 0x4d, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0f, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x2e,
	0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x65,
	0x72, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f,
	0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f,
	0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x3b, 0x70, 0x6f, 0x6f, 0x6c, 0x4d, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pool_miners_proto_rawDescOnce sync.Once
	file_pool_miners_proto_rawDescData = file_pool_miners_proto_rawDesc
)

func file_pool_miners_proto_rawDescGZIP() []byte {
	file_pool_miners_proto_rawDescOnce.Do(func() {
		file_pool_miners_proto_rawDescData = protoimpl.X.CompressGZIP(file_pool_miners_proto_rawDescData)
	})
	return file_pool_miners_proto_rawDescData
}

var file_pool_miners_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pool_miners_proto_goTypes = []interface{}{
	(*MinerAddressRequest)(nil),     // 0: pool_miners.MinerAddressRequest
	(*MinerAddressesRequest)(nil),   // 1: pool_miners.MinerAddressesRequest
	(*ValidateAddressResponse)(nil), // 2: pool_miners.ValidateAddressResponse
	(*MinerWorker)(nil),             // 3: pool_miners.MinerWorker
	(*MinerWorkers)(nil),            // 4: pool_miners.MinerWorkers
	(*MinersWorkers)(nil),           // 5: pool_miners.MinersWorkers
	nil,                             // 6: pool_miners.MinersWorkers.WorkersEntry
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_pool_miners_proto_depIdxs = []int32{
	7, // 0: pool_miners.MinerWorker.connected_at:type_name -> google.protobuf.Timestamp
	3, // 1: pool_miners.MinerWorkers.workers:type_name -> pool_miners.MinerWorker
	6, // 2: pool_miners.MinersWorkers.workers:type_name -> pool_miners.MinersWorkers.WorkersEntry
	4, // 3: pool_miners.MinersWorkers.WorkersEntry.value:type_name -> pool_miners.MinerWorkers
	0, // 4: pool_miners.PoolMinersService.ValidateAddress:input_type -> pool_miners.MinerAddressRequest
	0, // 5: pool_miners.PoolMinersService.GetMinerWorkers:input_type -> pool_miners.MinerAddressRequest
	1, // 6: pool_miners.PoolMinersService.GetWorkers:input_type -> pool_miners.MinerAddressesRequest
	2, // 7: pool_miners.PoolMinersService.ValidateAddress:output_type -> pool_miners.ValidateAddressResponse
	4, // 8: pool_miners.PoolMinersService.GetMinerWorkers:output_type -> pool_miners.MinerWorkers
	5, // 9: pool_miners.PoolMinersService.GetWorkers:output_type -> pool_miners.MinersWorkers
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pool_miners_proto_init() }
func file_pool_miners_proto_init() {
	if File_pool_miners_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pool_miners_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinerAddressRequest); i {
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
		file_pool_miners_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinerAddressesRequest); i {
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
		file_pool_miners_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAddressResponse); i {
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
		file_pool_miners_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinerWorker); i {
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
		file_pool_miners_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinerWorkers); i {
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
		file_pool_miners_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinersWorkers); i {
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
			RawDescriptor: file_pool_miners_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pool_miners_proto_goTypes,
		DependencyIndexes: file_pool_miners_proto_depIdxs,
		MessageInfos:      file_pool_miners_proto_msgTypes,
	}.Build()
	File_pool_miners_proto = out.File
	file_pool_miners_proto_rawDesc = nil
	file_pool_miners_proto_goTypes = nil
	file_pool_miners_proto_depIdxs = nil
}
