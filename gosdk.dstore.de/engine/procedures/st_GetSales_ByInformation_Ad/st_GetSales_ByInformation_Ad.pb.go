// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetSales_ByInformation_Ad.proto
// DO NOT EDIT!

/*
Package st_GetSales_ByInformation_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetSales_ByInformation_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetSales_ByInformation_Ad

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
	FromDate                      *dstore_values.TimestampValue `protobuf:"bytes,1,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                  bool                          `protobuf:"varint,1001,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                        *dstore_values.TimestampValue `protobuf:"bytes,2,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                    bool                          `protobuf:"varint,1002,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	InformationTypeId             *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull         bool                          `protobuf:"varint,1003,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
	InformationIsValidInHours     *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=information_is_valid_in_hours,json=informationIsValidInHours" json:"information_is_valid_in_hours,omitempty"`
	InformationIsValidInHoursNull bool                          `protobuf:"varint,1004,opt,name=information_is_valid_in_hours_null,json=informationIsValidInHoursNull" json:"information_is_valid_in_hours_null,omitempty"`
	SelectAffectedOrderIds        *dstore_values.BooleanValue   `protobuf:"bytes,5,opt,name=select_affected_order_ids,json=selectAffectedOrderIds" json:"select_affected_order_ids,omitempty"`
	SelectAffectedOrderIdsNull    bool                          `protobuf:"varint,1005,opt,name=select_affected_order_ids_null,json=selectAffectedOrderIdsNull" json:"select_affected_order_ids_null,omitempty"`
	OrderDesc                     *dstore_values.BooleanValue   `protobuf:"bytes,6,opt,name=order_desc,json=orderDesc" json:"order_desc,omitempty"`
	OrderDescNull                 bool                          `protobuf:"varint,1006,opt,name=order_desc_null,json=orderDescNull" json:"order_desc_null,omitempty"`
	Information                   *dstore_values.StringValue    `protobuf:"bytes,7,opt,name=information" json:"information,omitempty"`
	InformationNull               bool                          `protobuf:"varint,1007,opt,name=information_null,json=informationNull" json:"information_null,omitempty"`
	PurchaseOrder                 *dstore_values.BooleanValue   `protobuf:"bytes,8,opt,name=purchase_order,json=purchaseOrder" json:"purchase_order,omitempty"`
	PurchaseOrderNull             bool                          `protobuf:"varint,1008,opt,name=purchase_order_null,json=purchaseOrderNull" json:"purchase_order_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetFromDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Parameters) GetToDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Parameters) GetInformationIsValidInHours() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationIsValidInHours
	}
	return nil
}

func (m *Parameters) GetSelectAffectedOrderIds() *dstore_values.BooleanValue {
	if m != nil {
		return m.SelectAffectedOrderIds
	}
	return nil
}

func (m *Parameters) GetOrderDesc() *dstore_values.BooleanValue {
	if m != nil {
		return m.OrderDesc
	}
	return nil
}

func (m *Parameters) GetInformation() *dstore_values.StringValue {
	if m != nil {
		return m.Information
	}
	return nil
}

func (m *Parameters) GetPurchaseOrder() *dstore_values.BooleanValue {
	if m != nil {
		return m.PurchaseOrder
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
	RowId            int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Information      *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=information" json:"information,omitempty"`
	NettoOrderSales  *dstore_values.DecimalValue   `protobuf:"bytes,10002,opt,name=netto_order_sales,json=nettoOrderSales" json:"netto_order_sales,omitempty"`
	OrderDateAndTime *dstore_values.TimestampValue `protobuf:"bytes,20001,opt,name=order_date_and_time,json=orderDateAndTime" json:"order_date_and_time,omitempty"`
	PersonId         *dstore_values.IntegerValue   `protobuf:"bytes,20002,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	NettoSum         *dstore_values.DecimalValue   `protobuf:"bytes,20003,opt,name=netto_sum,json=nettoSum" json:"netto_sum,omitempty"`
	OrderId          *dstore_values.IntegerValue   `protobuf:"bytes,20005,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetInformation() *dstore_values.StringValue {
	if m != nil {
		return m.Information
	}
	return nil
}

func (m *Response_Row) GetNettoOrderSales() *dstore_values.DecimalValue {
	if m != nil {
		return m.NettoOrderSales
	}
	return nil
}

func (m *Response_Row) GetOrderDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.OrderDateAndTime
	}
	return nil
}

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetNettoSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.NettoSum
	}
	return nil
}

func (m *Response_Row) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetSales_ByInformation_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetSales_ByInformation_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetSales_ByInformation_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetSales_ByInformation_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 789 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x4e, 0x13, 0x41,
	0x14, 0x0e, 0xd4, 0x5e, 0x18, 0x2e, 0x85, 0x6d, 0x24, 0x4b, 0x09, 0x04, 0x31, 0x46, 0xc3, 0x8f,
	0xad, 0xd1, 0xc4, 0x10, 0xbc, 0x24, 0x20, 0x46, 0xaa, 0xa1, 0x9a, 0x85, 0x90, 0x68, 0x42, 0x36,
	0x4b, 0x67, 0x5a, 0x36, 0x76, 0x77, 0x9a, 0x99, 0xa9, 0x84, 0xb7, 0x50, 0x7f, 0xfb, 0x47, 0x8d,
	0x89, 0x8f, 0xe2, 0x4b, 0xf8, 0x00, 0xde, 0x7d, 0x04, 0xcf, 0xcc, 0xd9, 0xc2, 0x6e, 0x29, 0x50,
	0xfe, 0xb4, 0x99, 0x3d, 0xdf, 0xe5, 0x9c, 0x99, 0x9d, 0x6f, 0xc9, 0x5d, 0x2a, 0x15, 0x17, 0xac,
	0xc2, 0xa2, 0x66, 0x10, 0xb1, 0x4a, 0x5b, 0xf0, 0x3a, 0xa3, 0x1d, 0xc1, 0x64, 0x45, 0x2a, 0xef,
	0x31, 0x53, 0x5b, 0x7e, 0x8b, 0x49, 0x6f, 0xed, 0xb0, 0x1a, 0x35, 0xb8, 0x08, 0x7d, 0x15, 0xf0,
	0xc8, 0x5b, 0xa5, 0x0e, 0xe0, 0x14, 0xb7, 0x96, 0x90, 0xec, 0x20, 0xd9, 0x39, 0x8b, 0x51, 0x2e,
	0xc5, 0x46, 0xaf, 0xfd, 0x56, 0x87, 0x49, 0x14, 0x28, 0xcf, 0xa4, 0xdd, 0x99, 0x10, 0x5c, 0xc4,
	0xa5, 0xd9, 0x74, 0x29, 0x64, 0x52, 0xfa, 0x4d, 0x16, 0x17, 0xaf, 0xf6, 0x16, 0x95, 0x1f, 0x1c,
	0xdb, 0x21, 0x68, 0xf1, 0x5b, 0x9e, 0x90, 0xe7, 0xbe, 0xf0, 0xa1, 0xca, 0x84, 0xb4, 0x56, 0xc8,
	0x48, 0x43, 0xf0, 0xd0, 0xa3, 0xbe, 0x62, 0xf6, 0xd0, 0xc2, 0xd0, 0x8d, 0xd1, 0x5b, 0x73, 0x4e,
	0x3c, 0x40, 0xdc, 0x94, 0x0a, 0xc0, 0x46, 0xf9, 0x61, 0x7b, 0x47, 0xaf, 0xdd, 0x82, 0xc6, 0xaf,
	0x03, 0xdc, 0xba, 0x46, 0x26, 0x8e, 0xb8, 0x5e, 0xd4, 0x69, 0xb5, 0xec, 0xef, 0x79, 0x50, 0x28,
	0xb8, 0x63, 0x5d, 0x48, 0x0d, 0x1e, 0x5a, 0x77, 0x48, 0x5e, 0x71, 0x34, 0x18, 0x1e, 0xc4, 0x20,
	0xa7, 0xb8, 0x91, 0xbf, 0x42, 0xc6, 0x62, 0x1e, 0x8a, 0xff, 0x40, 0x71, 0x82, 0x65, 0x23, 0xfd,
	0x94, 0x94, 0x12, 0x13, 0x7a, 0xea, 0xb0, 0xcd, 0xbc, 0x80, 0xda, 0x19, 0x63, 0x33, 0xdb, 0x63,
	0x13, 0x44, 0x8a, 0x35, 0x99, 0x40, 0x93, 0xa9, 0x04, 0x6f, 0x1b, 0x68, 0x55, 0x6a, 0x2d, 0x13,
	0xbb, 0x8f, 0x18, 0x7a, 0xff, 0x44, 0xef, 0xcb, 0x27, 0x58, 0xa6, 0x8d, 0x5d, 0x32, 0x97, 0x64,
	0x06, 0xd2, 0x03, 0x4b, 0xa0, 0x06, 0x91, 0xb7, 0xcf, 0x3b, 0x42, 0xda, 0x97, 0xce, 0x6f, 0x68,
	0x26, 0xa1, 0x50, 0x95, 0x3b, 0x9a, 0x5f, 0x8d, 0x36, 0x34, 0xdb, 0x7a, 0x42, 0x16, 0xcf, 0x94,
	0xc7, 0x16, 0x7f, 0x61, 0x8b, 0x73, 0xa7, 0xea, 0x98, 0x56, 0x77, 0xc8, 0x8c, 0x64, 0x2d, 0x56,
	0x57, 0x9e, 0xdf, 0x68, 0xc0, 0x1f, 0xa3, 0x1e, 0x17, 0x94, 0x09, 0x98, 0x54, 0xda, 0xd9, 0xbe,
	0x6d, 0xee, 0x71, 0xde, 0x62, 0x7e, 0x84, 0x6d, 0x4e, 0x23, 0x7b, 0x35, 0x26, 0x3f, 0xd3, 0xdc,
	0x2a, 0x95, 0xd6, 0x43, 0x32, 0x7f, 0xaa, 0x2e, 0xf6, 0xf7, 0x1b, 0xfb, 0x2b, 0xf7, 0x17, 0x30,
	0xcd, 0xad, 0x10, 0x82, 0x24, 0xca, 0x64, 0xdd, 0xce, 0x9d, 0xdf, 0xcd, 0x88, 0x81, 0xaf, 0x03,
	0xda, 0xba, 0x4e, 0x8a, 0xc7, 0x5c, 0x74, 0xfc, 0x83, 0x8e, 0xe3, 0x47, 0x20, 0x63, 0x72, 0x8f,
	0x8c, 0x26, 0xb6, 0xc8, 0xce, 0x1b, 0x97, 0x72, 0x8f, 0x8b, 0x54, 0x22, 0x88, 0x9a, 0x68, 0x92,
	0x84, 0x5b, 0x4b, 0x64, 0x32, 0x79, 0x16, 0xc6, 0xe7, 0x2f, 0xfa, 0x14, 0x13, 0x05, 0xe3, 0xb4,
	0x46, 0x26, 0xda, 0x1d, 0x51, 0xdf, 0xf7, 0x25, 0xc3, 0xcd, 0xb0, 0x0b, 0xe7, 0x8f, 0x34, 0xde,
	0xa5, 0x98, 0x9d, 0xb1, 0x2a, 0xa4, 0x94, 0xd6, 0x40, 0xcb, 0x7f, 0x68, 0x39, 0x95, 0x02, 0x6b,
	0xd3, 0xc5, 0xaf, 0x59, 0x52, 0x70, 0x99, 0x6c, 0xf3, 0x48, 0x32, 0xeb, 0x26, 0xc9, 0x9a, 0xf4,
	0x88, 0x6f, 0xf6, 0xd1, 0x94, 0x71, 0x34, 0x61, 0xb2, 0x3c, 0xd2, 0xbf, 0x2e, 0x02, 0xad, 0x17,
	0x64, 0x52, 0xe7, 0x86, 0x97, 0xdc, 0xa2, 0xe1, 0x85, 0x0c, 0x90, 0x9d, 0x1e, 0x72, 0x6f, 0xbc,
	0x6c, 0xc2, 0x3a, 0x91, 0x6e, 0x6e, 0x31, 0x4c, 0x3f, 0x80, 0xfb, 0x95, 0x8f, 0xf3, 0x0a, 0x2e,
	0xa8, 0x56, 0x9c, 0x3f, 0xa1, 0x88, 0x69, 0xb6, 0x89, 0xff, 0x6e, 0x17, 0x0e, 0x17, 0x20, 0x23,
	0xf8, 0x01, 0xdc, 0x22, 0xcd, 0x5a, 0x76, 0x06, 0xcf, 0x57, 0xa7, 0xbb, 0x13, 0x8e, 0xcb, 0x0f,
	0x5c, 0x2d, 0x52, 0xfe, 0x92, 0x21, 0x19, 0x58, 0x58, 0xd3, 0x24, 0x07, 0x4b, 0x9d, 0x16, 0x6f,
	0x6a, 0xb0, 0x39, 0x59, 0x37, 0x0b, 0x4b, 0x48, 0x81, 0xfb, 0xe9, 0xd7, 0xe3, 0x6d, 0xed, 0x62,
	0xef, 0xc7, 0x06, 0x99, 0x8a, 0x98, 0x82, 0xdc, 0xc2, 0xc3, 0x92, 0xba, 0x2d, 0xfb, 0x5d, 0xad,
	0xef, 0xb9, 0x53, 0x56, 0x0f, 0x42, 0xbf, 0x85, 0x2a, 0x45, 0x43, 0x33, 0xe7, 0x68, 0x66, 0xb1,
	0x6a, 0xa4, 0x14, 0xbf, 0xd0, 0x3a, 0x01, 0xfd, 0x88, 0x7a, 0x3a, 0x27, 0xed, 0x0f, 0xef, 0x07,
	0x4a, 0xe9, 0x49, 0x7c, 0xe9, 0x81, 0xba, 0x1a, 0xd1, 0x6d, 0xa8, 0xe9, 0xa4, 0x6f, 0x43, 0xe2,
	0xeb, 0x00, 0xa1, 0xf6, 0xc7, 0x58, 0xe5, 0xcc, 0x48, 0x2a, 0x20, 0x1e, 0x36, 0x05, 0xb8, 0x38,
	0x95, 0xec, 0x84, 0xf6, 0xa7, 0x53, 0xb8, 0xa9, 0x71, 0x0a, 0x06, 0xbf, 0xd5, 0x09, 0xe1, 0xd8,
	0x0b, 0xdd, 0x24, 0xb0, 0x3f, 0x0f, 0x62, 0x9b, 0xe7, 0x98, 0x09, 0x6b, 0xbb, 0x64, 0x36, 0xe0,
	0x3d, 0xa7, 0x7d, 0xfc, 0x29, 0x7e, 0xf9, 0xa0, 0xc9, 0x25, 0x7d, 0xd5, 0xad, 0xd3, 0x8b, 0x7e,
	0xad, 0xf7, 0x72, 0xe6, 0x83, 0x78, 0xfb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0xfe, 0x1b,
	0x87, 0xed, 0x07, 0x00, 0x00,
}