// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_SearchMemberSettings_Ad.proto
// DO NOT EDIT!

/*
Package co_SearchMemberSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_SearchMemberSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_SearchMemberSettings_Ad

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
	CommunityId            *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull        bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	KeyVariable            *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull        bool                        `protobuf:"varint,1002,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	Value                  *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ValueNull              bool                        `protobuf:"varint,1003,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
	FromRowNumber          *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=from_row_number,json=fromRowNumber" json:"from_row_number,omitempty"`
	FromRowNumberNull      bool                        `protobuf:"varint,1004,opt,name=from_row_number_null,json=fromRowNumberNull" json:"from_row_number_null,omitempty"`
	MaxNumberOfRows        *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=max_number_of_rows,json=maxNumberOfRows" json:"max_number_of_rows,omitempty"`
	MaxNumberOfRowsNull    bool                        `protobuf:"varint,1005,opt,name=max_number_of_rows_null,json=maxNumberOfRowsNull" json:"max_number_of_rows_null,omitempty"`
	OrderByColumn          *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=order_by_column,json=orderByColumn" json:"order_by_column,omitempty"`
	OrderByColumnNull      bool                        `protobuf:"varint,1006,opt,name=order_by_column_null,json=orderByColumnNull" json:"order_by_column_null,omitempty"`
	OrderByLowerValues     *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=order_by_lower_values,json=orderByLowerValues" json:"order_by_lower_values,omitempty"`
	OrderByLowerValuesNull bool                        `protobuf:"varint,1007,opt,name=order_by_lower_values_null,json=orderByLowerValuesNull" json:"order_by_lower_values_null,omitempty"`
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

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Parameters) GetFromRowNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromRowNumber
	}
	return nil
}

func (m *Parameters) GetMaxNumberOfRows() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfRows
	}
	return nil
}

func (m *Parameters) GetOrderByColumn() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderByColumn
	}
	return nil
}

func (m *Parameters) GetOrderByLowerValues() *dstore_values.BooleanValue {
	if m != nil {
		return m.OrderByLowerValues
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	IdentifyingValue  *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=identifying_value,json=identifyingValue" json:"identifying_value,omitempty"`
	SettingValue      *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=setting_value,json=settingValue" json:"setting_value,omitempty"`
	CommunityMemberId *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetIdentifyingValue() *dstore_values.StringValue {
	if m != nil {
		return m.IdentifyingValue
	}
	return nil
}

func (m *Response_Row) GetSettingValue() *dstore_values.StringValue {
	if m != nil {
		return m.SettingValue
	}
	return nil
}

func (m *Response_Row) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_SearchMemberSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_SearchMemberSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_SearchMemberSettings_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_SearchMemberSettings_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x6f, 0xd3, 0x3e,
	0x14, 0xd7, 0xfe, 0x5d, 0xbb, 0xfd, 0xbd, 0x4d, 0xdb, 0x3c, 0x18, 0xa1, 0x93, 0x26, 0x34, 0x5e,
	0x26, 0x21, 0xa5, 0x08, 0x04, 0x02, 0x71, 0x11, 0x6c, 0x42, 0x5a, 0xa5, 0x35, 0xa0, 0x4c, 0x9a,
	0x04, 0x3c, 0x44, 0x69, 0xe3, 0x96, 0x68, 0x89, 0x3d, 0xd9, 0xe9, 0x4a, 0xbf, 0x05, 0x17, 0xf1,
	0x49, 0xf8, 0x54, 0xdc, 0x1f, 0x79, 0xe5, 0xd8, 0xc7, 0x6d, 0xd3, 0x74, 0x28, 0xe3, 0xa5, 0x8d,
	0xed, 0xdf, 0xe5, 0xf8, 0xe4, 0x9c, 0x13, 0x72, 0x3f, 0x52, 0x99, 0x90, 0xac, 0xc1, 0x78, 0x2f,
	0xe6, 0xac, 0x71, 0x2a, 0x45, 0x87, 0x45, 0x7d, 0xc9, 0x54, 0xa3, 0x23, 0x82, 0x23, 0x16, 0xca,
	0xce, 0x9b, 0x16, 0x4b, 0xdb, 0x4c, 0x1e, 0xb1, 0x2c, 0x8b, 0x79, 0x4f, 0x05, 0x4f, 0x23, 0x17,
	0x50, 0x99, 0xa0, 0xbb, 0x48, 0x75, 0x91, 0xea, 0xfe, 0x1d, 0x5f, 0xdf, 0xb0, 0x26, 0x67, 0x61,
	0xd2, 0x67, 0x0a, 0xe9, 0xf5, 0xab, 0xd3, 0xce, 0x4c, 0x4a, 0x21, 0xed, 0xd1, 0xd6, 0xf4, 0x51,
	0xca, 0x94, 0x0a, 0x7b, 0xcc, 0x1e, 0x5e, 0x2f, 0x1e, 0x66, 0x61, 0xcc, 0xbb, 0x42, 0xa6, 0x61,
	0x16, 0x0b, 0x8e, 0xa0, 0x9d, 0xcf, 0x35, 0x42, 0x5e, 0x84, 0x32, 0x84, 0x53, 0x26, 0x15, 0x7d,
	0x4c, 0x96, 0x3b, 0x22, 0x4d, 0xfb, 0x3c, 0xce, 0x86, 0x41, 0x1c, 0x39, 0x73, 0xd7, 0xe6, 0x76,
	0x97, 0x6e, 0x6d, 0xb9, 0xf6, 0x06, 0x36, 0xae, 0x98, 0x67, 0xac, 0xc7, 0xe4, 0xb1, 0x5e, 0xf9,
	0x4b, 0x63, 0x42, 0x33, 0xa2, 0x37, 0xc8, 0x7a, 0x9e, 0x1f, 0xf0, 0x7e, 0x92, 0x38, 0x5f, 0x16,
	0x40, 0x65, 0xd1, 0x5f, 0xcd, 0x01, 0x3d, 0xd8, 0xa7, 0x8f, 0xc8, 0xf2, 0x09, 0x1b, 0x06, 0x67,
	0xa1, 0x8c, 0xc3, 0x76, 0xc2, 0x9c, 0xff, 0x8c, 0x59, 0xbd, 0x60, 0xa6, 0x32, 0x09, 0xe9, 0xb1,
	0x5e, 0x80, 0x3f, 0xb6, 0x70, 0xed, 0x95, 0xa7, 0xa3, 0xd7, 0x57, 0xeb, 0x95, 0x03, 0x1a, 0xaf,
	0x9b, 0xa4, 0x6a, 0xf4, 0x9c, 0x4a, 0xa9, 0x09, 0x02, 0xe9, 0x36, 0x21, 0xe6, 0x01, 0x75, 0xbf,
	0xa1, 0xee, 0xff, 0x66, 0xcb, 0x28, 0xee, 0x93, 0xd5, 0xae, 0x14, 0x69, 0x20, 0xc5, 0x00, 0x20,
	0xfa, 0x4d, 0x3a, 0xf3, 0xe5, 0xd9, 0x5a, 0xd1, 0x1c, 0x5f, 0x0c, 0x3c, 0xc3, 0x80, 0xb0, 0x2e,
	0x15, 0x44, 0xd0, 0xee, 0x3b, 0xda, 0xad, 0x4f, 0xa1, 0x8d, 0xed, 0x01, 0xa1, 0x69, 0xf8, 0x76,
	0x04, 0x16, 0x5d, 0x4d, 0x55, 0x4e, 0xb5, 0xdc, 0x79, 0x15, 0x68, 0xa8, 0xf3, 0xbc, 0x0b, 0x92,
	0x8a, 0xde, 0x21, 0x57, 0x66, 0x95, 0xd0, 0xfe, 0x07, 0xda, 0x6f, 0x14, 0x28, 0xa3, 0x7b, 0x0b,
	0x19, 0x01, 0xa3, 0x3d, 0x0c, 0x3a, 0x22, 0xe9, 0xa7, 0xdc, 0xa9, 0x5d, 0xe0, 0xde, 0x86, 0xb3,
	0x37, 0xdc, 0x37, 0x0c, 0x7d, 0xef, 0x82, 0x08, 0x1a, 0xff, 0xb4, 0xf7, 0x9e, 0x42, 0x1b, 0x5b,
	0x8f, 0x5c, 0x1e, 0x33, 0x12, 0x31, 0x80, 0x07, 0xb4, 0x71, 0x16, 0xce, 0x35, 0x6f, 0x0b, 0x91,
	0xb0, 0x90, 0xa3, 0x39, 0xb5, 0x72, 0x87, 0x9a, 0x67, 0xb6, 0x14, 0x7d, 0x40, 0xea, 0xe7, 0xea,
	0x61, 0x1c, 0xbf, 0x30, 0x8e, 0xcd, 0x59, 0xa2, 0x0e, 0x66, 0xe7, 0xd3, 0x3c, 0x59, 0xf4, 0x99,
	0x3a, 0x15, 0x5c, 0x31, 0x5d, 0x5a, 0xa6, 0x27, 0x6d, 0xb3, 0x8c, 0x4b, 0xcb, 0xb6, 0x3b, 0xf6,
	0xeb, 0x33, 0xfd, 0xeb, 0x23, 0x90, 0xbe, 0x24, 0x6b, 0xba, 0x1b, 0x83, 0x5c, 0x3b, 0x42, 0xf1,
	0x57, 0x80, 0xec, 0x16, 0xc8, 0xc5, 0xa6, 0x6d, 0xc1, 0xba, 0x39, 0x59, 0xc3, 0x4b, 0x9d, 0xde,
	0xa0, 0xf7, 0xc8, 0x82, 0x9d, 0x02, 0x50, 0xe9, 0x5a, 0x71, 0x7b, 0x46, 0x11, 0x67, 0x44, 0x0b,
	0xff, 0xfd, 0x11, 0x1c, 0x0a, 0xab, 0x02, 0x05, 0x00, 0x35, 0xac, 0x59, 0x77, 0xdd, 0x8b, 0xce,
	0x2c, 0x77, 0x94, 0x07, 0x17, 0x2a, 0xc4, 0xd7, 0x12, 0xf5, 0xdf, 0x73, 0xa4, 0x02, 0x0b, 0xba,
	0x49, 0x6a, 0xba, 0xae, 0x61, 0x8c, 0xbc, 0xf3, 0x20, 0x35, 0x55, 0xbf, 0x0a, 0x4b, 0x18, 0x12,
	0x07, 0x64, 0x3d, 0x8e, 0x18, 0xcf, 0xe2, 0xee, 0x10, 0x44, 0x30, 0xef, 0xce, 0x7b, 0xaf, 0xb4,
	0x31, 0xd7, 0x72, 0x2c, 0xb3, 0x43, 0x9f, 0x90, 0x15, 0x85, 0xa1, 0x58, 0x95, 0x0f, 0xe5, 0x2a,
	0xcb, 0x96, 0x81, 0x0a, 0x87, 0x64, 0x63, 0x32, 0xb0, 0x52, 0x73, 0x39, 0x1d, 0xf0, 0x47, 0xaf,
	0xbc, 0xa4, 0x27, 0x93, 0x0e, 0x93, 0xd2, 0x8c, 0xf6, 0x5e, 0x93, 0xad, 0x58, 0x14, 0x52, 0x37,
	0xf9, 0x52, 0xbc, 0x7a, 0xd8, 0x13, 0x2a, 0x3a, 0x19, 0x9d, 0x47, 0xff, 0xf6, 0x31, 0x69, 0xd7,
	0xcc, 0xc4, 0xbe, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x36, 0x3f, 0xa9, 0x2f, 0x8a, 0x06, 0x00,
	0x00,
}
