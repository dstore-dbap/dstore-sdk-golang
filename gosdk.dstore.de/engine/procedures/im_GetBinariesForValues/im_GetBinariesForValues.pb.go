// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetBinariesForValues.proto
// DO NOT EDIT!

/*
Package im_GetBinariesForValues is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetBinariesForValues.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetBinariesForValues

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
	ValueId                         *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	ValueIdNull                     bool                        `protobuf:"varint,1001,opt,name=value_id_null,json=valueIdNull" json:"value_id_null,omitempty"`
	NodeCharacteristicId            *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull        bool                        `protobuf:"varint,1002,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	Value                           *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ValueNull                       bool                        `protobuf:"varint,1003,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
	ValuesInAnyValues               *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=values_in_any_values,json=valuesInAnyValues" json:"values_in_any_values,omitempty"`
	ValuesInAnyValuesNull           bool                        `protobuf:"varint,1004,opt,name=values_in_any_values_null,json=valuesInAnyValuesNull" json:"values_in_any_values_null,omitempty"`
	IncludeBinaryCode               *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=include_binary_code,json=includeBinaryCode" json:"include_binary_code,omitempty"`
	IncludeBinaryCodeNull           bool                        `protobuf:"varint,1005,opt,name=include_binary_code_null,json=includeBinaryCodeNull" json:"include_binary_code_null,omitempty"`
	FilterByBinaryCharacValueId     *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=filter_by_binary_charac_value_id,json=filterByBinaryCharacValueId" json:"filter_by_binary_charac_value_id,omitempty"`
	FilterByBinaryCharacValueIdNull bool                        `protobuf:"varint,1006,opt,name=filter_by_binary_charac_value_id_null,json=filterByBinaryCharacValueIdNull" json:"filter_by_binary_charac_value_id_null,omitempty"`
	BinaryCharacteristicId          *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=binary_characteristic_id,json=binaryCharacteristicId" json:"binary_characteristic_id,omitempty"`
	BinaryCharacteristicIdNull      bool                        `protobuf:"varint,1007,opt,name=binary_characteristic_id_null,json=binaryCharacteristicIdNull" json:"binary_characteristic_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Parameters) GetValuesInAnyValues() *dstore_values.BooleanValue {
	if m != nil {
		return m.ValuesInAnyValues
	}
	return nil
}

func (m *Parameters) GetIncludeBinaryCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.IncludeBinaryCode
	}
	return nil
}

func (m *Parameters) GetFilterByBinaryCharacValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilterByBinaryCharacValueId
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicId
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
	RowId                int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	BinaryValueId        *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=binary_value_id,json=binaryValueId" json:"binary_value_id,omitempty"`
	BinaryCharacValue    *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=binary_charac_value,json=binaryCharacValue" json:"binary_charac_value,omitempty"`
	BinaryCodeId         *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	NodeCharacteristicId *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	Value                *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=value" json:"value,omitempty"`
	ValueId              *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	ThumbnailCode        *dstore_values.BytesValue   `protobuf:"bytes,10007,opt,name=thumbnail_code,json=thumbnailCode" json:"thumbnail_code,omitempty"`
	BinaryCode           *dstore_values.BytesValue   `protobuf:"bytes,10008,opt,name=binary_code,json=binaryCode" json:"binary_code,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetBinaryValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryValueId
	}
	return nil
}

func (m *Response_Row) GetBinaryCharacValue() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryCharacValue
	}
	return nil
}

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Response_Row) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func (m *Response_Row) GetThumbnailCode() *dstore_values.BytesValue {
	if m != nil {
		return m.ThumbnailCode
	}
	return nil
}

func (m *Response_Row) GetBinaryCode() *dstore_values.BytesValue {
	if m != nil {
		return m.BinaryCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetBinariesForValues.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetBinariesForValues.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetBinariesForValues.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetBinariesForValues.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x4e, 0xd4, 0x4a,
	0x18, 0x0e, 0x67, 0xd9, 0x5d, 0xce, 0xcb, 0x01, 0x0e, 0x85, 0x43, 0xca, 0x6e, 0x0e, 0x12, 0x08,
	0x89, 0xbf, 0x8a, 0x4a, 0x82, 0x12, 0x13, 0x13, 0x16, 0xd1, 0x10, 0xa5, 0x6a, 0x13, 0x89, 0xfa,
	0xa7, 0x69, 0xb7, 0xc3, 0x32, 0xb1, 0x3b, 0x43, 0x3a, 0x5d, 0x49, 0xef, 0xc2, 0x6f, 0xbd, 0x0a,
	0x7f, 0x79, 0x1f, 0x5e, 0x87, 0xdf, 0xb7, 0xe0, 0x74, 0xde, 0xd9, 0xa5, 0x5b, 0x8a, 0x05, 0xff,
	0xec, 0x66, 0x3e, 0x9e, 0xf7, 0x79, 0xfa, 0xce, 0x33, 0xcf, 0xc0, 0x7a, 0x20, 0x62, 0x1e, 0x91,
	0x55, 0xc2, 0x3a, 0x94, 0x91, 0xd5, 0xc3, 0x88, 0xb7, 0x49, 0xd0, 0x8b, 0x88, 0x58, 0xa5, 0x5d,
	0xf7, 0x36, 0x89, 0x5b, 0x94, 0x79, 0x11, 0x25, 0xe2, 0x16, 0x8f, 0xf6, 0xbc, 0xb0, 0x47, 0x84,
	0x25, 0xb7, 0xc4, 0xdc, 0x58, 0x41, 0x9c, 0x85, 0x38, 0xeb, 0x94, 0xcd, 0x8d, 0x19, 0x5d, 0xfe,
	0x59, 0x06, 0xdb, 0x98, 0x1f, 0xe6, 0x24, 0x51, 0xc4, 0x23, 0xbd, 0xd4, 0x1c, 0x5e, 0xea, 0x12,
	0x21, 0xbc, 0x0e, 0xd1, 0x8b, 0xcb, 0xf9, 0xc5, 0xd8, 0xa3, 0x6c, 0x9f, 0x47, 0x5d, 0x2f, 0xa6,
	0x9c, 0xe1, 0xa6, 0xa5, 0x0f, 0x75, 0x80, 0xfb, 0x5e, 0xe4, 0xc9, 0x55, 0x12, 0x09, 0x63, 0x1d,
	0xc6, 0x14, 0xb7, 0x4b, 0x03, 0x73, 0x64, 0x71, 0xe4, 0xe2, 0xf8, 0x95, 0xa6, 0xa5, 0xa5, 0x6b,
	0x4d, 0x94, 0xc5, 0xa4, 0x43, 0x50, 0xb0, 0x53, 0x57, 0x93, 0x3b, 0x81, 0xb1, 0x0c, 0x13, 0x7d,
	0x9c, 0xcb, 0x7a, 0x61, 0x68, 0x7e, 0xae, 0x4b, 0xf4, 0x98, 0x33, 0xae, 0x37, 0xd8, 0x72, 0xce,
	0x78, 0x00, 0x73, 0x8c, 0x07, 0xc4, 0x6d, 0x1f, 0x48, 0xc2, 0xb6, 0xe4, 0xa3, 0x22, 0xa6, 0xed,
	0x94, 0xea, 0xaf, 0x72, 0xaa, 0xd9, 0x14, 0xba, 0x35, 0x84, 0x94, 0xbc, 0x37, 0xa0, 0x59, 0x5c,
	0x12, 0x55, 0x7c, 0x41, 0x15, 0x66, 0x11, 0x56, 0x49, 0xba, 0x04, 0x55, 0x45, 0x66, 0x56, 0x94,
	0x82, 0x46, 0x4e, 0x81, 0x88, 0x23, 0xca, 0x3a, 0x28, 0x00, 0x37, 0x1a, 0x0b, 0x00, 0xf8, 0xa5,
	0x8a, 0xe0, 0x2b, 0x12, 0xfc, 0xad, 0xa6, 0x54, 0xc5, 0xbb, 0x30, 0x8b, 0x60, 0x97, 0x32, 0xd7,
	0x63, 0x89, 0x8b, 0x23, 0x73, 0xb4, 0xf0, 0x13, 0x7d, 0xce, 0x43, 0xe2, 0x31, 0x64, 0x98, 0xc6,
	0xc9, 0x1d, 0xb6, 0xc9, 0x12, 0x34, 0x84, 0xb1, 0x01, 0xf3, 0x45, 0xd5, 0x90, 0xfc, 0x1b, 0x92,
	0xff, 0x77, 0x02, 0xa6, 0x84, 0xdc, 0x81, 0x19, 0xca, 0xda, 0x61, 0x4f, 0x76, 0xc7, 0x4f, 0x8d,
	0x96, 0xb8, 0x6d, 0xd9, 0x05, 0xb3, 0x5a, 0xde, 0xea, 0x69, 0x8d, 0x53, 0xfe, 0x4c, 0xb6, 0x24,
	0xca, 0xb8, 0x06, 0x66, 0x41, 0x31, 0x94, 0xf1, 0x5d, 0xcb, 0x38, 0x81, 0x52, 0x32, 0x7c, 0x58,
	0xdc, 0xa7, 0xa1, 0x6c, 0xbb, 0xeb, 0x27, 0x03, 0xac, 0x3a, 0x0a, 0x77, 0xe0, 0xb4, 0x5a, 0xb9,
	0xa6, 0x26, 0x16, 0x69, 0x25, 0xba, 0xbc, 0xaa, 0xb0, 0xa7, 0xdd, 0x77, 0x0f, 0x56, 0xca, 0x38,
	0x50, 0xea, 0x0f, 0x94, 0x7a, 0xe1, 0x37, 0xc5, 0x94, 0xe8, 0x87, 0x60, 0x0e, 0x95, 0xc9, 0x7a,
	0xb5, 0x5e, 0x2e, 0x76, 0xce, 0xcf, 0xd4, 0xcd, 0xb8, 0xb5, 0x05, 0xff, 0x9f, 0x56, 0x16, 0xf5,
	0xfd, 0x44, 0x7d, 0x8d, 0x62, 0x7c, 0x2a, 0x6d, 0xe9, 0x53, 0x0d, 0xc6, 0x1c, 0x22, 0x0e, 0x39,
	0x13, 0x24, 0xb5, 0xaf, 0x8a, 0x03, 0x7d, 0x57, 0x07, 0xf6, 0xd5, 0x31, 0x83, 0x51, 0xb1, 0x9d,
	0xfe, 0x3a, 0xb8, 0xd1, 0x78, 0x0c, 0xff, 0xa6, 0x41, 0xe0, 0x66, 0x92, 0x40, 0xde, 0xbe, 0x8a,
	0x04, 0x5b, 0x39, 0x70, 0x3e, 0x2f, 0x76, 0xe5, 0x78, 0xe7, 0x78, 0xec, 0x4c, 0x75, 0x87, 0x27,
	0xa4, 0x47, 0xea, 0x3a, 0x80, 0xe4, 0x6d, 0x4a, 0x2b, 0x2e, 0x9c, 0xa8, 0x88, 0xf1, 0xb4, 0x8b,
	0xff, 0x4e, 0x7f, 0xbb, 0xb1, 0x0d, 0x95, 0x88, 0x1f, 0xc9, 0x2b, 0x92, 0xa2, 0xd6, 0xac, 0x33,
	0x65, 0xa5, 0xd5, 0x6f, 0x82, 0xe5, 0xf0, 0x23, 0x27, 0xc5, 0x37, 0x3e, 0x8e, 0x42, 0x45, 0x0e,
	0x8c, 0x39, 0xa8, 0xc9, 0x61, 0x7a, 0x56, 0xcf, 0x6d, 0xd9, 0x97, 0xaa, 0x53, 0x95, 0x43, 0xd9,
	0xfe, 0x9b, 0x30, 0xa5, 0xdb, 0x3f, 0x70, 0xde, 0x0b, 0xbb, 0xfc, 0x34, 0x27, 0x10, 0xd4, 0x37,
	0x9b, 0xbc, 0x57, 0x05, 0x16, 0x33, 0x5f, 0xda, 0xa5, 0x09, 0x32, 0xed, 0xe7, 0xed, 0x26, 0x1d,
	0x31, 0x99, 0xbd, 0x4f, 0x52, 0xd1, 0xab, 0x33, 0x28, 0xfa, 0xc7, 0x1f, 0xdc, 0x31, 0x29, 0xc8,
	0x39, 0x35, 0x56, 0x5f, 0xdb, 0x7f, 0x9a, 0xab, 0x97, 0xfb, 0xb9, 0xf8, 0xc6, 0x3e, 0x6b, 0x30,
	0x5e, 0xcd, 0x3c, 0x1d, 0x6f, 0xed, 0x73, 0xbc, 0x1d, 0x9b, 0x30, 0x19, 0x1f, 0xf4, 0xba, 0x3e,
	0xf3, 0x68, 0x88, 0x19, 0xf5, 0x0e, 0xe1, 0xf3, 0xf9, 0xb0, 0x4c, 0x62, 0x22, 0xf4, 0x99, 0x0c,
	0x10, 0x2a, 0x9e, 0xae, 0xc3, 0x78, 0x36, 0xe3, 0xde, 0x97, 0xe2, 0xe1, 0xb8, 0x83, 0xad, 0x47,
	0xd0, 0xa4, 0x3c, 0x67, 0xba, 0xe3, 0x87, 0xfd, 0xc9, 0x46, 0x87, 0x8b, 0xe0, 0x69, 0x7f, 0x3d,
	0x38, 0xc7, 0xdb, 0xef, 0xd7, 0xd4, 0x1b, 0xbb, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x37,
	0xdf, 0x43, 0x36, 0x08, 0x00, 0x00,
}