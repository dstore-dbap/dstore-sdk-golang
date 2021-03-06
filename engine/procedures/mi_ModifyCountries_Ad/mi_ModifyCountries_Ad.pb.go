// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyCountries_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyCountries_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyCountries_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyCountries_Ad

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
	CountryId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	CountryIdNull          bool                        `protobuf:"varint,1001,opt,name=country_id_null,json=countryIdNull" json:"country_id_null,omitempty"`
	CountryDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=country_description,json=countryDescription" json:"country_description,omitempty"`
	CountryDescriptionNull bool                        `protobuf:"varint,1002,opt,name=country_description_null,json=countryDescriptionNull" json:"country_description_null,omitempty"`
	CountryCode            *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=country_code,json=countryCode" json:"country_code,omitempty"`
	CountryCodeNull        bool                        `protobuf:"varint,1003,opt,name=country_code_null,json=countryCodeNull" json:"country_code_null,omitempty"`
	SortNo                 *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull             bool                        `protobuf:"varint,1004,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	LanguageId             *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull         bool                        `protobuf:"varint,1005,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	DeleteCountry          *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=delete_country,json=deleteCountry" json:"delete_country,omitempty"`
	DeleteCountryNull      bool                        `protobuf:"varint,1006,opt,name=delete_country_null,json=deleteCountryNull" json:"delete_country_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCountryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountryId
	}
	return nil
}

func (m *Parameters) GetCountryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CountryDescription
	}
	return nil
}

func (m *Parameters) GetCountryCode() *dstore_values.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetDeleteCountry() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCountry
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyCountries_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyCountries_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyCountries_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifyCountries_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x1b, 0x92, 0x94, 0x49, 0xdb, 0xb4, 0x8e, 0x54, 0x2d, 0x89, 0x40, 0xd0, 0x1e, 0x00,
	0x21, 0x6d, 0xf8, 0xa9, 0xc4, 0x8f, 0xe0, 0x40, 0x03, 0x87, 0x08, 0x25, 0x42, 0x7b, 0xa8, 0x04,
	0x97, 0xd5, 0x36, 0x76, 0x57, 0x16, 0x1b, 0x3b, 0xb2, 0x37, 0x54, 0x79, 0x0b, 0x1e, 0x80, 0xd7,
	0xe0, 0xa1, 0xf8, 0x7d, 0x06, 0xec, 0x1d, 0x6f, 0x37, 0xbb, 0x8d, 0x94, 0x5e, 0x92, 0x8c, 0xe7,
	0xfb, 0x99, 0x78, 0x66, 0x0c, 0xc7, 0x54, 0xa7, 0x52, 0xb1, 0x3e, 0x13, 0x31, 0x17, 0xac, 0x3f,
	0x53, 0x72, 0xc2, 0xe8, 0x5c, 0x31, 0xdd, 0x9f, 0xf2, 0x70, 0x24, 0x29, 0x3f, 0x5f, 0x0c, 0xe4,
	0x5c, 0xa4, 0x8a, 0x33, 0x1d, 0xbe, 0xa5, 0xbe, 0x01, 0xa4, 0x92, 0x1c, 0x21, 0xcb, 0x47, 0x96,
	0xbf, 0x12, 0xda, 0xed, 0x38, 0xe9, 0xaf, 0x51, 0x32, 0x67, 0x1a, 0x99, 0xdd, 0x5b, 0x65, 0x3f,
	0xa6, 0x94, 0x54, 0x2e, 0xd5, 0x2b, 0xa7, 0xa6, 0x4c, 0xeb, 0x28, 0x66, 0x2e, 0x79, 0x54, 0x4d,
	0xa6, 0x11, 0x17, 0xe7, 0x52, 0x4d, 0xa3, 0x94, 0x4b, 0x81, 0xa0, 0xc3, 0xef, 0x75, 0x80, 0x8f,
	0x91, 0x8a, 0x4c, 0x96, 0x29, 0x4d, 0x5e, 0x01, 0x4c, 0xb2, 0x82, 0x16, 0x21, 0xa7, 0xde, 0xc6,
	0xdd, 0x8d, 0x07, 0xad, 0xa7, 0x3d, 0xdf, 0x95, 0xee, 0xaa, 0xe2, 0x22, 0x65, 0x31, 0x53, 0xa7,
	0x36, 0x0a, 0x6e, 0x3a, 0xf8, 0x90, 0x92, 0xfb, 0xd0, 0x2e, 0xb8, 0xa1, 0x98, 0x27, 0x89, 0xf7,
	0xb3, 0x69, 0x14, 0xb6, 0x82, 0x9d, 0x4b, 0xd0, 0xd8, 0x9c, 0x92, 0x0f, 0xd0, 0xc9, 0x81, 0x94,
	0xe9, 0x89, 0xe2, 0x33, 0x5b, 0x90, 0xb7, 0x99, 0xb9, 0x75, 0x2b, 0x6e, 0xda, 0x5c, 0x8e, 0x88,
	0xd1, 0x8c, 0x38, 0xda, 0xbb, 0x82, 0x45, 0x5e, 0x82, 0xb7, 0x42, 0x0c, 0xed, 0x7f, 0xa1, 0xfd,
	0xc1, 0x55, 0x5a, 0x56, 0xc7, 0x1b, 0xd8, 0xce, 0xa9, 0x13, 0x49, 0x99, 0x57, 0x5b, 0x5b, 0x40,
	0xcb, 0xe1, 0x07, 0x06, 0x4e, 0x1e, 0xc1, 0xfe, 0x32, 0x1d, 0x2d, 0x7f, 0xa3, 0x65, 0x7b, 0x09,
	0x98, 0x79, 0x1d, 0x43, 0x53, 0x4b, 0x95, 0x86, 0x42, 0x7a, 0x37, 0xd6, 0xdf, 0x6a, 0xc3, 0x62,
	0xc7, 0x92, 0xdc, 0x83, 0x6d, 0xc7, 0x42, 0xf5, 0x3f, 0xa8, 0x0e, 0x98, 0xce, 0x84, 0x5f, 0x43,
	0x2b, 0x89, 0x44, 0x3c, 0x37, 0x7d, 0xb7, 0x2d, 0xab, 0xaf, 0x17, 0x87, 0x1c, 0x6f, 0x7a, 0xf6,
	0x10, 0xf6, 0x96, 0xd8, 0x68, 0xf2, 0x17, 0x4d, 0x76, 0x0b, 0x58, 0x66, 0x74, 0x02, 0xbb, 0x94,
	0x25, 0x66, 0x4c, 0x42, 0xf7, 0xdf, 0xbc, 0xc6, 0x4a, 0xaf, 0x33, 0x29, 0x13, 0x16, 0x09, 0xf4,
	0xda, 0x41, 0x0a, 0x0e, 0xf9, 0x82, 0xf4, 0xa1, 0x53, 0xd6, 0x40, 0xc7, 0x7f, 0xe8, 0xb8, 0x5f,
	0x02, 0x5b, 0xd3, 0xc3, 0x1f, 0x9b, 0xb0, 0x15, 0x30, 0x3d, 0x93, 0x42, 0x33, 0xf2, 0x18, 0xea,
	0xd9, 0xf0, 0xbb, 0xb9, 0xbc, 0x6c, 0x94, 0x5b, 0x29, 0x5c, 0x8c, 0xf7, 0xf6, 0x33, 0x40, 0x20,
	0xf9, 0x04, 0x7b, 0x76, 0xec, 0xc3, 0xa5, 0xb9, 0x37, 0x63, 0x56, 0x33, 0x64, 0xbf, 0x42, 0xae,
	0x6e, 0xc7, 0xc8, 0xc4, 0xc3, 0x22, 0x0e, 0xda, 0xd3, 0xf2, 0x01, 0x79, 0x01, 0x4d, 0xb7, 0x6e,
	0x66, 0x6e, 0xac, 0xe2, 0x9d, 0x2b, 0x8a, 0xb8, 0x8c, 0x23, 0xfc, 0x0e, 0x72, 0x38, 0x19, 0x40,
	0x4d, 0xc9, 0x0b, 0x33, 0x06, 0x96, 0xf5, 0xc4, 0xbf, 0xc6, 0xbb, 0xe0, 0xe7, 0x57, 0xe0, 0x07,
	0xf2, 0x22, 0xb0, 0xec, 0xee, 0x6d, 0xa8, 0x99, 0xdf, 0xe4, 0x00, 0x1a, 0x26, 0xb2, 0x8d, 0xff,
	0x36, 0x36, 0x97, 0x52, 0x0f, 0xea, 0x26, 0x1c, 0xd2, 0x93, 0x53, 0xe8, 0x71, 0x59, 0x91, 0x2e,
	0x1e, 0xaa, 0xcf, 0xcf, 0x63, 0xa9, 0xe9, 0x97, 0x3c, 0x4f, 0xaf, 0xfd, 0x96, 0x9d, 0x35, 0xb2,
	0x57, 0xe3, 0xd9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xf8, 0xa7, 0x24, 0x04, 0x05, 0x00,
	0x00,
}
