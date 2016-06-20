// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyPredefinedValue_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyPredefinedValue_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyPredefinedValue_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyPredefinedValue_Ad

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
	ValueId                       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	ValueIdNull                   bool                        `protobuf:"varint,1001,opt,name=value_id_null,json=valueIdNull" json:"value_id_null,omitempty"`
	NewValue                      *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=new_value,json=newValue" json:"new_value,omitempty"`
	NewValueNull                  bool                        `protobuf:"varint,1002,opt,name=new_value_null,json=newValueNull" json:"new_value_null,omitempty"`
	SortNo                        *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull                    bool                        `protobuf:"varint,1003,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	ProhibitValueChangeIfUsed     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=prohibit_value_change_if_used,json=prohibitValueChangeIfUsed" json:"prohibit_value_change_if_used,omitempty"`
	ProhibitValueChangeIfUsedNull bool                        `protobuf:"varint,1004,opt,name=prohibit_value_change_if_used_null,json=prohibitValueChangeIfUsedNull" json:"prohibit_value_change_if_used_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func (m *Parameters) GetNewValue() *dstore_values.StringValue {
	if m != nil {
		return m.NewValue
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetProhibitValueChangeIfUsed() *dstore_values.BooleanValue {
	if m != nil {
		return m.ProhibitValueChangeIfUsed
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyPredefinedValue_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyPredefinedValue_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyPredefinedValue_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x5d, 0x9b, 0xc4, 0xd3, 0xfa, 0xc3, 0x08, 0xb2, 0x4d, 0x88, 0x68, 0x8a, 0xa0, 0x37,
	0x1b, 0xd1, 0x62, 0xc5, 0x3b, 0x15, 0x2f, 0x56, 0x48, 0x28, 0x2b, 0x0a, 0x0a, 0xb2, 0x6c, 0x3a,
	0x67, 0xb7, 0x03, 0xc9, 0x4c, 0x99, 0xd9, 0x18, 0x7c, 0x0b, 0xdf, 0xc7, 0x67, 0xf0, 0x41, 0xfc,
	0x79, 0x08, 0xcf, 0xfc, 0x6c, 0x6b, 0x56, 0xad, 0xe2, 0x4d, 0x32, 0xe7, 0xe7, 0xfb, 0xce, 0x97,
	0x33, 0xdf, 0x04, 0x1e, 0x73, 0x53, 0x2b, 0x8d, 0x63, 0x94, 0x95, 0x90, 0x38, 0x3e, 0xd1, 0xea,
	0x08, 0xf9, 0x52, 0xa3, 0x19, 0x8b, 0x45, 0x3e, 0x51, 0x5c, 0x94, 0x1f, 0x0e, 0x35, 0x72, 0x2c,
	0xa9, 0xca, 0x5f, 0x17, 0xf3, 0x25, 0xe6, 0x4f, 0x78, 0x42, 0x6d, 0xb5, 0x62, 0x77, 0x3d, 0x36,
	0xf1, 0xd8, 0xe4, 0x1c, 0x40, 0xff, 0x5a, 0x18, 0xf3, 0xde, 0x26, 0x8c, 0xc7, 0xf7, 0x77, 0xd7,
	0x67, 0xa3, 0xd6, 0x4a, 0x87, 0xd2, 0x60, 0xbd, 0xb4, 0x40, 0x63, 0x8a, 0x0a, 0x43, 0x71, 0xaf,
	0x5d, 0xac, 0x0b, 0x21, 0x4b, 0xa5, 0x17, 0x45, 0x2d, 0x94, 0xf4, 0x4d, 0xa3, 0xcf, 0x11, 0xc0,
	0x61, 0xa1, 0x0b, 0xaa, 0xa2, 0x36, 0xec, 0x21, 0xf4, 0xdc, 0xec, 0x5c, 0xf0, 0x78, 0xe3, 0xe6,
	0xc6, 0x9d, 0xed, 0xfb, 0x83, 0x24, 0xc8, 0x0f, 0x9a, 0x84, 0xac, 0xb1, 0x42, 0xed, 0x24, 0x67,
	0x5d, 0x97, 0x4c, 0x39, 0xdb, 0x83, 0x4b, 0x0d, 0x2e, 0x97, 0xcb, 0xf9, 0x3c, 0xfe, 0xd2, 0x25,
	0x74, 0x2f, 0xdb, 0x0e, 0x0d, 0x53, 0xca, 0xb1, 0x03, 0xb8, 0x28, 0x71, 0x95, 0xbb, 0x54, 0xbc,
	0xe9, 0xd8, 0xfb, 0x2d, 0x76, 0x53, 0x6b, 0x21, 0x2b, 0x4f, 0xde, 0xa3, 0x66, 0x77, 0x62, 0xb7,
	0xe1, 0xf2, 0x29, 0xd0, 0xd3, 0x7f, 0xf5, 0xf4, 0x3b, 0x4d, 0x8b, 0xe3, 0xdf, 0x87, 0xae, 0x51,
	0xba, 0xce, 0xa5, 0x8a, 0xa3, 0xbf, 0x6b, 0xef, 0xd8, 0xde, 0xa9, 0x62, 0xb7, 0x60, 0x27, 0xa0,
	0x3c, 0xf5, 0x37, 0x4f, 0x0d, 0xbe, 0xec, 0x88, 0xdf, 0xc1, 0x90, 0xb6, 0x75, 0x2c, 0x66, 0xa2,
	0x0e, 0x22, 0x8e, 0x8e, 0x0b, 0x59, 0xd1, 0xaf, 0x2d, 0xf3, 0xa5, 0x41, 0x1e, 0x5f, 0xf8, 0xed,
	0xb8, 0x99, 0x52, 0x73, 0x2c, 0xa4, 0x1f, 0xb7, 0xdb, 0x30, 0xb8, 0xf0, 0x99, 0xc3, 0xa7, 0xe5,
	0x2b, 0x42, 0xb3, 0x17, 0x30, 0x3a, 0x97, 0xde, 0xeb, 0xfa, 0xee, 0x75, 0x0d, 0xff, 0xc8, 0x63,
	0xa5, 0x8e, 0x3e, 0x6d, 0x42, 0x2f, 0x43, 0x73, 0xa2, 0xa4, 0x41, 0x76, 0x0f, 0xb6, 0x9c, 0x5b,
	0xc2, 0x55, 0x9e, 0x2e, 0x3b, 0x38, 0xd1, 0x3b, 0xe9, 0xb9, 0xfd, 0xcc, 0x7c, 0x23, 0x7b, 0x03,
	0x57, 0xad, 0x4f, 0xf2, 0x9f, 0x8c, 0x42, 0x37, 0x15, 0x11, 0x38, 0x69, 0x81, 0xdb, 0x76, 0x9a,
	0x50, 0x9c, 0x9e, 0xc5, 0xd9, 0x95, 0xc5, 0x7a, 0x82, 0x3d, 0x82, 0x6e, 0xf0, 0x27, 0xdd, 0x8e,
	0x65, 0xbc, 0xf1, 0x0b, 0xa3, 0x77, 0xef, 0xc4, 0x7f, 0x67, 0x4d, 0x3b, 0x4b, 0x21, 0xd2, 0x6a,
	0x45, 0x4b, 0xb6, 0xa8, 0x83, 0xe4, 0x9f, 0x9f, 0x53, 0xd2, 0x2c, 0x22, 0xc9, 0xd4, 0x2a, 0xb3,
	0x1c, 0xfd, 0x21, 0x44, 0x74, 0x66, 0xd7, 0xa1, 0x43, 0x91, 0x35, 0xf9, 0xc7, 0x29, 0xad, 0x66,
	0x2b, 0xdb, 0xa2, 0x30, 0xe5, 0x4f, 0x5f, 0xc2, 0x40, 0xa8, 0xd6, 0x80, 0xb3, 0xb7, 0xfe, 0x76,
	0xff, 0x7f, 0xfe, 0x05, 0x66, 0x1d, 0xf7, 0xd2, 0x1e, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd8,
	0x48, 0xfa, 0xe6, 0x44, 0x04, 0x00, 0x00,
}