// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetPersonInfoForOrders_Ad.proto
// DO NOT EDIT!

/*
Package om_GetPersonInfoForOrders_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetPersonInfoForOrders_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetPersonInfoForOrders_Ad

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
	OrderId                 *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	OrderIdNull             bool                          `protobuf:"varint,1001,opt,name=order_id_null,json=orderIdNull" json:"order_id_null,omitempty"`
	LanguageId              *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull          bool                          `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	DateAndTime             *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=date_and_time,json=dateAndTime" json:"date_and_time,omitempty"`
	DateAndTimeNull         bool                          `protobuf:"varint,1003,opt,name=date_and_time_null,json=dateAndTimeNull" json:"date_and_time_null,omitempty"`
	GetActualProperties     *dstore_values.BooleanValue   `protobuf:"bytes,4,opt,name=get_actual_properties,json=getActualProperties" json:"get_actual_properties,omitempty"`
	GetActualPropertiesNull bool                          `protobuf:"varint,1004,opt,name=get_actual_properties_null,json=getActualPropertiesNull" json:"get_actual_properties_null,omitempty"`
	DateFormat              *dstore_values.StringValue    `protobuf:"bytes,5,opt,name=date_format,json=dateFormat" json:"date_format,omitempty"`
	DateFormatNull          bool                          `protobuf:"varint,1005,opt,name=date_format_null,json=dateFormatNull" json:"date_format_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.DateAndTime
	}
	return nil
}

func (m *Parameters) GetGetActualProperties() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetActualProperties
	}
	return nil
}

func (m *Parameters) GetDateFormat() *dstore_values.StringValue {
	if m != nil {
		return m.DateFormat
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
	ValueRestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value_restricted_by_pattern,json=valueRestrictedByPattern" json:"value_restricted_by_pattern,omitempty"`
	Description              *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=description" json:"description,omitempty"`
	PersonId                 *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	InvoiceOrDeliveryAdress  *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=invoice_or_delivery_adress,json=invoiceOrDeliveryAdress" json:"invoice_or_delivery_adress,omitempty"`
	ActualValue              *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=actual_value,json=actualValue" json:"actual_value,omitempty"`
	PersonCharacteristicId   *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	Value                    *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=value" json:"value,omitempty"`
	SortNo                   *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	PersonTypeId             *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValueRestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.ValueRestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetInvoiceOrDeliveryAdress() *dstore_values.StringValue {
	if m != nil {
		return m.InvoiceOrDeliveryAdress
	}
	return nil
}

func (m *Response_Row) GetActualValue() *dstore_values.StringValue {
	if m != nil {
		return m.ActualValue
	}
	return nil
}

func (m *Response_Row) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetPersonInfoForOrders_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetPersonInfoForOrders_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetPersonInfoForOrders_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetPersonInfoForOrders_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x6a, 0x1b, 0x47,
	0x14, 0xc6, 0x95, 0x65, 0xa9, 0x23, 0xdf, 0x18, 0x53, 0x7b, 0x2b, 0xd1, 0x52, 0xec, 0x3f, 0x6d,
	0x29, 0xeb, 0xb6, 0x21, 0xc1, 0xc1, 0x89, 0x41, 0x4e, 0xe2, 0xa0, 0x80, 0x65, 0xb1, 0x18, 0x93,
	0x18, 0xc2, 0xb0, 0xd6, 0x4c, 0x94, 0x21, 0xd2, 0xcc, 0x32, 0x33, 0xb2, 0xd1, 0x5b, 0x24, 0x71,
	0xae, 0xcf, 0x91, 0x7f, 0x79, 0xa3, 0xdc, 0xde, 0x21, 0x67, 0x2e, 0xb2, 0x2e, 0x31, 0x91, 0xf2,
	0x47, 0xe2, 0xec, 0x7c, 0xdf, 0xf9, 0xbe, 0x73, 0xe6, 0x9c, 0x5d, 0xb4, 0x4d, 0xb5, 0x91, 0x8a,
	0x6d, 0x32, 0xd1, 0xe2, 0x82, 0x6d, 0x66, 0x4a, 0x36, 0x19, 0xed, 0x2a, 0xa6, 0x37, 0x65, 0x87,
	0xdc, 0x65, 0xa6, 0xc1, 0x94, 0x96, 0xa2, 0x26, 0x1e, 0xc9, 0x3d, 0xa9, 0x0e, 0x14, 0x85, 0x90,
	0x54, 0x69, 0x0c, 0x38, 0x23, 0xf1, 0xdf, 0x9e, 0x1c, 0x7b, 0x72, 0xfc, 0x3d, 0x46, 0x79, 0x25,
	0x08, 0x9d, 0xa6, 0xed, 0x2e, 0xd3, 0x3e, 0x41, 0xf9, 0xd7, 0x51, 0x75, 0xa6, 0x94, 0x54, 0xe1,
	0xa8, 0x32, 0x7a, 0xd4, 0x61, 0x5a, 0xa7, 0x2d, 0x16, 0x0e, 0x37, 0xc6, 0x0f, 0x4d, 0xca, 0x41,
	0x50, 0x75, 0x52, 0xc3, 0xa5, 0xf0, 0xa0, 0xf5, 0xf7, 0xb3, 0x08, 0x35, 0x52, 0x95, 0xc2, 0x29,
	0x78, 0xc0, 0xd7, 0x50, 0x51, 0x5a, 0x37, 0x84, 0xd3, 0x68, 0xe6, 0x8f, 0x99, 0x3f, 0x4b, 0xff,
	0x57, 0xe2, 0xe0, 0x3f, 0x78, 0xe2, 0xc2, 0xb0, 0x16, 0x53, 0x47, 0x36, 0x4a, 0x0a, 0x0e, 0x5c,
	0xa3, 0x78, 0x03, 0x2d, 0xf4, 0x79, 0x44, 0x74, 0xdb, 0xed, 0xe8, 0x43, 0x01, 0xd8, 0xc5, 0xa4,
	0x14, 0x00, 0x75, 0x78, 0x86, 0x6f, 0xa0, 0x52, 0x3b, 0x15, 0xad, 0x2e, 0x58, 0xb4, 0xf9, 0x7f,
	0x9a, 0x9c, 0x1f, 0xf5, 0xf1, 0x20, 0xf1, 0x17, 0x5a, 0x1e, 0x62, 0x7b, 0x95, 0x8f, 0x5e, 0x65,
	0x71, 0x00, 0x73, 0x42, 0x55, 0xb4, 0x40, 0x53, 0xc3, 0x48, 0x2a, 0x28, 0x31, 0xbc, 0xc3, 0xa2,
	0x9c, 0x93, 0xfa, 0x6d, 0x4c, 0xca, 0x1e, 0x69, 0x93, 0x76, 0x32, 0x2f, 0x56, 0xb2, 0x9c, 0xaa,
	0xa0, 0x87, 0xf0, 0x18, 0xff, 0x83, 0xf0, 0x48, 0x0a, 0xaf, 0xf7, 0xc9, 0xeb, 0x2d, 0x0d, 0x21,
	0x9d, 0xe0, 0x01, 0xfa, 0xa5, 0xc5, 0x0c, 0x49, 0x9b, 0xa6, 0x9b, 0xb6, 0x09, 0x74, 0x36, 0x63,
	0xca, 0x70, 0xa6, 0xa3, 0xd9, 0x4b, 0x6b, 0x3c, 0x91, 0xb2, 0xcd, 0x52, 0xe1, 0x65, 0x57, 0x80,
	0x59, 0x75, 0xc4, 0xc6, 0x05, 0x0f, 0x5a, 0x55, 0xbe, 0x34, 0xa1, 0xb7, 0xf1, 0xd9, 0xdb, 0x58,
	0xbb, 0x84, 0xe9, 0xec, 0x6c, 0x23, 0x57, 0x0b, 0xf1, 0x97, 0x1d, 0xe5, 0x9d, 0x89, 0xf2, 0x98,
	0x09, 0x6d, 0x14, 0x17, 0xad, 0xd0, 0x67, 0x0b, 0xdf, 0x73, 0x68, 0xdb, 0xe7, 0x21, 0xb2, 0x17,
	0xfc, 0x12, 0xfa, 0x3c, 0x80, 0x59, 0x9d, 0xf5, 0x77, 0x05, 0x54, 0x4c, 0x98, 0xce, 0xa4, 0xd0,
	0x0c, 0xff, 0x8b, 0xf2, 0x6e, 0x34, 0xc3, 0xdc, 0x5c, 0xc8, 0x85, 0xb9, 0xf7, 0x63, 0x7b, 0xc7,
	0xfe, 0x26, 0x1e, 0x88, 0x1f, 0xa0, 0x65, 0x3b, 0x94, 0x64, 0x68, 0x2a, 0x61, 0x28, 0x72, 0x40,
	0x8e, 0xc7, 0xc8, 0xe3, 0xb3, 0xbb, 0x0f, 0x71, 0x6d, 0x10, 0x27, 0x4b, 0x9d, 0xd1, 0x07, 0x78,
	0x0b, 0x15, 0xc2, 0x32, 0xc0, 0xdd, 0xdb, 0x8c, 0xbf, 0x7f, 0x93, 0xd1, 0xaf, 0xca, 0xbe, 0xff,
	0x4f, 0xfa, 0x70, 0x7c, 0x0f, 0xe5, 0x94, 0x3c, 0x83, 0x8b, 0xb3, 0xac, 0xad, 0x78, 0xfa, 0xe5,
	0x8d, 0xfb, 0x9d, 0x88, 0x13, 0x79, 0x96, 0xd8, 0x24, 0xe5, 0xf3, 0x3c, 0xca, 0x41, 0x80, 0x57,
	0xd1, 0x1c, 0x84, 0x76, 0xe6, 0x9f, 0xd6, 0xa1, 0x39, 0xf9, 0x24, 0x0f, 0x21, 0x8c, 0xf4, 0x31,
	0xaa, 0xb8, 0xcb, 0x20, 0xf0, 0x2a, 0x81, 0xdb, 0x68, 0x1a, 0x46, 0xc9, 0x49, 0x8f, 0x64, 0xa9,
	0x81, 0x65, 0x14, 0xd1, 0xb3, 0xfa, 0xc4, 0x8b, 0x8b, 0xdc, 0xb3, 0xe4, 0x82, 0xbe, 0xdb, 0x6b,
	0x78, 0x32, 0xbe, 0x09, 0x33, 0xc0, 0x74, 0x53, 0xf1, 0xcc, 0xf5, 0xf5, 0xf9, 0xe4, 0x5c, 0xc3,
	0x78, 0x7c, 0x1d, 0xfd, 0x9c, 0xb9, 0x32, 0xad, 0xeb, 0xf3, 0xfa, 0xe4, 0x55, 0x2d, 0x7a, 0x38,
	0x54, 0x75, 0x1f, 0x95, 0xb9, 0x38, 0x95, 0xbc, 0xc9, 0x88, 0x54, 0x84, 0xb2, 0x36, 0x3f, 0x65,
	0xaa, 0x47, 0x52, 0x0a, 0x55, 0xea, 0xe8, 0xc5, 0x64, 0x23, 0x6b, 0x81, 0x7e, 0xa0, 0x6e, 0x07,
	0x72, 0xd5, 0x71, 0xf1, 0x0e, 0x9a, 0x0f, 0x1b, 0xe1, 0x58, 0xd1, 0xcb, 0x29, 0x8a, 0xf2, 0x04,
	0x17, 0xe0, 0x23, 0x14, 0x85, 0xa2, 0x9a, 0x8f, 0xe1, 0x9d, 0x07, 0x0d, 0x53, 0x5c, 0x1b, 0xde,
	0xb4, 0x35, 0xbe, 0x9a, 0xa2, 0xc6, 0x55, 0xcf, 0xbe, 0x35, 0x42, 0x86, 0x8a, 0xff, 0x43, 0x79,
	0x6f, 0xe8, 0xf5, 0x64, 0x43, 0x1e, 0x89, 0xaf, 0xa2, 0x82, 0x96, 0x0a, 0xd6, 0x4b, 0x46, 0x6f,
	0xa6, 0x50, 0x9e, 0xb3, 0xe0, 0xba, 0xc4, 0xbb, 0x68, 0x31, 0x54, 0x60, 0x7a, 0x99, 0x7b, 0x8b,
	0xbe, 0x9d, 0x82, 0x3d, 0xef, 0x39, 0x87, 0x40, 0xa9, 0xd1, 0xdd, 0x87, 0xa8, 0xc2, 0xe5, 0xd8,
	0x60, 0x0f, 0x3e, 0x69, 0xc7, 0x3b, 0x2d, 0xa9, 0xe9, 0x93, 0xfe, 0x39, 0xfd, 0xd1, 0xaf, 0xde,
	0xc9, 0x9c, 0xfb, 0xb0, 0x5c, 0xf9, 0x1a, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x27, 0xe4, 0x2f, 0x35,
	0x07, 0x00, 0x00,
}
