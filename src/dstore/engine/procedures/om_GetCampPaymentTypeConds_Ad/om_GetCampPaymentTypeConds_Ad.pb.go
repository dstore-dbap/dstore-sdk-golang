// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampPaymentTypeConds_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampPaymentTypeConds_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampPaymentTypeConds_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampPaymentTypeConds_Ad

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
	ConditionId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
	ConditionIdNull bool                        `protobuf:"varint,1001,opt,name=condition_id_null,json=conditionIdNull" json:"condition_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
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
	RowId                  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	FilterByPaymentTypeIds *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=filter_by_payment_type_ids,json=filterByPaymentTypeIds" json:"filter_by_payment_type_ids,omitempty"`
	ConditionId            *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
	NegateFilter           *dstore_values.BooleanValue `protobuf:"bytes,10003,opt,name=negate_filter,json=negateFilter" json:"negate_filter,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetFilterByPaymentTypeIds() *dstore_values.StringValue {
	if m != nil {
		return m.FilterByPaymentTypeIds
	}
	return nil
}

func (m *Response_Row) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
	}
	return nil
}

func (m *Response_Row) GetNegateFilter() *dstore_values.BooleanValue {
	if m != nil {
		return m.NegateFilter
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampPaymentTypeConds_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampPaymentTypeConds_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampPaymentTypeConds_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x6b, 0x14, 0x31,
	0x14, 0x65, 0x9d, 0x6e, 0x5b, 0xd2, 0x4a, 0x35, 0x42, 0x19, 0x67, 0x41, 0x4a, 0x7d, 0x11, 0x0a,
	0x59, 0x51, 0x10, 0x05, 0x51, 0x6a, 0x51, 0x59, 0xa4, 0x43, 0x09, 0x7e, 0xa0, 0x2f, 0x21, 0xdb,
	0xb9, 0x1d, 0x02, 0x33, 0xc9, 0x90, 0x64, 0x2d, 0xf3, 0x2f, 0xfc, 0x78, 0xf4, 0x17, 0xfa, 0x17,
	0x7c, 0x32, 0x33, 0xc9, 0xba, 0x33, 0xa3, 0x08, 0xdb, 0x97, 0x19, 0x6e, 0xee, 0x39, 0xe7, 0xde,
	0x7b, 0x6e, 0x82, 0x9e, 0x66, 0xc6, 0x2a, 0x0d, 0x53, 0x90, 0xb9, 0x90, 0x30, 0xad, 0xb4, 0x3a,
	0x87, 0x6c, 0xa1, 0xc1, 0x4c, 0x55, 0xc9, 0x5e, 0x83, 0x3d, 0xe1, 0x65, 0x75, 0xc6, 0xeb, 0x12,
	0xa4, 0x7d, 0x5b, 0x57, 0x70, 0xa2, 0x64, 0x66, 0xd8, 0x71, 0x46, 0x1c, 0xd0, 0x2a, 0x7c, 0xe4,
	0xd9, 0xc4, 0xb3, 0xc9, 0x7f, 0x29, 0xc9, 0xad, 0x50, 0xea, 0x33, 0x2f, 0x16, 0x60, 0xbc, 0x42,
	0x72, 0xbb, 0x5f, 0x1f, 0xb4, 0x56, 0x3a, 0xa4, 0x26, 0xfd, 0x54, 0x09, 0xc6, 0xf0, 0x1c, 0x42,
	0xf2, 0xee, 0x30, 0x69, 0xb9, 0x90, 0x17, 0x4a, 0x97, 0xdc, 0x0a, 0x25, 0x3d, 0xe8, 0xb0, 0x46,
	0xe8, 0x8c, 0x6b, 0xee, 0x92, 0xa0, 0x0d, 0x7e, 0x86, 0x76, 0xcf, 0x5d, 0x2f, 0xa2, 0x01, 0x30,
	0x91, 0xc5, 0xa3, 0x83, 0xd1, 0xbd, 0x9d, 0x07, 0x13, 0x12, 0x66, 0x08, 0x6d, 0x09, 0x69, 0x21,
	0x07, 0xfd, 0xbe, 0x89, 0xe8, 0xce, 0x1f, 0xc2, 0x2c, 0xc3, 0x47, 0xe8, 0x66, 0x97, 0xcf, 0xe4,
	0xa2, 0x28, 0xe2, 0x9f, 0x5b, 0x4e, 0x65, 0x9b, 0xee, 0x75, 0x80, 0xa9, 0x3b, 0x3f, 0xfc, 0xb1,
	0x81, 0xb6, 0x29, 0x98, 0x4a, 0x49, 0x03, 0xf8, 0x3e, 0x1a, 0xb7, 0x83, 0x85, 0x92, 0x09, 0xe9,
	0xdb, 0xe6, 0x87, 0x7e, 0xd9, 0x7c, 0xa9, 0x07, 0xe2, 0x8f, 0xe8, 0x46, 0x33, 0x12, 0xeb, 0xcc,
	0x14, 0x5f, 0x3b, 0x88, 0x1c, 0x99, 0x0c, 0xc8, 0xc3, 0xc9, 0x4f, 0x5d, 0x3c, 0x5b, 0xc5, 0x74,
	0xaf, 0xec, 0x1f, 0xe0, 0xc7, 0x68, 0x2b, 0x58, 0x19, 0x47, 0xad, 0xe2, 0x9d, 0xbf, 0x14, 0xbd,
	0xd1, 0xa7, 0xfe, 0x4f, 0x97, 0x70, 0xfc, 0x06, 0x45, 0x5a, 0x5d, 0xc6, 0x1b, 0x2d, 0xeb, 0x09,
	0x59, 0x63, 0xf7, 0x64, 0x69, 0x05, 0xa1, 0xea, 0x92, 0x36, 0x2a, 0xc9, 0xaf, 0x11, 0x8a, 0x5c,
	0x80, 0xf7, 0xd1, 0xa6, 0x0b, 0x9b, 0x7d, 0x7c, 0x49, 0x9d, 0x3b, 0x63, 0x3a, 0x76, 0xa1, 0x73,
	0xfb, 0x03, 0x4a, 0x2e, 0x44, 0xe1, 0x16, 0xc7, 0xe6, 0x35, 0xab, 0xbc, 0x24, 0xb3, 0x4e, 0xd3,
	0x41, 0x4d, 0xfc, 0x35, 0xed, 0x3b, 0x19, 0x96, 0x67, 0xac, 0x16, 0x32, 0xf7, 0xbb, 0xdb, 0xf7,
	0xf4, 0x17, 0x75, 0xa7, 0x9f, 0x59, 0x66, 0xf0, 0xf3, 0xc1, 0x35, 0xf8, 0x96, 0xae, 0x79, 0x0f,
	0x8e, 0xd1, 0x75, 0x09, 0x39, 0xb7, 0xc0, 0x7c, 0x85, 0xf8, 0xfb, 0xbf, 0x15, 0xe6, 0x4a, 0x15,
	0xc0, 0xa5, 0x57, 0xd8, 0xf5, 0x94, 0x57, 0xbe, 0xa7, 0x77, 0x68, 0x22, 0xd4, 0xc0, 0xc0, 0xd5,
	0xd3, 0xfb, 0xf4, 0xe8, 0x6a, 0x8f, 0x72, 0xbe, 0xd9, 0x5e, 0xfb, 0x87, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0xe6, 0xcc, 0x89, 0xd5, 0x03, 0x00, 0x00,
}