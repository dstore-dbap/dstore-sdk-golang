// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetLanguageDescriptions_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetLanguageDescriptions_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetLanguageDescriptions_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetLanguageDescriptions_Ad

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
	LanguageId                *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull            bool                        `protobuf:"varint,1001,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	TranslationLanguageId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=translation_language_id,json=translationLanguageId" json:"translation_language_id,omitempty"`
	TranslationLanguageIdNull bool                        `protobuf:"varint,1002,opt,name=translation_language_id_null,json=translationLanguageIdNull" json:"translation_language_id_null,omitempty"`
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

func (m *Parameters) GetTranslationLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TranslationLanguageId
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
	RowId                   int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TranslationLanguageId   *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=translation_language_id,json=translationLanguageId" json:"translation_language_id,omitempty"`
	TranslationLanguageName *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=translation_language_name,json=translationLanguageName" json:"translation_language_name,omitempty"`
	LanguageName            *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=language_name,json=languageName" json:"language_name,omitempty"`
	TranslatedLanguageName  *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=translated_language_name,json=translatedLanguageName" json:"translated_language_name,omitempty"`
	LanguageId              *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTranslationLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TranslationLanguageId
	}
	return nil
}

func (m *Response_Row) GetTranslationLanguageName() *dstore_values.StringValue {
	if m != nil {
		return m.TranslationLanguageName
	}
	return nil
}

func (m *Response_Row) GetLanguageName() *dstore_values.StringValue {
	if m != nil {
		return m.LanguageName
	}
	return nil
}

func (m *Response_Row) GetTranslatedLanguageName() *dstore_values.StringValue {
	if m != nil {
		return m.TranslatedLanguageName
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetLanguageDescriptions_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetLanguageDescriptions_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetLanguageDescriptions_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetLanguageDescriptions_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd5, 0x9a, 0xb4, 0xd5, 0x94, 0x43, 0x65, 0x44, 0xeb, 0x24, 0x08, 0xa1, 0x72, 0x03,
	0x42, 0xda, 0x20, 0xb8, 0x01, 0xa9, 0x88, 0x83, 0x40, 0xa8, 0xa2, 0xb5, 0xd0, 0x72, 0x12, 0x5c,
	0x60, 0x99, 0x7a, 0xb0, 0x56, 0x38, 0xbb, 0xd1, 0xae, 0x43, 0x5f, 0x82, 0x0b, 0x8e, 0x6f, 0xc0,
	0xab, 0xf0, 0x30, 0xf0, 0x14, 0x8c, 0xbd, 0x6b, 0xf9, 0x40, 0x42, 0x14, 0x6e, 0x12, 0x8d, 0x67,
	0xfe, 0x6f, 0xc6, 0xf3, 0xaf, 0x17, 0xf6, 0x12, 0x93, 0x2b, 0x8d, 0x23, 0x94, 0xa9, 0x90, 0x38,
	0x9a, 0x68, 0x75, 0x84, 0xc9, 0x54, 0xa3, 0x19, 0x8d, 0x45, 0xf4, 0x08, 0xf3, 0x83, 0x58, 0xa6,
	0xd3, 0x38, 0xc5, 0x07, 0x68, 0x8e, 0xb4, 0x98, 0xe4, 0x42, 0x49, 0x13, 0xdd, 0x4b, 0x18, 0x15,
	0xe6, 0xca, 0xbf, 0x6a, 0xd5, 0xcc, 0xaa, 0xd9, 0x3f, 0x25, 0x83, 0xb3, 0xae, 0xd5, 0x87, 0x38,
	0x9b, 0xa2, 0xb1, 0x84, 0x41, 0xbf, 0xdd, 0x1f, 0xb5, 0x56, 0xda, 0xa5, 0x86, 0xed, 0xd4, 0x18,
	0x8d, 0x21, 0xaa, 0x4b, 0x5e, 0xea, 0x26, 0xf3, 0x58, 0xc8, 0x77, 0x4a, 0x8f, 0xe3, 0xa2, 0x9f,
	0x2d, 0xda, 0xfd, 0xb8, 0x0a, 0xf0, 0x24, 0xd6, 0x31, 0x65, 0x51, 0x1b, 0x7f, 0x0f, 0x36, 0x33,
	0x37, 0x5b, 0x24, 0x92, 0x60, 0xe5, 0xe2, 0xca, 0xe5, 0xcd, 0xeb, 0x43, 0xe6, 0xde, 0xc1, 0x8d,
	0x25, 0x64, 0x8e, 0x29, 0xea, 0x17, 0x45, 0xc4, 0xa1, 0xaa, 0xdf, 0x4f, 0xfc, 0x2b, 0xb0, 0xd5,
	0x50, 0x47, 0x72, 0x9a, 0x65, 0xc1, 0xaf, 0x75, 0x62, 0x6c, 0xf0, 0xd3, 0x75, 0x59, 0x48, 0x8f,
	0xfd, 0xa7, 0xb0, 0x93, 0xeb, 0x58, 0x9a, 0xac, 0x1c, 0x26, 0x6a, 0x36, 0x5d, 0x5d, 0xdc, 0xf4,
	0x5c, 0x43, 0x7b, 0x50, 0xf7, 0xbf, 0x0b, 0xe7, 0xe7, 0x40, 0xed, 0x2c, 0xbf, 0xed, 0x2c, 0xfd,
	0x99, 0xea, 0x62, 0xac, 0xdd, 0x9f, 0x3d, 0xd8, 0xe0, 0x68, 0x26, 0xe4, 0x07, 0xfa, 0xd7, 0xa0,
	0x57, 0x2e, 0xdb, 0xad, 0x61, 0xc0, 0xda, 0x56, 0x5a, 0x23, 0x1e, 0x16, 0xbf, 0xdc, 0x16, 0xfa,
	0xaf, 0x60, 0xab, 0x58, 0x73, 0xd4, 0xd8, 0x33, 0xbd, 0x8e, 0x47, 0x62, 0xd6, 0x11, 0x77, 0xdd,
	0x38, 0xa4, 0x78, 0xbf, 0x8e, 0xf9, 0x99, 0x71, 0xfb, 0x81, 0x7f, 0x13, 0xd6, 0x9d, 0xbd, 0x81,
	0x57, 0x12, 0x2f, 0xfc, 0x45, 0xb4, 0xe6, 0x1f, 0xda, 0x7f, 0x5e, 0x95, 0xfb, 0x8f, 0xc1, 0xd3,
	0xea, 0x38, 0x38, 0x51, 0xaa, 0x6e, 0xb1, 0x25, 0xce, 0x23, 0xab, 0x56, 0xc1, 0xb8, 0x3a, 0xe6,
	0x05, 0x65, 0xf0, 0xc3, 0x03, 0x8f, 0x02, 0x7f, 0x1b, 0xd6, 0x28, 0x2c, 0xec, 0xfa, 0x14, 0xd2,
	0x76, 0x7a, 0xbc, 0x47, 0x21, 0x59, 0xf0, 0x6c, 0xbe, 0xaf, 0x9f, 0xc3, 0xff, 0x36, 0xf6, 0x25,
	0xf4, 0x67, 0x52, 0x25, 0x9d, 0xdb, 0xe0, 0x4b, 0xd8, 0xb6, 0xc7, 0x71, 0x4d, 0xae, 0x85, 0x4c,
	0x2d, 0x76, 0x67, 0x06, 0x36, 0x24, 0x2d, 0x9d, 0x98, 0x53, 0x6d, 0xd8, 0xd7, 0xc5, 0xb0, 0x93,
	0x59, 0x93, 0xf0, 0x1c, 0x82, 0x0a, 0x8e, 0x49, 0x67, 0xb2, 0x6f, 0x8b, 0x61, 0xdb, 0xb5, 0xb8,
	0x35, 0xd8, 0xed, 0xf6, 0x87, 0xf8, 0x3d, 0x5c, 0xea, 0x4b, 0xbc, 0xff, 0x06, 0x86, 0x42, 0x75,
	0xac, 0xae, 0x2f, 0xae, 0xd7, 0x77, 0x52, 0x65, 0x92, 0xf7, 0x55, 0x3e, 0x59, 0xfa, 0x6e, 0x7b,
	0xbb, 0x56, 0xde, 0x1e, 0x37, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x48, 0x67, 0x59, 0x8b, 0x1c,
	0x05, 0x00, 0x00,
}