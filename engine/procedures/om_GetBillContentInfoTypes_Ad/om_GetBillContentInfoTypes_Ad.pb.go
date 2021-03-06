// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetBillContentInfoTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetBillContentInfoTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetBillContentInfoTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetBillContentInfoTypes_Ad

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
	RowId                        int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	FieldTypeId                  *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	InformationType              *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	BillContentInformationTypeId *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=bill_content_information_type_id,json=billContentInformationTypeId" json:"bill_content_information_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

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

func (m *Response_Row) GetBillContentInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BillContentInformationTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetBillContentInfoTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetBillContentInfoTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetBillContentInfoTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetBillContentInfoTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x25, 0xdd, 0xa6, 0x2d, 0x53, 0xa4, 0x75, 0x8b, 0xb2, 0x6e, 0x44, 0x42, 0x7d, 0x11, 0x84,
	0x89, 0xe8, 0x8b, 0x82, 0xa0, 0x56, 0x54, 0x82, 0x74, 0x91, 0x41, 0x04, 0x7d, 0x70, 0xd8, 0x74,
	0x6e, 0x97, 0xc1, 0xdd, 0x99, 0x30, 0x33, 0xb1, 0xf8, 0x11, 0x82, 0xda, 0xff, 0xf0, 0xbf, 0xfc,
	0x0b, 0xef, 0xee, 0x4c, 0x4c, 0x76, 0x13, 0xd4, 0xbc, 0x24, 0xdc, 0xbd, 0xf7, 0x9c, 0xb9, 0xe7,
	0x9c, 0x19, 0xf2, 0x58, 0x58, 0xa7, 0x0d, 0x8c, 0x40, 0x15, 0x52, 0xc1, 0x68, 0x6a, 0xf4, 0x19,
	0x88, 0x99, 0x01, 0x3b, 0xd2, 0x15, 0x7f, 0x05, 0xee, 0x44, 0x96, 0xe5, 0x73, 0xad, 0x1c, 0x28,
	0x37, 0x56, 0xe7, 0xfa, 0xed, 0x97, 0x29, 0x58, 0xfe, 0x4c, 0x50, 0x1c, 0x74, 0x3a, 0xbe, 0xeb,
	0xd1, 0xd4, 0xa3, 0xe9, 0x5f, 0x21, 0xe9, 0x51, 0x38, 0xea, 0x73, 0x5e, 0xce, 0xc0, 0x7a, 0x86,
	0xf4, 0x46, 0xfb, 0x7c, 0x30, 0x46, 0x9b, 0xd0, 0x1a, 0xb4, 0x5b, 0x15, 0x58, 0x9b, 0x17, 0x10,
	0x9a, 0xb7, 0xbb, 0x4d, 0x97, 0x4b, 0x3c, 0xcc, 0x54, 0xb9, 0x93, 0x5a, 0xf9, 0xa1, 0xe3, 0xcb,
	0x1e, 0x21, 0x6f, 0x72, 0x93, 0x63, 0x17, 0x8c, 0x8d, 0x5f, 0x93, 0xa3, 0xa5, 0x19, 0xee, 0x70,
	0x31, 0x2e, 0x45, 0xd2, 0x1b, 0xf6, 0xee, 0xec, 0xdf, 0x1f, 0xd0, 0xa0, 0x25, 0xac, 0x27, 0x51,
	0x40, 0x01, 0xe6, 0x5d, 0x5d, 0xb1, 0xab, 0x4b, 0xb8, 0x5a, 0xcf, 0x58, 0xc4, 0x0f, 0x49, 0xb2,
	0x86, 0x8c, 0xab, 0x59, 0x59, 0x26, 0xbf, 0x76, 0x91, 0x72, 0x8f, 0x5d, 0x5b, 0x41, 0x65, 0xd8,
	0x3d, 0xfe, 0xb9, 0x4d, 0xf6, 0x18, 0xd8, 0xa9, 0x56, 0x16, 0xe2, 0x7b, 0xa4, 0xdf, 0x68, 0x0e,
	0x5b, 0xa4, 0xb4, 0xed, 0xa8, 0xf7, 0xe3, 0x45, 0xfd, 0xcb, 0xfc, 0x60, 0xfc, 0x9e, 0x1c, 0xd6,
	0x6a, 0xf9, 0x12, 0x79, 0xb2, 0x35, 0x8c, 0x10, 0x4c, 0x3b, 0xe0, 0xae, 0x29, 0xa7, 0x58, 0x8f,
	0x17, 0x35, 0x3b, 0xa8, 0xda, 0x1f, 0x50, 0xd3, 0x6e, 0x70, 0x39, 0x89, 0x1a, 0xc6, 0x5b, 0x2b,
	0x8c, 0x3e, 0x83, 0x53, 0xff, 0xcf, 0xe6, 0xe3, 0x68, 0x6d, 0x64, 0xf4, 0x45, 0xb2, 0xdd, 0xa0,
	0x1e, 0xd1, 0x0d, 0xae, 0x05, 0x9d, 0x5b, 0x41, 0x99, 0xbe, 0x60, 0x35, 0x4b, 0xfa, 0x75, 0x8b,
	0x44, 0x58, 0xc4, 0xd7, 0xc9, 0x0e, 0x96, 0x75, 0x44, 0xdf, 0x32, 0x74, 0xa7, 0xcf, 0xfa, 0x58,
	0xa2, 0xf5, 0x4f, 0xc9, 0x95, 0x73, 0x09, 0xa5, 0xf8, 0x93, 0xe0, 0xf7, 0xec, 0xdf, 0x11, 0xee,
	0x37, 0x90, 0x10, 0xde, 0x4b, 0x72, 0xd8, 0x0d, 0x2f, 0xf9, 0x91, 0xb5, 0x13, 0x08, 0x24, 0xd6,
	0x19, 0xa9, 0x0a, 0xcf, 0x71, 0xd0, 0x09, 0x34, 0x16, 0x64, 0x38, 0x41, 0x59, 0xfc, 0xcc, 0xeb,
	0xe2, 0xeb, 0xae, 0xd7, 0xe5, 0x7f, 0x2c, 0x77, 0x73, 0xd2, 0x36, 0xa7, 0x75, 0x69, 0x4e, 0x3e,
	0x92, 0x81, 0xd4, 0x1d, 0x4f, 0x17, 0x0f, 0xf5, 0xc3, 0x93, 0x42, 0x5b, 0xf1, 0x69, 0xde, 0x17,
	0x1b, 0xbf, 0xe5, 0xc9, 0x4e, 0xf3, 0x5a, 0x1e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xbe,
	0x03, 0x02, 0x0c, 0x04, 0x00, 0x00,
}
