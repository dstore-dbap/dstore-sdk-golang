// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyRelationshipSetts_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyRelationshipSetts_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyRelationshipSetts_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyRelationshipSetts_Ad

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
	RelationshipId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	RelationshipIdNull   bool                        `protobuf:"varint,1001,opt,name=relationship_id_null,json=relationshipIdNull" json:"relationship_id_null,omitempty"`
	FromPersonTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=from_person_type_id,json=fromPersonTypeId" json:"from_person_type_id,omitempty"`
	FromPersonTypeIdNull bool                        `protobuf:"varint,1002,opt,name=from_person_type_id_null,json=fromPersonTypeIdNull" json:"from_person_type_id_null,omitempty"`
	ToPersonTypeId       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=to_person_type_id,json=toPersonTypeId" json:"to_person_type_id,omitempty"`
	ToPersonTypeIdNull   bool                        `protobuf:"varint,1003,opt,name=to_person_type_id_null,json=toPersonTypeIdNull" json:"to_person_type_id_null,omitempty"`
	KeyVariable          *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull      bool                        `protobuf:"varint,1004,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	Value                *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=value" json:"value,omitempty"`
	ValueNull            bool                        `protobuf:"varint,1005,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetRelationshipId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelationshipId
	}
	return nil
}

func (m *Parameters) GetFromPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromPersonTypeId
	}
	return nil
}

func (m *Parameters) GetToPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToPersonTypeId
	}
	return nil
}

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyRelationshipSetts_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyRelationshipSetts_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyRelationshipSetts_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x6b, 0x13, 0x41,
	0x10, 0xa7, 0x4d, 0xd3, 0xd6, 0xa9, 0x98, 0x76, 0x5b, 0x4a, 0x4c, 0xb0, 0x48, 0x7d, 0x11, 0x0a,
	0x17, 0xb5, 0xe0, 0x1f, 0xd0, 0x07, 0x45, 0x85, 0x28, 0x29, 0x65, 0xd5, 0x82, 0xbe, 0x1c, 0x57,
	0x6f, 0x12, 0x17, 0xef, 0x6e, 0x8f, 0xdd, 0x4d, 0x4b, 0xbe, 0x85, 0x5f, 0xc8, 0x47, 0x3f, 0x8c,
	0xff, 0xbe, 0x83, 0xb3, 0x7f, 0x5a, 0xef, 0xae, 0x62, 0xb5, 0x2f, 0xc9, 0xcd, 0xcd, 0xef, 0xcf,
	0xdc, 0xce, 0xcc, 0xc2, 0xc3, 0x54, 0x1b, 0xa9, 0x70, 0x80, 0xc5, 0x44, 0x14, 0x38, 0x28, 0x95,
	0x7c, 0x8f, 0xe9, 0x54, 0xa1, 0x1e, 0x94, 0x79, 0x3c, 0x92, 0xa9, 0x18, 0xcf, 0x38, 0x66, 0x89,
	0x11, 0xb2, 0xd0, 0x1f, 0x44, 0xf9, 0x0a, 0x8d, 0xd1, 0xf1, 0xe3, 0x34, 0x22, 0xa0, 0x91, 0x6c,
	0xc7, 0xb3, 0x23, 0xcf, 0x8e, 0xfe, 0x4a, 0xe9, 0xad, 0x07, 0xab, 0xa3, 0x24, 0x9b, 0xa2, 0xf6,
	0x0a, 0xbd, 0xab, 0x75, 0x7f, 0x54, 0x4a, 0xaa, 0x90, 0xea, 0xd7, 0x53, 0x39, 0x6a, 0x9d, 0x4c,
	0x30, 0x24, 0x6f, 0x34, 0x93, 0x26, 0x11, 0xc5, 0x58, 0xaa, 0xdc, 0xf9, 0x7a, 0xd0, 0xf6, 0x97,
	0x05, 0x80, 0xfd, 0x44, 0x25, 0x94, 0x45, 0xa5, 0xd9, 0x53, 0xe8, 0xa8, 0x4a, 0x61, 0xb1, 0x48,
	0xbb, 0x73, 0xd7, 0xe7, 0x6e, 0xae, 0xdc, 0xe9, 0x47, 0xe1, 0x3b, 0x42, 0x69, 0xa2, 0x30, 0x38,
	0x41, 0x75, 0x60, 0x23, 0x7e, 0xa5, 0xca, 0x19, 0xa6, 0xec, 0x36, 0x6c, 0x34, 0x54, 0xe2, 0x62,
	0x9a, 0x65, 0xdd, 0xaf, 0x4b, 0xa4, 0xb5, 0xcc, 0x59, 0x1d, 0xbe, 0x47, 0x29, 0xf6, 0x02, 0xd6,
	0xc7, 0x4a, 0xe6, 0x71, 0x49, 0x55, 0xc8, 0x22, 0x36, 0xb3, 0x12, 0xad, 0xf9, 0xfc, 0xf9, 0xe6,
	0xab, 0x96, 0xb7, 0xef, 0x68, 0xaf, 0x89, 0x45, 0xf6, 0xf7, 0xa0, 0xfb, 0x07, 0x2d, 0x5f, 0xc2,
	0x37, 0x5f, 0xc2, 0x46, 0x93, 0xe4, 0x8a, 0x78, 0x0e, 0x6b, 0x46, 0x36, 0x4b, 0x68, 0xfd, 0xc3,
	0xf7, 0x1b, 0x59, 0x2b, 0x60, 0x17, 0x36, 0xcf, 0xe8, 0x78, 0xfb, 0xef, 0xe1, 0x04, 0xea, 0x04,
	0x67, 0xfe, 0x08, 0x2e, 0x7f, 0xc4, 0x59, 0x7c, 0x94, 0x28, 0x91, 0x1c, 0x66, 0xd8, 0x5d, 0x70,
	0xbe, 0xbd, 0x86, 0xaf, 0x36, 0x4a, 0x14, 0x13, 0x6f, 0xbb, 0x42, 0xf8, 0x83, 0x00, 0x67, 0x3b,
	0xb0, 0x56, 0xa5, 0x7b, 0xbb, 0x1f, 0xde, 0xae, 0x53, 0x01, 0x3a, 0xaf, 0x5b, 0xd0, 0x76, 0x7a,
	0xdd, 0xf6, 0xb9, 0x26, 0x1e, 0xc8, 0xb6, 0x00, 0xdc, 0x83, 0xd7, 0xfd, 0xe9, 0x75, 0x2f, 0xb9,
	0x57, 0x56, 0x71, 0xfb, 0xf3, 0x3c, 0x2c, 0x73, 0xd4, 0x25, 0x35, 0x15, 0xad, 0xbc, 0x9b, 0xd2,
	0x30, 0x3b, 0xa7, 0xf2, 0x61, 0x07, 0xfc, 0x04, 0x3f, 0xb3, 0xbf, 0xdc, 0x03, 0xd9, 0x5b, 0x58,
	0xb5, 0xf3, 0x19, 0x57, 0x06, 0x94, 0x7a, 0xdf, 0x22, 0x72, 0xd4, 0x20, 0x37, 0xc7, 0x78, 0x44,
	0xf1, 0xf0, 0x77, 0xcc, 0x3b, 0x79, 0xfd, 0x05, 0xbb, 0x0f, 0x4b, 0x61, 0x2f, 0xa8, 0x95, 0x56,
	0x71, 0xeb, 0x8c, 0xa2, 0xdf, 0x9a, 0x91, 0xff, 0xe7, 0x27, 0x70, 0xf6, 0x12, 0x5a, 0x4a, 0x1e,
	0x53, 0x23, 0x2c, 0xeb, 0x41, 0xf4, 0x1f, 0x8b, 0x1c, 0x9d, 0x1c, 0x45, 0xc4, 0xe5, 0x31, 0xb7,
	0x2a, 0xbd, 0x6b, 0xd0, 0xa2, 0x67, 0xb6, 0x09, 0x8b, 0x14, 0xd9, 0xb9, 0xfa, 0xb4, 0x47, 0x87,
	0xd3, 0xe6, 0x6d, 0x0a, 0x87, 0xe9, 0x93, 0x37, 0xd0, 0x17, 0xb2, 0x69, 0x71, 0x7a, 0xd3, 0xbc,
	0xbb, 0x7b, 0xb1, 0x3b, 0xe8, 0x70, 0xd1, 0x6d, 0xf9, 0xee, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x82, 0x90, 0xb6, 0x17, 0xc4, 0x04, 0x00, 0x00,
}