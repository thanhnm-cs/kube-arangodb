//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: pkg/api/ml/storage/v1/bucketservice.proto

package v1

import (
	v1 "github.com/arangodb/kube-arangodb/pkg/api/common/v1"
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

// Request params for Bucket related calls
type BucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key-value pairs stored at the providers (as labels/tags) during CreateBucket
	Tags []*v1.KeyValuePair `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *BucketRequest) Reset() {
	*x = BucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketRequest) ProtoMessage() {}

func (x *BucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketRequest.ProtoReflect.Descriptor instead.
func (*BucketRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{0}
}

func (x *BucketRequest) GetTags() []*v1.KeyValuePair {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Request params for Path related requests
type PathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path (Specify as "." to indicate the root folder)
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PathRequest) Reset() {
	*x = PathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathRequest) ProtoMessage() {}

func (x *PathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathRequest.ProtoReflect.Descriptor instead.
func (*PathRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{1}
}

func (x *PathRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Response from GetRepositoryURL request
type RepositoryURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URL
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The URL without provider information
	BucketPath string `protobuf:"bytes,2,opt,name=bucket_path,json=bucketPath,proto3" json:"bucket_path,omitempty"`
}

func (x *RepositoryURL) Reset() {
	*x = RepositoryURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryURL) ProtoMessage() {}

func (x *RepositoryURL) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryURL.ProtoReflect.Descriptor instead.
func (*RepositoryURL) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{2}
}

func (x *RepositoryURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RepositoryURL) GetBucketPath() string {
	if x != nil {
		return x.BucketPath
	}
	return ""
}

// Response from GetPathSize request
type PathSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The size in bytes
	SizeInBytes uint64 `protobuf:"varint,1,opt,name=size_in_bytes,json=sizeInBytes,proto3" json:"size_in_bytes,omitempty"`
	// Number of files
	NumberOfFiles uint32 `protobuf:"varint,2,opt,name=number_of_files,json=numberOfFiles,proto3" json:"number_of_files,omitempty"`
	// Number of folders
	NumberOfFolders uint32 `protobuf:"varint,3,opt,name=number_of_folders,json=numberOfFolders,proto3" json:"number_of_folders,omitempty"`
}

func (x *PathSize) Reset() {
	*x = PathSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathSize) ProtoMessage() {}

func (x *PathSize) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathSize.ProtoReflect.Descriptor instead.
func (*PathSize) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{3}
}

func (x *PathSize) GetSizeInBytes() uint64 {
	if x != nil {
		return x.SizeInBytes
	}
	return 0
}

func (x *PathSize) GetNumberOfFiles() uint32 {
	if x != nil {
		return x.NumberOfFiles
	}
	return 0
}

func (x *PathSize) GetNumberOfFolders() uint32 {
	if x != nil {
		return x.NumberOfFolders
	}
	return 0
}

// Response from GetObjectInfo request
type ObjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates if the object is locked
	IsLocked bool `protobuf:"varint,1,opt,name=is_locked,json=isLocked,proto3" json:"is_locked,omitempty"`
	// Indicates the size of the object in bytes
	SizeInBytes uint64 `protobuf:"varint,2,opt,name=size_in_bytes,json=sizeInBytes,proto3" json:"size_in_bytes,omitempty"`
	// The timestamp this object has last been modified
	LastUpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_updated_at,json=lastUpdatedAt,proto3" json:"last_updated_at,omitempty"`
}

func (x *ObjectInfo) Reset() {
	*x = ObjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectInfo) ProtoMessage() {}

func (x *ObjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectInfo.ProtoReflect.Descriptor instead.
func (*ObjectInfo) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectInfo) GetIsLocked() bool {
	if x != nil {
		return x.IsLocked
	}
	return false
}

func (x *ObjectInfo) GetSizeInBytes() uint64 {
	if x != nil {
		return x.SizeInBytes
	}
	return 0
}

func (x *ObjectInfo) GetLastUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedAt
	}
	return nil
}

// Output message for ReadObject.
type ReadObjectChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Raw output
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *ReadObjectChunk) Reset() {
	*x = ReadObjectChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadObjectChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadObjectChunk) ProtoMessage() {}

func (x *ReadObjectChunk) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadObjectChunk.ProtoReflect.Descriptor instead.
func (*ReadObjectChunk) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{5}
}

func (x *ReadObjectChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// Input message for WriteObject
type WriteObjectChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base request params for Path related requests.
	// This field cannot change during the stream.
	Path *PathRequest `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Raw input
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	// If set, the caller wants to send a next message with more input data.
	// If not set, no more control message will be sent.
	HasMore bool `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *WriteObjectChunk) Reset() {
	*x = WriteObjectChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteObjectChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteObjectChunk) ProtoMessage() {}

func (x *WriteObjectChunk) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteObjectChunk.ProtoReflect.Descriptor instead.
func (*WriteObjectChunk) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{6}
}

func (x *WriteObjectChunk) GetPath() *PathRequest {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *WriteObjectChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *WriteObjectChunk) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// Output message for WriteObject
type WriteObjectControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, the next chunk of input is allowed.
	// If unset, the caller is expected to terminate the call.
	AllowMoreOutput bool `protobuf:"varint,1,opt,name=allow_more_output,json=allowMoreOutput,proto3" json:"allow_more_output,omitempty"`
	// Maximum number of bytes the caller is allowed to send in the next input chunk.
	MaxChunkBytes int32 `protobuf:"varint,2,opt,name=max_chunk_bytes,json=maxChunkBytes,proto3" json:"max_chunk_bytes,omitempty"`
}

func (x *WriteObjectControl) Reset() {
	*x = WriteObjectControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteObjectControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteObjectControl) ProtoMessage() {}

func (x *WriteObjectControl) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteObjectControl.ProtoReflect.Descriptor instead.
func (*WriteObjectControl) Descriptor() ([]byte, []int) {
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP(), []int{7}
}

func (x *WriteObjectControl) GetAllowMoreOutput() bool {
	if x != nil {
		return x.AllowMoreOutput
	}
	return false
}

func (x *WriteObjectControl) GetMaxChunkBytes() int32 {
	if x != nil {
		return x.MaxChunkBytes
	}
	return 0
}

var File_pkg_api_ml_storage_v1_bucketservice_proto protoreflect.FileDescriptor

var file_pkg_api_ml_storage_v1_bucketservice_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6c, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x72, 0x61,
	0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d,
	0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a,
	0x0d, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61,
	0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x21, 0x0a,
	0x0b, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x42, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x52,
	0x4c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6e,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6f, 0x66, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x0a, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x69,
	0x7a, 0x65, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x27, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x85, 0x01, 0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x40, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x72, 0x61, 0x6e,
	0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x68,
	0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x6f,
	0x72, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x72, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x32, 0xa9, 0x08, 0x0a, 0x0d, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x72,
	0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x24, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x64, 0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x59, 0x65, 0x73, 0x4f, 0x72, 0x4e, 0x6f, 0x12, 0x62, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2e, 0x2e, 0x61, 0x72,
	0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x72,
	0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x62, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x2e, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x70, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x2c, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f,
	0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x5e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x2c, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x66, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x6e, 0x0a,
	0x0a, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x2e, 0x61, 0x72,
	0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x72, 0x61, 0x6e,
	0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x30, 0x01, 0x12, 0x79, 0x0a,
	0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31, 0x2e, 0x61,
	0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a,
	0x33, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x28, 0x01, 0x30, 0x01, 0x12, 0x6a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x2e, 0x61, 0x72, 0x61, 0x6e,
	0x67, 0x6f, 0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f,
	0x64, 0x62, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6c, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x2d, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescOnce sync.Once
	file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescData = file_pkg_api_ml_storage_v1_bucketservice_proto_rawDesc
)

func file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescGZIP() []byte {
	file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescOnce.Do(func() {
		file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescData)
	})
	return file_pkg_api_ml_storage_v1_bucketservice_proto_rawDescData
}

var file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_api_ml_storage_v1_bucketservice_proto_goTypes = []interface{}{
	(*BucketRequest)(nil),         // 0: arangodb.operator.ml.storage.v1.BucketRequest
	(*PathRequest)(nil),           // 1: arangodb.operator.ml.storage.v1.PathRequest
	(*RepositoryURL)(nil),         // 2: arangodb.operator.ml.storage.v1.RepositoryURL
	(*PathSize)(nil),              // 3: arangodb.operator.ml.storage.v1.PathSize
	(*ObjectInfo)(nil),            // 4: arangodb.operator.ml.storage.v1.ObjectInfo
	(*ReadObjectChunk)(nil),       // 5: arangodb.operator.ml.storage.v1.ReadObjectChunk
	(*WriteObjectChunk)(nil),      // 6: arangodb.operator.ml.storage.v1.WriteObjectChunk
	(*WriteObjectControl)(nil),    // 7: arangodb.operator.ml.storage.v1.WriteObjectControl
	(*v1.KeyValuePair)(nil),       // 8: arangodb.operator.common.v1.KeyValuePair
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*v1.Empty)(nil),              // 10: arangodb.operator.common.v1.Empty
	(*v1.Version)(nil),            // 11: arangodb.operator.common.v1.Version
	(*v1.YesOrNo)(nil),            // 12: arangodb.operator.common.v1.YesOrNo
}
var file_pkg_api_ml_storage_v1_bucketservice_proto_depIdxs = []int32{
	8,  // 0: arangodb.operator.ml.storage.v1.BucketRequest.tags:type_name -> arangodb.operator.common.v1.KeyValuePair
	9,  // 1: arangodb.operator.ml.storage.v1.ObjectInfo.last_updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: arangodb.operator.ml.storage.v1.WriteObjectChunk.path:type_name -> arangodb.operator.ml.storage.v1.PathRequest
	10, // 3: arangodb.operator.ml.storage.v1.BucketService.GetAPIVersion:input_type -> arangodb.operator.common.v1.Empty
	0,  // 4: arangodb.operator.ml.storage.v1.BucketService.BucketExists:input_type -> arangodb.operator.ml.storage.v1.BucketRequest
	0,  // 5: arangodb.operator.ml.storage.v1.BucketService.CreateBucket:input_type -> arangodb.operator.ml.storage.v1.BucketRequest
	0,  // 6: arangodb.operator.ml.storage.v1.BucketService.DeleteBucket:input_type -> arangodb.operator.ml.storage.v1.BucketRequest
	1,  // 7: arangodb.operator.ml.storage.v1.BucketService.GetRepositoryURL:input_type -> arangodb.operator.ml.storage.v1.PathRequest
	1,  // 8: arangodb.operator.ml.storage.v1.BucketService.DeletePath:input_type -> arangodb.operator.ml.storage.v1.PathRequest
	1,  // 9: arangodb.operator.ml.storage.v1.BucketService.GetPathSize:input_type -> arangodb.operator.ml.storage.v1.PathRequest
	1,  // 10: arangodb.operator.ml.storage.v1.BucketService.ReadObject:input_type -> arangodb.operator.ml.storage.v1.PathRequest
	6,  // 11: arangodb.operator.ml.storage.v1.BucketService.WriteObject:input_type -> arangodb.operator.ml.storage.v1.WriteObjectChunk
	1,  // 12: arangodb.operator.ml.storage.v1.BucketService.GetObjectInfo:input_type -> arangodb.operator.ml.storage.v1.PathRequest
	11, // 13: arangodb.operator.ml.storage.v1.BucketService.GetAPIVersion:output_type -> arangodb.operator.common.v1.Version
	12, // 14: arangodb.operator.ml.storage.v1.BucketService.BucketExists:output_type -> arangodb.operator.common.v1.YesOrNo
	10, // 15: arangodb.operator.ml.storage.v1.BucketService.CreateBucket:output_type -> arangodb.operator.common.v1.Empty
	10, // 16: arangodb.operator.ml.storage.v1.BucketService.DeleteBucket:output_type -> arangodb.operator.common.v1.Empty
	2,  // 17: arangodb.operator.ml.storage.v1.BucketService.GetRepositoryURL:output_type -> arangodb.operator.ml.storage.v1.RepositoryURL
	10, // 18: arangodb.operator.ml.storage.v1.BucketService.DeletePath:output_type -> arangodb.operator.common.v1.Empty
	3,  // 19: arangodb.operator.ml.storage.v1.BucketService.GetPathSize:output_type -> arangodb.operator.ml.storage.v1.PathSize
	5,  // 20: arangodb.operator.ml.storage.v1.BucketService.ReadObject:output_type -> arangodb.operator.ml.storage.v1.ReadObjectChunk
	7,  // 21: arangodb.operator.ml.storage.v1.BucketService.WriteObject:output_type -> arangodb.operator.ml.storage.v1.WriteObjectControl
	4,  // 22: arangodb.operator.ml.storage.v1.BucketService.GetObjectInfo:output_type -> arangodb.operator.ml.storage.v1.ObjectInfo
	13, // [13:23] is the sub-list for method output_type
	3,  // [3:13] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_api_ml_storage_v1_bucketservice_proto_init() }
func file_pkg_api_ml_storage_v1_bucketservice_proto_init() {
	if File_pkg_api_ml_storage_v1_bucketservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketRequest); i {
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
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathRequest); i {
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
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryURL); i {
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
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathSize); i {
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
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectInfo); i {
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
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadObjectChunk); i {
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
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteObjectChunk); i {
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
		file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteObjectControl); i {
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
			RawDescriptor: file_pkg_api_ml_storage_v1_bucketservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_ml_storage_v1_bucketservice_proto_goTypes,
		DependencyIndexes: file_pkg_api_ml_storage_v1_bucketservice_proto_depIdxs,
		MessageInfos:      file_pkg_api_ml_storage_v1_bucketservice_proto_msgTypes,
	}.Build()
	File_pkg_api_ml_storage_v1_bucketservice_proto = out.File
	file_pkg_api_ml_storage_v1_bucketservice_proto_rawDesc = nil
	file_pkg_api_ml_storage_v1_bucketservice_proto_goTypes = nil
	file_pkg_api_ml_storage_v1_bucketservice_proto_depIdxs = nil
}