// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonDetails_Pu.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonDetails_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonDetails_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonDetails_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	CaseSensitive                  *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull              bool                        `protobuf:"varint,1003,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
	PersonCharacteristicId         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	PersonCharacteristicIdNull     bool                        `protobuf:"varint,1004,opt,name=person_characteristic_id_null,json=personCharacteristicIdNull" json:"person_characteristic_id_null,omitempty"`
	PersonId                       *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                   bool                        `protobuf:"varint,1005,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
	}
	return nil
}

func (m *Parameters) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonDetails_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonDetails_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonDetails_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xd6, 0x56, 0xba, 0x15, 0x33, 0xc6, 0x70, 0xa5, 0x29, 0xa4, 0x6c, 0x82, 0x21, 0x04, 0x57,
	0x29, 0x1a, 0x42, 0xda, 0x2d, 0x05, 0x84, 0x8a, 0xd4, 0xa9, 0x98, 0x1f, 0x09, 0x84, 0x14, 0x65,
	0xc9, 0x69, 0xb1, 0xd4, 0xda, 0x95, 0xed, 0x6e, 0xda, 0x5b, 0x70, 0xcd, 0x9b, 0xf0, 0x48, 0xfc,
	0xbd, 0x03, 0xc7, 0x39, 0x49, 0x59, 0xc2, 0xca, 0xe0, 0xa6, 0xad, 0x7b, 0xbe, 0x3f, 0xfb, 0x1c,
	0x9b, 0x1d, 0x64, 0xd6, 0x69, 0x03, 0x5d, 0x50, 0x63, 0xa9, 0xa0, 0x3b, 0x33, 0x3a, 0x85, 0x6c,
	0x6e, 0xc0, 0x76, 0x67, 0xd3, 0x78, 0xa0, 0x33, 0x39, 0x3a, 0x1d, 0x82, 0xb1, 0x5a, 0x3d, 0x05,
	0x97, 0xc8, 0x89, 0x8d, 0x87, 0xf3, 0x08, 0x41, 0x4e, 0xf3, 0x7b, 0xc4, 0x8c, 0x88, 0x19, 0x2d,
	0x85, 0x87, 0xed, 0xc2, 0xe2, 0x38, 0x99, 0xcc, 0xc1, 0x12, 0x3b, 0xbc, 0x51, 0xf5, 0x05, 0x63,
	0xb4, 0x29, 0x4a, 0x9d, 0x6a, 0x69, 0x0a, 0xd6, 0x26, 0x63, 0x28, 0x8a, 0x77, 0xea, 0x45, 0xb4,
	0x51, 0x23, 0x6d, 0xa6, 0x89, 0x93, 0x5a, 0x11, 0x68, 0xef, 0x73, 0x93, 0xb1, 0x61, 0x62, 0x12,
	0xac, 0x62, 0x18, 0xfe, 0x81, 0xdd, 0x9c, 0xe5, 0xa1, 0x62, 0x99, 0x81, 0x72, 0x72, 0x24, 0xd3,
	0x1c, 0x1d, 0x53, 0xa2, 0x60, 0xe5, 0xd6, 0xca, 0xfd, 0x2b, 0xfb, 0x61, 0x54, 0x6c, 0xa8, 0xc8,
	0x69, 0x9d, 0x91, 0x6a, 0xfc, 0xd6, 0x2f, 0x44, 0x48, 0xfc, 0x7e, 0x85, 0x9e, 0x97, 0x2c, 0x7f,
	0xc1, 0x6e, 0xff, 0x4d, 0x3d, 0x56, 0xf3, 0xc9, 0x24, 0xf8, 0xba, 0x8e, 0x1e, 0x2d, 0xb1, 0xbb,
	0x5c, 0xe7, 0x10, 0x61, 0xfc, 0x31, 0xdb, 0x2c, 0xb4, 0xdc, 0xe9, 0x0c, 0x50, 0x30, 0x58, 0xcd,
	0xb3, 0x75, 0x6a, 0xd9, 0xa4, 0x72, 0x30, 0x06, 0x43, 0xe1, 0x36, 0x88, 0xf2, 0x1a, 0x19, 0xfd,
	0x8c, 0x47, 0xac, 0x5d, 0x95, 0xa0, 0x00, 0xdf, 0x28, 0xc0, 0xd6, 0x59, 0x6c, 0x6e, 0xd9, 0x63,
	0x9b, 0x69, 0x62, 0x21, 0xb6, 0xa0, 0xac, 0x74, 0xf2, 0x18, 0x82, 0xc6, 0xb9, 0x96, 0x47, 0x5a,
	0x4f, 0x20, 0xa1, 0xb0, 0xe2, 0xaa, 0xa7, 0xbc, 0x2a, 0x19, 0xbc, 0xcb, 0xda, 0x55, 0x0d, 0xf2,
	0xfc, 0x4e, 0x9e, 0xd7, 0x2b, 0xe0, 0xdc, 0xf4, 0x0d, 0x0b, 0x8a, 0x90, 0xe9, 0x47, 0xec, 0x53,
	0x8a, 0x6d, 0x92, 0xd6, 0xc9, 0xd4, 0xef, 0xf8, 0xd2, 0xc5, 0x3b, 0xde, 0x26, 0xf2, 0x93, 0x0a,
	0x17, 0xf7, 0xde, 0x63, 0x3b, 0xcb, 0x64, 0x29, 0xd1, 0x0f, 0x4a, 0x14, 0x9e, 0xcf, 0xcf, 0xa3,
	0x1d, 0xb0, 0xcb, 0x8b, 0x76, 0x06, 0xcd, 0x8b, 0xb3, 0xb4, 0xca, 0x96, 0xf2, 0xbb, 0x8b, 0xe6,
	0x95, 0x76, 0x3f, 0xc9, 0x6e, 0xa3, 0x84, 0x78, 0x83, 0xbd, 0x2f, 0xab, 0xac, 0x25, 0xc0, 0xce,
	0xb4, 0xb2, 0xc0, 0x1f, 0xb0, 0x66, 0x3e, 0xfa, 0xf5, 0x19, 0x2c, 0x2e, 0x15, 0x5d, 0x8b, 0x67,
	0xfe, 0x53, 0x10, 0x90, 0xbf, 0x63, 0x5b, 0x7e, 0xe8, 0xe3, 0x33, 0x53, 0x8f, 0x43, 0xd2, 0x40,
	0x72, 0x54, 0x23, 0xd7, 0xef, 0xc6, 0x00, 0xd7, 0xfd, 0xdf, 0x6b, 0x71, 0x6d, 0x5a, 0xfd, 0x03,
	0xb7, 0xbe, 0x5e, 0x5c, 0x36, 0x9c, 0x01, 0xaf, 0xb8, 0xfb, 0x87, 0x22, 0x5d, 0xc5, 0x01, 0x7d,
	0x8b, 0x12, 0xce, 0x9f, 0xb3, 0x86, 0xd1, 0x27, 0xd8, 0x3a, 0xcf, 0x7a, 0x14, 0xfd, 0xe3, 0xcb,
	0x10, 0x95, 0xc7, 0x10, 0x09, 0x7d, 0x22, 0xbc, 0x42, 0xb8, 0xc3, 0x1a, 0xf8, 0x9b, 0x6f, 0xb3,
	0x35, 0x5c, 0xf9, 0x0e, 0x7c, 0x3a, 0xc4, 0x83, 0x69, 0x8a, 0x26, 0x2e, 0xfb, 0x59, 0xef, 0x25,
	0xeb, 0x48, 0x5d, 0x97, 0x5f, 0x3c, 0x59, 0xef, 0xf7, 0xff, 0xff, 0x31, 0x3b, 0x5a, 0xcb, 0x9f,
	0x8c, 0x87, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xc5, 0x99, 0xb7, 0x09, 0x05, 0x00, 0x00,
}