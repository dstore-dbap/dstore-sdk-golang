// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_AddBinaryToValues_Ad.proto
// DO NOT EDIT!

/*
Package im_AddBinaryToValues_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_AddBinaryToValues_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_AddBinaryToValues_Ad

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
	BinaryCodeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	BinaryCodeIdNull bool                        `protobuf:"varint,1001,opt,name=binary_code_id_null,json=binaryCodeIdNull" json:"binary_code_id_null,omitempty"`
	ValueIds         *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=value_ids,json=valueIds" json:"value_ids,omitempty"`
	ValueIdsNull     bool                        `protobuf:"varint,1002,opt,name=value_ids_null,json=valueIdsNull" json:"value_ids_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Parameters) GetValueIds() *dstore_values.StringValue {
	if m != nil {
		return m.ValueIds
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_AddBinaryToValues_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_AddBinaryToValues_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_AddBinaryToValues_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_AddBinaryToValues_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0x26, 0xf1, 0xf2, 0x33, 0x2d, 0x64, 0x41, 0x81, 0xe1, 0x39, 0x6c, 0x8c, 0x8c, 0xc0, 0xae,
	0x94, 0xb1, 0xc0, 0x7e, 0x2e, 0x93, 0x91, 0x0b, 0x5f, 0x24, 0x0c, 0x31, 0xc6, 0xb6, 0x1b, 0xe3,
	0x44, 0xaa, 0x11, 0x8d, 0xa5, 0x20, 0x39, 0x0d, 0x7d, 0x8b, 0xbe, 0x4d, 0x1f, 0xa7, 0xd7, 0xed,
	0x53, 0x54, 0x96, 0xe4, 0x24, 0x76, 0x29, 0xb4, 0x37, 0xb6, 0xcf, 0x39, 0xdf, 0xf7, 0x9d, 0xa3,
	0xef, 0xc8, 0xe0, 0x2b, 0x51, 0x99, 0x90, 0x74, 0x4c, 0x79, 0xc2, 0x38, 0x1d, 0x6f, 0xa5, 0x58,
	0x53, 0xb2, 0x93, 0x54, 0x8d, 0x59, 0x1a, 0x4d, 0x09, 0x99, 0x31, 0x1e, 0xcb, 0xcb, 0xdf, 0xe2,
	0x4f, 0xbc, 0xd9, 0x51, 0xa5, 0x33, 0x48, 0x43, 0x32, 0x01, 0x47, 0x96, 0x87, 0x2c, 0x0f, 0x3d,
	0x02, 0x0e, 0xfa, 0x4e, 0xfe, 0xc2, 0x64, 0x2c, 0x37, 0x78, 0x5b, 0xee, 0x49, 0xa5, 0x14, 0xd2,
	0x95, 0x06, 0xe5, 0x52, 0x4a, 0x95, 0x8a, 0x13, 0xea, 0x8a, 0x1f, 0xab, 0xc5, 0x2c, 0x66, 0xfc,
	0x4c, 0xc8, 0x34, 0xce, 0x98, 0xe0, 0x16, 0x34, 0xbc, 0xa9, 0x01, 0xf0, 0x2b, 0x96, 0xb1, 0xae,
	0x52, 0xa9, 0xe0, 0x14, 0x74, 0x57, 0x66, 0xaa, 0x68, 0x2d, 0x08, 0x8d, 0x18, 0xf1, 0x6b, 0x1f,
	0x6a, 0x9f, 0x5e, 0x7d, 0x19, 0x20, 0x77, 0x00, 0x37, 0x19, 0xe3, 0x19, 0x4d, 0xa8, 0x34, 0x93,
	0xe3, 0x8e, 0xa5, 0xfc, 0xd4, 0x8c, 0x90, 0x40, 0x04, 0xfa, 0x65, 0x89, 0x88, 0xef, 0x36, 0x1b,
	0xff, 0xb6, 0xa5, 0x85, 0xda, 0xb8, 0x77, 0x8a, 0x5d, 0xea, 0x02, 0xfc, 0x06, 0x5e, 0x1a, 0x51,
	0x8d, 0x54, 0x7e, 0xdd, 0x74, 0x0b, 0x2a, 0xdd, 0x54, 0x26, 0x19, 0x4f, 0x6c, 0xb3, 0xb6, 0xc9,
	0x85, 0x44, 0xc1, 0x11, 0xe8, 0x1e, 0x88, 0xb6, 0xc7, 0x9d, 0xed, 0xd1, 0x29, 0x20, 0xb9, 0xfe,
	0xf0, 0xba, 0x0e, 0xda, 0x98, 0xaa, 0xad, 0xe0, 0x8a, 0xc2, 0xcf, 0xa0, 0x61, 0xfc, 0x73, 0xc7,
	0x3a, 0x34, 0x72, 0x7b, 0xb1, 0xde, 0xce, 0xf3, 0x27, 0xb6, 0x40, 0xf8, 0x0f, 0xf4, 0x72, 0xe7,
	0xa2, 0x13, 0xeb, 0xf4, 0x94, 0x9e, 0x26, 0xa3, 0x0a, 0xb9, 0x6a, 0xf0, 0x42, 0xc7, 0xe1, 0x31,
	0xc6, 0xaf, 0xd3, 0x72, 0x02, 0x7e, 0x07, 0x2d, 0xb7, 0x31, 0xdf, 0x33, 0x8a, 0xef, 0x1f, 0x28,
	0xda, 0x7d, 0x2e, 0xec, 0x1b, 0x17, 0x70, 0x38, 0x07, 0x9e, 0x14, 0x7b, 0xff, 0x85, 0x61, 0x4d,
	0xd0, 0x93, 0x2e, 0x17, 0x2a, 0x4c, 0x40, 0x58, 0xec, 0x71, 0xce, 0x0f, 0xde, 0x01, 0x4f, 0x7f,
	0xc3, 0x37, 0xa0, 0xa9, 0xa3, 0x7c, 0xd9, 0x57, 0x4b, 0x6d, 0x4b, 0x03, 0x37, 0x74, 0x18, 0x92,
	0xd9, 0x5f, 0x30, 0x60, 0xa2, 0x22, 0x7e, 0xbc, 0xf1, 0xff, 0x7f, 0x24, 0x42, 0x91, 0xf3, 0xa2,
	0x4e, 0x9e, 0xf1, 0x53, 0xac, 0x9a, 0xe6, 0xf2, 0x4d, 0xee, 0x03, 0x00, 0x00, 0xff, 0xff, 0xce,
	0x10, 0x27, 0xef, 0x4f, 0x03, 0x00, 0x00,
}
