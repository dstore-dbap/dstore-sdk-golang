// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetPHStatistics_Ad.proto
// DO NOT EDIT!

/*
Package st_GetPHStatistics_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetPHStatistics_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetPHStatistics_Ad

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
	FromDateAndTime                *dstore_values.TimestampValue `protobuf:"bytes,1,opt,name=from_date_and_time,json=fromDateAndTime" json:"from_date_and_time,omitempty"`
	FromDateAndTimeNull            bool                          `protobuf:"varint,1001,opt,name=from_date_and_time_null,json=fromDateAndTimeNull" json:"from_date_and_time_null,omitempty"`
	ToDateAndTime                  *dstore_values.TimestampValue `protobuf:"bytes,2,opt,name=to_date_and_time,json=toDateAndTime" json:"to_date_and_time,omitempty"`
	ToDateAndTimeNull              bool                          `protobuf:"varint,1002,opt,name=to_date_and_time_null,json=toDateAndTimeNull" json:"to_date_and_time_null,omitempty"`
	BasicCharacteristicNumbers     *dstore_values.StringValue    `protobuf:"bytes,3,opt,name=basic_characteristic_numbers,json=basicCharacteristicNumbers" json:"basic_characteristic_numbers,omitempty"`
	BasicCharacteristicNumbersNull bool                          `protobuf:"varint,1003,opt,name=basic_characteristic_numbers_null,json=basicCharacteristicNumbersNull" json:"basic_characteristic_numbers_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetFromDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDateAndTime
	}
	return nil
}

func (m *Parameters) GetToDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDateAndTime
	}
	return nil
}

func (m *Parameters) GetBasicCharacteristicNumbers() *dstore_values.StringValue {
	if m != nil {
		return m.BasicCharacteristicNumbers
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
	Description               *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	BasicCharacteristic       *dstore_values.StringValue    `protobuf:"bytes,10002,opt,name=basic_characteristic,json=basicCharacteristic" json:"basic_characteristic,omitempty"`
	Value                     *dstore_values.DecimalValue   `protobuf:"bytes,10003,opt,name=value" json:"value,omitempty"`
	DateAndTime               *dstore_values.TimestampValue `protobuf:"bytes,10004,opt,name=date_and_time,json=dateAndTime" json:"date_and_time,omitempty"`
	BasicCharacteristicNumber *dstore_values.IntegerValue   `protobuf:"bytes,10005,opt,name=basic_characteristic_number,json=basicCharacteristicNumber" json:"basic_characteristic_number,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetBasicCharacteristic() *dstore_values.StringValue {
	if m != nil {
		return m.BasicCharacteristic
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.DateAndTime
	}
	return nil
}

func (m *Response_Row) GetBasicCharacteristicNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicCharacteristicNumber
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetPHStatistics_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetPHStatistics_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetPHStatistics_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetPHStatistics_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0x0e, 0xd6, 0x05, 0x72, 0x08, 0x01, 0x67, 0xfd, 0x29, 0x5d, 0x25, 0x0a, 0x37, 0x5e, 0x75,
	0x05, 0x35, 0x7a, 0xe3, 0x05, 0xe0, 0x1f, 0x26, 0x34, 0xa4, 0x1a, 0x12, 0x0d, 0x49, 0x33, 0xdb,
	0x19, 0xea, 0xc4, 0xed, 0xcc, 0x66, 0x66, 0x56, 0xee, 0x7d, 0x02, 0x7f, 0xaf, 0x7c, 0x0d, 0x5f,
	0x4a, 0x7d, 0x09, 0x4f, 0x3b, 0x5d, 0xd9, 0x96, 0x05, 0xf6, 0x66, 0x37, 0xa7, 0xe7, 0x7c, 0xdf,
	0xf9, 0xe6, 0x9c, 0x6f, 0x06, 0x1e, 0x30, 0x63, 0x95, 0xe6, 0x5d, 0x2e, 0x33, 0x21, 0x79, 0x77,
	0xa0, 0x55, 0xca, 0xd9, 0x50, 0x73, 0xd3, 0x35, 0x36, 0x79, 0xc1, 0xed, 0xfe, 0xcb, 0xd7, 0x96,
	0x5a, 0x61, 0xac, 0x48, 0x4d, 0xb2, 0xc5, 0x42, 0x2c, 0xb0, 0x8a, 0xac, 0x3b, 0x54, 0xe8, 0x50,
	0xe1, 0xc4, 0xd2, 0xa0, 0x5d, 0x51, 0x7f, 0xa4, 0xfd, 0x21, 0x37, 0x0e, 0x19, 0xac, 0xd4, 0xfb,
	0x71, 0xad, 0x95, 0xae, 0x52, 0x9d, 0x7a, 0x2a, 0xe7, 0xc6, 0xd0, 0x8c, 0x57, 0xc9, 0xf5, 0x66,
	0xd2, 0x52, 0x21, 0x8f, 0x94, 0xce, 0xb1, 0xa3, 0x92, 0xae, 0x68, 0xed, 0x97, 0x07, 0xb0, 0x4f,
	0x35, 0xc5, 0x2c, 0xd7, 0x86, 0xbc, 0x02, 0x72, 0xa4, 0x55, 0x9e, 0x30, 0x6a, 0x79, 0x42, 0x25,
	0x4b, 0xac, 0xc8, 0xb9, 0x3f, 0x73, 0x7b, 0xe6, 0xee, 0xc2, 0xe6, 0xad, 0xb0, 0x3a, 0x42, 0xa5,
	0xae, 0x48, 0x19, 0x4b, 0xf3, 0xc1, 0x41, 0x11, 0xc7, 0x4b, 0x05, 0xf0, 0x29, 0xe2, 0xb6, 0x24,
	0x7b, 0x83, 0x29, 0xf2, 0x10, 0x6e, 0x9c, 0xe6, 0x4a, 0xe4, 0xb0, 0xdf, 0xf7, 0x7f, 0xcf, 0x21,
	0xe3, 0x7c, 0xdc, 0x6e, 0x40, 0x22, 0xcc, 0x91, 0xe7, 0xb0, 0x6c, 0x55, 0x43, 0xc0, 0xa5, 0x69,
	0x04, 0x2c, 0x5a, 0x35, 0xde, 0x7e, 0x03, 0xae, 0x35, 0x79, 0x5c, 0xf3, 0x3f, 0xae, 0xf9, 0x95,
	0x5a, 0x79, 0xd9, 0xfa, 0x10, 0x6e, 0xf6, 0xa8, 0x11, 0x69, 0x92, 0xbe, 0xc7, 0x91, 0xa4, 0x38,
	0x91, 0x72, 0x35, 0x08, 0xcb, 0x7b, 0x38, 0x1d, 0xdf, 0x2b, 0x65, 0x04, 0x0d, 0x19, 0xc6, 0x6a,
	0x21, 0x33, 0xa7, 0x21, 0x28, 0xf1, 0x3b, 0x35, 0x78, 0xe4, 0xd0, 0x38, 0xdb, 0x3b, 0xe7, 0xb1,
	0x3b, 0x71, 0x7f, 0x9d, 0xb8, 0xd5, 0xb3, 0x79, 0x0a, 0xa5, 0x6b, 0x3f, 0x5b, 0x30, 0x1f, 0x73,
	0x33, 0x50, 0xd2, 0x70, 0x72, 0x0f, 0x5a, 0xa5, 0x29, 0xaa, 0x3d, 0xfd, 0xd7, 0x57, 0x59, 0xcd,
	0x19, 0xe6, 0x59, 0xf1, 0x1b, 0xbb, 0x42, 0xf2, 0x16, 0x96, 0x0b, 0x3b, 0x24, 0x63, 0x7e, 0xc0,
	0x19, 0x7b, 0x08, 0x0e, 0x1b, 0xe0, 0xa6, 0x6b, 0xf6, 0x30, 0xde, 0x3d, 0x89, 0xe3, 0xa5, 0xbc,
	0xfe, 0x81, 0x3c, 0x86, 0xb9, 0xca, 0x86, 0x38, 0xae, 0x82, 0x71, 0xf5, 0x14, 0xa3, 0x33, 0xe9,
	0x9e, 0xfb, 0x8f, 0x47, 0xe5, 0x64, 0x07, 0x3c, 0xad, 0x8e, 0xfd, 0xcb, 0x25, 0x6a, 0x23, 0x9c,
	0xe2, 0xbe, 0x84, 0xa3, 0x11, 0x84, 0xb1, 0x3a, 0x8e, 0x0b, 0x74, 0xf0, 0xc9, 0x03, 0x0f, 0x03,
	0x72, 0x1d, 0x66, 0x31, 0x4c, 0x04, 0xf3, 0x3f, 0x47, 0x38, 0x95, 0x56, 0xdc, 0xc2, 0x70, 0x97,
	0x91, 0x27, 0xb0, 0xc0, 0xb8, 0x49, 0xb5, 0x18, 0x94, 0x87, 0xfe, 0x12, 0x5d, 0xb8, 0xd2, 0xf1,
	0x7a, 0x12, 0xc1, 0xd5, 0x49, 0x3b, 0xf4, 0xbf, 0x5e, 0xcc, 0xd3, 0x9e, 0xb0, 0x52, 0xb2, 0x09,
	0xad, 0xb2, 0xd4, 0xff, 0xe6, 0x08, 0x3a, 0x0d, 0x02, 0xc6, 0x53, 0x91, 0xd3, 0xbe, 0x63, 0x70,
	0xa5, 0x64, 0x1b, 0x16, 0xeb, 0xb7, 0xe3, 0x7b, 0x34, 0xcd, 0xf5, 0x58, 0x60, 0x63, 0x97, 0xe3,
	0x10, 0x3a, 0xe7, 0x78, 0xd1, 0xff, 0x31, 0x59, 0x8d, 0x90, 0x96, 0x67, 0x5c, 0x3b, 0xbe, 0x95,
	0x33, 0x2d, 0xba, 0x7d, 0x00, 0x1d, 0xa1, 0x1a, 0x0b, 0x3c, 0x79, 0x26, 0xdf, 0x3d, 0xca, 0x94,
	0x61, 0x1f, 0x46, 0x79, 0x36, 0xf5, 0x4b, 0xda, 0x9b, 0x2d, 0xdf, 0xac, 0xfb, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xcd, 0xf4, 0x7f, 0xa2, 0x82, 0x05, 0x00, 0x00,
}