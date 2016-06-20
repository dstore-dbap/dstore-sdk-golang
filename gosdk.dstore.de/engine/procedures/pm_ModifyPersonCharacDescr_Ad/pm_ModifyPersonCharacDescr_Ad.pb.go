// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonCharacDescr_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonCharacDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonCharacDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonCharacDescr_Ad

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
	PersonCharacteristicId        *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	PersonCharacteristicIdNull    bool                        `protobuf:"varint,1001,opt,name=person_characteristic_id_null,json=personCharacteristicIdNull" json:"person_characteristic_id_null,omitempty"`
	LanguageId                    *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	CharacteristicDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	CharacteristicDescriptionNull bool                        `protobuf:"varint,1003,opt,name=characteristic_description_null,json=characteristicDescriptionNull" json:"characteristic_description_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonCharacDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonCharacDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonCharacDescr_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyPersonCharacDescr_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xa6, 0x8d, 0xdb, 0x96, 0x53, 0xd0, 0x32, 0x42, 0x49, 0xb3, 0xac, 0x4a, 0xbd, 0x51, 0x84,
	0x59, 0xd1, 0x1b, 0x85, 0x82, 0x58, 0x2b, 0xb2, 0xc8, 0x96, 0x32, 0xa0, 0xa0, 0x17, 0x86, 0x98,
	0x99, 0xc6, 0xc1, 0xdd, 0x99, 0x30, 0x93, 0x58, 0x7c, 0x0b, 0x1f, 0xc6, 0x5b, 0x1f, 0x48, 0x7d,
	0x09, 0xcf, 0x64, 0x26, 0xdd, 0x4d, 0x74, 0x2d, 0x7b, 0xb3, 0x9b, 0x93, 0x73, 0xbe, 0x9f, 0x39,
	0xf9, 0x06, 0x8e, 0xb8, 0xad, 0xb4, 0x11, 0x63, 0xa1, 0x0a, 0xa9, 0xc4, 0xb8, 0x34, 0x3a, 0x17,
	0xbc, 0x36, 0xc2, 0x8e, 0xcb, 0x79, 0x3a, 0xd5, 0x5c, 0x9e, 0x7f, 0x3d, 0x13, 0xc6, 0x6a, 0xf5,
	0xe2, 0x53, 0x66, 0xb2, 0xfc, 0x44, 0xd8, 0xdc, 0xa4, 0xcf, 0x39, 0xc5, 0xc1, 0x4a, 0x93, 0x07,
	0x1e, 0x4d, 0x3d, 0x9a, 0xfe, 0x17, 0x92, 0xdc, 0x0c, 0x52, 0x5f, 0xb2, 0x59, 0x2d, 0xac, 0x67,
	0x48, 0x0e, 0xba, 0xfa, 0xc2, 0x18, 0x6d, 0x42, 0x6b, 0xd8, 0x6d, 0xcd, 0x85, 0xb5, 0x59, 0x21,
	0x42, 0xf3, 0x6e, 0xbf, 0x59, 0x65, 0x52, 0x9d, 0x6b, 0x33, 0xcf, 0x2a, 0xa9, 0x95, 0x1f, 0x3a,
	0xfc, 0x1e, 0x01, 0x9c, 0xa1, 0x07, 0xec, 0xa2, 0x21, 0xf2, 0x06, 0xe2, 0xb2, 0x31, 0x96, 0xe6,
	0x8d, 0x33, 0x7c, 0x29, 0x6d, 0x25, 0xf3, 0x54, 0xf2, 0x78, 0xe3, 0xce, 0xc6, 0xbd, 0xdd, 0x47,
	0x43, 0x1a, 0x0e, 0x14, 0x3c, 0x4a, 0x55, 0x89, 0x42, 0x98, 0xb7, 0xae, 0x62, 0xfb, 0xe5, 0xd2,
	0xa9, 0x5a, 0xec, 0x84, 0x93, 0x63, 0x18, 0xad, 0xa2, 0x4d, 0x55, 0x3d, 0x9b, 0xc5, 0x3f, 0xb7,
	0x91, 0x7c, 0x87, 0x25, 0xff, 0xc6, 0x9f, 0xe2, 0x08, 0x39, 0x82, 0xdd, 0x59, 0xa6, 0x8a, 0x1a,
	0x0f, 0xe8, 0xdc, 0x6c, 0x5e, 0xed, 0x06, 0xda, 0x79, 0x74, 0x70, 0x1f, 0xf6, 0x96, 0xd0, 0x5e,
	0xf4, 0x97, 0x17, 0xbd, 0xbe, 0x18, 0x6b, 0x84, 0xde, 0x41, 0xd2, 0x73, 0xc9, 0xdd, 0xf7, 0x91,
	0xa5, 0x5b, 0x5b, 0x1c, 0x35, 0xba, 0x49, 0x4f, 0xd7, 0x56, 0x46, 0xaa, 0xc2, 0xcb, 0x1e, 0x74,
	0xd1, 0x27, 0x0b, 0x30, 0x79, 0x05, 0xb7, 0x57, 0x53, 0x7b, 0x53, 0xbf, 0xbd, 0xa9, 0xd1, 0x4a,
	0x12, 0xe7, 0xf1, 0xf0, 0xc7, 0x26, 0xec, 0x30, 0x61, 0x4b, 0xad, 0xac, 0x20, 0x0f, 0x61, 0xd0,
	0x84, 0x22, 0x7c, 0xa1, 0x4b, 0x6f, 0x21, 0x72, 0x3e, 0x30, 0x2f, 0xdd, 0x2f, 0xf3, 0x83, 0x78,
	0xc4, 0x3d, 0x17, 0x87, 0x74, 0x29, 0x0f, 0xb8, 0xd0, 0x08, 0xc1, 0xb4, 0x07, 0xee, 0xa7, 0x66,
	0x8a, 0xf5, 0x64, 0x51, 0xb3, 0x1b, 0xf3, 0xee, 0x0b, 0xf2, 0x04, 0xb6, 0x43, 0x0c, 0x71, 0x55,
	0x8e, 0xf1, 0xd6, 0x5f, 0x8c, 0x3e, 0xa4, 0x53, 0xff, 0xcf, 0xda, 0x71, 0xf2, 0x1a, 0x22, 0xa3,
	0x2f, 0xe2, 0x6b, 0x0d, 0xea, 0x29, 0x5d, 0xe3, 0xde, 0xd0, 0x76, 0x15, 0x94, 0xe9, 0x0b, 0xe6,
	0x58, 0x92, 0x11, 0x44, 0xf8, 0x4c, 0xf6, 0x61, 0x0b, 0x2b, 0x97, 0x97, 0x6f, 0xa7, 0xb8, 0x9c,
	0x01, 0x1b, 0x60, 0x39, 0xe1, 0xc7, 0x1f, 0x60, 0x28, 0x75, 0x5f, 0xe2, 0xf2, 0x62, 0xbf, 0x7f,
	0x56, 0x68, 0xcb, 0x3f, 0xb7, 0x7d, 0xbe, 0xf6, 0xdd, 0xff, 0xb8, 0xd5, 0xdc, 0xae, 0xc7, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x78, 0xee, 0xfd, 0x3c, 0x04, 0x00, 0x00,
}