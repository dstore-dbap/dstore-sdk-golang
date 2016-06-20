// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetNodeSymbols_Ad.proto
// DO NOT EDIT!

/*
Package im_GetNodeSymbols_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetNodeSymbols_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetNodeSymbols_Ad

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
	SymbolId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=symbol_id,json=symbolId" json:"symbol_id,omitempty"`
	SymbolIdNull bool                        `protobuf:"varint,1001,opt,name=symbol_id_null,json=symbolIdNull" json:"symbol_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSymbolId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SymbolId
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
	Symbol   *dstore_values.BytesValue   `protobuf:"bytes,10001,opt,name=symbol" json:"symbol,omitempty"`
	SymbolId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=symbol_id,json=symbolId" json:"symbol_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetSymbol() *dstore_values.BytesValue {
	if m != nil {
		return m.Symbol
	}
	return nil
}

func (m *Response_Row) GetSymbolId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SymbolId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetNodeSymbols_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetNodeSymbols_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetNodeSymbols_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetNodeSymbols_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x4a, 0xe3, 0x40,
	0x18, 0xa5, 0x9b, 0xed, 0xcf, 0xce, 0x2e, 0xbb, 0xcb, 0x2c, 0x2c, 0x69, 0x0a, 0x52, 0x2a, 0x82,
	0x57, 0xd3, 0xd2, 0x82, 0xe8, 0xa5, 0x05, 0x91, 0x5e, 0x34, 0xc8, 0x88, 0x82, 0xde, 0x84, 0xd4,
	0xf9, 0x0c, 0xc1, 0x24, 0x53, 0x66, 0x52, 0x4b, 0xdf, 0xc1, 0x0b, 0xf5, 0x19, 0x7c, 0x39, 0xdf,
	0xc2, 0x49, 0x66, 0x62, 0x9b, 0x28, 0xa8, 0x37, 0x2d, 0x67, 0xce, 0x39, 0x33, 0x73, 0xce, 0x37,
	0x41, 0x23, 0x26, 0x53, 0x2e, 0xa0, 0x0f, 0x49, 0x10, 0x26, 0xd0, 0x9f, 0x0b, 0x7e, 0x05, 0x6c,
	0x21, 0x40, 0xf6, 0xc3, 0xd8, 0x3b, 0x86, 0xd4, 0xe5, 0x0c, 0x4e, 0x57, 0xf1, 0x8c, 0x47, 0xd2,
	0x3b, 0x64, 0x44, 0xf1, 0x29, 0xc7, 0x3d, 0x6d, 0x22, 0xda, 0x44, 0xde, 0x53, 0x3a, 0xff, 0xcc,
	0xc6, 0xb7, 0x7e, 0xb4, 0x00, 0xa9, 0x8d, 0x4e, 0xbb, 0x7c, 0x1a, 0x08, 0xc1, 0x85, 0xa1, 0x3a,
	0x65, 0x2a, 0x06, 0x29, 0xfd, 0x00, 0x0c, 0xb9, 0x5d, 0x25, 0x53, 0x3f, 0x4c, 0xae, 0xb9, 0x88,
	0xfd, 0x34, 0xe4, 0x89, 0x16, 0xf5, 0x62, 0x84, 0x4e, 0x7c, 0xe1, 0x2b, 0x12, 0x84, 0xc4, 0xfb,
	0xe8, 0x87, 0xcc, 0x6f, 0xe3, 0x85, 0xcc, 0xae, 0x75, 0x6b, 0xbb, 0x3f, 0x87, 0x1d, 0x62, 0xee,
	0x6d, 0xee, 0x14, 0x26, 0x29, 0x04, 0x20, 0xce, 0x33, 0x44, 0x5b, 0x5a, 0x3d, 0x61, 0x78, 0x07,
	0xfd, 0x7e, 0x75, 0x7a, 0xc9, 0x22, 0x8a, 0xec, 0xe7, 0xa6, 0xf2, 0xb7, 0xe8, 0xaf, 0x42, 0xe2,
	0xaa, 0xc5, 0xde, 0x93, 0x85, 0x5a, 0x14, 0xe4, 0x9c, 0x27, 0x12, 0xf0, 0x00, 0xd5, 0xf3, 0x30,
	0xe6, 0x24, 0x87, 0x94, 0x1b, 0xd2, 0x41, 0x8f, 0xb2, 0x5f, 0xaa, 0x85, 0xf8, 0x02, 0xfd, 0xcd,
	0x62, 0x78, 0x1b, 0x39, 0xec, 0x6f, 0x5d, 0x4b, 0x99, 0x49, 0xc5, 0x5c, 0x4d, 0x3b, 0x55, 0x78,
	0xb2, 0xc6, 0xf4, 0x4f, 0x5c, 0x5e, 0x50, 0xd1, 0x9b, 0xa6, 0x3e, 0xdb, 0xca, 0x77, 0xdc, 0x7a,
	0xb3, 0xa3, 0x2e, 0x77, 0xaa, 0xff, 0x69, 0x21, 0xc7, 0x63, 0x64, 0x09, 0xbe, 0xb4, 0xbf, 0xe7,
	0xae, 0x01, 0xf9, 0x78, 0xcc, 0xa4, 0x68, 0x80, 0x50, 0xbe, 0xa4, 0x99, 0xd9, 0xb9, 0xab, 0x21,
	0x4b, 0x01, 0xfc, 0x1f, 0x35, 0x14, 0xcc, 0xda, 0xbf, 0x77, 0x55, 0x29, 0x75, 0x5a, 0x57, 0x50,
	0xd5, 0x3b, 0x44, 0x0d, 0xdd, 0xa3, 0xfd, 0xe0, 0xe6, 0x65, 0xb5, 0x2b, 0x63, 0x99, 0xad, 0x52,
	0x90, 0x7a, 0x28, 0x46, 0x89, 0x0f, 0x36, 0x87, 0xf9, 0xe8, 0x7e, 0x61, 0x9a, 0xe3, 0x33, 0xd4,
	0x09, 0x79, 0x25, 0xc9, 0xfa, 0x95, 0x5f, 0xee, 0x05, 0x5c, 0xb2, 0x9b, 0x82, 0x67, 0x9f, 0xfd,
	0x10, 0x66, 0x8d, 0xfc, 0xcd, 0x8d, 0x5e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x5a, 0xad, 0x4c,
	0x40, 0x03, 0x00, 0x00,
}