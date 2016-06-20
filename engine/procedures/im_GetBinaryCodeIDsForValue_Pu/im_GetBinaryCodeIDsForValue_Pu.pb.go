// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetBinaryCodeIDsForValue_Pu.proto
// DO NOT EDIT!

/*
Package im_GetBinaryCodeIDsForValue_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetBinaryCodeIDsForValue_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetBinaryCodeIDsForValue_Pu

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
	ValueId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	ValueIdNull               bool                        `protobuf:"varint,1001,opt,name=value_id_null,json=valueIdNull" json:"value_id_null,omitempty"`
	BinaryPropertyValueId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=binary_property_value_id,json=binaryPropertyValueId" json:"binary_property_value_id,omitempty"`
	BinaryPropertyValueIdNull bool                        `protobuf:"varint,1002,opt,name=binary_property_value_id_null,json=binaryPropertyValueIdNull" json:"binary_property_value_id_null,omitempty"`
	Value                     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ValueNull                 bool                        `protobuf:"varint,1003,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
	NodeCharacteristicId      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull  bool                        `protobuf:"varint,1004,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	ValueIdsInOneId           *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=value_ids_in_one_id,json=valueIdsInOneId" json:"value_ids_in_one_id,omitempty"`
	ValueIdsInOneIdNull       bool                        `protobuf:"varint,1005,opt,name=value_ids_in_one_id_null,json=valueIdsInOneIdNull" json:"value_ids_in_one_id_null,omitempty"`
	ValuesInAnyValues         *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=values_in_any_values,json=valuesInAnyValues" json:"values_in_any_values,omitempty"`
	ValuesInAnyValuesNull     bool                        `protobuf:"varint,1006,opt,name=values_in_any_values_null,json=valuesInAnyValuesNull" json:"values_in_any_values_null,omitempty"`
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

func (m *Parameters) GetBinaryPropertyValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryPropertyValueId
	}
	return nil
}

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetValueIdsInOneId() *dstore_values.BooleanValue {
	if m != nil {
		return m.ValueIdsInOneId
	}
	return nil
}

func (m *Parameters) GetValuesInAnyValues() *dstore_values.BooleanValue {
	if m != nil {
		return m.ValuesInAnyValues
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
	BinaryCodeId         *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	ValueId              *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	NodeCharacteristicId *dstore_values.IntegerValue `protobuf:"bytes,30002,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	Value                *dstore_values.StringValue  `protobuf:"bytes,30003,opt,name=value" json:"value,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
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

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetBinaryCodeIDsForValue_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetBinaryCodeIDsForValue_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetBinaryCodeIDsForValue_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetBinaryCodeIDsForValue_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x56, 0x9a, 0xa6, 0xed, 0x3f, 0xfd, 0xa1, 0xe0, 0xb4, 0x95, 0x93, 0x88, 0x0a, 0xb5, 0x37,
	0x5c, 0x20, 0x07, 0x15, 0xa9, 0x1c, 0x24, 0x10, 0x4d, 0x39, 0xc8, 0x12, 0x0d, 0xc1, 0xa0, 0x4a,
	0x70, 0x63, 0x39, 0xf1, 0x12, 0x56, 0x24, 0xbb, 0xd1, 0xae, 0x43, 0xd5, 0xb7, 0x00, 0x5e, 0x81,
	0x37, 0x80, 0x67, 0xe0, 0x61, 0x08, 0x87, 0x1b, 0x5e, 0x80, 0xdd, 0x9d, 0x75, 0xd3, 0xb8, 0x2e,
	0x21, 0xdc, 0x24, 0xda, 0x9d, 0xf9, 0x66, 0xbe, 0x9d, 0x99, 0x6f, 0x0c, 0x77, 0x62, 0x99, 0x70,
	0x41, 0xea, 0x84, 0x75, 0x29, 0x23, 0xf5, 0x81, 0xe0, 0x1d, 0x12, 0x0f, 0x05, 0x91, 0x75, 0xda,
	0x0f, 0x1f, 0x91, 0xa4, 0x41, 0x59, 0x24, 0x8e, 0xf6, 0x78, 0x4c, 0xfc, 0xfb, 0xf2, 0x21, 0x17,
	0x07, 0x51, 0x6f, 0x48, 0xc2, 0xd6, 0xd0, 0x53, 0x9e, 0x09, 0x77, 0xae, 0x22, 0xdc, 0x43, 0xb8,
	0xf7, 0x67, 0x4c, 0xb5, 0x6c, 0x93, 0xbd, 0xd5, 0x17, 0x12, 0x43, 0x54, 0x2b, 0x93, 0x0c, 0x88,
	0x10, 0x5c, 0x58, 0x53, 0x6d, 0xd2, 0xd4, 0x27, 0x52, 0x46, 0x5d, 0x62, 0x8d, 0x5b, 0x59, 0x63,
	0x12, 0x51, 0xf6, 0x8a, 0x8b, 0x7e, 0x94, 0x50, 0xce, 0xd0, 0x69, 0xf3, 0x57, 0x09, 0xa0, 0x15,
	0x89, 0x48, 0x59, 0x89, 0x90, 0xce, 0x0e, 0x2c, 0x99, 0xdc, 0x21, 0x8d, 0xdd, 0xc2, 0xe5, 0xc2,
	0x95, 0xe5, 0xed, 0x9a, 0x67, 0x5f, 0x60, 0x39, 0x51, 0x96, 0x90, 0x2e, 0x41, 0xca, 0xc1, 0xa2,
	0xb9, 0xf4, 0x63, 0x67, 0x0b, 0xce, 0xa5, 0xb8, 0x90, 0x0d, 0x7b, 0x3d, 0xf7, 0xeb, 0xa2, 0x42,
	0x2f, 0x05, 0xcb, 0xd6, 0xa1, 0xa9, 0xee, 0x9c, 0xe7, 0xe0, 0xb6, 0xcd, 0xcb, 0x43, 0x95, 0x7b,
	0x40, 0x44, 0x72, 0x14, 0x1e, 0x27, 0x9b, 0x9b, 0x9e, 0x6c, 0x0d, 0xc1, 0x2d, 0x8b, 0x3d, 0xb0,
	0xa9, 0x77, 0xe1, 0xd2, 0x59, 0x51, 0x91, 0xca, 0x08, 0xa9, 0x54, 0x72, 0xe1, 0x86, 0xd8, 0x35,
	0x28, 0x19, 0x88, 0x5b, 0x34, 0x2c, 0xaa, 0x19, 0x16, 0x32, 0x11, 0x94, 0x75, 0x91, 0x04, 0x3a,
	0x3a, 0x1b, 0x00, 0x98, 0xc4, 0x64, 0xf8, 0x86, 0x19, 0xfe, 0x33, 0x57, 0x26, 0xe2, 0x53, 0x58,
	0x67, 0xaa, 0xbd, 0x61, 0xe7, 0xb5, 0xaa, 0x6d, 0x47, 0x95, 0x96, 0xca, 0x84, 0x76, 0xf4, 0x43,
	0xe7, 0xa7, 0x3f, 0x74, 0x55, 0x43, 0xf7, 0x26, 0x90, 0xea, 0x9d, 0x77, 0xa1, 0x96, 0x1f, 0x12,
	0x39, 0x7c, 0x47, 0x0e, 0x6e, 0x1e, 0xd6, 0x50, 0xf2, 0xa1, 0x9c, 0xd6, 0x45, 0x86, 0x94, 0x85,
	0x9c, 0x99, 0xc2, 0x97, 0x72, 0xf9, 0xb4, 0x39, 0xef, 0x91, 0x88, 0x21, 0x9f, 0x15, 0xdb, 0x44,
	0xe9, 0xb3, 0x27, 0x4c, 0x97, 0x7c, 0x07, 0xdc, 0x9c, 0x50, 0xc8, 0xe3, 0x07, 0xf2, 0x28, 0x67,
	0x30, 0x86, 0xc2, 0x63, 0x58, 0xc5, 0xf8, 0x1a, 0x14, 0x31, 0xdb, 0x28, 0xe9, 0x2e, 0x4c, 0xe7,
	0x70, 0x11, 0x2f, 0x7d, 0xb6, 0xcb, 0xb0, 0x75, 0xd2, 0xb9, 0x05, 0x95, 0xbc, 0x68, 0x48, 0xe3,
	0x27, 0xd2, 0x58, 0x3b, 0x05, 0xd3, 0x44, 0x36, 0xbf, 0xcc, 0xc3, 0x52, 0x40, 0xe4, 0x80, 0x33,
	0x49, 0x74, 0xf7, 0x8d, 0xa6, 0xec, 0xc0, 0x1f, 0x77, 0xdf, 0x4a, 0x16, 0xf5, 0xf6, 0x40, 0xff,
	0x06, 0xe8, 0xe8, 0xbc, 0x80, 0x0b, 0x5a, 0x4d, 0xe1, 0x09, 0x39, 0xa9, 0x01, 0x2e, 0x2a, 0xb0,
	0x97, 0x01, 0x67, 0x45, 0xb7, 0xaf, 0xce, 0xfe, 0xf8, 0x1c, 0xac, 0xf4, 0x27, 0x2f, 0x9c, 0x9b,
	0xb0, 0x68, 0x55, 0xac, 0x86, 0x51, 0x47, 0xdc, 0x38, 0x15, 0x11, 0x35, 0xbe, 0x8f, 0xff, 0x41,
	0xea, 0xae, 0x8a, 0x5b, 0x14, 0xfc, 0x50, 0xcd, 0x97, 0x46, 0xdd, 0xf6, 0x66, 0xd9, 0x3b, 0x5e,
	0x5a, 0x0b, 0x2f, 0xe0, 0x87, 0x81, 0x0e, 0x53, 0xfd, 0x38, 0x07, 0x45, 0x75, 0x70, 0xd6, 0x61,
	0x41, 0x1d, 0xf5, 0xa0, 0xbc, 0x6b, 0xaa, 0xf2, 0x94, 0x82, 0x92, 0x3a, 0xaa, 0x11, 0x68, 0xc0,
	0x79, 0xab, 0xba, 0x8e, 0x1e, 0x4a, 0x65, 0x7f, 0xdf, 0x9c, 0x3e, 0xd9, 0xff, 0xb7, 0xc7, 0x0c,
	0x62, 0xe7, 0xc6, 0x89, 0x65, 0xf3, 0xa1, 0x39, 0xc3, 0xb6, 0x79, 0x76, 0xa6, 0xba, 0x3e, 0x8d,
	0x0a, 0xff, 0xaa, 0xaf, 0xed, 0x74, 0x09, 0x7c, 0x1e, 0x15, 0xfe, 0x72, 0x0d, 0x34, 0x42, 0xa8,
	0x51, 0x9e, 0x29, 0xf5, 0xf8, 0x0b, 0xf1, 0xf2, 0x5e, 0x97, 0xcb, 0xf8, 0x4d, 0x6a, 0x8f, 0x67,
	0xff, 0x88, 0xb4, 0x17, 0xcc, 0x96, 0xbe, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x27, 0x74,
	0xce, 0x86, 0x06, 0x00, 0x00,
}