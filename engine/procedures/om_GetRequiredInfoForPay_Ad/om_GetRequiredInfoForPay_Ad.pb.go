// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetRequiredInfoForPay_Ad.proto
// DO NOT EDIT!

/*
Package om_GetRequiredInfoForPay_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetRequiredInfoForPay_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetRequiredInfoForPay_Ad

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
	PaymentForShippingId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=payment_for_shipping_id,json=paymentForShippingId" json:"payment_for_shipping_id,omitempty"`
	PaymentForShippingIdNull bool                        `protobuf:"varint,1001,opt,name=payment_for_shipping_id_null,json=paymentForShippingIdNull" json:"payment_for_shipping_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	PaymentForShippingId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=payment_for_shipping_id,json=paymentForShippingId" json:"payment_for_shipping_id,omitempty"`
	CategoryDescription           *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	PaymentForShippingDescription *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=payment_for_shipping_description,json=paymentForShippingDescription" json:"payment_for_shipping_description,omitempty"`
	PersonCharacCategoryId        *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=person_charac_category_id,json=personCharacCategoryId" json:"person_charac_category_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPaymentForShippingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentForShippingId
	}
	return nil
}

func (m *Response_Row) GetCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CategoryDescription
	}
	return nil
}

func (m *Response_Row) GetPaymentForShippingDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentForShippingDescription
	}
	return nil
}

func (m *Response_Row) GetPersonCharacCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacCategoryId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetRequiredInfoForPay_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetRequiredInfoForPay_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetRequiredInfoForPay_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetRequiredInfoForPay_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0x15, 0x42, 0xda, 0xca, 0x1c, 0x40, 0xdb, 0xaa, 0x6c, 0x13, 0x40, 0x51, 0xb9, 0xc0,
	0xc5, 0x41, 0x70, 0x00, 0x21, 0x21, 0x04, 0xe5, 0x43, 0x7b, 0xe8, 0xaa, 0x72, 0x25, 0x10, 0x08,
	0xc9, 0x32, 0xeb, 0x61, 0x6b, 0x91, 0xb5, 0x17, 0x7b, 0x43, 0x95, 0x23, 0x6f, 0xc0, 0xd7, 0x03,
	0xc0, 0xe3, 0xf1, 0x16, 0xcc, 0xae, 0x1d, 0x92, 0x4d, 0x5b, 0x28, 0xbd, 0x24, 0x9a, 0xcc, 0xfc,
	0xfe, 0x9e, 0xf9, 0x8f, 0x63, 0x72, 0x5f, 0xba, 0xca, 0x58, 0x18, 0x81, 0xce, 0x95, 0x86, 0x51,
	0x69, 0x4d, 0x06, 0x72, 0x62, 0xc1, 0x8d, 0x4c, 0xc1, 0x9f, 0x43, 0xc5, 0xe0, 0xc3, 0x44, 0x59,
	0x90, 0x89, 0x7e, 0x67, 0x9e, 0x19, 0xbb, 0x27, 0xa6, 0xfc, 0x91, 0xa4, 0x58, 0x56, 0x99, 0xe8,
	0xa6, 0x67, 0xa9, 0x67, 0xe9, 0x5f, 0x80, 0xfe, 0x7a, 0x38, 0xe6, 0xa3, 0x18, 0x4f, 0xc0, 0x79,
	0xbe, 0xbf, 0xd5, 0x3e, 0x1b, 0xac, 0x35, 0x36, 0xa4, 0x06, 0xed, 0x54, 0x01, 0xce, 0x89, 0x1c,
	0x42, 0xf2, 0xfa, 0x72, 0xb2, 0x12, 0x0a, 0x4f, 0xb3, 0x85, 0xa8, 0x94, 0xd1, 0xbe, 0x68, 0xfb,
	0x67, 0x87, 0x90, 0x3d, 0x61, 0x05, 0x66, 0xc1, 0xba, 0x88, 0x91, 0xcb, 0xa5, 0x98, 0x16, 0xa0,
	0x2b, 0x8e, 0x95, 0xdc, 0x1d, 0xa8, 0xb2, 0x54, 0x3a, 0xe7, 0x4a, 0xc6, 0x9d, 0x61, 0xe7, 0xc6,
	0x85, 0xdb, 0x03, 0x1a, 0xa6, 0x09, 0x2d, 0x2a, 0x5d, 0x41, 0x0e, 0xf6, 0x45, 0x1d, 0xb1, 0x8d,
	0xc0, 0xe2, 0x40, 0xfb, 0x81, 0x4c, 0x64, 0xf4, 0x90, 0x5c, 0x39, 0x41, 0x93, 0xeb, 0xc9, 0x78,
	0x1c, 0xff, 0x5a, 0x45, 0xe5, 0x35, 0x16, 0x1f, 0x07, 0xa7, 0x58, 0xb0, 0xfd, 0xa3, 0x47, 0xd6,
	0x18, 0xb8, 0xd2, 0x68, 0x07, 0xd1, 0x2d, 0xd2, 0x6b, 0x1c, 0x08, 0xfd, 0xf4, 0x69, 0xdb, 0x5d,
	0xef, 0xce, 0xd3, 0xfa, 0x93, 0xf9, 0xc2, 0xe8, 0x15, 0xb9, 0x54, 0xcf, 0xce, 0x17, 0x86, 0x8f,
	0xcf, 0x0d, 0xbb, 0x08, 0xd3, 0x25, 0x78, 0xd9, 0xa2, 0x5d, 0x8c, 0x93, 0x79, 0xcc, 0x2e, 0x16,
	0xed, 0x1f, 0xa2, 0x7b, 0x64, 0x35, 0x78, 0x1e, 0x77, 0x1b, 0xc5, 0x6b, 0x47, 0x14, 0xfd, 0x46,
	0x76, 0xfd, 0x37, 0x9b, 0x95, 0x47, 0x09, 0xe9, 0x5a, 0x73, 0x18, 0x9f, 0x6f, 0xa8, 0xbb, 0xf4,
	0xd4, 0x57, 0x84, 0xce, 0x8c, 0xa0, 0xcc, 0x1c, 0xb2, 0x5a, 0xa3, 0xff, 0xa9, 0x4b, 0xba, 0x18,
	0x44, 0x9b, 0x64, 0x05, 0xc3, 0x7a, 0x55, 0x9f, 0x53, 0xf4, 0xa6, 0xc7, 0x7a, 0x18, 0xa2, 0xff,
	0xfb, 0x27, 0xef, 0xf4, 0x4b, 0x7a, 0xd6, 0xa5, 0xa6, 0x64, 0x23, 0x13, 0x58, 0x65, 0xec, 0x94,
	0x4b, 0x70, 0x99, 0x55, 0x65, 0x63, 0xec, 0xd7, 0xb4, 0xbd, 0x96, 0xa0, 0xe8, 0x2a, 0x8b, 0x9c,
	0x17, 0x5c, 0x9f, 0x81, 0x4f, 0xe6, 0x5c, 0x24, 0xc9, 0xf0, 0xd8, 0x26, 0x17, 0xb5, 0xbf, 0xfd,
	0x5b, 0xfb, 0xea, 0xd1, 0x66, 0x17, 0x4f, 0x79, 0x49, 0xb6, 0x4a, 0xbc, 0xe6, 0x46, 0xf3, 0xec,
	0x00, 0x2f, 0x7d, 0xc6, 0xff, 0xcc, 0x80, 0x66, 0x7c, 0x3f, 0x85, 0x19, 0x9b, 0x1e, 0xdf, 0x69,
	0xe8, 0x9d, 0x00, 0x27, 0xf2, 0xf1, 0x1b, 0x32, 0x50, 0x66, 0x69, 0x8b, 0xf3, 0x47, 0xe2, 0xf5,
	0x83, 0xdc, 0x38, 0xf9, 0x7e, 0x96, 0x97, 0xff, 0xf9, 0x8e, 0xbc, 0x5d, 0x69, 0xfe, 0xab, 0x77,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x08, 0xf3, 0xec, 0x86, 0x04, 0x00, 0x00,
}
