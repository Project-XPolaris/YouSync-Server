// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: service.proto

package pb

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

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ChunkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size     uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Offset   uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	CheckSum string `protobuf:"bytes,4,opt,name=checkSum,proto3" json:"checkSum,omitempty"`
	Path     string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	FolderId uint64 `protobuf:"varint,6,opt,name=folderId,proto3" json:"folderId,omitempty"`
}

func (x *ChunkInfo) Reset() {
	*x = ChunkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkInfo) ProtoMessage() {}

func (x *ChunkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkInfo.ProtoReflect.Descriptor instead.
func (*ChunkInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkInfo) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ChunkInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ChunkInfo) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ChunkInfo) GetCheckSum() string {
	if x != nil {
		return x.CheckSum
	}
	return ""
}

func (x *ChunkInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ChunkInfo) GetFolderId() uint64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

type CheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckResult) Reset() {
	*x = CheckResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResult) ProtoMessage() {}

func (x *CheckResult) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResult.ProtoReflect.Descriptor instead.
func (*CheckResult) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size      uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Offset    uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	FolderId  uint64 `protobuf:"varint,5,opt,name=folderId,proto3" json:"folderId,omitempty"`
	LastChunk bool   `protobuf:"varint,6,opt,name=LastChunk,proto3" json:"LastChunk,omitempty"`
	Data      []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *Chunk) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Chunk) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Chunk) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Chunk) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Chunk) GetFolderId() uint64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *Chunk) GetLastChunk() bool {
	if x != nil {
		return x.LastChunk
	}
	return false
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SyncChunkResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SyncChunkResult) Reset() {
	*x = SyncChunkResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncChunkResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncChunkResult) ProtoMessage() {}

func (x *SyncChunkResult) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncChunkResult.ProtoReflect.Descriptor instead.
func (*SyncChunkResult) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *SyncChunkResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RemoteFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	FolderId uint64 `protobuf:"varint,2,opt,name=folderId,proto3" json:"folderId,omitempty"`
	Size     uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *RemoteFiles) Reset() {
	*x = RemoteFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteFiles) ProtoMessage() {}

func (x *RemoteFiles) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteFiles.ProtoReflect.Descriptor instead.
func (*RemoteFiles) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *RemoteFiles) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RemoteFiles) GetFolderId() uint64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *RemoteFiles) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type RemoteFilesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*RemoteFiles `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *RemoteFilesResult) Reset() {
	*x = RemoteFilesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteFilesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteFilesResult) ProtoMessage() {}

func (x *RemoteFilesResult) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteFilesResult.ProtoReflect.Descriptor instead.
func (*RemoteFilesResult) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *RemoteFilesResult) GetFiles() []*RemoteFiles {
	if x != nil {
		return x.Files
	}
	return nil
}

type RemoteFilesMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FolderId uint64 `protobuf:"varint,1,opt,name=folderId,proto3" json:"folderId,omitempty"`
}

func (x *RemoteFilesMessage) Reset() {
	*x = RemoteFilesMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteFilesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteFilesMessage) ProtoMessage() {}

func (x *RemoteFilesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteFilesMessage.ProtoReflect.Descriptor instead.
func (*RemoteFilesMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *RemoteFilesMessage) GetFolderId() uint64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

type GetRemoteChunkInfoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	FolderId uint64 `protobuf:"varint,2,opt,name=folderId,proto3" json:"folderId,omitempty"`
	Size     uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Offset   uint64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetRemoteChunkInfoMessage) Reset() {
	*x = GetRemoteChunkInfoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteChunkInfoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteChunkInfoMessage) ProtoMessage() {}

func (x *GetRemoteChunkInfoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteChunkInfoMessage.ProtoReflect.Descriptor instead.
func (*GetRemoteChunkInfoMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetRemoteChunkInfoMessage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetRemoteChunkInfoMessage) GetFolderId() uint64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *GetRemoteChunkInfoMessage) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetRemoteChunkInfoMessage) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type RemoteChunkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastChunk bool   `protobuf:"varint,1,opt,name=LastChunk,proto3" json:"LastChunk,omitempty"`
	Checksum  string `protobuf:"bytes,2,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
}

func (x *RemoteChunkInfo) Reset() {
	*x = RemoteChunkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteChunkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteChunkInfo) ProtoMessage() {}

func (x *RemoteChunkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteChunkInfo.ProtoReflect.Descriptor instead.
func (*RemoteChunkInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9}
}

func (x *RemoteChunkInfo) GetLastChunk() bool {
	if x != nil {
		return x.LastChunk
	}
	return false
}

func (x *RemoteChunkInfo) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

type GetRemoteChunkMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size     uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Offset   uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	FolderId uint64 `protobuf:"varint,4,opt,name=folderId,proto3" json:"folderId,omitempty"`
}

func (x *GetRemoteChunkMessage) Reset() {
	*x = GetRemoteChunkMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteChunkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteChunkMessage) ProtoMessage() {}

func (x *GetRemoteChunkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteChunkMessage.ProtoReflect.Descriptor instead.
func (*GetRemoteChunkMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetRemoteChunkMessage) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetRemoteChunkMessage) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetRemoteChunkMessage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetRemoteChunkMessage) GetFolderId() uint64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

type RemoteChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RemoteChunk) Reset() {
	*x = RemoteChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteChunk) ProtoMessage() {}

func (x *RemoteChunk) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteChunk.ProtoReflect.Descriptor instead.
func (*RemoteChunk) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{11}
}

func (x *RemoteChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SyncFileItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *SyncFileItem) Reset() {
	*x = SyncFileItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncFileItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncFileItem) ProtoMessage() {}

func (x *SyncFileItem) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncFileItem.ProtoReflect.Descriptor instead.
func (*SyncFileItem) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{12}
}

func (x *SyncFileItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SyncFileListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folder   []*SyncFileItem `protobuf:"bytes,1,rep,name=folder,proto3" json:"folder,omitempty"`
	File     []*SyncFileItem `protobuf:"bytes,2,rep,name=file,proto3" json:"file,omitempty"`
	FolderId uint64          `protobuf:"varint,3,opt,name=folderId,proto3" json:"folderId,omitempty"`
}

func (x *SyncFileListMessage) Reset() {
	*x = SyncFileListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncFileListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncFileListMessage) ProtoMessage() {}

func (x *SyncFileListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncFileListMessage.ProtoReflect.Descriptor instead.
func (*SyncFileListMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{13}
}

func (x *SyncFileListMessage) GetFolder() []*SyncFileItem {
	if x != nil {
		return x.Folder
	}
	return nil
}

func (x *SyncFileListMessage) GetFile() []*SyncFileItem {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *SyncFileListMessage) GetFolderId() uint64 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x09, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x53, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x53, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xab,
	0x01, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x61,
	0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4c,
	0x61, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x0f,
	0x53, 0x79, 0x6e, 0x63, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x51, 0x0a, 0x0b, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x37, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x4b, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22, 0x73, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x7b, 0x0a, 0x13, 0x53, 0x79, 0x6e,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0xde, 0x02, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x28, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x0a, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2b, 0x0a,
	0x0d, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x06,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x10, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0f, 0x52, 0x65,
	0x61, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x13, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_service_proto_goTypes = []interface{}{
	(*BaseResponse)(nil),              // 0: BaseResponse
	(*ChunkInfo)(nil),                 // 1: ChunkInfo
	(*CheckResult)(nil),               // 2: CheckResult
	(*Chunk)(nil),                     // 3: Chunk
	(*SyncChunkResult)(nil),           // 4: SyncChunkResult
	(*RemoteFiles)(nil),               // 5: RemoteFiles
	(*RemoteFilesResult)(nil),         // 6: RemoteFilesResult
	(*RemoteFilesMessage)(nil),        // 7: RemoteFilesMessage
	(*GetRemoteChunkInfoMessage)(nil), // 8: GetRemoteChunkInfoMessage
	(*RemoteChunkInfo)(nil),           // 9: RemoteChunkInfo
	(*GetRemoteChunkMessage)(nil),     // 10: GetRemoteChunkMessage
	(*RemoteChunk)(nil),               // 11: RemoteChunk
	(*SyncFileItem)(nil),              // 12: SyncFileItem
	(*SyncFileListMessage)(nil),       // 13: SyncFileListMessage
}
var file_service_proto_depIdxs = []int32{
	5,  // 0: RemoteFilesResult.files:type_name -> RemoteFiles
	12, // 1: SyncFileListMessage.folder:type_name -> SyncFileItem
	12, // 2: SyncFileListMessage.file:type_name -> SyncFileItem
	1,  // 3: FileSync.CheckChunk:input_type -> ChunkInfo
	3,  // 4: FileSync.SyncFileChunk:input_type -> Chunk
	7,  // 5: FileSync.ReadFolderFiles:input_type -> RemoteFilesMessage
	8,  // 6: FileSync.GetRemoteFileChunkInfo:input_type -> GetRemoteChunkInfoMessage
	10, // 7: FileSync.GetRemoteFileChunk:input_type -> GetRemoteChunkMessage
	13, // 8: FileSync.SyncFileList:input_type -> SyncFileListMessage
	2,  // 9: FileSync.CheckChunk:output_type -> CheckResult
	4,  // 10: FileSync.SyncFileChunk:output_type -> SyncChunkResult
	6,  // 11: FileSync.ReadFolderFiles:output_type -> RemoteFilesResult
	9,  // 12: FileSync.GetRemoteFileChunkInfo:output_type -> RemoteChunkInfo
	11, // 13: FileSync.GetRemoteFileChunk:output_type -> RemoteChunk
	0,  // 14: FileSync.SyncFileList:output_type -> BaseResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkInfo); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResult); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncChunkResult); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteFiles); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteFilesResult); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteFilesMessage); i {
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
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteChunkInfoMessage); i {
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
		file_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteChunkInfo); i {
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
		file_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteChunkMessage); i {
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
		file_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteChunk); i {
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
		file_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncFileItem); i {
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
		file_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncFileListMessage); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
