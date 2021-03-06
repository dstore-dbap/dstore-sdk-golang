// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetAllDatabaseUsers_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetAllDatabaseUsers_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetAllDatabaseUsers_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetAllDatabaseUsers_Ad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dstore_values "gosdk.dstore.de/values"
import dstore_engine_error "gosdk.dstore.de/engine/error"
import dstore_engine_message "gosdk.dstore.de/engine/message"
import dstore_engine_metainformation "gosdk.dstore.de/engine/metainformation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Parameters struct {
	IncludeSpecialUser     *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=include_special_user,json=includeSpecialUser" json:"include_special_user,omitempty"`
	IncludeSpecialUserNull bool                        `protobuf:"varint,1001,opt,name=include_special_user_null,json=includeSpecialUserNull" json:"include_special_user_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetIncludeSpecialUser() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeSpecialUser
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetError() *dstore_engine_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetMetaInformation() []*dstore_engine_metainformation.MetaInformation {
	if m != nil {
		return m.MetaInformation
	}
	return nil
}

func (m *Response) GetMessage() []*dstore_engine_message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Response) GetRow() []*Response_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

type Response_Row struct {
	RowId int32                      `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Name  *dstore_values.StringValue `protobuf:"bytes,10001,opt,name=name" json:"name,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetName() *dstore_values.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetAllDatabaseUsers_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetAllDatabaseUsers_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetAllDatabaseUsers_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetAllDatabaseUsers_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x51, 0xab, 0xd3, 0x30,
	0x14, 0x66, 0xb7, 0x77, 0xf7, 0x5e, 0xe2, 0x83, 0x12, 0x65, 0x74, 0x1d, 0xc8, 0x98, 0x0f, 0xfa,
	0x94, 0x8a, 0x22, 0x0c, 0x7d, 0x9a, 0x28, 0x63, 0x0f, 0x2b, 0x12, 0x51, 0x70, 0x2f, 0x25, 0x5b,
	0x8e, 0x25, 0x98, 0x26, 0x23, 0x69, 0xdd, 0xdf, 0xd0, 0x27, 0x7f, 0xa3, 0xfe, 0x0a, 0xd3, 0x26,
	0x63, 0xb6, 0x53, 0xf0, 0xbe, 0xb4, 0x3d, 0xfd, 0xbe, 0xef, 0xe4, 0x9c, 0xef, 0x9c, 0xa0, 0x39,
	0xb7, 0x95, 0x36, 0x90, 0x82, 0x2a, 0x84, 0x82, 0x74, 0x6f, 0xf4, 0x0e, 0x78, 0x6d, 0xc0, 0xa6,
	0xa5, 0xc8, 0x97, 0x50, 0x2d, 0xa4, 0x7c, 0xc3, 0x2a, 0xb6, 0x65, 0x16, 0x3e, 0x58, 0x30, 0x36,
	0x5f, 0x70, 0xe2, 0x48, 0x95, 0xc6, 0x8f, 0xbd, 0x92, 0x78, 0x25, 0xf9, 0x27, 0x3d, 0xb9, 0x1f,
	0x8e, 0xf8, 0xca, 0x64, 0x0d, 0xd6, 0xab, 0x93, 0x71, 0xf7, 0x5c, 0x30, 0x46, 0x9b, 0x00, 0x4d,
	0xba, 0x50, 0x09, 0xd6, 0xb2, 0x02, 0x02, 0xf8, 0xa8, 0x0f, 0x56, 0x4c, 0xa8, 0xcf, 0xda, 0x94,
	0xac, 0x12, 0x5a, 0x79, 0xd2, 0xec, 0xc7, 0x00, 0xa1, 0x77, 0xcc, 0x30, 0x87, 0xba, 0x1a, 0xf0,
	0x1a, 0x3d, 0x10, 0x6a, 0x27, 0x6b, 0x0e, 0xb9, 0xdd, 0xc3, 0x4e, 0x30, 0x99, 0xd7, 0xae, 0xb8,
	0x78, 0x30, 0x1d, 0x3c, 0xb9, 0xf3, 0x6c, 0x42, 0x42, 0x23, 0xa1, 0xbe, 0xad, 0xd6, 0x12, 0x98,
	0xfa, 0xd8, 0x44, 0x14, 0x07, 0xe1, 0x7b, 0xaf, 0x6b, 0x7a, 0xc2, 0x2f, 0xd1, 0xf8, 0x6f, 0xe9,
	0x72, 0x55, 0x4b, 0x19, 0xff, 0xbc, 0x76, 0x49, 0x6f, 0xe8, 0xe8, 0x5c, 0x97, 0x39, 0x78, 0xf6,
	0xeb, 0x02, 0xdd, 0x50, 0xb0, 0x7b, 0xad, 0x2c, 0xe0, 0xa7, 0x68, 0xd8, 0xf6, 0x1d, 0x0a, 0x49,
	0x48, 0xd7, 0x51, 0xef, 0xc9, 0xdb, 0xe6, 0x49, 0x3d, 0x11, 0x7f, 0x42, 0xf7, 0x9a, 0x8e, 0xf3,
	0x3f, 0x5a, 0x8e, 0x2f, 0xa6, 0x91, 0x13, 0x93, 0x9e, 0xb8, 0x6f, 0xcc, 0xda, 0xc5, 0xab, 0x53,
	0x4c, 0xef, 0x96, 0xdd, 0x1f, 0x78, 0x8e, 0xae, 0x83, 0xd3, 0x71, 0xd4, 0x66, 0x7c, 0x78, 0x96,
	0xd1, 0xcf, 0x61, 0xed, 0xdf, 0xf4, 0x48, 0xc7, 0x4b, 0x14, 0x19, 0x7d, 0x88, 0x2f, 0x5b, 0xd5,
	0x0b, 0xf2, 0x9f, 0x6b, 0x41, 0x8e, 0x36, 0x10, 0xaa, 0x0f, 0xb4, 0xc9, 0x90, 0x64, 0x28, 0x72,
	0xdf, 0x78, 0x84, 0xae, 0x5c, 0x94, 0x0b, 0x1e, 0x7f, 0xcb, 0x9c, 0x31, 0x43, 0x3a, 0x74, 0xe1,
	0x8a, 0xe3, 0x14, 0x5d, 0x2a, 0x37, 0xd2, 0xf8, 0x7b, 0xd6, 0xb5, 0x2b, 0xcc, 0xcd, 0x56, 0x46,
	0xa8, 0xc2, 0x8f, 0xad, 0x25, 0xbe, 0xde, 0xa0, 0x89, 0xd0, 0xbd, 0x7a, 0x4e, 0x0b, 0xbe, 0x79,
	0x55, 0x68, 0xcb, 0xbf, 0x1c, 0x71, 0x7e, 0xab, 0x3b, 0xb0, 0xbd, 0x6a, 0x37, 0xed, 0xf9, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x77, 0xc4, 0x5c, 0x40, 0x03, 0x00, 0x00,
}
