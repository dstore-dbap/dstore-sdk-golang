// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_SearchMembers_Ad.proto
// DO NOT EDIT!

/*
Package co_SearchMembers_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_SearchMembers_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_SearchMembers_Ad

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
	CommunityId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull               bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	SearchString                  *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=search_string,json=searchString" json:"search_string,omitempty"`
	SearchStringNull              bool                        `protobuf:"varint,1002,opt,name=search_string_null,json=searchStringNull" json:"search_string_null,omitempty"`
	MaxNumberOfRows               *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=max_number_of_rows,json=maxNumberOfRows" json:"max_number_of_rows,omitempty"`
	MaxNumberOfRowsNull           bool                        `protobuf:"varint,1003,opt,name=max_number_of_rows_null,json=maxNumberOfRowsNull" json:"max_number_of_rows_null,omitempty"`
	CaseSensitive                 *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull             bool                        `protobuf:"varint,1004,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
	OutputCharacteristicId1       *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=output_characteristic_id1,json=outputCharacteristicId1" json:"output_characteristic_id1,omitempty"`
	OutputCharacteristicId1Null   bool                        `protobuf:"varint,1005,opt,name=output_characteristic_id1_null,json=outputCharacteristicId1Null" json:"output_characteristic_id1_null,omitempty"`
	OutputCharacteristicId2       *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=output_characteristic_id2,json=outputCharacteristicId2" json:"output_characteristic_id2,omitempty"`
	OutputCharacteristicId2Null   bool                        `protobuf:"varint,1006,opt,name=output_characteristic_id2_null,json=outputCharacteristicId2Null" json:"output_characteristic_id2_null,omitempty"`
	OutputCharacteristicId3       *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=output_characteristic_id3,json=outputCharacteristicId3" json:"output_characteristic_id3,omitempty"`
	OutputCharacteristicId3Null   bool                        `protobuf:"varint,1007,opt,name=output_characteristic_id3_null,json=outputCharacteristicId3Null" json:"output_characteristic_id3_null,omitempty"`
	CommunityBinaryCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,8,opt,name=community_binary_category_id,json=communityBinaryCategoryId" json:"community_binary_category_id,omitempty"`
	CommunityBinaryCategoryIdNull bool                        `protobuf:"varint,1008,opt,name=community_binary_category_id_null,json=communityBinaryCategoryIdNull" json:"community_binary_category_id_null,omitempty"`
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

func (m *Parameters) GetSearchString() *dstore_values.StringValue {
	if m != nil {
		return m.SearchString
	}
	return nil
}

func (m *Parameters) GetMaxNumberOfRows() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfRows
	}
	return nil
}

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
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

func (m *Parameters) GetCommunityBinaryCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityBinaryCategoryId
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value1RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value1_restricted_by_pattern,json=value1RestrictedByPattern" json:"value1_restricted_by_pattern,omitempty"`
	BinaryId                  *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=binary_id,json=binaryId" json:"binary_id,omitempty"`
	Value2RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value2_restricted_by_pattern,json=value2RestrictedByPattern" json:"value2_restricted_by_pattern,omitempty"`
	CommunityMemberId         *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	Value3                    *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=value3" json:"value3,omitempty"`
	Value1                    *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=value1" json:"value1,omitempty"`
	Value2                    *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=value2" json:"value2,omitempty"`
	Value3RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=value3_restricted_by_pattern,json=value3RestrictedByPattern" json:"value3_restricted_by_pattern,omitempty"`
	Nickname                  *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=nickname" json:"nickname,omitempty"`
	IsOnline                  *dstore_values.BooleanValue `protobuf:"bytes,10010,opt,name=is_online,json=isOnline" json:"is_online,omitempty"`
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_SearchMembers_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_SearchMembers_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_SearchMembers_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/co_SearchMembers_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0xd3, 0x4a,
	0x14, 0x55, 0x4f, 0x4e, 0x2e, 0x67, 0xda, 0x9e, 0xb6, 0x53, 0xe9, 0xd4, 0x49, 0x0e, 0x15, 0x2d,
	0x2f, 0x48, 0x08, 0x87, 0xd8, 0xaa, 0x80, 0x17, 0x50, 0x53, 0x90, 0x88, 0x44, 0xd3, 0xca, 0x45,
	0x20, 0x10, 0x92, 0xe5, 0xd8, 0xd3, 0x74, 0xd4, 0x64, 0x26, 0xf2, 0x38, 0x85, 0xfc, 0x05, 0xf7,
	0xdb, 0x6f, 0xf1, 0x0d, 0x3c, 0x50, 0xae, 0x9f, 0xc0, 0xf6, 0x8c, 0xe3, 0x5c, 0x9a, 0xc4, 0x51,
	0x5f, 0x5a, 0x8d, 0xf7, 0x5e, 0x7b, 0xad, 0xec, 0x3d, 0x7b, 0xd9, 0xc8, 0xf0, 0x44, 0xc0, 0x7d,
	0x52, 0x22, 0xac, 0x41, 0x19, 0x29, 0xb5, 0x7d, 0xee, 0x12, 0xaf, 0xe3, 0x13, 0x51, 0x72, 0xb9,
	0x7d, 0x40, 0x1c, 0xdf, 0x3d, 0xda, 0x25, 0xad, 0x3a, 0xf1, 0x85, 0xbd, 0xed, 0xe9, 0x10, 0x0e,
	0x38, 0xde, 0x50, 0x18, 0x5d, 0x61, 0xf4, 0x31, 0x89, 0x85, 0xd5, 0xa8, 0xec, 0x89, 0xd3, 0xec,
	0x10, 0xa1, 0x70, 0x85, 0xfc, 0x30, 0x17, 0xf1, 0x7d, 0xee, 0x47, 0xa1, 0xe2, 0x70, 0xa8, 0x45,
	0x84, 0x70, 0x1a, 0x24, 0x0a, 0x5e, 0x1a, 0x0d, 0x06, 0x0e, 0x65, 0x87, 0xdc, 0x6f, 0x39, 0x01,
	0xe5, 0x4c, 0x25, 0x6d, 0x7e, 0xc9, 0x21, 0xb4, 0xef, 0xf8, 0x0e, 0x44, 0x41, 0x03, 0xbe, 0x85,
	0x16, 0x5c, 0xde, 0x6a, 0x75, 0x18, 0x0d, 0xba, 0x36, 0xf5, 0xb4, 0xb9, 0x8b, 0x73, 0x97, 0xe7,
	0x8d, 0xa2, 0x1e, 0x49, 0x8f, 0x74, 0x51, 0x16, 0x90, 0x06, 0xf1, 0x1f, 0x86, 0x27, 0x6b, 0x3e,
	0x06, 0x54, 0x3d, 0x7c, 0x05, 0xad, 0x0c, 0xe2, 0x6d, 0xd6, 0x69, 0x36, 0xb5, 0xaf, 0x59, 0xa8,
	0x92, 0xb3, 0x96, 0x06, 0x12, 0x6b, 0xf0, 0x1c, 0xdf, 0x46, 0x8b, 0x42, 0x76, 0xc0, 0x16, 0x81,
	0x4f, 0x59, 0x43, 0xfb, 0x4b, 0xb2, 0x15, 0x46, 0xd8, 0x54, 0x50, 0x91, 0x2d, 0x28, 0xc0, 0x81,
	0x7c, 0x84, 0xaf, 0x22, 0x3c, 0x54, 0x40, 0xd1, 0x9d, 0x2a, 0xba, 0xe5, 0xc1, 0x54, 0xc9, 0x77,
	0x0f, 0xe1, 0x96, 0xf3, 0x1c, 0x92, 0xc2, 0x7e, 0xdb, 0xfc, 0xd0, 0xf6, 0xf9, 0x33, 0xa1, 0xa5,
	0x92, 0x7f, 0xe2, 0x12, 0xc0, 0x6a, 0x12, 0xb5, 0x77, 0x68, 0x01, 0x06, 0x6f, 0xa1, 0xb5, 0xb3,
	0x95, 0x14, 0xfb, 0x37, 0xc5, 0xbe, 0x3a, 0x02, 0x91, 0x02, 0x2a, 0xe8, 0x5f, 0xd7, 0x11, 0xc4,
	0x16, 0x84, 0x09, 0x1a, 0xd0, 0x13, 0xa2, 0xfd, 0x3d, 0x96, 0xbc, 0xce, 0x79, 0x93, 0x38, 0x4c,
	0x91, 0x2f, 0x86, 0x90, 0x83, 0x1e, 0x02, 0x97, 0xd0, 0xea, 0x70, 0x0d, 0x45, 0xfb, 0x5d, 0xd1,
	0xae, 0x0c, 0x25, 0x4b, 0xd2, 0x47, 0x28, 0xcf, 0x3b, 0x41, 0xbb, 0x13, 0xd8, 0xee, 0x11, 0x0c,
	0xda, 0x85, 0x39, 0x53, 0x11, 0x50, 0x17, 0xc6, 0x53, 0xd6, 0xd2, 0xc9, 0x3f, 0x7e, 0x4d, 0xa1,
	0x77, 0x86, 0xc0, 0x55, 0xaf, 0x8c, 0xef, 0xa0, 0xf5, 0x89, 0x85, 0x95, 0xa8, 0x1f, 0x4a, 0x54,
	0x71, 0x42, 0x85, 0x24, 0x79, 0x86, 0x96, 0x39, 0xb7, 0x3c, 0x63, 0x9a, 0x3c, 0x43, 0xc9, 0xfb,
	0x39, 0x55, 0x9e, 0x91, 0x24, 0xcf, 0xd4, 0xb2, 0xe7, 0x96, 0x67, 0x4e, 0x93, 0x67, 0x2a, 0x79,
	0xbf, 0xa6, 0xca, 0x33, 0xa5, 0xbc, 0xa7, 0xe8, 0xff, 0xfe, 0xbe, 0xd5, 0x29, 0x73, 0xfc, 0xae,
	0xed, 0x3a, 0xc0, 0xcf, 0x7d, 0xb9, 0xbf, 0xb9, 0x64, 0x85, 0xf9, 0xb8, 0x40, 0x45, 0xe2, 0x77,
	0x22, 0x38, 0x6c, 0x73, 0x15, 0x6d, 0x4c, 0xab, 0xae, 0x64, 0xfe, 0x56, 0x32, 0x2f, 0x4c, 0x2c,
	0x13, 0x0a, 0xdd, 0x3c, 0xcd, 0xa2, 0x9c, 0x45, 0x44, 0x9b, 0x33, 0x41, 0xf0, 0x35, 0x94, 0x96,
	0x2e, 0x16, 0xd9, 0x4b, 0xbc, 0xf0, 0x91, 0x33, 0x2a, 0x87, 0xbb, 0x1b, 0xfe, 0xb5, 0x54, 0x22,
	0x7e, 0x8c, 0x96, 0x43, 0xff, 0xb2, 0x07, 0x0c, 0x0c, 0xdc, 0x22, 0x05, 0x60, 0x7d, 0x04, 0x3c,
	0x6a, 0x73, 0xbb, 0x70, 0xae, 0xf6, 0xcf, 0xb0, 0xcb, 0xc3, 0x0f, 0xf0, 0x0d, 0x94, 0x8d, 0x7c,
	0x13, 0xac, 0x20, 0xac, 0xb8, 0x7e, 0xa6, 0xa2, 0x72, 0xd5, 0x5d, 0xf5, 0xdf, 0xea, 0xa5, 0xe3,
	0x6d, 0x94, 0x82, 0xbd, 0x87, 0x1d, 0x0e, 0x51, 0x25, 0x3d, 0xd1, 0xde, 0xf5, 0x5e, 0x03, 0x74,
	0x70, 0x04, 0x2b, 0xc4, 0x16, 0x3e, 0xa7, 0x51, 0x0a, 0x0e, 0xf8, 0x3f, 0x94, 0x81, 0x63, 0x38,
	0xb1, 0x17, 0x35, 0xe8, 0x49, 0xda, 0x4a, 0xc3, 0x11, 0x26, 0x00, 0xf3, 0x95, 0x33, 0x2b, 0xdb,
	0xf0, 0x6e, 0x01, 0x27, 0x83, 0x0b, 0xe0, 0xd9, 0xf5, 0xae, 0xdd, 0x76, 0x02, 0xb8, 0x0a, 0x4c,
	0x7b, 0x59, 0x4b, 0xb4, 0xcc, 0xbc, 0x2a, 0x60, 0xc5, 0xf8, 0x4a, 0x77, 0x5f, 0xa1, 0xf1, 0x4d,
	0xf4, 0x4f, 0x34, 0x55, 0x20, 0x7e, 0x55, 0x4b, 0xbe, 0x2b, 0x39, 0x95, 0x3e, 0x20, 0xcc, 0x98,
	0x20, 0xec, 0xf5, 0xac, 0xc2, 0x8c, 0x71, 0xc2, 0xee, 0x83, 0xc9, 0xc5, 0x17, 0xaf, 0x25, 0x1b,
	0x18, 0x4a, 0x7c, 0x33, 0x83, 0xc4, 0xfe, 0xfb, 0x47, 0x35, 0x1e, 0xb4, 0x9a, 0x28, 0x23, 0x33,
	0x4d, 0xed, 0x6d, 0xb2, 0xaa, 0x28, 0x35, 0x06, 0x95, 0xb5, 0x77, 0xb3, 0x82, 0xca, 0x31, 0xc8,
	0xd0, 0xde, 0xcf, 0x0a, 0x32, 0xe2, 0x56, 0x9a, 0x13, 0x5a, 0xf9, 0x61, 0xd6, 0x56, 0x9a, 0xe3,
	0x5a, 0x79, 0x1d, 0xe5, 0x18, 0x75, 0x8f, 0x19, 0xbc, 0xe1, 0xb5, 0x8f, 0xc9, 0x95, 0xe2, 0xe4,
	0xf0, 0x72, 0x50, 0x61, 0x73, 0xd6, 0x84, 0xdb, 0xac, 0x7d, 0xaa, 0x25, 0xbf, 0xa8, 0x72, 0x54,
	0xec, 0xc9, 0xec, 0xca, 0x03, 0x54, 0xa4, 0x7c, 0x64, 0x1f, 0xfa, 0x9f, 0x48, 0x4f, 0xb6, 0x1a,
	0x5c, 0x78, 0xc7, 0xbd, 0xb8, 0x37, 0xe3, 0x57, 0x54, 0x3d, 0x23, 0xbf, 0x58, 0xcc, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0x46, 0xe8, 0x05, 0x7c, 0x09, 0x00, 0x00,
}
