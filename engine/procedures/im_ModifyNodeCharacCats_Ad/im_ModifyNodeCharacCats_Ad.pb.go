// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyNodeCharacCats_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyNodeCharacCats_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyNodeCharacCats_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyNodeCharacCats_Ad

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
	NodeCharacCategoryId      *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=node_charac_category_id,json=nodeCharacCategoryId" json:"node_charac_category_id,omitempty"`
	NodeCharacCategoryIdNull  bool                        `protobuf:"varint,1001,opt,name=node_charac_category_id_null,json=nodeCharacCategoryIdNull" json:"node_charac_category_id_null,omitempty"`
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

func (m *Parameters) GetNodeCharacCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacCategoryId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyNodeCharacCats_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyNodeCharacCats_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyNodeCharacCats_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_ModifyNodeCharacCats_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x1b, 0xd2, 0x56, 0x03, 0x6a, 0x61, 0x1b, 0x88, 0x93, 0x14, 0x84, 0xca, 0x4b, 0xc5,
	0x83, 0xc3, 0x45, 0x42, 0x20, 0x40, 0x5c, 0x5a, 0x24, 0xf2, 0x90, 0x08, 0xb9, 0x12, 0x12, 0xf0,
	0x60, 0xb9, 0xde, 0x69, 0xb0, 0x48, 0x76, 0xa2, 0x5d, 0xa7, 0x55, 0xff, 0x82, 0xcf, 0xe0, 0x1b,
	0xf8, 0x23, 0x6e, 0xff, 0xc0, 0xac, 0xd7, 0xb9, 0x19, 0x47, 0xa5, 0x2f, 0x89, 0xc7, 0x73, 0xce,
	0x9c, 0xb3, 0xe3, 0x99, 0x85, 0x27, 0xd2, 0xa4, 0xa4, 0xb1, 0x8d, 0xaa, 0x9f, 0x28, 0x6c, 0x8f,
	0x34, 0xc5, 0x28, 0xc7, 0x1a, 0x4d, 0x3b, 0x19, 0x86, 0x5d, 0x92, 0xc9, 0xf1, 0x59, 0x8f, 0x24,
	0xee, 0x7f, 0x8e, 0x74, 0x14, 0xef, 0x47, 0xa9, 0x09, 0x5f, 0x49, 0x9f, 0x51, 0x29, 0x89, 0x3d,
	0x47, 0xf5, 0x1d, 0xd5, 0x5f, 0x8e, 0x6f, 0x6e, 0xe7, 0x22, 0x27, 0xd1, 0x60, 0x8c, 0xc6, 0xd1,
	0x9b, 0x8d, 0x45, 0x65, 0xd4, 0x9a, 0x74, 0x9e, 0x6a, 0x2d, 0xa6, 0x86, 0x68, 0x4c, 0xd4, 0xc7,
	0x3c, 0x79, 0xa7, 0x98, 0x4c, 0xa3, 0x44, 0x1d, 0x93, 0x1e, 0x46, 0x69, 0x42, 0xca, 0x81, 0x76,
	0xbf, 0x55, 0x01, 0xde, 0xb1, 0x05, 0xce, 0xa2, 0x36, 0x22, 0x80, 0xba, 0x62, 0x57, 0x61, 0x9c,
	0xd9, 0x0a, 0xe3, 0x28, 0xc5, 0x3e, 0xe9, 0xb3, 0x30, 0x91, 0xde, 0xca, 0xed, 0x95, 0xbd, 0xcb,
	0x0f, 0x5a, 0x7e, 0x7e, 0x98, 0xdc, 0x62, 0xa2, 0x18, 0x81, 0xfa, 0xbd, 0x8d, 0x82, 0x9a, 0x9a,
	0x3f, 0x51, 0xc6, 0xec, 0x48, 0xf1, 0x02, 0x76, 0x96, 0xd4, 0x0c, 0xd5, 0x78, 0x30, 0xf0, 0x7e,
	0xac, 0x73, 0xe5, 0x8d, 0xc0, 0x2b, 0x23, 0xf7, 0x18, 0x20, 0xba, 0x50, 0x9b, 0x92, 0x24, 0x9a,
	0x58, 0x27, 0x23, 0x7b, 0x02, 0x6f, 0x35, 0x73, 0xd4, 0x2c, 0x38, 0x32, 0xa9, 0x4e, 0x54, 0xdf,
	0x19, 0xda, 0x9e, 0xf0, 0x0e, 0x66, 0x34, 0xf1, 0x14, 0x1a, 0x65, 0xe5, 0x9c, 0x99, 0x9f, 0xce,
	0x4c, 0xbd, 0x84, 0x98, 0x79, 0x39, 0x84, 0xfa, 0x48, 0xa3, 0xc4, 0x98, 0x5b, 0x4d, 0x7a, 0xa1,
	0x41, 0x95, 0xf3, 0x1b, 0x74, 0x7d, 0x8e, 0x3b, 0xd7, 0xa1, 0x97, 0xb0, 0xb3, 0xa4, 0xa8, 0x33,
	0xf5, 0xcb, 0x99, 0x6a, 0x94, 0xb2, 0x33, 0x5b, 0xcf, 0xe1, 0xca, 0x90, 0x4e, 0x30, 0xe4, 0x4c,
	0x1a, 0x2a, 0xf2, 0x2e, 0x9d, 0xef, 0x05, 0x2c, 0xe1, 0x90, 0xf1, 0x3d, 0x12, 0x77, 0xe1, 0xda,
	0x3c, 0xdd, 0xa9, 0xfe, 0x76, 0xaa, 0x9b, 0x33, 0x5c, 0x26, 0x75, 0x00, 0x5b, 0x12, 0x07, 0x3c,
	0x2e, 0x53, 0x9f, 0x5e, 0xb5, 0x54, 0xed, 0x88, 0x68, 0x80, 0x91, 0x72, 0x6a, 0x9b, 0x8e, 0x33,
	0xb1, 0x2d, 0xee, 0x43, 0xad, 0x50, 0xc5, 0x89, 0xfe, 0x71, 0xa2, 0x62, 0x11, 0x6e, 0x85, 0x77,
	0xbf, 0xaf, 0xc2, 0x46, 0x80, 0x66, 0x44, 0xca, 0xa0, 0xb8, 0x07, 0xd5, 0x6c, 0x11, 0xf2, 0xb1,
	0x9c, 0x0e, 0x41, 0xbe, 0x63, 0x6e, 0x49, 0xde, 0xd8, 0xdf, 0xc0, 0x01, 0xc5, 0x07, 0xb8, 0x6a,
	0x57, 0x20, 0x9c, 0xdb, 0x01, 0x9e, 0xa0, 0x0a, 0x93, 0xfd, 0x02, 0xb9, 0xb8, 0x29, 0x5d, 0x8e,
	0x3b, 0xb3, 0x38, 0xd8, 0x1a, 0x2e, 0xbe, 0x10, 0x8f, 0x61, 0x3d, 0x5f, 0x3d, 0x1e, 0x02, 0x5b,
	0xf1, 0xd6, 0x3f, 0x15, 0xdd, 0x62, 0x76, 0xdd, 0x7f, 0x30, 0x81, 0x8b, 0xb7, 0x50, 0xd1, 0x74,
	0xca, 0x9f, 0xcb, 0xb2, 0x1e, 0xf9, 0xff, 0x7b, 0x51, 0xf8, 0x93, 0x3e, 0xf8, 0x01, 0x9d, 0x06,
	0xb6, 0x44, 0xf3, 0x26, 0x54, 0xf8, 0x59, 0xdc, 0x80, 0x35, 0x8e, 0xec, 0x38, 0x7e, 0xed, 0x71,
	0x67, 0xaa, 0x41, 0x95, 0xc3, 0x8e, 0x7c, 0xfd, 0x09, 0x5a, 0x09, 0x15, 0xea, 0xcf, 0xee, 0xb0,
	0x8f, 0xcf, 0xfa, 0x64, 0xe4, 0x97, 0x49, 0x5e, 0x5e, 0xec, 0x9a, 0x3b, 0x5a, 0xcb, 0xee, 0x92,
	0x87, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x14, 0xe5, 0xe6, 0x24, 0x05, 0x00, 0x00,
}
