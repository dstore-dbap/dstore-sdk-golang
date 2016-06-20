// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ResetSymbolIDs_Ad.proto
// DO NOT EDIT!

/*
Package im_ResetSymbolIDs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ResetSymbolIDs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ResetSymbolIDs_Ad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "dstore/values"
import dstore_engine_error "dstore/engine/error"
import dstore_engine_message "dstore/engine/message"
import dstore_engine_metainformation "dstore/engine/metainformation"

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
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ResetSymbolIDs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ResetSymbolIDs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ResetSymbolIDs_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xa6, 0x8d, 0xad, 0x32, 0x0a, 0xca, 0x0a, 0x12, 0x53, 0x94, 0x52, 0x2f, 0x5e, 0xdc, 0x14,
	0x7b, 0xf1, 0x6a, 0xd1, 0x43, 0xc0, 0x8a, 0xac, 0x27, 0xbd, 0x84, 0xad, 0x19, 0xcb, 0x42, 0x93,
	0x29, 0xbb, 0xa9, 0xc5, 0xb7, 0xf0, 0x61, 0x7c, 0x40, 0x37, 0xd9, 0x94, 0x9a, 0x28, 0x88, 0x97,
	0x24, 0x93, 0xef, 0x6f, 0x66, 0x77, 0x60, 0x94, 0x98, 0x9c, 0x34, 0x86, 0x98, 0xcd, 0x54, 0x86,
	0xe1, 0x42, 0xd3, 0x0b, 0x26, 0x4b, 0x8d, 0x26, 0x54, 0x69, 0x2c, 0xd0, 0x60, 0xfe, 0xf8, 0x9e,
	0x4e, 0x69, 0x1e, 0xdd, 0x98, 0xf8, 0x3a, 0xe1, 0x16, 0xcf, 0x89, 0x0d, 0x9c, 0x88, 0x3b, 0x11,
	0xff, 0x8d, 0x19, 0x1c, 0x56, 0xc6, 0x6f, 0x72, 0xbe, 0x44, 0xe3, 0x84, 0xc1, 0x71, 0x3d, 0x0d,
	0xb5, 0x26, 0x5d, 0x41, 0xbd, 0x3a, 0x94, 0xa2, 0x31, 0x72, 0x86, 0x15, 0x78, 0xd6, 0x04, 0x73,
	0xa9, 0xb2, 0x57, 0xd2, 0xa9, 0xcc, 0x15, 0x65, 0x8e, 0x34, 0xd8, 0x03, 0x78, 0x90, 0x5a, 0x5a,
	0x10, 0xb5, 0x19, 0x7c, 0xb6, 0x61, 0xc7, 0x76, 0xb5, 0xa0, 0xcc, 0x20, 0x1b, 0x42, 0xa7, 0xcc,
	0xf2, 0x5b, 0xfd, 0xd6, 0xf9, 0xee, 0x65, 0xc0, 0xeb, 0x03, 0xb8, 0x3e, 0x6e, 0x8b, 0xa7, 0x70,
	0x44, 0xf6, 0x04, 0x07, 0x45, 0x4a, 0xfc, 0x2d, 0xc6, 0x6f, 0xf7, 0x3d, 0x2b, 0xe6, 0x0d, 0x71,
	0xb3, 0x99, 0x89, 0xad, 0xa3, 0x4d, 0x2d, 0xf6, 0xd3, 0xfa, 0x0f, 0x76, 0x05, 0xdb, 0xd5, 0x74,
	0xbe, 0x57, 0x3a, 0x9e, 0xfe, 0x70, 0x74, 0xb3, 0x4f, 0xdc, 0x5b, 0xac, 0xe9, 0x6c, 0x0c, 0x9e,
	0xa6, 0x95, 0xbf, 0x55, 0xaa, 0x86, 0xfc, 0xef, 0x5b, 0xe0, 0xeb, 0x13, 0xe0, 0x82, 0x56, 0xa2,
	0x10, 0x07, 0x27, 0xe0, 0xd9, 0x6f, 0x76, 0x04, 0x5d, 0x5b, 0xc5, 0x2a, 0xf1, 0x3f, 0xee, 0xed,
	0x99, 0x74, 0x44, 0xc7, 0x96, 0x51, 0x32, 0xbe, 0x83, 0x9e, 0xa2, 0x86, 0xf3, 0x66, 0x29, 0x9e,
	0x2f, 0xfe, 0xb5, 0x2e, 0xd3, 0x6e, 0x79, 0x33, 0xa3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0x86, 0x94, 0xa6, 0x66, 0x02, 0x00, 0x00,
}