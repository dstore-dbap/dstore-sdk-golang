// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_UpdateVisitorProperties_Pu.proto
// DO NOT EDIT!

/*
Package mi_UpdateVisitorProperties_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_UpdateVisitorProperties_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_UpdateVisitorProperties_Pu

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
	UniqueId       *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull   bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	CountryId      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	CountryIdNull  bool                        `protobuf:"varint,1002,opt,name=country_id_null,json=countryIdNull" json:"country_id_null,omitempty"`
	LanguageId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull bool                        `protobuf:"varint,1003,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	CurrencyId     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencyIdNull bool                        `protobuf:"varint,1004,opt,name=currency_id_null,json=currencyIdNull" json:"currency_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetCountryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountryId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_UpdateVisitorProperties_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_UpdateVisitorProperties_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_UpdateVisitorProperties_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_UpdateVisitorProperties_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0x69, 0xd7, 0xb4, 0xe9, 0xa9, 0xb6, 0x65, 0x05, 0xa9, 0x09, 0x8a, 0x54, 0x44, 0x45,
	0xd8, 0x88, 0x5e, 0xf8, 0x41, 0x41, 0x10, 0xbc, 0x08, 0xd2, 0x12, 0x06, 0x2c, 0xe8, 0x85, 0xcb,
	0x9a, 0x3d, 0x2e, 0x83, 0xc9, 0x4c, 0x3c, 0x33, 0x63, 0xf1, 0x2d, 0x7c, 0x07, 0x9f, 0xc3, 0x07,
	0xf2, 0xe3, 0x21, 0x3c, 0xb3, 0x33, 0xdb, 0x4d, 0x56, 0x50, 0x7b, 0x93, 0xe4, 0xec, 0xf9, 0xff,
	0xcf, 0x2f, 0x7b, 0x3e, 0xe0, 0xb0, 0x34, 0x56, 0x13, 0x8e, 0x50, 0x55, 0x52, 0xe1, 0x68, 0x41,
	0x7a, 0x8a, 0xa5, 0x23, 0x34, 0xa3, 0xb9, 0xcc, 0x5f, 0x2d, 0xca, 0xc2, 0xe2, 0x89, 0x34, 0x92,
	0x35, 0x13, 0xd2, 0x0b, 0x24, 0x2b, 0xd1, 0xe4, 0x13, 0x97, 0xb1, 0xd0, 0xea, 0xf4, 0x5e, 0x70,
	0x67, 0xc1, 0x9d, 0xfd, 0xd5, 0x32, 0xb8, 0x1c, 0x51, 0x9f, 0x8a, 0x99, 0x43, 0x13, 0x2a, 0x0c,
	0xae, 0xae, 0xf2, 0x91, 0x48, 0x53, 0x4c, 0x0d, 0x57, 0x53, 0x73, 0x34, 0xa6, 0xa8, 0x30, 0x26,
	0x6f, 0x76, 0x93, 0xb6, 0x90, 0xea, 0xbd, 0xa6, 0x79, 0x61, 0xa5, 0x56, 0x41, 0x74, 0xf0, 0x35,
	0x01, 0x98, 0x14, 0x54, 0x70, 0x16, 0xc9, 0xa4, 0x8f, 0x60, 0xcb, 0x29, 0xf9, 0xd1, 0x61, 0x2e,
	0xcb, 0xfd, 0xb5, 0x1b, 0x6b, 0x77, 0xb6, 0x1f, 0x0c, 0xb2, 0xf8, 0x06, 0xf1, 0x4f, 0x19, 0x4b,
	0x52, 0x55, 0x27, 0x3e, 0x10, 0xfd, 0x20, 0x1e, 0x97, 0xe9, 0x2d, 0xd8, 0x39, 0x33, 0xe6, 0xca,
	0xcd, 0x66, 0xfb, 0xdf, 0x37, 0xd9, 0xde, 0x17, 0x17, 0x1b, 0xc9, 0x31, 0x3f, 0x4c, 0x9f, 0x02,
	0x4c, 0xb5, 0x53, 0x96, 0x3e, 0x7b, 0xc0, 0x7a, 0x0d, 0x18, 0x76, 0x00, 0x52, 0x59, 0xac, 0x90,
	0x02, 0x61, 0x2b, 0xca, 0x19, 0x71, 0x1b, 0x76, 0x5b, 0x6f, 0x60, 0xfc, 0x08, 0x8c, 0x4b, 0x67,
	0xa2, 0x1a, 0x72, 0x08, 0xdb, 0xb3, 0x42, 0x55, 0x8e, 0x5b, 0xe1, 0x29, 0xc9, 0xbf, 0x29, 0xd0,
	0xe8, 0x19, 0x73, 0x17, 0xf6, 0x96, 0xdc, 0x81, 0xf3, 0x33, 0x70, 0x76, 0x5a, 0x59, 0x03, 0x9a,
	0x3a, 0x22, 0x54, 0xd3, 0xfa, 0x75, 0x2e, 0xfc, 0x07, 0xa8, 0xd1, 0x07, 0xd0, 0x92, 0x3b, 0x80,
	0x7e, 0x45, 0x50, 0x2b, 0xf3, 0xa0, 0x83, 0x6f, 0xeb, 0xd0, 0x17, 0x68, 0x16, 0x5a, 0x19, 0x4c,
	0xef, 0x43, 0xaf, 0xde, 0x81, 0xee, 0x7c, 0xe2, 0x86, 0x85, 0xfd, 0x78, 0xe1, 0x3f, 0x45, 0x10,
	0xa6, 0xaf, 0x61, 0xcf, 0x4f, 0x3f, 0x5f, 0x1a, 0x3f, 0xf7, 0x3e, 0x61, 0x73, 0xd6, 0x31, 0x77,
	0x97, 0xe4, 0x88, 0xe3, 0x71, 0x1b, 0x8b, 0xdd, 0xf9, 0xea, 0x83, 0xf4, 0x31, 0x6c, 0xc6, 0xad,
	0xe3, 0x3e, 0xfb, 0x8a, 0xd7, 0xff, 0xa8, 0x18, 0x76, 0xf2, 0x28, 0x7c, 0x8b, 0x46, 0x9e, 0xbe,
	0x84, 0x84, 0xf4, 0x29, 0x37, 0xcd, 0xbb, 0x9e, 0x64, 0xe7, 0x38, 0x93, 0xac, 0x69, 0x45, 0x26,
	0xf4, 0xa9, 0xf0, 0x55, 0x06, 0xd7, 0x20, 0xe1, 0xdf, 0xe9, 0x15, 0xd8, 0xe0, 0xc8, 0xcf, 0xe2,
	0xcb, 0x31, 0x37, 0xa7, 0x27, 0x7a, 0x1c, 0x8e, 0xcb, 0xe7, 0x6f, 0x61, 0x28, 0x75, 0x07, 0xd1,
	0xde, 0xf1, 0x9b, 0x67, 0x95, 0x36, 0xe5, 0x87, 0x26, 0x5f, 0x9e, 0xfb, 0xd4, 0xdf, 0x6d, 0xd4,
	0xc7, 0xf4, 0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x3f, 0x89, 0xc8, 0x2b, 0x04, 0x00,
	0x00,
}
