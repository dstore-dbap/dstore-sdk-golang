// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetShippingCost_Pu.proto
// DO NOT EDIT!

/*
Package om_GetShippingCost_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetShippingCost_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetShippingCost_Pu

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
	ShippingTypeId                *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=shipping_type_id,json=shippingTypeId" json:"shipping_type_id,omitempty"`
	ShippingTypeIdNull            bool                          `protobuf:"varint,1001,opt,name=shipping_type_id_null,json=shippingTypeIdNull" json:"shipping_type_id_null,omitempty"`
	CurrencyId                    *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencyIdNull                bool                          `protobuf:"varint,1002,opt,name=currency_id_null,json=currencyIdNull" json:"currency_id_null,omitempty"`
	TotalBruttoPrice              *dstore_values.DecimalValue   `protobuf:"bytes,3,opt,name=total_brutto_price,json=totalBruttoPrice" json:"total_brutto_price,omitempty"`
	TotalBruttoPriceNull          bool                          `protobuf:"varint,1003,opt,name=total_brutto_price_null,json=totalBruttoPriceNull" json:"total_brutto_price_null,omitempty"`
	TotalNettoPrice               *dstore_values.DecimalValue   `protobuf:"bytes,4,opt,name=total_netto_price,json=totalNettoPrice" json:"total_netto_price,omitempty"`
	TotalNettoPriceNull           bool                          `protobuf:"varint,1004,opt,name=total_netto_price_null,json=totalNettoPriceNull" json:"total_netto_price_null,omitempty"`
	SelectResult                  *dstore_values.BooleanValue   `protobuf:"bytes,5,opt,name=select_result,json=selectResult" json:"select_result,omitempty"`
	SelectResultNull              bool                          `protobuf:"varint,1005,opt,name=select_result_null,json=selectResultNull" json:"select_result_null,omitempty"`
	Date                          *dstore_values.TimestampValue `protobuf:"bytes,6,opt,name=date" json:"date,omitempty"`
	DateNull                      bool                          `protobuf:"varint,1006,opt,name=date_null,json=dateNull" json:"date_null,omitempty"`
	UniqueId                      *dstore_values.StringValue    `protobuf:"bytes,7,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                  bool                          `protobuf:"varint,1007,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonId                      *dstore_values.IntegerValue   `protobuf:"bytes,8,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                  bool                          `protobuf:"varint,1008,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	DeliveryPersonId              *dstore_values.IntegerValue   `protobuf:"bytes,9,opt,name=delivery_person_id,json=deliveryPersonId" json:"delivery_person_id,omitempty"`
	DeliveryPersonIdNull          bool                          `protobuf:"varint,1009,opt,name=delivery_person_id_null,json=deliveryPersonIdNull" json:"delivery_person_id_null,omitempty"`
	PriceNodeCharacteristicId     *dstore_values.IntegerValue   `protobuf:"bytes,10,opt,name=price_node_characteristic_id,json=priceNodeCharacteristicId" json:"price_node_characteristic_id,omitempty"`
	PriceNodeCharacteristicIdNull bool                          `protobuf:"varint,1010,opt,name=price_node_characteristic_id_null,json=priceNodeCharacteristicIdNull" json:"price_node_characteristic_id_null,omitempty"`
	PaymentTypeId                 *dstore_values.IntegerValue   `protobuf:"bytes,11,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	PaymentTypeIdNull             bool                          `protobuf:"varint,1011,opt,name=payment_type_id_null,json=paymentTypeIdNull" json:"payment_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetShippingTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ShippingTypeId
	}
	return nil
}

func (m *Parameters) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func (m *Parameters) GetTotalBruttoPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalBruttoPrice
	}
	return nil
}

func (m *Parameters) GetTotalNettoPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalNettoPrice
	}
	return nil
}

func (m *Parameters) GetSelectResult() *dstore_values.BooleanValue {
	if m != nil {
		return m.SelectResult
	}
	return nil
}

func (m *Parameters) GetDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetDeliveryPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.DeliveryPersonId
	}
	return nil
}

func (m *Parameters) GetPriceNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PriceNodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetPaymentTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentTypeId
	}
	return nil
}

type Response struct {
	Error              *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation    []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message            []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row                []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ShippingCost       *dstore_values.DecimalValue                      `protobuf:"bytes,101,opt,name=shipping_cost,json=shippingCost" json:"shipping_cost,omitempty"`
	ShippingCostBrutto *dstore_values.DecimalValue                      `protobuf:"bytes,102,opt,name=shipping_cost_brutto,json=shippingCostBrutto" json:"shipping_cost_brutto,omitempty"`
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

func (m *Response) GetShippingCost() *dstore_values.DecimalValue {
	if m != nil {
		return m.ShippingCost
	}
	return nil
}

func (m *Response) GetShippingCostBrutto() *dstore_values.DecimalValue {
	if m != nil {
		return m.ShippingCostBrutto
	}
	return nil
}

type Response_Row struct {
	RowId              int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ShippingCost       *dstore_values.DecimalValue `protobuf:"bytes,10001,opt,name=shipping_cost,json=shippingCost" json:"shipping_cost,omitempty"`
	ShippingCostBrutto *dstore_values.DecimalValue `protobuf:"bytes,10002,opt,name=shipping_cost_brutto,json=shippingCostBrutto" json:"shipping_cost_brutto,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetShippingCost() *dstore_values.DecimalValue {
	if m != nil {
		return m.ShippingCost
	}
	return nil
}

func (m *Response_Row) GetShippingCostBrutto() *dstore_values.DecimalValue {
	if m != nil {
		return m.ShippingCostBrutto
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetShippingCost_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetShippingCost_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetShippingCost_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetShippingCost_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 808 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x4e, 0xdb, 0x58,
	0x14, 0x16, 0x13, 0x02, 0xc9, 0xe1, 0x2f, 0x18, 0x86, 0x09, 0x01, 0x46, 0x33, 0xa0, 0x91, 0x66,
	0x16, 0x63, 0x0a, 0x45, 0xa5, 0x8b, 0x2e, 0x5a, 0x10, 0x42, 0x59, 0x10, 0x45, 0x6e, 0x85, 0xd4,
	0xaa, 0x92, 0x65, 0xec, 0x4b, 0xb0, 0xea, 0xf8, 0xba, 0xd7, 0xd7, 0xa0, 0xbc, 0x45, 0xdb, 0x55,
	0x1f, 0xa3, 0xcf, 0xd3, 0x37, 0xe8, 0xff, 0xdf, 0x0b, 0xf4, 0xf8, 0x1e, 0x3b, 0x89, 0x1d, 0xa8,
	0x69, 0x37, 0x58, 0xd7, 0xe7, 0xfb, 0xc3, 0x1c, 0x3e, 0x1b, 0x76, 0x9c, 0x50, 0x72, 0xc1, 0x36,
	0x99, 0xdf, 0x71, 0x7d, 0xb6, 0x19, 0x08, 0x6e, 0x33, 0x27, 0x12, 0x2c, 0xdc, 0xe4, 0x5d, 0xf3,
	0x90, 0xc9, 0xfb, 0x67, 0x6e, 0x10, 0xb8, 0x7e, 0x67, 0x9f, 0x87, 0xd2, 0x6c, 0x47, 0x3a, 0x02,
	0x24, 0xd7, 0x36, 0x88, 0xa5, 0x13, 0x4b, 0xbf, 0x14, 0xda, 0x58, 0x48, 0xa4, 0xcf, 0x2d, 0x2f,
	0x62, 0x21, 0x31, 0x1b, 0xcb, 0x59, 0x3f, 0x26, 0x04, 0x17, 0xc9, 0x68, 0x25, 0x3b, 0xea, 0xb2,
	0x30, 0xb4, 0x3a, 0x2c, 0x19, 0x6e, 0xe4, 0x87, 0xd2, 0x72, 0xfd, 0x53, 0x2e, 0xba, 0x96, 0x74,
	0xb9, 0x4f, 0xa0, 0xf5, 0x97, 0x00, 0xd0, 0xb6, 0x84, 0x85, 0x53, 0x26, 0x42, 0xed, 0x00, 0x6a,
	0x61, 0x92, 0xc9, 0x94, 0xbd, 0x80, 0x99, 0xae, 0x53, 0x1f, 0xfb, 0x6b, 0xec, 0xdf, 0xa9, 0xed,
	0x15, 0x3d, 0xf9, 0x05, 0x92, 0x6c, 0xae, 0x2f, 0x59, 0x87, 0x89, 0xe3, 0xf8, 0x64, 0xcc, 0xa6,
	0xa4, 0x07, 0xc8, 0x69, 0x3a, 0xda, 0x36, 0xfc, 0x9e, 0x97, 0x31, 0xfd, 0xc8, 0xf3, 0xea, 0x6f,
	0x26, 0x51, 0xac, 0x62, 0x68, 0x59, 0x7c, 0x0b, 0x47, 0xda, 0x1d, 0x98, 0xb2, 0x23, 0x21, 0x98,
	0x6f, 0xf7, 0x62, 0xd7, 0xdf, 0x8a, 0x5d, 0x21, 0xc5, 0xa3, 0xe3, 0x7f, 0x50, 0x1b, 0x62, 0x93,
	0xd9, 0x5b, 0x32, 0x9b, 0x1d, 0xc0, 0x94, 0x51, 0x13, 0x34, 0xc9, 0xa5, 0xe5, 0x99, 0x27, 0x22,
	0x92, 0x92, 0x9b, 0x81, 0x70, 0x6d, 0x56, 0x2f, 0x5d, 0xea, 0xe7, 0x30, 0xdb, 0xed, 0x5a, 0x1e,
	0xf9, 0xd5, 0x14, 0x6d, 0x4f, 0xb1, 0xda, 0x31, 0x49, 0xbb, 0x05, 0x7f, 0x8c, 0x4a, 0x91, 0xf9,
	0x3b, 0x32, 0x5f, 0xcc, 0x73, 0x54, 0x84, 0x43, 0x98, 0x27, 0x9e, 0xcf, 0x06, 0x09, 0xc6, 0x8b,
	0x13, 0xcc, 0x29, 0x56, 0x8b, 0xf5, 0x03, 0xec, 0xc0, 0xd2, 0x88, 0x10, 0xf9, 0xbf, 0x27, 0xff,
	0x85, 0x1c, 0x43, 0xd9, 0xdf, 0x85, 0x99, 0x90, 0x79, 0xcc, 0x96, 0x26, 0x6e, 0x6d, 0xe4, 0xc9,
	0x7a, 0xf9, 0x52, 0xeb, 0x13, 0xce, 0x3d, 0x66, 0xf9, 0x64, 0x3d, 0x4d, 0x0c, 0x43, 0x11, 0xb4,
	0xff, 0x41, 0xcb, 0x28, 0x90, 0xe7, 0x07, 0xf2, 0xac, 0x0d, 0x43, 0x95, 0xe1, 0x16, 0x8c, 0x3b,
	0x96, 0x64, 0xf5, 0x09, 0xe5, 0xb3, 0x96, 0xf3, 0x91, 0x2e, 0x2e, 0xae, 0xb4, 0xba, 0x01, 0x39,
	0x29, 0xa8, 0xb6, 0x0a, 0xd5, 0xf8, 0x4a, 0xc2, 0x1f, 0x49, 0xb8, 0x12, 0xdf, 0x51, 0x82, 0xbb,
	0x50, 0x8d, 0x7c, 0xf7, 0x69, 0xa4, 0x16, 0x74, 0x52, 0xa9, 0x36, 0x72, 0xaa, 0xa1, 0x14, 0xb8,
	0x60, 0x24, 0x59, 0x21, 0x30, 0xee, 0xc9, 0x3f, 0x30, 0xdb, 0x27, 0x92, 0xf6, 0x27, 0xd2, 0x9e,
	0x4e, 0x21, 0x4a, 0xff, 0x36, 0x54, 0x03, 0xfc, 0x7f, 0xe0, 0x7e, 0xac, 0x5f, 0x29, 0x5e, 0xc5,
	0x0a, 0xa1, 0xc9, 0xa0, 0xcf, 0x24, 0x83, 0xcf, 0x89, 0x41, 0x0a, 0x49, 0x97, 0xd0, 0x61, 0x9e,
	0x7b, 0xce, 0x44, 0xcf, 0x1c, 0x38, 0x55, 0x8b, 0x9d, 0x6a, 0x29, 0xad, 0x9d, 0x3a, 0xe2, 0x12,
	0x8e, 0x4a, 0x91, 0xf5, 0x97, 0x64, 0x09, 0xf3, 0x1c, 0x15, 0xe1, 0x31, 0xac, 0x26, 0xfb, 0xc2,
	0x1d, 0x66, 0xda, 0x67, 0xd8, 0x02, 0x36, 0x96, 0x80, 0x1b, 0x4a, 0xd7, 0x8e, 0xc3, 0x40, 0x71,
	0x98, 0x65, 0x25, 0xd0, 0x42, 0xfe, 0x7e, 0x86, 0x8e, 0xa9, 0x9a, 0xf0, 0xf7, 0x8f, 0xd4, 0x29,
	0xdf, 0x57, 0xca, 0xb7, 0x76, 0xa5, 0x8c, 0x0a, 0xba, 0x0f, 0x73, 0x81, 0xd5, 0xeb, 0x32, 0x5f,
	0xf6, 0x3b, 0x69, 0xaa, 0x38, 0xdb, 0x4c, 0xc2, 0x49, 0x2a, 0xe9, 0x06, 0x2c, 0xe6, 0x44, 0x28,
	0xc2, 0x37, 0x8a, 0x30, 0x9f, 0x41, 0xc7, 0xb6, 0xeb, 0xaf, 0xc7, 0xa1, 0x82, 0x3b, 0x1c, 0x70,
	0x3f, 0x64, 0x48, 0x2f, 0xab, 0xe2, 0x4d, 0xda, 0xb0, 0xbf, 0x6c, 0x49, 0x9d, 0x53, 0x29, 0x1f,
	0xc4, 0x3f, 0x0d, 0x02, 0x6a, 0x0f, 0xa1, 0x16, 0x57, 0xae, 0x39, 0xd4, 0xb9, 0x58, 0x6a, 0x25,
	0x24, 0xeb, 0x39, 0x72, 0xbe, 0x99, 0x8f, 0xf0, 0xdc, 0x1c, 0x9c, 0x8d, 0xb9, 0x6e, 0xf6, 0x06,
	0x6e, 0xe7, 0x64, 0x52, 0xf5, 0x58, 0x5b, 0xb1, 0xe2, 0x9f, 0x23, 0x8a, 0xf4, 0x22, 0x38, 0xa2,
	0xab, 0x91, 0xc2, 0xf1, 0x51, 0x96, 0x04, 0xbf, 0xc0, 0xaa, 0x89, 0x59, 0x5b, 0xfa, 0x35, 0xde,
	0x49, 0x7a, 0xfa, 0x08, 0x74, 0x83, 0x5f, 0x18, 0x31, 0x5b, 0xd5, 0x47, 0xda, 0xee, 0x36, 0xa2,
	0xea, 0xac, 0xb8, 0xb9, 0xa6, 0xc3, 0x21, 0x59, 0xed, 0x08, 0x16, 0x33, 0x0a, 0x49, 0x7f, 0xd6,
	0x4f, 0x8b, 0x85, 0xb4, 0x61, 0x21, 0xea, 0xd5, 0xc6, 0xab, 0x31, 0x28, 0x61, 0x3a, 0x6d, 0x09,
	0x26, 0x30, 0x5f, 0xbc, 0x1f, 0xcf, 0x5a, 0xa8, 0x54, 0x36, 0xca, 0x78, 0xc4, 0xbf, 0xfd, 0xbd,
	0x7c, 0xe0, 0xe7, 0xad, 0x9f, 0x4d, 0xdc, 0xba, 0x22, 0xf1, 0x8b, 0xd6, 0x2f, 0x45, 0xde, 0x3b,
	0x86, 0x15, 0x97, 0xe7, 0x9e, 0xff, 0xe0, 0x4b, 0xe2, 0xd1, 0x6e, 0x87, 0x87, 0xce, 0x93, 0x74,
	0xee, 0x5c, 0xfb, 0x63, 0xe3, 0x64, 0x42, 0xbd, 0xd6, 0x6f, 0x7e, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x06, 0xe3, 0x59, 0xb7, 0xa5, 0x08, 0x00, 0x00,
}