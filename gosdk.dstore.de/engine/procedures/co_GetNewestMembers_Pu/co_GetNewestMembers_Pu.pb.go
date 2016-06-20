// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetNewestMembers_Pu.proto
// DO NOT EDIT!

/*
Package co_GetNewestMembers_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetNewestMembers_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetNewestMembers_Pu

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
	UniqueId                          *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                      bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues        *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull    bool                        `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                   bool                        `protobuf:"varint,1003,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	RowNumber                         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=row_number,json=rowNumber" json:"row_number,omitempty"`
	RowNumberNull                     bool                        `protobuf:"varint,1004,opt,name=row_number_null,json=rowNumberNull" json:"row_number_null,omitempty"`
	DateTimeFormat                    *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=date_time_format,json=dateTimeFormat" json:"date_time_format,omitempty"`
	DateTimeFormatNull                bool                        `protobuf:"varint,1005,opt,name=date_time_format_null,json=dateTimeFormatNull" json:"date_time_format_null,omitempty"`
	OutputCharacteristicId1           *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=output_characteristic_id1,json=outputCharacteristicId1" json:"output_characteristic_id1,omitempty"`
	OutputCharacteristicId1Null       bool                        `protobuf:"varint,1006,opt,name=output_characteristic_id1_null,json=outputCharacteristicId1Null" json:"output_characteristic_id1_null,omitempty"`
	OutputCharacteristicId2           *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=output_characteristic_id2,json=outputCharacteristicId2" json:"output_characteristic_id2,omitempty"`
	OutputCharacteristicId2Null       bool                        `protobuf:"varint,1007,opt,name=output_characteristic_id2_null,json=outputCharacteristicId2Null" json:"output_characteristic_id2_null,omitempty"`
	OutputCharacteristicId3           *dstore_values.IntegerValue `protobuf:"bytes,8,opt,name=output_characteristic_id3,json=outputCharacteristicId3" json:"output_characteristic_id3,omitempty"`
	OutputCharacteristicId3Null       bool                        `protobuf:"varint,1008,opt,name=output_characteristic_id3_null,json=outputCharacteristicId3Null" json:"output_characteristic_id3_null,omitempty"`
	AvgMinutesUntilNewMembership      *dstore_values.IntegerValue `protobuf:"bytes,9,opt,name=avg_minutes_until_new_membership,json=avgMinutesUntilNewMembership" json:"avg_minutes_until_new_membership,omitempty"`
	AvgMinutesUntilNewMembershipNull  bool                        `protobuf:"varint,1009,opt,name=avg_minutes_until_new_membership_null,json=avgMinutesUntilNewMembershipNull" json:"avg_minutes_until_new_membership_null,omitempty"`
	CommunityBinaryCategoryId         *dstore_values.IntegerValue `protobuf:"bytes,10,opt,name=community_binary_category_id,json=communityBinaryCategoryId" json:"community_binary_category_id,omitempty"`
	CommunityBinaryCategoryIdNull     bool                        `protobuf:"varint,1010,opt,name=community_binary_category_id_null,json=communityBinaryCategoryIdNull" json:"community_binary_category_id_null,omitempty"`
	OnlineStatusInsteadOfIsOnline     *dstore_values.BooleanValue `protobuf:"bytes,11,opt,name=online_status_instead_of_is_online,json=onlineStatusInsteadOfIsOnline" json:"online_status_instead_of_is_online,omitempty"`
	OnlineStatusInsteadOfIsOnlineNull bool                        `protobuf:"varint,1011,opt,name=online_status_instead_of_is_online_null,json=onlineStatusInsteadOfIsOnlineNull" json:"online_status_instead_of_is_online_null,omitempty"`
	SeparatorInIdentVals              *dstore_values.StringValue  `protobuf:"bytes,12,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull          bool                        `protobuf:"varint,1012,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetRowNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.RowNumber
	}
	return nil
}

func (m *Parameters) GetDateTimeFormat() *dstore_values.IntegerValue {
	if m != nil {
		return m.DateTimeFormat
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId1() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId1
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId2() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId2
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId3() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId3
	}
	return nil
}

func (m *Parameters) GetAvgMinutesUntilNewMembership() *dstore_values.IntegerValue {
	if m != nil {
		return m.AvgMinutesUntilNewMembership
	}
	return nil
}

func (m *Parameters) GetCommunityBinaryCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityBinaryCategoryId
	}
	return nil
}

func (m *Parameters) GetOnlineStatusInsteadOfIsOnline() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlineStatusInsteadOfIsOnline
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
	RowId                     int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value1RestrictedByPattern *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=value1_restricted_by_pattern,json=value1RestrictedByPattern" json:"value1_restricted_by_pattern,omitempty"`
	BinaryId                  *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=binary_id,json=binaryId" json:"binary_id,omitempty"`
	Value2RestrictedByPattern *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=value2_restricted_by_pattern,json=value2RestrictedByPattern" json:"value2_restricted_by_pattern,omitempty"`
	CommunityMemberId         *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	Value3                    *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=value3" json:"value3,omitempty"`
	OnlineStatus              *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=online_status,json=onlineStatus" json:"online_status,omitempty"`
	MemberSince               *dstore_values.TimestampValue `protobuf:"bytes,10007,opt,name=member_since,json=memberSince" json:"member_since,omitempty"`
	Value1                    *dstore_values.StringValue    `protobuf:"bytes,10008,opt,name=value1" json:"value1,omitempty"`
	Value2                    *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=value2" json:"value2,omitempty"`
	Value3RestrictedByPattern *dstore_values.StringValue    `protobuf:"bytes,10010,opt,name=value3_restricted_by_pattern,json=value3RestrictedByPattern" json:"value3_restricted_by_pattern,omitempty"`
	Nickname                  *dstore_values.StringValue    `protobuf:"bytes,10011,opt,name=nickname" json:"nickname,omitempty"`
	IsOnline                  *dstore_values.BooleanValue   `protobuf:"bytes,10012,opt,name=is_online,json=isOnline" json:"is_online,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue1RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value1RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetBinaryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryId
	}
	return nil
}

func (m *Response_Row) GetValue2RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value2RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Response_Row) GetValue3() *dstore_values.StringValue {
	if m != nil {
		return m.Value3
	}
	return nil
}

func (m *Response_Row) GetOnlineStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.OnlineStatus
	}
	return nil
}

func (m *Response_Row) GetMemberSince() *dstore_values.TimestampValue {
	if m != nil {
		return m.MemberSince
	}
	return nil
}

func (m *Response_Row) GetValue1() *dstore_values.StringValue {
	if m != nil {
		return m.Value1
	}
	return nil
}

func (m *Response_Row) GetValue2() *dstore_values.StringValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Response_Row) GetValue3RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value3RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetNickname() *dstore_values.StringValue {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *Response_Row) GetIsOnline() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsOnline
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetNewestMembers_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetNewestMembers_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetNewestMembers_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_GetNewestMembers_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1038 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x97, 0xeb, 0x6e, 0x1b, 0x45,
	0x14, 0xc7, 0x55, 0x42, 0x5c, 0x67, 0xe2, 0x34, 0x65, 0x0b, 0x74, 0xe2, 0x5c, 0x94, 0x04, 0xaa,
	0x22, 0x21, 0x39, 0x78, 0x2d, 0xd4, 0xc2, 0x07, 0x10, 0x29, 0x05, 0x19, 0x61, 0x37, 0x6c, 0x80,
	0x0a, 0x84, 0x34, 0x5a, 0xef, 0x8e, 0xdd, 0x51, 0xbd, 0x33, 0x66, 0x67, 0xb6, 0x51, 0xde, 0x82,
	0xfb, 0xfd, 0xc5, 0x78, 0x04, 0xee, 0x94, 0xf2, 0x00, 0x9c, 0x9d, 0xb3, 0x5e, 0x5f, 0xb0, 0xbd,
	0x56, 0xbf, 0x24, 0xda, 0x9d, 0xf3, 0xff, 0x9f, 0x9f, 0x76, 0xe6, 0x9c, 0x33, 0x26, 0x2f, 0x87,
	0xda, 0xa8, 0x98, 0x1f, 0x71, 0xd9, 0x13, 0x92, 0x1f, 0x0d, 0x62, 0x15, 0xf0, 0x30, 0x89, 0xb9,
	0x3e, 0x0a, 0x14, 0x7b, 0x9b, 0x9b, 0x36, 0x3f, 0xe3, 0xda, 0xb4, 0x78, 0xd4, 0xe1, 0xb1, 0x66,
	0x27, 0x49, 0x0d, 0x22, 0x8c, 0x72, 0x9e, 0x47, 0x59, 0x0d, 0x65, 0xb5, 0xd9, 0xb1, 0xd5, 0x2b,
	0x99, 0xf9, 0x03, 0xbf, 0x9f, 0x70, 0x8d, 0xd2, 0xea, 0xd6, 0x64, 0x46, 0x1e, 0xc7, 0x2a, 0xce,
	0x96, 0xb6, 0x27, 0x97, 0x22, 0xae, 0xb5, 0xdf, 0xe3, 0xd9, 0xe2, 0x73, 0xd3, 0x8b, 0xc6, 0x17,
	0xb2, 0xab, 0xe2, 0xc8, 0x37, 0x42, 0x49, 0x0c, 0x3a, 0x7c, 0xb8, 0x41, 0xc8, 0x89, 0x1f, 0xfb,
	0xb0, 0x0a, 0x0c, 0xce, 0x0d, 0xb2, 0x96, 0x48, 0xf1, 0x69, 0xc2, 0x99, 0x08, 0xe9, 0x85, 0xfd,
	0x0b, 0x2f, 0xac, 0xbb, 0xd5, 0x5a, 0x86, 0x9e, 0x41, 0x69, 0x13, 0x0b, 0xd9, 0xfb, 0x30, 0x7d,
	0xf0, 0xca, 0x18, 0xdc, 0x0c, 0x9d, 0x6b, 0xe4, 0x52, 0x2e, 0x64, 0x32, 0xe9, 0xf7, 0xe9, 0x2f,
	0x17, 0x41, 0x5e, 0xf6, 0x2a, 0xc3, 0x90, 0x36, 0xbc, 0x74, 0x3e, 0x21, 0x3b, 0x03, 0xc8, 0xa3,
	0x24, 0x84, 0x71, 0x69, 0x44, 0x57, 0x04, 0x96, 0x86, 0xa1, 0x39, 0x7d, 0xa2, 0x30, 0x65, 0x15,
	0xf5, 0xcd, 0x09, 0xb9, 0x5d, 0xd2, 0xce, 0x3b, 0xe4, 0x60, 0x91, 0x3b, 0x72, 0xfd, 0x8a, 0x5c,
	0x7b, 0xf3, 0x7d, 0x2c, 0xe9, 0x6b, 0xa4, 0x12, 0xa8, 0x28, 0x02, 0x7a, 0x73, 0x9e, 0x7e, 0x8c,
	0x15, 0x4b, 0xb6, 0x3d, 0x45, 0x26, 0xa4, 0xe1, 0x3d, 0x1e, 0x23, 0xda, 0x7a, 0x2e, 0x80, 0x0f,
	0xf2, 0x22, 0x79, 0x6a, 0x5c, 0x8f, 0xb9, 0x7f, 0xc3, 0xdc, 0x9b, 0x63, 0x81, 0x36, 0xd9, 0xab,
	0x84, 0xc4, 0xea, 0x0c, 0x62, 0xd2, 0x93, 0x40, 0x9f, 0x2c, 0x4e, 0xb5, 0x06, 0xe1, 0x6d, 0x1b,
	0xed, 0x5c, 0x27, 0x9b, 0x23, 0x2d, 0xa6, 0xf9, 0x1d, 0xd3, 0x6c, 0xe4, 0x41, 0x36, 0xc9, 0x6d,
	0x72, 0x39, 0xf4, 0x0d, 0x67, 0x46, 0x44, 0x9c, 0xe1, 0x39, 0xa0, 0xab, 0xc5, 0xa9, 0x2e, 0xa5,
	0xa2, 0xf7, 0x41, 0xf3, 0x96, 0x95, 0x38, 0x2e, 0x79, 0x66, 0xda, 0x06, 0xb3, 0xfe, 0x81, 0x59,
	0x9d, 0xc9, 0x78, 0x9b, 0xfa, 0x2e, 0xd9, 0x52, 0x89, 0x19, 0x24, 0x86, 0x05, 0xf7, 0xe0, 0xb0,
	0x05, 0x70, 0xd6, 0x84, 0x36, 0x22, 0x80, 0x0f, 0x53, 0xa7, 0xa5, 0x62, 0x86, 0xab, 0xa8, 0xbe,
	0x35, 0x21, 0x6e, 0x86, 0x75, 0xe7, 0x4d, 0xb2, 0x37, 0xd7, 0x18, 0xa9, 0xfe, 0x44, 0xaa, 0xed,
	0x39, 0x0e, 0x45, 0x78, 0x2e, 0xbd, 0xf8, 0xd8, 0x78, 0xee, 0x22, 0x3c, 0x17, 0xf1, 0xfe, 0x5a,
	0x88, 0xe7, 0x16, 0xe1, 0x35, 0x68, 0xf9, 0xb1, 0xf1, 0x1a, 0x8b, 0xf0, 0x1a, 0x88, 0xf7, 0xf7,
	0x42, 0xbc, 0x86, 0xc5, 0x0b, 0xc8, 0xbe, 0xff, 0xa0, 0xc7, 0x22, 0x21, 0x13, 0x03, 0x45, 0x96,
	0x40, 0x3d, 0xf5, 0x99, 0xe4, 0x67, 0x2c, 0xc2, 0xc6, 0x76, 0x4f, 0x0c, 0xe8, 0x5a, 0x31, 0xe5,
	0x0e, 0x98, 0xb4, 0xd0, 0xe3, 0x83, 0xd4, 0x02, 0x9a, 0x63, 0x2b, 0x37, 0x70, 0x4e, 0xc8, 0xb5,
	0xa2, 0x24, 0x48, 0xfc, 0x10, 0x89, 0xf7, 0x17, 0xb9, 0x0d, 0x5b, 0xd1, 0xa8, 0x40, 0x3b, 0x42,
	0xfa, 0xf1, 0x39, 0x83, 0x2e, 0xc0, 0x7b, 0x2a, 0xb6, 0x05, 0x4f, 0x8a, 0x91, 0xb7, 0x72, 0x83,
	0x63, 0xab, 0xbf, 0x95, 0xc9, 0xa1, 0xfc, 0x9b, 0xe4, 0x60, 0x91, 0x3b, 0xb2, 0xfe, 0x83, 0xac,
	0xbb, 0x73, 0x6d, 0x2c, 0x68, 0x97, 0x1c, 0x2a, 0xd9, 0x87, 0x16, 0xce, 0xb4, 0xf1, 0x4d, 0xa2,
	0x99, 0x90, 0xda, 0x70, 0x3f, 0x64, 0xaa, 0xcb, 0x84, 0x66, 0xb8, 0x46, 0xd7, 0x67, 0xe2, 0x76,
	0x94, 0xea, 0x73, 0x1f, 0xfb, 0x9a, 0xb7, 0x8b, 0xa1, 0xa7, 0xd6, 0xa5, 0x89, 0x26, 0x77, 0xba,
	0x4d, 0x7d, 0xc7, 0xbe, 0x76, 0x4e, 0xc9, 0xf5, 0xe2, 0x3c, 0x08, 0xfe, 0x08, 0xc1, 0x0f, 0x16,
	0x1a, 0x5a, 0xf8, 0xf7, 0xc8, 0x55, 0xcd, 0x07, 0x70, 0x6c, 0x80, 0x0a, 0x0c, 0xb1, 0x31, 0xa7,
	0xfd, 0x58, 0xd3, 0x4a, 0x61, 0xaf, 0x7f, 0x3a, 0x97, 0x36, 0xb1, 0x51, 0xc3, 0x6b, 0xed, 0xbc,
	0x4e, 0x76, 0xe6, 0x58, 0x22, 0xdc, 0xbf, 0x08, 0x47, 0x67, 0x89, 0x53, 0xa6, 0xc3, 0x47, 0x65,
	0x52, 0xf6, 0xb8, 0x1e, 0x28, 0xa9, 0xb9, 0xf3, 0x12, 0x59, 0xb5, 0x13, 0x75, 0x7a, 0xda, 0x65,
	0x83, 0x1a, 0xa7, 0xed, 0xed, 0xf4, 0xaf, 0x87, 0x81, 0xce, 0x47, 0xe4, 0x72, 0x3a, 0x4b, 0xd9,
	0xd8, 0x30, 0x85, 0xb9, 0xb5, 0x02, 0xe2, 0xda, 0x94, 0x78, 0x7a, 0xe4, 0xb6, 0xe0, 0xb9, 0x39,
	0x7a, 0xf6, 0x36, 0xa3, 0xc9, 0x17, 0xce, 0x4d, 0x72, 0x31, 0x9b, 0xe1, 0x30, 0x6f, 0x52, 0xc7,
	0xbd, 0xff, 0x39, 0xe2, 0x84, 0x6f, 0xe1, 0x7f, 0x6f, 0x18, 0x0e, 0xa5, 0xbc, 0x02, 0xdd, 0x1e,
	0x46, 0x47, 0xaa, 0x72, 0x6b, 0xcb, 0xdc, 0x36, 0x6a, 0xc3, 0x6f, 0x50, 0xf3, 0xd4, 0x99, 0x97,
	0xca, 0xab, 0x3f, 0x97, 0xc8, 0x0a, 0x3c, 0x38, 0xcf, 0x92, 0x52, 0x3a, 0x53, 0xa0, 0x0a, 0x3e,
	0x6b, 0xc3, 0x67, 0x59, 0xf5, 0x56, 0xe1, 0x11, 0x4e, 0x35, 0xd4, 0x8c, 0xdd, 0xa6, 0x3a, 0x83,
	0x0b, 0x0f, 0x6c, 0x14, 0xf4, 0x82, 0x90, 0x75, 0xce, 0xd9, 0xc0, 0x37, 0xd0, 0x15, 0x24, 0xfd,
	0xbc, 0x5d, 0xb8, 0xa7, 0x5b, 0x68, 0xe0, 0xe5, 0xfa, 0xe3, 0xf3, 0x13, 0x54, 0x3b, 0xaf, 0x90,
	0xb5, 0xac, 0x52, 0x20, 0xf1, 0x17, 0xed, 0xe2, 0xfa, 0x2b, 0x63, 0xf8, 0x18, 0x98, 0x3b, 0x07,
	0xec, 0xcb, 0x65, 0xc1, 0xdc, 0x59, 0x60, 0xef, 0x92, 0x2b, 0xa3, 0x62, 0xc6, 0x86, 0x93, 0x22,
	0x7e, 0xb5, 0x04, 0xe2, 0xe8, 0x12, 0x80, 0x1f, 0x1e, 0x58, 0x1b, 0xa4, 0x64, 0x23, 0x1b, 0xf4,
	0xeb, 0x62, 0xaa, 0x2c, 0xd4, 0x79, 0x83, 0x6c, 0x4c, 0x14, 0x27, 0xfd, 0x66, 0x89, 0xe4, 0x95,
	0xf1, 0xfa, 0x04, 0x8b, 0x4a, 0xc6, 0xae, 0x85, 0x0c, 0x38, 0xfd, 0x16, 0x1d, 0x76, 0xa7, 0x1c,
	0xd2, 0xb9, 0x0e, 0x39, 0xa2, 0x41, 0x76, 0xa9, 0x41, 0xcd, 0x69, 0x2a, 0xc9, 0xd1, 0xeb, 0xf4,
	0xbb, 0x65, 0xd1, 0xeb, 0xb9, 0xc8, 0xa5, 0xdf, 0x2f, 0x2b, 0x72, 0xf3, 0x0d, 0x6d, 0xcc, 0xd9,
	0xd0, 0x1f, 0x96, 0xdd, 0xd0, 0xc6, 0xac, 0x0d, 0xbd, 0x41, 0xca, 0x52, 0x04, 0xf7, 0x25, 0x5c,
	0x7b, 0xe9, 0x8f, 0xc5, 0x4e, 0x79, 0x70, 0x7a, 0x44, 0x47, 0x2d, 0xf7, 0xa7, 0x76, 0x71, 0xcf,
	0x2d, 0x8b, 0xac, 0x1b, 0x1e, 0xdf, 0x25, 0xdb, 0x42, 0x4d, 0x15, 0xe6, 0xe8, 0xd7, 0xc3, 0xc7,
	0x37, 0x7b, 0x4a, 0x87, 0xf7, 0x87, 0xeb, 0xe1, 0xf2, 0x3f, 0x30, 0x3a, 0x25, 0x7b, 0x93, 0x6f,
	0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x95, 0xdb, 0x5a, 0x4c, 0x9a, 0x0c, 0x00, 0x00,
}