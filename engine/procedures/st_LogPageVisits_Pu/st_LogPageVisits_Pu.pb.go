// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_LogPageVisits_Pu.proto
// DO NOT EDIT!

/*
Package st_LogPageVisits_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_LogPageVisits_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_LogPageVisits_Pu

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
	UniqueId     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PageNo       *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=page_no,json=pageNo" json:"page_no,omitempty"`
	PageNoNull   bool                        `protobuf:"varint,1002,opt,name=page_no_null,json=pageNoNull" json:"page_no_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPageNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageNo
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_LogPageVisits_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_LogPageVisits_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_LogPageVisits_Pu.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/st_LogPageVisits_Pu.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x4a, 0xe3, 0x40,
	0x14, 0xa6, 0xcd, 0xf6, 0x67, 0xcf, 0x96, 0xdd, 0x65, 0x16, 0x96, 0x6c, 0xca, 0x8a, 0x56, 0x04,
	0xaf, 0x26, 0x52, 0x15, 0xbd, 0x55, 0xf0, 0xa2, 0x60, 0x4b, 0x19, 0xa4, 0xa0, 0x37, 0x21, 0x9a,
	0x31, 0x0c, 0xb6, 0x33, 0x75, 0x66, 0x62, 0x5f, 0xc3, 0x77, 0xf1, 0x19, 0x7c, 0x18, 0x7d, 0x0a,
	0x27, 0x99, 0x89, 0xb5, 0x51, 0xd0, 0x9b, 0x24, 0x67, 0xbe, 0x9f, 0x9c, 0xf9, 0xce, 0x81, 0x7e,
	0xa2, 0xb4, 0x90, 0x34, 0xa4, 0x3c, 0x65, 0x9c, 0x86, 0x73, 0x29, 0xae, 0x68, 0x92, 0x49, 0xaa,
	0x42, 0xa5, 0xa3, 0x53, 0x91, 0x8e, 0xe3, 0x94, 0x4e, 0x98, 0x62, 0x5a, 0x45, 0xe3, 0x0c, 0x1b,
	0x58, 0x0b, 0xb4, 0x61, 0x35, 0xd8, 0x6a, 0xf0, 0x07, 0xc4, 0xe0, 0x8f, 0xb3, 0xbd, 0x8b, 0xa7,
	0x19, 0x55, 0x56, 0x17, 0xfc, 0x5b, 0xfd, 0x17, 0x95, 0x52, 0x48, 0x07, 0x75, 0x57, 0xa1, 0x19,
	0x55, 0xca, 0xf8, 0x39, 0x70, 0xb3, 0x0a, 0xea, 0x98, 0xf1, 0x6b, 0x21, 0x67, 0xb1, 0x66, 0x82,
	0x5b, 0x52, 0xef, 0xb1, 0x06, 0x30, 0x8e, 0x65, 0x6c, 0x50, 0x2a, 0x15, 0x3a, 0x80, 0xef, 0x19,
	0x67, 0xb7, 0x19, 0x8d, 0x58, 0xe2, 0xd7, 0xd6, 0x6b, 0xdb, 0x3f, 0xfa, 0x01, 0x76, 0x7d, 0xbb,
	0xa6, 0x94, 0x96, 0x8c, 0xa7, 0x93, 0xbc, 0x20, 0x6d, 0x4b, 0x1e, 0x24, 0x68, 0x0b, 0x7e, 0xbe,
	0x0a, 0x23, 0x9e, 0x4d, 0xa7, 0xfe, 0x53, 0xcb, 0xc8, 0xdb, 0xa4, 0x53, 0x52, 0x46, 0xe6, 0x10,
	0xed, 0x41, 0x6b, 0x6e, 0x3a, 0x8c, 0xb8, 0xf0, 0xeb, 0x85, 0x7b, 0xb7, 0xe2, 0xce, 0xb8, 0xa6,
	0x29, 0x95, 0xd6, 0xbe, 0x99, 0x73, 0x47, 0x26, 0x39, 0xe8, 0x38, 0x95, 0xb5, 0x7e, 0xb6, 0xd6,
	0x60, 0xe1, 0xdc, 0xb8, 0xf7, 0x50, 0x87, 0x36, 0xa1, 0x6a, 0x2e, 0xb8, 0xa2, 0x68, 0x07, 0x1a,
	0x45, 0x4a, 0xd5, 0x1b, 0xb8, 0xe4, 0x6d, 0x82, 0x27, 0xf9, 0x93, 0x58, 0x22, 0x3a, 0x87, 0xdf,
	0x79, 0x3e, 0xd1, 0x9b, 0x80, 0x4c, 0x83, 0x9e, 0x11, 0xe3, 0x8a, 0xb8, 0x1a, 0xe3, 0xd0, 0xd4,
	0x83, 0x65, 0x4d, 0x7e, 0xcd, 0x56, 0x0f, 0xd0, 0x21, 0xb4, 0xdc, 0x5c, 0x7c, 0xaf, 0x70, 0x5c,
	0x7b, 0xe7, 0x68, 0xa7, 0x36, 0xb4, 0x6f, 0x52, 0xd2, 0xd1, 0x11, 0x78, 0x52, 0x2c, 0xfc, 0x6f,
	0x85, 0x2a, 0xc4, 0x9f, 0xae, 0x0f, 0x2e, 0x03, 0xc0, 0x44, 0x2c, 0x48, 0xae, 0x0d, 0xfe, 0x83,
	0x67, 0xbe, 0xd1, 0x5f, 0x68, 0x9a, 0x2a, 0x9f, 0xe9, 0xfd, 0xc8, 0x44, 0xd2, 0x20, 0x0d, 0x53,
	0x0e, 0x92, 0xe3, 0x33, 0xe8, 0x32, 0x51, 0x31, 0x5e, 0xee, 0xf2, 0xc5, 0x7e, 0x2a, 0x54, 0x72,
	0x53, 0xe2, 0xc9, 0x17, 0xd7, 0xfd, 0xb2, 0x59, 0xac, 0xd6, 0xee, 0x4b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf8, 0x62, 0x12, 0x3e, 0x25, 0x03, 0x00, 0x00,
}
