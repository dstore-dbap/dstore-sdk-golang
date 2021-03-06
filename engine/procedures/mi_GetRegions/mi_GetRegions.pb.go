// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetRegions.proto
// DO NOT EDIT!

/*
Package mi_GetRegions is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetRegions.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetRegions

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
	RegionId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	RegionIdNull bool                        `protobuf:"varint,1001,opt,name=region_id_null,json=regionIdNull" json:"region_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetRegionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RegionId
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
	Region   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=region" json:"region,omitempty"`
	RegionId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetRegion() *dstore_values.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Response_Row) GetRegionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RegionId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetRegions.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetRegions.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetRegions.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_GetRegions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x4b, 0xe3, 0x40,
	0x14, 0xa5, 0x9b, 0xed, 0xc7, 0xce, 0x2e, 0xbb, 0xcb, 0x2c, 0x2c, 0xd9, 0x04, 0x96, 0x52, 0x11,
	0x14, 0x64, 0x22, 0xf6, 0x45, 0xf1, 0x4d, 0x10, 0xe9, 0x43, 0x83, 0xcc, 0x83, 0xa0, 0x2f, 0x21,
	0x9a, 0x6b, 0x18, 0x6c, 0x32, 0x65, 0x26, 0xb5, 0xbf, 0x42, 0x50, 0x7f, 0x80, 0xff, 0xcf, 0x7f,
	0xe1, 0x4d, 0x66, 0x62, 0x9b, 0x08, 0x8a, 0x2f, 0x09, 0x77, 0xce, 0x39, 0x77, 0xe6, 0x9c, 0x7b,
	0xc9, 0x4e, 0xa2, 0x0b, 0xa9, 0x20, 0x80, 0x3c, 0x15, 0x39, 0x04, 0x73, 0x25, 0xaf, 0x20, 0x59,
	0x28, 0xd0, 0x41, 0x26, 0xa2, 0x13, 0x28, 0x38, 0xa4, 0x42, 0xe6, 0x9a, 0x21, 0x50, 0x48, 0xea,
	0x1b, 0x36, 0x33, 0x6c, 0xd6, 0xa0, 0x78, 0x7f, 0x6c, 0xab, 0xdb, 0x78, 0xb6, 0x00, 0xab, 0xf0,
	0xfe, 0x35, 0xfb, 0x83, 0x52, 0x52, 0x59, 0xc8, 0x6f, 0x42, 0x19, 0x68, 0x1d, 0xa7, 0x60, 0xc1,
	0x8d, 0x36, 0x58, 0xc4, 0x22, 0xbf, 0x96, 0x2a, 0x8b, 0x0b, 0xbc, 0xcc, 0x90, 0x46, 0x19, 0x21,
	0xa7, 0xb1, 0x8a, 0x11, 0x04, 0xa5, 0xe9, 0x3e, 0xf9, 0xa6, 0xaa, 0xa7, 0x44, 0x22, 0x71, 0x3b,
	0xc3, 0xce, 0xd6, 0xf7, 0x3d, 0x9f, 0xd9, 0x07, 0xdb, 0x37, 0x89, 0xbc, 0x80, 0x14, 0xd4, 0x59,
	0x59, 0xf1, 0x81, 0x61, 0x4f, 0x12, 0xba, 0x49, 0x7e, 0xbe, 0x2a, 0xa3, 0x7c, 0x31, 0x9b, 0xb9,
	0xcf, 0x7d, 0xd4, 0x0f, 0xf8, 0x8f, 0x9a, 0x12, 0xe2, 0xe1, 0xe8, 0xc9, 0x21, 0x03, 0x0e, 0x7a,
	0x8e, 0x66, 0x81, 0xee, 0x92, 0x6e, 0x65, 0xc6, 0xde, 0xe4, 0xb1, 0x66, 0x34, 0xc6, 0xe8, 0x71,
	0xf9, 0xe5, 0x86, 0x48, 0xcf, 0xc9, 0xef, 0xd2, 0x46, 0xb4, 0xe6, 0xc3, 0xfd, 0x32, 0x74, 0x50,
	0xcc, 0x5a, 0xe2, 0xb6, 0xdb, 0x29, 0xd6, 0x93, 0x55, 0xcd, 0x7f, 0x65, 0xcd, 0x03, 0xb4, 0xde,
	0xb7, 0xf1, 0xb9, 0x4e, 0xd5, 0xf1, 0xff, 0x9b, 0x8e, 0x26, 0xdc, 0xa9, 0xf9, 0xf3, 0x9a, 0x4e,
	0x0f, 0x89, 0xa3, 0xe4, 0xd2, 0xfd, 0x5a, 0xa9, 0xb6, 0xd9, 0x3b, 0xf3, 0x65, 0xb5, 0x75, 0xc6,
	0xe5, 0x92, 0x97, 0x2a, 0xef, 0xae, 0x43, 0x1c, 0x2c, 0xe8, 0x5f, 0xd2, 0xc3, 0xb2, 0x8c, 0xfd,
	0x3e, 0xc4, 0x34, 0xba, 0xbc, 0x8b, 0x25, 0xe6, 0x3a, 0xc6, 0xf3, 0x4a, 0xec, 0x3e, 0x84, 0xcd,
	0x94, 0xec, 0x3c, 0x74, 0xa1, 0x44, 0x9e, 0x9a, 0x71, 0x58, 0x2a, 0x3d, 0x58, 0x1f, 0xe3, 0x63,
	0xf8, 0x89, 0x39, 0x1e, 0x85, 0xc4, 0x17, 0xb2, 0xe5, 0x61, 0xb5, 0xd1, 0x17, 0x41, 0x2a, 0x75,
	0x72, 0x53, 0xe3, 0xc9, 0x87, 0x4b, 0x7f, 0xd9, 0xab, 0xd6, 0x6c, 0xfc, 0x12, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0x4d, 0x6b, 0xf2, 0x25, 0x03, 0x00, 0x00,
}
