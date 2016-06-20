// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetForums_Pu.proto
// DO NOT EDIT!

/*
Package fo_GetForums_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetForums_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetForums_Pu

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dstore_values "dstore/values"
import dstore_engine_error "dstore/engine/error"
import dstore_engine_message "dstore/engine/message"
import dstore_engine_metainformation "dstore/engine/metainformation"

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
	ForumId                        *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumIdNull                    bool                        `protobuf:"varint,1004,opt,name=forum_id_null,json=forumIdNull" json:"forum_id_null,omitempty"`
	LanguageId                     *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                 bool                        `protobuf:"varint,1005,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	ForumCategoryIdList            *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=forum_category_id_list,json=forumCategoryIdList" json:"forum_category_id_list,omitempty"`
	ForumCategoryIdListNull        bool                        `protobuf:"varint,1006,opt,name=forum_category_id_list_null,json=forumCategoryIdListNull" json:"forum_category_id_list_null,omitempty"`
	GetCategoryInformation         *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=get_category_information,json=getCategoryInformation" json:"get_category_information,omitempty"`
	GetCategoryInformationNull     bool                        `protobuf:"varint,1007,opt,name=get_category_information_null,json=getCategoryInformationNull" json:"get_category_information_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1008,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetForumCategoryIdList() *dstore_values.StringValue {
	if m != nil {
		return m.ForumCategoryIdList
	}
	return nil
}

func (m *Parameters) GetGetCategoryInformation() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetCategoryInformation
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
	RowId                 int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Language              *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=language" json:"language,omitempty"`
	ForumId               *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumName             *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=forum_name,json=forumName" json:"forum_name,omitempty"`
	LanguageId            *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	PredecessorCategoryId *dstore_values.IntegerValue `protobuf:"bytes,20002,opt,name=predecessor_category_id,json=predecessorCategoryId" json:"predecessor_category_id,omitempty"`
	ForumCategoryId       *dstore_values.IntegerValue `protobuf:"bytes,20003,opt,name=forum_category_id,json=forumCategoryId" json:"forum_category_id,omitempty"`
	CategoryDescription   *dstore_values.StringValue  `protobuf:"bytes,20006,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	SortNo                *dstore_values.IntegerValue `protobuf:"bytes,20007,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetLanguage() *dstore_values.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *Response_Row) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Response_Row) GetForumName() *dstore_values.StringValue {
	if m != nil {
		return m.ForumName
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Response_Row) GetPredecessorCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredecessorCategoryId
	}
	return nil
}

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

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetForums_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetForums_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetForums_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xeb, 0x4e, 0x13, 0x41,
	0x14, 0x0e, 0x94, 0x5e, 0x38, 0x20, 0xe0, 0x16, 0x61, 0x6d, 0x81, 0x20, 0xc4, 0x04, 0x63, 0xb2,
	0x18, 0x4d, 0xc4, 0x18, 0x8d, 0x11, 0x2f, 0xa4, 0x46, 0x2a, 0x6e, 0xc4, 0x44, 0x63, 0xd2, 0x2c,
	0xed, 0xb4, 0xd9, 0xa4, 0xdd, 0xa9, 0x33, 0xbb, 0x12, 0x5e, 0xc2, 0x78, 0xfb, 0xc9, 0x2f, 0x4d,
	0xf4, 0x09, 0x7c, 0x07, 0x5f, 0xc5, 0xbb, 0x8f, 0xe0, 0x99, 0x39, 0xbb, 0xbd, 0x2c, 0x85, 0x85,
	0x3f, 0xdd, 0xcc, 0xce, 0xf7, 0x7d, 0xe7, 0xdb, 0xd3, 0xf9, 0xce, 0x2e, 0x58, 0x35, 0xe9, 0x73,
	0xc1, 0x56, 0x99, 0xd7, 0x70, 0x3d, 0xb6, 0xda, 0x16, 0xbc, 0xca, 0x6a, 0x81, 0x60, 0x72, 0xb5,
	0xce, 0x2b, 0x1b, 0xcc, 0xbf, 0xcf, 0x45, 0xd0, 0x92, 0x95, 0xad, 0xc0, 0xc2, 0x2d, 0x9f, 0x1b,
	0xf3, 0x84, 0xb7, 0x08, 0x6f, 0xc5, 0x40, 0x85, 0x7c, 0x28, 0xf7, 0xca, 0x69, 0x06, 0x4c, 0x12,
	0xa7, 0x70, 0xb6, 0xbf, 0x06, 0x13, 0x82, 0x8b, 0x70, 0xab, 0xd8, 0xbf, 0xd5, 0x62, 0x52, 0x3a,
	0x0d, 0x16, 0x6e, 0x2e, 0xc7, 0x37, 0x7d, 0xc7, 0xf5, 0xea, 0x5c, 0xb4, 0x1c, 0xdf, 0xe5, 0x1e,
	0x81, 0x96, 0x5e, 0xe7, 0x00, 0xb6, 0x1c, 0xe1, 0xe0, 0x2e, 0x13, 0xd2, 0x78, 0x01, 0x73, 0x6d,
	0xbc, 0x72, 0xaf, 0xe2, 0xd6, 0x98, 0xe7, 0xbb, 0x75, 0xb7, 0xaa, 0xd1, 0x15, 0x72, 0x64, 0x0e,
	0x2d, 0x0e, 0xad, 0x8c, 0x5d, 0x2e, 0x84, 0x8f, 0x6d, 0x85, 0x3e, 0xa5, 0x2f, 0x5c, 0xaf, 0xf1,
	0x54, 0x2d, 0xec, 0x02, 0xf1, 0x4b, 0x7d, 0x74, 0xbd, 0x25, 0x8d, 0x07, 0x70, 0xee, 0x28, 0xf5,
	0x8a, 0x17, 0x34, 0x9b, 0xe6, 0xf7, 0x2c, 0xd6, 0xc8, 0xd9, 0x0b, 0x87, 0xeb, 0x94, 0x11, 0x66,
	0xdc, 0x86, 0x89, 0x50, 0xcb, 0xdf, 0x6b, 0x33, 0x14, 0x34, 0x87, 0xb5, 0xb7, 0x62, 0xcc, 0x9b,
	0xeb, 0xf9, 0xac, 0xc1, 0x04, 0x99, 0x1b, 0x27, 0xca, 0x13, 0x64, 0x94, 0x6a, 0x86, 0x05, 0xf9,
	0x7e, 0x09, 0x32, 0xf0, 0x83, 0x0c, 0x4c, 0xf5, 0x62, 0x75, 0xc9, 0x35, 0x18, 0x0d, 0x3c, 0xf7,
	0x65, 0xa0, 0xab, 0xa5, 0x12, 0x3b, 0x91, 0x23, 0x30, 0x16, 0x3a, 0x0f, 0x13, 0x1d, 0x22, 0xd5,
	0xf8, 0x49, 0x35, 0xc6, 0x23, 0x88, 0xd6, 0xbf, 0x0a, 0xb9, 0xba, 0x3a, 0x0a, 0x4a, 0x7e, 0x24,
	0xf9, 0x61, 0xb2, 0x1a, 0x8c, 0xf2, 0xcb, 0x70, 0x2a, 0xe2, 0x91, 0xfa, 0x2f, 0x52, 0x1f, 0x0b,
	0x01, 0x5a, 0xfc, 0x06, 0x8c, 0x35, 0x1d, 0xaf, 0x11, 0xe0, 0xf9, 0x50, 0xfa, 0xe9, 0x64, 0x7d,
	0x88, 0xf0, 0x58, 0xe2, 0x02, 0x4c, 0xf5, 0xb0, 0xa9, 0xca, 0x6f, 0xaa, 0x32, 0xd1, 0x85, 0xe9,
	0x42, 0x8f, 0x60, 0x86, 0xdc, 0xe0, 0x5f, 0xc6, 0x1a, 0x5c, 0xec, 0x29, 0x42, 0xd3, 0x95, 0xbe,
	0x99, 0x49, 0x6c, 0x59, 0x5e, 0x33, 0xef, 0x84, 0xc4, 0x52, 0xed, 0x21, 0xd2, 0x8c, 0x9b, 0x50,
	0x1c, 0x2c, 0x48, 0x36, 0xfe, 0x90, 0x8d, 0xd9, 0x01, 0x54, 0xed, 0x67, 0x1b, 0xcc, 0x06, 0xf3,
	0x7b, 0xc8, 0xdd, 0x0c, 0x98, 0xd9, 0x81, 0x5d, 0xd8, 0xe1, 0xbc, 0xc9, 0x1c, 0x3a, 0x6c, 0xf6,
	0x0c, 0x92, 0x3b, 0xaa, 0x5d, 0xaa, 0xb1, 0x0e, 0xf3, 0x87, 0xc9, 0x92, 0xaf, 0xbf, 0xe4, 0xab,
	0x30, 0x98, 0xaf, 0xad, 0x3d, 0x86, 0x59, 0xc9, 0xda, 0x98, 0x3e, 0xac, 0x8e, 0x02, 0x94, 0x0a,
	0x15, 0x06, 0x69, 0xe6, 0x12, 0x7b, 0x35, 0xdd, 0xa1, 0x96, 0x28, 0x25, 0x78, 0x5b, 0x1a, 0xb7,
	0x60, 0xee, 0x10, 0x49, 0x72, 0xf5, 0x8f, 0x5c, 0x99, 0x83, 0xc8, 0xca, 0xd3, 0xd2, 0xb7, 0x0c,
	0xe4, 0x6c, 0x26, 0xdb, 0xdc, 0x93, 0xcc, 0xb8, 0x04, 0x69, 0x3d, 0x6e, 0xe2, 0xb9, 0x0f, 0xc7,
	0x17, 0x8d, 0xa2, 0x7b, 0xea, 0xd7, 0x26, 0xa0, 0xf1, 0x0c, 0xa6, 0xd4, 0xa0, 0xe9, 0xeb, 0xf2,
	0xf0, 0x62, 0x0a, 0xc9, 0x56, 0x8c, 0x1c, 0x9f, 0x47, 0x9b, 0xb8, 0xee, 0x69, 0x90, 0x3d, 0xd9,
	0xea, 0xbf, 0x61, 0x5c, 0x83, 0x6c, 0x38, 0xe0, 0x30, 0x7c, 0x4a, 0x71, 0xe1, 0x80, 0x22, 0x8d,
	0xbf, 0x4d, 0xba, 0xda, 0x11, 0x1c, 0x4f, 0x50, 0x4a, 0xf0, 0x5d, 0xcc, 0x94, 0x62, 0x5d, 0xb4,
	0x8e, 0x9c, 0xc1, 0x56, 0xf4, 0xf0, 0x96, 0xcd, 0x77, 0x6d, 0xc5, 0x2b, 0x7c, 0x1d, 0x81, 0x14,
	0x2e, 0x8c, 0x19, 0xc8, 0xe0, 0x52, 0xa5, 0xe7, 0x4d, 0x19, 0xfb, 0x91, 0xb6, 0xd3, 0xb8, 0xc4,
	0x70, 0xac, 0x41, 0x2e, 0xca, 0x80, 0xf9, 0xb6, 0x9c, 0x3c, 0x17, 0x22, 0xb0, 0x22, 0x76, 0x02,
	0xff, 0xae, 0x7c, 0x82, 0xc4, 0x5f, 0x07, 0x20, 0xa2, 0x87, 0x83, 0xdb, 0x7c, 0x9f, 0x5c, 0x73,
	0x54, 0xc3, 0xcb, 0x88, 0xc6, 0x66, 0xf4, 0x0d, 0x82, 0x0f, 0xe5, 0x93, 0x4d, 0x82, 0x6d, 0x98,
	0x6d, 0x0b, 0x56, 0x63, 0x55, 0xec, 0x2d, 0x1e, 0xb1, 0x9e, 0x4c, 0x9a, 0x1f, 0xf7, 0x87, 0x92,
	0xb5, 0xce, 0xf4, 0xb0, 0xbb, 0x69, 0x35, 0x4a, 0x70, 0xfa, 0x40, 0xc8, 0xcd, 0x4f, 0xc7, 0x11,
	0x9c, 0x8c, 0x05, 0x1f, 0x07, 0xd0, 0x74, 0x47, 0xa4, 0xc6, 0x64, 0x55, 0xb8, 0x6d, 0x7d, 0x0c,
	0x3f, 0xef, 0x27, 0xbf, 0xbd, 0xf2, 0x11, 0xf3, 0x6e, 0x97, 0x88, 0x73, 0x39, 0x8b, 0x66, 0x71,
	0xdc, 0x70, 0xf3, 0xcb, 0x71, 0x1c, 0x65, 0x14, 0xba, 0xcc, 0xd7, 0x37, 0xa0, 0xe8, 0xf2, 0xd8,
	0x69, 0xeb, 0x7e, 0x21, 0x3c, 0x5f, 0x39, 0xee, 0xb7, 0xc3, 0x4e, 0x46, 0xbf, 0xab, 0xaf, 0xfc,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x84, 0x3e, 0x6f, 0x20, 0x6e, 0x08, 0x00, 0x00,
}