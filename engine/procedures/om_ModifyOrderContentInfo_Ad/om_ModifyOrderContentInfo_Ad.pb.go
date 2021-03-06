// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyOrderContentInfo_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyOrderContentInfo_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyOrderContentInfo_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyOrderContentInfo_Ad

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
	ChangeAllOrNothing     *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=change_all_or_nothing,json=changeAllOrNothing" json:"change_all_or_nothing,omitempty"`
	ChangeAllOrNothingNull bool                        `protobuf:"varint,1001,opt,name=change_all_or_nothing_null,json=changeAllOrNothingNull" json:"change_all_or_nothing_null,omitempty"`
	OnlyNewData            *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=only_new_data,json=onlyNewData" json:"only_new_data,omitempty"`
	OnlyNewDataNull        bool                        `protobuf:"varint,1002,opt,name=only_new_data_null,json=onlyNewDataNull" json:"only_new_data_null,omitempty"`
	Country                *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=country" json:"country,omitempty"`
	CountryNull            bool                        `protobuf:"varint,1003,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetChangeAllOrNothing() *dstore_values.BooleanValue {
	if m != nil {
		return m.ChangeAllOrNothing
	}
	return nil
}

func (m *Parameters) GetOnlyNewData() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyNewData
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	InformationTypeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	ErrorCode         *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	OrderContentId    *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=order_content_id,json=orderContentId" json:"order_content_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Response_Row) GetErrorCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Response_Row) GetOrderContentId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderContentId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyOrderContentInfo_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyOrderContentInfo_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyOrderContentInfo_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyOrderContentInfo_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6b, 0x13, 0x41,
	0x10, 0x27, 0x4d, 0xd3, 0xd4, 0x89, 0xda, 0xba, 0xc5, 0x72, 0x26, 0x20, 0x12, 0x5f, 0x44, 0xe4,
	0x22, 0xea, 0x43, 0x21, 0xa0, 0xd4, 0xaa, 0x10, 0x31, 0x57, 0x39, 0x44, 0x50, 0x90, 0x65, 0x9b,
	0x9d, 0x5e, 0x0f, 0x2f, 0x3b, 0x61, 0x6f, 0x63, 0xc8, 0xb7, 0xa8, 0xfa, 0x0d, 0x7d, 0x53, 0x5f,
	0xfc, 0x08, 0xee, 0xdd, 0x6e, 0x6c, 0x92, 0x16, 0x8d, 0x2f, 0x49, 0xe6, 0xe6, 0xf7, 0x67, 0xfe,
	0xe5, 0xa0, 0x2b, 0x73, 0x43, 0x1a, 0x3b, 0xa8, 0x92, 0x54, 0x61, 0x67, 0xa4, 0x69, 0x80, 0x72,
	0xac, 0x31, 0xef, 0xd0, 0x90, 0xf7, 0x49, 0xa6, 0xc7, 0xd3, 0x43, 0x2d, 0x51, 0x1f, 0x90, 0x32,
	0xa8, 0x4c, 0x4f, 0x1d, 0x13, 0xdf, 0x97, 0xa1, 0xc5, 0x19, 0x62, 0x77, 0x1d, 0x39, 0x74, 0xe4,
	0xf0, 0x6f, 0x8c, 0xe6, 0x8e, 0x37, 0xfa, 0x24, 0xb2, 0x31, 0xe6, 0x4e, 0xa0, 0x79, 0x63, 0xd1,
	0x1d, 0xb5, 0x26, 0xed, 0x53, 0xad, 0xc5, 0xd4, 0x10, 0xf3, 0x5c, 0x24, 0xe8, 0x93, 0xb7, 0x97,
	0x93, 0x46, 0xa4, 0xd6, 0x47, 0x0f, 0x85, 0x49, 0x49, 0x39, 0x50, 0xfb, 0xdb, 0x1a, 0xc0, 0x6b,
	0xa1, 0x85, 0xcd, 0xa2, 0xce, 0x59, 0x04, 0xd7, 0x07, 0x27, 0x42, 0x25, 0xc8, 0x45, 0x96, 0x71,
	0xd2, 0x5c, 0x91, 0x39, 0x49, 0x55, 0x12, 0x54, 0x6e, 0x55, 0xee, 0x34, 0x1e, 0xb4, 0x42, 0xdf,
	0x8c, 0x2f, 0xf0, 0x88, 0x28, 0x43, 0xa1, 0xde, 0x16, 0x51, 0xcc, 0x1c, 0x73, 0x3f, 0xcb, 0x0e,
	0x75, 0xe4, 0x68, 0xac, 0x0b, 0xcd, 0x0b, 0xf5, 0xb8, 0x1a, 0x67, 0x59, 0xf0, 0xbd, 0x6e, 0x55,
	0x37, 0xe3, 0xdd, 0xf3, 0xc4, 0xc8, 0xa6, 0xd9, 0x13, 0xb8, 0x42, 0x2a, 0x9b, 0x72, 0x85, 0x13,
	0x2e, 0x85, 0x11, 0xc1, 0xda, 0xbf, 0x8b, 0x68, 0x14, 0x8c, 0x08, 0x27, 0xcf, 0x2c, 0x9e, 0xdd,
	0x03, 0xb6, 0x20, 0xe0, 0x5c, 0x7f, 0x38, 0xd7, 0xad, 0x39, 0x64, 0x69, 0xf7, 0x08, 0xea, 0x03,
	0x1a, 0x2b, 0xa3, 0xa7, 0x41, 0xb5, 0x34, 0x6a, 0x2e, 0x19, 0xe5, 0x46, 0xdb, 0xd2, 0x9c, 0xcf,
	0x0c, 0xca, 0xda, 0x70, 0xd9, 0xff, 0x74, 0xea, 0x3f, 0x9d, 0x7a, 0xc3, 0x3f, 0x2c, 0x94, 0xdb,
	0xa7, 0xeb, 0xb0, 0x19, 0x63, 0x3e, 0x22, 0x95, 0x23, 0xbb, 0x0f, 0xb5, 0x72, 0x85, 0x7e, 0xa4,
	0x7f, 0x4c, 0xfc, 0x7d, 0xb8, 0xf5, 0x3e, 0x2f, 0x3e, 0x63, 0x07, 0x64, 0xef, 0x60, 0xbb, 0x58,
	0x1e, 0x9f, 0xdb, 0x9e, 0x1d, 0x45, 0xd5, 0x92, 0xc3, 0x25, 0xf2, 0xf2, 0x8e, 0xfb, 0x36, 0xee,
	0x9d, 0xc5, 0xf1, 0xd6, 0x70, 0xf1, 0x01, 0xdb, 0x83, 0xba, 0x3f, 0x1a, 0xdb, 0x73, 0xa1, 0x78,
	0xf3, 0x9c, 0xa2, 0x3b, 0xa9, 0xbe, 0xfb, 0x8e, 0x67, 0x70, 0xf6, 0x12, 0xaa, 0x9a, 0x26, 0xc1,
	0x7a, 0xc9, 0xda, 0x0b, 0x57, 0x3f, 0xf2, 0x70, 0x36, 0x89, 0x30, 0xa6, 0x49, 0x5c, 0x88, 0x34,
	0x7f, 0x55, 0xa0, 0x6a, 0x03, 0xb6, 0x0b, 0x1b, 0x36, 0xe4, 0xa9, 0x0c, 0x4e, 0x23, 0x3b, 0x9c,
	0x5a, 0x5c, 0xb3, 0x61, 0x4f, 0xb2, 0x57, 0xb0, 0x33, 0xd7, 0x15, 0x37, 0xd3, 0x11, 0x16, 0xa0,
	0xcf, 0xd1, 0x85, 0xf7, 0x90, 0x5a, 0xb7, 0x04, 0xb5, 0xdb, 0xd3, 0xb5, 0x39, 0xe2, 0x1b, 0xcb,
	0xb3, 0x6a, 0x5d, 0x80, 0x72, 0xae, 0x7c, 0x40, 0x12, 0x83, 0x2f, 0x2b, 0x88, 0x5c, 0x2a, 0xf1,
	0x07, 0x16, 0xce, 0x5e, 0xc0, 0x36, 0x15, 0x3d, 0x59, 0x72, 0xd9, 0x54, 0x51, 0xc7, 0xd7, 0x15,
	0x24, 0xae, 0xd2, 0xfc, 0x24, 0xe4, 0xd3, 0x0f, 0xd0, 0x4a, 0x69, 0x69, 0x6a, 0x67, 0xef, 0x95,
	0xf7, 0x8f, 0x13, 0xca, 0xe5, 0xc7, 0x59, 0x5e, 0xfe, 0xef, 0xab, 0xe7, 0x68, 0xa3, 0xfc, 0x77,
	0x3f, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x50, 0x83, 0x07, 0x72, 0xba, 0x04, 0x00, 0x00,
}
