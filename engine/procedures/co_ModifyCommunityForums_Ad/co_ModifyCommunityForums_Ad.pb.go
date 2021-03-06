// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_ModifyCommunityForums_Ad.proto
// DO NOT EDIT!

/*
Package co_ModifyCommunityForums_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_ModifyCommunityForums_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_ModifyCommunityForums_Ad

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
	CommunityId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	ForumId         *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumIdNull     bool                        `protobuf:"varint,1002,opt,name=forum_id_null,json=forumIdNull" json:"forum_id_null,omitempty"`
	SortNo          *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull      bool                        `protobuf:"varint,1003,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	Delete          *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull      bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_ModifyCommunityForums_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_ModifyCommunityForums_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_ModifyCommunityForums_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_ModifyCommunityForums_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x6b, 0x14, 0x31,
	0x10, 0xa7, 0x5d, 0xef, 0x0f, 0x73, 0x95, 0x6a, 0x04, 0x39, 0xef, 0x50, 0x6a, 0xfb, 0xa2, 0x08,
	0x39, 0xb1, 0xa2, 0x22, 0x28, 0xa8, 0x28, 0xdc, 0xc3, 0x1d, 0x92, 0x07, 0x41, 0x11, 0x96, 0xed,
	0x25, 0x77, 0x04, 0x77, 0x33, 0x25, 0xd9, 0xb5, 0xf4, 0x0b, 0xf8, 0xec, 0xf7, 0xf1, 0x13, 0xf9,
	0xe7, 0x43, 0x98, 0xec, 0x64, 0xed, 0xdd, 0x2a, 0x56, 0x5f, 0x76, 0x33, 0x99, 0xdf, 0x9f, 0x64,
	0x32, 0x03, 0x8f, 0xa5, 0x2b, 0xd1, 0xaa, 0x89, 0x32, 0x2b, 0x6d, 0xd4, 0xe4, 0xd8, 0xe2, 0x42,
	0xc9, 0xca, 0x2a, 0x37, 0x59, 0x60, 0x3a, 0x43, 0xa9, 0x97, 0xa7, 0x2f, 0xb0, 0x28, 0x2a, 0xa3,
	0xcb, 0xd3, 0x57, 0x68, 0xab, 0xc2, 0xa5, 0xcf, 0x24, 0xf7, 0xb0, 0x12, 0xd9, 0x6d, 0xe2, 0x72,
	0xe2, 0xf2, 0xbf, 0x10, 0x46, 0x57, 0xa2, 0xcd, 0xc7, 0x2c, 0xaf, 0x94, 0x23, 0xfe, 0xe8, 0xda,
	0xa6, 0xb7, 0xb2, 0x16, 0x6d, 0x4c, 0x8d, 0x37, 0x53, 0x85, 0x72, 0x2e, 0x5b, 0xa9, 0x98, 0x3c,
	0x68, 0x27, 0xcb, 0x4c, 0x9b, 0x25, 0xda, 0x22, 0x2b, 0x35, 0x1a, 0x02, 0xed, 0x7f, 0x4a, 0x00,
	0x5e, 0x67, 0x36, 0xf3, 0x59, 0x65, 0x1d, 0x7b, 0x0a, 0x3b, 0x8b, 0xe6, 0x58, 0xa9, 0x96, 0xc3,
	0xad, 0xbd, 0xad, 0x5b, 0x83, 0x7b, 0x63, 0x1e, 0xaf, 0x10, 0xcf, 0xa5, 0x4d, 0xa9, 0x56, 0xca,
	0xbe, 0x09, 0x91, 0x18, 0xfc, 0x22, 0x4c, 0x25, 0xbb, 0x03, 0x97, 0xd7, 0xf9, 0xa9, 0xa9, 0xf2,
	0x7c, 0xf8, 0xb5, 0xe7, 0x55, 0xfa, 0x62, 0x77, 0x0d, 0x38, 0xf7, 0xfb, 0xec, 0x01, 0xf4, 0x97,
	0xe1, 0xea, 0xc1, 0x68, 0xfb, 0x7c, 0xa3, 0x5e, 0x0d, 0xf6, 0x26, 0x07, 0x70, 0xb1, 0xe1, 0x91,
	0xc1, 0x37, 0x32, 0x18, 0x44, 0x40, 0x2d, 0x7e, 0x1f, 0x7a, 0x0e, 0x6d, 0x99, 0x1a, 0x1c, 0x26,
	0xe7, 0x6b, 0x77, 0x03, 0x76, 0x8e, 0xec, 0x26, 0xec, 0x44, 0x16, 0x29, 0x7f, 0x27, 0x65, 0xa0,
	0x74, 0x2d, 0x7c, 0x08, 0x5d, 0xa9, 0x72, 0x5f, 0xae, 0xe1, 0x85, 0x3f, 0xea, 0x1e, 0x21, 0xe6,
	0x2a, 0x33, 0x51, 0x97, 0xa0, 0x6c, 0x0f, 0x06, 0xb4, 0x22, 0xd9, 0x1f, 0x51, 0x96, 0xf6, 0x82,
	0xec, 0xfe, 0x97, 0x6d, 0xe8, 0x0b, 0xe5, 0x8e, 0xd1, 0x38, 0xc5, 0xee, 0x42, 0xa7, 0x7e, 0xe6,
	0x58, 0xff, 0x11, 0xdf, 0x6c, 0x21, 0x6a, 0x81, 0x97, 0xe1, 0x2b, 0x08, 0xc8, 0xde, 0xc2, 0xa5,
	0xf0, 0xc0, 0xe9, 0xda, 0x0b, 0xfb, 0x9a, 0x26, 0x9e, 0xcc, 0x5b, 0xe4, 0x76, 0x1f, 0xcc, 0x7c,
	0x3c, 0x3d, 0x8b, 0xc5, 0x6e, 0xb1, 0xb9, 0xc1, 0x1e, 0x41, 0x2f, 0x36, 0x96, 0xaf, 0x64, 0x50,
	0xbc, 0xf1, 0x9b, 0x22, 0xb5, 0xdd, 0x8c, 0xfe, 0xa2, 0x81, 0xb3, 0x29, 0x24, 0x16, 0x4f, 0x7c,
	0x9d, 0x02, 0xeb, 0x21, 0xff, 0xe7, 0x39, 0xe0, 0x4d, 0x21, 0xb8, 0xc0, 0x13, 0x11, 0x34, 0x46,
	0xd7, 0x21, 0xf1, 0x6b, 0x76, 0x15, 0xba, 0x3e, 0x0a, 0x0d, 0xf3, 0x79, 0xee, 0x4b, 0xd3, 0x11,
	0x1d, 0x1f, 0x4e, 0xe5, 0xf3, 0xf7, 0x30, 0xd6, 0xd8, 0x32, 0x38, 0x1b, 0xd2, 0x77, 0x4f, 0x56,
	0xe8, 0xe4, 0x87, 0x26, 0x2f, 0xff, 0x73, 0x8e, 0x8f, 0xba, 0xf5, 0xac, 0x1c, 0xfe, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xe1, 0xdb, 0x08, 0xfc, 0x06, 0x04, 0x00, 0x00,
}
