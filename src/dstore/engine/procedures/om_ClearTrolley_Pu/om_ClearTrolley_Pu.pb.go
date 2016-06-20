// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ClearTrolley_Pu.proto
// DO NOT EDIT!

/*
Package om_ClearTrolley_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ClearTrolley_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ClearTrolley_Pu

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
	UniqueId                 *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull             bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	InsertIntoStatistics     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=insert_into_statistics,json=insertIntoStatistics" json:"insert_into_statistics,omitempty"`
	InsertIntoStatisticsNull bool                        `protobuf:"varint,1002,opt,name=insert_into_statistics_null,json=insertIntoStatisticsNull" json:"insert_into_statistics_null,omitempty"`
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

func (m *Parameters) GetInsertIntoStatistics() *dstore_values.BooleanValue {
	if m != nil {
		return m.InsertIntoStatistics
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ClearTrolley_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ClearTrolley_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ClearTrolley_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x4a, 0xe3, 0x40,
	0x14, 0xa6, 0xcd, 0xf6, 0x67, 0x67, 0x97, 0xdd, 0x65, 0x76, 0x29, 0xd9, 0x94, 0x5d, 0x4a, 0x45,
	0x10, 0x84, 0xd4, 0x9f, 0x0b, 0xbd, 0x12, 0x51, 0xbc, 0x88, 0xd0, 0x52, 0x47, 0x11, 0xf4, 0x26,
	0xa4, 0xcd, 0x58, 0x06, 0x92, 0x99, 0x3a, 0x33, 0xb1, 0xf8, 0x16, 0x3e, 0x8b, 0x6f, 0xa5, 0x2f,
	0xa1, 0x27, 0x99, 0xb4, 0xb5, 0x69, 0x6e, 0xbc, 0x69, 0x73, 0xe6, 0x7c, 0xdf, 0x39, 0xdf, 0x7c,
	0xe7, 0x0c, 0xda, 0x0d, 0x95, 0x16, 0x92, 0xf6, 0x28, 0x9f, 0x30, 0x4e, 0x7b, 0x53, 0x29, 0xc6,
	0x34, 0x4c, 0x24, 0x55, 0x3d, 0x11, 0xfb, 0xa7, 0x11, 0x0d, 0xe4, 0x95, 0x14, 0x51, 0x44, 0x1f,
	0xfd, 0x61, 0xe2, 0x42, 0x56, 0x0b, 0xdc, 0x31, 0x14, 0xd7, 0x50, 0xdc, 0x75, 0x9c, 0xf3, 0x3b,
	0x2f, 0xfa, 0x10, 0x44, 0x09, 0x55, 0x86, 0xe6, 0xfc, 0x5d, 0xed, 0x44, 0xa5, 0x14, 0x32, 0x4f,
	0xb5, 0x57, 0x53, 0x31, 0x55, 0x2a, 0x98, 0xd0, 0x3c, 0xb9, 0x51, 0x4c, 0xea, 0x80, 0xf1, 0x3b,
	0x21, 0xe3, 0x40, 0x33, 0xc1, 0x0d, 0xa8, 0xfb, 0x56, 0x41, 0x68, 0x18, 0xc8, 0x00, 0xb2, 0x54,
	0x2a, 0x7c, 0x80, 0xbe, 0x26, 0x9c, 0xdd, 0x27, 0xd4, 0x67, 0xa1, 0x5d, 0xe9, 0x54, 0xb6, 0xbe,
	0xed, 0x39, 0x6e, 0x2e, 0x3b, 0x17, 0xa5, 0xb4, 0x64, 0x7c, 0x72, 0x9d, 0x06, 0xa4, 0x69, 0xc0,
	0x5e, 0x88, 0x37, 0xd1, 0x8f, 0x05, 0xd1, 0xe7, 0x49, 0x14, 0xd9, 0x2f, 0x0d, 0xa0, 0x37, 0xc9,
	0xf7, 0x39, 0x64, 0x00, 0x87, 0xf8, 0x02, 0xb5, 0x18, 0x57, 0x54, 0x6a, 0x9f, 0x71, 0x2d, 0x7c,
	0xa5, 0x41, 0x8b, 0xd2, 0x6c, 0xac, 0xec, 0x6a, 0xd6, 0xac, 0x5d, 0x68, 0x36, 0x12, 0x02, 0x2c,
	0xe2, 0xa6, 0xdb, 0x1f, 0x43, 0xf5, 0x80, 0x79, 0xb9, 0x20, 0xe2, 0x23, 0xd4, 0x2e, 0x2f, 0x69,
	0x64, 0xbc, 0x1a, 0x19, 0x76, 0x19, 0x37, 0x95, 0xd4, 0x7d, 0xae, 0xa2, 0x26, 0xa1, 0x6a, 0x2a,
	0x00, 0x80, 0x77, 0x50, 0x2d, 0xf3, 0xb7, 0x78, 0xf7, 0x7c, 0x64, 0xc6, 0xfb, 0xb3, 0xf4, 0x97,
	0x18, 0x20, 0xbe, 0x41, 0xbf, 0x52, 0x67, 0xfd, 0x0f, 0xd6, 0xc2, 0x5d, 0x2c, 0x20, 0xbb, 0x05,
	0x72, 0x71, 0x00, 0x7d, 0x88, 0xbd, 0x65, 0x4c, 0x7e, 0xc6, 0xab, 0x07, 0xf8, 0x10, 0x35, 0xf2,
	0x89, 0xda, 0x56, 0x56, 0xf1, 0xff, 0x5a, 0x45, 0x33, 0xef, 0xbe, 0xf9, 0x27, 0x73, 0x38, 0x3e,
	0x46, 0x96, 0x14, 0x33, 0xfb, 0x4b, 0xa9, 0x8e, 0x92, 0xfd, 0x9c, 0xdf, 0xdf, 0x25, 0x62, 0x46,
	0x52, 0xaa, 0xf3, 0x0f, 0x59, 0xf0, 0x8d, 0x5b, 0xa8, 0x0e, 0x51, 0xba, 0x0c, 0x4f, 0x03, 0x70,
	0xa4, 0x46, 0x6a, 0x10, 0x7a, 0xe1, 0xc9, 0x39, 0x98, 0x2e, 0x0a, 0x75, 0x97, 0x4f, 0xe0, 0x76,
	0xfb, 0x13, 0x8f, 0x63, 0x54, 0xcf, 0x36, 0x71, 0xff, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xee,
	0xf6, 0xa0, 0x52, 0x03, 0x00, 0x00,
}