// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetVisitorProperties_Pu.proto
// DO NOT EDIT!

/*
Package mi_GetVisitorProperties_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetVisitorProperties_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetVisitorProperties_Pu

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
	UniqueId     *dstore_values.StringValue `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull bool                       `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
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
	RowId      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CountryId  *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	LanguageId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	CurrencyId *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCountryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountryId
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Response_Row) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetVisitorProperties_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetVisitorProperties_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetVisitorProperties_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetVisitorProperties_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xa9, 0xb1, 0xbb, 0xdd, 0xb3, 0xa2, 0x32, 0x82, 0xc4, 0x16, 0x64, 0x59, 0x11, 0xf6,
	0x6a, 0x2a, 0x0a, 0xfe, 0x41, 0xbd, 0x11, 0x44, 0x7b, 0xb1, 0xa5, 0xcc, 0xc5, 0x82, 0x7a, 0x11,
	0x62, 0x72, 0x0c, 0x83, 0xc9, 0x4c, 0x3d, 0x33, 0xe3, 0xe2, 0x5b, 0xf8, 0xe7, 0xe5, 0x7c, 0x05,
	0x5f, 0xc0, 0x6b, 0x27, 0x99, 0x89, 0xb5, 0x11, 0xa1, 0xde, 0xb4, 0x9c, 0x7c, 0xdf, 0x6f, 0xce,
	0x9c, 0xef, 0x24, 0xf0, 0xa8, 0x34, 0x56, 0x13, 0xce, 0x51, 0x55, 0x52, 0xe1, 0x7c, 0x4d, 0xba,
	0xc0, 0xd2, 0x11, 0x9a, 0x79, 0x23, 0xb3, 0x17, 0x68, 0xcf, 0xa4, 0x91, 0xde, 0xb0, 0x22, 0xbd,
	0x46, 0xb2, 0x12, 0x4d, 0xb6, 0x72, 0xdc, 0xbb, 0xac, 0x66, 0x27, 0x01, 0xe5, 0x01, 0xe5, 0xff,
	0xf6, 0x4f, 0xaf, 0xc5, 0x26, 0x1f, 0xf3, 0xda, 0xa1, 0x09, 0xf8, 0xf4, 0xc6, 0x76, 0x67, 0x24,
	0xd2, 0x14, 0xa5, 0xd9, 0xb6, 0xd4, 0xa0, 0x31, 0x79, 0x85, 0x51, 0xbc, 0x35, 0x14, 0x6d, 0x2e,
	0xd5, 0x3b, 0x4d, 0x4d, 0x6e, 0xa5, 0x56, 0xc1, 0x74, 0x5c, 0x03, 0xac, 0x72, 0xca, 0xbd, 0x88,
	0x64, 0xd8, 0x03, 0x38, 0x70, 0x4a, 0x7e, 0x70, 0x98, 0xc9, 0x32, 0x1d, 0x1d, 0x8d, 0x4e, 0x0e,
	0xef, 0x4e, 0x79, 0xbc, 0x7d, 0xbc, 0x93, 0xb1, 0x24, 0x55, 0x75, 0xd6, 0x16, 0x62, 0x12, 0xcc,
	0x8b, 0x92, 0xdd, 0x86, 0xcb, 0xbf, 0xc1, 0x4c, 0xb9, 0xba, 0x4e, 0x7f, 0xec, 0x7b, 0x7c, 0x22,
	0x2e, 0xf5, 0x96, 0xa5, 0x7f, 0x78, 0xfc, 0x33, 0x81, 0x89, 0x40, 0xb3, 0xd6, 0xca, 0x20, 0xbb,
	0x03, 0xe3, 0x6e, 0x96, 0x61, 0xa3, 0x18, 0x53, 0x98, 0xf3, 0x79, 0xfb, 0x2b, 0x82, 0x91, 0xbd,
	0x82, 0xab, 0xed, 0x14, 0xd9, 0x1f, 0x63, 0xa4, 0x17, 0x8e, 0x12, 0x0f, 0xf3, 0x01, 0x3c, 0x1c,
	0xf6, 0xd4, 0xd7, 0x8b, 0x4d, 0x2d, 0xae, 0x34, 0xdb, 0x0f, 0xd8, 0x43, 0xd8, 0x8f, 0xe9, 0xa5,
	0x49, 0x77, 0xe2, 0xcd, 0xbf, 0x4e, 0x0c, 0xd9, 0x9e, 0x86, 0x7f, 0xd1, 0xdb, 0xd9, 0x4b, 0x48,
	0x48, 0x9f, 0xa7, 0x17, 0x3b, 0xea, 0x3e, 0xdf, 0x75, 0xd7, 0xbc, 0xcf, 0x81, 0x0b, 0x7d, 0x2e,
	0xda, 0x23, 0xa6, 0xdf, 0x47, 0x90, 0xf8, 0x82, 0x5d, 0x87, 0x3d, 0x5f, 0xb6, 0x2b, 0xf8, 0xbc,
	0xf4, 0xd1, 0x8c, 0xc5, 0xd8, 0x97, 0x3e, 0xe4, 0xc7, 0x00, 0x85, 0x76, 0xca, 0xd2, 0xa7, 0x56,
	0xfb, 0xb2, 0xec, 0x62, 0x9b, 0x0d, 0xf6, 0x23, 0x95, 0xc5, 0x0a, 0x29, 0x2c, 0xe8, 0x20, 0xfa,
	0x3d, 0xfc, 0x14, 0x0e, 0xeb, 0x5c, 0x55, 0xce, 0x5f, 0xb9, 0xa5, 0xbf, 0xee, 0x40, 0x43, 0x0f,
	0x04, 0xbc, 0x70, 0x44, 0xa8, 0x8a, 0xae, 0xf9, 0xb7, 0x5d, 0xf0, 0x1e, 0x58, 0x94, 0xcf, 0xde,
	0xc0, 0x4c, 0xea, 0x41, 0x36, 0x9b, 0x4f, 0xe8, 0xf5, 0x93, 0x4a, 0x9b, 0xf2, 0x7d, 0xaf, 0x97,
	0xff, 0xf7, 0x95, 0xbd, 0xdd, 0xeb, 0x5e, 0xe5, 0x7b, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x14, 0xdd, 0x3c, 0xa3, 0x03, 0x00, 0x00,
}
