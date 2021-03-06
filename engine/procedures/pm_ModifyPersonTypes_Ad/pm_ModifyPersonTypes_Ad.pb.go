// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonTypes_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonTypes_Ad

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
	PersonTypeId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull          bool                        `protobuf:"varint,1001,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	PersonTypeDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_type_description,json=personTypeDescription" json:"person_type_description,omitempty"`
	PersonTypeDescriptionNull bool                        `protobuf:"varint,1002,opt,name=person_type_description_null,json=personTypeDescriptionNull" json:"person_type_description_null,omitempty"`
	CountryId                 *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	CountryIdNull             bool                        `protobuf:"varint,1003,opt,name=country_id_null,json=countryIdNull" json:"country_id_null,omitempty"`
	SortNo                    *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull                bool                        `protobuf:"varint,1004,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	DeletePersonType          *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_person_type,json=deletePersonType" json:"delete_person_type,omitempty"`
	DeletePersonTypeNull      bool                        `protobuf:"varint,1005,opt,name=delete_person_type_null,json=deletePersonTypeNull" json:"delete_person_type_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PersonTypeDescription
	}
	return nil
}

func (m *Parameters) GetCountryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountryId
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetDeletePersonType() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeletePersonType
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyPersonTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x9b, 0x26, 0x29, 0x43, 0xa1, 0xd1, 0x16, 0x68, 0x9a, 0x00, 0x82, 0x22, 0x04, 0x27,
	0x07, 0x51, 0x54, 0x01, 0x27, 0x8a, 0xe8, 0x21, 0x87, 0x44, 0xd5, 0x0a, 0x21, 0xe0, 0x62, 0xb9,
	0xd9, 0xad, 0x65, 0xe1, 0xec, 0x58, 0xbb, 0x0e, 0x55, 0xde, 0x82, 0xb7, 0xe1, 0x4d, 0x78, 0x07,
	0xfe, 0xde, 0x81, 0x5d, 0xcf, 0x26, 0x4e, 0xdc, 0x56, 0xa5, 0x97, 0x24, 0x9b, 0xf9, 0x7e, 0xc6,
	0xb3, 0xdf, 0x18, 0xf6, 0x85, 0xc9, 0x51, 0xcb, 0x9e, 0x54, 0x71, 0xa2, 0x64, 0x2f, 0xd3, 0x38,
	0x92, 0x62, 0xa2, 0xa5, 0xe9, 0x65, 0xe3, 0x70, 0x80, 0x22, 0x39, 0x99, 0x1e, 0x49, 0x6d, 0x50,
	0xbd, 0x9f, 0x66, 0xd2, 0x84, 0x07, 0x22, 0xb0, 0x90, 0x1c, 0xd9, 0x63, 0xe2, 0x05, 0xc4, 0x0b,
	0x2e, 0x00, 0x77, 0xb6, 0xbc, 0xfc, 0xd7, 0x28, 0x9d, 0x48, 0x43, 0xdc, 0xce, 0xce, 0xb2, 0xa7,
	0xd4, 0x1a, 0xb5, 0x2f, 0x75, 0x97, 0x4b, 0x63, 0x69, 0x4c, 0x14, 0x4b, 0x5f, 0x7c, 0x54, 0x2d,
	0xe6, 0x51, 0xa2, 0x4e, 0x50, 0x8f, 0xa3, 0x3c, 0x41, 0x45, 0xa0, 0xdd, 0x1f, 0x6b, 0x00, 0x47,
	0x91, 0x8e, 0x6c, 0xd5, 0xb6, 0xc2, 0x0e, 0xe0, 0x66, 0x56, 0xb4, 0x14, 0xe6, 0xb6, 0xa7, 0x30,
	0x11, 0xed, 0x95, 0x07, 0x2b, 0x4f, 0xaf, 0x3f, 0xef, 0x06, 0xfe, 0x01, 0x7c, 0x67, 0x89, 0xca,
	0x65, 0x2c, 0xf5, 0x07, 0x77, 0xe2, 0x1b, 0xd9, 0xfc, 0x29, 0xfa, 0x82, 0x05, 0xb0, 0xb5, 0x2c,
	0x11, 0xaa, 0x49, 0x9a, 0xb6, 0x7f, 0x36, 0xad, 0xd0, 0x3a, 0x6f, 0x2d, 0x62, 0x87, 0xb6, 0xc0,
	0x38, 0x6c, 0x2f, 0xe2, 0x85, 0x34, 0x23, 0x9d, 0x64, 0xae, 0xc5, 0xf6, 0x6a, 0xe1, 0xdd, 0xa9,
	0x78, 0x9b, 0x5c, 0x27, 0x2a, 0x26, 0xeb, 0xdb, 0xa5, 0xdc, 0xbb, 0x92, 0xc8, 0xde, 0xc0, 0xdd,
	0x0b, 0x34, 0xa9, 0x99, 0x5f, 0xd4, 0xcc, 0xce, 0xb9, 0xec, 0xa2, 0xab, 0xd7, 0x00, 0x23, 0x9c,
	0xa8, 0x5c, 0x4f, 0xdd, 0x10, 0x6a, 0x97, 0x0f, 0xe1, 0x9a, 0x87, 0xdb, 0x09, 0x3c, 0x81, 0xcd,
	0x92, 0x4b, 0x86, 0xbf, 0xc9, 0xf0, 0xc6, 0x1c, 0x54, 0x98, 0xbc, 0x80, 0xa6, 0x41, 0x9d, 0x87,
	0x0a, 0xdb, 0x6b, 0x97, 0x3b, 0x34, 0x1c, 0x76, 0x88, 0xec, 0x21, 0x6c, 0x78, 0x16, 0x69, 0xff,
	0x21, 0x6d, 0xa0, 0x72, 0x21, 0xdc, 0x07, 0x26, 0x64, 0x6a, 0xaf, 0x34, 0x5c, 0x18, 0x43, 0xbb,
	0x7e, 0xae, 0xc7, 0x31, 0x62, 0x2a, 0x23, 0x45, 0x1e, 0x2d, 0xa2, 0x95, 0xb1, 0x64, 0xfb, 0xb0,
	0x7d, 0x56, 0x8a, 0x8c, 0xff, 0x92, 0xf1, 0xad, 0x2a, 0xc7, 0xb5, 0xb0, 0xfb, 0x7d, 0x15, 0xd6,
	0xb9, 0x34, 0x19, 0x2a, 0x23, 0xd9, 0x33, 0xa8, 0x17, 0xb1, 0xf5, 0x69, 0x9a, 0xdf, 0xa8, 0x5f,
	0x07, 0x8a, 0xf4, 0xa1, 0xfb, 0xe4, 0x04, 0x64, 0x9f, 0xa0, 0xe5, 0x02, 0x1b, 0x2e, 0x24, 0xd6,
	0xc6, 0xa1, 0x66, 0xc9, 0x41, 0x85, 0x5c, 0xcd, 0xf5, 0xc0, 0x9e, 0xfb, 0xe5, 0x99, 0x6f, 0x8e,
	0x97, 0xff, 0x60, 0x2f, 0xa1, 0xe9, 0x17, 0xc5, 0xde, 0xab, 0x53, 0xbc, 0x7f, 0x46, 0x91, 0xd6,
	0x68, 0x40, 0xdf, 0x7c, 0x06, 0x67, 0x87, 0x50, 0xd3, 0x78, 0x6a, 0xef, 0xca, 0xb1, 0xf6, 0x82,
	0xff, 0xda, 0xe9, 0x60, 0x36, 0x84, 0x80, 0xe3, 0x29, 0x77, 0xfc, 0xce, 0x3d, 0xa8, 0xd9, 0xdf,
	0xec, 0x0e, 0x34, 0xec, 0xc9, 0xc5, 0xeb, 0xdb, 0xd0, 0x8e, 0xa5, 0xce, 0xeb, 0xf6, 0xd8, 0x17,
	0x6f, 0x3f, 0x42, 0x37, 0xc1, 0xaa, 0xf8, 0xfc, 0x45, 0xf3, 0xf9, 0x55, 0x8c, 0x46, 0x7c, 0x99,
	0xd5, 0xc5, 0x15, 0xde, 0x45, 0xc7, 0x8d, 0x62, 0xe7, 0xf7, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0xfd, 0x64, 0xde, 0xc6, 0x04, 0x00, 0x00,
}
