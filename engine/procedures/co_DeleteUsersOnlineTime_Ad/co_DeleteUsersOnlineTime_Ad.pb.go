// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_DeleteUsersOnlineTime_Ad.proto
// DO NOT EDIT!

/*
Package co_DeleteUsersOnlineTime_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_DeleteUsersOnlineTime_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_DeleteUsersOnlineTime_Ad

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
	CommunityId           *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull       bool                          `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunityMemberId     *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	CommunityMemberIdNull bool                          `protobuf:"varint,1002,opt,name=community_member_id_null,json=communityMemberIdNull" json:"community_member_id_null,omitempty"`
	MaxOutDate            *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=max_out_date,json=maxOutDate" json:"max_out_date,omitempty"`
	MaxOutDateNull        bool                          `protobuf:"varint,1003,opt,name=max_out_date_null,json=maxOutDateNull" json:"max_out_date_null,omitempty"`
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

func (m *Parameters) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Parameters) GetMaxOutDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.MaxOutDate
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_DeleteUsersOnlineTime_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_DeleteUsersOnlineTime_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_DeleteUsersOnlineTime_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_DeleteUsersOnlineTime_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x6d, 0x6b, 0x13, 0x41,
	0x10, 0xa6, 0x89, 0x69, 0xcb, 0xb4, 0x58, 0xbb, 0x45, 0x39, 0x13, 0x2a, 0x52, 0xbf, 0xf8, 0x02,
	0x1b, 0xd1, 0x0f, 0x8a, 0xa0, 0xa2, 0xd4, 0x0f, 0x41, 0xd2, 0xca, 0xa2, 0x82, 0x22, 0x1c, 0xdb,
	0xec, 0x18, 0x0e, 0x6f, 0x77, 0xc3, 0xee, 0x9e, 0xd5, 0x7f, 0xe1, 0xff, 0xf1, 0x17, 0xa9, 0xf8,
	0x1f, 0x9c, 0xbb, 0xbd, 0x78, 0xc9, 0x45, 0xd4, 0x7e, 0xb9, 0xbb, 0xd9, 0xe7, 0x79, 0xe6, 0x99,
	0x9b, 0x99, 0x85, 0x07, 0xca, 0x07, 0xeb, 0x70, 0x88, 0x66, 0x9a, 0x19, 0x1c, 0xce, 0x9c, 0x9d,
	0xa0, 0x2a, 0x1c, 0xfa, 0xe1, 0xc4, 0xa6, 0x87, 0x98, 0x63, 0xc0, 0x57, 0x1e, 0x9d, 0x3f, 0x36,
	0x39, 0xe1, 0x2f, 0x33, 0x8d, 0xe9, 0x13, 0xc5, 0x89, 0x16, 0x2c, 0xbb, 0x11, 0xb5, 0x3c, 0x6a,
	0xf9, 0x5f, 0x04, 0xfd, 0xbd, 0xda, 0xe6, 0xa3, 0xcc, 0x0b, 0xf4, 0x51, 0xdf, 0xbf, 0xbc, 0xec,
	0x8d, 0xce, 0x59, 0x57, 0x43, 0x83, 0x65, 0x48, 0xa3, 0xf7, 0x72, 0x8a, 0x35, 0x78, 0xad, 0x0d,
	0x06, 0x99, 0x99, 0xf7, 0xd6, 0x69, 0x19, 0x32, 0x6b, 0x22, 0xe9, 0xe0, 0x67, 0x07, 0xe0, 0x85,
	0x74, 0x92, 0x50, 0xaa, 0x86, 0x3d, 0x82, 0xed, 0x89, 0xd5, 0xba, 0x30, 0x59, 0xf8, 0x9c, 0x66,
	0x2a, 0x59, 0xbb, 0xba, 0x76, 0x7d, 0xeb, 0xce, 0x80, 0xd7, 0xbf, 0x50, 0xd7, 0x95, 0x99, 0x80,
	0x53, 0x74, 0xaf, 0xcb, 0x48, 0x6c, 0xfd, 0x16, 0x8c, 0x14, 0xbb, 0x05, 0xbb, 0x8b, 0xfa, 0xd4,
	0x14, 0x79, 0x9e, 0x7c, 0xdb, 0xa0, 0x2c, 0x9b, 0x62, 0x67, 0x81, 0x78, 0x44, 0xe7, 0xec, 0x39,
	0xec, 0x35, 0x64, 0x8d, 0xfa, 0x04, 0x5d, 0xe9, 0xd9, 0xf9, 0xb7, 0x67, 0x63, 0x32, 0xae, 0x64,
	0xe4, 0x7c, 0x1f, 0x92, 0x3f, 0x24, 0x8b, 0x05, 0x7c, 0x8f, 0x05, 0x5c, 0x5c, 0x51, 0x55, 0x65,
	0x3c, 0x86, 0x6d, 0x2d, 0x3f, 0xa5, 0xb6, 0x08, 0xa9, 0x92, 0x01, 0x93, 0x6e, 0xe5, 0xbf, 0xdf,
	0xf2, 0x0f, 0x34, 0x22, 0x1f, 0xa4, 0x9e, 0xc5, 0x0a, 0x80, 0x24, 0xc7, 0x45, 0x38, 0x24, 0x01,
	0xbb, 0x09, 0xbb, 0x8b, 0x09, 0xa2, 0xe7, 0x8f, 0xe8, 0x79, 0xbe, 0xe1, 0x95, 0x66, 0x07, 0x5f,
	0x3b, 0xb0, 0x29, 0xd0, 0xcf, 0xac, 0xf1, 0xc8, 0x6e, 0x43, 0xaf, 0x9a, 0x66, 0xdd, 0xe6, 0x3e,
	0x5f, 0xde, 0x94, 0x38, 0xe9, 0x67, 0xe5, 0x53, 0x44, 0x22, 0x7b, 0x03, 0x17, 0xca, 0x39, 0xa6,
	0x0b, 0x83, 0xa4, 0x7e, 0x75, 0x49, 0xcc, 0x5b, 0xe2, 0xf6, 0xb8, 0xc7, 0x14, 0x8f, 0x9a, 0x58,
	0xec, 0xe8, 0xe5, 0x03, 0x6a, 0xe0, 0x46, 0xbd, 0x3f, 0xd4, 0x81, 0x32, 0xe3, 0x95, 0x95, 0x8c,
	0x71, 0xbb, 0xc6, 0xf1, 0x2d, 0xe6, 0x74, 0x36, 0x82, 0xae, 0xb3, 0xa7, 0xc9, 0xb9, 0x4a, 0x75,
	0x8f, 0xff, 0xf7, 0xba, 0xf3, 0x79, 0x23, 0xb8, 0xb0, 0xa7, 0xa2, 0xcc, 0xd1, 0xdf, 0x87, 0x2e,
	0x7d, 0xb3, 0x4b, 0xb0, 0x4e, 0x51, 0xb9, 0x0c, 0x5f, 0x8e, 0xa8, 0x35, 0x3d, 0xd1, 0xa3, 0x70,
	0xa4, 0x9e, 0xbe, 0x83, 0x41, 0x66, 0x5b, 0x06, 0xcd, 0x5d, 0x7c, 0xfb, 0x70, 0x6a, 0xbd, 0xfa,
	0x30, 0xc7, 0xd5, 0x19, 0xaf, 0xeb, 0xc9, 0x7a, 0x75, 0x25, 0xee, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x8e, 0x2d, 0x03, 0xda, 0xed, 0x03, 0x00, 0x00,
}
