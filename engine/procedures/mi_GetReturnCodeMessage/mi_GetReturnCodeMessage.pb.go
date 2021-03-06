// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetReturnCodeMessage.proto
// DO NOT EDIT!

/*
Package mi_GetReturnCodeMessage is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetReturnCodeMessage.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetReturnCodeMessage

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
	ReturnCode     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=return_code,json=returnCode" json:"return_code,omitempty"`
	ReturnCodeNull bool                        `protobuf:"varint,1001,opt,name=return_code_null,json=returnCodeNull" json:"return_code_null,omitempty"`
	LanguageId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetReturnCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
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
	RowId               int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ReturnCode          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=return_code,json=returnCode" json:"return_code,omitempty"`
	Description         *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=description" json:"description,omitempty"`
	DetailedDescription *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=detailed_description,json=detailedDescription" json:"detailed_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetReturnCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetDetailedDescription() *dstore_values.StringValue {
	if m != nil {
		return m.DetailedDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetReturnCodeMessage.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetReturnCodeMessage.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetReturnCodeMessage.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetReturnCodeMessage.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6b, 0x15, 0x31,
	0x10, 0xe5, 0x7a, 0x7b, 0xdb, 0x92, 0x82, 0x96, 0x54, 0x64, 0xdd, 0x0b, 0x22, 0x15, 0x41, 0x5f,
	0x72, 0xc5, 0x82, 0x28, 0xe8, 0x8b, 0x5a, 0xa4, 0x0f, 0x5d, 0x24, 0x0f, 0xa2, 0xbe, 0x2c, 0xeb,
	0xcd, 0xb8, 0x04, 0x77, 0x93, 0xcb, 0x24, 0x6b, 0xff, 0x81, 0xcf, 0x7e, 0xfc, 0x3f, 0xdf, 0xd5,
	0x3f, 0x61, 0xb2, 0xc9, 0xba, 0x1f, 0x22, 0xb6, 0x2f, 0x0b, 0xb3, 0x73, 0xce, 0x9c, 0xc9, 0x99,
	0x19, 0xf2, 0x40, 0x18, 0xab, 0x11, 0x56, 0xa0, 0x4a, 0xa9, 0x60, 0xb5, 0x41, 0xbd, 0x06, 0xd1,
	0x20, 0x98, 0x55, 0x2d, 0xf3, 0x17, 0x60, 0x39, 0xd8, 0x06, 0xd5, 0x33, 0x2d, 0xe0, 0x14, 0x8c,
	0x29, 0x4a, 0x60, 0x0e, 0x62, 0x35, 0xbd, 0x1d, 0x78, 0x2c, 0xf0, 0xd8, 0x3f, 0xc0, 0xe9, 0x41,
	0x2c, 0xff, 0xb1, 0xa8, 0x1a, 0x30, 0x81, 0x9b, 0x5e, 0x1f, 0x6b, 0x02, 0xa2, 0xc6, 0x98, 0x5a,
	0x8e, 0x53, 0xf5, 0x50, 0x33, 0xbd, 0x35, 0x4d, 0xda, 0x42, 0xaa, 0xf7, 0x1a, 0xeb, 0xc2, 0x4a,
	0xad, 0x02, 0xe8, 0xf0, 0xfb, 0x8c, 0x90, 0x97, 0x05, 0x16, 0x2e, 0x0b, 0x68, 0xe8, 0x63, 0xb2,
	0x87, 0x6d, 0x57, 0xf9, 0xda, 0xb5, 0x95, 0xcc, 0x6e, 0xce, 0xee, 0xec, 0xdd, 0x5f, 0xb2, 0xd8,
	0x7d, 0x6c, 0x4b, 0x2a, 0x0b, 0x25, 0xe0, 0x2b, 0x1f, 0x71, 0x82, 0x7f, 0x5e, 0x41, 0xef, 0x92,
	0xfd, 0x01, 0x3b, 0x57, 0x4d, 0x55, 0x25, 0x3f, 0x76, 0x5c, 0x8d, 0x5d, 0x7e, 0xb9, 0x87, 0x65,
	0xee, 0xb7, 0x17, 0xaa, 0x0a, 0x55, 0x36, 0xae, 0xdd, 0x5c, 0x8a, 0xe4, 0xd2, 0x39, 0x84, 0x3a,
	0xfc, 0x89, 0xf0, 0x42, 0x03, 0x76, 0x10, 0xfa, 0x19, 0x85, 0x7a, 0x98, 0x17, 0x3a, 0xfc, 0xb4,
	0x45, 0x76, 0x39, 0x98, 0x8d, 0x56, 0x06, 0xe8, 0x3d, 0xb2, 0x68, 0xed, 0x8b, 0x0f, 0x4b, 0xd9,
	0x78, 0x2c, 0xc1, 0xda, 0x63, 0xff, 0xe5, 0x01, 0x48, 0xdf, 0x90, 0x7d, 0x6f, 0x5c, 0x3e, 0x70,
	0xce, 0x35, 0x3b, 0x77, 0x64, 0x36, 0x21, 0x4f, 0xfd, 0x3d, 0x75, 0xf1, 0x49, 0x1f, 0xf3, 0x2b,
	0xf5, 0xf8, 0x07, 0x7d, 0x48, 0x76, 0xe2, 0xc0, 0x92, 0x79, 0x5b, 0xf1, 0xc6, 0x5f, 0x15, 0xc3,
	0x38, 0xe3, 0x76, 0xf0, 0x0e, 0x4e, 0x8f, 0xc9, 0x1c, 0xf5, 0x59, 0xb2, 0xd5, 0xb2, 0x8e, 0xd8,
	0xb9, 0x76, 0x8b, 0x75, 0x26, 0x30, 0xae, 0xcf, 0xb8, 0xe7, 0xa7, 0xbf, 0x66, 0x64, 0xee, 0x02,
	0x7a, 0x8d, 0x6c, 0xbb, 0xd0, 0x8f, 0xe1, 0x73, 0xe6, 0x7c, 0x59, 0xf0, 0x85, 0x0b, 0x9d, 0xcb,
	0x4f, 0xc6, 0xcb, 0xf0, 0x25, 0xbb, 0xd8, 0x36, 0x38, 0xba, 0x00, 0xb3, 0x46, 0xb9, 0x69, 0x5d,
	0xfb, 0x9a, 0x8d, 0x3d, 0x8f, 0x74, 0x63, 0x51, 0xaa, 0x32, 0xb0, 0x87, 0x78, 0x9a, 0x91, 0xab,
	0xc2, 0x5b, 0x5a, 0x81, 0xc8, 0x87, 0x75, 0xbe, 0xfd, 0xbf, 0xce, 0x41, 0x47, 0x7c, 0xde, 0xf3,
	0x9e, 0xbe, 0x26, 0x4b, 0xa9, 0x27, 0x5e, 0xf5, 0xf7, 0xfb, 0xf6, 0x51, 0xa9, 0x8d, 0xf8, 0xd0,
	0xe5, 0xc5, 0x05, 0x4e, 0xfc, 0xdd, 0x76, 0x7b, 0x4a, 0x47, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x7a, 0x4c, 0xfb, 0xd5, 0x1d, 0x04, 0x00, 0x00,
}
