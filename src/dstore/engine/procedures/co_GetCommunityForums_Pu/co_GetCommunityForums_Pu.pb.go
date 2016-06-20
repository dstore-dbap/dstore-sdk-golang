// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetCommunityForums_Pu.proto
// DO NOT EDIT!

/*
Package co_GetCommunityForums_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetCommunityForums_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetCommunityForums_Pu

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
	CommunityId                    *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1002,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1003,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1004,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
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
	Error             *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation   []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message           []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row               []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	CommunityMemberId *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
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

func (m *Response) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

type Response_Row struct {
	RowId         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CommunityId   *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	ForumId       *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumName     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=forum_name,json=forumName" json:"forum_name,omitempty"`
	CommunityName *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	LanguageId    *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	AccessLevel   *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=access_level,json=accessLevel" json:"access_level,omitempty"`
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

func (m *Response_Row) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Response_Row) GetAccessLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessLevel
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetCommunityForums_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetCommunityForums_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetCommunityForums_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0x6b, 0x6b, 0x13, 0x4f,
	0x14, 0xc6, 0xe9, 0x3f, 0xbd, 0xa4, 0xa7, 0xfd, 0xb7, 0x3a, 0x15, 0x5d, 0x93, 0x52, 0xb4, 0xa2,
	0x08, 0xc2, 0xc6, 0x1b, 0x44, 0x04, 0x2d, 0x54, 0xac, 0x44, 0x4d, 0x88, 0xfb, 0x42, 0x50, 0x84,
	0x65, 0xbb, 0x39, 0x09, 0x0b, 0xbb, 0x33, 0x71, 0x66, 0xb7, 0xc5, 0x77, 0x7e, 0x04, 0xef, 0xfd,
	0x8c, 0x5e, 0x3e, 0x84, 0x67, 0x66, 0x36, 0xb7, 0xd5, 0x98, 0xf4, 0x4d, 0xcb, 0xec, 0x9c, 0xdf,
	0x73, 0x9e, 0x9d, 0x79, 0xce, 0x06, 0xea, 0x1d, 0x95, 0x0a, 0x89, 0x35, 0xe4, 0xbd, 0x88, 0x63,
	0xad, 0x2f, 0x45, 0x88, 0x9d, 0x4c, 0xa2, 0xaa, 0x85, 0xc2, 0x7f, 0x82, 0xe9, 0x23, 0x91, 0x24,
	0x19, 0x8f, 0xd2, 0x77, 0x07, 0x42, 0x66, 0x89, 0xf2, 0xdb, 0x99, 0x4b, 0x35, 0xa9, 0x60, 0xd7,
	0x2c, 0xe8, 0x5a, 0xd0, 0x9d, 0x56, 0x5d, 0xd9, 0xca, 0x1b, 0x1c, 0x05, 0x71, 0x86, 0xca, 0xc2,
	0x95, 0x8b, 0x93, 0x5d, 0x51, 0x4a, 0x21, 0xf3, 0xad, 0xea, 0xe4, 0x56, 0x82, 0x4a, 0x05, 0x3d,
	0xcc, 0x37, 0xaf, 0x14, 0x37, 0xd3, 0x20, 0xe2, 0x5d, 0x21, 0x93, 0x20, 0x8d, 0x04, 0xb7, 0x45,
	0xbb, 0x27, 0x8b, 0x00, 0xed, 0x40, 0x06, 0xb4, 0x8b, 0x52, 0xb1, 0x87, 0xb0, 0x1e, 0x0e, 0x6c,
	0xf9, 0x51, 0xc7, 0x59, 0xb8, 0xb4, 0x70, 0x7d, 0xed, 0x76, 0xd5, 0xcd, 0xfd, 0xe7, 0xbe, 0x22,
	0x9e, 0x62, 0x0f, 0xe5, 0x4b, 0xbd, 0xf2, 0xd6, 0x86, 0x40, 0xa3, 0xc3, 0x6e, 0xc0, 0xd9, 0x71,
	0xde, 0xe7, 0x59, 0x1c, 0x3b, 0xdf, 0x57, 0x48, 0xa5, 0xec, 0x6d, 0x8e, 0x15, 0xb6, 0xe8, 0x39,
	0xab, 0xc3, 0x2a, 0x2d, 0xdf, 0x66, 0xa8, 0x3b, 0xfd, 0x67, 0x3a, 0x55, 0x0a, 0x9d, 0x54, 0x2a,
	0x23, 0xde, 0xb3, 0x8d, 0xca, 0xb6, 0x98, 0xba, 0x5c, 0x85, 0x8d, 0x21, 0x68, 0x5b, 0xfc, 0xb0,
	0x2d, 0xd6, 0x07, 0x25, 0x46, 0xff, 0x0d, 0x6c, 0xf7, 0xe9, 0xa5, 0x04, 0xa7, 0x32, 0xe4, 0x69,
	0xd4, 0x8d, 0x42, 0xf3, 0xea, 0xbe, 0x15, 0x77, 0x4a, 0x33, 0x5b, 0x56, 0x2c, 0xdf, 0x98, 0xc0,
	0xcd, 0x96, 0x62, 0x4f, 0xe1, 0xf2, 0xbf, 0xd4, 0xad, 0xaf, 0x9f, 0xd6, 0xd7, 0xce, 0x74, 0x1d,
	0xe3, 0xf4, 0x05, 0x5c, 0x50, 0xd8, 0xa7, 0x6b, 0x20, 0x23, 0x7e, 0x94, 0x2b, 0x6a, 0x21, 0xe5,
	0x2c, 0xce, 0x34, 0x79, 0x6e, 0x88, 0x36, 0x6c, 0x07, 0x7a, 0xac, 0xd8, 0x1e, 0x6c, 0x4f, 0x91,
	0xb4, 0xce, 0x7e, 0x59, 0x67, 0xce, 0xdf, 0x60, 0xed, 0x69, 0xf7, 0xfd, 0x32, 0x94, 0x3d, 0x54,
	0x7d, 0xc1, 0x15, 0xb2, 0x9b, 0xb0, 0x64, 0x72, 0x97, 0x07, 0x62, 0x68, 0x27, 0x0f, 0xb4, 0xcd,
	0xe4, 0x63, 0xfd, 0xd7, 0xb3, 0x85, 0xec, 0x15, 0x9c, 0xd1, 0x89, 0xf3, 0xc7, 0x22, 0x47, 0x77,
	0x5c, 0x22, 0xd8, 0x2d, 0xc0, 0xc5, 0x60, 0x36, 0x69, 0xdd, 0x18, 0xad, 0xbd, 0xcd, 0x64, 0xf2,
	0x01, 0xbb, 0x07, 0x2b, 0x79, 0xd2, 0xe9, 0x0a, 0xb5, 0xe2, 0xce, 0x1f, 0x8a, 0x76, 0x0e, 0x9a,
	0xf6, 0xbf, 0x37, 0x28, 0x67, 0x07, 0x50, 0x92, 0xe2, 0x98, 0xce, 0x54, 0x53, 0x77, 0xdd, 0xf9,
	0xa6, 0xd2, 0x1d, 0x9c, 0x82, 0xeb, 0x89, 0x63, 0x4f, 0x0b, 0xb0, 0x67, 0xb0, 0x35, 0x8a, 0x79,
	0x82, 0xc9, 0x21, 0x4a, 0x9d, 0x61, 0x9c, 0x3d, 0x2d, 0xa3, 0xf1, 0x68, 0x1a, 0xac, 0xd1, 0xa9,
	0x9c, 0x94, 0xa0, 0x44, 0xca, 0xec, 0x3c, 0x2c, 0x93, 0xb6, 0xd6, 0xf9, 0xd0, 0x22, 0xa1, 0x25,
	0x6f, 0x89, 0x96, 0x94, 0xf6, 0xbd, 0xc2, 0x4c, 0x7e, 0x6c, 0x9d, 0x72, 0x28, 0xeb, 0x50, 0xee,
	0xea, 0x97, 0xd1, 0xf0, 0xa7, 0x39, 0xe0, 0x15, 0x53, 0x4d, 0xe0, 0x7d, 0x00, 0x0b, 0x72, 0xfa,
	0x3e, 0x38, 0x9f, 0x5b, 0x33, 0xa3, 0xb8, 0x6a, 0xca, 0x5b, 0x54, 0xcd, 0xf6, 0x61, 0x63, 0xe4,
	0xda, 0xf0, 0x5f, 0x66, 0xf3, 0xff, 0x0f, 0x11, 0xa3, 0xf1, 0x00, 0xd6, 0xe2, 0x80, 0xf7, 0x32,
	0xba, 0x3a, 0xed, 0xfd, 0xeb, 0x1c, 0xde, 0x61, 0x00, 0xd8, 0x83, 0x0b, 0xc2, 0x90, 0xae, 0xde,
	0x8f, 0xf1, 0x08, 0x63, 0xe7, 0xdb, 0x3c, 0x07, 0x67, 0x89, 0xe7, 0x1a, 0xd8, 0x6f, 0x43, 0x35,
	0x12, 0x85, 0x94, 0x8c, 0x3e, 0xfa, 0xaf, 0x6f, 0x9d, 0xfa, 0xe7, 0xe0, 0x70, 0xd9, 0x7c, 0x75,
	0xef, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x06, 0xe4, 0x9f, 0x6d, 0x4a, 0x06, 0x00, 0x00,
}