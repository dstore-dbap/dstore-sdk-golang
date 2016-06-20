// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetGroupPaymentForShip_Ad.proto
// DO NOT EDIT!

/*
Package om_GetGroupPaymentForShip_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetGroupPaymentForShip_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetGroupPaymentForShip_Ad

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
	GroupId                  *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	GroupIdNull              bool                        `protobuf:"varint,1001,opt,name=group_id_null,json=groupIdNull" json:"group_id_null,omitempty"`
	PaymentForShippingId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=payment_for_shipping_id,json=paymentForShippingId" json:"payment_for_shipping_id,omitempty"`
	PaymentForShippingIdNull bool                        `protobuf:"varint,1002,opt,name=payment_for_shipping_id_null,json=paymentForShippingIdNull" json:"payment_for_shipping_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *Parameters) GetPaymentForShippingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentForShippingId
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
	RowId                         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	GroupDescription              *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=group_description,json=groupDescription" json:"group_description,omitempty"`
	PaymentForShippingId          *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=payment_for_shipping_id,json=paymentForShippingId" json:"payment_for_shipping_id,omitempty"`
	DescriptionForAdmin           *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=description_for_admin,json=descriptionForAdmin" json:"description_for_admin,omitempty"`
	PaymentForShippingDescription *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=payment_for_shipping_description,json=paymentForShippingDescription" json:"payment_for_shipping_description,omitempty"`
	GroupId                       *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetGroupDescription() *dstore_values.StringValue {
	if m != nil {
		return m.GroupDescription
	}
	return nil
}

func (m *Response_Row) GetPaymentForShippingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentForShippingId
	}
	return nil
}

func (m *Response_Row) GetDescriptionForAdmin() *dstore_values.StringValue {
	if m != nil {
		return m.DescriptionForAdmin
	}
	return nil
}

func (m *Response_Row) GetPaymentForShippingDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentForShippingDescription
	}
	return nil
}

func (m *Response_Row) GetGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetGroupPaymentForShip_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetGroupPaymentForShip_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetGroupPaymentForShip_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x09, 0x5d, 0xab, 0x33, 0x21, 0x86, 0xc7, 0x4f, 0x68, 0x01, 0x4d, 0xdb, 0x0d, 0xe2,
	0xc2, 0x45, 0x20, 0x60, 0x12, 0x17, 0x68, 0x08, 0x18, 0x43, 0x5a, 0x98, 0x3c, 0x84, 0x04, 0x37,
	0x51, 0x58, 0x4c, 0xb0, 0xd4, 0xd8, 0x91, 0x9d, 0x32, 0xf1, 0x16, 0xfc, 0xbe, 0x05, 0x8f, 0xc2,
	0x8b, 0xc0, 0x43, 0x20, 0x8e, 0x63, 0x57, 0x4d, 0xb2, 0x8d, 0x4e, 0xbd, 0x69, 0x75, 0x72, 0xce,
	0x77, 0xce, 0x77, 0xbe, 0xcf, 0x36, 0x3c, 0x4c, 0x4d, 0xa9, 0x34, 0x1f, 0x71, 0x99, 0x09, 0xc9,
	0x47, 0x85, 0x56, 0x07, 0x3c, 0x9d, 0x68, 0x6e, 0x46, 0x2a, 0x8f, 0xb7, 0x79, 0xb9, 0xad, 0xd5,
	0xa4, 0xd8, 0x4b, 0x3e, 0xe5, 0x5c, 0x96, 0xcf, 0x94, 0xde, 0xff, 0x20, 0x8a, 0x78, 0x2b, 0xa5,
	0x58, 0x57, 0x2a, 0x72, 0xcb, 0x81, 0xa9, 0x03, 0xd3, 0xff, 0x21, 0x06, 0xab, 0x7e, 0xd0, 0xc7,
	0x64, 0x3c, 0xe1, 0xc6, 0x35, 0x18, 0x5c, 0x6d, 0x4e, 0xe7, 0x5a, 0x2b, 0xed, 0x53, 0xc3, 0x66,
	0x2a, 0xe7, 0xc6, 0x24, 0x19, 0xf7, 0xc9, 0x8d, 0x76, 0xb2, 0x4c, 0x84, 0x7c, 0xaf, 0x74, 0x9e,
	0x94, 0x42, 0x49, 0x57, 0xb4, 0xfe, 0xb7, 0x03, 0xb0, 0x97, 0xe8, 0x04, 0xb3, 0x5c, 0x1b, 0x72,
	0x1f, 0xfa, 0x99, 0xa5, 0x16, 0x8b, 0x34, 0xec, 0xac, 0x75, 0x6e, 0x2e, 0xdf, 0x19, 0x52, 0xcf,
	0xdf, 0x73, 0x12, 0xb2, 0xe4, 0x19, 0xd7, 0xaf, 0x6d, 0xc4, 0x7a, 0x55, 0xf1, 0x4e, 0x4a, 0x36,
	0xe0, 0xdc, 0x14, 0x17, 0xcb, 0xc9, 0x78, 0x1c, 0xfe, 0xee, 0x21, 0xba, 0xcf, 0x96, 0x7d, 0x41,
	0x84, 0xdf, 0x08, 0x83, 0x2b, 0x85, 0x5b, 0x39, 0x46, 0x1a, 0xb1, 0xc1, 0xa5, 0x0b, 0x21, 0x33,
	0x3b, 0xeb, 0xcc, 0xfc, 0x59, 0x17, 0x8b, 0x86, 0x5c, 0x16, 0x89, 0x83, 0x1f, 0xc1, 0xb5, 0x13,
	0x7a, 0x3a, 0x1e, 0x7f, 0x1c, 0x8f, 0xf0, 0x38, 0xb0, 0x25, 0xb5, 0xfe, 0xab, 0x0b, 0x7d, 0xc6,
	0x4d, 0xa1, 0xa4, 0xe1, 0xe4, 0x36, 0x74, 0x2b, 0x79, 0xfd, 0xee, 0x03, 0xda, 0xf4, 0xce, 0x49,
	0xff, 0xd4, 0xfe, 0x32, 0x57, 0x48, 0xde, 0xc0, 0x8a, 0x15, 0x36, 0xae, 0x29, 0x8b, 0xcb, 0x04,
	0x08, 0xa6, 0x2d, 0x70, 0x5b, 0xff, 0x5d, 0x8c, 0x77, 0x66, 0x31, 0x3b, 0x9f, 0x37, 0x3f, 0x90,
	0x4d, 0xe8, 0x79, 0x43, 0xc3, 0xa0, 0xea, 0x78, 0xe3, 0x48, 0x47, 0x67, 0xf7, 0xae, 0xfb, 0x67,
	0xd3, 0x72, 0xf2, 0x02, 0x02, 0xad, 0x0e, 0xc3, 0xb3, 0x15, 0x6a, 0x93, 0x9e, 0xfe, 0x00, 0xd2,
	0xa9, 0x12, 0x94, 0xa9, 0x43, 0x66, 0x9b, 0x0c, 0x7e, 0x06, 0x10, 0x60, 0x40, 0x2e, 0xc3, 0x12,
	0x86, 0xd6, 0xab, 0xcf, 0x11, 0x8a, 0xd3, 0x65, 0x5d, 0x0c, 0xd1, 0x80, 0xe7, 0x70, 0xc1, 0x39,
	0x9f, 0x72, 0x73, 0xa0, 0x45, 0x51, 0x29, 0xf0, 0x25, 0x6a, 0xea, 0xe7, 0xfd, 0x34, 0xa5, 0x46,
	0xe1, 0x9d, 0x9d, 0x2b, 0x15, 0xea, 0xc9, 0x0c, 0x44, 0xf6, 0x4f, 0x3e, 0x1e, 0x5f, 0xa3, 0x45,
	0xcf, 0xc7, 0x4b, 0xb8, 0x54, 0x23, 0x56, 0x35, 0x4e, 0xd2, 0x5c, 0xc8, 0xf0, 0xdb, 0x7c, 0x8a,
	0xab, 0x35, 0x24, 0x76, 0xdd, 0xb2, 0x38, 0x92, 0xc2, 0xda, 0xb1, 0x2c, 0xeb, 0xeb, 0x7f, 0x9f,
	0xdf, 0xfb, 0xfa, 0x51, 0xb6, 0x75, 0x2d, 0x1e, 0xd4, 0xee, 0xe1, 0x8f, 0xe8, 0xf4, 0x17, 0xf1,
	0xf1, 0x2b, 0x18, 0x0a, 0xd5, 0x72, 0x7c, 0xf6, 0x5e, 0xbd, 0xbd, 0xb7, 0xd0, 0x4b, 0xf6, 0x6e,
	0xa9, 0x7a, 0x2c, 0xee, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x16, 0xbd, 0xcd, 0x90, 0x09, 0x05,
	0x00, 0x00,
}