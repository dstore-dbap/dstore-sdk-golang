// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_ModifyForumCategories_Ad.proto
// DO NOT EDIT!

/*
Package fo_ModifyForumCategories_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_ModifyForumCategories_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_ModifyForumCategories_Ad

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
	ForumCategoryId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=forum_category_id,json=forumCategoryId" json:"forum_category_id,omitempty"`
	ForumCategoryIdNull       bool                        `protobuf:"varint,1001,opt,name=forum_category_id_null,json=forumCategoryIdNull" json:"forum_category_id_null,omitempty"`
	CategoryDescription       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	CategoryDescriptionNull   bool                        `protobuf:"varint,1002,opt,name=category_description_null,json=categoryDescriptionNull" json:"category_description_null,omitempty"`
	PredecessorCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=predecessor_category_id,json=predecessorCategoryId" json:"predecessor_category_id,omitempty"`
	PredecessorCategoryIdNull bool                        `protobuf:"varint,1003,opt,name=predecessor_category_id_null,json=predecessorCategoryIdNull" json:"predecessor_category_id_null,omitempty"`
	MoveSortNo                *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=move_sort_no,json=moveSortNo" json:"move_sort_no,omitempty"`
	MoveSortNoNull            bool                        `protobuf:"varint,1004,opt,name=move_sort_no_null,json=moveSortNoNull" json:"move_sort_no_null,omitempty"`
	DeleteCategory            *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_category,json=deleteCategory" json:"delete_category,omitempty"`
	DeleteCategoryNull        bool                        `protobuf:"varint,1005,opt,name=delete_category_null,json=deleteCategoryNull" json:"delete_category_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetForumCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumCategoryId
	}
	return nil
}

func (m *Parameters) GetCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CategoryDescription
	}
	return nil
}

func (m *Parameters) GetPredecessorCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredecessorCategoryId
	}
	return nil
}

func (m *Parameters) GetMoveSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.MoveSortNo
	}
	return nil
}

func (m *Parameters) GetDeleteCategory() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCategory
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_ModifyForumCategories_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_ModifyForumCategories_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_ModifyForumCategories_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_ModifyForumCategories_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x6b, 0x6b, 0x13, 0x41,
	0x14, 0xa5, 0x8d, 0x69, 0xcb, 0x55, 0x5a, 0x3b, 0xa9, 0x6d, 0x1e, 0x2a, 0x52, 0xbf, 0xa8, 0x1f,
	0x36, 0xbe, 0x40, 0x51, 0x0a, 0x3e, 0xaa, 0x92, 0x0f, 0x09, 0xb2, 0x05, 0x41, 0x11, 0x96, 0x6d,
	0xe6, 0x26, 0x0c, 0x26, 0x73, 0xc3, 0xcc, 0xa6, 0x25, 0xff, 0x42, 0xfc, 0x3b, 0xfe, 0x22, 0x5f,
	0xff, 0xc1, 0x99, 0xbd, 0x9b, 0xc7, 0x6e, 0xa3, 0xd5, 0x2f, 0x49, 0x26, 0xf7, 0x9c, 0x7b, 0xce,
	0xde, 0x3d, 0x77, 0xe0, 0x89, 0xb4, 0x09, 0x19, 0x6c, 0xa2, 0xee, 0x2b, 0x8d, 0xcd, 0x91, 0xa1,
	0x2e, 0xca, 0xb1, 0x41, 0xdb, 0xec, 0x51, 0xd4, 0x26, 0xa9, 0x7a, 0x93, 0xd7, 0x64, 0xc6, 0xc3,
	0x97, 0x71, 0x82, 0x7d, 0x32, 0x0a, 0x6d, 0xf4, 0x5c, 0x06, 0x0e, 0x96, 0x90, 0xb8, 0xcd, 0xdc,
	0x80, 0xb9, 0xc1, 0x5f, 0x08, 0xf5, 0x4a, 0x26, 0x73, 0x12, 0x0f, 0xc6, 0x68, 0x99, 0x5f, 0xaf,
	0xe5, 0xb5, 0xd1, 0x18, 0x32, 0x59, 0xa9, 0x91, 0x2f, 0x0d, 0xd1, 0xda, 0xb8, 0x8f, 0x59, 0xf1,
	0x66, 0xb1, 0x98, 0xc4, 0x4a, 0xf7, 0xc8, 0x0c, 0xe3, 0x44, 0x91, 0x66, 0xd0, 0xfe, 0x97, 0x32,
	0xc0, 0xdb, 0xd8, 0xc4, 0xae, 0x8a, 0xc6, 0x8a, 0x37, 0xb0, 0xdd, 0xf3, 0xb6, 0xa2, 0x2e, 0xfb,
	0x9a, 0x44, 0x4a, 0x56, 0x57, 0x6e, 0xac, 0xdc, 0xba, 0x78, 0xbf, 0x11, 0x64, 0xcf, 0x91, 0x99,
	0x53, 0xda, 0x21, 0xd0, 0xbc, 0xf3, 0xa7, 0x70, 0xab, 0xb7, 0xf0, 0x30, 0x93, 0x96, 0x14, 0x0f,
	0x61, 0xf7, 0x4c, 0xa3, 0x48, 0x8f, 0x07, 0x83, 0xea, 0xb7, 0x75, 0xd7, 0x6e, 0x23, 0xac, 0x14,
	0x18, 0x1d, 0x57, 0x13, 0x6d, 0xd8, 0x99, 0xe1, 0x25, 0xda, 0xae, 0x51, 0x23, 0xef, 0xb5, 0xba,
	0x9a, 0x3a, 0xa8, 0x17, 0x1c, 0xd8, 0xc4, 0x28, 0xdd, 0x67, 0x03, 0x95, 0x29, 0xef, 0x70, 0x4e,
	0x13, 0x4f, 0xa1, 0xb6, 0xac, 0x1d, 0xfb, 0xf8, 0xce, 0x3e, 0xf6, 0x96, 0x10, 0x53, 0x2f, 0x47,
	0xb0, 0x37, 0x32, 0x28, 0xb1, 0xeb, 0x86, 0x4a, 0x26, 0x37, 0x90, 0xd2, 0xf9, 0x03, 0xb9, 0xb2,
	0xc0, 0x5d, 0x18, 0xcb, 0x33, 0xb8, 0xfa, 0x87, 0xa6, 0x6c, 0xea, 0x07, 0x9b, 0xaa, 0x2d, 0x65,
	0xa7, 0xb6, 0x0e, 0xe0, 0xd2, 0x90, 0x4e, 0x30, 0x72, 0x95, 0x24, 0xd2, 0x54, 0xbd, 0x70, 0xbe,
	0x17, 0xf0, 0x84, 0x23, 0x87, 0xef, 0x90, 0xb8, 0x03, 0xdb, 0x8b, 0x74, 0x56, 0xfd, 0xc9, 0xaa,
	0x9b, 0x73, 0x5c, 0x2a, 0x75, 0x08, 0x5b, 0x12, 0x07, 0x2e, 0x18, 0x33, 0x9f, 0xd5, 0xf2, 0x52,
	0xb5, 0x63, 0xa2, 0x01, 0xc6, 0x9a, 0xd5, 0x36, 0x99, 0x33, 0xb5, 0x2d, 0xee, 0xc1, 0x4e, 0xa1,
	0x0b, 0x8b, 0xfe, 0x62, 0x51, 0x91, 0x87, 0x7b, 0xe1, 0xfd, 0xaf, 0xab, 0xb0, 0x11, 0xa2, 0x1d,
	0x91, 0xb6, 0x28, 0xee, 0x42, 0x39, 0x8d, 0x7c, 0x16, 0xc3, 0x59, 0x08, 0xb2, 0x75, 0xe2, 0x75,
	0x78, 0xe5, 0x3f, 0x43, 0x06, 0x8a, 0xf7, 0x70, 0xd9, 0x87, 0x3d, 0x5a, 0x48, 0xbb, 0x4b, 0x50,
	0xc9, 0x91, 0x83, 0x02, 0xb9, 0xb8, 0x13, 0x6d, 0x77, 0x6e, 0xcd, 0xcf, 0xe1, 0xd6, 0x30, 0xff,
	0x87, 0x78, 0x0c, 0xeb, 0xd9, 0x92, 0xb9, 0x10, 0xf8, 0x8e, 0xd7, 0xcf, 0x74, 0xe4, 0x15, 0x6c,
	0xf3, 0x77, 0x38, 0x85, 0x8b, 0x16, 0x94, 0x0c, 0x9d, 0xba, 0xd7, 0xe5, 0x59, 0x8f, 0x82, 0x7f,
	0xbe, 0x13, 0x82, 0xe9, 0x20, 0x82, 0x90, 0x4e, 0x43, 0xdf, 0xa3, 0x7e, 0x0d, 0x4a, 0xee, 0xb7,
	0xd8, 0x85, 0x35, 0x77, 0xf2, 0x79, 0xfc, 0xdc, 0x71, 0xa3, 0x29, 0x87, 0x65, 0x77, 0x6c, 0xc9,
	0x17, 0x1f, 0xa1, 0xa1, 0xa8, 0x20, 0x30, 0xbf, 0xb0, 0x3e, 0x1c, 0xf4, 0xc9, 0xca, 0x4f, 0xd3,
	0xba, 0xfc, 0xcf, 0x3b, 0xed, 0x78, 0x2d, 0xbd, 0x37, 0x1e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x26, 0x9d, 0x02, 0x12, 0x05, 0x00, 0x00,
}
