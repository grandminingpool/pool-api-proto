// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: filters.proto

package filtersProto

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

type UInt32RangeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Operators:
	//
	//	*UInt32RangeFilter_Equals
	//	*UInt32RangeFilter_Min
	//	*UInt32RangeFilter_Max
	Operators isUInt32RangeFilter_Operators `protobuf_oneof:"operators"`
}

func (x *UInt32RangeFilter) Reset() {
	*x = UInt32RangeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt32RangeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt32RangeFilter) ProtoMessage() {}

func (x *UInt32RangeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt32RangeFilter.ProtoReflect.Descriptor instead.
func (*UInt32RangeFilter) Descriptor() ([]byte, []int) {
	return file_filters_proto_rawDescGZIP(), []int{0}
}

func (m *UInt32RangeFilter) GetOperators() isUInt32RangeFilter_Operators {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (x *UInt32RangeFilter) GetEquals() uint32 {
	if x, ok := x.GetOperators().(*UInt32RangeFilter_Equals); ok {
		return x.Equals
	}
	return 0
}

func (x *UInt32RangeFilter) GetMin() uint32 {
	if x, ok := x.GetOperators().(*UInt32RangeFilter_Min); ok {
		return x.Min
	}
	return 0
}

func (x *UInt32RangeFilter) GetMax() uint32 {
	if x, ok := x.GetOperators().(*UInt32RangeFilter_Max); ok {
		return x.Max
	}
	return 0
}

type isUInt32RangeFilter_Operators interface {
	isUInt32RangeFilter_Operators()
}

type UInt32RangeFilter_Equals struct {
	Equals uint32 `protobuf:"varint,1,opt,name=equals,proto3,oneof"`
}

type UInt32RangeFilter_Min struct {
	Min uint32 `protobuf:"varint,2,opt,name=min,proto3,oneof"`
}

type UInt32RangeFilter_Max struct {
	Max uint32 `protobuf:"varint,3,opt,name=max,proto3,oneof"`
}

func (*UInt32RangeFilter_Equals) isUInt32RangeFilter_Operators() {}

func (*UInt32RangeFilter_Min) isUInt32RangeFilter_Operators() {}

func (*UInt32RangeFilter_Max) isUInt32RangeFilter_Operators() {}

type UInt64RangeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Operators:
	//
	//	*UInt64RangeFilter_Min
	//	*UInt64RangeFilter_Max
	Operators isUInt64RangeFilter_Operators `protobuf_oneof:"operators"`
}

func (x *UInt64RangeFilter) Reset() {
	*x = UInt64RangeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt64RangeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt64RangeFilter) ProtoMessage() {}

func (x *UInt64RangeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt64RangeFilter.ProtoReflect.Descriptor instead.
func (*UInt64RangeFilter) Descriptor() ([]byte, []int) {
	return file_filters_proto_rawDescGZIP(), []int{1}
}

func (m *UInt64RangeFilter) GetOperators() isUInt64RangeFilter_Operators {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (x *UInt64RangeFilter) GetMin() uint64 {
	if x, ok := x.GetOperators().(*UInt64RangeFilter_Min); ok {
		return x.Min
	}
	return 0
}

func (x *UInt64RangeFilter) GetMax() uint64 {
	if x, ok := x.GetOperators().(*UInt64RangeFilter_Max); ok {
		return x.Max
	}
	return 0
}

type isUInt64RangeFilter_Operators interface {
	isUInt64RangeFilter_Operators()
}

type UInt64RangeFilter_Min struct {
	Min uint64 `protobuf:"varint,1,opt,name=min,proto3,oneof"`
}

type UInt64RangeFilter_Max struct {
	Max uint64 `protobuf:"varint,2,opt,name=max,proto3,oneof"`
}

func (*UInt64RangeFilter_Min) isUInt64RangeFilter_Operators() {}

func (*UInt64RangeFilter_Max) isUInt64RangeFilter_Operators() {}

type BigIntRangeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Operators:
	//
	//	*BigIntRangeFilter_Min
	//	*BigIntRangeFilter_Max
	Operators isBigIntRangeFilter_Operators `protobuf_oneof:"operators"`
}

func (x *BigIntRangeFilter) Reset() {
	*x = BigIntRangeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigIntRangeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigIntRangeFilter) ProtoMessage() {}

func (x *BigIntRangeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigIntRangeFilter.ProtoReflect.Descriptor instead.
func (*BigIntRangeFilter) Descriptor() ([]byte, []int) {
	return file_filters_proto_rawDescGZIP(), []int{2}
}

func (m *BigIntRangeFilter) GetOperators() isBigIntRangeFilter_Operators {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (x *BigIntRangeFilter) GetMin() []byte {
	if x, ok := x.GetOperators().(*BigIntRangeFilter_Min); ok {
		return x.Min
	}
	return nil
}

func (x *BigIntRangeFilter) GetMax() []byte {
	if x, ok := x.GetOperators().(*BigIntRangeFilter_Max); ok {
		return x.Max
	}
	return nil
}

type isBigIntRangeFilter_Operators interface {
	isBigIntRangeFilter_Operators()
}

type BigIntRangeFilter_Min struct {
	Min []byte `protobuf:"bytes,1,opt,name=min,proto3,oneof"`
}

type BigIntRangeFilter_Max struct {
	Max []byte `protobuf:"bytes,2,opt,name=max,proto3,oneof"`
}

func (*BigIntRangeFilter_Min) isBigIntRangeFilter_Operators() {}

func (*BigIntRangeFilter_Max) isBigIntRangeFilter_Operators() {}

type DateFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Operators:
	//
	//	*DateFilter_Start
	//	*DateFilter_End
	Operators isDateFilter_Operators `protobuf_oneof:"operators"`
}

func (x *DateFilter) Reset() {
	*x = DateFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateFilter) ProtoMessage() {}

func (x *DateFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateFilter.ProtoReflect.Descriptor instead.
func (*DateFilter) Descriptor() ([]byte, []int) {
	return file_filters_proto_rawDescGZIP(), []int{3}
}

func (m *DateFilter) GetOperators() isDateFilter_Operators {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (x *DateFilter) GetStart() *timestamppb.Timestamp {
	if x, ok := x.GetOperators().(*DateFilter_Start); ok {
		return x.Start
	}
	return nil
}

func (x *DateFilter) GetEnd() *timestamppb.Timestamp {
	if x, ok := x.GetOperators().(*DateFilter_End); ok {
		return x.End
	}
	return nil
}

type isDateFilter_Operators interface {
	isDateFilter_Operators()
}

type DateFilter_Start struct {
	Start *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start,proto3,oneof"`
}

type DateFilter_End struct {
	End *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end,proto3,oneof"`
}

func (*DateFilter_Start) isDateFilter_Operators() {}

func (*DateFilter_End) isDateFilter_Operators() {}

var File_filters_proto protoreflect.FileDescriptor

var file_filters_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x11, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x06, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x06, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x48, 0x0a,
	0x11, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48,
	0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x48, 0x0a, 0x11, 0x42, 0x69, 0x67, 0x49, 0x6e,
	0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x03,
	0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e,
	0x12, 0x12, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x22, 0x7d, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x72, 0x61, 0x6e, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70,
	0x6f, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x3b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filters_proto_rawDescOnce sync.Once
	file_filters_proto_rawDescData = file_filters_proto_rawDesc
)

func file_filters_proto_rawDescGZIP() []byte {
	file_filters_proto_rawDescOnce.Do(func() {
		file_filters_proto_rawDescData = protoimpl.X.CompressGZIP(file_filters_proto_rawDescData)
	})
	return file_filters_proto_rawDescData
}

var file_filters_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_filters_proto_goTypes = []interface{}{
	(*UInt32RangeFilter)(nil),     // 0: filters.UInt32RangeFilter
	(*UInt64RangeFilter)(nil),     // 1: filters.UInt64RangeFilter
	(*BigIntRangeFilter)(nil),     // 2: filters.BigIntRangeFilter
	(*DateFilter)(nil),            // 3: filters.DateFilter
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_filters_proto_depIdxs = []int32{
	4, // 0: filters.DateFilter.start:type_name -> google.protobuf.Timestamp
	4, // 1: filters.DateFilter.end:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_filters_proto_init() }
func file_filters_proto_init() {
	if File_filters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt32RangeFilter); i {
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
		file_filters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt64RangeFilter); i {
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
		file_filters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigIntRangeFilter); i {
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
		file_filters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateFilter); i {
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
	file_filters_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UInt32RangeFilter_Equals)(nil),
		(*UInt32RangeFilter_Min)(nil),
		(*UInt32RangeFilter_Max)(nil),
	}
	file_filters_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UInt64RangeFilter_Min)(nil),
		(*UInt64RangeFilter_Max)(nil),
	}
	file_filters_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BigIntRangeFilter_Min)(nil),
		(*BigIntRangeFilter_Max)(nil),
	}
	file_filters_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*DateFilter_Start)(nil),
		(*DateFilter_End)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filters_proto_goTypes,
		DependencyIndexes: file_filters_proto_depIdxs,
		MessageInfos:      file_filters_proto_msgTypes,
	}.Build()
	File_filters_proto = out.File
	file_filters_proto_rawDesc = nil
	file_filters_proto_goTypes = nil
	file_filters_proto_depIdxs = nil
}
