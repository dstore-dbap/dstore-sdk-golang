// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetPostingCharacs_Ad.proto
// DO NOT EDIT!

/*
Package fo_GetPostingCharacs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetPostingCharacs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetPostingCharacs_Ad

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
	PostingCharacteristicId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=posting_characteristic_id,json=postingCharacteristicId" json:"posting_characteristic_id,omitempty"`
	PostingCharacteristicIdNull       bool                        `protobuf:"varint,1001,opt,name=posting_characteristic_id_null,json=postingCharacteristicIdNull" json:"posting_characteristic_id_null,omitempty"`
	GetAssignedForumsOrCategories     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=get_assigned_forums_or_categories,json=getAssignedForumsOrCategories" json:"get_assigned_forums_or_categories,omitempty"`
	GetAssignedForumsOrCategoriesNull bool                        `protobuf:"varint,1002,opt,name=get_assigned_forums_or_categories_null,json=getAssignedForumsOrCategoriesNull" json:"get_assigned_forums_or_categories_null,omitempty"`
	CharacsAssignedToForumId          *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=characs_assigned_to_forum_id,json=characsAssignedToForumId" json:"characs_assigned_to_forum_id,omitempty"`
	CharacsAssignedToForumIdNull      bool                        `protobuf:"varint,1003,opt,name=characs_assigned_to_forum_id_null,json=characsAssignedToForumIdNull" json:"characs_assigned_to_forum_id_null,omitempty"`
	CharacsAssignedToCategoryId       *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=characs_assigned_to_category_id,json=characsAssignedToCategoryId" json:"characs_assigned_to_category_id,omitempty"`
	CharacsAssignedToCategoryIdNull   bool                        `protobuf:"varint,1004,opt,name=characs_assigned_to_category_id_null,json=characsAssignedToCategoryIdNull" json:"characs_assigned_to_category_id_null,omitempty"`
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

func (m *Parameters) GetGetAssignedForumsOrCategories() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetAssignedForumsOrCategories
	}
	return nil
}

func (m *Parameters) GetCharacsAssignedToForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CharacsAssignedToForumId
	}
	return nil
}

func (m *Parameters) GetCharacsAssignedToCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CharacsAssignedToCategoryId
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
	RowId                       int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CharacteristicDescription   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	FieldTypeId                 *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	PostingCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=posting_characteristic_id,json=postingCharacteristicId" json:"posting_characteristic_id,omitempty"`
	PrecisionValue              *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=precision_value,json=precisionValue" json:"precision_value,omitempty"`
	CommonCharacteristic        *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=common_characteristic,json=commonCharacteristic" json:"common_characteristic,omitempty"`
	MaxLength                   *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=max_length,json=maxLength" json:"max_length,omitempty"`
	PredefinedValues            *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=predefined_values,json=predefinedValues" json:"predefined_values,omitempty"`
	CheckPostingVisibility      *dstore_values.BooleanValue `protobuf:"bytes,10008,opt,name=check_posting_visibility,json=checkPostingVisibility" json:"check_posting_visibility,omitempty"`
	Format                      *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=format" json:"format,omitempty"`
	BasicFieldType              *dstore_values.StringValue  `protobuf:"bytes,10010,opt,name=basic_field_type,json=basicFieldType" json:"basic_field_type,omitempty"`
	FieldTypeDescription        *dstore_values.StringValue  `protobuf:"bytes,10011,opt,name=field_type_description,json=fieldTypeDescription" json:"field_type_description,omitempty"`
	BasicFieldTypeId            *dstore_values.IntegerValue `protobuf:"bytes,10012,opt,name=basic_field_type_id,json=basicFieldTypeId" json:"basic_field_type_id,omitempty"`
	PropertyModificationAllowed *dstore_values.IntegerValue `protobuf:"bytes,10013,opt,name=property_modification_allowed,json=propertyModificationAllowed" json:"property_modification_allowed,omitempty"`
	MaxNumberOfProperties       *dstore_values.IntegerValue `protobuf:"bytes,10014,opt,name=max_number_of_properties,json=maxNumberOfProperties" json:"max_number_of_properties,omitempty"`
	ForumId                     *dstore_values.IntegerValue `protobuf:"bytes,40003,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumName                   *dstore_values.StringValue  `protobuf:"bytes,40004,opt,name=forum_name,json=forumName" json:"forum_name,omitempty"`
	ForumCategoryId             *dstore_values.IntegerValue `protobuf:"bytes,50011,opt,name=forum_category_id,json=forumCategoryId" json:"forum_category_id,omitempty"`
	CategoryDescription         *dstore_values.StringValue  `protobuf:"bytes,50014,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
	}
	return nil
}

func (m *Response_Row) GetPostingCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetPrecisionValue() *dstore_values.IntegerValue {
	if m != nil {
		return m.PrecisionValue
	}
	return nil
}

func (m *Response_Row) GetCommonCharacteristic() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommonCharacteristic
	}
	return nil
}

func (m *Response_Row) GetMaxLength() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxLength
	}
	return nil
}

func (m *Response_Row) GetPredefinedValues() *dstore_values.BooleanValue {
	if m != nil {
		return m.PredefinedValues
	}
	return nil
}

func (m *Response_Row) GetCheckPostingVisibility() *dstore_values.BooleanValue {
	if m != nil {
		return m.CheckPostingVisibility
	}
	return nil
}

func (m *Response_Row) GetFormat() *dstore_values.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *Response_Row) GetBasicFieldType() *dstore_values.StringValue {
	if m != nil {
		return m.BasicFieldType
	}
	return nil
}

func (m *Response_Row) GetFieldTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.FieldTypeDescription
	}
	return nil
}

func (m *Response_Row) GetBasicFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicFieldTypeId
	}
	return nil
}

func (m *Response_Row) GetPropertyModificationAllowed() *dstore_values.IntegerValue {
	if m != nil {
		return m.PropertyModificationAllowed
	}
	return nil
}

func (m *Response_Row) GetMaxNumberOfProperties() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfProperties
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

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetPostingCharacs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetPostingCharacs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetPostingCharacs_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_GetPostingCharacs_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0x56, 0x09, 0xb9, 0xf4, 0x54, 0x6d, 0xc2, 0x24, 0x0d, 0x1b, 0xbb, 0x17, 0x52, 0x2e, 0xe2,
	0x69, 0x83, 0x88, 0x84, 0x40, 0xbc, 0x10, 0xda, 0x14, 0x22, 0x88, 0x13, 0x56, 0x25, 0x94, 0xf2,
	0x30, 0xac, 0x77, 0xc7, 0xee, 0xa8, 0xbb, 0x3b, 0xd6, 0xcc, 0xba, 0xc1, 0x7f, 0x80, 0x67, 0x28,
	0xf7, 0xeb, 0xef, 0xe0, 0x11, 0x09, 0xfe, 0x05, 0x12, 0xef, 0x5c, 0xde, 0x79, 0xe5, 0xcc, 0x65,
	0xed, 0x5d, 0xb7, 0xf6, 0x2e, 0xe2, 0xc5, 0xd6, 0xec, 0x9c, 0xef, 0x3b, 0xdf, 0xcc, 0xf9, 0xce,
	0xcc, 0xc0, 0x4b, 0xb1, 0xca, 0x85, 0x64, 0x3b, 0x2c, 0xeb, 0xf3, 0x8c, 0xed, 0x0c, 0xa4, 0x88,
	0x58, 0x3c, 0x94, 0x4c, 0xed, 0xf4, 0x04, 0x7d, 0x83, 0xe5, 0xc7, 0x42, 0xe5, 0x3c, 0xeb, 0x5f,
	0xbf, 0x1b, 0xca, 0x30, 0x52, 0x74, 0x2f, 0xf6, 0x31, 0x24, 0x17, 0xe4, 0x59, 0x8b, 0xf3, 0x2d,
	0xce, 0x9f, 0x11, 0xdc, 0x5a, 0x77, 0xf4, 0xf7, 0xc3, 0x64, 0xc8, 0x94, 0xc5, 0xb6, 0xb6, 0xaa,
	0x39, 0x99, 0x94, 0x42, 0xba, 0xa9, 0x76, 0x75, 0x2a, 0x65, 0x4a, 0x85, 0x7d, 0xe6, 0x26, 0x9f,
	0x9e, 0x9e, 0xcc, 0x43, 0x9e, 0xf5, 0x84, 0x4c, 0xc3, 0x9c, 0x8b, 0xcc, 0x06, 0x5d, 0xfb, 0x79,
	0x11, 0xe0, 0x18, 0xf3, 0xe3, 0x2c, 0x93, 0x8a, 0xbc, 0x07, 0x5b, 0x03, 0xab, 0x8a, 0x46, 0x46,
	0x16, 0x7e, 0xe5, 0x38, 0x8e, 0x28, 0x8f, 0xbd, 0x33, 0x4f, 0x9d, 0x79, 0xfe, 0xdc, 0x8b, 0x6d,
	0xdf, 0xad, 0xc5, 0x89, 0xe4, 0x59, 0xce, 0xfa, 0x4c, 0x9e, 0xe8, 0x51, 0xf0, 0xe4, 0xa0, 0xbc,
	0xa6, 0x02, 0x7c, 0x10, 0x93, 0x1b, 0x70, 0x65, 0x26, 0x31, 0xcd, 0x86, 0x49, 0xe2, 0xfd, 0xb1,
	0x8c, 0xf4, 0x2b, 0x41, 0x7b, 0x06, 0x43, 0x07, 0x63, 0x08, 0x83, 0xed, 0x3e, 0xcb, 0x69, 0xa8,
	0x14, 0xef, 0x67, 0x2c, 0xa6, 0xb8, 0x9a, 0x61, 0xaa, 0xa8, 0x90, 0x34, 0x0a, 0x51, 0x85, 0x90,
	0x9c, 0x29, 0xef, 0xb1, 0x7a, 0x99, 0x97, 0x91, 0x65, 0xcf, 0x91, 0xdc, 0x34, 0x1c, 0x47, 0xf2,
	0xfa, 0x98, 0x81, 0x04, 0xf0, 0x5c, 0x6d, 0x1a, 0x2b, 0xfa, 0x4f, 0x2b, 0x7a, 0x7b, 0x2e, 0x9f,
	0x91, 0xfe, 0x01, 0x5c, 0x8a, 0x5c, 0xa1, 0xc7, 0xbc, 0xb9, 0xb0, 0xd4, 0x7a, 0x73, 0x17, 0xea,
	0x55, 0x7b, 0x8e, 0xa0, 0xc8, 0x74, 0x4b, 0x98, 0x5c, 0xb8, 0xbb, 0x6f, 0xc2, 0xf6, 0x3c, 0x72,
	0xab, 0xf5, 0x2f, 0xab, 0xf5, 0xd2, 0x2c, 0x16, 0x23, 0x33, 0x84, 0xab, 0x8f, 0x62, 0x72, 0xeb,
	0x1e, 0x69, 0xa5, 0x8f, 0xd7, 0x2b, 0x6d, 0x3f, 0x94, 0xc3, 0xed, 0xc6, 0x08, 0xc5, 0x76, 0xe0,
	0x99, 0x9a, 0x14, 0x56, 0xef, 0xdf, 0x56, 0xef, 0xd5, 0x39, 0x5c, 0x5a, 0xf2, 0xb5, 0x07, 0xe7,
	0x61, 0x25, 0x60, 0x6a, 0x20, 0x32, 0xc5, 0xc8, 0x0b, 0xb0, 0x68, 0x1a, 0xc4, 0x99, 0xb5, 0xe5,
	0x57, 0x1b, 0xcf, 0x36, 0xcf, 0xbe, 0xfe, 0x0d, 0x6c, 0x20, 0x79, 0x1f, 0xd6, 0x74, 0x6b, 0xd0,
	0x52, 0x6f, 0xa0, 0x85, 0x16, 0x10, 0xec, 0x4f, 0x81, 0xa7, 0x3b, 0xe8, 0x10, 0xc7, 0x07, 0x93,
	0x71, 0xb0, 0x9a, 0x56, 0x3f, 0x90, 0x97, 0x61, 0xd9, 0xb5, 0x24, 0x96, 0x57, 0x33, 0x5e, 0x79,
	0x88, 0xd1, 0x36, 0xec, 0xa1, 0xfd, 0x0f, 0x8a, 0x70, 0xb2, 0x0f, 0x0b, 0x52, 0x9c, 0xe2, 0x56,
	0x6b, 0xd4, 0xae, 0xdf, 0xe8, 0xf4, 0xf0, 0x8b, 0x4d, 0xf0, 0x03, 0x71, 0x1a, 0x68, 0x7c, 0xeb,
	0x1f, 0x80, 0x05, 0x1c, 0x90, 0x4d, 0x58, 0xc2, 0xa1, 0x2e, 0xde, 0x27, 0x1d, 0xdc, 0x97, 0xc5,
	0x60, 0x11, 0x87, 0x58, 0x8a, 0x3b, 0xd0, 0x9a, 0xea, 0xc6, 0x98, 0xa9, 0x48, 0xf2, 0x81, 0xd9,
	0x85, 0x4f, 0x3b, 0xd5, 0x3d, 0x74, 0x95, 0x56, 0xb9, 0xc4, 0xbc, 0xb6, 0xd0, 0x5b, 0x55, 0xf8,
	0x8d, 0x09, 0x9a, 0xbc, 0x06, 0xe7, 0x7b, 0x9c, 0x25, 0x58, 0xdb, 0xd1, 0x80, 0xe9, 0xd4, 0x0f,
	0x3a, 0xf5, 0xc6, 0x39, 0x67, 0x20, 0xb7, 0x10, 0x81, 0xea, 0x6e, 0xcf, 0x3b, 0x8c, 0x3e, 0xeb,
	0xfc, 0x8f, 0xd3, 0x68, 0x1f, 0x56, 0x07, 0x92, 0x45, 0x5c, 0xa1, 0x50, 0x6a, 0x90, 0xde, 0xe7,
	0x0d, 0xf8, 0x2e, 0x8c, 0x41, 0x66, 0x4c, 0xde, 0x81, 0x8b, 0x91, 0x48, 0x53, 0xe4, 0xa8, 0xea,
	0xf3, 0xbe, 0x68, 0x40, 0xb6, 0x61, 0xa1, 0x55, 0x6d, 0xe4, 0x55, 0x80, 0x34, 0xfc, 0x88, 0x26,
	0x58, 0xea, 0xfc, 0xae, 0xf7, 0x65, 0x03, 0x9e, 0xb3, 0x18, 0xff, 0xb6, 0x09, 0x27, 0x07, 0xf0,
	0x04, 0x2a, 0x8c, 0x59, 0x8f, 0xeb, 0x9e, 0xb2, 0xc1, 0xde, 0x57, 0x8f, 0xe6, 0xe8, 0x0a, 0x91,
	0xb0, 0xd0, 0x2e, 0x24, 0x58, 0x9b, 0xc0, 0xcc, 0x07, 0x45, 0x4e, 0x00, 0x4f, 0x1b, 0x16, 0xdd,
	0xa3, 0x45, 0x05, 0xee, 0xe3, 0xba, 0xbb, 0x3c, 0xe1, 0xf9, 0xc8, 0xfb, 0xba, 0x01, 0xe3, 0xa6,
	0x41, 0x3b, 0x8f, 0x9e, 0x8c, 0xb1, 0x64, 0x17, 0x96, 0x6c, 0x7f, 0x78, 0xdf, 0xd4, 0xbb, 0xcb,
	0x85, 0x62, 0xb9, 0xd6, 0xba, 0xa1, 0xc2, 0xba, 0x4f, 0x0c, 0xe5, 0x7d, 0x5b, 0x0f, 0xbf, 0x60,
	0x40, 0x37, 0x0b, 0x47, 0x61, 0xb9, 0x36, 0x4b, 0x8e, 0x2c, 0x3b, 0xfd, 0xbb, 0x7a, 0xb2, 0x8d,
	0xb1, 0x33, 0xcb, 0x26, 0x7f, 0x0b, 0xd6, 0xa7, 0x95, 0x69, 0x73, 0x7e, 0xdf, 0xa0, 0x6e, 0x6b,
	0x55, 0x75, 0xe8, 0xca, 0x0f, 0xe1, 0x32, 0x5e, 0xca, 0x03, 0x26, 0xf3, 0x11, 0x4d, 0x45, 0xcc,
	0x7b, 0x3c, 0x32, 0xe7, 0x08, 0x0d, 0x93, 0x44, 0x9c, 0xb2, 0xd8, 0xfb, 0xa1, 0x01, 0x6d, 0xbb,
	0xa0, 0x38, 0x2c, 0x31, 0xec, 0x59, 0x02, 0xf2, 0x2e, 0x78, 0xda, 0x5d, 0xd9, 0x30, 0xed, 0x32,
	0x49, 0x45, 0x8f, 0xba, 0x60, 0x7d, 0x6d, 0xfe, 0xd8, 0x80, 0xfc, 0x22, 0xa2, 0x3b, 0x06, 0x7c,
	0xd4, 0x3b, 0x1e, 0x43, 0xf1, 0x9c, 0x5b, 0x19, 0xdf, 0x63, 0xbf, 0xfc, 0xd4, 0xe0, 0xfe, 0x5d,
	0xee, 0xb9, 0x8b, 0x0b, 0xed, 0x6e, 0x91, 0x19, 0xbe, 0x40, 0xbc, 0x5f, 0x1d, 0x76, 0x5e, 0x1d,
	0xce, 0x9a, 0xf8, 0x0e, 0x86, 0x6b, 0xbb, 0x5b, 0x70, 0xf9, 0x76, 0xfa, 0xed, 0xe3, 0x06, 0x37,
	0xe9, 0xaa, 0xc1, 0x95, 0xee, 0xa4, 0x23, 0xd8, 0x18, 0x93, 0x94, 0x8d, 0xf1, 0xbb, 0x63, 0x9b,
	0xa7, 0x68, 0xbd, 0x40, 0x96, 0x8c, 0xf1, 0xfa, 0x6d, 0x68, 0x73, 0x31, 0x75, 0x6e, 0x4f, 0x5e,
	0x8b, 0x77, 0x5e, 0xe9, 0x0b, 0x15, 0xdf, 0x2b, 0xe6, 0xe3, 0xff, 0xf0, 0xa0, 0xec, 0x2e, 0x99,
	0x87, 0xdb, 0xee, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0xa1, 0xf1, 0xae, 0x8b, 0x0a, 0x00,
	0x00,
}