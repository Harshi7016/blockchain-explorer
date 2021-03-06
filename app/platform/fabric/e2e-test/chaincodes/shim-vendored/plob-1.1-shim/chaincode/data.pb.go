// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	data.proto

It has these top-level messages:
	PermissionedBlob
*/
package main

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

type PermissionedBlob struct {
	Owner []byte `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Blob  []byte `protobuf:"bytes,2,opt,name=Blob,proto3" json:"Blob,omitempty"`
}

func (m *PermissionedBlob) Reset()                    { *m = PermissionedBlob{} }
func (m *PermissionedBlob) String() string            { return proto.CompactTextString(m) }
func (*PermissionedBlob) ProtoMessage()               {}
func (*PermissionedBlob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*PermissionedBlob)(nil), "main.PermissionedBlob")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0xb2, 0xe1, 0x12,
	0x08, 0x48, 0x2d, 0xca, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0x4b, 0x4d, 0x71, 0xca, 0xc9, 0x4f,
	0x12, 0x12, 0xe1, 0x62, 0xf5, 0x2f, 0xcf, 0x4b, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09,
	0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0x40, 0xb2, 0x12, 0x4c, 0x60, 0x41, 0x30, 0x3b, 0x89, 0x0d,
	0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe0, 0xd3, 0xcf, 0x58, 0x00, 0x00, 0x00,
}
