// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gogoproto/gogo.proto

package gogoproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
	Filename:      "gogoproto/gogo.proto",
}

func init() {
	proto.RegisterExtension(E_Nullable)
}

func init() { proto.RegisterFile("gogoproto/gogo.proto", fileDescriptor_c586470e9b64aee7) }

var fileDescriptor_c586470e9b64aee7 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0xb1, 0xf4, 0xc0, 0x4c, 0x21, 0x4e, 0xb8, 0xa8, 0x94,
	0x42, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x92, 0x5a,
	0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x5f, 0x04, 0x51, 0x6c, 0x65, 0xcd, 0xc5, 0x91, 0x57, 0x9a,
	0x93, 0x93, 0x98, 0x94, 0x93, 0x2a, 0x24, 0xab, 0x07, 0x51, 0xae, 0x07, 0x53, 0xae, 0xe7, 0x96,
	0x99, 0x9a, 0x93, 0xe2, 0x5f, 0x50, 0x92, 0x99, 0x9f, 0x57, 0x2c, 0xf1, 0xf2, 0x37, 0xb3, 0x02,
	0xa3, 0x06, 0x47, 0x10, 0x5c, 0x83, 0x93, 0x4a, 0x94, 0x52, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0x2e, 0xd8, 0x09, 0x08, 0x9b, 0xe0, 0x8e, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x48, 0x2a, 0x59, 0xa6, 0x00, 0x00, 0x00,
}
