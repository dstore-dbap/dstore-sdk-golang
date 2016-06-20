// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonTypeMetaInfo_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPersonTypeMetaInfo_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonTypeMetaInfo_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonTypeMetaInfo_Ad

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
	PersonTypeId                  *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull              bool                        `protobuf:"varint,1001,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	MetaInformationTypeIdList     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=meta_information_type_id_list,json=metaInformationTypeIdList" json:"meta_information_type_id_list,omitempty"`
	MetaInformationTypeIdListNull bool                        `protobuf:"varint,1002,opt,name=meta_information_type_id_list_null,json=metaInformationTypeIdListNull" json:"meta_information_type_id_list_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetMetaInformationTypeIdList() *dstore_values.StringValue {
	if m != nil {
		return m.MetaInformationTypeIdList
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	MetaInformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=meta_information_type_id,json=metaInformationTypeId" json:"meta_information_type_id,omitempty"`
	MetaInformationType       *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=meta_information_type,json=metaInformationType" json:"meta_information_type,omitempty"`
	MetaInformation           *dstore_values.DecimalValue `protobuf:"bytes,10003,opt,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	PersonTypeId              *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	TranslatedMetaInformation *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=translated_meta_information,json=translatedMetaInformation" json:"translated_meta_information,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetMetaInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.MetaInformationTypeId
	}
	return nil
}

func (m *Response_Row) GetMetaInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.MetaInformationType
	}
	return nil
}

func (m *Response_Row) GetMetaInformation() *dstore_values.DecimalValue {
	if m != nil {
		return m.MetaInformation
	}
	return nil
}

func (m *Response_Row) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Response_Row) GetTranslatedMetaInformation() *dstore_values.StringValue {
	if m != nil {
		return m.TranslatedMetaInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonTypeMetaInfo_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonTypeMetaInfo_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonTypeMetaInfo_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6a, 0x14, 0x31,
	0x14, 0xc7, 0x69, 0xc7, 0x6d, 0xcb, 0x51, 0xb4, 0xa4, 0x54, 0x66, 0x77, 0xa9, 0xc8, 0x7a, 0xa3,
	0x37, 0x59, 0x51, 0x41, 0xf1, 0xae, 0x05, 0x29, 0x2b, 0xee, 0x58, 0xe2, 0x07, 0xf8, 0x01, 0x43,
	0xec, 0x1c, 0x97, 0xc0, 0x4c, 0x32, 0x24, 0x59, 0x8b, 0x6f, 0xe1, 0xe7, 0x43, 0x78, 0xe5, 0x9b,
	0xf8, 0x1e, 0xfa, 0x14, 0x66, 0x26, 0xe9, 0xda, 0x9d, 0x9d, 0x76, 0xd5, 0x9b, 0x5d, 0xc2, 0x39,
	0xff, 0x7f, 0x7e, 0xc9, 0xff, 0x4c, 0xe0, 0x7e, 0x66, 0xac, 0xd2, 0x38, 0x44, 0x39, 0x11, 0x12,
	0x87, 0xa5, 0x56, 0x87, 0x98, 0x4d, 0x35, 0x9a, 0x61, 0x59, 0xa4, 0xfb, 0x68, 0x0f, 0x50, 0x1b,
	0x25, 0x9f, 0xbe, 0x2f, 0x71, 0x8c, 0x96, 0x8f, 0xe4, 0x5b, 0x95, 0xee, 0x66, 0xd4, 0xb5, 0x59,
	0x45, 0x6e, 0x78, 0x2d, 0xf5, 0x5a, 0x7a, 0x86, 0xa0, 0xb7, 0x15, 0xb6, 0x79, 0xc7, 0xf3, 0x29,
	0x1a, 0xaf, 0xef, 0x75, 0xe7, 0xf7, 0x46, 0xad, 0x95, 0x0e, 0xa5, 0xfe, 0x7c, 0xa9, 0x40, 0x63,
	0xf8, 0x04, 0x43, 0xf1, 0x5a, 0xb3, 0x68, 0xb9, 0x70, 0xfb, 0xe8, 0x82, 0x5b, 0xa1, 0xa4, 0x6f,
	0x1a, 0x7c, 0x5b, 0x05, 0x38, 0xe0, 0x9a, 0xbb, 0xaa, 0x23, 0x22, 0xbb, 0x70, 0xb1, 0xac, 0xc9,
	0x52, 0xeb, 0xd0, 0x52, 0x91, 0xc5, 0x2b, 0x57, 0x57, 0xae, 0x9f, 0xbf, 0xd5, 0xa7, 0xe1, 0x10,
	0x81, 0x4c, 0x48, 0x8b, 0x13, 0xd4, 0xcf, 0xab, 0x15, 0xbb, 0x50, 0xce, 0x0e, 0x33, 0xca, 0x08,
	0x85, 0xad, 0x79, 0x8b, 0x54, 0x4e, 0xf3, 0x3c, 0xfe, 0xb9, 0xee, 0x8c, 0x36, 0xd8, 0xe6, 0xc9,
	0xde, 0xc4, 0x15, 0xc8, 0x6b, 0xd8, 0xa9, 0xd0, 0xd2, 0x13, 0x6c, 0x33, 0x65, 0x2e, 0x8c, 0x8d,
	0x57, 0x6b, 0x82, 0x5e, 0x83, 0xc0, 0x58, 0x2d, 0xe4, 0xc4, 0x03, 0x74, 0x8b, 0x70, 0x87, 0x41,
	0xef, 0xdd, 0x1f, 0x39, 0x31, 0x79, 0x08, 0x83, 0x33, 0xdd, 0x3d, 0xdc, 0x2f, 0x0f, 0xb7, 0x73,
	0xaa, 0x4f, 0x45, 0x3a, 0xf8, 0xd1, 0x81, 0x0d, 0x86, 0xa6, 0x54, 0xd2, 0x20, 0xb9, 0x09, 0x9d,
	0x3a, 0x89, 0x70, 0x41, 0x33, 0xbc, 0x90, 0xb2, 0x4f, 0xe9, 0x41, 0xf5, 0xcb, 0x7c, 0x23, 0x79,
	0x01, 0x9b, 0x4d, 0x14, 0x77, 0xb6, 0xc8, 0x89, 0x69, 0x43, 0xdc, 0x8c, 0x6a, 0x3c, 0x8f, 0xc5,
	0x2e, 0x35, 0x38, 0xc9, 0x3d, 0x58, 0x0f, 0xd9, 0xc7, 0x51, 0xed, 0x78, 0x65, 0xc1, 0xd1, 0x4f,
	0xc6, 0xd8, 0xff, 0xb3, 0xe3, 0x76, 0x32, 0x82, 0x48, 0xab, 0xa3, 0xf8, 0x5c, 0xad, 0xba, 0x4b,
	0xff, 0x7a, 0x54, 0xe9, 0xf1, 0x45, 0x50, 0xa6, 0x8e, 0x58, 0xe5, 0xd1, 0xfb, 0x1e, 0x41, 0xe4,
	0x16, 0xe4, 0x32, 0xac, 0xb9, 0x65, 0x35, 0x3b, 0x1f, 0x12, 0x77, 0x37, 0x1d, 0xd6, 0x71, 0x4b,
	0x37, 0x18, 0xcf, 0x20, 0x3e, 0x2d, 0x8a, 0xf8, 0x63, 0xb2, 0x7c, 0xcc, 0xb6, 0x5b, 0xd3, 0x21,
	0x8f, 0x61, 0xbb, 0xd5, 0x36, 0xfe, 0x94, 0x2c, 0x1d, 0x9c, 0xad, 0x16, 0x4b, 0xb2, 0xdf, 0x92,
	0xd3, 0xe7, 0x76, 0xbe, 0x0c, 0x0f, 0x45, 0xc1, 0x73, 0x6f, 0xb6, 0x90, 0xca, 0xde, 0xc2, 0xc7,
	0xf4, 0x25, 0xf9, 0xd7, 0xaf, 0xe9, 0x15, 0xf4, 0xad, 0xe6, 0xd2, 0xe4, 0xdc, 0x62, 0x96, 0x2e,
	0x70, 0x7d, 0x5d, 0x7e, 0xc6, 0xee, 0x1f, 0x7d, 0x63, 0x8e, 0xf6, 0x9e, 0x40, 0x5f, 0xa8, 0x66,
	0xe6, 0xb3, 0xa7, 0xed, 0xe5, 0x9d, 0xff, 0x79, 0xf4, 0xde, 0xac, 0xd5, 0x0f, 0xcb, 0xed, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x61, 0x2d, 0x34, 0x33, 0x05, 0x00, 0x00,
}