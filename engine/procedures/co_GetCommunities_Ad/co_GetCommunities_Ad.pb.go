// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetCommunities_Ad.proto
// DO NOT EDIT!

/*
Package co_GetCommunities_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetCommunities_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetCommunities_Ad

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
	CommunityId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull   bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunityName     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	CommunityNameNull bool                        `protobuf:"varint,1002,opt,name=community_name_null,json=communityNameNull" json:"community_name_null,omitempty"`
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

func (m *Parameters) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
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
	RowId         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CommunityId   *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityName *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Response_Row) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetCommunities_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetCommunities_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetCommunities_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_GetCommunities_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdf, 0xea, 0xd3, 0x30,
	0x14, 0xa6, 0xd6, 0xfd, 0x7e, 0x23, 0x53, 0xe7, 0x32, 0x90, 0xda, 0x81, 0xc8, 0xbc, 0x11, 0x84,
	0x74, 0x38, 0x10, 0xaf, 0x94, 0x4d, 0x44, 0x76, 0xb1, 0x22, 0x01, 0x05, 0xbd, 0x29, 0x75, 0x3d,
	0x96, 0xe0, 0x9a, 0x8c, 0x24, 0x75, 0xf8, 0x16, 0xea, 0x23, 0xf8, 0x18, 0xbe, 0x91, 0x82, 0xef,
	0x60, 0xda, 0x64, 0x7f, 0xda, 0x09, 0xfe, 0x76, 0xb3, 0x71, 0x7a, 0xbe, 0xef, 0x7c, 0x27, 0xdf,
	0x97, 0xa0, 0x69, 0xa6, 0xb4, 0x90, 0x10, 0x01, 0xcf, 0x19, 0x87, 0x68, 0x23, 0xc5, 0x0a, 0xb2,
	0x52, 0x82, 0x8a, 0x56, 0x22, 0x79, 0x05, 0xfa, 0x85, 0x28, 0x8a, 0x92, 0x33, 0xcd, 0x40, 0x25,
	0xb3, 0x8c, 0x98, 0xbe, 0x16, 0x78, 0x6c, 0x49, 0xc4, 0x92, 0xc8, 0xbf, 0x90, 0xe1, 0xd0, 0x0d,
	0xfe, 0x9c, 0xae, 0x4b, 0x50, 0x96, 0x18, 0xde, 0x6d, 0xaa, 0x81, 0x94, 0x42, 0xba, 0xd6, 0xa8,
	0xd9, 0x2a, 0x40, 0xa9, 0x34, 0x07, 0xd7, 0x7c, 0xd0, 0x6e, 0xea, 0x94, 0xf1, 0x8f, 0x42, 0x16,
	0xa9, 0x66, 0x82, 0x5b, 0xd0, 0xf8, 0x8f, 0x87, 0xd0, 0xeb, 0x54, 0xa6, 0xa6, 0x0b, 0x52, 0xe1,
	0x67, 0xe8, 0xc6, 0xca, 0xad, 0xf4, 0x25, 0x61, 0x59, 0xe0, 0xdd, 0xf7, 0x1e, 0xf6, 0x1e, 0x8f,
	0x88, 0xdb, 0xdd, 0xed, 0xc5, 0xb8, 0x86, 0x1c, 0xe4, 0xdb, 0xaa, 0xa2, 0xbd, 0x3d, 0x61, 0x91,
	0xe1, 0x47, 0x68, 0x70, 0xcc, 0x4f, 0x78, 0xb9, 0x5e, 0x07, 0xbf, 0x2e, 0xcd, 0x94, 0x2e, 0xed,
	0x1f, 0x01, 0x63, 0xf3, 0x1d, 0xcf, 0xd0, 0xad, 0x03, 0x98, 0x9b, 0x15, 0x82, 0x6b, 0xb5, 0x5c,
	0xd8, 0x92, 0x53, 0x5a, 0x32, 0x9e, 0x5b, 0xb5, 0x9b, 0x7b, 0x46, 0x6c, 0x08, 0x38, 0x42, 0xc3,
	0xe6, 0x08, 0xab, 0xf8, 0xdb, 0x2a, 0x0e, 0x1a, 0xe0, 0x4a, 0x73, 0xfc, 0xd3, 0x47, 0x5d, 0x0a,
	0x6a, 0x23, 0xb8, 0x02, 0x3c, 0x41, 0x9d, 0xda, 0x4d, 0x77, 0xcc, 0xbd, 0xae, 0x8b, 0xc8, 0x3a,
	0xfd, 0xb2, 0xfa, 0xa5, 0x16, 0x88, 0xdf, 0xa1, 0xdb, 0x95, 0x8f, 0xc9, 0x91, 0x91, 0x66, 0x69,
	0xdf, 0x90, 0x49, 0x8b, 0xdc, 0xb6, 0x7b, 0x69, 0xea, 0xc5, 0xa1, 0xa6, 0xfd, 0xa2, 0xf9, 0x01,
	0x3f, 0x45, 0x97, 0x2e, 0xbf, 0xc0, 0xaf, 0x27, 0xde, 0x3b, 0x99, 0x68, 0xd3, 0x5d, 0xda, 0x7f,
	0xba, 0x83, 0xe3, 0x39, 0xf2, 0xa5, 0xd8, 0x06, 0xd7, 0x6b, 0xd6, 0x84, 0xfc, 0xff, 0x9e, 0x91,
	0x9d, 0x03, 0x84, 0x8a, 0x2d, 0xad, 0xc8, 0xe1, 0x0f, 0x0f, 0xf9, 0xa6, 0xc0, 0x77, 0xd0, 0x85,
	0x29, 0xab, 0xe8, 0xbf, 0xc6, 0xc6, 0x94, 0x0e, 0xed, 0x98, 0xd2, 0x04, 0xfb, 0xbc, 0x75, 0x31,
	0xbe, 0xc5, 0x67, 0xde, 0x8c, 0xf9, 0x49, 0xd8, 0xdf, 0xe3, 0x33, 0xd3, 0x9e, 0xbf, 0x41, 0x23,
	0x26, 0x5a, 0xe7, 0x3b, 0x3c, 0xbe, 0xf7, 0x4f, 0x72, 0xa1, 0xb2, 0x4f, 0xbb, 0x7e, 0x76, 0xd5,
	0xf7, 0xf9, 0xe1, 0xa2, 0x7e, 0x0a, 0xd3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x81, 0x7f,
	0x90, 0xd7, 0x03, 0x00, 0x00,
}
