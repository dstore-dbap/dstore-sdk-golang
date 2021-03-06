// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetForumCategories_Pu.proto
// DO NOT EDIT!

/*
Package fo_GetForumCategories_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetForumCategories_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetForumCategories_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	ForumCategoryId                *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=forum_category_id,json=forumCategoryId" json:"forum_category_id,omitempty"`
	ForumCategoryIdNull            bool                        `protobuf:"varint,1004,opt,name=forum_category_id_null,json=forumCategoryIdNull" json:"forum_category_id_null,omitempty"`
	DoNotCheckAccessRights         *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=do_not_check_access_rights,json=doNotCheckAccessRights" json:"do_not_check_access_rights,omitempty"`
	DoNotCheckAccessRightsNull     bool                        `protobuf:"varint,1005,opt,name=do_not_check_access_rights_null,json=doNotCheckAccessRightsNull" json:"do_not_check_access_rights_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1006,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetForumCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumCategoryId
	}
	return nil
}

func (m *Parameters) GetDoNotCheckAccessRights() *dstore_values.BooleanValue {
	if m != nil {
		return m.DoNotCheckAccessRights
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
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
	RowId                      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ForumCategoryId            *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=forum_category_id,json=forumCategoryId" json:"forum_category_id,omitempty"`
	CategoryDescription        *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	SortNo                     *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	ContainsCategoriesOrForums *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=contains_categories_or_forums,json=containsCategoriesOrForums" json:"contains_categories_or_forums,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetForumCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumCategoryId
	}
	return nil
}

func (m *Response_Row) GetCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CategoryDescription
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetContainsCategoriesOrForums() *dstore_values.IntegerValue {
	if m != nil {
		return m.ContainsCategoriesOrForums
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetForumCategories_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetForumCategories_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetForumCategories_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_GetForumCategories_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x4e, 0x13, 0x4d,
	0x14, 0x0e, 0x94, 0x16, 0xde, 0x79, 0x09, 0xe0, 0x94, 0xe0, 0xba, 0x20, 0x2a, 0x46, 0xe3, 0xaf,
	0xc5, 0x28, 0x06, 0xe3, 0x1f, 0x83, 0x20, 0x58, 0x13, 0x2a, 0x6e, 0x8c, 0x06, 0x63, 0xb2, 0x59,
	0x76, 0xa7, 0x65, 0x42, 0xbb, 0x53, 0x67, 0x66, 0x25, 0xdc, 0x85, 0x5f, 0x77, 0xe3, 0x45, 0x78,
	0x1d, 0x8a, 0x7a, 0x0d, 0x9e, 0x99, 0xb3, 0x4b, 0xe9, 0xd2, 0x5a, 0xf4, 0x4f, 0x9b, 0xd9, 0x73,
	0x9e, 0x8f, 0x3d, 0xfb, 0x9c, 0x5d, 0xb2, 0x1a, 0x2b, 0x2d, 0x24, 0x5b, 0x66, 0x49, 0x93, 0x27,
	0x6c, 0xb9, 0x23, 0x45, 0xc4, 0xe2, 0x54, 0x32, 0xb5, 0xdc, 0x10, 0xc1, 0x16, 0xd3, 0x9b, 0x42,
	0xa6, 0xed, 0xf5, 0x50, 0xb3, 0xa6, 0x90, 0x9c, 0xa9, 0x60, 0x27, 0xf5, 0xa0, 0x47, 0x0b, 0x7a,
	0x13, 0x81, 0x1e, 0x02, 0xbd, 0x41, 0xdd, 0x6e, 0x35, 0x13, 0x78, 0x17, 0xb6, 0x52, 0xa6, 0x10,
	0xec, 0x5e, 0xea, 0x55, 0x65, 0x52, 0x0a, 0x99, 0x95, 0xe6, 0x7b, 0x4b, 0x6d, 0xa6, 0x54, 0xd8,
	0x64, 0x59, 0xf1, 0x7a, 0xb1, 0xa8, 0x43, 0x9e, 0x34, 0x84, 0x6c, 0x87, 0x9a, 0x8b, 0x04, 0x9b,
	0x96, 0xbe, 0x54, 0x08, 0xd9, 0x09, 0x65, 0x08, 0x55, 0x26, 0x15, 0x7d, 0x43, 0x16, 0x3a, 0xf0,
	0x2f, 0x92, 0x80, 0xc7, 0x2c, 0xd1, 0xbc, 0xc1, 0x23, 0xdb, 0x1d, 0xa0, 0x23, 0x67, 0xe4, 0xea,
	0xc8, 0xad, 0xff, 0xef, 0xb8, 0x5e, 0x76, 0x3f, 0x99, 0x4f, 0xa5, 0x25, 0x4f, 0x9a, 0x2f, 0xcd,
	0xc1, 0x77, 0x11, 0x5f, 0xeb, 0x81, 0xdb, 0x92, 0xa2, 0x4f, 0xc9, 0xb5, 0x3f, 0xb1, 0x07, 0x49,
	0xda, 0x6a, 0x39, 0xdf, 0xc6, 0x41, 0x63, 0xc2, 0x5f, 0x1c, 0xcc, 0x53, 0x87, 0x36, 0xba, 0x46,
	0xa6, 0x32, 0x2e, 0x7d, 0xd4, 0x61, 0x40, 0xe8, 0x8c, 0x5a, 0x6f, 0xf3, 0x05, 0x6f, 0x3c, 0x81,
	0x01, 0x33, 0x89, 0xe6, 0x26, 0x11, 0xf2, 0x02, 0x10, 0xb5, 0x98, 0x7a, 0xa4, 0xda, 0x4b, 0x81,
	0x06, 0xbe, 0xa3, 0x81, 0x99, 0xd3, 0xbd, 0x56, 0x72, 0x95, 0xfc, 0x97, 0x26, 0xfc, 0x6d, 0x6a,
	0xd5, 0x4a, 0x43, 0x27, 0x31, 0x81, 0xcd, 0x20, 0x74, 0x83, 0x4c, 0x9d, 0x00, 0x51, 0xe3, 0x18,
	0x35, 0x26, 0xf3, 0x16, 0xcb, 0xbf, 0x45, 0x2e, 0x34, 0x4c, 0x26, 0x82, 0x08, 0x43, 0x71, 0x64,
	0x74, 0xc6, 0x86, 0xdf, 0xd5, 0x74, 0xe3, 0x54, 0x92, 0x8e, 0x40, 0x6f, 0x85, 0xcc, 0x9d, 0x21,
	0x42, 0xdd, 0x1f, 0xa8, 0x5b, 0x2d, 0x20, 0xac, 0xfc, 0x2b, 0xe2, 0xc6, 0x22, 0x48, 0x84, 0x0e,
	0xa2, 0x7d, 0x16, 0x1d, 0x04, 0x61, 0x14, 0x41, 0x9e, 0x02, 0xc9, 0x9b, 0xfb, 0x5a, 0x39, 0xe5,
	0xbe, 0x3e, 0xf6, 0x84, 0x68, 0xb1, 0x10, 0x9f, 0x8b, 0x3f, 0x17, 0x8b, 0xba, 0xd0, 0xeb, 0x06,
	0xbc, 0x66, 0xb1, 0xbe, 0x85, 0xd2, 0x0d, 0x72, 0x65, 0x30, 0x31, 0xfa, 0xfa, 0x89, 0xbe, 0xdc,
	0xfe, 0x0c, 0xd6, 0xde, 0x73, 0x72, 0x51, 0xb1, 0x0e, 0x44, 0x15, 0xf4, 0x03, 0x9e, 0x45, 0xc8,
	0x24, 0x47, 0x39, 0x95, 0xa1, 0xcf, 0x62, 0xf6, 0x04, 0x5a, 0xc3, 0x48, 0xc1, 0x65, 0x45, 0x1f,
	0x92, 0x85, 0x01, 0x94, 0xe8, 0xea, 0x17, 0xba, 0x72, 0xfa, 0x81, 0x8d, 0xa7, 0xa5, 0xe3, 0x31,
	0x32, 0xe1, 0x33, 0xd5, 0x11, 0x89, 0x62, 0xf4, 0x36, 0x29, 0xdb, 0xdd, 0x2c, 0x2e, 0x49, 0xb6,
	0xf4, 0xb8, 0xb7, 0x8f, 0xcd, 0xaf, 0x8f, 0x8d, 0x74, 0x97, 0xcc, 0x98, 0xad, 0x0c, 0x4e, 0xad,
	0x25, 0xa4, 0xb8, 0x04, 0x60, 0xaf, 0x00, 0x2e, 0x2e, 0xef, 0x36, 0x9c, 0x6b, 0xdd, 0xb3, 0x3f,
	0xdd, 0xee, 0xbd, 0x40, 0xef, 0x93, 0xf1, 0xec, 0x6d, 0x00, 0x49, 0x35, 0x8c, 0x8b, 0x67, 0x18,
	0xf1, 0x5d, 0xb1, 0x8d, 0xff, 0x7e, 0xde, 0x4e, 0x37, 0x49, 0x49, 0x8a, 0x43, 0xc8, 0x9d, 0x41,
	0xad, 0x78, 0xe7, 0x7b, 0x73, 0x79, 0xf9, 0x14, 0x3c, 0x5f, 0x1c, 0xfa, 0x86, 0xc0, 0xfd, 0x3a,
	0x4a, 0x4a, 0x70, 0xa0, 0x73, 0xa4, 0x02, 0x47, 0x13, 0xe5, 0xf7, 0x75, 0x18, 0x4c, 0xd9, 0x2f,
	0xc3, 0x11, 0x42, 0xfa, 0xa4, 0x5f, 0xda, 0x3f, 0xd4, 0xff, 0x21, 0xee, 0x75, 0x32, 0x7b, 0xc2,
	0x11, 0x33, 0x15, 0x49, 0xde, 0xb1, 0xa3, 0xfc, 0x58, 0x1f, 0x9a, 0x8b, 0x6a, 0x0e, 0xdc, 0xe8,
	0xe2, 0xe8, 0x3d, 0x32, 0xae, 0x84, 0xd4, 0x90, 0x58, 0xe7, 0xd3, 0x39, 0xfc, 0x54, 0x4c, 0x73,
	0x5d, 0xd0, 0x80, 0x5c, 0x8e, 0x44, 0x62, 0x9e, 0x93, 0xca, 0xef, 0xc9, 0x0c, 0x07, 0xa2, 0x65,
	0x0d, 0x2b, 0xe7, 0xf3, 0x39, 0xc8, 0xdc, 0x9c, 0xa2, 0x3b, 0xde, 0x67, 0xd2, 0xce, 0x5b, 0x3d,
	0xda, 0x25, 0xf3, 0x5c, 0x14, 0x1e, 0x48, 0xf7, 0x1b, 0xf4, 0xfa, 0x41, 0x53, 0xa8, 0xf8, 0x20,
	0xaf, 0xc7, 0x7f, 0xf3, 0x99, 0xda, 0xab, 0xd8, 0xaf, 0xc1, 0xdd, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x57, 0xbe, 0xcc, 0x2e, 0xe2, 0x06, 0x00, 0x00,
}
