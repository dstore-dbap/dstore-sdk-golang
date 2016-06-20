// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_SetProductDescriptions_Ad.proto
// DO NOT EDIT!

/*
Package im_SetProductDescriptions_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_SetProductDescriptions_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_SetProductDescriptions_Ad

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
	LanguageId                       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                   bool                        `protobuf:"varint,1001,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	SkipElementsWithProductDescr     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=skip_elements_with_product_descr,json=skipElementsWithProductDescr" json:"skip_elements_with_product_descr,omitempty"`
	SkipElementsWithProductDescrNull bool                        `protobuf:"varint,1002,opt,name=skip_elements_with_product_descr_null,json=skipElementsWithProductDescrNull" json:"skip_elements_with_product_descr_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetSkipElementsWithProductDescr() *dstore_values.BooleanValue {
	if m != nil {
		return m.SkipElementsWithProductDescr
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_SetProductDescriptions_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_SetProductDescriptions_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_SetProductDescriptions_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x65, 0x5b, 0xbb, 0xbb, 0xdc, 0x05, 0x5d, 0x22, 0xc8, 0xd8, 0xaa, 0x94, 0x15, 0x41, 0x7d,
	0x48, 0x45, 0x11, 0x16, 0xf4, 0x45, 0x71, 0x1f, 0x2a, 0xec, 0x52, 0xa2, 0x28, 0xfa, 0x12, 0xb2,
	0xcd, 0xb5, 0x06, 0x67, 0x92, 0x21, 0xc9, 0xd8, 0xdf, 0xf0, 0x43, 0xfc, 0x04, 0x7f, 0x48, 0xbf,
	0xc2, 0xa4, 0x49, 0x69, 0x67, 0x94, 0x55, 0xf6, 0xa5, 0x9d, 0x9b, 0x7b, 0xcf, 0x39, 0x37, 0xf7,
	0xdc, 0xc0, 0x33, 0xe9, 0xbc, 0xb1, 0x38, 0x41, 0xbd, 0x50, 0x1a, 0x27, 0xb5, 0x35, 0x73, 0x94,
	0x8d, 0x45, 0x37, 0x51, 0x15, 0x7f, 0x83, 0x7e, 0x66, 0x8d, 0x6c, 0xe6, 0xfe, 0x15, 0xba, 0xb9,
	0x55, 0xb5, 0x57, 0x46, 0x3b, 0xfe, 0x42, 0xd2, 0x50, 0xe7, 0x0d, 0x79, 0x98, 0xc0, 0x34, 0x81,
	0xe9, 0x45, 0x88, 0xe1, 0xf5, 0x2c, 0xf4, 0x55, 0x94, 0x0d, 0xba, 0x44, 0x30, 0xbc, 0xd9, 0x56,
	0x47, 0x6b, 0x8d, 0xcd, 0xa9, 0x51, 0x3b, 0x55, 0xa1, 0x73, 0x62, 0x81, 0x39, 0x79, 0xb7, 0x9b,
	0xf4, 0x42, 0xe9, 0x4f, 0xc6, 0x56, 0x22, 0xea, 0xa5, 0xa2, 0xa3, 0xef, 0x3d, 0x80, 0x99, 0xb0,
	0x22, 0x64, 0xd1, 0x3a, 0xf2, 0x1c, 0x0e, 0x4a, 0xa1, 0x17, 0x4d, 0x60, 0xe1, 0x4a, 0x16, 0x3b,
	0xe3, 0x9d, 0xfb, 0x07, 0x8f, 0x47, 0x34, 0x5f, 0x21, 0xb7, 0xa5, 0xb4, 0xc7, 0x05, 0xda, 0x77,
	0x31, 0x62, 0xb0, 0xae, 0x9f, 0x4a, 0xf2, 0x00, 0x0e, 0xb7, 0xd0, 0x5c, 0x37, 0x65, 0x59, 0xfc,
	0xdc, 0x0b, 0x1c, 0xfb, 0xec, 0xea, 0xa6, 0xec, 0x2c, 0x1c, 0x93, 0x39, 0x8c, 0xdd, 0x17, 0x55,
	0x73, 0x2c, 0xb1, 0x42, 0xed, 0x1d, 0x5f, 0x2a, 0xff, 0x99, 0xd7, 0x69, 0x2c, 0x5c, 0xc6, 0xb9,
	0x14, 0xbd, 0xbf, 0xaa, 0x9f, 0x1b, 0x53, 0xa2, 0xd0, 0x49, 0xfd, 0x56, 0x24, 0x39, 0xc9, 0x1c,
	0xef, 0x03, 0xc5, 0xf6, 0x60, 0xc9, 0x0c, 0xee, 0xfd, 0x4b, 0x24, 0x35, 0xf9, 0x2b, 0x35, 0x39,
	0xbe, 0x88, 0x2d, 0xb6, 0x7d, 0xf4, 0xa3, 0x07, 0xfb, 0x0c, 0x5d, 0x1d, 0xfc, 0x42, 0xf2, 0x08,
	0x06, 0x2b, 0x33, 0xf2, 0x98, 0x86, 0xb4, 0xed, 0x74, 0x32, 0xea, 0x24, 0xfe, 0xb2, 0x54, 0x48,
	0x3e, 0xc0, 0x61, 0xb4, 0x81, 0x6f, 0xf9, 0x10, 0x6e, 0xd9, 0x0f, 0x60, 0xda, 0x01, 0x77, 0xdd,
	0x3a, 0x0d, 0xf1, 0x74, 0x13, 0xb3, 0x6b, 0x55, 0xfb, 0x80, 0x1c, 0xc3, 0x5e, 0xb6, 0xbf, 0xe8,
	0xaf, 0x18, 0xef, 0xfc, 0xc1, 0x98, 0x96, 0xe3, 0x34, 0xfd, 0xb3, 0x75, 0x39, 0x79, 0x0d, 0x7d,
	0x6b, 0x96, 0xc5, 0x95, 0x15, 0xea, 0x98, 0xfe, 0xff, 0xba, 0xd2, 0xf5, 0x24, 0x28, 0x33, 0x4b,
	0x16, 0x49, 0x86, 0xb7, 0xa1, 0x1f, 0xbe, 0xc9, 0x0d, 0xd8, 0x0d, 0x51, 0xdc, 0xa0, 0x6f, 0x67,
	0x61, 0x36, 0x03, 0x36, 0x08, 0xe1, 0x54, 0xbe, 0x7c, 0x0b, 0x23, 0x65, 0x3a, 0x0a, 0x9b, 0xd7,
	0xf4, 0xf1, 0xe9, 0xa5, 0xde, 0xd9, 0xf9, 0xee, 0x6a, 0x95, 0x9f, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0x8c, 0x87, 0x75, 0xa7, 0x03, 0x00, 0x00,
}