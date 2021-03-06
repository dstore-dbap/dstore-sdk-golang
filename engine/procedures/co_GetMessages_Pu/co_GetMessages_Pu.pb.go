// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetMessages_Pu.proto
// DO NOT EDIT!

/*
Package co_GetMessages_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetMessages_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetMessages_Pu

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
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                    *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                bool                        `protobuf:"varint,1003,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	GetNewMessages                 *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=get_new_messages,json=getNewMessages" json:"get_new_messages,omitempty"`
	GetNewMessagesNull             bool                        `protobuf:"varint,1004,opt,name=get_new_messages_null,json=getNewMessagesNull" json:"get_new_messages_null,omitempty"`
	DateAndTimeFormat              *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=date_and_time_format,json=dateAndTimeFormat" json:"date_and_time_format,omitempty"`
	DateAndTimeFormatNull          bool                        `protobuf:"varint,1005,opt,name=date_and_time_format_null,json=dateAndTimeFormatNull" json:"date_and_time_format_null,omitempty"`
	MessageNo                      *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=message_no,json=messageNo" json:"message_no,omitempty"`
	MessageNoNull                  bool                        `protobuf:"varint,1006,opt,name=message_no_null,json=messageNoNull" json:"message_no_null,omitempty"`
	MaxNumberOfRows                *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=max_number_of_rows,json=maxNumberOfRows" json:"max_number_of_rows,omitempty"`
	MaxNumberOfRowsNull            bool                        `protobuf:"varint,1007,opt,name=max_number_of_rows_null,json=maxNumberOfRowsNull" json:"max_number_of_rows_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1008,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetGetNewMessages() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetNewMessages
	}
	return nil
}

func (m *Parameters) GetDateAndTimeFormat() *dstore_values.IntegerValue {
	if m != nil {
		return m.DateAndTimeFormat
	}
	return nil
}

func (m *Parameters) GetMessageNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.MessageNo
	}
	return nil
}

func (m *Parameters) GetMaxNumberOfRows() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfRows
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
	RowId                       int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	FromCommunityMemberNickname *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=from_community_member_nickname,json=fromCommunityMemberNickname" json:"from_community_member_nickname,omitempty"`
	MessageStatus               *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=message_status,json=messageStatus" json:"message_status,omitempty"`
	Message                     *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=message" json:"message,omitempty"`
	MessageNo                   *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=message_no,json=messageNo" json:"message_no,omitempty"`
	MessageDateAndTime          *dstore_values.TimestampValue `protobuf:"bytes,10005,opt,name=message_date_and_time,json=messageDateAndTime" json:"message_date_and_time,omitempty"`
	FromCommunityMemberId       *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=from_community_member_id,json=fromCommunityMemberId" json:"from_community_member_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetFromCommunityMemberNickname() *dstore_values.StringValue {
	if m != nil {
		return m.FromCommunityMemberNickname
	}
	return nil
}

func (m *Response_Row) GetMessageStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.MessageStatus
	}
	return nil
}

func (m *Response_Row) GetMessage() *dstore_values.StringValue {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Response_Row) GetMessageNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.MessageNo
	}
	return nil
}

func (m *Response_Row) GetMessageDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.MessageDateAndTime
	}
	return nil
}

func (m *Response_Row) GetFromCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromCommunityMemberId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetMessages_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetMessages_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetMessages_Pu.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/co_GetMessages_Pu.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x59, 0x6f, 0xd3, 0x40,
	0x10, 0x56, 0x09, 0x69, 0xd3, 0x6d, 0xe9, 0xb1, 0x6d, 0xa8, 0x9b, 0x94, 0x02, 0x45, 0x08, 0x24,
	0x84, 0x5b, 0xb5, 0x42, 0x5c, 0x12, 0x08, 0x4a, 0x81, 0x20, 0x6a, 0x5a, 0x73, 0x48, 0x20, 0xa4,
	0xc5, 0x8d, 0x37, 0x91, 0xd5, 0x78, 0x37, 0x78, 0x1d, 0x0a, 0x8f, 0xfc, 0x03, 0xee, 0x1f, 0xc8,
	0x1b, 0x37, 0x6f, 0xbc, 0x32, 0xde, 0xb1, 0x93, 0xc6, 0x49, 0x71, 0x79, 0xb1, 0xb4, 0x9e, 0xf9,
	0x0e, 0xcf, 0xce, 0x8c, 0xc9, 0x92, 0xab, 0x42, 0x19, 0xf0, 0x45, 0x2e, 0xea, 0x9e, 0xe0, 0x8b,
	0xcd, 0x40, 0x56, 0xb9, 0xdb, 0x0a, 0xb8, 0x5a, 0xac, 0x4a, 0x76, 0x8b, 0x87, 0xeb, 0x5c, 0x29,
	0xa7, 0xce, 0x15, 0xdb, 0x68, 0x99, 0x10, 0x0c, 0x25, 0x3d, 0x8a, 0x08, 0x13, 0x11, 0x66, 0x4f,
	0x5a, 0x69, 0x2a, 0xa6, 0x7c, 0xe1, 0x34, 0x5a, 0x5c, 0x21, 0xaa, 0x34, 0xdb, 0xad, 0xc3, 0x83,
	0x40, 0x06, 0x71, 0xa8, 0xdc, 0x1d, 0xf2, 0x91, 0x2a, 0x0e, 0x9e, 0x48, 0x07, 0x43, 0xc7, 0x13,
	0x35, 0x19, 0xf8, 0x4e, 0xe8, 0x49, 0x81, 0x49, 0x0b, 0xaf, 0x0b, 0x84, 0x6c, 0x38, 0x81, 0x03,
	0x51, 0x1e, 0x28, 0x7a, 0x9e, 0x0c, 0xb7, 0x84, 0xf7, 0xbc, 0xc5, 0x99, 0xe7, 0x1a, 0x03, 0xc7,
	0x06, 0x4e, 0x8f, 0x2c, 0x97, 0xcc, 0xd8, 0x75, 0x6c, 0x4a, 0x85, 0x81, 0x27, 0xea, 0x8f, 0xa2,
	0x83, 0x5d, 0xc0, 0xe4, 0x8a, 0x4b, 0x4f, 0x92, 0xb1, 0x36, 0x90, 0x89, 0x56, 0xa3, 0x61, 0x7c,
	0x19, 0x02, 0x78, 0xc1, 0x1e, 0x4d, 0x52, 0x2c, 0x78, 0x49, 0x9f, 0x92, 0xb9, 0x26, 0xe8, 0x48,
	0x01, 0x69, 0x5c, 0x84, 0x5e, 0xcd, 0xab, 0x6a, 0x37, 0x0c, 0xc9, 0x8d, 0x03, 0x99, 0x92, 0x25,
	0xc4, 0x57, 0xba, 0xe0, 0x3a, 0xa4, 0xe8, 0x1d, 0x72, 0xfc, 0x5f, 0xec, 0xe8, 0xeb, 0x2b, 0xfa,
	0x9a, 0xdf, 0x9b, 0x47, 0x3b, 0xbd, 0x42, 0x46, 0xab, 0xd2, 0xf7, 0xc1, 0x7d, 0xf8, 0x2a, 0x2a,
	0x46, 0x4e, 0x3b, 0x2b, 0xa7, 0x9c, 0x79, 0x22, 0xe4, 0x75, 0x1e, 0xa0, 0xb5, 0x91, 0x36, 0x00,
	0x0a, 0x72, 0x86, 0x4c, 0xee, 0xc6, 0xa3, 0xf6, 0x37, 0xd4, 0x1e, 0xdf, 0x95, 0xa8, 0xc5, 0xd6,
	0xc8, 0x44, 0x9d, 0x87, 0x4c, 0xf0, 0x1d, 0x16, 0xdf, 0xa1, 0x32, 0x0e, 0xf6, 0x15, 0xdc, 0x92,
	0xb2, 0xc1, 0x1d, 0x34, 0x6a, 0x8f, 0x01, 0xc8, 0xe2, 0x3b, 0x49, 0x07, 0xd1, 0x65, 0x52, 0x4c,
	0xd3, 0xa0, 0xee, 0x77, 0xd4, 0xa5, 0xdd, 0xf9, 0x5a, 0xfa, 0x2e, 0x99, 0x76, 0x9d, 0x90, 0x33,
	0x47, 0xb8, 0x2c, 0xf4, 0x7c, 0xce, 0xb0, 0x43, 0x8c, 0x7c, 0xf6, 0xf7, 0x4e, 0x46, 0xc0, 0x6b,
	0xc2, 0x7d, 0x00, 0xb0, 0x9b, 0x1a, 0x45, 0x2f, 0x92, 0xd9, 0x7e, 0x6c, 0xe8, 0xe2, 0x07, 0xba,
	0x28, 0xf6, 0xc0, 0xb4, 0x91, 0x4b, 0x84, 0xc4, 0xa6, 0x99, 0x90, 0xc6, 0x60, 0xb6, 0xfc, 0x70,
	0x9c, 0x6e, 0x49, 0x7a, 0x8a, 0x8c, 0x77, 0xb0, 0x28, 0xf6, 0x13, 0xc5, 0x0e, 0xb5, 0x93, 0xb4,
	0xc8, 0x6d, 0x42, 0x7d, 0xe7, 0x25, 0x64, 0xf8, 0x5b, 0x3c, 0x60, 0xb2, 0xc6, 0x02, 0xb9, 0xa3,
	0x8c, 0xa1, 0x6c, 0xb1, 0x71, 0x80, 0x59, 0x1a, 0x75, 0xaf, 0x66, 0x03, 0x86, 0x9e, 0x23, 0x33,
	0xbd, 0x4c, 0x28, 0xfd, 0x0b, 0xa5, 0xa7, 0x52, 0x10, 0x6d, 0x60, 0x93, 0xcc, 0x28, 0xde, 0x84,
	0x81, 0x03, 0x25, 0xe6, 0xc5, 0x8d, 0x1a, 0xf5, 0xa7, 0x32, 0x0a, 0x99, 0xbd, 0x3f, 0xdd, 0x86,
	0x56, 0xb0, 0x71, 0xe1, 0xb5, 0xa2, 0x57, 0xc9, 0xdc, 0x1e, 0x94, 0x68, 0xe7, 0x37, 0xda, 0x31,
	0xfa, 0x81, 0x23, 0x4f, 0x0b, 0x7f, 0xf2, 0xa4, 0x60, 0x73, 0xd5, 0x94, 0x42, 0x71, 0xba, 0x44,
	0xf2, 0x7a, 0xc3, 0xa4, 0xa7, 0x3f, 0xde, 0x59, 0xb8, 0x7d, 0xd6, 0xa2, 0xa7, 0x8d, 0x89, 0xf4,
	0x31, 0x99, 0x88, 0x76, 0x0b, 0xdb, 0xb5, 0x5c, 0x60, 0x8e, 0x73, 0x00, 0x36, 0x53, 0xe0, 0xf4,
	0x0a, 0x5a, 0x87, 0x73, 0xa5, 0x73, 0x86, 0x22, 0x77, 0xbf, 0xa0, 0x17, 0xc8, 0x50, 0x7c, 0x7f,
	0x30, 0x7f, 0x11, 0xe3, 0x7c, 0x0f, 0x23, 0x6e, 0xbc, 0xb8, 0xa5, 0xed, 0x24, 0x1d, 0x8a, 0x92,
	0x83, 0x0b, 0x81, 0x21, 0x8a, 0x50, 0x67, 0xcd, 0x8c, 0xc5, 0x6b, 0x26, 0x9f, 0x6f, 0xc2, 0x45,
	0xd9, 0x11, 0xb2, 0xf4, 0x39, 0x47, 0x72, 0x70, 0xa0, 0x87, 0xc9, 0x20, 0x1c, 0xa3, 0x0d, 0xf0,
	0xc6, 0x82, 0x8a, 0xe4, 0xed, 0x3c, 0x1c, 0x61, 0xbe, 0x9f, 0x91, 0xf9, 0x5a, 0x20, 0x7d, 0xd6,
	0x19, 0x72, 0x9f, 0xeb, 0x56, 0x10, 0x5e, 0x75, 0x5b, 0xc0, 0x3e, 0x35, 0xde, 0x5a, 0x99, 0x17,
	0x5a, 0x8e, 0x28, 0x56, 0x13, 0x86, 0x75, 0x4d, 0x60, 0xc5, 0x78, 0xba, 0x4a, 0xc6, 0x92, 0xa6,
	0x56, 0xa1, 0x13, 0xb6, 0x94, 0xf1, 0xce, 0xca, 0x6e, 0xd4, 0xa4, 0xe1, 0xef, 0x6b, 0x08, 0xb4,
	0x69, 0xbb, 0x82, 0xef, 0xb3, 0xfd, 0xb4, 0xcb, 0x77, 0xb9, 0x6b, 0x18, 0x3f, 0x58, 0xff, 0x35,
	0x8d, 0x9b, 0xa4, 0x98, 0x80, 0xbb, 0x96, 0x81, 0xf1, 0x11, 0x79, 0x8e, 0xa4, 0x78, 0xa2, 0x18,
	0x7c, 0xa1, 0xdf, 0x44, 0x26, 0x1a, 0x83, 0x6f, 0x74, 0xd6, 0x04, 0x7d, 0x48, 0x8c, 0xfe, 0xd5,
	0x86, 0x7b, 0xf9, 0xb4, 0x0f, 0x77, 0xc5, 0x3e, 0x85, 0xae, 0xb8, 0xd7, 0x6d, 0x52, 0xf6, 0x64,
	0xaa, 0x39, 0x3a, 0xff, 0xf1, 0x27, 0x2b, 0x75, 0xa9, 0xdc, 0xed, 0x24, 0xee, 0xee, 0xeb, 0x57,
	0xbf, 0x35, 0xa8, 0x7f, 0xac, 0x2b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x8f, 0xce, 0x3c,
	0x1f, 0x08, 0x00, 0x00,
}
