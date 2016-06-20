// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/do_GetdStoreTables_Ad.proto
// DO NOT EDIT!

/*
Package do_GetdStoreTables_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/do_GetdStoreTables_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package do_GetdStoreTables_Ad

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
	TableName           *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	TableNameNull       bool                        `protobuf:"varint,1001,opt,name=table_name_null,json=tableNameNull" json:"table_name_null,omitempty"`
	TableCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=table_category_id,json=tableCategoryId" json:"table_category_id,omitempty"`
	TableCategoryIdNull bool                        `protobuf:"varint,1002,opt,name=table_category_id_null,json=tableCategoryIdNull" json:"table_category_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTableName() *dstore_values.StringValue {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *Parameters) GetTableCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableCategoryId
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
	RowId           int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TableName       *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	TableCategoryId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=table_category_id,json=tableCategoryId" json:"table_category_id,omitempty"`
	Description     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=description" json:"description,omitempty"`
	TableCategory   *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=table_category,json=tableCategory" json:"table_category,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTableName() *dstore_values.StringValue {
	if m != nil {
		return m.TableName
	}
	return nil
}

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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.do_GetdStoreTables_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.do_GetdStoreTables_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.do_GetdStoreTables_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0x9a, 0xa6, 0x2d, 0xb7, 0x2a, 0x05, 0x57, 0xaa, 0x86, 0x44, 0x42, 0xa8, 0x5d, 0xc0,
	0xca, 0xe1, 0xd1, 0x05, 0x20, 0xb1, 0xa0, 0x15, 0x2a, 0x5d, 0x64, 0x84, 0x0c, 0x42, 0x82, 0xcd,
	0xc8, 0x8d, 0x2f, 0x23, 0x4b, 0x13, 0x3b, 0xb2, 0x1d, 0x2a, 0xfe, 0x82, 0xd7, 0x37, 0xf0, 0x1b,
	0x7c, 0x0f, 0x6c, 0xf8, 0x05, 0xec, 0xb1, 0x4b, 0x3a, 0xd3, 0x4a, 0x25, 0xdd, 0x24, 0xba, 0xbe,
	0xe7, 0xdc, 0x73, 0x72, 0xae, 0x63, 0xd8, 0x13, 0xd6, 0x69, 0x83, 0x43, 0x54, 0xa5, 0x54, 0x38,
	0x9c, 0x1a, 0x3d, 0x46, 0x31, 0x33, 0x68, 0x87, 0x42, 0x17, 0x87, 0xe8, 0xc4, 0xeb, 0xd0, 0x7e,
	0xc3, 0x8f, 0x2b, 0xb4, 0xc5, 0x73, 0x41, 0x3d, 0xc0, 0x69, 0xb2, 0x1b, 0x59, 0x34, 0xb2, 0xe8,
	0x85, 0xd0, 0xfe, 0x56, 0x1a, 0xfd, 0x91, 0x57, 0x33, 0xb4, 0x91, 0xd9, 0xbf, 0xd5, 0xd4, 0x43,
	0x63, 0xb4, 0x49, 0xad, 0x41, 0xb3, 0x35, 0x41, 0x6b, 0x79, 0x89, 0xa9, 0xb9, 0xdb, 0x6e, 0x3a,
	0x2e, 0xd5, 0x07, 0x6d, 0x26, 0xdc, 0x49, 0xad, 0x22, 0x68, 0xe7, 0x4f, 0x07, 0xe0, 0x15, 0x37,
	0xdc, 0x77, 0xd1, 0x58, 0xf2, 0x04, 0xc0, 0x05, 0x37, 0x85, 0xf2, 0x27, 0x59, 0xe7, 0x4e, 0xe7,
	0xde, 0xfa, 0xc3, 0x3e, 0x4d, 0xd6, 0x93, 0x2b, 0xeb, 0x8c, 0x54, 0xe5, 0xdb, 0x50, 0xb0, 0x6b,
	0x35, 0x3a, 0xf7, 0x60, 0x72, 0x17, 0x36, 0xe7, 0xd4, 0x42, 0xcd, 0xaa, 0x2a, 0xfb, 0xb5, 0xea,
	0x07, 0xac, 0xb1, 0x8d, 0x7f, 0xa0, 0xdc, 0x9f, 0x92, 0x43, 0xb8, 0x19, 0x81, 0x63, 0xee, 0xb0,
	0xd4, 0xe6, 0x53, 0x21, 0x45, 0xb6, 0x54, 0x4b, 0x0d, 0x5a, 0x52, 0x52, 0x79, 0x04, 0x9a, 0xa8,
	0x15, 0xc7, 0x1f, 0x24, 0xd2, 0x91, 0x20, 0x7b, 0xb0, 0x7d, 0x6e, 0x50, 0x14, 0xfe, 0x1d, 0x85,
	0xb7, 0x5a, 0x8c, 0x20, 0xbf, 0xf3, 0x73, 0x19, 0xd6, 0x18, 0xda, 0xa9, 0x56, 0x16, 0xc9, 0x7d,
	0xe8, 0xd5, 0x79, 0xb6, 0x7f, 0x6a, 0xda, 0x52, 0xcc, 0xfa, 0x45, 0xf8, 0x64, 0x11, 0x48, 0xde,
	0xc1, 0x8d, 0x90, 0x64, 0x71, 0x26, 0x4a, 0x6f, 0xbe, 0xeb, 0xc9, 0xb4, 0x45, 0x6e, 0x07, 0x3e,
	0xf2, 0xf5, 0xd1, 0xbc, 0x66, 0x9b, 0x93, 0xe6, 0x01, 0x79, 0x0c, 0xab, 0x69, 0x83, 0x59, 0xb7,
	0x9e, 0x78, 0xfb, 0xdc, 0xc4, 0xb8, 0xdf, 0x51, 0xfc, 0x66, 0xa7, 0x70, 0x72, 0x00, 0x5d, 0xa3,
	0x4f, 0xb2, 0xe5, 0x9a, 0xf5, 0x80, 0xfe, 0xc7, 0x55, 0xa3, 0xa7, 0x11, 0x50, 0xa6, 0x4f, 0x58,
	0x60, 0xf7, 0x7f, 0x2c, 0x41, 0xd7, 0x17, 0x64, 0x1b, 0x56, 0x7c, 0x19, 0x96, 0xf2, 0x39, 0xf7,
	0xa9, 0xf4, 0x58, 0xcf, 0x97, 0x3e, 0xee, 0xa7, 0x8d, 0xbb, 0xf1, 0x25, 0x5f, 0xe4, 0x72, 0xbc,
	0xbc, 0x68, 0xe7, 0x5f, 0xf3, 0x2b, 0x2c, 0xfd, 0x19, 0xac, 0x0b, 0xb4, 0x63, 0x23, 0xa7, 0x75,
	0xf4, 0xdf, 0x2e, 0xb7, 0x71, 0x16, 0x4f, 0xf6, 0xe1, 0x7a, 0xd3, 0x48, 0xf6, 0xfd, 0xf2, 0x09,
	0x1b, 0x0d, 0x13, 0xfb, 0x23, 0x18, 0x48, 0xdd, 0x0a, 0x79, 0xfe, 0x0a, 0xbc, 0xa7, 0x8b, 0xbd,
	0x0f, 0xc7, 0x2b, 0xf5, 0x3f, 0xf1, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x11, 0x68,
	0xb6, 0x58, 0x04, 0x00, 0x00,
}