// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyPayForShipDescr_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyPayForShipDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyPayForShipDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyPayForShipDescr_Ad

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
	PaymentForShippingId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=payment_for_shipping_id,json=paymentForShippingId" json:"payment_for_shipping_id,omitempty"`
	PaymentForShippingIdNull          bool                        `protobuf:"varint,1001,opt,name=payment_for_shipping_id_null,json=paymentForShippingIdNull" json:"payment_for_shipping_id_null,omitempty"`
	LanguageId                        *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                    bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	PaymentForShippingDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=payment_for_shipping_description,json=paymentForShippingDescription" json:"payment_for_shipping_description,omitempty"`
	PaymentForShippingDescriptionNull bool                        `protobuf:"varint,1003,opt,name=payment_for_shipping_description_null,json=paymentForShippingDescriptionNull" json:"payment_for_shipping_description_null,omitempty"`
	DeleteTranslation                 *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_translation,json=deleteTranslation" json:"delete_translation,omitempty"`
	DeleteTranslationNull             bool                        `protobuf:"varint,1004,opt,name=delete_translation_null,json=deleteTranslationNull" json:"delete_translation_null,omitempty"`
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

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetPaymentForShippingDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentForShippingDescription
	}
	return nil
}

func (m *Parameters) GetDeleteTranslation() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteTranslation
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyPayForShipDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyPayForShipDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyPayForShipDescr_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xed, 0x6a, 0x13, 0x41,
	0x14, 0xa5, 0x4d, 0xd3, 0x96, 0x5b, 0xd0, 0x3a, 0x7e, 0x74, 0x4d, 0xac, 0xd4, 0x8a, 0x60, 0xff,
	0x6c, 0x44, 0x85, 0x8a, 0x08, 0xa2, 0x54, 0x21, 0x42, 0x4a, 0x9d, 0x8a, 0xa0, 0x7f, 0x96, 0x69,
	0x66, 0xba, 0x0e, 0xec, 0xce, 0x2c, 0x33, 0x1b, 0x4b, 0xde, 0xc2, 0xf7, 0xf1, 0x01, 0x7c, 0x16,
	0x3f, 0x1e, 0xc2, 0x3b, 0x3b, 0x13, 0xb7, 0xd9, 0xc4, 0x46, 0xfc, 0x93, 0x64, 0x72, 0xcf, 0xb9,
	0xe7, 0xcc, 0xfd, 0x18, 0x78, 0xca, 0x6d, 0xa9, 0x8d, 0xe8, 0x09, 0x95, 0x4a, 0x25, 0x7a, 0x85,
	0xd1, 0x43, 0xc1, 0x47, 0x46, 0xd8, 0x9e, 0xce, 0x93, 0x81, 0xe6, 0xf2, 0x74, 0x7c, 0xc4, 0xc6,
	0xaf, 0xb5, 0x39, 0xfe, 0x24, 0x8b, 0x03, 0x61, 0x87, 0x26, 0x79, 0xc1, 0x63, 0x84, 0x95, 0x9a,
	0xec, 0x79, 0x6e, 0xec, 0xb9, 0xf1, 0x05, 0x84, 0xce, 0xd5, 0x20, 0xf3, 0x99, 0x65, 0x23, 0x61,
	0x3d, 0xbf, 0x73, 0x73, 0x5a, 0x5b, 0x18, 0xa3, 0x4d, 0x08, 0x75, 0xa7, 0x43, 0xb9, 0xb0, 0x96,
	0xa5, 0x22, 0x04, 0xef, 0x36, 0x83, 0x25, 0x93, 0xea, 0x54, 0x9b, 0x9c, 0x95, 0x52, 0x2b, 0x0f,
	0xda, 0xfd, 0xb6, 0x02, 0x70, 0xc4, 0x0c, 0xc3, 0xa8, 0x30, 0x96, 0x50, 0xd8, 0x2a, 0xd8, 0x38,
	0x17, 0xaa, 0x4c, 0x10, 0x99, 0x58, 0x34, 0x57, 0x48, 0x95, 0x26, 0x92, 0x47, 0x4b, 0x3b, 0x4b,
	0xf7, 0x37, 0x1e, 0x76, 0xe3, 0x70, 0x9b, 0x60, 0x51, 0xaa, 0x52, 0xa4, 0xc2, 0xbc, 0x77, 0x27,
	0x7a, 0x2d, 0x70, 0xc3, 0xb5, 0x1c, 0xb3, 0xcf, 0xc9, 0x73, 0xb8, 0xf5, 0x97, 0x9c, 0x89, 0x1a,
	0x65, 0x59, 0xf4, 0x7d, 0x0d, 0x33, 0xaf, 0xd3, 0x68, 0x1e, 0xf9, 0x10, 0x01, 0xe4, 0x19, 0x6c,
	0x64, 0x4c, 0xa5, 0x23, 0xbc, 0x9a, 0x33, 0xb2, 0xbc, 0xd8, 0x08, 0x4c, 0xf0, 0x28, 0xbf, 0x07,
	0x9b, 0xe7, 0xd8, 0x5e, 0xf2, 0x87, 0x97, 0xbc, 0x54, 0xc3, 0x2a, 0xa1, 0x21, 0xec, 0xcc, 0x75,
	0xca, 0x5d, 0x7f, 0x64, 0xe1, 0xca, 0x16, 0xb5, 0x2a, 0xf5, 0x4e, 0x43, 0xdd, 0x96, 0x06, 0x81,
	0x5e, 0x7c, 0x7b, 0xf6, 0x22, 0x07, 0x75, 0x02, 0xf2, 0x16, 0xee, 0x2d, 0x12, 0xf1, 0x26, 0x7f,
	0x7a, 0x93, 0x77, 0x2e, 0x4c, 0x57, 0xf9, 0x7e, 0x03, 0x84, 0x8b, 0x0c, 0x3b, 0x98, 0x94, 0x86,
	0x29, 0x9b, 0x55, 0x0d, 0x8e, 0x56, 0xe6, 0xd6, 0xe9, 0x44, 0xeb, 0x4c, 0x30, 0xe5, 0xad, 0x5e,
	0xf1, 0xb4, 0x77, 0x35, 0x8b, 0xec, 0xc3, 0xd6, 0x6c, 0x2e, 0x6f, 0xe8, 0x97, 0x37, 0x74, 0x7d,
	0x86, 0xe4, 0x4c, 0xec, 0x7e, 0x5d, 0x86, 0x75, 0x2a, 0x6c, 0xa1, 0x95, 0x15, 0xe4, 0x01, 0xb4,
	0xab, 0x39, 0x0d, 0x53, 0xf3, 0xa7, 0x5c, 0x61, 0x07, 0xfc, 0x0c, 0xbf, 0x72, 0x9f, 0xd4, 0x03,
	0xc9, 0x07, 0xd8, 0x74, 0x13, 0x9a, 0x9c, 0x1b, 0x51, 0xec, 0x74, 0x0b, 0xc9, 0x71, 0x83, 0xdc,
	0x1c, 0xe4, 0x01, 0x9e, 0xfb, 0xf5, 0x99, 0x5e, 0xce, 0xa7, 0xff, 0x20, 0x4f, 0x60, 0x2d, 0x6c,
	0x06, 0x76, 0xcf, 0x65, 0xbc, 0x3d, 0x93, 0xd1, 0xef, 0xcd, 0xc0, 0x7f, 0xd3, 0x09, 0x9c, 0xf4,
	0xa1, 0x65, 0xf4, 0x19, 0x56, 0xd2, 0xb1, 0xf6, 0xe3, 0x7f, 0x5e, 0xe4, 0x78, 0x52, 0x88, 0x98,
	0xea, 0x33, 0xea, 0x72, 0x74, 0xb6, 0xa1, 0x85, 0xbf, 0xc9, 0x0d, 0x58, 0xc5, 0x93, 0x1b, 0xe3,
	0x2f, 0x87, 0x58, 0x9a, 0x36, 0x6d, 0xe3, 0xb1, 0xcf, 0x5f, 0x1e, 0x43, 0x57, 0xea, 0x86, 0x40,
	0xfd, 0xca, 0x7c, 0x7c, 0xfc, 0x3f, 0xef, 0xcf, 0xc9, 0x6a, 0xb5, 0xe3, 0x8f, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x7c, 0x6a, 0x07, 0x8a, 0xbe, 0x04, 0x00, 0x00,
}