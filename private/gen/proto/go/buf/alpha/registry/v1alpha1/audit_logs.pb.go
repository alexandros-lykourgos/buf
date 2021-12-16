// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: buf/alpha/registry/v1alpha1/audit_logs.proto

package registryv1alpha1

import (
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/audit/v1alpha1"
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

type ListAuditLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize uint32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The first page is returned if this is empty.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Reverse   bool   `protobuf:"varint,3,opt,name=reverse,proto3" json:"reverse,omitempty"`
	// Optionally specifies a start time for the query.
	// Only audit logs from later than or equal to this time will be returned.
	// This value is inclusive so that start_time and end_time together
	// form a closed-open range [start_time, end_time), allowing
	// consecutive pages to cover exclusive logs.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Optionally specifies an end time for the query.
	// Only audit logs from earlier than this time will be returned.
	// This value is exclusive, so that start_time and end_time together
	// form a closed-open range [start_time, end_time), allowing
	// consecutive pages to cover exclusive logs.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *ListAuditLogsRequest) Reset() {
	*x = ListAuditLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogsRequest) ProtoMessage() {}

func (x *ListAuditLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogsRequest.ProtoReflect.Descriptor instead.
func (*ListAuditLogsRequest) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescGZIP(), []int{0}
}

func (x *ListAuditLogsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAuditLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAuditLogsRequest) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

func (x *ListAuditLogsRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ListAuditLogsRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type ListAuditLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuditLogs []*v1alpha1.Event `protobuf:"bytes,1,rep,name=audit_logs,json=auditLogs,proto3" json:"audit_logs,omitempty"`
	// There are no more pages if this is empty.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListAuditLogsResponse) Reset() {
	*x = ListAuditLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogsResponse) ProtoMessage() {}

func (x *ListAuditLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogsResponse.ProtoReflect.Descriptor instead.
func (*ListAuditLogsResponse) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescGZIP(), []int{1}
}

func (x *ListAuditLogsResponse) GetAuditLogs() []*v1alpha1.Event {
	if x != nil {
		return x.AuditLogs
	}
	return nil
}

func (x *ListAuditLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_buf_alpha_registry_v1alpha1_audit_logs_proto protoreflect.FileDescriptor

var file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x8a, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x31, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x9b, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x52, 0xaa, 0x02, 0x1b, 0x42, 0x75,
	0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x5c,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x27, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescOnce sync.Once
	file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescData = file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDesc
)

func file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescGZIP() []byte {
	file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescOnce.Do(func() {
		file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescData)
	})
	return file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDescData
}

var file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_alpha_registry_v1alpha1_audit_logs_proto_goTypes = []interface{}{
	(*ListAuditLogsRequest)(nil),  // 0: buf.alpha.registry.v1alpha1.ListAuditLogsRequest
	(*ListAuditLogsResponse)(nil), // 1: buf.alpha.registry.v1alpha1.ListAuditLogsResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*v1alpha1.Event)(nil),        // 3: buf.alpha.audit.v1alpha1.Event
}
var file_buf_alpha_registry_v1alpha1_audit_logs_proto_depIdxs = []int32{
	2, // 0: buf.alpha.registry.v1alpha1.ListAuditLogsRequest.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: buf.alpha.registry.v1alpha1.ListAuditLogsRequest.end_time:type_name -> google.protobuf.Timestamp
	3, // 2: buf.alpha.registry.v1alpha1.ListAuditLogsResponse.audit_logs:type_name -> buf.alpha.audit.v1alpha1.Event
	0, // 3: buf.alpha.registry.v1alpha1.AuditLogsService.ListAuditLogs:input_type -> buf.alpha.registry.v1alpha1.ListAuditLogsRequest
	1, // 4: buf.alpha.registry.v1alpha1.AuditLogsService.ListAuditLogs:output_type -> buf.alpha.registry.v1alpha1.ListAuditLogsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_buf_alpha_registry_v1alpha1_audit_logs_proto_init() }
func file_buf_alpha_registry_v1alpha1_audit_logs_proto_init() {
	if File_buf_alpha_registry_v1alpha1_audit_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogsRequest); i {
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
		file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogsResponse); i {
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
			RawDescriptor: file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_alpha_registry_v1alpha1_audit_logs_proto_goTypes,
		DependencyIndexes: file_buf_alpha_registry_v1alpha1_audit_logs_proto_depIdxs,
		MessageInfos:      file_buf_alpha_registry_v1alpha1_audit_logs_proto_msgTypes,
	}.Build()
	File_buf_alpha_registry_v1alpha1_audit_logs_proto = out.File
	file_buf_alpha_registry_v1alpha1_audit_logs_proto_rawDesc = nil
	file_buf_alpha_registry_v1alpha1_audit_logs_proto_goTypes = nil
	file_buf_alpha_registry_v1alpha1_audit_logs_proto_depIdxs = nil
}