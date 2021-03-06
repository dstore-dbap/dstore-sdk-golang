// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifySessionManagement_Pu.proto
// DO NOT EDIT!

/*
Package mi_ModifySessionManagement_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifySessionManagement_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifySessionManagement_Pu

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
	SessionId       *dstore_values.StringValue `protobuf:"bytes,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	SessionIdNull   bool                       `protobuf:"varint,1001,opt,name=session_id_null,json=sessionIdNull" json:"session_id_null,omitempty"`
	KeyVariable     *dstore_values.StringValue `protobuf:"bytes,2,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull bool                       `protobuf:"varint,1002,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	Value           *dstore_values.StringValue `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ValueNull       bool                       `protobuf:"varint,1003,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSessionId() *dstore_values.StringValue {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifySessionManagement_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifySessionManagement_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifySessionManagement_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifySessionManagement_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0x14, 0x31,
	0x14, 0xa6, 0x3b, 0x6e, 0x7f, 0x4e, 0x95, 0xd5, 0x08, 0x32, 0x4e, 0xb1, 0x48, 0xbd, 0x50, 0x28,
	0x64, 0x45, 0x6f, 0x2c, 0x28, 0x82, 0xe0, 0xc5, 0x22, 0x53, 0x4a, 0x84, 0x82, 0x5e, 0x38, 0xa4,
	0x26, 0x1d, 0x42, 0x67, 0x92, 0x92, 0xcc, 0x74, 0xd9, 0xb7, 0xf0, 0x49, 0x7c, 0x03, 0x1f, 0x48,
	0x7d, 0x09, 0x4f, 0x26, 0x59, 0x77, 0x77, 0x04, 0x97, 0xbd, 0x99, 0xc9, 0xc9, 0xf9, 0x7e, 0x92,
	0x73, 0x4e, 0xe0, 0xb5, 0x70, 0x8d, 0xb1, 0x72, 0x2c, 0x75, 0xa9, 0xb4, 0x1c, 0x5f, 0x5b, 0xf3,
	0x55, 0x8a, 0xd6, 0x4a, 0x37, 0xae, 0x55, 0x91, 0x1b, 0xa1, 0x2e, 0x67, 0x1f, 0xa5, 0x73, 0xca,
	0xe8, 0x9c, 0x6b, 0x5e, 0xca, 0x5a, 0xea, 0xa6, 0x38, 0x6b, 0x29, 0x02, 0x1b, 0x43, 0x8e, 0x03,
	0x9b, 0x06, 0x36, 0xfd, 0x2f, 0x25, 0xbb, 0x1f, 0xad, 0x6e, 0x78, 0xd5, 0x4a, 0x17, 0x14, 0xb2,
	0x87, 0xab, 0xfe, 0xd2, 0x5a, 0x63, 0x63, 0xea, 0x60, 0x35, 0x55, 0xa3, 0x26, 0xaa, 0xc5, 0xe4,
	0x93, 0x7e, 0xb2, 0xe1, 0x4a, 0x5f, 0x1a, 0x5b, 0xf3, 0x06, 0x8d, 0x03, 0xe8, 0xe8, 0xfb, 0x00,
	0xe0, 0x8c, 0x5b, 0x8e, 0x59, 0x69, 0x1d, 0x39, 0x01, 0x70, 0xe1, 0x60, 0x85, 0x12, 0xe9, 0xd6,
	0xe3, 0xad, 0x67, 0xfb, 0x2f, 0x32, 0x1a, 0xaf, 0x10, 0x4f, 0xe5, 0x1a, 0xab, 0x74, 0x79, 0xee,
	0x03, 0xb6, 0x17, 0xd1, 0x13, 0x41, 0x9e, 0xc2, 0x68, 0x41, 0x2d, 0x74, 0x5b, 0x55, 0xe9, 0xcf,
	0x1d, 0x14, 0xd8, 0x65, 0x77, 0xfe, 0x82, 0x4e, 0x71, 0x97, 0xbc, 0x81, 0xdb, 0x57, 0x72, 0x56,
	0xdc, 0x70, 0xab, 0xf8, 0x45, 0x25, 0xd3, 0xc1, 0x5a, 0x97, 0x7d, 0xc4, 0x9f, 0x47, 0x38, 0x39,
	0x86, 0x7b, 0xcb, 0xf4, 0xe0, 0xf4, 0x2b, 0x38, 0x8d, 0x96, 0x80, 0x9d, 0xd7, 0x73, 0x18, 0x76,
	0x7a, 0x69, 0xb2, 0xd6, 0x24, 0x00, 0xc9, 0x21, 0x40, 0xb7, 0x08, 0xba, 0xbf, 0x83, 0xee, 0x5e,
	0xb7, 0xe5, 0x15, 0x8f, 0x7e, 0x0c, 0x60, 0x97, 0x49, 0x77, 0x6d, 0xb4, 0x93, 0x5e, 0xbe, 0x6b,
	0x47, 0xbf, 0x52, 0xb1, 0xd9, 0xa1, 0x55, 0xef, 0xfd, 0x97, 0x05, 0x20, 0xf9, 0x04, 0x77, 0x7d,
	0x23, 0x8a, 0xa5, 0x4e, 0x60, 0x01, 0x12, 0x24, 0xd3, 0x1e, 0xb9, 0xdf, 0xaf, 0x1c, 0xe3, 0xc9,
	0x22, 0x66, 0xa3, 0x7a, 0x75, 0x83, 0xbc, 0x82, 0x9d, 0x38, 0x00, 0x78, 0x5b, 0xaf, 0x78, 0xf8,
	0x8f, 0x62, 0x18, 0x8f, 0x3c, 0xfc, 0xd9, 0x1c, 0x4e, 0x3e, 0x40, 0x62, 0xcd, 0x34, 0xbd, 0xd5,
	0xb1, 0x4e, 0xe8, 0x06, 0x13, 0x4b, 0xe7, 0xa5, 0xa0, 0xcc, 0x4c, 0x99, 0x57, 0xc9, 0x1e, 0x41,
	0x82, 0x6b, 0xf2, 0x00, 0xb6, 0x31, 0xf2, 0x53, 0xf4, 0xed, 0x14, 0x8b, 0x33, 0x64, 0x43, 0x0c,
	0x27, 0xe2, 0xdd, 0x17, 0x38, 0x50, 0xa6, 0x67, 0xb1, 0x78, 0x52, 0x9f, 0xdf, 0x96, 0xc6, 0x89,
	0xab, 0x79, 0x5e, 0x6c, 0xfc, 0xea, 0x2e, 0xb6, 0xbb, 0xb9, 0x7e, 0xf9, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x1a, 0xeb, 0x74, 0x14, 0xb6, 0x03, 0x00, 0x00,
}
