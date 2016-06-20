// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ValidateTRITrigger_Ad.proto
// DO NOT EDIT!

/*
Package mi_ValidateTRITrigger_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ValidateTRITrigger_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ValidateTRITrigger_Ad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dstore_values "dstore/values"
import dstore_engine_error "dstore/engine/error"
import dstore_engine_message "dstore/engine/message"
import dstore_engine_metainformation "dstore/engine/metainformation"

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
	TriggerId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=trigger_id,json=triggerId" json:"trigger_id,omitempty"`
	TriggerIdNull bool                        `protobuf:"varint,1001,opt,name=trigger_id_null,json=triggerIdNull" json:"trigger_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTriggerId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TriggerId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ValidateTRITrigger_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ValidateTRITrigger_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ValidateTRITrigger_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0x66, 0xeb, 0x6f, 0x7f, 0x7e, 0x11, 0x99, 0x44, 0x90, 0xda, 0xa1, 0x8c, 0x09, 0xea, 0x55,
	0xe6, 0x3f, 0x50, 0xbc, 0x53, 0x50, 0xe8, 0xc5, 0xc6, 0x08, 0x63, 0xa0, 0x37, 0x25, 0xda, 0x63,
	0x09, 0xb4, 0xcd, 0x4c, 0x3a, 0xf7, 0x1a, 0x3e, 0x8e, 0xaf, 0xe4, 0x5b, 0x98, 0x36, 0x99, 0xb3,
	0x15, 0x41, 0x6f, 0xda, 0x9c, 0x9c, 0xef, 0x3b, 0xe7, 0x7c, 0x5f, 0x0e, 0x3a, 0x0f, 0x55, 0x26,
	0x24, 0x0c, 0x20, 0x8d, 0x78, 0x0a, 0x83, 0x99, 0x14, 0x8f, 0x10, 0xce, 0x25, 0xa8, 0x41, 0xc2,
	0x83, 0x29, 0x8b, 0x79, 0xc8, 0x32, 0x98, 0x50, 0x7f, 0x22, 0x79, 0x14, 0x81, 0x0c, 0xae, 0x42,
	0xa2, 0x31, 0x99, 0xc0, 0xfb, 0x86, 0x48, 0x0c, 0x91, 0xfc, 0x84, 0xf6, 0x36, 0x6d, 0x83, 0x17,
	0x16, 0xcf, 0x41, 0x19, 0xb2, 0xb7, 0x5d, 0xee, 0x0a, 0x52, 0x0a, 0x69, 0x53, 0xdd, 0x72, 0x2a,
	0x01, 0xa5, 0x58, 0x04, 0x36, 0xb9, 0x57, 0x4d, 0x66, 0x8c, 0xa7, 0x4f, 0x42, 0x26, 0x2c, 0xe3,
	0x22, 0x35, 0xa0, 0xfe, 0x33, 0x42, 0x63, 0x26, 0x99, 0x4e, 0x82, 0x54, 0xf8, 0x12, 0xa1, 0xcc,
	0x4e, 0xc3, 0x43, 0xb7, 0xd6, 0xab, 0x1d, 0xae, 0x9d, 0x74, 0x89, 0x1d, 0xde, 0x0e, 0xc5, 0xd3,
	0x0c, 0x34, 0x60, 0x9a, 0x47, 0xf4, 0xbf, 0x85, 0xfb, 0x21, 0x3e, 0x40, 0x9d, 0x15, 0x37, 0x48,
	0xe7, 0x71, 0xec, 0xbe, 0xb7, 0x74, 0x85, 0x36, 0x5d, 0xff, 0x04, 0x8d, 0xf4, 0x6d, 0xff, 0xad,
	0x8e, 0xda, 0x14, 0xd4, 0x4c, 0xa4, 0x0a, 0xf0, 0x11, 0x6a, 0x14, 0x82, 0x6c, 0x33, 0x8f, 0x94,
	0x9d, 0x32, 0x62, 0x6f, 0xf2, 0x2f, 0x35, 0x40, 0x7c, 0x87, 0x36, 0x72, 0x29, 0xc1, 0x17, 0x2d,
	0x6e, 0xbd, 0xe7, 0x68, 0x32, 0xa9, 0x90, 0xab, 0x8a, 0x87, 0x3a, 0xf6, 0x57, 0x31, 0xed, 0x24,
	0xe5, 0x0b, 0x7c, 0x81, 0x5a, 0xd6, 0x42, 0xd7, 0x29, 0x2a, 0xee, 0x7e, 0xab, 0x68, 0x0c, 0x1e,
	0x9a, 0x3f, 0x5d, 0xc2, 0xf1, 0x2d, 0x72, 0xa4, 0x58, 0xb8, 0xff, 0x0a, 0xd6, 0x19, 0xf9, 0xdd,
	0x73, 0x93, 0xa5, 0x0b, 0x84, 0x8a, 0x05, 0xcd, 0x0b, 0x78, 0x3b, 0xc8, 0xd1, 0x67, 0xbc, 0x85,
	0x9a, 0x3a, 0xca, 0xdf, 0xe0, 0x75, 0xa4, 0x7d, 0x69, 0xd0, 0x86, 0x0e, 0xfd, 0xf0, 0x7a, 0x8c,
	0xba, 0x5c, 0x54, 0xaa, 0xaf, 0xb6, 0xf0, 0xfe, 0xf8, 0xcf, 0xfb, 0xf9, 0xd0, 0x2c, 0xd6, 0xe0,
	0xf4, 0x23, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x64, 0xce, 0x6b, 0xdb, 0x02, 0x00, 0x00,
}