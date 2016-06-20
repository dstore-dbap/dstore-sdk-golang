// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetRessourceUsage.proto
// DO NOT EDIT!

/*
Package mi_GetRessourceUsage is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetRessourceUsage.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetRessourceUsage

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
	SelectResult     *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=select_result,json=selectResult" json:"select_result,omitempty"`
	SelectResultNull bool                        `protobuf:"varint,1001,opt,name=select_result_null,json=selectResultNull" json:"select_result_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSelectResult() *dstore_values.BooleanValue {
	if m != nil {
		return m.SelectResult
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	Cpu             *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=cpu" json:"cpu,omitempty"`
	Physicalio      *dstore_values.IntegerValue                      `protobuf:"bytes,102,opt,name=physicalio" json:"physicalio,omitempty"`
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

func (m *Response) GetCpu() *dstore_values.IntegerValue {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Response) GetPhysicalio() *dstore_values.IntegerValue {
	if m != nil {
		return m.Physicalio
	}
	return nil
}

type Response_Row struct {
	RowId      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CPU        *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=c_p_u,json=cPU" json:"c_p_u,omitempty"`
	PhysicalIO *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=physical_i_o,json=physicalIO" json:"physical_i_o,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCPU() *dstore_values.IntegerValue {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *Response_Row) GetPhysicalIO() *dstore_values.IntegerValue {
	if m != nil {
		return m.PhysicalIO
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetRessourceUsage.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetRessourceUsage.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetRessourceUsage.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetRessourceUsage.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x6b, 0xd4, 0x40,
	0x10, 0xe7, 0x4c, 0xaf, 0x2d, 0x63, 0xc5, 0xb2, 0x82, 0xac, 0x39, 0x90, 0x72, 0xbe, 0xf8, 0xe2,
	0x5e, 0xb1, 0x20, 0x82, 0x20, 0x52, 0x10, 0xb9, 0x87, 0x9e, 0x65, 0xe1, 0x04, 0x7d, 0x59, 0xd2,
	0x64, 0x7a, 0x2e, 0xe6, 0x76, 0xc2, 0xee, 0xc6, 0xe2, 0x83, 0x9f, 0xc1, 0x7f, 0x9f, 0xd2, 0x6f,
	0xe1, 0x26, 0x9b, 0xb3, 0x97, 0xb3, 0xd0, 0xfa, 0x92, 0x30, 0xf9, 0xfd, 0x99, 0x99, 0xdf, 0x66,
	0xe1, 0xa8, 0x70, 0x9e, 0x2c, 0x4e, 0xd0, 0x2c, 0xb4, 0xc1, 0x49, 0x65, 0x29, 0xc7, 0xa2, 0xb6,
	0xe8, 0x26, 0x4b, 0xad, 0xde, 0xa0, 0x97, 0xe8, 0x1c, 0xd5, 0x36, 0xc7, 0xb9, 0xcb, 0x16, 0x28,
	0x02, 0xee, 0x89, 0x8d, 0xa3, 0x48, 0x44, 0x91, 0xb8, 0x8a, 0x99, 0xde, 0xeb, 0x8c, 0x3f, 0x67,
	0x65, 0x8d, 0x2e, 0x0a, 0xd3, 0x07, 0xfd, 0x6e, 0x68, 0x2d, 0xd9, 0x0e, 0x1a, 0xf5, 0xa1, 0x65,
	0x70, 0xfb, 0xdb, 0x30, 0x7d, 0xb4, 0x09, 0xfa, 0x4c, 0x9b, 0x73, 0xb2, 0xcb, 0xcc, 0x6b, 0x32,
	0x91, 0x34, 0xfe, 0x0a, 0x70, 0x9a, 0xd9, 0x2c, 0x80, 0x68, 0x1d, 0x7b, 0x05, 0x77, 0x1c, 0x96,
	0x98, 0x7b, 0x15, 0x96, 0xa9, 0x4b, 0xcf, 0x07, 0x07, 0x83, 0xc7, 0xb7, 0x9f, 0x8e, 0x44, 0x37,
	0x7b, 0x37, 0xd7, 0x19, 0x51, 0x89, 0x99, 0x79, 0xd7, 0x54, 0x72, 0x2f, 0x2a, 0x64, 0x2b, 0x60,
	0x4f, 0x80, 0xf5, 0x1c, 0x94, 0xa9, 0xcb, 0x92, 0xff, 0xde, 0x09, 0x3e, 0xbb, 0x72, 0x7f, 0x9d,
	0x3a, 0x0b, 0xc0, 0xf8, 0xd7, 0x16, 0xec, 0x86, 0xb2, 0x22, 0xe3, 0x90, 0x1d, 0xc2, 0xb0, 0x5d,
	0xae, 0xeb, 0x9a, 0x8a, 0x7e, 0x62, 0x71, 0xf1, 0xd7, 0xcd, 0x53, 0x46, 0x22, 0x7b, 0x0f, 0xfb,
	0xcd, 0x5a, 0x6a, 0x6d, 0x2f, 0x7e, 0xeb, 0x20, 0x09, 0x62, 0xb1, 0x21, 0xde, 0xdc, 0xfe, 0x24,
	0xd4, 0xd3, 0xcb, 0x5a, 0xde, 0x5d, 0xf6, 0x3f, 0xb0, 0xe7, 0xb0, 0xd3, 0xc5, 0xc9, 0x93, 0xd6,
	0xf1, 0xe1, 0x3f, 0x8e, 0x31, 0xec, 0x93, 0xf8, 0x96, 0x2b, 0x3a, 0x3b, 0x86, 0xc4, 0xd2, 0x05,
	0xdf, 0x6a, 0x55, 0x87, 0xe2, 0xfa, 0x63, 0x17, 0xab, 0x04, 0x84, 0xa4, 0x0b, 0xd9, 0x88, 0x43,
	0x8c, 0x49, 0x5e, 0xd5, 0x1c, 0xaf, 0x8c, 0x5f, 0x1b, 0x8f, 0x0b, 0xb4, 0x31, 0xfe, 0x86, 0xc7,
	0x5e, 0x00, 0x54, 0x1f, 0xbf, 0x38, 0x9d, 0x67, 0xa5, 0x26, 0x7e, 0x7e, 0xbd, 0x6a, 0x8d, 0x9e,
	0x7e, 0x1b, 0x40, 0x12, 0x1a, 0xb3, 0xfb, 0xb0, 0x1d, 0x5a, 0x2b, 0x5d, 0xf0, 0xef, 0xb3, 0xe0,
	0x30, 0x94, 0xc3, 0x50, 0x4e, 0x8b, 0xe6, 0x58, 0x72, 0x55, 0xa9, 0x9a, 0xff, 0x98, 0xdd, 0x64,
	0x9c, 0xd3, 0x39, 0x7b, 0x09, 0x7b, 0x2b, 0x7f, 0xa5, 0x15, 0xf1, 0x9f, 0xb3, 0xff, 0x98, 0x68,
	0xfa, 0xf6, 0x78, 0x0e, 0x23, 0x4d, 0x1b, 0xc1, 0x5d, 0x5e, 0xb2, 0x0f, 0xcf, 0x16, 0xe4, 0x8a,
	0x4f, 0x2b, 0xbc, 0xb8, 0xe9, 0x3d, 0x3c, 0xdb, 0x6e, 0x7f, 0xf9, 0xa3, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0xa5, 0xeb, 0xe5, 0xbf, 0x03, 0x00, 0x00,
}