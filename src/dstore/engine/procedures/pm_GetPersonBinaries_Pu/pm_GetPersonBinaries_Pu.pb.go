// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonBinaries_Pu.proto
// DO NOT EDIT!

/*
Package pm_GetPersonBinaries_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonBinaries_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonBinaries_Pu

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
	PersonId                       *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                   bool                        `protobuf:"varint,1004,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	IncludeBinaryCode              *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=include_binary_code,json=includeBinaryCode" json:"include_binary_code,omitempty"`
	IncludeBinaryCodeNull          bool                        `protobuf:"varint,1005,opt,name=include_binary_code_null,json=includeBinaryCodeNull" json:"include_binary_code_null,omitempty"`
	BinaryCharacteristicId         *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=binary_characteristic_id,json=binaryCharacteristicId" json:"binary_characteristic_id,omitempty"`
	BinaryCharacteristicIdNull     bool                        `protobuf:"varint,1006,opt,name=binary_characteristic_id_null,json=binaryCharacteristicIdNull" json:"binary_characteristic_id_null,omitempty"`
	Value                          *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=value" json:"value,omitempty"`
	ValueNull                      bool                        `protobuf:"varint,1007,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
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

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetIncludeBinaryCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.IncludeBinaryCode
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicId
	}
	return nil
}

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
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
	PersonId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	BinaryDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=binary_description,json=binaryDescription" json:"binary_description,omitempty"`
	BinaryCodeId      *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	ThumbnailCode     *dstore_values.BytesValue   `protobuf:"bytes,10004,opt,name=thumbnail_code,json=thumbnailCode" json:"thumbnail_code,omitempty"`
	BinaryCode        *dstore_values.BytesValue   `protobuf:"bytes,10005,opt,name=binary_code,json=binaryCode" json:"binary_code,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetBinaryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryDescription
	}
	return nil
}

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Response_Row) GetThumbnailCode() *dstore_values.BytesValue {
	if m != nil {
		return m.ThumbnailCode
	}
	return nil
}

func (m *Response_Row) GetBinaryCode() *dstore_values.BytesValue {
	if m != nil {
		return m.BinaryCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonBinaries_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonBinaries_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonBinaries_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x4f, 0x13, 0x4f,
	0x14, 0x0f, 0xf4, 0xdf, 0x02, 0x03, 0x7f, 0x84, 0x6d, 0x24, 0x4b, 0x51, 0xa2, 0x18, 0x12, 0x9f,
	0xb6, 0x44, 0x12, 0xa3, 0xf1, 0xc9, 0x22, 0x31, 0x60, 0xa8, 0x64, 0xbd, 0x24, 0x1a, 0x93, 0xcd,
	0x76, 0xf7, 0x50, 0x26, 0x69, 0x67, 0x9a, 0x99, 0x2d, 0xa4, 0xcf, 0x7e, 0x01, 0xaf, 0xdf, 0xd1,
	0xbb, 0x0f, 0x7e, 0x01, 0xcf, 0xcc, 0xd9, 0x2d, 0xdd, 0x05, 0x2c, 0xbc, 0x40, 0x67, 0xe7, 0xfc,
	0x2e, 0x73, 0xe6, 0x9c, 0x33, 0xec, 0x6e, 0xac, 0x13, 0xa9, 0xa0, 0x0e, 0xa2, 0xcd, 0x05, 0xd4,
	0x7b, 0x4a, 0x46, 0x10, 0xf7, 0x15, 0xe8, 0x7a, 0xaf, 0x1b, 0x3c, 0x86, 0x64, 0x1f, 0x94, 0x96,
	0xa2, 0xc1, 0x45, 0xa8, 0x38, 0xe8, 0x60, 0xbf, 0xef, 0x61, 0x48, 0x22, 0x9d, 0x75, 0xc2, 0x79,
	0x84, 0xf3, 0xce, 0x09, 0xae, 0x55, 0x53, 0xfa, 0xa3, 0xb0, 0xd3, 0x07, 0x4d, 0xd8, 0xda, 0x72,
	0x5e, 0x13, 0x94, 0x92, 0x2a, 0xdd, 0x5a, 0xc9, 0x6f, 0x75, 0x41, 0xeb, 0xb0, 0x0d, 0xe9, 0xe6,
	0xad, 0xe2, 0x66, 0x12, 0x72, 0x71, 0x20, 0x55, 0x37, 0x4c, 0xb8, 0x14, 0x14, 0xb4, 0xf6, 0xa7,
	0xc2, 0xd8, 0x7e, 0xa8, 0x42, 0xdc, 0x45, 0x37, 0xce, 0x1b, 0x76, 0xad, 0x67, 0x5d, 0x05, 0x3c,
	0x06, 0x91, 0xf0, 0x03, 0x1e, 0xd9, 0xe8, 0x80, 0x1c, 0xb9, 0x13, 0x37, 0x26, 0x6e, 0xcf, 0xde,
	0xa9, 0x79, 0xe9, 0x71, 0x52, 0x9f, 0x3a, 0x51, 0x5c, 0xb4, 0x5f, 0x9a, 0x85, 0x5f, 0x23, 0xfc,
	0x4e, 0x0e, 0x6e, 0xb7, 0xb4, 0xb3, 0xcb, 0x6e, 0xfe, 0x8b, 0x3d, 0x10, 0xfd, 0x4e, 0xc7, 0xfd,
	0x3a, 0x85, 0x1a, 0xd3, 0xfe, 0xea, 0xf9, 0x3c, 0x4d, 0x0c, 0x73, 0x1e, 0xb2, 0xf9, 0x94, 0x2b,
	0x19, 0xf4, 0x00, 0x09, 0xdd, 0x49, 0xeb, 0x6d, 0xa5, 0xe0, 0x8d, 0x8b, 0x04, 0xda, 0xa0, 0xc8,
	0xdc, 0x1c, 0x41, 0x9e, 0x23, 0x62, 0x27, 0x76, 0x3c, 0x56, 0xcd, 0x53, 0x90, 0x81, 0x6f, 0x64,
	0x60, 0x61, 0x34, 0xd6, 0x4a, 0x36, 0xd8, 0x7c, 0x14, 0x6a, 0x08, 0x34, 0x08, 0xcd, 0x13, 0x7e,
	0x04, 0x6e, 0xe9, 0x4c, 0xc9, 0x96, 0x94, 0x1d, 0x08, 0xc9, 0xac, 0xff, 0xbf, 0x81, 0x3c, 0xcb,
	0x10, 0x4e, 0x9d, 0x55, 0xf3, 0x1c, 0xa4, 0xf9, 0x9d, 0x34, 0x17, 0x73, 0xc1, 0x56, 0xf4, 0x1e,
	0x9b, 0x19, 0xe6, 0xcc, 0xfd, 0x6f, 0xfc, 0x11, 0xa7, 0xb3, 0xbc, 0x39, 0xeb, 0xc3, 0x0c, 0x65,
	0x27, 0xfb, 0x41, 0x2a, 0x73, 0x59, 0x88, 0x15, 0x78, 0xc2, 0xaa, 0x5c, 0x44, 0x9d, 0x7e, 0x0c,
	0x41, 0xcb, 0x94, 0xe2, 0x20, 0x88, 0x64, 0x0c, 0x6e, 0x79, 0xbc, 0xd4, 0x62, 0x8a, 0xb3, 0x15,
	0x3c, 0xd8, 0x42, 0x14, 0xba, 0x75, 0xcf, 0x20, 0x23, 0xf5, 0x9f, 0xa4, 0x7e, 0xf5, 0x14, 0xca,
	0xda, 0x78, 0xc1, 0xdc, 0x0c, 0x71, 0x88, 0xf5, 0x18, 0x61, 0x39, 0x72, 0x9d, 0xf0, 0xc8, 0x1c,
	0xbb, 0x32, 0xde, 0xcb, 0x12, 0x81, 0xb7, 0x72, 0x58, 0x4c, 0x42, 0x83, 0x5d, 0x3f, 0x8f, 0x96,
	0x5c, 0xfd, 0x22, 0x57, 0xb5, 0xb3, 0xf1, 0xd6, 0xda, 0x06, 0x2b, 0x5b, 0x49, 0x77, 0x6a, 0x6c,
	0xf5, 0x53, 0xa0, 0xb3, 0xca, 0x98, 0xfd, 0x41, 0x12, 0xbf, 0x49, 0x62, 0xc6, 0x7e, 0x32, 0x8c,
	0x6b, 0x6f, 0xcb, 0x6c, 0xda, 0x07, 0xdd, 0x93, 0x42, 0x83, 0xa1, 0xb7, 0x3d, 0x5d, 0x6c, 0xae,
	0x74, 0x56, 0x50, 0xbf, 0x6f, 0x9b, 0xbf, 0x3e, 0x05, 0x3a, 0xaf, 0xd8, 0x82, 0xe9, 0xe6, 0x60,
	0xa4, 0x9d, 0xb1, 0xfa, 0x4b, 0x08, 0xf6, 0x0a, 0xe0, 0x62, 0xd3, 0xef, 0xe1, 0x7a, 0xe7, 0x64,
	0xed, 0x5f, 0xe9, 0xe6, 0x3f, 0xe0, 0x05, 0x4e, 0xa5, 0x53, 0x04, 0x8b, 0xdb, 0x30, 0xae, 0x9e,
	0x62, 0xa4, 0x19, 0xb3, 0x47, 0xff, 0xfd, 0x2c, 0xdc, 0xd9, 0x66, 0x25, 0x25, 0x8f, 0xb1, 0x44,
	0x0d, 0x6a, 0xd3, 0xbb, 0xd0, 0xc0, 0xf3, 0xb2, 0x24, 0x78, 0xbe, 0x3c, 0xf6, 0x0d, 0xbe, 0xf6,
	0x65, 0x92, 0x95, 0x70, 0xe1, 0x2c, 0xb1, 0x0a, 0x2e, 0xcd, 0xed, 0xbf, 0x6b, 0x62, 0x5e, 0xca,
	0x7e, 0x19, 0x97, 0x78, 0xa1, 0xf7, 0x47, 0xfb, 0xe1, 0x7d, 0xf3, 0x32, 0x0d, 0xb1, 0xcb, 0x9c,
	0xb4, 0x16, 0x62, 0xd0, 0x91, 0xe2, 0x3d, 0x9b, 0xb8, 0x0f, 0xcd, 0xb1, 0xb7, 0xba, 0x48, 0xb0,
	0x47, 0x27, 0x28, 0x33, 0x0b, 0x46, 0x0b, 0x1c, 0xbd, 0x7c, 0xbc, 0x80, 0x97, 0xb9, 0xd6, 0xb0,
	0xe8, 0xd1, 0x0f, 0x8e, 0xb0, 0xe4, 0xb0, 0xdf, 0x6d, 0x89, 0x90, 0x77, 0xa8, 0xe9, 0x3e, 0x11,
	0xc7, 0x72, 0x71, 0xa0, 0x0c, 0x12, 0xd0, 0xe9, 0x38, 0x19, 0x22, 0x6c, 0xbf, 0x3d, 0x60, 0xb3,
	0xa3, 0x4d, 0xfb, 0x79, 0x2c, 0x9e, 0x9d, 0x38, 0x68, 0x3c, 0x65, 0x2b, 0x5c, 0x16, 0x2f, 0x6a,
	0xf8, 0xa2, 0xbd, 0xde, 0xb8, 0xec, 0x5b, 0xd7, 0xaa, 0xd8, 0x37, 0x65, 0xf3, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x09, 0x07, 0xcb, 0x22, 0x26, 0x07, 0x00, 0x00,
}