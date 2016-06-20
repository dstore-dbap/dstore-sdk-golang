// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyValueDetails_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyValueDetails_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyValueDetails_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyValueDetails_Ad

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
	ValueId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	ValueIdNull bool                        `protobuf:"varint,1001,opt,name=value_id_null,json=valueIdNull" json:"value_id_null,omitempty"`
	SortNo      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull  bool                        `protobuf:"varint,1002,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyValueDetails_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyValueDetails_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyValueDetails_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_ModifyValueDetails_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x6b, 0xdb, 0x30,
	0x14, 0x27, 0xf1, 0x12, 0x87, 0x97, 0x8d, 0x0d, 0x0d, 0x86, 0xe7, 0xb0, 0xb1, 0x25, 0x30, 0x76,
	0x52, 0xc6, 0x16, 0xb6, 0xd1, 0x5b, 0x4b, 0x5b, 0xc8, 0x21, 0xa1, 0xe8, 0x50, 0x48, 0x2f, 0xc6,
	0xad, 0x14, 0x23, 0x6a, 0x4b, 0x41, 0x72, 0x1a, 0xfa, 0x2d, 0xfa, 0x71, 0x7a, 0xee, 0xb7, 0x69,
	0x3f, 0x45, 0x65, 0x4b, 0x26, 0xb5, 0x4b, 0x69, 0x7b, 0x49, 0xfc, 0xde, 0xef, 0xcf, 0x7b, 0xfc,
	0xf4, 0xe0, 0x1f, 0xd5, 0xb9, 0x54, 0x6c, 0xcc, 0x44, 0xc2, 0x05, 0x1b, 0xaf, 0x94, 0x3c, 0x63,
	0x74, 0xad, 0x98, 0x1e, 0xf3, 0x2c, 0x9a, 0x49, 0xca, 0x97, 0x97, 0xc7, 0x71, 0xba, 0x66, 0xfb,
	0x2c, 0x8f, 0x79, 0xaa, 0xa3, 0x5d, 0x8a, 0x0d, 0x27, 0x97, 0xe8, 0x87, 0x15, 0x62, 0x2b, 0xc4,
	0x4f, 0xb1, 0xc3, 0x8f, 0x6e, 0xc0, 0x45, 0xd1, 0xd7, 0x56, 0x1c, 0x7e, 0xae, 0x4f, 0x65, 0x4a,
	0x49, 0xe5, 0xa0, 0x41, 0x1d, 0xca, 0x98, 0xd6, 0x71, 0xc2, 0x1c, 0x38, 0x6a, 0x82, 0x66, 0x8c,
	0x58, 0x4a, 0x95, 0xc5, 0x39, 0x97, 0xc2, 0x92, 0x86, 0x37, 0x2d, 0x80, 0xa3, 0x58, 0xc5, 0x06,
	0x65, 0x4a, 0xa3, 0xbf, 0xd0, 0x2b, 0x67, 0x47, 0x9c, 0x06, 0xad, 0x6f, 0xad, 0x9f, 0xfd, 0xdf,
	0x03, 0xec, 0x76, 0x77, 0x3b, 0x71, 0x91, 0xb3, 0x84, 0xa9, 0x72, 0x73, 0xe2, 0x97, 0xcd, 0x29,
	0x45, 0x23, 0x78, 0x57, 0xe9, 0x22, 0xb1, 0x4e, 0xd3, 0xe0, 0xd6, 0x37, 0xea, 0x1e, 0xe9, 0x3b,
	0xc2, 0xdc, 0xf4, 0xd0, 0x04, 0x7c, 0x2d, 0x55, 0x1e, 0x09, 0x19, 0xb4, 0x9f, 0xf7, 0xee, 0x16,
	0xdc, 0xb9, 0x44, 0xdf, 0xe1, 0xad, 0x53, 0x59, 0xe7, 0x3b, 0xeb, 0x0c, 0x16, 0x2e, 0x8c, 0x87,
	0xd7, 0x6d, 0xe8, 0x11, 0xa6, 0x57, 0x52, 0x68, 0x86, 0x7e, 0x41, 0xa7, 0x8c, 0xc8, 0xed, 0x1f,
	0xe2, 0x7a, 0xf6, 0x36, 0xbe, 0x83, 0xe2, 0x97, 0x58, 0x22, 0x5a, 0xc0, 0x87, 0x22, 0x9c, 0xe8,
	0x41, 0x3a, 0x66, 0x41, 0xcf, 0x88, 0x71, 0x43, 0xdc, 0xcc, 0x70, 0x66, 0xea, 0xe9, 0xb6, 0x26,
	0xef, 0xb3, 0x7a, 0x03, 0xfd, 0x07, 0xdf, 0x3d, 0x4a, 0xe0, 0x95, 0x8e, 0x5f, 0x1f, 0x39, 0xda,
	0x27, 0x9b, 0xd9, 0x7f, 0x52, 0xd1, 0xd1, 0x21, 0x78, 0x4a, 0x6e, 0x82, 0x37, 0xa5, 0x6a, 0x82,
	0x5f, 0x76, 0x40, 0xb8, 0x4a, 0x01, 0x13, 0xb9, 0x21, 0x85, 0x41, 0xf8, 0x05, 0x3c, 0xf3, 0x8d,
	0x3e, 0x41, 0xd7, 0x54, 0xc5, 0xb3, 0x5e, 0xcd, 0x4d, 0x2e, 0x1d, 0xd2, 0x31, 0xe5, 0x94, 0xee,
	0x2d, 0x60, 0xc0, 0x65, 0xc3, 0x7d, 0x7b, 0xd7, 0x27, 0x3b, 0x89, 0xd4, 0xf4, 0xbc, 0xc2, 0xe9,
	0x6b, 0x4e, 0xff, 0xb4, 0x5b, 0x5e, 0xd8, 0x9f, 0xfb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xab,
	0xf4, 0x1f, 0x36, 0x03, 0x00, 0x00,
}