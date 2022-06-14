// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tracker_service.proto

package trackerpb

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

type AnnounceQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoHash   string `protobuf:"bytes,1,opt,name=infoHash,proto3" json:"infoHash,omitempty"`
	PeerID     string `protobuf:"bytes,2,opt,name=peerID,proto3" json:"peerID,omitempty"`
	IP         string `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	Port       int32  `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	Uploaded   int32  `protobuf:"varint,5,opt,name=Uploaded,proto3" json:"Uploaded,omitempty"`
	Downloaded int32  `protobuf:"varint,6,opt,name=Downloaded,proto3" json:"Downloaded,omitempty"`
	Left       int32  `protobuf:"varint,7,opt,name=Left,proto3" json:"Left,omitempty"`
	Event      int32  `protobuf:"varint,8,opt,name=Event,proto3" json:"Event,omitempty"`
	NumWant    int32  `protobuf:"varint,9,opt,name=NumWant,proto3" json:"NumWant,omitempty"`
}

func (x *AnnounceQuery) Reset() {
	*x = AnnounceQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceQuery) ProtoMessage() {}

func (x *AnnounceQuery) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceQuery.ProtoReflect.Descriptor instead.
func (*AnnounceQuery) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{0}
}

func (x *AnnounceQuery) GetInfoHash() string {
	if x != nil {
		return x.InfoHash
	}
	return ""
}

func (x *AnnounceQuery) GetPeerID() string {
	if x != nil {
		return x.PeerID
	}
	return ""
}

func (x *AnnounceQuery) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *AnnounceQuery) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AnnounceQuery) GetUploaded() int32 {
	if x != nil {
		return x.Uploaded
	}
	return 0
}

func (x *AnnounceQuery) GetDownloaded() int32 {
	if x != nil {
		return x.Downloaded
	}
	return 0
}

func (x *AnnounceQuery) GetLeft() int32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *AnnounceQuery) GetEvent() int32 {
	if x != nil {
		return x.Event
	}
	return 0
}

func (x *AnnounceQuery) GetNumWant() int32 {
	if x != nil {
		return x.NumWant
	}
	return 0
}

type AnnounceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval   int32  `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
	TrackerID  int32  `protobuf:"varint,2,opt,name=trackerID,proto3" json:"trackerID,omitempty"`
	Complete   int32  `protobuf:"varint,3,opt,name=complete,proto3" json:"complete,omitempty"`
	Incomplete int32  `protobuf:"varint,4,opt,name=incomplete,proto3" json:"incomplete,omitempty"`
	Peers      string `protobuf:"bytes,5,opt,name=peers,proto3" json:"peers,omitempty"`
}

func (x *AnnounceResponse) Reset() {
	*x = AnnounceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceResponse) ProtoMessage() {}

func (x *AnnounceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceResponse.ProtoReflect.Descriptor instead.
func (*AnnounceResponse) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{1}
}

func (x *AnnounceResponse) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *AnnounceResponse) GetTrackerID() int32 {
	if x != nil {
		return x.TrackerID
	}
	return 0
}

func (x *AnnounceResponse) GetComplete() int32 {
	if x != nil {
		return x.Complete
	}
	return 0
}

func (x *AnnounceResponse) GetIncomplete() int32 {
	if x != nil {
		return x.Incomplete
	}
	return 0
}

func (x *AnnounceResponse) GetPeers() string {
	if x != nil {
		return x.Peers
	}
	return ""
}

var File_tracker_service_proto protoreflect.FileDescriptor

var file_tracker_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x22, 0xe7, 0x01, 0x0a, 0x0d, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x4e, 0x75, 0x6d, 0x57, 0x61, 0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x10, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x32, 0x4a, 0x0a, 0x07, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x12, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracker_service_proto_rawDescOnce sync.Once
	file_tracker_service_proto_rawDescData = file_tracker_service_proto_rawDesc
)

func file_tracker_service_proto_rawDescGZIP() []byte {
	file_tracker_service_proto_rawDescOnce.Do(func() {
		file_tracker_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracker_service_proto_rawDescData)
	})
	return file_tracker_service_proto_rawDescData
}

var file_tracker_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tracker_service_proto_goTypes = []interface{}{
	(*AnnounceQuery)(nil),    // 0: tracker.AnnounceQuery
	(*AnnounceResponse)(nil), // 1: tracker.AnnounceResponse
}
var file_tracker_service_proto_depIdxs = []int32{
	0, // 0: tracker.Tracker.Announce:input_type -> tracker.AnnounceQuery
	1, // 1: tracker.Tracker.Announce:output_type -> tracker.AnnounceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tracker_service_proto_init() }
func file_tracker_service_proto_init() {
	if File_tracker_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracker_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceQuery); i {
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
		file_tracker_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceResponse); i {
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
			RawDescriptor: file_tracker_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracker_service_proto_goTypes,
		DependencyIndexes: file_tracker_service_proto_depIdxs,
		MessageInfos:      file_tracker_service_proto_msgTypes,
	}.Build()
	File_tracker_service_proto = out.File
	file_tracker_service_proto_rawDesc = nil
	file_tracker_service_proto_goTypes = nil
	file_tracker_service_proto_depIdxs = nil
}