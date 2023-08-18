// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/experimental/v1/editor_service.proto

package v1

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

// The type of the editor, currently browser or desktop.
type EditorOption_Type int32

const (
	EditorOption_TYPE_UNSPECIFIED EditorOption_Type = 0
	EditorOption_TYPE_BROWSER     EditorOption_Type = 1
	EditorOption_TYPE_DESKTOP     EditorOption_Type = 2
)

// Enum value maps for EditorOption_Type.
var (
	EditorOption_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_BROWSER",
		2: "TYPE_DESKTOP",
	}
	EditorOption_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_BROWSER":     1,
		"TYPE_DESKTOP":     2,
	}
)

func (x EditorOption_Type) Enum() *EditorOption_Type {
	p := new(EditorOption_Type)
	*p = x
	return p
}

func (x EditorOption_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EditorOption_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_gitpod_experimental_v1_editor_service_proto_enumTypes[0].Descriptor()
}

func (EditorOption_Type) Type() protoreflect.EnumType {
	return &file_gitpod_experimental_v1_editor_service_proto_enumTypes[0]
}

func (x EditorOption_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EditorOption_Type.Descriptor instead.
func (EditorOption_Type) EnumDescriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_editor_service_proto_rawDescGZIP(), []int{2, 0}
}

type ListEditorOptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListEditorOptionsRequest) Reset() {
	*x = ListEditorOptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEditorOptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEditorOptionsRequest) ProtoMessage() {}

func (x *ListEditorOptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEditorOptionsRequest.ProtoReflect.Descriptor instead.
func (*ListEditorOptionsRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_editor_service_proto_rawDescGZIP(), []int{0}
}

type ListEditorOptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*EditorOption `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListEditorOptionsResponse) Reset() {
	*x = ListEditorOptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEditorOptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEditorOptionsResponse) ProtoMessage() {}

func (x *ListEditorOptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEditorOptionsResponse.ProtoReflect.Descriptor instead.
func (*ListEditorOptionsResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_editor_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListEditorOptionsResponse) GetResult() []*EditorOption {
	if x != nil {
		return x.Result
	}
	return nil
}

type EditorOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier for an editor.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human readable title text of the editor (plain text only).
	Title string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type  EditorOption_Type `protobuf:"varint,3,opt,name=type,proto3,enum=gitpod.experimental.v1.EditorOption_Type" json:"type,omitempty"`
	// The logo for the editor
	Logo string `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	// Text of an optional label next to the editor option like “Insiders” (plain
	// text only).
	Label  string             `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Stable *EditorOption_Kind `protobuf:"bytes,6,opt,name=stable,proto3" json:"stable,omitempty"`
	Latest *EditorOption_Kind `protobuf:"bytes,7,opt,name=latest,proto3" json:"latest,omitempty"`
}

func (x *EditorOption) Reset() {
	*x = EditorOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditorOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditorOption) ProtoMessage() {}

func (x *EditorOption) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditorOption.ProtoReflect.Descriptor instead.
func (*EditorOption) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_editor_service_proto_rawDescGZIP(), []int{2}
}

func (x *EditorOption) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EditorOption) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditorOption) GetType() EditorOption_Type {
	if x != nil {
		return x.Type
	}
	return EditorOption_TYPE_UNSPECIFIED
}

func (x *EditorOption) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *EditorOption) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *EditorOption) GetStable() *EditorOption_Kind {
	if x != nil {
		return x.Stable
	}
	return nil
}

func (x *EditorOption) GetLatest() *EditorOption_Kind {
	if x != nil {
		return x.Latest
	}
	return nil
}

type EditorOption_Kind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The semantic version of the editor.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *EditorOption_Kind) Reset() {
	*x = EditorOption_Kind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditorOption_Kind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditorOption_Kind) ProtoMessage() {}

func (x *EditorOption_Kind) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_editor_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditorOption_Kind.ProtoReflect.Descriptor instead.
func (*EditorOption_Kind) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_editor_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *EditorOption_Kind) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_gitpod_experimental_v1_editor_service_proto protoreflect.FileDescriptor

var file_gitpod_experimental_v1_editor_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x59, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x87, 0x03, 0x0a,
	0x0c, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x06, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x41, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x4f, 0x57, 0x53,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x53,
	0x4b, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x32, 0x8b, 0x01, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x2e,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74,
	0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_experimental_v1_editor_service_proto_rawDescOnce sync.Once
	file_gitpod_experimental_v1_editor_service_proto_rawDescData = file_gitpod_experimental_v1_editor_service_proto_rawDesc
)

func file_gitpod_experimental_v1_editor_service_proto_rawDescGZIP() []byte {
	file_gitpod_experimental_v1_editor_service_proto_rawDescOnce.Do(func() {
		file_gitpod_experimental_v1_editor_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_experimental_v1_editor_service_proto_rawDescData)
	})
	return file_gitpod_experimental_v1_editor_service_proto_rawDescData
}

var file_gitpod_experimental_v1_editor_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gitpod_experimental_v1_editor_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gitpod_experimental_v1_editor_service_proto_goTypes = []interface{}{
	(EditorOption_Type)(0),            // 0: gitpod.experimental.v1.EditorOption.Type
	(*ListEditorOptionsRequest)(nil),  // 1: gitpod.experimental.v1.ListEditorOptionsRequest
	(*ListEditorOptionsResponse)(nil), // 2: gitpod.experimental.v1.ListEditorOptionsResponse
	(*EditorOption)(nil),              // 3: gitpod.experimental.v1.EditorOption
	(*EditorOption_Kind)(nil),         // 4: gitpod.experimental.v1.EditorOption.Kind
}
var file_gitpod_experimental_v1_editor_service_proto_depIdxs = []int32{
	3, // 0: gitpod.experimental.v1.ListEditorOptionsResponse.result:type_name -> gitpod.experimental.v1.EditorOption
	0, // 1: gitpod.experimental.v1.EditorOption.type:type_name -> gitpod.experimental.v1.EditorOption.Type
	4, // 2: gitpod.experimental.v1.EditorOption.stable:type_name -> gitpod.experimental.v1.EditorOption.Kind
	4, // 3: gitpod.experimental.v1.EditorOption.latest:type_name -> gitpod.experimental.v1.EditorOption.Kind
	1, // 4: gitpod.experimental.v1.EditorService.ListEditorOptions:input_type -> gitpod.experimental.v1.ListEditorOptionsRequest
	2, // 5: gitpod.experimental.v1.EditorService.ListEditorOptions:output_type -> gitpod.experimental.v1.ListEditorOptionsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gitpod_experimental_v1_editor_service_proto_init() }
func file_gitpod_experimental_v1_editor_service_proto_init() {
	if File_gitpod_experimental_v1_editor_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_experimental_v1_editor_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEditorOptionsRequest); i {
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
		file_gitpod_experimental_v1_editor_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEditorOptionsResponse); i {
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
		file_gitpod_experimental_v1_editor_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditorOption); i {
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
		file_gitpod_experimental_v1_editor_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditorOption_Kind); i {
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
			RawDescriptor: file_gitpod_experimental_v1_editor_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitpod_experimental_v1_editor_service_proto_goTypes,
		DependencyIndexes: file_gitpod_experimental_v1_editor_service_proto_depIdxs,
		EnumInfos:         file_gitpod_experimental_v1_editor_service_proto_enumTypes,
		MessageInfos:      file_gitpod_experimental_v1_editor_service_proto_msgTypes,
	}.Build()
	File_gitpod_experimental_v1_editor_service_proto = out.File
	file_gitpod_experimental_v1_editor_service_proto_rawDesc = nil
	file_gitpod_experimental_v1_editor_service_proto_goTypes = nil
	file_gitpod_experimental_v1_editor_service_proto_depIdxs = nil
}
