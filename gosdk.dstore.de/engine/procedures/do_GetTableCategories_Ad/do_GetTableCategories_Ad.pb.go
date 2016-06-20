// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/do_GetTableCategories_Ad.proto
// DO NOT EDIT!

/*
Package do_GetTableCategories_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/do_GetTableCategories_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package do_GetTableCategories_Ad

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
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	RowId           int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TableCategoryId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=table_category_id,json=tableCategoryId" json:"table_category_id,omitempty"`
	Description     *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=description" json:"description,omitempty"`
	TableCategory   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=table_category,json=tableCategory" json:"table_category,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTableCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableCategoryId
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetTableCategory() *dstore_values.StringValue {
	if m != nil {
		return m.TableCategory
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.do_GetTableCategories_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.do_GetTableCategories_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.do_GetTableCategories_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/do_GetTableCategories_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0x26, 0x73, 0x92, 0x0d, 0x65, 0x5b, 0x36, 0x0d, 0x86, 0xe7, 0xc0, 0x08, 0x19, 0x8c, 0x5d,
	0x29, 0x63, 0x1b, 0xac, 0x14, 0x7a, 0xd1, 0x94, 0xfe, 0xe4, 0x22, 0xa1, 0x88, 0x52, 0x48, 0x6f,
	0x8c, 0x12, 0x9d, 0x1a, 0xd1, 0xc4, 0x0a, 0x92, 0xd2, 0x90, 0x17, 0xe8, 0x75, 0x7f, 0xde, 0xb2,
	0x4f, 0x51, 0xd9, 0xb2, 0x49, 0xec, 0x52, 0xda, 0xde, 0xd8, 0x1c, 0x9f, 0xef, 0xe7, 0x9c, 0x4f,
	0x32, 0xfa, 0xcf, 0xb5, 0x91, 0x0a, 0xba, 0x10, 0x47, 0x22, 0x86, 0xee, 0x5c, 0xc9, 0x09, 0xf0,
	0x85, 0x02, 0xdd, 0xe5, 0x32, 0x3c, 0x04, 0x73, 0xc2, 0xc6, 0x53, 0xd8, 0x63, 0x06, 0x22, 0xa9,
	0x04, 0xe8, 0x70, 0x97, 0x13, 0x8b, 0x31, 0x12, 0xff, 0x74, 0x44, 0xe2, 0x88, 0xe4, 0x29, 0x74,
	0xf0, 0x25, 0x33, 0xb8, 0x64, 0xd3, 0x05, 0x68, 0x47, 0x0e, 0xbe, 0x15, 0x5d, 0x41, 0x29, 0xa9,
	0xb2, 0x56, 0xab, 0xd8, 0x9a, 0x81, 0xd6, 0x2c, 0x82, 0xac, 0xf9, 0xa3, 0xdc, 0x34, 0x4c, 0xc4,
	0xe7, 0x52, 0xcd, 0x98, 0x11, 0x32, 0x76, 0xa0, 0xce, 0x7b, 0x84, 0x8e, 0x99, 0x62, 0xb6, 0x09,
	0x4a, 0x77, 0xae, 0xaa, 0xe8, 0x1d, 0x05, 0x3d, 0x97, 0xb1, 0x06, 0xfc, 0x1b, 0xd5, 0x52, 0x2f,
	0xbf, 0xd2, 0xae, 0xfc, 0x6a, 0xfc, 0x09, 0x48, 0x71, 0x09, 0x37, 0xc7, 0x7e, 0xf2, 0xa4, 0x0e,
	0x88, 0x47, 0xe8, 0x53, 0xe2, 0x12, 0x6e, 0xd8, 0xf8, 0x6f, 0xda, 0x9e, 0x25, 0x93, 0x12, 0xb9,
	0x3c, 0xcc, 0xc0, 0xd6, 0xfd, 0x75, 0x4d, 0x9b, 0xb3, 0xe2, 0x07, 0xbc, 0x85, 0xde, 0x66, 0xdb,
	0xf9, 0x5e, 0xaa, 0xf8, 0xfd, 0x91, 0xa2, 0xdb, 0x7d, 0xe0, 0xde, 0x34, 0x87, 0xe3, 0x03, 0xe4,
	0x29, 0xb9, 0xf4, 0xab, 0x29, 0xeb, 0x1f, 0x79, 0xd9, 0x49, 0x90, 0x3c, 0x05, 0x42, 0xe5, 0x92,
	0x26, 0x02, 0xc1, 0x7d, 0x05, 0x79, 0xb6, 0xc0, 0x5f, 0x51, 0xdd, 0x96, 0xa1, 0xe0, 0xfe, 0xf5,
	0xd0, 0x06, 0x53, 0xa3, 0x35, 0x5b, 0xf6, 0x39, 0x3e, 0x42, 0x9f, 0x4d, 0xa2, 0x13, 0x4e, 0x9c,
	0xd0, 0x2a, 0x81, 0xdc, 0x0c, 0xd3, 0xec, 0x5a, 0xb9, 0x6d, 0x76, 0xb0, 0x22, 0xb6, 0x10, 0x50,
	0xa7, 0x49, 0x45, 0x9b, 0x66, 0xc3, 0x7e, 0x65, 0x95, 0x76, 0x50, 0x83, 0x83, 0x9e, 0x28, 0x31,
	0x4f, 0x13, 0xbc, 0x1d, 0x16, 0xf3, 0xcf, 0x34, 0xb4, 0x51, 0x22, 0x8e, 0x9c, 0xc4, 0x26, 0x1e,
	0xf7, 0xd0, 0xc7, 0xe2, 0x20, 0xfe, 0xdd, 0xf3, 0x0a, 0x1f, 0x0a, 0x43, 0xf4, 0x46, 0xa8, 0x25,
	0x64, 0x29, 0xab, 0xf5, 0x75, 0x3f, 0xdb, 0x8e, 0xa4, 0xe6, 0x17, 0x79, 0x9f, 0xbf, 0xe6, 0x8f,
	0x18, 0xd7, 0xd3, 0x8b, 0xf7, 0xf7, 0x21, 0x00, 0x00, 0xff, 0xff, 0x57, 0x7d, 0x62, 0x16, 0x4d,
	0x03, 0x00, 0x00,
}