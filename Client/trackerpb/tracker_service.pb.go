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

type PublishQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoHash []byte `protobuf:"bytes,1,opt,name=infoHash,proto3" json:"infoHash,omitempty"`
	PeerID   string `protobuf:"bytes,2,opt,name=peerID,proto3" json:"peerID,omitempty"`
	IP       string `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	Port     int32  `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *PublishQuery) Reset() {
	*x = PublishQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishQuery) ProtoMessage() {}

func (x *PublishQuery) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PublishQuery.ProtoReflect.Descriptor instead.
func (*PublishQuery) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{0}
}

func (x *PublishQuery) GetInfoHash() []byte {
	if x != nil {
		return x.InfoHash
	}
	return nil
}

func (x *PublishQuery) GetPeerID() string {
	if x != nil {
		return x.PeerID
	}
	return ""
}

func (x *PublishQuery) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *PublishQuery) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{1}
}

func (x *PublishResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type AnnounceQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoHash   []byte `protobuf:"bytes,1,opt,name=infoHash,proto3" json:"infoHash,omitempty"`
	PeerID     string `protobuf:"bytes,2,opt,name=peerID,proto3" json:"peerID,omitempty"`
	IP         string `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	Port       int32  `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	Uploaded   uint64 `protobuf:"varint,5,opt,name=Uploaded,proto3" json:"Uploaded,omitempty"`
	Downloaded uint64 `protobuf:"varint,6,opt,name=Downloaded,proto3" json:"Downloaded,omitempty"`
	Left       uint64 `protobuf:"varint,7,opt,name=Left,proto3" json:"Left,omitempty"`
	Event      string `protobuf:"bytes,8,opt,name=Event,proto3" json:"Event,omitempty"`
	NumWant    uint32 `protobuf:"varint,9,opt,name=NumWant,proto3" json:"NumWant,omitempty"`
	Request    bool   `protobuf:"varint,10,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *AnnounceQuery) Reset() {
	*x = AnnounceQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceQuery) ProtoMessage() {}

func (x *AnnounceQuery) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[2]
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
	return file_tracker_service_proto_rawDescGZIP(), []int{2}
}

func (x *AnnounceQuery) GetInfoHash() []byte {
	if x != nil {
		return x.InfoHash
	}
	return nil
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

func (x *AnnounceQuery) GetUploaded() uint64 {
	if x != nil {
		return x.Uploaded
	}
	return 0
}

func (x *AnnounceQuery) GetDownloaded() uint64 {
	if x != nil {
		return x.Downloaded
	}
	return 0
}

func (x *AnnounceQuery) GetLeft() uint64 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *AnnounceQuery) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *AnnounceQuery) GetNumWant() uint32 {
	if x != nil {
		return x.NumWant
	}
	return 0
}

func (x *AnnounceQuery) GetRequest() bool {
	if x != nil {
		return x.Request
	}
	return false
}

type AnnounceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval      uint32            `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
	TrackerID     string            `protobuf:"bytes,2,opt,name=trackerID,proto3" json:"trackerID,omitempty"`
	Complete      int64             `protobuf:"varint,3,opt,name=complete,proto3" json:"complete,omitempty"`
	Incomplete    int64             `protobuf:"varint,4,opt,name=incomplete,proto3" json:"incomplete,omitempty"`
	Peers         map[string]string `protobuf:"bytes,5,rep,name=peers,proto3" json:"peers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FailureReason string            `protobuf:"bytes,6,opt,name=failureReason,proto3" json:"failureReason,omitempty"`
}

func (x *AnnounceResponse) Reset() {
	*x = AnnounceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceResponse) ProtoMessage() {}

func (x *AnnounceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[3]
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
	return file_tracker_service_proto_rawDescGZIP(), []int{3}
}

func (x *AnnounceResponse) GetInterval() uint32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *AnnounceResponse) GetTrackerID() string {
	if x != nil {
		return x.TrackerID
	}
	return ""
}

func (x *AnnounceResponse) GetComplete() int64 {
	if x != nil {
		return x.Complete
	}
	return 0
}

func (x *AnnounceResponse) GetIncomplete() int64 {
	if x != nil {
		return x.Incomplete
	}
	return 0
}

func (x *AnnounceResponse) GetPeers() map[string]string {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *AnnounceResponse) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

type ScraperQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoHash [][]byte `protobuf:"bytes,1,rep,name=infoHash,proto3" json:"infoHash,omitempty"`
}

func (x *ScraperQuery) Reset() {
	*x = ScraperQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScraperQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScraperQuery) ProtoMessage() {}

func (x *ScraperQuery) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScraperQuery.ProtoReflect.Descriptor instead.
func (*ScraperQuery) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{4}
}

func (x *ScraperQuery) GetInfoHash() [][]byte {
	if x != nil {
		return x.InfoHash
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incomplete int64 `protobuf:"varint,1,opt,name=incomplete,proto3" json:"incomplete,omitempty"`
	Complete   int64 `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
	Downloaded int64 `protobuf:"varint,3,opt,name=downloaded,proto3" json:"downloaded,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{5}
}

func (x *File) GetIncomplete() int64 {
	if x != nil {
		return x.Incomplete
	}
	return 0
}

func (x *File) GetComplete() int64 {
	if x != nil {
		return x.Complete
	}
	return 0
}

func (x *File) GetDownloaded() int64 {
	if x != nil {
		return x.Downloaded
	}
	return 0
}

//ScraperResponse representa el contenido de una respuesta a peticion Scraper
type ScraperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files         map[string]*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FailureReason string           `protobuf:"bytes,2,opt,name=failureReason,proto3" json:"failureReason,omitempty"`
}

func (x *ScraperResponse) Reset() {
	*x = ScraperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScraperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScraperResponse) ProtoMessage() {}

func (x *ScraperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScraperResponse.ProtoReflect.Descriptor instead.
func (*ScraperResponse) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{6}
}

func (x *ScraperResponse) GetFiles() map[string]*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *ScraperResponse) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

type RePublishQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedKey   []byte `protobuf:"bytes,1,opt,name=redKey,proto3" json:"redKey,omitempty"`
	InfoHash []byte `protobuf:"bytes,2,opt,name=infoHash,proto3" json:"infoHash,omitempty"`
	PeerID   string `protobuf:"bytes,3,opt,name=peerID,proto3" json:"peerID,omitempty"`
	IP       string `protobuf:"bytes,4,opt,name=IP,proto3" json:"IP,omitempty"`
	Port     int32  `protobuf:"varint,5,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *RePublishQuery) Reset() {
	*x = RePublishQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RePublishQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RePublishQuery) ProtoMessage() {}

func (x *RePublishQuery) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RePublishQuery.ProtoReflect.Descriptor instead.
func (*RePublishQuery) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{7}
}

func (x *RePublishQuery) GetRedKey() []byte {
	if x != nil {
		return x.RedKey
	}
	return nil
}

func (x *RePublishQuery) GetInfoHash() []byte {
	if x != nil {
		return x.InfoHash
	}
	return nil
}

func (x *RePublishQuery) GetPeerID() string {
	if x != nil {
		return x.PeerID
	}
	return ""
}

func (x *RePublishQuery) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *RePublishQuery) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type RePublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RePublishResponse) Reset() {
	*x = RePublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RePublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RePublishResponse) ProtoMessage() {}

func (x *RePublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RePublishResponse.ProtoReflect.Descriptor instead.
func (*RePublishResponse) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{8}
}

func (x *RePublishResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type KnowMeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedKey []byte `protobuf:"bytes,1,opt,name=redKey,proto3" json:"redKey,omitempty"`
	IP     string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port   int32  `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *KnowMeRequest) Reset() {
	*x = KnowMeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowMeRequest) ProtoMessage() {}

func (x *KnowMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowMeRequest.ProtoReflect.Descriptor instead.
func (*KnowMeRequest) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{9}
}

func (x *KnowMeRequest) GetRedKey() []byte {
	if x != nil {
		return x.RedKey
	}
	return nil
}

func (x *KnowMeRequest) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *KnowMeRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type KnowMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	RepeatInterval int32 `protobuf:"varint,2,opt,name=repeatInterval,proto3" json:"repeatInterval,omitempty"`
}

func (x *KnowMeResponse) Reset() {
	*x = KnowMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowMeResponse) ProtoMessage() {}

func (x *KnowMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowMeResponse.ProtoReflect.Descriptor instead.
func (*KnowMeResponse) Descriptor() ([]byte, []int) {
	return file_tracker_service_proto_rawDescGZIP(), []int{10}
}

func (x *KnowMeResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *KnowMeResponse) GetRepeatInterval() int32 {
	if x != nil {
		return x.RepeatInterval
	}
	return 0
}

var File_tracker_service_proto protoreflect.FileDescriptor

var file_tracker_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x22, 0x66, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x81, 0x02, 0x0a, 0x0d, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x65, 0x66,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4e, 0x75, 0x6d, 0x57, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa4, 0x02, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a,
	0x0a, 0x0c, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x22, 0x62, 0x0a, 0x04, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x22, 0xbb,
	0x01, 0x0a, 0x0f, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x1a, 0x47, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x22,
	0x2b, 0x0a, 0x11, 0x52, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x0d,
	0x4b, 0x6e, 0x6f, 0x77, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x50, 0x0a, 0x0e, 0x4b, 0x6e, 0x6f,
	0x77, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x32, 0xc6, 0x02, 0x0a, 0x07,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x12, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x52, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x12, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x4b, 0x6e, 0x6f, 0x77, 0x4d,
	0x65, 0x12, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x4b, 0x6e, 0x6f, 0x77,
	0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_tracker_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tracker_service_proto_goTypes = []interface{}{
	(*PublishQuery)(nil),      // 0: tracker.PublishQuery
	(*PublishResponse)(nil),   // 1: tracker.PublishResponse
	(*AnnounceQuery)(nil),     // 2: tracker.AnnounceQuery
	(*AnnounceResponse)(nil),  // 3: tracker.AnnounceResponse
	(*ScraperQuery)(nil),      // 4: tracker.ScraperQuery
	(*File)(nil),              // 5: tracker.File
	(*ScraperResponse)(nil),   // 6: tracker.ScraperResponse
	(*RePublishQuery)(nil),    // 7: tracker.RePublishQuery
	(*RePublishResponse)(nil), // 8: tracker.RePublishResponse
	(*KnowMeRequest)(nil),     // 9: tracker.KnowMeRequest
	(*KnowMeResponse)(nil),    // 10: tracker.KnowMeResponse
	nil,                       // 11: tracker.AnnounceResponse.PeersEntry
	nil,                       // 12: tracker.ScraperResponse.FilesEntry
}
var file_tracker_service_proto_depIdxs = []int32{
	11, // 0: tracker.AnnounceResponse.peers:type_name -> tracker.AnnounceResponse.PeersEntry
	12, // 1: tracker.ScraperResponse.files:type_name -> tracker.ScraperResponse.FilesEntry
	5,  // 2: tracker.ScraperResponse.FilesEntry.value:type_name -> tracker.File
	0,  // 3: tracker.Tracker.Publish:input_type -> tracker.PublishQuery
	2,  // 4: tracker.Tracker.Announce:input_type -> tracker.AnnounceQuery
	4,  // 5: tracker.Tracker.Scrape:input_type -> tracker.ScraperQuery
	7,  // 6: tracker.Tracker.RePublish:input_type -> tracker.RePublishQuery
	9,  // 7: tracker.Tracker.KnowMe:input_type -> tracker.KnowMeRequest
	1,  // 8: tracker.Tracker.Publish:output_type -> tracker.PublishResponse
	3,  // 9: tracker.Tracker.Announce:output_type -> tracker.AnnounceResponse
	6,  // 10: tracker.Tracker.Scrape:output_type -> tracker.ScraperResponse
	8,  // 11: tracker.Tracker.RePublish:output_type -> tracker.RePublishResponse
	10, // 12: tracker.Tracker.KnowMe:output_type -> tracker.KnowMeResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_tracker_service_proto_init() }
func file_tracker_service_proto_init() {
	if File_tracker_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracker_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishQuery); i {
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
			switch v := v.(*PublishResponse); i {
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
		file_tracker_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_tracker_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_tracker_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScraperQuery); i {
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
		file_tracker_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_tracker_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScraperResponse); i {
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
		file_tracker_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RePublishQuery); i {
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
		file_tracker_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RePublishResponse); i {
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
		file_tracker_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowMeRequest); i {
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
		file_tracker_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowMeResponse); i {
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
			NumMessages:   13,
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
