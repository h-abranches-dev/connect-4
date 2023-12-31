// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: game-engine/game_engine.proto

package gameengine

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

type VerifyCompatibilityPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameServerVersion string `protobuf:"bytes,1,opt,name=gameServerVersion,proto3" json:"gameServerVersion,omitempty"`
}

func (x *VerifyCompatibilityPayload) Reset() {
	*x = VerifyCompatibilityPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCompatibilityPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCompatibilityPayload) ProtoMessage() {}

func (x *VerifyCompatibilityPayload) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCompatibilityPayload.ProtoReflect.Descriptor instead.
func (*VerifyCompatibilityPayload) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyCompatibilityPayload) GetGameServerVersion() string {
	if x != nil {
		return x.GameServerVersion
	}
	return ""
}

type VerifyCompatibilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyCompatibilityResponse) Reset() {
	*x = VerifyCompatibilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCompatibilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCompatibilityResponse) ProtoMessage() {}

func (x *VerifyCompatibilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCompatibilityResponse.ProtoReflect.Descriptor instead.
func (*VerifyCompatibilityResponse) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{1}
}

type ConnectPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectPayload) Reset() {
	*x = ConnectPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectPayload) ProtoMessage() {}

func (x *ConnectPayload) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectPayload.ProtoReflect.Descriptor instead.
func (*ConnectPayload) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{2}
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionToken string `protobuf:"bytes,1,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectResponse) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type POLPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionToken  string `protobuf:"bytes,1,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
	EncodedBoards string `protobuf:"bytes,2,opt,name=encodedBoards,proto3" json:"encodedBoards,omitempty"`
}

func (x *POLPayload) Reset() {
	*x = POLPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *POLPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*POLPayload) ProtoMessage() {}

func (x *POLPayload) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use POLPayload.ProtoReflect.Descriptor instead.
func (*POLPayload) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{4}
}

func (x *POLPayload) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

func (x *POLPayload) GetEncodedBoards() string {
	if x != nil {
		return x.EncodedBoards
	}
	return ""
}

type POLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *POLResponse) Reset() {
	*x = POLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *POLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*POLResponse) ProtoMessage() {}

func (x *POLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use POLResponse.ProtoReflect.Descriptor instead.
func (*POLResponse) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{5}
}

type ServeBoardPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardID string `protobuf:"bytes,1,opt,name=boardID,proto3" json:"boardID,omitempty"`
}

func (x *ServeBoardPayload) Reset() {
	*x = ServeBoardPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServeBoardPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServeBoardPayload) ProtoMessage() {}

func (x *ServeBoardPayload) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServeBoardPayload.ProtoReflect.Descriptor instead.
func (*ServeBoardPayload) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{6}
}

func (x *ServeBoardPayload) GetBoardID() string {
	if x != nil {
		return x.BoardID
	}
	return ""
}

type ServeBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board string `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *ServeBoardResponse) Reset() {
	*x = ServeBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServeBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServeBoardResponse) ProtoMessage() {}

func (x *ServeBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServeBoardResponse.ProtoReflect.Descriptor instead.
func (*ServeBoardResponse) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{7}
}

func (x *ServeBoardResponse) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

type PlayPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardID    string `protobuf:"bytes,1,opt,name=boardID,proto3" json:"boardID,omitempty"`
	PlayerCode string `protobuf:"bytes,2,opt,name=playerCode,proto3" json:"playerCode,omitempty"`
	Column     int32  `protobuf:"varint,3,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *PlayPayload) Reset() {
	*x = PlayPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayPayload) ProtoMessage() {}

func (x *PlayPayload) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayPayload.ProtoReflect.Descriptor instead.
func (*PlayPayload) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{8}
}

func (x *PlayPayload) GetBoardID() string {
	if x != nil {
		return x.BoardID
	}
	return ""
}

func (x *PlayPayload) GetPlayerCode() string {
	if x != nil {
		return x.PlayerCode
	}
	return ""
}

func (x *PlayPayload) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

type PlayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThereIsWinner bool `protobuf:"varint,1,opt,name=thereIsWinner,proto3" json:"thereIsWinner,omitempty"`
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_game_engine_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_game_engine_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_game_engine_game_engine_proto_rawDescGZIP(), []int{9}
}

func (x *PlayResponse) GetThereIsWinner() bool {
	if x != nil {
		return x.ThereIsWinner
	}
	return false
}

var File_game_engine_game_engine_proto protoreflect.FileDescriptor

var file_game_engine_game_engine_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x22, 0x4a, 0x0a, 0x1a, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x67, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x56, 0x0a, 0x0a, 0x50, 0x4f, 0x4c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x4f, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x22, 0x5f, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x22, 0x34, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x68, 0x65, 0x72, 0x65, 0x49, 0x73, 0x57, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x49, 0x73, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x32, 0xfd, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x68, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x26, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x03, 0x50, 0x4f, 0x4c, 0x12, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x4f, 0x4c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50,
	0x4f, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x50,
	0x6c, 0x61, 0x79, 0x12, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x18, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x67, 0x61,
	0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_engine_game_engine_proto_rawDescOnce sync.Once
	file_game_engine_game_engine_proto_rawDescData = file_game_engine_game_engine_proto_rawDesc
)

func file_game_engine_game_engine_proto_rawDescGZIP() []byte {
	file_game_engine_game_engine_proto_rawDescOnce.Do(func() {
		file_game_engine_game_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_engine_game_engine_proto_rawDescData)
	})
	return file_game_engine_game_engine_proto_rawDescData
}

var file_game_engine_game_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_game_engine_game_engine_proto_goTypes = []interface{}{
	(*VerifyCompatibilityPayload)(nil),  // 0: gameengine.VerifyCompatibilityPayload
	(*VerifyCompatibilityResponse)(nil), // 1: gameengine.VerifyCompatibilityResponse
	(*ConnectPayload)(nil),              // 2: gameengine.ConnectPayload
	(*ConnectResponse)(nil),             // 3: gameengine.ConnectResponse
	(*POLPayload)(nil),                  // 4: gameengine.POLPayload
	(*POLResponse)(nil),                 // 5: gameengine.POLResponse
	(*ServeBoardPayload)(nil),           // 6: gameengine.ServeBoardPayload
	(*ServeBoardResponse)(nil),          // 7: gameengine.ServeBoardResponse
	(*PlayPayload)(nil),                 // 8: gameengine.PlayPayload
	(*PlayResponse)(nil),                // 9: gameengine.PlayResponse
}
var file_game_engine_game_engine_proto_depIdxs = []int32{
	0, // 0: gameengine.Route.VerifyCompatibility:input_type -> gameengine.VerifyCompatibilityPayload
	2, // 1: gameengine.Route.Connect:input_type -> gameengine.ConnectPayload
	4, // 2: gameengine.Route.POL:input_type -> gameengine.POLPayload
	6, // 3: gameengine.Route.ServeBoard:input_type -> gameengine.ServeBoardPayload
	8, // 4: gameengine.Route.Play:input_type -> gameengine.PlayPayload
	1, // 5: gameengine.Route.VerifyCompatibility:output_type -> gameengine.VerifyCompatibilityResponse
	3, // 6: gameengine.Route.Connect:output_type -> gameengine.ConnectResponse
	5, // 7: gameengine.Route.POL:output_type -> gameengine.POLResponse
	7, // 8: gameengine.Route.ServeBoard:output_type -> gameengine.ServeBoardResponse
	9, // 9: gameengine.Route.Play:output_type -> gameengine.PlayResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_engine_game_engine_proto_init() }
func file_game_engine_game_engine_proto_init() {
	if File_game_engine_game_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_engine_game_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCompatibilityPayload); i {
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
		file_game_engine_game_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCompatibilityResponse); i {
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
		file_game_engine_game_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectPayload); i {
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
		file_game_engine_game_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_game_engine_game_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*POLPayload); i {
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
		file_game_engine_game_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*POLResponse); i {
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
		file_game_engine_game_engine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServeBoardPayload); i {
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
		file_game_engine_game_engine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServeBoardResponse); i {
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
		file_game_engine_game_engine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayPayload); i {
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
		file_game_engine_game_engine_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse); i {
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
			RawDescriptor: file_game_engine_game_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_engine_game_engine_proto_goTypes,
		DependencyIndexes: file_game_engine_game_engine_proto_depIdxs,
		MessageInfos:      file_game_engine_game_engine_proto_msgTypes,
	}.Build()
	File_game_engine_game_engine_proto = out.File
	file_game_engine_game_engine_proto_rawDesc = nil
	file_game_engine_game_engine_proto_goTypes = nil
	file_game_engine_game_engine_proto_depIdxs = nil
}
