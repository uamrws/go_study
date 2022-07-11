// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: trip.proto

package trippb

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

type TripStatus int32

const (
	TripStatus_TS_NOT_SPECIFIED TripStatus = 0
	TripStatus_NOT_STARTED      TripStatus = 1
	TripStatus_IN_PROGRESS      TripStatus = 2
	TripStatus_FINISHED         TripStatus = 3
	TripStatus_PAID             TripStatus = 4
)

// Enum value maps for TripStatus.
var (
	TripStatus_name = map[int32]string{
		0: "TS_NOT_SPECIFIED",
		1: "NOT_STARTED",
		2: "IN_PROGRESS",
		3: "FINISHED",
		4: "PAID",
	}
	TripStatus_value = map[string]int32{
		"TS_NOT_SPECIFIED": 0,
		"NOT_STARTED":      1,
		"IN_PROGRESS":      2,
		"FINISHED":         3,
		"PAID":             4,
	}
)

func (x TripStatus) Enum() *TripStatus {
	p := new(TripStatus)
	*p = x
	return p
}

func (x TripStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TripStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_trip_proto_enumTypes[0].Descriptor()
}

func (TripStatus) Type() protoreflect.EnumType {
	return &file_trip_proto_enumTypes[0]
}

func (x TripStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TripStatus.Descriptor instead.
func (TripStatus) EnumDescriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{0}
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude float64 `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start         string      `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	StartPos      *Location   `protobuf:"bytes,5,opt,name=start_pos,json=startPos,proto3" json:"start_pos,omitempty"`
	PathLocations []*Location `protobuf:"bytes,7,rep,name=path_locations,json=pathLocations,proto3" json:"path_locations,omitempty"`
	End           string      `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	EndPos        *Location   `protobuf:"bytes,6,opt,name=end_pos,json=endPos,proto3" json:"end_pos,omitempty"`
	DurationSec   uint32      `protobuf:"varint,3,opt,name=duration_sec,json=durationSec,proto3" json:"duration_sec,omitempty"`
	FeeCent       uint32      `protobuf:"varint,4,opt,name=fee_cent,json=feeCent,proto3" json:"fee_cent,omitempty"`
	Status        TripStatus  `protobuf:"varint,8,opt,name=status,proto3,enum=coolcar.TripStatus" json:"status,omitempty"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{1}
}

func (x *Trip) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Trip) GetStartPos() *Location {
	if x != nil {
		return x.StartPos
	}
	return nil
}

func (x *Trip) GetPathLocations() []*Location {
	if x != nil {
		return x.PathLocations
	}
	return nil
}

func (x *Trip) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *Trip) GetEndPos() *Location {
	if x != nil {
		return x.EndPos
	}
	return nil
}

func (x *Trip) GetDurationSec() uint32 {
	if x != nil {
		return x.DurationSec
	}
	return 0
}

func (x *Trip) GetFeeCent() uint32 {
	if x != nil {
		return x.FeeCent
	}
	return 0
}

func (x *Trip) GetStatus() TripStatus {
	if x != nil {
		return x.Status
	}
	return TripStatus_TS_NOT_SPECIFIED
}

type GetTripRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Trip *Trip  `protobuf:"bytes,2,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *GetTripRespone) Reset() {
	*x = GetTripRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripRespone) ProtoMessage() {}

func (x *GetTripRespone) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripRespone.ProtoReflect.Descriptor instead.
func (*GetTripRespone) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{2}
}

func (x *GetTripRespone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTripRespone) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

type GetTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTripRequest) Reset() {
	*x = GetTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripRequest) ProtoMessage() {}

func (x *GetTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripRequest.ProtoReflect.Descriptor instead.
func (*GetTripRequest) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{3}
}

func (x *GetTripRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_trip_proto protoreflect.FileDescriptor

var file_trip_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f,
	0x6f, 0x6c, 0x63, 0x61, 0x72, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xaf, 0x02, 0x0a, 0x04,
	0x54, 0x72, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x38, 0x0a, 0x0e, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x61, 0x74, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61,
	0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x50,
	0x6f, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x65, 0x65, 0x43, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72,
	0x69, 0x70, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x2a, 0x5c, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x49, 0x44,
	0x10, 0x04, 0x32, 0x49, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x12, 0x17, 0x2e, 0x63, 0x6f,
	0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x72, 0x69, 0x70, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trip_proto_rawDescOnce sync.Once
	file_trip_proto_rawDescData = file_trip_proto_rawDesc
)

func file_trip_proto_rawDescGZIP() []byte {
	file_trip_proto_rawDescOnce.Do(func() {
		file_trip_proto_rawDescData = protoimpl.X.CompressGZIP(file_trip_proto_rawDescData)
	})
	return file_trip_proto_rawDescData
}

var file_trip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trip_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_trip_proto_goTypes = []interface{}{
	(TripStatus)(0),        // 0: coolcar.TripStatus
	(*Location)(nil),       // 1: coolcar.Location
	(*Trip)(nil),           // 2: coolcar.Trip
	(*GetTripRespone)(nil), // 3: coolcar.GetTripRespone
	(*GetTripRequest)(nil), // 4: coolcar.GetTripRequest
}
var file_trip_proto_depIdxs = []int32{
	1, // 0: coolcar.Trip.start_pos:type_name -> coolcar.Location
	1, // 1: coolcar.Trip.path_locations:type_name -> coolcar.Location
	1, // 2: coolcar.Trip.end_pos:type_name -> coolcar.Location
	0, // 3: coolcar.Trip.status:type_name -> coolcar.TripStatus
	2, // 4: coolcar.GetTripRespone.trip:type_name -> coolcar.Trip
	4, // 5: coolcar.TripServer.GetTrip:input_type -> coolcar.GetTripRequest
	3, // 6: coolcar.TripServer.GetTrip:output_type -> coolcar.GetTripRespone
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_trip_proto_init() }
func file_trip_proto_init() {
	if File_trip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_trip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trip); i {
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
		file_trip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripRespone); i {
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
		file_trip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripRequest); i {
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
			RawDescriptor: file_trip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trip_proto_goTypes,
		DependencyIndexes: file_trip_proto_depIdxs,
		EnumInfos:         file_trip_proto_enumTypes,
		MessageInfos:      file_trip_proto_msgTypes,
	}.Build()
	File_trip_proto = out.File
	file_trip_proto_rawDesc = nil
	file_trip_proto_goTypes = nil
	file_trip_proto_depIdxs = nil
}