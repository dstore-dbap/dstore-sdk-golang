// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ResetBatchJob_Ad.proto
// DO NOT EDIT!

/*
Package mi_ResetBatchJob_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ResetBatchJob_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ResetBatchJob_Ad

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
	ProcedureName     *dstore_values.StringValue `protobuf:"bytes,1,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	ProcedureNameNull bool                       `protobuf:"varint,1001,opt,name=procedure_name_null,json=procedureNameNull" json:"procedure_name_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ResetBatchJob_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ResetBatchJob_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ResetBatchJob_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_ResetBatchJob_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x4a, 0xf3, 0x30,
	0x18, 0x66, 0xeb, 0xb7, 0x1f, 0xf2, 0xe1, 0x5f, 0x07, 0x52, 0x3b, 0x94, 0x39, 0x4f, 0x3c, 0x4a,
	0x65, 0x22, 0x78, 0xba, 0x81, 0x07, 0x13, 0x36, 0x24, 0x88, 0xa0, 0x27, 0xa5, 0x5b, 0x63, 0x0d,
	0xae, 0xc9, 0x48, 0x32, 0x77, 0xea, 0x25, 0x78, 0x2f, 0x5e, 0x95, 0x77, 0xe1, 0xdb, 0x26, 0xdb,
	0x6c, 0x15, 0xf4, 0xa4, 0xed, 0x9b, 0xe7, 0x27, 0x79, 0x9e, 0x06, 0xf5, 0x62, 0xa5, 0x85, 0xa4,
	0x01, 0xe5, 0x09, 0xe3, 0x34, 0x98, 0x4b, 0x31, 0xa5, 0xf1, 0x42, 0x52, 0x15, 0xa4, 0x2c, 0x24,
	0x54, 0x51, 0x3d, 0x88, 0xf4, 0xf4, 0xe9, 0x5a, 0x4c, 0xc2, 0x7e, 0x8c, 0x01, 0xd6, 0xc2, 0x3d,
	0x36, 0x1a, 0x6c, 0x34, 0xf8, 0x07, 0xa2, 0xdf, 0xb2, 0xb6, 0x2f, 0xd1, 0x6c, 0x41, 0x95, 0xd1,
	0xf9, 0x07, 0xc5, 0xbd, 0xa8, 0x94, 0x42, 0x5a, 0xa8, 0x5d, 0x84, 0x52, 0xaa, 0x54, 0x94, 0x50,
	0x0b, 0x9e, 0x94, 0x41, 0x1d, 0x31, 0xfe, 0x28, 0x64, 0x1a, 0x69, 0x26, 0xb8, 0x21, 0x75, 0x5f,
	0x2b, 0x08, 0xdd, 0x44, 0x32, 0x02, 0x94, 0x4a, 0xe5, 0xf6, 0xd1, 0xf6, 0x3a, 0x4b, 0xc8, 0x61,
	0xd5, 0xab, 0x74, 0x2a, 0xa7, 0xff, 0x7b, 0x3e, 0xb6, 0x87, 0xb7, 0x27, 0x53, 0x5a, 0x32, 0x9e,
	0xdc, 0x65, 0x03, 0xd9, 0x5a, 0x2b, 0xc6, 0x20, 0x70, 0x03, 0xd4, 0x2a, 0x5a, 0x84, 0x7c, 0x31,
	0x9b, 0x79, 0x1f, 0x0d, 0x30, 0x6a, 0x92, 0xbd, 0x02, 0x79, 0x0c, 0x48, 0xf7, 0xbd, 0x8a, 0x9a,
	0xd0, 0xc4, 0x5c, 0x70, 0x45, 0xdd, 0x33, 0x54, 0xcb, 0x03, 0x96, 0xf7, 0xb5, 0xa5, 0x99, 0xf0,
	0x57, 0xd9, 0x93, 0x18, 0xa2, 0x7b, 0x8f, 0x76, 0xb3, 0x68, 0xe1, 0x97, 0x6c, 0x5e, 0xb5, 0xe3,
	0x80, 0x18, 0x97, 0xc4, 0xe5, 0x06, 0x46, 0x30, 0x0f, 0x37, 0x33, 0xd9, 0x49, 0x8b, 0x0b, 0xee,
	0x25, 0x6a, 0xd8, 0x4a, 0x3d, 0x27, 0x77, 0x3c, 0xfa, 0xe6, 0x68, 0x0a, 0x1f, 0x99, 0x37, 0x59,
	0xd1, 0xa1, 0x47, 0x47, 0x8a, 0xa5, 0xf7, 0x2f, 0x57, 0x05, 0xf8, 0xd7, 0x3f, 0x8f, 0x57, 0x05,
	0x60, 0x22, 0x96, 0x24, 0xd3, 0xfa, 0x87, 0xc8, 0x81, 0x6f, 0x77, 0x1f, 0xd5, 0x61, 0x0a, 0x59,
	0xec, 0xbd, 0x8d, 0xa1, 0x92, 0x1a, 0xa9, 0xc1, 0x38, 0x8c, 0x07, 0xb7, 0xa8, 0xcd, 0x44, 0xc9,
	0x78, 0x73, 0x0d, 0x1f, 0x2e, 0x12, 0xa1, 0xe2, 0xe7, 0x15, 0x1e, 0xff, 0xf1, 0xa6, 0x4e, 0xea,
	0xf9, 0xad, 0x38, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x7b, 0x3d, 0xe6, 0xe0, 0x02, 0x00,
	0x00,
}
