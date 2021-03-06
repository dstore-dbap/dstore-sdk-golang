// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetUserInfo.proto
// DO NOT EDIT!

/*
Package mi_GetUserInfo is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetUserInfo.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetUserInfo

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
	UserId          *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=user_id,json=userId" json:"user_id,omitempty"`
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

func (m *Response) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type Response_Row struct {
	RowId  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	UserId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetUserInfo.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetUserInfo.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetUserInfo.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_GetUserInfo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x4b, 0xe3, 0x40,
	0x18, 0xa6, 0x9b, 0xed, 0x07, 0xb3, 0xbb, 0x6c, 0x19, 0x41, 0x62, 0x2a, 0x52, 0xea, 0x45, 0x84,
	0x4e, 0x8b, 0x1f, 0xe0, 0xc1, 0x83, 0x08, 0x22, 0x39, 0xb4, 0xca, 0xa0, 0x82, 0x5e, 0x42, 0xda,
	0xbc, 0x86, 0x60, 0x32, 0x53, 0x66, 0x26, 0xf6, 0xe4, 0x7f, 0xd0, 0xb3, 0xbf, 0xd0, 0x7f, 0xe1,
	0x24, 0x93, 0xd2, 0x26, 0x8a, 0x7a, 0x69, 0x79, 0xf3, 0x7c, 0xbc, 0xf3, 0x3c, 0x99, 0xa0, 0x7e,
	0x20, 0x15, 0x17, 0x30, 0x00, 0x16, 0x46, 0x0c, 0x06, 0x33, 0xc1, 0xa7, 0x10, 0xa4, 0x02, 0xe4,
	0x20, 0x89, 0xbc, 0x73, 0x50, 0xd7, 0x12, 0x84, 0xcb, 0xee, 0x39, 0xd1, 0x88, 0xe2, 0x78, 0xd3,
	0xd0, 0x89, 0xa1, 0x93, 0x32, 0xc7, 0x59, 0x2b, 0xcc, 0x1e, 0xfd, 0x38, 0x05, 0x69, 0x24, 0xce,
	0x46, 0x79, 0x03, 0x08, 0xc1, 0x45, 0x01, 0x75, 0xca, 0x50, 0x02, 0x52, 0xfa, 0x21, 0x14, 0xe0,
	0x76, 0x15, 0x54, 0x7e, 0xa4, 0x97, 0x88, 0xc4, 0x57, 0x11, 0x67, 0x86, 0xd4, 0x7b, 0x42, 0xe8,
	0xd2, 0x17, 0xbe, 0x06, 0x41, 0x48, 0x7c, 0x82, 0xfe, 0x49, 0x88, 0x61, 0xaa, 0x3c, 0x1d, 0x20,
	0x8d, 0x95, 0x5d, 0xeb, 0xd6, 0x76, 0xfe, 0xec, 0x75, 0x48, 0x71, 0xea, 0xe2, 0x5c, 0x13, 0xce,
	0x63, 0xf0, 0xd9, 0x4d, 0x36, 0xd1, 0xbf, 0x46, 0x41, 0x73, 0x01, 0xee, 0x23, 0x5c, 0x72, 0xf0,
	0x58, 0x1a, 0xc7, 0xf6, 0x5b, 0x53, 0xfb, 0xb4, 0x68, 0x7b, 0x95, 0x3a, 0xd6, 0x40, 0xef, 0xd5,
	0x42, 0x2d, 0x3d, 0xce, 0x38, 0x93, 0x80, 0x87, 0xa8, 0x9e, 0x87, 0x2b, 0xb6, 0x3a, 0xa4, 0xdc,
	0x95, 0x09, 0x7e, 0x96, 0xfd, 0x52, 0x43, 0xc4, 0xb7, 0xa8, 0x9d, 0xc5, 0xf2, 0x56, 0x72, 0xd9,
	0xbf, 0xba, 0x96, 0x16, 0x93, 0x8a, 0xb8, 0x9a, 0x7e, 0xa4, 0x67, 0x77, 0x39, 0xd3, 0xff, 0x49,
	0xf9, 0x01, 0x3e, 0x42, 0xcd, 0xa2, 0x4e, 0xdb, 0xca, 0x1d, 0xb7, 0x3e, 0x38, 0x9a, 0xb2, 0x47,
	0xe6, 0x9f, 0x2e, 0xe8, 0xf8, 0x18, 0x59, 0x82, 0xcf, 0xed, 0xdf, 0xb9, 0x6a, 0x97, 0x7c, 0xf5,
	0xc2, 0xc9, 0x22, 0x3b, 0xa1, 0x7c, 0x4e, 0x33, 0x19, 0x3e, 0x40, 0xcd, 0x54, 0xa3, 0x5e, 0x14,
	0xd8, 0xf0, 0x69, 0xf9, 0x11, 0x53, 0x10, 0x82, 0x30, 0xe5, 0x37, 0x32, 0xae, 0x1b, 0x38, 0x57,
	0xc8, 0xd2, 0x0e, 0x78, 0x1d, 0x35, 0xb4, 0x47, 0xa6, 0x7d, 0x1e, 0x6b, 0x71, 0x9d, 0xd6, 0xf5,
	0xe8, 0x06, 0xf8, 0x70, 0x69, 0xfa, 0x32, 0xfe, 0xb1, 0xeb, 0xe9, 0x05, 0xea, 0x44, 0xbc, 0x12,
	0x60, 0x79, 0xc1, 0xef, 0x86, 0x21, 0x97, 0xc1, 0xc3, 0x02, 0x0f, 0xbe, 0xff, 0x06, 0x26, 0x8d,
	0xfc, 0xd2, 0xed, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x19, 0xb8, 0xc8, 0xe0, 0x35, 0x03, 0x00,
	0x00,
}
