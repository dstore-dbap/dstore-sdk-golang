// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyRegions_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyRegions_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyRegions_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyRegions_Ad

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
	RegionId         *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	RegionIdNull     bool                        `protobuf:"varint,1001,opt,name=region_id_null,json=regionIdNull" json:"region_id_null,omitempty"`
	Region           *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
	RegionNull       bool                        `protobuf:"varint,1002,opt,name=region_null,json=regionNull" json:"region_null,omitempty"`
	DeleteRegion     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete_region,json=deleteRegion" json:"delete_region,omitempty"`
	DeleteRegionNull bool                        `protobuf:"varint,1003,opt,name=delete_region_null,json=deleteRegionNull" json:"delete_region_null,omitempty"`
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

func (m *Parameters) GetRegion() *dstore_values.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Parameters) GetDeleteRegion() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteRegion
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyRegions_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyRegions_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyRegions_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6a, 0x14, 0x41,
	0x10, 0x25, 0xbb, 0x6e, 0xb2, 0x56, 0xa2, 0x86, 0x16, 0x64, 0x9c, 0x45, 0x89, 0x11, 0xc1, 0x83,
	0xf6, 0xc8, 0x7a, 0xf1, 0x68, 0x04, 0x0f, 0x8b, 0x6c, 0x90, 0x3e, 0x08, 0x7a, 0x19, 0x26, 0x76,
	0x65, 0x68, 0x98, 0xe9, 0x0e, 0xdd, 0xb3, 0x2e, 0xfe, 0x85, 0xff, 0xa2, 0x3f, 0xa5, 0xfe, 0x84,
	0x35, 0x53, 0xbd, 0x49, 0x66, 0x5c, 0x10, 0x2f, 0xbb, 0x53, 0x5d, 0xef, 0xbd, 0xaa, 0xae, 0x57,
	0x0d, 0x73, 0x1d, 0x1a, 0xe7, 0x31, 0x43, 0x5b, 0x1a, 0x8b, 0xd9, 0x85, 0x77, 0x9f, 0x51, 0xaf,
	0x3c, 0x86, 0xac, 0x36, 0xf9, 0xd2, 0x69, 0x73, 0xfe, 0x55, 0x61, 0x69, 0x9c, 0x0d, 0xf9, 0x89,
	0x96, 0x94, 0x6e, 0x9c, 0x78, 0xc4, 0x1c, 0xc9, 0x1c, 0xb9, 0x05, 0x98, 0xde, 0x8d, 0xb2, 0x5f,
	0x8a, 0x6a, 0x85, 0x81, 0x79, 0xe9, 0xfd, 0x7e, 0x2d, 0xf4, 0xde, 0xf9, 0x98, 0x9a, 0xf5, 0x53,
	0x35, 0x86, 0x50, 0x94, 0x18, 0x93, 0x8f, 0x87, 0xc9, 0xa6, 0x30, 0xf6, 0xdc, 0xf9, 0xba, 0x68,
	0xa8, 0x1e, 0x83, 0x8e, 0x7f, 0x8c, 0x00, 0xde, 0x17, 0xbe, 0xa0, 0x2c, 0xfa, 0x20, 0x5e, 0xc1,
	0x4d, 0xdf, 0xb5, 0x93, 0x1b, 0x9d, 0xec, 0x1c, 0xed, 0x3c, 0xdd, 0x9f, 0xcf, 0x64, 0xec, 0x3b,
	0x36, 0x65, 0x6c, 0x83, 0x25, 0xfa, 0x0f, 0x6d, 0xa4, 0xa6, 0x8c, 0x5e, 0x68, 0xf1, 0x04, 0x6e,
	0x5f, 0x32, 0x73, 0xbb, 0xaa, 0xaa, 0xe4, 0xe7, 0x1e, 0xf1, 0xa7, 0xea, 0x60, 0x03, 0x39, 0xa5,
	0x43, 0x31, 0x87, 0x5d, 0x8e, 0x93, 0x51, 0xa7, 0x9e, 0x0e, 0xd4, 0x43, 0xe3, 0x8d, 0x2d, 0x59,
	0x3c, 0x22, 0xc5, 0x11, 0xec, 0x47, 0xe9, 0x4e, 0xf7, 0x17, 0xeb, 0x02, 0x9f, 0x75, 0xaa, 0xaf,
	0xe1, 0x96, 0xc6, 0x8a, 0xae, 0x90, 0x47, 0xf1, 0xf1, 0xd6, 0xd6, 0xcf, 0x9c, 0xab, 0xb0, 0xb0,
	0xac, 0x7e, 0xc0, 0x0c, 0x9e, 0xbe, 0x78, 0x0e, 0xa2, 0xa7, 0xc0, 0xa5, 0x7e, 0x73, 0xa9, 0xc3,
	0xeb, 0xd0, 0xb6, 0xe0, 0xf1, 0xf7, 0x11, 0x4c, 0x15, 0x86, 0x0b, 0xf2, 0x0d, 0xc5, 0x0b, 0x98,
	0x74, 0xa6, 0xc4, 0x81, 0x5d, 0x5e, 0x29, 0x1a, 0xcd, 0x86, 0xbd, 0x6d, 0x7f, 0x15, 0x03, 0xc5,
	0x47, 0x38, 0x6c, 0xed, 0xc8, 0xaf, 0xf9, 0x41, 0xf3, 0x18, 0x13, 0x59, 0x0e, 0xc8, 0x43, 0xd7,
	0x96, 0x14, 0x2f, 0xae, 0x62, 0x75, 0xa7, 0xee, 0x1f, 0x90, 0x83, 0x7b, 0x71, 0x0d, 0x68, 0x08,
	0xad, 0xe2, 0xc3, 0xbf, 0x14, 0x79, 0x49, 0x96, 0xfc, 0xaf, 0x36, 0x70, 0x71, 0x02, 0x63, 0xef,
	0xd6, 0xc9, 0x8d, 0x8e, 0x95, 0xc9, 0x7f, 0x6e, 0xab, 0xdc, 0x0c, 0x40, 0x2a, 0xb7, 0x56, 0x2d,
	0x37, 0x7d, 0x00, 0x63, 0xfa, 0x16, 0xf7, 0xc8, 0x64, 0xb7, 0x6e, 0x57, 0xe8, 0xdb, 0x29, 0x8d,
	0x64, 0xa2, 0x26, 0x14, 0x2e, 0xf4, 0x9b, 0x77, 0x30, 0x33, 0x6e, 0x20, 0x7c, 0xf5, 0x74, 0x3e,
	0x3d, 0xfb, 0x9f, 0x47, 0x75, 0xb6, 0xdb, 0x2d, 0xf0, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0x77, 0xdb, 0xa6, 0x8b, 0x03, 0x00, 0x00,
}