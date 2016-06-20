// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_InsertVisitorInformation_Pu.proto
// DO NOT EDIT!

/*
Package mi_InsertVisitorInformation_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_InsertVisitorInformation_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_InsertVisitorInformation_Pu

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
	UniqueId                     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                 bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	VisitorInformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=visitor_information_type_id,json=visitorInformationTypeId" json:"visitor_information_type_id,omitempty"`
	VisitorInformationTypeIdNull bool                        `protobuf:"varint,1002,opt,name=visitor_information_type_id_null,json=visitorInformationTypeIdNull" json:"visitor_information_type_id_null,omitempty"`
	VisitorInformation           *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=visitor_information,json=visitorInformation" json:"visitor_information,omitempty"`
	VisitorInformationNull       bool                        `protobuf:"varint,1003,opt,name=visitor_information_null,json=visitorInformationNull" json:"visitor_information_null,omitempty"`
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

func (m *Parameters) GetVisitorInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VisitorInformationTypeId
	}
	return nil
}

func (m *Parameters) GetVisitorInformation() *dstore_values.StringValue {
	if m != nil {
		return m.VisitorInformation
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_InsertVisitorInformation_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_InsertVisitorInformation_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_InsertVisitorInformation_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xdf, 0xea, 0xd3, 0x30,
	0x14, 0x66, 0xab, 0xfb, 0xe3, 0x51, 0x54, 0x32, 0x18, 0xb5, 0x55, 0x19, 0x13, 0xc1, 0x0b, 0xc9,
	0x44, 0x2f, 0xa6, 0x82, 0x37, 0x82, 0x48, 0xd1, 0x8d, 0x11, 0x64, 0xe0, 0x6e, 0x4a, 0xb5, 0xb1,
	0x04, 0xd6, 0x64, 0x26, 0xe9, 0x86, 0x6f, 0xe1, 0xad, 0x2f, 0xe3, 0xfb, 0xa8, 0x2f, 0x61, 0xda,
	0x74, 0x6e, 0x6d, 0xe7, 0xe4, 0xf7, 0xbb, 0xd9, 0x7a, 0x72, 0xce, 0xf7, 0x9d, 0x2f, 0x27, 0xdf,
	0x81, 0x97, 0xb1, 0xd2, 0x42, 0xd2, 0x09, 0xe5, 0x09, 0xe3, 0x74, 0xb2, 0x91, 0xe2, 0x13, 0x8d,
	0x33, 0x49, 0xd5, 0x24, 0x65, 0x61, 0xc0, 0x15, 0x95, 0x7a, 0xc9, 0x14, 0x33, 0x35, 0x01, 0xff,
	0x2c, 0x64, 0x1a, 0x69, 0x26, 0x78, 0xb8, 0xc8, 0xb0, 0xa9, 0xd4, 0x02, 0x3d, 0xb2, 0x70, 0x6c,
	0xe1, 0xf8, 0x3c, 0xc6, 0x1b, 0x94, 0xcd, 0xb6, 0xd1, 0x3a, 0xa3, 0xca, 0x52, 0x78, 0xb7, 0xab,
	0x0a, 0xa8, 0x94, 0x42, 0x96, 0x29, 0xbf, 0x9a, 0x4a, 0xa9, 0x52, 0x51, 0x42, 0xcb, 0xe4, 0xfd,
	0x7a, 0x52, 0x47, 0xec, 0xd0, 0xce, 0x16, 0x8d, 0xbf, 0x3b, 0x00, 0x8b, 0x48, 0x46, 0x26, 0x4b,
	0xa5, 0x42, 0x53, 0xb8, 0x9a, 0x71, 0xf6, 0x25, 0xa3, 0x21, 0x8b, 0xdd, 0xd6, 0xa8, 0xf5, 0xf0,
	0xda, 0x13, 0x0f, 0x97, 0x57, 0x28, 0x45, 0x29, 0x2d, 0x19, 0x4f, 0x96, 0x79, 0x40, 0xfa, 0xb6,
	0x38, 0x88, 0xd1, 0x03, 0xb8, 0xf1, 0x17, 0x18, 0xf2, 0x6c, 0xbd, 0x76, 0x7f, 0xf6, 0x0c, 0xbc,
	0x4f, 0xae, 0xef, 0x4b, 0xe6, 0xe6, 0x10, 0xad, 0xc0, 0xdf, 0xda, 0x9b, 0x87, 0x47, 0x5a, 0x42,
	0xfd, 0x75, 0x53, 0x74, 0x6c, 0x17, 0x1d, 0xfd, 0x5a, 0x47, 0xc6, 0x35, 0x4d, 0xa8, 0xb4, 0x2d,
	0xdd, 0x6d, 0x63, 0x72, 0xef, 0x0d, 0xda, 0x48, 0x78, 0x03, 0xa3, 0x33, 0xdc, 0x56, 0xd4, 0x2f,
	0x2b, 0xea, 0xce, 0xbf, 0x48, 0x0a, 0x91, 0x6f, 0x61, 0x70, 0x82, 0xc8, 0x75, 0xfe, 0x3b, 0x0e,
	0xd4, 0xa4, 0x45, 0xcf, 0xc1, 0x3d, 0xa5, 0xaa, 0x50, 0xf3, 0xdb, 0xaa, 0x19, 0x36, 0x61, 0xb9,
	0x8e, 0xf1, 0x8f, 0x36, 0xf4, 0x09, 0x55, 0x1b, 0x61, 0x1c, 0x83, 0x1e, 0x43, 0xa7, 0x78, 0xf9,
	0xfa, 0xab, 0x94, 0xc6, 0xb2, 0xae, 0x78, 0x9d, 0xff, 0x12, 0x5b, 0x88, 0x3e, 0xc0, 0xad, 0xfc,
	0xcd, 0x2b, 0x77, 0x68, 0x8f, 0x1c, 0x03, 0xc6, 0x35, 0x70, 0xdd, 0x1a, 0x33, 0x13, 0x1f, 0x89,
	0x21, 0x37, 0xd3, 0xea, 0x01, 0x7a, 0x06, 0xbd, 0xd2, 0x6b, 0x66, 0x2a, 0x39, 0xe3, 0xbd, 0x06,
	0xa3, 0x75, 0xe2, 0xcc, 0xfe, 0x93, 0x7d, 0x39, 0x7a, 0x07, 0x8e, 0x14, 0x3b, 0xf7, 0x4a, 0x81,
	0x7a, 0x81, 0x2f, 0xb2, 0x1d, 0x78, 0x3f, 0x0b, 0x4c, 0xc4, 0x8e, 0xe4, 0x34, 0xde, 0x5d, 0x70,
	0xcc, 0x37, 0x1a, 0x42, 0xd7, 0x44, 0xb9, 0x81, 0xbe, 0xcd, 0xcd, 0x74, 0x3a, 0xa4, 0x63, 0xc2,
	0x20, 0x7e, 0xb5, 0x04, 0x9f, 0x89, 0x5a, 0x8f, 0xc3, 0x02, 0xaf, 0xa6, 0x97, 0x5c, 0xed, 0x8f,
	0xdd, 0x62, 0x77, 0x9e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x90, 0x39, 0xbf, 0x05, 0x1c, 0x04,
	0x00, 0x00,
}