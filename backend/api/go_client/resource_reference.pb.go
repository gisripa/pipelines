// Copyright 2019 Google LLC
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
// source: backend/api/resource_reference.proto

package go_client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResourceType int32

const (
	ResourceType_UNKNOWN_RESOURCE_TYPE ResourceType = 0
	ResourceType_EXPERIMENT            ResourceType = 1
	ResourceType_JOB                   ResourceType = 2
)

var ResourceType_name = map[int32]string{
	0: "UNKNOWN_RESOURCE_TYPE",
	1: "EXPERIMENT",
	2: "JOB",
}
var ResourceType_value = map[string]int32{
	"UNKNOWN_RESOURCE_TYPE": 0,
	"EXPERIMENT":            1,
	"JOB":                   2,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_resource_reference_4b926f4cbb4ee3f2, []int{0}
}

type Relationship int32

const (
	Relationship_UNKNOWN_RELATIONSHIP Relationship = 0
	Relationship_OWNER                Relationship = 1
	Relationship_CREATOR              Relationship = 2
)

var Relationship_name = map[int32]string{
	0: "UNKNOWN_RELATIONSHIP",
	1: "OWNER",
	2: "CREATOR",
}
var Relationship_value = map[string]int32{
	"UNKNOWN_RELATIONSHIP": 0,
	"OWNER":                1,
	"CREATOR":              2,
}

func (x Relationship) String() string {
	return proto.EnumName(Relationship_name, int32(x))
}
func (Relationship) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_resource_reference_4b926f4cbb4ee3f2, []int{1}
}

type ResourceKey struct {
	Type                 ResourceType `protobuf:"varint,1,opt,name=type,proto3,enum=api.ResourceType" json:"type,omitempty"`
	Id                   string       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResourceKey) Reset()         { *m = ResourceKey{} }
func (m *ResourceKey) String() string { return proto.CompactTextString(m) }
func (*ResourceKey) ProtoMessage()    {}
func (*ResourceKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_reference_4b926f4cbb4ee3f2, []int{0}
}
func (m *ResourceKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceKey.Unmarshal(m, b)
}
func (m *ResourceKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceKey.Marshal(b, m, deterministic)
}
func (dst *ResourceKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceKey.Merge(dst, src)
}
func (m *ResourceKey) XXX_Size() int {
	return xxx_messageInfo_ResourceKey.Size(m)
}
func (m *ResourceKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceKey.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceKey proto.InternalMessageInfo

func (m *ResourceKey) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_UNKNOWN_RESOURCE_TYPE
}

func (m *ResourceKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ResourceReference struct {
	Key                  *ResourceKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Relationship         Relationship `protobuf:"varint,2,opt,name=relationship,proto3,enum=api.Relationship" json:"relationship,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResourceReference) Reset()         { *m = ResourceReference{} }
func (m *ResourceReference) String() string { return proto.CompactTextString(m) }
func (*ResourceReference) ProtoMessage()    {}
func (*ResourceReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_reference_4b926f4cbb4ee3f2, []int{1}
}
func (m *ResourceReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceReference.Unmarshal(m, b)
}
func (m *ResourceReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceReference.Marshal(b, m, deterministic)
}
func (dst *ResourceReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceReference.Merge(dst, src)
}
func (m *ResourceReference) XXX_Size() int {
	return xxx_messageInfo_ResourceReference.Size(m)
}
func (m *ResourceReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceReference.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceReference proto.InternalMessageInfo

func (m *ResourceReference) GetKey() *ResourceKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ResourceReference) GetRelationship() Relationship {
	if m != nil {
		return m.Relationship
	}
	return Relationship_UNKNOWN_RELATIONSHIP
}

func init() {
	proto.RegisterType((*ResourceKey)(nil), "api.ResourceKey")
	proto.RegisterType((*ResourceReference)(nil), "api.ResourceReference")
	proto.RegisterEnum("api.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterEnum("api.Relationship", Relationship_name, Relationship_value)
}

func init() {
	proto.RegisterFile("backend/api/resource_reference.proto", fileDescriptor_resource_reference_4b926f4cbb4ee3f2)
}

var fileDescriptor_resource_reference_4b926f4cbb4ee3f2 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x6b, 0xf2, 0x40,
	0x10, 0x47, 0x4d, 0xfc, 0xbe, 0x8a, 0xa3, 0xc8, 0xba, 0xb4, 0x60, 0x6f, 0x22, 0x2d, 0x88, 0x07,
	0x05, 0x4b, 0xef, 0x55, 0xbb, 0x50, 0x9b, 0x76, 0x37, 0x8c, 0x11, 0xdb, 0x53, 0x88, 0xc9, 0x94,
	0x2e, 0x96, 0x64, 0x59, 0xd3, 0x43, 0xfe, 0xfb, 0x62, 0x68, 0x48, 0x3d, 0xbf, 0xc7, 0xfb, 0x0d,
	0x03, 0x37, 0xfb, 0x28, 0x3e, 0x50, 0x9a, 0xcc, 0x22, 0xa3, 0x67, 0x96, 0x8e, 0xd9, 0xb7, 0x8d,
	0x29, 0xb4, 0xf4, 0x41, 0x96, 0xd2, 0x98, 0xa6, 0xc6, 0x66, 0x79, 0xc6, 0x9b, 0x91, 0xd1, 0xa3,
	0x47, 0xe8, 0xe0, 0xaf, 0xe0, 0x51, 0xc1, 0x6f, 0xe1, 0x5f, 0x5e, 0x18, 0x1a, 0x38, 0x43, 0x67,
	0xdc, 0x9b, 0xf7, 0xa7, 0x91, 0xd1, 0xd3, 0x8a, 0x07, 0x85, 0x21, 0x2c, 0x31, 0xef, 0x81, 0xab,
	0x93, 0x81, 0x3b, 0x74, 0xc6, 0x6d, 0x74, 0x75, 0x32, 0x4a, 0xa1, 0x5f, 0x59, 0x58, 0xad, 0xf0,
	0x11, 0x34, 0x0f, 0x54, 0x94, 0xa9, 0xce, 0x9c, 0x9d, 0xa5, 0x3c, 0x2a, 0xf0, 0x04, 0xf9, 0x3d,
	0x74, 0x2d, 0x7d, 0x45, 0xb9, 0xce, 0xd2, 0xe3, 0xa7, 0x36, 0x65, 0xb2, 0xde, 0xad, 0x01, 0x9e,
	0x69, 0x93, 0x25, 0x74, 0xff, 0x5e, 0xc5, 0xaf, 0xe1, 0x6a, 0x2b, 0x3d, 0xa9, 0x76, 0x32, 0x44,
	0xb1, 0x51, 0x5b, 0x5c, 0x89, 0x30, 0x78, 0xf7, 0x05, 0x6b, 0xf0, 0x1e, 0x80, 0x78, 0xf3, 0x05,
	0xae, 0x5f, 0x85, 0x0c, 0x98, 0xc3, 0x5b, 0xd0, 0x7c, 0x56, 0x4b, 0xe6, 0x4e, 0x1e, 0x4e, 0x8d,
	0xba, 0xc9, 0x07, 0x70, 0x59, 0x37, 0x5e, 0x16, 0xc1, 0x5a, 0xc9, 0xcd, 0xd3, 0xda, 0x67, 0x0d,
	0xde, 0x86, 0xff, 0x6a, 0x27, 0x05, 0x32, 0x87, 0x77, 0xa0, 0xb5, 0x42, 0xb1, 0x08, 0x14, 0x32,
	0x77, 0x7f, 0x51, 0xfe, 0xf1, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xce, 0x2d, 0x81, 0x6f,
	0x01, 0x00, 0x00,
}
