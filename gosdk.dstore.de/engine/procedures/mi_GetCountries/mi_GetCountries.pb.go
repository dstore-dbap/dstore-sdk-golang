// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetCountries.proto
// DO NOT EDIT!

/*
Package mi_GetCountries is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetCountries.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetCountries

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
	CountryId      *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	CountryIdNull  bool                        `protobuf:"varint,1001,opt,name=country_id_null,json=countryIdNull" json:"country_id_null,omitempty"`
	RegionId       *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	RegionIdNull   bool                        `protobuf:"varint,1002,opt,name=region_id_null,json=regionIdNull" json:"region_id_null,omitempty"`
	LanguageId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull bool                        `protobuf:"varint,1003,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
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

func (m *Parameters) GetRegionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RegionId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
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
	RowId              int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CountryId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	CountryDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=country_description,json=countryDescription" json:"country_description,omitempty"`
	Region             *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=region" json:"region,omitempty"`
	RegionId           *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	CountryCode        *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=country_code,json=countryCode" json:"country_code,omitempty"`
	SortNo             *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
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

func (m *Response_Row) GetCountryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CountryDescription
	}
	return nil
}

func (m *Response_Row) GetRegion() *dstore_values.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Response_Row) GetRegionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RegionId
	}
	return nil
}

func (m *Response_Row) GetCountryCode() *dstore_values.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetCountries.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetCountries.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetCountries.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_GetCountries.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xed, 0x6a, 0xd4, 0x4c,
	0x14, 0xa6, 0xcd, 0xbb, 0x1f, 0x3d, 0xdb, 0xb7, 0x2d, 0x53, 0x90, 0xb8, 0x8b, 0x22, 0x15, 0x51,
	0x11, 0x66, 0xd5, 0x22, 0xf8, 0xfd, 0xc3, 0x2a, 0x52, 0xb0, 0xa1, 0xe4, 0x87, 0xa0, 0x7f, 0x42,
	0xdc, 0x1c, 0x43, 0x70, 0x77, 0x66, 0x99, 0x49, 0x2c, 0x5e, 0x85, 0xdf, 0x5e, 0x80, 0xb7, 0xe1,
	0x15, 0xa9, 0x37, 0xe1, 0x99, 0xcc, 0xa4, 0xd9, 0x44, 0x70, 0xdb, 0x3f, 0x09, 0x67, 0xce, 0xf3,
	0x9c, 0xe7, 0xcc, 0x73, 0x0e, 0x03, 0x3c, 0xd1, 0xb9, 0x54, 0x38, 0x46, 0x91, 0x66, 0x02, 0xc7,
	0x73, 0x25, 0x27, 0x98, 0x14, 0x0a, 0xf5, 0x78, 0x96, 0x45, 0x4f, 0x31, 0xdf, 0x93, 0x85, 0xc8,
	0x55, 0x86, 0x9a, 0x53, 0x2a, 0x97, 0xec, 0x9c, 0xc5, 0x73, 0x8b, 0xe7, 0x2d, 0xd0, 0x70, 0xdb,
	0x95, 0x7b, 0x1b, 0x4f, 0x8b, 0x8a, 0x33, 0x3c, 0xdb, 0xd4, 0x40, 0xa5, 0xa4, 0x72, 0xa9, 0x51,
	0x33, 0x35, 0x43, 0xad, 0xe3, 0x14, 0x5d, 0xf2, 0x62, 0x3b, 0x99, 0xc7, 0x99, 0x78, 0x2d, 0xd5,
	0x2c, 0xce, 0x33, 0x29, 0x2c, 0x68, 0xe7, 0xc7, 0x2a, 0xc0, 0x61, 0xac, 0x62, 0xca, 0xa2, 0xd2,
	0xec, 0x2e, 0xc0, 0xa4, 0xec, 0xe6, 0x5d, 0x94, 0x25, 0xfe, 0xca, 0x85, 0x95, 0x2b, 0x83, 0x9b,
	0x23, 0x77, 0x49, 0xee, 0xba, 0xca, 0x44, 0x8e, 0x29, 0xaa, 0xe7, 0x26, 0x0a, 0xd7, 0x1c, 0x7c,
	0x3f, 0x61, 0x97, 0x61, 0xb3, 0xe6, 0x46, 0xa2, 0x98, 0x4e, 0xfd, 0x9f, 0x3d, 0xaa, 0xd0, 0x0f,
	0xff, 0x3f, 0x06, 0x05, 0x74, 0xca, 0x6e, 0xc3, 0x9a, 0xc2, 0x94, 0x7a, 0x30, 0x1a, 0xab, 0xcb,
	0x35, 0xfa, 0x16, 0x4d, 0x12, 0x97, 0x60, 0xe3, 0x98, 0x69, 0x15, 0x7e, 0x59, 0x85, 0xf5, 0x0a,
	0x52, 0x0a, 0xdc, 0x87, 0xc1, 0x34, 0x16, 0x69, 0x41, 0x5e, 0x18, 0x09, 0x6f, 0xb9, 0x04, 0x54,
	0x78, 0x12, 0xb9, 0x0a, 0x5b, 0x0b, 0x6c, 0x2b, 0xf3, 0xdb, 0xca, 0x6c, 0xd4, 0x30, 0x23, 0xb4,
	0xf3, 0xbd, 0x03, 0xfd, 0x10, 0xf5, 0x5c, 0x0a, 0x8d, 0xec, 0x3a, 0x74, 0xca, 0xd9, 0x38, 0xdb,
	0x86, 0xbc, 0x39, 0x6b, 0x3b, 0xb7, 0x27, 0xe6, 0x1b, 0x5a, 0x20, 0x7b, 0x01, 0x5b, 0x66, 0x2a,
	0xd1, 0xc2, 0x58, 0xc8, 0x0f, 0x8f, 0xc8, 0xbc, 0x45, 0x6e, 0x0f, 0xef, 0x80, 0xe2, 0xfd, 0x3a,
	0x0e, 0x37, 0x67, 0xcd, 0x03, 0xf2, 0xb8, 0xe7, 0xb6, 0x81, 0xae, 0x6f, 0x2a, 0x9e, 0xff, 0xab,
	0xa2, 0xdd, 0x95, 0x03, 0xfb, 0x0f, 0x2b, 0x38, 0x7b, 0x00, 0x9e, 0x92, 0x47, 0xfe, 0x7f, 0x25,
	0xeb, 0x1a, 0xff, 0xe7, 0xc2, 0xf2, 0xea, 0xf2, 0x3c, 0x94, 0x47, 0xa1, 0xe1, 0x0d, 0xdf, 0x7b,
	0xe0, 0x51, 0xc0, 0xce, 0x40, 0x97, 0x42, 0x63, 0xff, 0x87, 0x80, 0xfc, 0xe8, 0x84, 0x1d, 0x0a,
	0xc9, 0xdd, 0x7b, 0x8d, 0x0d, 0xfb, 0x18, 0x9c, 0x6a, 0xc5, 0x9e, 0xc1, 0x76, 0x45, 0x4e, 0x50,
	0x4f, 0x54, 0x36, 0x2f, 0x3d, 0xfb, 0x14, 0x34, 0x1d, 0x77, 0x55, 0x34, 0xf5, 0x28, 0x52, 0x5b,
	0x84, 0x39, 0xde, 0xe3, 0x9a, 0xc6, 0x76, 0xa9, 0xc5, 0x72, 0x6d, 0xfc, 0xcf, 0xcb, 0x0b, 0x38,
	0x28, 0xbb, 0xb3, 0xb8, 0xbc, 0x5f, 0x82, 0xd3, 0x6c, 0xef, 0x43, 0x58, 0xaf, 0xba, 0x9f, 0xc8,
	0x04, 0xfd, 0xaf, 0xcb, 0x55, 0x07, 0x8e, 0xb0, 0x47, 0x78, 0x76, 0x0b, 0x7a, 0x5a, 0xaa, 0x3c,
	0x12, 0xd2, 0xff, 0x76, 0x02, 0xe1, 0xae, 0x01, 0x07, 0xf2, 0xd1, 0x21, 0x8c, 0x32, 0xd9, 0x9a,
	0x63, 0xfd, 0x50, 0xbd, 0xbc, 0x91, 0x4a, 0x9d, 0xbc, 0xa9, 0xf2, 0xc9, 0x09, 0xde, 0xb2, 0x57,
	0xdd, 0xf2, 0xed, 0xd8, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x1c, 0x22, 0xea, 0xfe, 0x04,
	0x00, 0x00,
}