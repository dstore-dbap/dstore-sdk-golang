// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetSetsForBonItBenefits_Ad.proto
// DO NOT EDIT!

/*
Package om_GetSetsForBonItBenefits_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetSetsForBonItBenefits_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetSetsForBonItBenefits_Ad

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
	BenefitId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	BenefitIdNull bool                        `protobuf:"varint,1001,opt,name=benefit_id_null,json=benefitIdNull" json:"benefit_id_null,omitempty"`
	ItemSetId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	ItemSetIdNull bool                        `protobuf:"varint,1002,opt,name=item_set_id_null,json=itemSetIdNull" json:"item_set_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
	}
	return nil
}

func (m *Parameters) GetItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemSetId
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
	RowId                    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ItemConditionId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	BenefitId                *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	ItemSetId                *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	MaxQuantity              *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=max_quantity,json=maxQuantity" json:"max_quantity,omitempty"`
	ItemConditionDescription *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=item_condition_description,json=itemConditionDescription" json:"item_condition_description,omitempty"`
	SortNo                   *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Response_Row) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
	}
	return nil
}

func (m *Response_Row) GetItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemSetId
	}
	return nil
}

func (m *Response_Row) GetMaxQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxQuantity
	}
	return nil
}

func (m *Response_Row) GetItemConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionDescription
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetSetsForBonItBenefits_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetSetsForBonItBenefits_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetSetsForBonItBenefits_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xed, 0x8a, 0x13, 0x31,
	0x14, 0xa5, 0x8e, 0x6d, 0xd7, 0x5b, 0xa5, 0x6b, 0x04, 0x19, 0xa7, 0x20, 0xb2, 0xfe, 0x70, 0x41,
	0x98, 0x8a, 0xa2, 0xa8, 0xbb, 0x20, 0xd6, 0xcf, 0x22, 0x3b, 0x68, 0x44, 0x41, 0xff, 0x0c, 0xb3,
	0x9d, 0xbb, 0x25, 0xd0, 0x49, 0x6a, 0x92, 0xba, 0xfa, 0x16, 0x7e, 0xbf, 0x89, 0xaf, 0x22, 0xbe,
	0x82, 0x3e, 0x85, 0xc9, 0x24, 0xdd, 0xee, 0x8c, 0xa0, 0xed, 0xfe, 0x99, 0xe1, 0xe6, 0x9e, 0x73,
	0xef, 0xcd, 0xb9, 0x87, 0xc0, 0x76, 0xae, 0xb4, 0x90, 0xd8, 0x47, 0x3e, 0x66, 0x1c, 0xfb, 0x53,
	0x29, 0x46, 0x98, 0xcf, 0x24, 0xaa, 0xbe, 0x28, 0xd2, 0x47, 0xa8, 0x9f, 0xa3, 0x56, 0x0f, 0x85,
	0x1c, 0x08, 0x3e, 0xd4, 0x03, 0xe4, 0xb8, 0xc7, 0xb4, 0x4a, 0xef, 0xe6, 0xb1, 0x01, 0x6a, 0x41,
	0x2e, 0x3b, 0x76, 0xec, 0xd8, 0xf1, 0x3f, 0x29, 0xd1, 0x19, 0xdf, 0xea, 0x6d, 0x36, 0x99, 0xa1,
	0x72, 0x15, 0xa2, 0x73, 0xd5, 0xfe, 0x28, 0xa5, 0x90, 0x3e, 0xd5, 0xab, 0xa6, 0x0a, 0x54, 0x2a,
	0x1b, 0xa3, 0x4f, 0x5e, 0xac, 0x27, 0x75, 0xc6, 0xf8, 0x9e, 0x90, 0x45, 0xa6, 0x99, 0xe0, 0x0e,
	0xb4, 0xf1, 0xb3, 0x01, 0xf0, 0x34, 0x93, 0x99, 0xc9, 0xa2, 0x54, 0xe4, 0x36, 0xc0, 0xae, 0x9b,
	0x27, 0x65, 0x79, 0xd8, 0xb8, 0xd0, 0xd8, 0xec, 0x5c, 0xed, 0xc5, 0xfe, 0x0a, 0x7e, 0x2a, 0xc6,
	0x35, 0x8e, 0x51, 0xbe, 0xb4, 0x11, 0x3d, 0xe1, 0xe1, 0xc3, 0x9c, 0x5c, 0x82, 0xee, 0x82, 0x9b,
	0xf2, 0xd9, 0x64, 0x12, 0xfe, 0x6a, 0x9b, 0x0a, 0x6b, 0xf4, 0xd4, 0x01, 0x28, 0x31, 0xa7, 0x64,
	0x0b, 0x3a, 0x4c, 0x63, 0x91, 0x2a, 0x2c, 0xbb, 0x1c, 0x5b, 0xa2, 0x8b, 0xc5, 0x1b, 0xd1, 0x4c,
	0x97, 0x4d, 0x58, 0x3f, 0x44, 0x76, 0x6d, 0x7e, 0xfb, 0x36, 0x07, 0x28, 0xdb, 0x66, 0xe3, 0x47,
	0x13, 0xd6, 0x28, 0xaa, 0xa9, 0xe0, 0x0a, 0xc9, 0x15, 0x68, 0x96, 0xc2, 0xf9, 0x3b, 0x45, 0x71,
	0x75, 0x2d, 0x4e, 0xd4, 0x07, 0xf6, 0x4b, 0x1d, 0x90, 0xbc, 0x82, 0x75, 0x2b, 0x59, 0x7a, 0x48,
	0x33, 0x33, 0x6a, 0x60, 0xc8, 0x71, 0x8d, 0x5c, 0x57, 0x76, 0xc7, 0xc4, 0xc3, 0x45, 0x4c, 0xbb,
	0x45, 0xf5, 0x80, 0xdc, 0x84, 0xb6, 0x5f, 0x55, 0x18, 0x94, 0x15, 0xcf, 0xff, 0x55, 0xd1, 0x2d,
	0x72, 0xc7, 0xfd, 0xe9, 0x1c, 0x4e, 0x9e, 0x40, 0x20, 0xc5, 0x7e, 0x78, 0xbc, 0x64, 0xdd, 0x8a,
	0x57, 0xf0, 0x56, 0x3c, 0x97, 0x22, 0xa6, 0x62, 0x9f, 0xda, 0x2a, 0xd1, 0xf7, 0x00, 0x02, 0x13,
	0x90, 0xb3, 0xd0, 0x32, 0xa1, 0x5d, 0xc5, 0x87, 0xc4, 0xa8, 0xd3, 0xa4, 0x4d, 0x13, 0x1a, 0xa9,
	0x1f, 0xc3, 0xe9, 0x52, 0xea, 0x91, 0xe0, 0x39, 0xb3, 0x83, 0x5b, 0xc8, 0xc7, 0xe4, 0xff, 0xeb,
	0xea, 0x5a, 0xda, 0xbd, 0x39, 0xcb, 0x54, 0xda, 0xaa, 0xd8, 0xea, 0x53, 0xb2, 0x92, 0xaf, 0xb6,
	0xab, 0x76, 0xf9, 0x9c, 0xac, 0xe4, 0x97, 0x3b, 0x70, 0xb2, 0xc8, 0xde, 0xa5, 0x6f, 0x66, 0x19,
	0xd7, 0x4c, 0xbf, 0x0f, 0xbf, 0x2c, 0x41, 0xef, 0x18, 0xc6, 0x33, 0x4f, 0x30, 0x3e, 0x88, 0x6a,
	0x2a, 0xe4, 0xa8, 0x46, 0x92, 0x4d, 0x4b, 0x47, 0x7c, 0x4d, 0xaa, 0x7e, 0xf2, 0xe5, 0x94, 0x96,
	0x8c, 0x8f, 0x5d, 0xb5, 0xb0, 0xa2, 0xc6, 0xfd, 0x05, 0x99, 0x5c, 0x87, 0xb6, 0x12, 0x52, 0xa7,
	0x5c, 0x84, 0xdf, 0x96, 0x18, 0xab, 0x65, 0xc1, 0x89, 0x18, 0xbc, 0x80, 0x1e, 0x13, 0xb5, 0xdd,
	0x2f, 0x5e, 0xa5, 0xd7, 0x37, 0x8e, 0xf6, 0x5e, 0xed, 0xb6, 0xca, 0x17, 0xe1, 0xda, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x57, 0xa7, 0x99, 0xeb, 0xf0, 0x04, 0x00, 0x00,
}