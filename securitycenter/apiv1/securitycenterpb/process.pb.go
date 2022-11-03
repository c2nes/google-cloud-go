// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/securitycenter/v1/process.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents an operating system process.
type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The process name visible in utilities like `top` and `ps`; it can
	// be accessed via `/proc/[pid]/comm` and changed with `prctl(PR_SET_NAME)`.
	Name string `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	// File information for the process executable.
	Binary *File `protobuf:"bytes,3,opt,name=binary,proto3" json:"binary,omitempty"`
	// File information for libraries loaded by the process.
	Libraries []*File `protobuf:"bytes,4,rep,name=libraries,proto3" json:"libraries,omitempty"`
	// When the process represents the invocation of a script,
	// `binary` provides information about the interpreter while `script`
	// provides information about the script file provided to the
	// interpreter.
	Script *File `protobuf:"bytes,5,opt,name=script,proto3" json:"script,omitempty"`
	// Process arguments as JSON encoded strings.
	Args []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	// True if `args` is incomplete.
	ArgumentsTruncated bool `protobuf:"varint,7,opt,name=arguments_truncated,json=argumentsTruncated,proto3" json:"arguments_truncated,omitempty"`
	// Process environment variables.
	EnvVariables []*EnvironmentVariable `protobuf:"bytes,8,rep,name=env_variables,json=envVariables,proto3" json:"env_variables,omitempty"`
	// True if `env_variables` is incomplete.
	EnvVariablesTruncated bool `protobuf:"varint,9,opt,name=env_variables_truncated,json=envVariablesTruncated,proto3" json:"env_variables_truncated,omitempty"`
	// The process id.
	Pid int64 `protobuf:"varint,10,opt,name=pid,proto3" json:"pid,omitempty"`
	// The parent process id.
	ParentPid int64 `protobuf:"varint,11,opt,name=parent_pid,json=parentPid,proto3" json:"parent_pid,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_process_proto_rawDescGZIP(), []int{0}
}

func (x *Process) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Process) GetBinary() *File {
	if x != nil {
		return x.Binary
	}
	return nil
}

func (x *Process) GetLibraries() []*File {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *Process) GetScript() *File {
	if x != nil {
		return x.Script
	}
	return nil
}

func (x *Process) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Process) GetArgumentsTruncated() bool {
	if x != nil {
		return x.ArgumentsTruncated
	}
	return false
}

func (x *Process) GetEnvVariables() []*EnvironmentVariable {
	if x != nil {
		return x.EnvVariables
	}
	return nil
}

func (x *Process) GetEnvVariablesTruncated() bool {
	if x != nil {
		return x.EnvVariablesTruncated
	}
	return false
}

func (x *Process) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Process) GetParentPid() int64 {
	if x != nil {
		return x.ParentPid
	}
	return 0
}

// EnvironmentVariable is a name-value pair to store environment variables for
// Process.
type EnvironmentVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Environment variable name as a JSON encoded string.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Environment variable value as a JSON encoded string.
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *EnvironmentVariable) Reset() {
	*x = EnvironmentVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentVariable) ProtoMessage() {}

func (x *EnvironmentVariable) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentVariable.ProtoReflect.Descriptor instead.
func (*EnvironmentVariable) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_process_proto_rawDescGZIP(), []int{1}
}

func (x *EnvironmentVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvironmentVariable) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_google_cloud_securitycenter_v1_process_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_process_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x29,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x42, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x2f, 0x0a,
	0x13, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x74, 0x72, 0x75, 0x6e, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x58,
	0x0a, 0x0d, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x6e, 0x76, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x65, 0x6e, 0x76, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x69,
	0x64, 0x22, 0x3b, 0x0a, 0x13, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0xe8,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_securitycenter_v1_process_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_process_proto_rawDescData = file_google_cloud_securitycenter_v1_process_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_process_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_process_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_process_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_process_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_process_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securitycenter_v1_process_proto_goTypes = []interface{}{
	(*Process)(nil),             // 0: google.cloud.securitycenter.v1.Process
	(*EnvironmentVariable)(nil), // 1: google.cloud.securitycenter.v1.EnvironmentVariable
	(*File)(nil),                // 2: google.cloud.securitycenter.v1.File
}
var file_google_cloud_securitycenter_v1_process_proto_depIdxs = []int32{
	2, // 0: google.cloud.securitycenter.v1.Process.binary:type_name -> google.cloud.securitycenter.v1.File
	2, // 1: google.cloud.securitycenter.v1.Process.libraries:type_name -> google.cloud.securitycenter.v1.File
	2, // 2: google.cloud.securitycenter.v1.Process.script:type_name -> google.cloud.securitycenter.v1.File
	1, // 3: google.cloud.securitycenter.v1.Process.env_variables:type_name -> google.cloud.securitycenter.v1.EnvironmentVariable
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_process_proto_init() }
func file_google_cloud_securitycenter_v1_process_proto_init() {
	if File_google_cloud_securitycenter_v1_process_proto != nil {
		return
	}
	file_google_cloud_securitycenter_v1_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
		file_google_cloud_securitycenter_v1_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentVariable); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_process_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_process_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_process_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_process_proto = out.File
	file_google_cloud_securitycenter_v1_process_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_process_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_process_proto_depIdxs = nil
}
