// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderStateHistory_Pu.proto
// DO NOT EDIT!

/*
Package om_GetOrderStateHistory_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderStateHistory_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderStateHistory_Pu

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
	InformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull bool                        `protobuf:"varint,1001,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
	Information           *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=information" json:"information,omitempty"`
	InformationNull       bool                        `protobuf:"varint,1002,opt,name=information_null,json=informationNull" json:"information_null,omitempty"`
	LanguageId            *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull        bool                        `protobuf:"varint,1003,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Parameters) GetInformation() *dstore_values.StringValue {
	if m != nil {
		return m.Information
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
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
	RowId                     int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ChangingDateAndTime       *dstore_values.TimestampValue `protobuf:"bytes,10001,opt,name=changing_date_and_time,json=changingDateAndTime" json:"changing_date_and_time,omitempty"`
	ToOrderStateId            *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=to_order_state_id,json=toOrderStateId" json:"to_order_state_id,omitempty"`
	UserName                  *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	ToOrderStateIdPublicDescr *dstore_values.StringValue    `protobuf:"bytes,10004,opt,name=to_order_state_id_public_descr,json=toOrderStateIdPublicDescr" json:"to_order_state_id_public_descr,omitempty"`
	OrderId                   *dstore_values.IntegerValue   `protobuf:"bytes,10005,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	ToOrderState              *dstore_values.StringValue    `protobuf:"bytes,10006,opt,name=to_order_state,json=toOrderState" json:"to_order_state,omitempty"`
	OrderContentId            *dstore_values.IntegerValue   `protobuf:"bytes,10007,opt,name=order_content_id,json=orderContentId" json:"order_content_id,omitempty"`
	FromOrderStatePublicDesc  *dstore_values.StringValue    `protobuf:"bytes,10008,opt,name=from_order_state_public_desc,json=fromOrderStatePublicDesc" json:"from_order_state_public_desc,omitempty"`
	FromOrderState            *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=from_order_state,json=fromOrderState" json:"from_order_state,omitempty"`
	UserId                    *dstore_values.IntegerValue   `protobuf:"bytes,10010,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	CompleteOrder             *dstore_values.BooleanValue   `protobuf:"bytes,10011,opt,name=complete_order,json=completeOrder" json:"complete_order,omitempty"`
	FromOrderStateId          *dstore_values.IntegerValue   `protobuf:"bytes,10012,opt,name=from_order_state_id,json=fromOrderStateId" json:"from_order_state_id,omitempty"`
	ChangingDateAndTimeChar   *dstore_values.StringValue    `protobuf:"bytes,10013,opt,name=changing_date_and_time_char,json=changingDateAndTimeChar" json:"changing_date_and_time_char,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetChangingDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.ChangingDateAndTime
	}
	return nil
}

func (m *Response_Row) GetToOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToOrderStateId
	}
	return nil
}

func (m *Response_Row) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Response_Row) GetToOrderStateIdPublicDescr() *dstore_values.StringValue {
	if m != nil {
		return m.ToOrderStateIdPublicDescr
	}
	return nil
}

func (m *Response_Row) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *Response_Row) GetToOrderState() *dstore_values.StringValue {
	if m != nil {
		return m.ToOrderState
	}
	return nil
}

func (m *Response_Row) GetOrderContentId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderContentId
	}
	return nil
}

func (m *Response_Row) GetFromOrderStatePublicDesc() *dstore_values.StringValue {
	if m != nil {
		return m.FromOrderStatePublicDesc
	}
	return nil
}

func (m *Response_Row) GetFromOrderState() *dstore_values.StringValue {
	if m != nil {
		return m.FromOrderState
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Response_Row) GetCompleteOrder() *dstore_values.BooleanValue {
	if m != nil {
		return m.CompleteOrder
	}
	return nil
}

func (m *Response_Row) GetFromOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromOrderStateId
	}
	return nil
}

func (m *Response_Row) GetChangingDateAndTimeChar() *dstore_values.StringValue {
	if m != nil {
		return m.ChangingDateAndTimeChar
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderStateHistory_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderStateHistory_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderStateHistory_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetOrderStateHistory_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x4f, 0xdb, 0x4a,
	0x10, 0x15, 0xe4, 0x42, 0xc2, 0x72, 0x6f, 0x08, 0x8b, 0x2e, 0x77, 0x49, 0x6e, 0x51, 0x45, 0x5f,
	0x68, 0x1f, 0x4c, 0xd5, 0xaa, 0x5f, 0x12, 0x2f, 0x14, 0x28, 0xa0, 0x8a, 0x14, 0x19, 0x54, 0x89,
	0x22, 0x75, 0x65, 0xe2, 0x25, 0xb5, 0x9a, 0xec, 0x46, 0xbb, 0xeb, 0x22, 0xfe, 0x45, 0xbf, 0xbf,
	0xfb, 0xe3, 0xfa, 0x52, 0xa9, 0xed, 0x8f, 0x68, 0x67, 0xbd, 0x76, 0x63, 0x3b, 0xa9, 0x9c, 0xbe,
	0x24, 0x5a, 0xcf, 0x9c, 0x39, 0xe7, 0xcc, 0x8c, 0xbd, 0xe8, 0x96, 0xaf, 0xb4, 0x90, 0x6c, 0x85,
	0xf1, 0x76, 0xc0, 0xd9, 0x4a, 0x4f, 0x8a, 0x16, 0xf3, 0x43, 0xc9, 0xd4, 0x8a, 0xe8, 0xd2, 0x2d,
	0xa6, 0xef, 0x49, 0x9f, 0xc9, 0x7d, 0xed, 0x69, 0xb6, 0x1d, 0x98, 0xcc, 0x33, 0xba, 0x17, 0x3a,
	0x90, 0xa5, 0x05, 0x5e, 0xb6, 0x50, 0xc7, 0x42, 0x9d, 0xdf, 0xe7, 0xd7, 0xe7, 0x62, 0x92, 0x27,
	0x5e, 0x27, 0x64, 0xca, 0xc2, 0xeb, 0x0b, 0x59, 0x66, 0x26, 0xa5, 0x90, 0x71, 0xa8, 0x91, 0x0d,
	0x75, 0x99, 0x52, 0x5e, 0x9b, 0xc5, 0xc1, 0x0b, 0xf9, 0xa0, 0xf6, 0x02, 0x7e, 0x22, 0x64, 0xd7,
	0xd3, 0x81, 0xe0, 0x36, 0x69, 0xe9, 0xcb, 0x38, 0x42, 0x7b, 0x9e, 0xf4, 0x20, 0xca, 0xa4, 0xc2,
	0x77, 0xd1, 0x5c, 0x2a, 0x87, 0xea, 0xb3, 0x1e, 0xa3, 0x81, 0x4f, 0xc6, 0xce, 0x8f, 0x2d, 0x4f,
	0x5f, 0x69, 0x38, 0xb1, 0x91, 0x58, 0x5e, 0xc0, 0x35, 0x6b, 0x33, 0x79, 0xdf, 0x9c, 0xdc, 0xd9,
	0x14, 0xee, 0x00, 0x60, 0x3b, 0x3e, 0xbe, 0x89, 0xc8, 0x90, 0x62, 0x94, 0x87, 0x9d, 0x0e, 0xf9,
	0x5a, 0x86, 0x92, 0x15, 0xf7, 0xdf, 0x01, 0x54, 0x13, 0xa2, 0x78, 0x15, 0x4d, 0xa7, 0x02, 0x64,
	0x3c, 0xa2, 0xaf, 0xe7, 0xe8, 0x95, 0x96, 0x01, 0x6f, 0x5b, 0xf6, 0x74, 0x3a, 0xbe, 0x84, 0x6a,
	0x69, 0xde, 0x88, 0xef, 0x9b, 0xe5, 0x9b, 0x49, 0x05, 0x12, 0xa6, 0x8e, 0xc7, 0xdb, 0x21, 0xb4,
	0xcd, 0x18, 0x2d, 0x15, 0x1b, 0x45, 0x49, 0x3e, 0x38, 0xbc, 0x88, 0x6a, 0x29, 0xb4, 0x65, 0xfa,
	0x6e, 0x99, 0xaa, 0xfd, 0x34, 0x43, 0xb4, 0xf4, 0x63, 0x0a, 0x55, 0x5c, 0xa6, 0x7a, 0x82, 0x2b,
	0x86, 0x2f, 0xa3, 0x89, 0x68, 0x8c, 0x71, 0x63, 0x7f, 0x39, 0x8b, 0x37, 0xc4, 0x8e, 0x78, 0xd3,
	0xfc, 0xba, 0x36, 0x11, 0x1f, 0xa2, 0x9a, 0x19, 0x20, 0xcd, 0xb6, 0xa5, 0x04, 0x60, 0x27, 0x07,
	0xce, 0xcf, 0x79, 0x17, 0xce, 0x3b, 0xfd, 0xb3, 0x3b, 0xd3, 0xcd, 0x3e, 0x80, 0x31, 0x95, 0xe3,
	0xc5, 0x01, 0xfb, 0xa6, 0xe2, 0xe2, 0x40, 0x45, 0xbb, 0x56, 0xbb, 0xf6, 0xdf, 0x4d, 0xd2, 0xf1,
	0x36, 0x2a, 0x49, 0x71, 0x4a, 0xfe, 0x8a, 0x50, 0xd7, 0x9d, 0x51, 0xd7, 0xdc, 0x49, 0xfa, 0xe0,
	0xb8, 0xe2, 0xd4, 0x35, 0x25, 0xea, 0x9f, 0xcb, 0xa8, 0x04, 0x07, 0x3c, 0x8f, 0x26, 0xe1, 0x68,
	0x26, 0xf1, 0xb4, 0x09, 0xad, 0x99, 0x70, 0x27, 0xe0, 0x08, 0x8d, 0xde, 0x47, 0xf3, 0xad, 0x47,
	0x9e, 0x29, 0xdc, 0xa6, 0x3e, 0x14, 0xa3, 0x1e, 0xf7, 0xa9, 0x0e, 0xba, 0x8c, 0x3c, 0x6b, 0x46,
	0x2d, 0x3c, 0x97, 0x1b, 0x99, 0x89, 0x29, 0xed, 0x75, 0x7b, 0x76, 0x68, 0x73, 0x09, 0x7a, 0x03,
	0xc0, 0x6b, 0xdc, 0x3f, 0x80, 0x30, 0xde, 0x42, 0xb3, 0x5a, 0x50, 0x61, 0x14, 0x52, 0x65, 0x24,
	0x1a, 0xde, 0xe7, 0xcd, 0xe2, 0x15, 0xa8, 0x6a, 0xd1, 0xf7, 0x15, 0x2d, 0xfa, 0x54, 0xa8, 0xa0,
	0x08, 0x87, 0xd7, 0x88, 0xbc, 0x68, 0x16, 0x6e, 0x6b, 0xc5, 0x64, 0x37, 0x21, 0x19, 0x3f, 0x44,
	0x8b, 0x03, 0x12, 0x68, 0x2f, 0x3c, 0xee, 0x04, 0x2d, 0xea, 0x33, 0xd5, 0x92, 0xe4, 0x65, 0x71,
	0xb9, 0x85, 0xac, 0x9c, 0xbd, 0x08, 0xbe, 0x61, 0xd0, 0xf8, 0x06, 0xaa, 0xd8, 0xe2, 0xe0, 0xec,
	0xd5, 0x08, 0xce, 0xca, 0x51, 0x36, 0x58, 0x5a, 0x43, 0xd5, 0xac, 0x30, 0xf2, 0xba, 0x58, 0xc8,
	0xdf, 0x69, 0x21, 0xf8, 0x0e, 0xaa, 0x59, 0x7c, 0x4b, 0x00, 0x05, 0xd7, 0x46, 0xc3, 0x9b, 0x51,
	0xba, 0x1b, 0xa1, 0xd6, 0x2d, 0x08, 0xa4, 0x1c, 0xa1, 0xff, 0x4f, 0x24, 0x6c, 0x53, 0xba, 0x4b,
	0xa9, 0x16, 0x91, 0xb7, 0xc5, 0xc2, 0x88, 0x29, 0xd0, 0x97, 0xd6, 0xef, 0x10, 0xde, 0x44, 0xb5,
	0x7c, 0x71, 0xf2, 0xae, 0xb8, 0x60, 0x35, 0x5b, 0x10, 0x5f, 0x43, 0xe5, 0x68, 0x03, 0xc0, 0xe2,
	0xfb, 0x11, 0x2c, 0x4e, 0x9a, 0x64, 0xb0, 0xb6, 0x8e, 0xaa, 0x2d, 0xd1, 0xed, 0x75, 0xe0, 0xe3,
	0x6b, 0x15, 0x90, 0x0f, 0xc3, 0xd1, 0xc7, 0x42, 0x74, 0x98, 0xc7, 0x2d, 0xfa, 0x9f, 0x04, 0x13,
	0x09, 0x30, 0xdf, 0xec, 0x81, 0xfe, 0x80, 0x8e, 0x8f, 0x23, 0xe8, 0xa8, 0x65, 0x6d, 0x80, 0xa2,
	0x43, 0xd4, 0x18, 0xfe, 0xa2, 0x51, 0x78, 0x2c, 0xc9, 0xa7, 0xe2, 0xd6, 0xfc, 0x37, 0xe4, 0x55,
	0x5b, 0x07, 0xec, 0xed, 0x23, 0xd4, 0x08, 0x44, 0xee, 0x23, 0xd1, 0xbf, 0x46, 0x1f, 0xac, 0xb6,
	0x85, 0xf2, 0x1f, 0x27, 0x71, 0xff, 0xcf, 0x6e, 0xda, 0xe3, 0xc9, 0xe8, 0x3a, 0xbb, 0xfa, 0x33,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0xff, 0x65, 0xad, 0xa7, 0x07, 0x00, 0x00,
}
