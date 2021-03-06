// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetTrolleySurchInfoTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetTrolleySurchInfoTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetTrolleySurchInfoTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetTrolleySurchInfoTypes_Ad

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
	InformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull bool                        `protobuf:"varint,1001,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	InformationTypeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	FieldTypeId       *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	InformationType   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Response_Row) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
	}
	return nil
}

func (m *Response_Row) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetTrolleySurchInfoTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetTrolleySurchInfoTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetTrolleySurchInfoTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetTrolleySurchInfoTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0x66, 0xbb, 0xbb, 0xcc, 0x22, 0xbb, 0x66, 0x51, 0x62, 0x0a, 0xb2, 0xd4, 0x17, 0x1f,
	0x64, 0x2a, 0xfa, 0x22, 0x82, 0x50, 0x05, 0x95, 0x62, 0x1b, 0x64, 0x2c, 0x82, 0xbe, 0x84, 0xd8,
	0xb9, 0x8d, 0x83, 0xc9, 0x4c, 0x99, 0x99, 0x58, 0xfa, 0x17, 0x6a, 0x1f, 0xfc, 0x46, 0x3f, 0xc0,
	0x77, 0x27, 0x99, 0xa9, 0x6d, 0xd2, 0x22, 0xd6, 0x97, 0x84, 0x9b, 0x73, 0xcf, 0x99, 0x73, 0xcf,
	0xcd, 0xa0, 0xa7, 0x54, 0x69, 0x21, 0xa1, 0x0f, 0x3c, 0x63, 0x1c, 0xfa, 0x73, 0x29, 0xa6, 0x40,
	0x4b, 0x09, 0xaa, 0x2f, 0x8a, 0xe4, 0x15, 0xe8, 0x89, 0x14, 0x79, 0x0e, 0xcb, 0xb7, 0xa5, 0x9c,
	0x7e, 0x1a, 0xf2, 0x99, 0x98, 0x2c, 0xe7, 0xa0, 0x92, 0x67, 0x14, 0x9b, 0x4e, 0x2d, 0x82, 0xfb,
	0x96, 0x8e, 0x2d, 0x1d, 0xff, 0x9d, 0x13, 0x5d, 0xba, 0xc3, 0xbe, 0xa4, 0x79, 0x09, 0xca, 0x4a,
	0x44, 0xb7, 0x9b, 0x0e, 0x40, 0x4a, 0x21, 0x1d, 0xd4, 0x6d, 0x42, 0x05, 0x28, 0x95, 0x66, 0xe0,
	0xc0, 0xbb, 0x6d, 0x50, 0xa7, 0xcc, 0x1c, 0x26, 0x8b, 0x54, 0x33, 0xc1, 0x6d, 0x53, 0x6f, 0xe5,
	0x21, 0xf4, 0x26, 0x95, 0xa9, 0x41, 0x41, 0xaa, 0xe0, 0x35, 0xba, 0xdc, 0xea, 0x49, 0xb4, 0x31,
	0x96, 0x30, 0x1a, 0x7a, 0x57, 0xde, 0xbd, 0xb3, 0x87, 0x5d, 0xec, 0x86, 0x71, 0xf6, 0x18, 0xd7,
	0x90, 0x81, 0x7c, 0x57, 0x55, 0xe4, 0xc6, 0x16, 0xaf, 0x9a, 0x67, 0x48, 0x83, 0xc7, 0x28, 0xdc,
	0x23, 0x96, 0xf0, 0x32, 0xcf, 0xc3, 0x9f, 0x27, 0x46, 0xf2, 0x94, 0xdc, 0xdc, 0x61, 0xc5, 0x06,
	0xed, 0xfd, 0x38, 0x42, 0xa7, 0x04, 0xd4, 0x5c, 0x70, 0x05, 0xc1, 0x03, 0xd4, 0xa9, 0x67, 0x76,
	0x2e, 0x22, 0xdc, 0x8c, 0xd4, 0xe6, 0xf1, 0xa2, 0x7a, 0x12, 0xdb, 0x18, 0xbc, 0x47, 0x17, 0xd5,
	0xb4, 0xc9, 0x96, 0x78, 0x78, 0xed, 0xca, 0x37, 0x64, 0xdc, 0x22, 0xb7, 0x43, 0x19, 0x9b, 0x7a,
	0xb8, 0xa9, 0xc9, 0x79, 0xd1, 0xfc, 0x60, 0x66, 0x3a, 0x71, 0x29, 0x87, 0x7e, 0xad, 0x78, 0x67,
	0x47, 0xd1, 0xee, 0x60, 0x6c, 0xdf, 0x64, 0xdd, 0x1e, 0x8c, 0x90, 0x2f, 0xc5, 0x22, 0x3c, 0xaa,
	0x59, 0x4f, 0xf0, 0x21, 0xff, 0x05, 0x5e, 0x67, 0x81, 0x89, 0x58, 0x90, 0x4a, 0x26, 0xfa, 0xe5,
	0x21, 0xdf, 0x14, 0xc1, 0x2d, 0x74, 0x6c, 0xca, 0x6a, 0x47, 0x5f, 0x63, 0x13, 0x4f, 0x87, 0x74,
	0x4c, 0x69, 0xb2, 0x1f, 0xed, 0x5f, 0xe4, 0xb7, 0xf8, 0xbf, 0x36, 0x39, 0x40, 0xd7, 0x67, 0x0c,
	0x72, 0xfa, 0x47, 0xe7, 0xfb, 0x3f, 0xe8, 0x9c, 0xd5, 0x14, 0xa7, 0xf0, 0x12, 0x5d, 0xb4, 0xfd,
	0x84, 0xab, 0xb8, 0xb9, 0x50, 0x27, 0xa2, 0xb4, 0x64, 0x3c, 0xb3, 0x1a, 0xe7, 0x2d, 0x2f, 0xcf,
	0x13, 0xd4, 0x65, 0xa2, 0x15, 0xde, 0xe6, 0x4e, 0x7e, 0x18, 0x64, 0x42, 0xd1, 0xcf, 0x6b, 0x9c,
	0x1e, 0x7e, 0x6d, 0x3f, 0x1e, 0xd7, 0xf7, 0xe2, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86,
	0x67, 0xd7, 0x17, 0xf8, 0x03, 0x00, 0x00,
}
