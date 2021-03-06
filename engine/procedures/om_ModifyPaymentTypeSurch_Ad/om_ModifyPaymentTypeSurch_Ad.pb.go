// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyPaymentTypeSurch_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyPaymentTypeSurch_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyPaymentTypeSurch_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyPaymentTypeSurch_Ad

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
	PaymentTypeId           *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	PaymentTypeIdNull       bool                          `protobuf:"varint,1001,opt,name=payment_type_id_null,json=paymentTypeIdNull" json:"payment_type_id_null,omitempty"`
	SurchargeTypeId         *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeIdNull     bool                          `protobuf:"varint,1002,opt,name=surcharge_type_id_null,json=surchargeTypeIdNull" json:"surcharge_type_id_null,omitempty"`
	SurchargeValue          *dstore_values.DecimalValue   `protobuf:"bytes,3,opt,name=surcharge_value,json=surchargeValue" json:"surcharge_value,omitempty"`
	SurchargeValueNull      bool                          `protobuf:"varint,1003,opt,name=surcharge_value_null,json=surchargeValueNull" json:"surcharge_value_null,omitempty"`
	ValidFrom               *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	ValidFromNull           bool                          `protobuf:"varint,1004,opt,name=valid_from_null,json=validFromNull" json:"valid_from_null,omitempty"`
	PriorityNo              *dstore_values.IntegerValue   `protobuf:"bytes,5,opt,name=priority_no,json=priorityNo" json:"priority_no,omitempty"`
	PriorityNoNull          bool                          `protobuf:"varint,1005,opt,name=priority_no_null,json=priorityNoNull" json:"priority_no_null,omitempty"`
	DeleteConfiguration     *dstore_values.BooleanValue   `protobuf:"bytes,6,opt,name=delete_configuration,json=deleteConfiguration" json:"delete_configuration,omitempty"`
	DeleteConfigurationNull bool                          `protobuf:"varint,1006,opt,name=delete_configuration_null,json=deleteConfigurationNull" json:"delete_configuration_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPaymentTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentTypeId
	}
	return nil
}

func (m *Parameters) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetSurchargeValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.SurchargeValue
	}
	return nil
}

func (m *Parameters) GetValidFrom() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *Parameters) GetPriorityNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PriorityNo
	}
	return nil
}

func (m *Parameters) GetDeleteConfiguration() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteConfiguration
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyPaymentTypeSurch_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyPaymentTypeSurch_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyPaymentTypeSurch_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyPaymentTypeSurch_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd5, 0x86, 0xa4, 0x65, 0xaa, 0x36, 0xed, 0x26, 0x2a, 0x6e, 0xa2, 0x22, 0x54, 0x2e,
	0x38, 0x5c, 0x38, 0x9c, 0x2e, 0x2a, 0x15, 0x21, 0x41, 0x39, 0x28, 0x48, 0x89, 0x2a, 0x83, 0x90,
	0x40, 0x42, 0xd6, 0x36, 0xbb, 0x31, 0x2b, 0xec, 0x5d, 0x6b, 0xed, 0x50, 0xe5, 0x2d, 0x78, 0x20,
	0x1e, 0x84, 0x57, 0xe0, 0xf8, 0x0c, 0xac, 0x77, 0x9c, 0x38, 0x76, 0x23, 0x0a, 0x37, 0x49, 0x26,
	0xf3, 0xff, 0xf3, 0x8d, 0xd7, 0x33, 0x0b, 0x47, 0x2c, 0x49, 0x95, 0xe6, 0x3d, 0x2e, 0x03, 0x21,
	0x79, 0x2f, 0xd6, 0x6a, 0xc4, 0xd9, 0x44, 0xf3, 0xa4, 0xa7, 0x22, 0x7f, 0xa0, 0x98, 0x18, 0x4f,
	0x4f, 0xe8, 0x34, 0xe2, 0x32, 0x7d, 0x3d, 0x8d, 0xf9, 0xab, 0x89, 0x1e, 0x7d, 0xf0, 0x1f, 0x33,
	0xd7, 0xe8, 0x52, 0x45, 0x6e, 0xa3, 0xd9, 0x45, 0xb3, 0xfb, 0x37, 0x47, 0xa7, 0x95, 0x83, 0x3e,
	0xd1, 0x70, 0xc2, 0x13, 0x2c, 0xd0, 0xd9, 0x2b, 0xd3, 0xb9, 0xd6, 0x4a, 0xe7, 0xa9, 0x6e, 0x39,
	0x15, 0xf1, 0x24, 0xa1, 0x01, 0xcf, 0x93, 0xd7, 0xab, 0xc9, 0x94, 0x0a, 0x39, 0x56, 0x3a, 0xa2,
	0xa9, 0x50, 0x12, 0x45, 0x07, 0x5f, 0xeb, 0x00, 0x27, 0x54, 0x53, 0x93, 0xe5, 0x3a, 0x21, 0xc7,
	0xd0, 0x8c, 0xb1, 0x2f, 0x3f, 0x35, 0x8d, 0xf9, 0x82, 0x39, 0x2b, 0xd7, 0x56, 0x6e, 0x6e, 0xdc,
	0xeb, 0xba, 0xf9, 0x63, 0xe4, 0xad, 0x09, 0x99, 0xf2, 0x80, 0xeb, 0x37, 0x59, 0xe4, 0x6d, 0xc6,
	0xc5, 0xb3, 0xf4, 0x19, 0xb9, 0x03, 0xed, 0x4a, 0x11, 0x5f, 0x4e, 0xc2, 0xd0, 0xf9, 0xb6, 0x66,
	0x4a, 0xad, 0x7b, 0x3b, 0x25, 0xf5, 0xd0, 0x64, 0xc8, 0x0b, 0xd8, 0x49, 0xb2, 0x33, 0xa0, 0x3a,
	0xe0, 0x73, 0xf0, 0xea, 0xc5, 0xe0, 0xe6, 0xdc, 0x95, 0xa3, 0x1f, 0xc0, 0xee, 0xb9, 0x42, 0x08,
	0xff, 0x8e, 0xf0, 0x56, 0xc5, 0x61, 0xf1, 0x4f, 0xa1, 0x28, 0xe4, 0x5b, 0x8e, 0x53, 0x5b, 0x0a,
	0x67, 0x7c, 0x24, 0x22, 0x1a, 0x22, 0x7c, 0x6b, 0xee, 0xb1, 0x31, 0xb9, 0x0b, 0xed, 0x4a, 0x15,
	0x24, 0xff, 0x40, 0x32, 0x29, 0xcb, 0x2d, 0xf8, 0x21, 0x80, 0x11, 0x9a, 0x16, 0xc7, 0x5a, 0x45,
	0xce, 0x25, 0xcb, 0xdc, 0xaf, 0x30, 0x53, 0x61, 0x5e, 0x6b, 0x4a, 0xa3, 0x18, 0xa9, 0x97, 0xad,
	0xe1, 0xb9, 0xd1, 0x93, 0x1b, 0xd0, 0x2c, 0xdc, 0xc8, 0xfa, 0x89, 0xac, 0xcd, 0xb9, 0x28, 0xc7,
	0x6c, 0xc4, 0x5a, 0x28, 0x2d, 0xd2, 0xa9, 0x2f, 0x95, 0x53, 0xbf, 0xf8, 0x60, 0x61, 0xa6, 0x1f,
	0x2a, 0x72, 0x0b, 0xb6, 0x17, 0xdc, 0xc8, 0xf9, 0x85, 0x9c, 0xad, 0x42, 0x66, 0x41, 0x43, 0x68,
	0x33, 0x1e, 0x9a, 0x51, 0xf2, 0x47, 0x4a, 0x8e, 0x45, 0x30, 0xd1, 0x76, 0xd6, 0x9c, 0xc6, 0x52,
	0xe2, 0xa9, 0x52, 0x21, 0xa7, 0x12, 0x89, 0x2d, 0x34, 0x1e, 0x2f, 0xfa, 0xc8, 0x11, 0xec, 0x2d,
	0xab, 0x87, 0x3d, 0xfc, 0xc6, 0x1e, 0xae, 0x2c, 0x31, 0x66, 0xcd, 0x1c, 0x7c, 0x59, 0x85, 0x75,
	0x8f, 0x27, 0xb1, 0x92, 0x09, 0x37, 0x33, 0x59, 0xb7, 0x8b, 0x93, 0x8f, 0x73, 0xc7, 0x2d, 0x6f,
	0x25, 0x2e, 0xd5, 0xb3, 0xec, 0xd3, 0x43, 0x21, 0x79, 0x0b, 0xdb, 0xd9, 0xca, 0xf8, 0x0b, 0x3b,
	0x63, 0x46, 0xb2, 0x66, 0xcc, 0x6e, 0xc5, 0x5c, 0xdd, 0xac, 0x81, 0x89, 0xfb, 0x45, 0xec, 0x35,
	0xa3, 0xf2, 0x1f, 0xe4, 0x10, 0xd6, 0xf2, 0x55, 0x35, 0x73, 0x96, 0x55, 0xbc, 0x7a, 0xae, 0x22,
	0x2e, 0xf2, 0x00, 0xbf, 0xbd, 0x99, 0x9c, 0xbc, 0x84, 0x9a, 0x56, 0x67, 0x66, 0x52, 0x32, 0xd7,
	0xa1, 0xfb, 0xef, 0x57, 0x8b, 0x3b, 0x3b, 0x09, 0xd7, 0x53, 0x67, 0x5e, 0x56, 0xa4, 0xb3, 0x0f,
	0x35, 0xf3, 0x9b, 0xec, 0x42, 0xc3, 0x44, 0xd9, 0xc2, 0x7d, 0x1e, 0x9a, 0xb3, 0xa9, 0x7b, 0x75,
	0x13, 0xf6, 0xd9, 0x93, 0xf7, 0xd0, 0x15, 0xaa, 0x42, 0x28, 0x6e, 0xbe, 0x77, 0x8f, 0x02, 0x95,
	0xb0, 0x8f, 0xb3, 0x3c, 0xfb, 0xdf, 0xcb, 0xf1, 0xb4, 0x61, 0xef, 0x9f, 0xfb, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0x6b, 0x35, 0x6c, 0x5c, 0x05, 0x00, 0x00,
}
