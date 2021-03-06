// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_CreateNewCommunityMember_Pu.proto
// DO NOT EDIT!

/*
Package co_CreateNewCommunityMember_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_CreateNewCommunityMember_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_CreateNewCommunityMember_Pu

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
	CommunityId                      *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                  bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	UniqueId                         *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                     bool                        `protobuf:"varint,1002,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	CharacteristicIdList1            *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=characteristic_id_list1,json=characteristicIdList1" json:"characteristic_id_list1,omitempty"`
	CharacteristicIdList1Null        bool                        `protobuf:"varint,1003,opt,name=characteristic_id_list1_null,json=characteristicIdList1Null" json:"characteristic_id_list1_null,omitempty"`
	ValueList1                       *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=value_list1,json=valueList1" json:"value_list1,omitempty"`
	ValueList1Null                   bool                        `protobuf:"varint,1004,opt,name=value_list1_null,json=valueList1Null" json:"value_list1_null,omitempty"`
	CharacteristicIdList2            *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=characteristic_id_list2,json=characteristicIdList2" json:"characteristic_id_list2,omitempty"`
	CharacteristicIdList2Null        bool                        `protobuf:"varint,1005,opt,name=characteristic_id_list2_null,json=characteristicIdList2Null" json:"characteristic_id_list2_null,omitempty"`
	ValueList2                       *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=value_list2,json=valueList2" json:"value_list2,omitempty"`
	ValueList2Null                   bool                        `protobuf:"varint,1006,opt,name=value_list2_null,json=valueList2Null" json:"value_list2_null,omitempty"`
	PersonCharacCategoryId           *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=person_charac_category_id,json=personCharacCategoryId" json:"person_charac_category_id,omitempty"`
	PersonCharacCategoryIdNull       bool                        `protobuf:"varint,1007,opt,name=person_charac_category_id_null,json=personCharacCategoryIdNull" json:"person_charac_category_id_null,omitempty"`
	ResultInErrorIdList              *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=result_in_error_id_list,json=resultInErrorIdList" json:"result_in_error_id_list,omitempty"`
	ResultInErrorIdListNull          bool                        `protobuf:"varint,1008,opt,name=result_in_error_id_list_null,json=resultInErrorIdListNull" json:"result_in_error_id_list_null,omitempty"`
	ValueIdsForPredefinedCharacs     *dstore_values.BooleanValue `protobuf:"bytes,9,opt,name=value_ids_for_predefined_characs,json=valueIdsForPredefinedCharacs" json:"value_ids_for_predefined_characs,omitempty"`
	ValueIdsForPredefinedCharacsNull bool                        `protobuf:"varint,1009,opt,name=value_ids_for_predefined_characs_null,json=valueIdsForPredefinedCharacsNull" json:"value_ids_for_predefined_characs_null,omitempty"`
	CancelOnError                    *dstore_values.BooleanValue `protobuf:"bytes,10,opt,name=cancel_on_error,json=cancelOnError" json:"cancel_on_error,omitempty"`
	CancelOnErrorNull                bool                        `protobuf:"varint,1010,opt,name=cancel_on_error_null,json=cancelOnErrorNull" json:"cancel_on_error_null,omitempty"`
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

func (m *Parameters) GetCharacteristicIdList1() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicIdList1
	}
	return nil
}

func (m *Parameters) GetValueList1() *dstore_values.StringValue {
	if m != nil {
		return m.ValueList1
	}
	return nil
}

func (m *Parameters) GetCharacteristicIdList2() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicIdList2
	}
	return nil
}

func (m *Parameters) GetValueList2() *dstore_values.StringValue {
	if m != nil {
		return m.ValueList2
	}
	return nil
}

func (m *Parameters) GetPersonCharacCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacCategoryId
	}
	return nil
}

func (m *Parameters) GetResultInErrorIdList() *dstore_values.BooleanValue {
	if m != nil {
		return m.ResultInErrorIdList
	}
	return nil
}

func (m *Parameters) GetValueIdsForPredefinedCharacs() *dstore_values.BooleanValue {
	if m != nil {
		return m.ValueIdsForPredefinedCharacs
	}
	return nil
}

func (m *Parameters) GetCancelOnError() *dstore_values.BooleanValue {
	if m != nil {
		return m.CancelOnError
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	PersonId        *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	ErrorIdList     *dstore_values.StringValue                       `protobuf:"bytes,102,opt,name=error_id_list,json=errorIdList" json:"error_id_list,omitempty"`
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

func (m *Response) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response) GetErrorIdList() *dstore_values.StringValue {
	if m != nil {
		return m.ErrorIdList
	}
	return nil
}

type Response_Row struct {
	RowId                  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PersonCharacteristicId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	ResultCode             *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=result_code,json=resultCode" json:"result_code,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetResultCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ResultCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_CreateNewCommunityMember_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_CreateNewCommunityMember_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_CreateNewCommunityMember_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_CreateNewCommunityMember_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 782 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x56, 0xfe, 0x34, 0x6d, 0x3a, 0xe9, 0xd1, 0xfd, 0xff, 0xd6, 0x4d, 0xab, 0xaa, 0xea, 0xaf,
	0x4a, 0x20, 0x90, 0x0b, 0xe1, 0x02, 0x04, 0x6a, 0x55, 0x11, 0x81, 0x14, 0xa9, 0x0d, 0xc1, 0x17,
	0x95, 0xe0, 0xc6, 0x72, 0xed, 0x6d, 0xb0, 0x48, 0xbc, 0x61, 0xd7, 0xa1, 0xe2, 0x2d, 0x80, 0x57,
	0x81, 0x07, 0xe1, 0x35, 0x38, 0x1f, 0x9e, 0x80, 0xf1, 0xce, 0xa6, 0x89, 0xdd, 0x24, 0x26, 0x5c,
	0xb5, 0xeb, 0x9d, 0xef, 0x30, 0xab, 0x9d, 0x6f, 0x03, 0xfb, 0xbe, 0x8c, 0xb8, 0x60, 0x7b, 0x2c,
	0x6c, 0x06, 0x21, 0xdb, 0xeb, 0x08, 0xee, 0x31, 0xbf, 0x2b, 0x98, 0xdc, 0xf3, 0xb8, 0x53, 0x15,
	0xcc, 0x8d, 0x58, 0x9d, 0x9d, 0x57, 0x79, 0xbb, 0xdd, 0x0d, 0x83, 0xe8, 0xd5, 0x31, 0x6b, 0x9f,
	0x32, 0xe1, 0x34, 0xba, 0x16, 0x56, 0x46, 0xdc, 0xb8, 0x4e, 0x70, 0x8b, 0xe0, 0xd6, 0x78, 0x4c,
	0x79, 0x45, 0x8b, 0xbd, 0x74, 0x5b, 0x5d, 0x26, 0x89, 0xa2, 0xbc, 0x9e, 0x74, 0xc0, 0x84, 0xe0,
	0x42, 0x6f, 0x6d, 0x24, 0xb7, 0xda, 0x4c, 0x4a, 0xb7, 0xc9, 0xf4, 0xe6, 0xff, 0xe9, 0xcd, 0xc8,
	0x0d, 0xc2, 0x33, 0x2e, 0xda, 0x6e, 0x14, 0xf0, 0x90, 0x8a, 0x76, 0xde, 0x01, 0x40, 0xc3, 0x15,
	0x2e, 0xee, 0x32, 0x21, 0x8d, 0x03, 0x98, 0xf3, 0x7a, 0xb6, 0x9c, 0xc0, 0x37, 0x73, 0xdb, 0xb9,
	0x2b, 0xa5, 0xca, 0x86, 0xa5, 0xbb, 0xd0, 0xbe, 0x82, 0x30, 0x62, 0x4d, 0x26, 0x4e, 0xe2, 0x95,
	0x5d, 0xba, 0x00, 0xd4, 0x7c, 0xe3, 0x1a, 0x2c, 0x0f, 0xe2, 0x9d, 0xb0, 0xdb, 0x6a, 0x99, 0x1f,
	0x67, 0x90, 0xa5, 0x68, 0x2f, 0x0e, 0x14, 0xd6, 0xf1, 0xbb, 0x71, 0x1b, 0x66, 0x71, 0xf9, 0xa2,
	0xcb, 0x62, 0xa5, 0x7f, 0x94, 0x52, 0x39, 0xa5, 0x24, 0x23, 0x11, 0x84, 0x4d, 0x12, 0x2a, 0x52,
	0x31, 0xaa, 0xec, 0xc2, 0xc2, 0x05, 0x90, 0x24, 0x3e, 0x91, 0xc4, 0x5c, 0xaf, 0x44, 0xf1, 0xdb,
	0xb0, 0xe6, 0x3d, 0xc3, 0xde, 0x3c, 0x6c, 0x2d, 0x90, 0x51, 0xe0, 0xc5, 0xe5, 0x2d, 0xfc, 0xef,
	0xa6, 0x99, 0xcf, 0x54, 0xfb, 0x2f, 0x09, 0xad, 0xf9, 0x47, 0x31, 0xd0, 0x38, 0x84, 0xcd, 0x11,
	0x9c, 0x64, 0xe4, 0x33, 0x19, 0x59, 0x1f, 0x8a, 0x56, 0xae, 0xee, 0x41, 0x49, 0xc9, 0x69, 0x27,
	0x53, 0x99, 0x4e, 0x40, 0x7d, 0x23, 0xf9, 0xab, 0xb0, 0x34, 0x00, 0x26, 0xc9, 0x2f, 0x24, 0xb9,
	0xd0, 0x2f, 0x1b, 0xdf, 0x7d, 0xc5, 0x2c, 0xfc, 0x5d, 0xf7, 0x95, 0xd1, 0xdd, 0x57, 0xc8, 0xca,
	0xd7, 0x31, 0xdd, 0x57, 0x2e, 0x77, 0x5f, 0x31, 0xa7, 0x27, 0xe8, 0xbe, 0x92, 0xec, 0x5e, 0x4b,
	0x7e, 0x4b, 0x77, 0x4f, 0x3a, 0x27, 0xb0, 0xde, 0xc1, 0x0b, 0xcd, 0x43, 0x87, 0xbc, 0x38, 0x1e,
	0x0e, 0x5d, 0x93, 0x0b, 0x75, 0xab, 0x67, 0xb2, 0x6f, 0xf5, 0x2a, 0xa1, 0xab, 0x0a, 0x5c, 0xd5,
	0x58, 0xbc, 0x7a, 0x55, 0xd8, 0x1a, 0xc9, 0x4b, 0x86, 0xbe, 0x93, 0xa1, 0xf2, 0x70, 0x02, 0x65,
	0xee, 0x31, 0xac, 0x61, 0x80, 0x74, 0x5b, 0x91, 0x13, 0x84, 0x8e, 0x9a, 0xe7, 0xde, 0x39, 0x9a,
	0xc5, 0xa1, 0xd6, 0x4e, 0x39, 0x6f, 0x31, 0x37, 0x24, 0x6b, 0x2b, 0x84, 0xad, 0x85, 0x0f, 0x62,
	0x24, 0x1d, 0x2e, 0x0e, 0xee, 0xe6, 0x08, 0x4a, 0x72, 0xf5, 0x83, 0x5c, 0xad, 0x0d, 0xc1, 0x2a,
	0x4b, 0x1e, 0x6c, 0xd3, 0xd1, 0x06, 0xbe, 0x74, 0x30, 0x24, 0x9c, 0x8e, 0x60, 0x3e, 0x3b, 0xc3,
	0xe0, 0xf0, 0x75, 0xa7, 0xd2, 0x9c, 0xcd, 0xf6, 0xb6, 0xa9, 0x3e, 0xd6, 0x7c, 0xf9, 0x90, 0x8b,
	0xc6, 0x05, 0x03, 0x1d, 0x82, 0x34, 0x1a, 0xb0, 0x9b, 0x25, 0x42, 0x6e, 0x7f, 0x92, 0xdb, 0xed,
	0x71, 0x6c, 0xca, 0x76, 0x15, 0x16, 0x3d, 0x37, 0xf4, 0x58, 0xcb, 0xe1, 0xba, 0x6d, 0x13, 0xb2,
	0x5d, 0xce, 0x13, 0xe6, 0x11, 0x9d, 0x82, 0x71, 0x03, 0xfe, 0x4d, 0x91, 0x90, 0x8b, 0x5f, 0xe4,
	0x62, 0x39, 0x51, 0x1d, 0xcb, 0xee, 0x7c, 0x98, 0x82, 0xa2, 0xcd, 0x64, 0x87, 0x87, 0x92, 0x21,
	0xbc, 0x40, 0xca, 0xb9, 0xe4, 0x65, 0xd6, 0x91, 0x4f, 0x79, 0xad, 0xb0, 0x36, 0x15, 0x1a, 0x4f,
	0x60, 0x29, 0x4e, 0x63, 0x67, 0x20, 0x8e, 0x31, 0xff, 0xf2, 0x08, 0xb6, 0x52, 0xe0, 0x74, 0x68,
	0x1f, 0xe3, 0xba, 0xd6, 0x5f, 0xdb, 0x8b, 0xed, 0xe4, 0x07, 0xe3, 0x0e, 0xcc, 0xe8, 0x57, 0x00,
	0x33, 0x2e, 0x66, 0xdc, 0xba, 0xc4, 0x48, 0x6f, 0xc4, 0x31, 0xfd, 0xb5, 0x7b, 0xe5, 0xc6, 0x11,
	0xe4, 0x05, 0x3f, 0xc7, 0x3c, 0x8a, 0x51, 0x77, 0xad, 0x49, 0xde, 0x2d, 0xab, 0x77, 0x16, 0x96,
	0xcd, 0xcf, 0xed, 0x98, 0x06, 0x7d, 0xcc, 0xea, 0x39, 0xc1, 0x79, 0x63, 0xd9, 0xf3, 0x56, 0xa4,
	0x6a, 0x9c, 0xb0, 0x03, 0x98, 0x4f, 0x8e, 0xc4, 0x59, 0x66, 0x46, 0x94, 0x58, 0xff, 0x36, 0x97,
	0xdf, 0xe7, 0x20, 0x8f, 0x36, 0x8c, 0x55, 0x98, 0x46, 0x23, 0xb1, 0xfc, 0xeb, 0x3a, 0x32, 0x14,
	0xec, 0x02, 0x2e, 0x91, 0xff, 0x04, 0xcc, 0xc4, 0x04, 0x0f, 0x44, 0x99, 0xf9, 0xa6, 0x3e, 0x61,
	0x32, 0xf4, 0x23, 0xce, 0xd8, 0x87, 0x92, 0x9e, 0x40, 0x8f, 0xfb, 0xcc, 0x7c, 0xfb, 0x07, 0x54,
	0x40, 0x80, 0x2a, 0xd6, 0xdf, 0x77, 0x60, 0x23, 0xe0, 0xa9, 0x53, 0xef, 0xff, 0xd8, 0x78, 0x7a,
	0xd8, 0xe4, 0xd2, 0x7f, 0xde, 0xdb, 0xf7, 0x27, 0xff, 0x3d, 0x72, 0x3a, 0xad, 0x1e, 0xfc, 0x5b,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x99, 0x93, 0xe1, 0xd1, 0x08, 0x00, 0x00,
}
