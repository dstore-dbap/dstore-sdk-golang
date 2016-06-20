// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyUnitCategoryDescr_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyUnitCategoryDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyUnitCategoryDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyUnitCategoryDescr_Ad

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
	UnitCategoryId                    *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=unit_category_id,json=unitCategoryId" json:"unit_category_id,omitempty"`
	UnitCategoryIdNull                bool                        `protobuf:"varint,1001,opt,name=unit_category_id_null,json=unitCategoryIdNull" json:"unit_category_id_null,omitempty"`
	LanguageId                        *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                    bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	UnitCategoryDescription           *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unit_category_description,json=unitCategoryDescription" json:"unit_category_description,omitempty"`
	UnitCategoryDescriptionNull       bool                        `protobuf:"varint,1003,opt,name=unit_category_description_null,json=unitCategoryDescriptionNull" json:"unit_category_description_null,omitempty"`
	DeleteUnitCategoryDescription     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_unit_category_description,json=deleteUnitCategoryDescription" json:"delete_unit_category_description,omitempty"`
	DeleteUnitCategoryDescriptionNull bool                        `protobuf:"varint,1004,opt,name=delete_unit_category_description_null,json=deleteUnitCategoryDescriptionNull" json:"delete_unit_category_description_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUnitCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitCategoryId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetUnitCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.UnitCategoryDescription
	}
	return nil
}

func (m *Parameters) GetDeleteUnitCategoryDescription() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteUnitCategoryDescription
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyUnitCategoryDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyUnitCategoryDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyUnitCategoryDescr_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifyUnitCategoryDescr_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x53, 0x41,
	0x10, 0xa6, 0x4d, 0xd3, 0x96, 0x29, 0xd4, 0xb0, 0xa2, 0xa6, 0x09, 0x2d, 0xb5, 0x22, 0x28, 0xc2,
	0x89, 0xd4, 0x1b, 0x85, 0x82, 0xa8, 0xed, 0x45, 0x90, 0x14, 0x5d, 0x68, 0x41, 0x2f, 0x3c, 0x6c,
	0xb3, 0xdb, 0xc3, 0xe2, 0xc9, 0x6e, 0xd8, 0x3d, 0xc7, 0xd2, 0xb7, 0xf0, 0x85, 0xbc, 0xf6, 0x59,
	0xfc, 0x79, 0x08, 0x67, 0x33, 0x1b, 0x93, 0x73, 0x34, 0x2d, 0xbd, 0x49, 0x32, 0x99, 0x99, 0xef,
	0xfb, 0x76, 0xf6, 0x9b, 0x85, 0x03, 0xe9, 0x0b, 0xeb, 0x54, 0x4f, 0x99, 0x4c, 0x1b, 0xd5, 0x1b,
	0x3b, 0x3b, 0x54, 0xb2, 0x74, 0xca, 0xf7, 0x46, 0x3a, 0x1d, 0x58, 0xa9, 0xcf, 0x2f, 0x4f, 0x8c,
	0x2e, 0xde, 0x88, 0x42, 0x65, 0xd6, 0x5d, 0x1e, 0x2a, 0x3f, 0x74, 0xe9, 0x2b, 0x99, 0x60, 0x61,
	0x61, 0xd9, 0x13, 0xea, 0x4e, 0xa8, 0x3b, 0xb9, 0xb2, 0xa5, 0x73, 0x3b, 0x52, 0x7d, 0x11, 0x79,
	0xa9, 0x3c, 0x21, 0x74, 0xb6, 0xaa, 0xfc, 0xca, 0x39, 0xeb, 0x62, 0xaa, 0x5b, 0x4d, 0x8d, 0x94,
	0xf7, 0x22, 0x53, 0x31, 0xf9, 0xa0, 0x9e, 0x2c, 0x84, 0x36, 0xe7, 0xd6, 0x8d, 0x44, 0xa1, 0xad,
	0xa1, 0xa2, 0xbd, 0xef, 0x2b, 0x00, 0xef, 0x84, 0x13, 0x98, 0x55, 0xce, 0xb3, 0x23, 0x68, 0x95,
	0x28, 0x2c, 0x1d, 0x46, 0x65, 0xa9, 0x96, 0xed, 0xa5, 0xdd, 0xa5, 0x47, 0x1b, 0xfb, 0xdd, 0x24,
	0x1e, 0x24, 0x6a, 0xd3, 0x06, 0x2b, 0x94, 0x3b, 0x0d, 0x11, 0xdf, 0x2c, 0xe7, 0x4e, 0xd3, 0x97,
	0x6c, 0x1f, 0xee, 0xd4, 0x61, 0x52, 0x53, 0xe6, 0x79, 0xfb, 0xc7, 0x1a, 0x82, 0xad, 0x73, 0x56,
	0xad, 0x3f, 0xc6, 0x14, 0x3b, 0x80, 0x8d, 0x5c, 0x98, 0xac, 0xc4, 0x03, 0x04, 0xd6, 0xe5, 0xeb,
	0x59, 0x61, 0x5a, 0x8f, 0x8c, 0x8f, 0xa1, 0x35, 0xd7, 0x4d, 0x64, 0x3f, 0x89, 0x6c, 0x73, 0x56,
	0x36, 0x21, 0x3a, 0x85, 0xad, 0xaa, 0x38, 0x19, 0xc6, 0xaf, 0xc7, 0x61, 0x2a, 0xed, 0xc6, 0x84,
	0xb6, 0x53, 0xa3, 0xf5, 0x85, 0xd3, 0x26, 0x23, 0xd6, 0x7b, 0x65, 0xfd, 0xe6, 0xa8, 0x95, 0x1d,
	0xc2, 0xce, 0x42, 0x5c, 0x12, 0xf4, 0x8b, 0x04, 0x75, 0x17, 0x20, 0x4c, 0xd4, 0x49, 0xd8, 0x95,
	0x2a, 0xc7, 0xdb, 0x48, 0x17, 0x8b, 0x5c, 0xf9, 0xef, 0x6c, 0xce, 0xac, 0xcd, 0x95, 0x30, 0xa4,
	0x72, 0x9b, 0x40, 0x4e, 0x16, 0x68, 0x7d, 0x0f, 0x0f, 0xaf, 0x63, 0x21, 0xc9, 0xbf, 0x49, 0xf2,
	0xfd, 0x2b, 0xe1, 0x82, 0xf0, 0xbd, 0x6f, 0xcb, 0xb0, 0xce, 0x95, 0x1f, 0x5b, 0xe3, 0x15, 0x7b,
	0x0a, 0xcd, 0x89, 0x4f, 0xa3, 0x79, 0xfe, 0xce, 0x33, 0x6e, 0x01, 0x79, 0xf8, 0x28, 0x7c, 0x72,
	0x2a, 0x64, 0x1f, 0xa0, 0x15, 0x1c, 0x9a, 0xce, 0x59, 0x14, 0x3d, 0xd0, 0xc0, 0xe6, 0xa4, 0xd6,
	0x5c, 0x37, 0xf2, 0x00, 0xe3, 0xfe, 0x2c, 0xe6, 0xb7, 0x46, 0xd5, 0x3f, 0xd8, 0x73, 0x58, 0x8b,
	0x9b, 0x81, 0xd7, 0x1b, 0x10, 0x77, 0xfe, 0x41, 0xa4, 0xbd, 0x19, 0xd0, 0x37, 0x9f, 0x96, 0xb3,
	0xb7, 0xd0, 0x70, 0xf6, 0x02, 0xe7, 0x1d, 0xba, 0x5e, 0x24, 0x37, 0x58, 0xe5, 0x64, 0x3a, 0x8a,
	0x84, 0xdb, 0x0b, 0x1e, 0x50, 0x3a, 0xdb, 0xd0, 0xc0, 0xdf, 0xec, 0x2e, 0xac, 0x62, 0x14, 0x2c,
	0xfe, 0xf5, 0x18, 0x87, 0xd3, 0xe4, 0x4d, 0x0c, 0xfb, 0xf2, 0xf5, 0x27, 0xe8, 0x6a, 0x5b, 0xa3,
	0x98, 0xbd, 0x35, 0x1f, 0x5f, 0x66, 0xd6, 0xcb, 0xcf, 0xd3, 0xbc, 0xbc, 0xf1, 0x73, 0x74, 0xb6,
	0x3a, 0x59, 0xf8, 0x67, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0x1b, 0x5e, 0xa7, 0xcf, 0x04,
	0x00, 0x00,
}