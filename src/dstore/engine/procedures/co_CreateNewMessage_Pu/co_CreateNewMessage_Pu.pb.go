// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_CreateNewMessage_Pu.proto
// DO NOT EDIT!

/*
Package co_CreateNewMessage_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_CreateNewMessage_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_CreateNewMessage_Pu

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
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                    *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                bool                        `protobuf:"varint,1003,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	ToCommunityMemberId            *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=to_community_member_id,json=toCommunityMemberId" json:"to_community_member_id,omitempty"`
	ToCommunityMemberIdNull        bool                        `protobuf:"varint,1004,opt,name=to_community_member_id_null,json=toCommunityMemberIdNull" json:"to_community_member_id_null,omitempty"`
	Message                        *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	MessageNull                    bool                        `protobuf:"varint,1005,opt,name=message_null,json=messageNull" json:"message_null,omitempty"`
	MessageStatus                  *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=message_status,json=messageStatus" json:"message_status,omitempty"`
	MessageStatusNull              bool                        `protobuf:"varint,1006,opt,name=message_status_null,json=messageStatusNull" json:"message_status_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1007,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetToCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToCommunityMemberId
	}
	return nil
}

func (m *Parameters) GetMessage() *dstore_values.StringValue {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Parameters) GetMessageStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.MessageStatus
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_CreateNewMessage_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_CreateNewMessage_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_CreateNewMessage_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x5b, 0x8b, 0xd3, 0x40,
	0x14, 0x66, 0xb7, 0xf6, 0xe2, 0x69, 0xdd, 0x75, 0xa7, 0xb2, 0x1b, 0xdb, 0x75, 0xd1, 0xaa, 0x20,
	0x08, 0xa9, 0xac, 0x8a, 0xbe, 0xa8, 0xb0, 0xab, 0x0f, 0x15, 0x5a, 0x6a, 0x04, 0x41, 0x11, 0x42,
	0xb6, 0x3d, 0x5b, 0x02, 0x4d, 0xa6, 0xce, 0x4c, 0x2c, 0xfe, 0x0b, 0x7f, 0x8d, 0xff, 0xc9, 0xfb,
	0xa3, 0xaf, 0x4e, 0xe6, 0x4c, 0x2f, 0x89, 0x5d, 0xab, 0x2f, 0x21, 0x33, 0xf3, 0xdd, 0x32, 0x39,
	0xe7, 0xc0, 0xfd, 0xa1, 0x54, 0x5c, 0x60, 0x1b, 0xe3, 0x51, 0x18, 0x63, 0x7b, 0x22, 0xf8, 0x00,
	0x87, 0x89, 0x40, 0xd9, 0x1e, 0x70, 0xff, 0x58, 0x60, 0xa0, 0xb0, 0x87, 0xd3, 0x2e, 0x4a, 0x19,
	0x8c, 0xd0, 0xef, 0x27, 0xae, 0x46, 0x28, 0xce, 0x6e, 0x10, 0xcd, 0x25, 0x9a, 0xbb, 0x1a, 0xdb,
	0xa8, 0x5b, 0xf1, 0xf7, 0xc1, 0x38, 0x41, 0x49, 0xd4, 0xc6, 0xe5, 0xac, 0x23, 0x0a, 0xc1, 0x85,
	0x3d, 0x6a, 0x66, 0x8f, 0x22, 0x52, 0xb2, 0x87, 0xd7, 0xf3, 0x87, 0x2a, 0x08, 0xe3, 0x53, 0x2e,
	0xa2, 0x40, 0x85, 0x3c, 0x26, 0x50, 0xeb, 0x57, 0x09, 0xa0, 0x1f, 0x88, 0x40, 0x9f, 0xa2, 0x90,
	0xec, 0x01, 0x9c, 0x4f, 0xe2, 0xf0, 0x5d, 0x82, 0x7e, 0x38, 0x74, 0x36, 0xae, 0x6e, 0xdc, 0xaa,
	0x1e, 0x36, 0x5c, 0x1b, 0xdd, 0x86, 0x92, 0x4a, 0x84, 0xf1, 0xe8, 0x55, 0xba, 0xf0, 0x2a, 0x04,
	0xee, 0x0c, 0xd9, 0x4d, 0xd8, 0x9a, 0x13, 0xfd, 0x38, 0x19, 0x8f, 0x9d, 0xcf, 0x65, 0x4d, 0xaf,
	0x78, 0xb5, 0x19, 0xa4, 0xa7, 0x37, 0xd9, 0x5b, 0xd8, 0x9f, 0x68, 0x1f, 0x1e, 0x6b, 0x18, 0xc6,
	0x2a, 0x3c, 0x0d, 0x07, 0x26, 0x8d, 0x4f, 0xe2, 0xce, 0xe6, 0x5a, 0xcb, 0x06, 0xf1, 0x3b, 0x19,
	0xba, 0x39, 0x92, 0xec, 0x39, 0x5c, 0xfb, 0x9b, 0x3a, 0xe5, 0xfa, 0x42, 0xb9, 0x0e, 0xce, 0xd6,
	0x31, 0x49, 0x1f, 0x43, 0x6d, 0xc0, 0xa3, 0x48, 0xa7, 0x57, 0x1f, 0xd2, 0xcb, 0x28, 0x98, 0x64,
	0xcd, 0x5c, 0xb2, 0x30, 0x56, 0x38, 0x42, 0x41, 0xd1, 0xaa, 0x73, 0x82, 0xbe, 0x90, 0xdb, 0xb0,
	0xb3, 0xcc, 0x27, 0xef, 0xaf, 0xe4, 0xbd, 0xbd, 0x04, 0x34, 0x66, 0x7d, 0xd8, 0x55, 0xdc, 0x5f,
	0xe0, 0x23, 0x8c, 0x4e, 0x50, 0xa4, 0xb6, 0xe7, 0xd6, 0xdb, 0xd6, 0x15, 0x3f, 0x9e, 0x31, 0xbb,
	0x86, 0xa8, 0xed, 0x1f, 0x41, 0x73, 0xb5, 0x22, 0x05, 0xf9, 0x46, 0x41, 0xf6, 0x56, 0x50, 0x4d,
	0xa0, 0x7b, 0x50, 0xb6, 0xc5, 0xe4, 0x14, 0xd7, 0xfe, 0x92, 0x19, 0x94, 0xb5, 0xa0, 0x66, 0x5f,
	0xc9, 0xe5, 0x3b, 0xb9, 0x54, 0xed, 0xa6, 0x51, 0x3e, 0x82, 0xad, 0x19, 0x46, 0xaa, 0x40, 0x25,
	0xd2, 0x29, 0xad, 0xff, 0xc4, 0x0b, 0x96, 0xf2, 0xd2, 0x30, 0x58, 0x1b, 0xea, 0x59, 0x0d, 0xb2,
	0xfb, 0x41, 0x76, 0x3b, 0x19, 0xb0, 0x31, 0x7d, 0x01, 0x7b, 0x12, 0x27, 0xba, 0xcc, 0xb5, 0x83,
	0x1f, 0xda, 0xf2, 0x48, 0xab, 0x42, 0x3a, 0xe5, 0xb5, 0x9f, 0x77, 0x69, 0x4e, 0xed, 0x50, 0xb9,
	0xe8, 0x6d, 0xc9, 0x9e, 0xc0, 0xfe, 0x19, 0x92, 0x14, 0xe6, 0x27, 0x85, 0x71, 0x56, 0x91, 0xd3,
	0x4c, 0xad, 0x4f, 0x9b, 0x50, 0xf1, 0x50, 0x4e, 0x78, 0x2c, 0x91, 0xdd, 0x81, 0xa2, 0xe9, 0xeb,
	0x7c, 0xcf, 0xd9, 0x71, 0x41, 0x3d, 0xff, 0x2c, 0x7d, 0x7a, 0x04, 0x64, 0xaf, 0xe1, 0x62, 0xda,
	0xd1, 0xfe, 0x52, 0x4b, 0xeb, 0xee, 0x29, 0x68, 0xb2, 0x9b, 0x23, 0xe7, 0x1b, 0xbf, 0xab, 0xd7,
	0x9d, 0xc5, 0xda, 0xdb, 0x8e, 0xb2, 0x1b, 0xec, 0xe1, 0xe2, 0xe7, 0x17, 0x8c, 0xe2, 0xc1, 0x1f,
	0x8a, 0x34, 0x67, 0xec, 0xe4, 0x5a, 0x14, 0xc0, 0x53, 0x28, 0x08, 0x3e, 0xd5, 0x45, 0x9b, 0xb2,
	0x0e, 0xdd, 0x7f, 0x99, 0x79, 0xee, 0xec, 0x0e, 0x5c, 0x8f, 0x4f, 0xbd, 0x94, 0xde, 0xb8, 0x02,
	0x05, 0xfd, 0xce, 0x76, 0xa1, 0xa4, 0x57, 0x69, 0x13, 0x7c, 0xec, 0xe9, 0x5b, 0x29, 0x7a, 0x45,
	0xbd, 0xec, 0x0c, 0x8f, 0x7a, 0xd0, 0x0c, 0x79, 0x4e, 0x7b, 0x31, 0x86, 0xdf, 0xb4, 0xff, 0x73,
	0x40, 0x9f, 0x94, 0xcc, 0x24, 0xbc, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xad, 0xfd, 0x25,
	0xda, 0x05, 0x00, 0x00,
}