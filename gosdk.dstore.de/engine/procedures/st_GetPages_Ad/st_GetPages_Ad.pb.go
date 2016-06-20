// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetPages_Ad.proto
// DO NOT EDIT!

/*
Package st_GetPages_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetPages_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetPages_Ad

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
	PageNo     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=page_no,json=pageNo" json:"page_no,omitempty"`
	PageNoNull bool                        `protobuf:"varint,1001,opt,name=page_no_null,json=pageNoNull" json:"page_no_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	RowId    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Filename *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=filename" json:"filename,omitempty"`
	PageNo   *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=page_no,json=pageNo" json:"page_no,omitempty"`
	Page     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=page" json:"page,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetFilename() *dstore_values.StringValue {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *Response_Row) GetPageNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageNo
	}
	return nil
}

func (m *Response_Row) GetPage() *dstore_values.StringValue {
	if m != nil {
		return m.Page
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetPages_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetPages_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetPages_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/st_GetPages_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0xaa, 0xd3, 0x40,
	0x14, 0xa5, 0xa6, 0x37, 0xb6, 0x82, 0x32, 0x82, 0x8c, 0xa9, 0x48, 0xad, 0x2f, 0x22, 0x38, 0x29,
	0x5e, 0xd0, 0x07, 0x5f, 0x14, 0x44, 0xfa, 0xd0, 0x58, 0xe6, 0x41, 0xd0, 0x97, 0x10, 0xcd, 0x6e,
	0x18, 0x4c, 0x67, 0xca, 0xcc, 0xd4, 0xfe, 0x86, 0x97, 0xff, 0xf1, 0x1b, 0xfc, 0x0d, 0xff, 0xc2,
	0x49, 0x66, 0x7a, 0x7a, 0x92, 0x03, 0xa7, 0xe7, 0x25, 0xc9, 0x9e, 0xb5, 0xd6, 0xbe, 0xad, 0x09,
	0x3c, 0x29, 0x8c, 0x55, 0x1a, 0x13, 0x94, 0xa5, 0x90, 0x98, 0x6c, 0xb5, 0xfa, 0x8a, 0xc5, 0x4e,
	0xa3, 0x49, 0x8c, 0xcd, 0xde, 0xa3, 0x5d, 0xe5, 0x25, 0x9a, 0xec, 0x4d, 0xc1, 0x1c, 0x62, 0x15,
	0xb9, 0xe7, 0xe9, 0xcc, 0xd3, 0x59, 0x9b, 0x13, 0xdf, 0x0e, 0xc9, 0xbe, 0xe7, 0xd5, 0x0e, 0x8d,
	0x97, 0xc4, 0x77, 0xdb, 0x15, 0x50, 0x6b, 0xa5, 0x03, 0x34, 0x69, 0x43, 0x1b, 0x34, 0xc6, 0xa5,
	0x0a, 0xe0, 0xc3, 0x2e, 0x68, 0x73, 0x21, 0xd7, 0x4a, 0x6f, 0x72, 0x2b, 0x94, 0xf4, 0xa4, 0x19,
	0x02, 0xac, 0x72, 0x9d, 0x3b, 0x10, 0xb5, 0x21, 0xcf, 0x61, 0xb4, 0x75, 0x09, 0x32, 0xa9, 0x68,
	0x6f, 0xda, 0x7b, 0x74, 0xfd, 0xe9, 0x84, 0x85, 0x7e, 0x43, 0x47, 0x42, 0x5a, 0x2c, 0x51, 0x7f,
	0xac, 0x23, 0x3e, 0xac, 0xb9, 0xa9, 0x22, 0x0f, 0xe0, 0x46, 0x50, 0x65, 0x72, 0x57, 0x55, 0xf4,
	0xdf, 0xc8, 0x69, 0xc7, 0x1c, 0x3c, 0x9c, 0xba, 0xa3, 0xd9, 0xdf, 0x08, 0xc6, 0x1c, 0xcd, 0x56,
	0x49, 0x83, 0x64, 0x0e, 0x83, 0x66, 0x88, 0x50, 0x23, 0x66, 0xed, 0x9d, 0xf8, 0x01, 0xdf, 0xd5,
	0x4f, 0xee, 0x89, 0xe4, 0x13, 0xdc, 0xaa, 0xdb, 0xcf, 0xce, 0xf5, 0x4f, 0xaf, 0x4d, 0x23, 0x27,
	0x66, 0x1d, 0x71, 0x77, 0xca, 0xa5, 0x8b, 0x17, 0xc7, 0x98, 0xdf, 0xdc, 0xb4, 0x0f, 0xc8, 0x2b,
	0x18, 0x85, 0xb5, 0xd1, 0xa8, 0xc9, 0x78, 0xff, 0x42, 0x46, 0xbf, 0xd4, 0xa5, 0x7f, 0xf3, 0x03,
	0x9d, 0xbc, 0x86, 0x48, 0xab, 0x3d, 0xed, 0x37, 0xaa, 0xc7, 0xec, 0x32, 0x63, 0xd9, 0x61, 0x76,
	0xc6, 0xd5, 0x9e, 0xd7, 0xb2, 0xf8, 0x4f, 0x0f, 0x22, 0x17, 0x90, 0x3b, 0x30, 0x74, 0x61, 0x26,
	0x0a, 0xfa, 0x23, 0x75, 0xeb, 0x18, 0xf0, 0x81, 0x0b, 0x17, 0x05, 0x79, 0x09, 0xe3, 0xb5, 0xa8,
	0x50, 0x3a, 0x6b, 0xe8, 0xcf, 0xb4, 0xbd, 0xa8, 0x60, 0x86, 0xb1, 0x5a, 0xc8, 0xd2, 0x7b, 0x71,
	0x46, 0x26, 0x2f, 0x8e, 0x1e, 0xfe, 0x4a, 0xaf, 0x6e, 0x62, 0x02, 0xfd, 0xfa, 0x8b, 0xfe, 0x3e,
	0x5d, 0xab, 0x21, 0xbe, 0xfd, 0x00, 0x13, 0xa1, 0x3a, 0x53, 0x1f, 0x6f, 0xff, 0xe7, 0x79, 0xa9,
	0x4c, 0xf1, 0xed, 0x80, 0x17, 0xa7, 0x7f, 0x90, 0x2f, 0xc3, 0xe6, 0x46, 0x3e, 0xfb, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x8f, 0x0f, 0x88, 0x52, 0x03, 0x00, 0x00,
}