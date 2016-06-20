// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetBasicCharacteristics_Ad.proto
// DO NOT EDIT!

/*
Package st_GetBasicCharacteristics_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetBasicCharacteristics_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetBasicCharacteristics_Ad

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
	BasicCharacteristicNumber     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=basic_characteristic_number,json=basicCharacteristicNumber" json:"basic_characteristic_number,omitempty"`
	BasicCharacteristicNumberNull bool                        `protobuf:"varint,1001,opt,name=basic_characteristic_number_null,json=basicCharacteristicNumberNull" json:"basic_characteristic_number_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetBasicCharacteristicNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicCharacteristicNumber
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	EvaluateForPersons        *dstore_values.BooleanValue `protobuf:"bytes,10001,opt,name=evaluate_for_persons,json=evaluateForPersons" json:"evaluate_for_persons,omitempty"`
	EvaluateForNodes          *dstore_values.BooleanValue `protobuf:"bytes,10002,opt,name=evaluate_for_nodes,json=evaluateForNodes" json:"evaluate_for_nodes,omitempty"`
	Description               *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=description" json:"description,omitempty"`
	BasicCharacteristic       *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=basic_characteristic,json=basicCharacteristic" json:"basic_characteristic,omitempty"`
	PersonIdRequired          *dstore_values.BooleanValue `protobuf:"bytes,10005,opt,name=person_id_required,json=personIdRequired" json:"person_id_required,omitempty"`
	BasicCharacteristicNumber *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=basic_characteristic_number,json=basicCharacteristicNumber" json:"basic_characteristic_number,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetEvaluateForPersons() *dstore_values.BooleanValue {
	if m != nil {
		return m.EvaluateForPersons
	}
	return nil
}

func (m *Response_Row) GetEvaluateForNodes() *dstore_values.BooleanValue {
	if m != nil {
		return m.EvaluateForNodes
	}
	return nil
}

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

func (m *Response_Row) GetPersonIdRequired() *dstore_values.BooleanValue {
	if m != nil {
		return m.PersonIdRequired
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetBasicCharacteristics_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetBasicCharacteristics_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetBasicCharacteristics_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetBasicCharacteristics_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0xa6, 0x8e, 0xed, 0x2e, 0xd9, 0x0b, 0x97, 0xac, 0x48, 0xb6, 0x45, 0x29, 0xeb, 0x8d, 0x20,
	0x4c, 0x45, 0x6f, 0x14, 0x14, 0x71, 0xc5, 0x9f, 0x2a, 0x3b, 0x2c, 0xb9, 0x10, 0xfc, 0xc1, 0x61,
	0x66, 0x72, 0x1c, 0x83, 0x6d, 0x52, 0x4f, 0x52, 0xf7, 0x35, 0xd4, 0x55, 0x5f, 0xc4, 0xa7, 0xf2,
	0xd6, 0x27, 0x30, 0x33, 0x49, 0xd9, 0x4e, 0xb7, 0xd6, 0x2d, 0xde, 0x34, 0x9c, 0x9e, 0xf3, 0x7d,
	0x39, 0xe7, 0xfb, 0xce, 0x84, 0xdc, 0x15, 0xc6, 0x6a, 0x84, 0x01, 0xa8, 0x52, 0x2a, 0x18, 0x4c,
	0x50, 0x17, 0x20, 0xa6, 0x08, 0x66, 0x60, 0x6c, 0xfa, 0x04, 0xec, 0x7e, 0x66, 0x64, 0xf1, 0xf0,
	0x7d, 0x86, 0x59, 0x61, 0x01, 0xa5, 0xb1, 0xb2, 0x30, 0xe9, 0x03, 0x11, 0xbb, 0x42, 0xab, 0xe9,
	0x75, 0x8f, 0x8e, 0x3d, 0x3a, 0x5e, 0x09, 0xe9, 0xee, 0x84, 0xab, 0x3e, 0x65, 0xa3, 0x29, 0x18,
	0xcf, 0xd0, 0xdd, 0x6d, 0xde, 0x0f, 0x88, 0x1a, 0x43, 0xaa, 0xd7, 0x4c, 0x8d, 0xc1, 0x98, 0xac,
	0x84, 0x90, 0xbc, 0xba, 0x98, 0xb4, 0x99, 0x54, 0xef, 0x34, 0x8e, 0x33, 0x2b, 0xb5, 0xf2, 0x45,
	0x7b, 0x3f, 0x5b, 0x84, 0x1c, 0xba, 0x3e, 0x5c, 0x16, 0xd0, 0xd0, 0xd7, 0xa4, 0x97, 0x57, 0xbd,
	0xa5, 0x45, 0xa3, 0xb9, 0x54, 0x4d, 0xc7, 0x39, 0x20, 0x6b, 0xf5, 0x5b, 0xd7, 0xb6, 0x6e, 0xf6,
	0xe2, 0x30, 0x53, 0x68, 0x53, 0x2a, 0x0b, 0x25, 0xe0, 0x8b, 0x2a, 0xe2, 0xbb, 0xf9, 0xe9, 0xd9,
	0x92, 0x1a, 0x4d, 0x9f, 0x92, 0xfe, 0x0a, 0x72, 0x77, 0x8c, 0x46, 0xec, 0xd7, 0x86, 0xbb, 0x62,
	0x93, 0x5f, 0xfe, 0x2b, 0x4b, 0xe2, 0xaa, 0xf6, 0x8e, 0x3b, 0x64, 0x93, 0x83, 0x99, 0x68, 0x65,
	0x80, 0xde, 0x20, 0xed, 0x5a, 0x93, 0xd0, 0x5d, 0x37, 0x6e, 0x2a, 0xee, 0xf5, 0x7a, 0x54, 0xfd,
	0x72, 0x5f, 0x48, 0x5f, 0x92, 0xed, 0x4a, 0x8d, 0x74, 0x4e, 0x0e, 0x76, 0xae, 0x1f, 0x39, 0x70,
	0xbc, 0x00, 0x5e, 0x14, 0xed, 0xc0, 0xc5, 0xc3, 0x93, 0x98, 0x5f, 0x18, 0x37, 0xff, 0xa0, 0xb7,
	0xc9, 0x46, 0x70, 0x81, 0x45, 0x35, 0xe3, 0x95, 0x53, 0x8c, 0xde, 0xa3, 0x03, 0x7f, 0xf2, 0x59,
	0x39, 0x7d, 0x4e, 0x22, 0xd4, 0x47, 0xec, 0x7c, 0x8d, 0xba, 0x13, 0xaf, 0xb1, 0x36, 0xf1, 0x4c,
	0x8a, 0x98, 0xeb, 0x23, 0x5e, 0xb1, 0x74, 0x7f, 0x47, 0x24, 0x72, 0x01, 0xbd, 0x44, 0x3a, 0x2e,
	0x4c, 0xa5, 0x60, 0x9f, 0x13, 0xa7, 0x4e, 0x9b, 0xb7, 0x5d, 0x38, 0x14, 0x34, 0x21, 0x17, 0xa1,
	0x72, 0x2f, 0xb3, 0x90, 0xba, 0xe6, 0xd3, 0x89, 0x33, 0xdf, 0x31, 0xb0, 0x2f, 0xc9, 0x52, 0x87,
	0x73, 0xad, 0x47, 0x90, 0x29, 0xef, 0x30, 0x9d, 0x21, 0x1f, 0x6b, 0x3c, 0xf4, 0x38, 0xfa, 0x8c,
	0xd0, 0x06, 0x9f, 0xd2, 0x02, 0x0c, 0xfb, 0x7a, 0x06, 0xb6, 0xed, 0x39, 0xb6, 0xa4, 0x42, 0xd1,
	0x7b, 0x64, 0xcb, 0x1d, 0x05, 0xca, 0x49, 0x6d, 0xcc, 0x71, 0xd2, 0xb4, 0x35, 0x90, 0x18, 0x8b,
	0x52, 0x95, 0x9e, 0x63, 0xbe, 0xbe, 0x1a, 0x6d, 0xd9, 0x96, 0xb1, 0x6f, 0xff, 0xe6, 0xd9, 0x59,
	0xb2, 0x75, 0xd5, 0x68, 0x5e, 0x1d, 0xa7, 0x62, 0x8a, 0xf0, 0x71, 0x2a, 0x11, 0x04, 0xfb, 0x7e,
	0x96, 0xd1, 0x3c, 0x6e, 0x28, 0x78, 0x40, 0xd1, 0x37, 0xab, 0x3f, 0xaf, 0x1f, 0xc9, 0xff, 0x7c,
	0x5f, 0xfb, 0x6f, 0x49, 0x4f, 0xea, 0x85, 0xc5, 0x39, 0x79, 0xad, 0x5e, 0xdd, 0x2f, 0xb5, 0x11,
	0x1f, 0x66, 0x79, 0xb1, 0xf6, 0x83, 0x96, 0x77, 0xea, 0x27, 0xe3, 0xd6, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x11, 0xa6, 0x43, 0x76, 0x11, 0x05, 0x00, 0x00,
}