// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_InsertTempdb_ImageData.proto
// DO NOT EDIT!

/*
Package mi_InsertTempdb_ImageData is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_InsertTempdb_ImageData.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_InsertTempdb_ImageData

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
	Data       *dstore_values.BytesValue   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	DataNull   bool                        `protobuf:"varint,1001,opt,name=data_null,json=dataNull" json:"data_null,omitempty"`
	SortNo     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull bool                        `protobuf:"varint,1002,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	Format     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=format" json:"format,omitempty"`
	FormatNull bool                        `protobuf:"varint,1003,opt,name=format_null,json=formatNull" json:"format_null,omitempty"`
	Delete     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetData() *dstore_values.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetFormat() *dstore_values.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_InsertTempdb_ImageData.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_InsertTempdb_ImageData.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_InsertTempdb_ImageData.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_InsertTempdb_ImageData.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x8a, 0x14, 0x31,
	0x10, 0xc7, 0xd9, 0x99, 0x9d, 0x0f, 0x6b, 0x05, 0x25, 0x82, 0xf4, 0xce, 0xa8, 0xac, 0xeb, 0x41,
	0x2f, 0xf6, 0xc8, 0xae, 0xc2, 0x82, 0x37, 0x51, 0x64, 0x0e, 0x3b, 0x48, 0x10, 0xc1, 0xbd, 0x34,
	0x19, 0x53, 0x36, 0x8d, 0xdd, 0xc9, 0x90, 0x64, 0x5c, 0x7c, 0x0b, 0x5f, 0xc7, 0xa7, 0xf1, 0xec,
	0xc7, 0x43, 0x58, 0xe9, 0xca, 0xb8, 0x4e, 0xab, 0xa0, 0x97, 0xe9, 0xa4, 0xea, 0xff, 0xff, 0xa5,
	0xa6, 0x2a, 0x81, 0x13, 0xed, 0x83, 0x75, 0x38, 0x43, 0x53, 0x56, 0x06, 0x67, 0x2b, 0x67, 0xdf,
	0xa0, 0x5e, 0x3b, 0xf4, 0xb3, 0xa6, 0x2a, 0xe6, 0xc6, 0xa3, 0x0b, 0x2f, 0xb1, 0x59, 0xe9, 0x65,
	0x31, 0x6f, 0x54, 0x89, 0x4f, 0x55, 0x50, 0x39, 0x89, 0x82, 0x15, 0x77, 0xd9, 0x99, 0xb3, 0x33,
	0xff, 0xab, 0x7c, 0x72, 0x2d, 0x1d, 0xf1, 0x5e, 0xd5, 0x6b, 0xf4, 0xec, 0x9e, 0xec, 0x6f, 0x9f,
	0x8b, 0xce, 0x59, 0x97, 0x52, 0xd3, 0xed, 0x54, 0x83, 0xde, 0x13, 0x2a, 0x25, 0xef, 0x74, 0x93,
	0x41, 0x55, 0xe6, 0xad, 0x75, 0x8d, 0x0a, 0x95, 0x35, 0x2c, 0x3a, 0xfc, 0xdc, 0x03, 0x78, 0xa1,
	0x9c, 0xa2, 0x2c, 0x3a, 0x2f, 0xee, 0xc3, 0xae, 0xa6, 0x42, 0xb2, 0x9d, 0x83, 0x9d, 0x7b, 0x7b,
	0x47, 0xfb, 0x79, 0x2a, 0x3c, 0xd5, 0xb3, 0xfc, 0x10, 0xd0, 0xbf, 0x8a, 0x6b, 0xd9, 0xca, 0xc4,
	0x0d, 0xb8, 0x14, 0xbf, 0x85, 0x59, 0xd7, 0x75, 0xf6, 0x65, 0x44, 0xa6, 0xb1, 0x1c, 0xc7, 0xc8,
	0x82, 0x02, 0xe2, 0x21, 0x8c, 0xbc, 0x75, 0xa1, 0x30, 0x36, 0xeb, 0xb5, 0xbc, 0x69, 0x87, 0x57,
	0x99, 0x80, 0x25, 0x3a, 0x26, 0x0e, 0xa3, 0x76, 0x61, 0xc5, 0x6d, 0xb8, 0x9c, 0x5c, 0x8c, 0xfd,
	0xca, 0x58, 0xe0, 0x74, 0x0b, 0x3e, 0x82, 0x21, 0xff, 0x8f, 0xac, 0xdf, 0x72, 0x27, 0x1d, 0xae,
	0x0f, 0xae, 0x32, 0x65, 0xc2, 0xb2, 0x52, 0x1c, 0xc0, 0x1e, 0xaf, 0x98, 0xfa, 0x2d, 0x51, 0x39,
	0xd6, 0x52, 0x8f, 0x61, 0xa8, 0xb1, 0xa6, 0x3e, 0x64, 0xbb, 0x7f, 0xac, 0x76, 0x69, 0x6d, 0x8d,
	0xca, 0x24, 0x2c, 0x4b, 0x23, 0x96, 0x57, 0x8c, 0xfd, 0x9e, 0xb0, 0x1c, 0x8b, 0xd8, 0xc3, 0x4f,
	0x3d, 0x18, 0x4b, 0xf4, 0x2b, 0x4b, 0x23, 0x17, 0x0f, 0x60, 0xd0, 0xce, 0x2f, 0x35, 0xf8, 0x67,
	0xe1, 0xe9, 0x66, 0xf0, 0x6c, 0x9f, 0xc5, 0x5f, 0xc9, 0x42, 0xf1, 0x1a, 0xae, 0xc6, 0xc9, 0x15,
	0xbf, 0x8c, 0x8e, 0xba, 0xd9, 0x27, 0x73, 0xde, 0x31, 0x77, 0x07, 0x7c, 0x4a, 0xfb, 0xf9, 0xc5,
	0x5e, 0x5e, 0x69, 0xb6, 0x03, 0xe2, 0x04, 0x46, 0xe9, 0xc6, 0x50, 0x1f, 0x23, 0xf1, 0xd6, 0x6f,
	0x44, 0xbe, 0x4f, 0xa7, 0xfc, 0x95, 0x1b, 0xb9, 0x78, 0x0e, 0x7d, 0x67, 0xcf, 0xa9, 0x4f, 0xd1,
	0xf5, 0x28, 0xff, 0xc7, 0xeb, 0x9d, 0x6f, 0xda, 0x90, 0x4b, 0x7b, 0x2e, 0x23, 0x61, 0x72, 0x13,
	0xfa, 0xb4, 0x16, 0xd7, 0x61, 0x48, 0xbb, 0xa2, 0xd2, 0xd9, 0xc7, 0x05, 0x35, 0x66, 0x20, 0x07,
	0xb4, 0x9d, 0xeb, 0x27, 0x67, 0x30, 0xad, 0x6c, 0x07, 0x7f, 0xf1, 0xee, 0xce, 0x1e, 0x97, 0xd6,
	0xeb, 0x77, 0x9b, 0xbc, 0xfe, 0xaf, 0xa7, 0xb9, 0x1c, 0xb6, 0x0f, 0xe0, 0xf8, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x31, 0x5b, 0x32, 0x1d, 0xd7, 0x03, 0x00, 0x00,
}
