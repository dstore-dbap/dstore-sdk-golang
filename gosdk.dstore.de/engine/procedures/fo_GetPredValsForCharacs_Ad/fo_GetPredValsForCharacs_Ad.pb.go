// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetPredValsForCharacs_Ad.proto
// DO NOT EDIT!

/*
Package fo_GetPredValsForCharacs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetPredValsForCharacs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetPredValsForCharacs_Ad

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
	PostingCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=posting_characteristic_id,json=postingCharacteristicId" json:"posting_characteristic_id,omitempty"`
	PostingCharacteristicIdNull bool                        `protobuf:"varint,1001,opt,name=posting_characteristic_id_null,json=postingCharacteristicIdNull" json:"posting_characteristic_id_null,omitempty"`
	DateFormat                  *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=date_format,json=dateFormat" json:"date_format,omitempty"`
	DateFormatNull              bool                        `protobuf:"varint,1002,opt,name=date_format_null,json=dateFormatNull" json:"date_format_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPostingCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingCharacteristicId
	}
	return nil
}

func (m *Parameters) GetDateFormat() *dstore_values.StringValue {
	if m != nil {
		return m.DateFormat
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
	RowId                 int32                      `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ValueInInternalFormat *dstore_values.StringValue `protobuf:"bytes,10001,opt,name=value_in_internal_format,json=valueInInternalFormat" json:"value_in_internal_format,omitempty"`
	Value                 *dstore_values.StringValue `protobuf:"bytes,10002,opt,name=value" json:"value,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValueInInternalFormat() *dstore_values.StringValue {
	if m != nil {
		return m.ValueInInternalFormat
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetPredValsForCharacs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetPredValsForCharacs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetPredValsForCharacs_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_GetPredValsForCharacs_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xd6, 0x36, 0x6c, 0x5b, 0x4d, 0x25, 0xa8, 0x8c, 0x80, 0x74, 0x57, 0xaa, 0x50, 0xb9, 0xd0,
	0x4b, 0x96, 0x9f, 0x03, 0x08, 0xc4, 0x01, 0x0a, 0x45, 0x39, 0x74, 0x55, 0x59, 0xa8, 0x08, 0x84,
	0x14, 0x99, 0xf5, 0x34, 0x44, 0x64, 0xed, 0x95, 0xed, 0xa5, 0xaf, 0xc1, 0xcf, 0x3b, 0xf0, 0x26,
	0x3c, 0x0c, 0x3c, 0x05, 0x13, 0xdb, 0xcb, 0x36, 0x41, 0xa5, 0xf4, 0x92, 0x68, 0x3c, 0xdf, 0x37,
	0xdf, 0x78, 0xbe, 0x31, 0x3c, 0x92, 0xd6, 0x69, 0x83, 0x23, 0x54, 0x65, 0xa5, 0x70, 0x34, 0x33,
	0x7a, 0x82, 0x72, 0x6e, 0xd0, 0x8e, 0x8e, 0x75, 0xf1, 0x12, 0xdd, 0xa1, 0x41, 0x79, 0x24, 0x6a,
	0xbb, 0xaf, 0xcd, 0xde, 0x07, 0x61, 0xc4, 0xc4, 0x16, 0x4f, 0x65, 0x46, 0x30, 0xa7, 0xd9, 0x6e,
	0xe0, 0x66, 0x81, 0x9b, 0xfd, 0x83, 0x30, 0xb8, 0x1a, 0x65, 0x3e, 0x89, 0x7a, 0x8e, 0x36, 0xf0,
	0x07, 0x5b, 0x6d, 0x6d, 0x34, 0x46, 0x9b, 0x98, 0x1a, 0xb6, 0x53, 0x53, 0xb4, 0x56, 0x94, 0x18,
	0x93, 0xb7, 0xba, 0x49, 0x27, 0x2a, 0x75, 0xac, 0xcd, 0x54, 0xb8, 0x4a, 0xab, 0x00, 0xda, 0xf9,
	0xb6, 0x02, 0x70, 0x48, 0xfa, 0x94, 0x45, 0x63, 0xd9, 0x6b, 0xd8, 0x9a, 0x69, 0xeb, 0x2a, 0x55,
	0x16, 0x13, 0xdf, 0x16, 0x9d, 0x56, 0x14, 0x4f, 0x8a, 0x4a, 0xa6, 0xbd, 0x9b, 0xbd, 0xdb, 0x1b,
	0xf7, 0x86, 0x59, 0xbc, 0x4f, 0x6c, 0xb2, 0x52, 0x0e, 0x4b, 0x34, 0x47, 0x4d, 0xc4, 0x6f, 0x44,
	0xf6, 0x5e, 0x8b, 0x9c, 0x4b, 0xf6, 0x1c, 0xb6, 0xcf, 0x2c, 0x5c, 0xa8, 0x79, 0x5d, 0xa7, 0x3f,
	0xd7, 0xa8, 0xfc, 0x3a, 0x1f, 0x9e, 0x51, 0x61, 0x4c, 0x18, 0xf6, 0x18, 0x36, 0xa4, 0x70, 0x58,
	0x84, 0x5b, 0xa4, 0x2b, 0xbe, 0xa1, 0x41, 0xa7, 0x21, 0xeb, 0x0c, 0xf1, 0x43, 0x3f, 0xd0, 0xc0,
	0xf7, 0x3d, 0x9a, 0xed, 0xc2, 0xe6, 0x29, 0x72, 0x10, 0xfd, 0x15, 0x44, 0x2f, 0x2f, 0x61, 0x8d,
	0xce, 0xce, 0x8f, 0x04, 0xd6, 0x39, 0xda, 0x99, 0x56, 0x16, 0xd9, 0x1d, 0xe8, 0xfb, 0x99, 0xc7,
	0xfb, 0xff, 0x91, 0x8b, 0x7e, 0x06, 0x3f, 0x5e, 0x34, 0x5f, 0x1e, 0x80, 0xec, 0x0d, 0x6c, 0x36,
	0xd3, 0x2e, 0x4e, 0x8d, 0x9b, 0x7a, 0x4d, 0x88, 0x9c, 0x75, 0xc8, 0x5d, 0x53, 0x0e, 0x28, 0xce,
	0x97, 0x31, 0xbf, 0x32, 0x6d, 0x1f, 0xb0, 0x87, 0xb0, 0x16, 0x5d, 0x4e, 0x13, 0x5f, 0x71, 0xfb,
	0xaf, 0x8a, 0x61, 0x07, 0x0e, 0xc2, 0x9f, 0x2f, 0xe0, 0x2c, 0x87, 0xc4, 0xe8, 0x93, 0xf4, 0x92,
	0x67, 0x3d, 0xc8, 0xfe, 0x7b, 0x29, 0xb3, 0xc5, 0x20, 0x32, 0xae, 0x4f, 0x78, 0x53, 0x63, 0xf0,
	0xbd, 0x07, 0x09, 0x05, 0xec, 0x3a, 0xac, 0x52, 0xd8, 0xac, 0xc6, 0xe7, 0x31, 0xcd, 0xa6, 0xcf,
	0xfb, 0x14, 0x92, 0xd9, 0xaf, 0x20, 0xf5, 0x5e, 0xd0, 0x00, 0x8a, 0x66, 0x3d, 0x8c, 0x12, 0xf5,
	0xc2, 0xb3, 0x2f, 0xe3, 0x73, 0x4d, 0xbb, 0xe6, 0xcf, 0x72, 0x95, 0x47, 0x6a, 0xf4, 0xef, 0x2e,
	0xf4, 0x7d, 0x22, 0xfd, 0x7a, 0x7e, 0x89, 0x80, 0x7c, 0xf6, 0x0e, 0x86, 0x95, 0xee, 0x5c, 0x75,
	0xf9, 0x76, 0xdf, 0x3e, 0x29, 0xb5, 0x95, 0x1f, 0x17, 0x79, 0x79, 0xc1, 0xe7, 0xfd, 0x7e, 0xd5,
	0x3f, 0xa1, 0xfb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x36, 0xaf, 0xc6, 0x1d, 0x04, 0x00,
	0x00,
}