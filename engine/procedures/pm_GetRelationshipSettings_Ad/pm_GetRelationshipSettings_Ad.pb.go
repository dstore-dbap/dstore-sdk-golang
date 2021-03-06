// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetRelationshipSettings_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetRelationshipSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetRelationshipSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetRelationshipSettings_Ad

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
	RelationshipId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	RelationshipIdNull   bool                        `protobuf:"varint,1001,opt,name=relationship_id_null,json=relationshipIdNull" json:"relationship_id_null,omitempty"`
	FromPersonTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=from_person_type_id,json=fromPersonTypeId" json:"from_person_type_id,omitempty"`
	FromPersonTypeIdNull bool                        `protobuf:"varint,1002,opt,name=from_person_type_id_null,json=fromPersonTypeIdNull" json:"from_person_type_id_null,omitempty"`
	ToPersonTypeId       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=to_person_type_id,json=toPersonTypeId" json:"to_person_type_id,omitempty"`
	ToPersonTypeIdNull   bool                        `protobuf:"varint,1003,opt,name=to_person_type_id_null,json=toPersonTypeIdNull" json:"to_person_type_id_null,omitempty"`
	KeyVariable          *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull      bool                        `protobuf:"varint,1004,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
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
	ToPersonTypeDescription   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=to_person_type_description,json=toPersonTypeDescription" json:"to_person_type_description,omitempty"`
	FromPersonTypeDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=from_person_type_description,json=fromPersonTypeDescription" json:"from_person_type_description,omitempty"`
	RelationshipId            *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	FromPersonTypeId          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=from_person_type_id,json=fromPersonTypeId" json:"from_person_type_id,omitempty"`
	Relationship              *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=relationship" json:"relationship,omitempty"`
	Value                     *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=value" json:"value,omitempty"`
	KeyVariable               *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	ToPersonTypeId            *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=to_person_type_id,json=toPersonTypeId" json:"to_person_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetToPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ToPersonTypeDescription
	}
	return nil
}

func (m *Response_Row) GetFromPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.FromPersonTypeDescription
	}
	return nil
}

func (m *Response_Row) GetRelationshipId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelationshipId
	}
	return nil
}

func (m *Response_Row) GetFromPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromPersonTypeId
	}
	return nil
}

func (m *Response_Row) GetRelationship() *dstore_values.StringValue {
	if m != nil {
		return m.Relationship
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Response_Row) GetToPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToPersonTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetRelationshipSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetRelationshipSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetRelationshipSettings_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_GetRelationshipSettings_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x4f, 0xd4, 0x40,
	0x10, 0x0e, 0x1e, 0xc7, 0x91, 0x81, 0x78, 0xb0, 0x10, 0x2c, 0x77, 0xc6, 0x18, 0x7c, 0x31, 0x21,
	0xe9, 0xa9, 0x3c, 0xa8, 0x89, 0x4a, 0x34, 0x20, 0x41, 0x42, 0x43, 0xaa, 0x21, 0x6a, 0x8c, 0x4d,
	0xa1, 0xc3, 0xd9, 0x70, 0xd7, 0x6d, 0x76, 0xf7, 0x20, 0xfc, 0x05, 0xbe, 0xfa, 0x5b, 0xff, 0x42,
	0x13, 0x7f, 0xbc, 0xf9, 0x0f, 0x38, 0xed, 0xf6, 0xa0, 0xdd, 0xbb, 0x78, 0x96, 0x17, 0xc8, 0xde,
	0xcc, 0xf7, 0x7d, 0xb3, 0x33, 0xdf, 0x6c, 0xe1, 0x5e, 0x20, 0x15, 0x17, 0xd8, 0xc2, 0xa8, 0x1d,
	0x46, 0xd8, 0x8a, 0x05, 0xdf, 0xc7, 0xa0, 0x27, 0x50, 0xb6, 0xe2, 0xae, 0xb7, 0x81, 0xca, 0xc5,
	0x8e, 0xaf, 0x42, 0x1e, 0xc9, 0x37, 0x61, 0xfc, 0x14, 0x95, 0x0a, 0xa3, 0xb6, 0xf4, 0x1e, 0x06,
	0x36, 0x25, 0x2a, 0xce, 0x96, 0x35, 0xda, 0xd6, 0x68, 0xfb, 0x9f, 0x90, 0xc6, 0x5c, 0x26, 0x75,
	0xe4, 0x77, 0x7a, 0x28, 0x35, 0x43, 0x63, 0xb1, 0xa8, 0x8f, 0x42, 0x70, 0x91, 0x85, 0x9a, 0xc5,
	0x50, 0x17, 0xa5, 0xf4, 0xdb, 0x98, 0x05, 0xaf, 0x99, 0x41, 0xe5, 0x87, 0xd1, 0x01, 0x17, 0xdd,
	0x54, 0x5a, 0x27, 0x2d, 0xfd, 0xa9, 0x00, 0xec, 0xf8, 0xc2, 0xa7, 0x28, 0x0a, 0xc9, 0xd6, 0xa0,
	0x2e, 0x72, 0xb5, 0x79, 0x61, 0x60, 0x8d, 0x5d, 0x1d, 0xbb, 0x3e, 0x75, 0xab, 0x69, 0x67, 0xf7,
	0xc8, 0x4a, 0x0b, 0x23, 0x85, 0x6d, 0x14, 0xbb, 0xc9, 0xc9, 0xbd, 0x98, 0xc7, 0x6c, 0x06, 0xec,
	0x26, 0xcc, 0x1b, 0x2c, 0x5e, 0xd4, 0xeb, 0x74, 0xac, 0x1f, 0x35, 0xe2, 0x9a, 0x74, 0x59, 0x31,
	0xdd, 0xa1, 0x10, 0x7b, 0x02, 0x73, 0x07, 0x82, 0x77, 0xbd, 0x98, 0xaa, 0xe0, 0x91, 0xa7, 0x4e,
	0x62, 0x4c, 0xc4, 0x2f, 0x8c, 0x16, 0x9f, 0x49, 0x70, 0x3b, 0x29, 0xec, 0x19, 0xa1, 0x48, 0xfe,
	0x36, 0x58, 0x43, 0xb8, 0x74, 0x09, 0x3f, 0x75, 0x09, 0xf3, 0x26, 0x28, 0x2d, 0xe2, 0x31, 0xcc,
	0x2a, 0x6e, 0x96, 0x50, 0xf9, 0x8f, 0xfb, 0x2b, 0x5e, 0x28, 0x60, 0x05, 0x16, 0x06, 0x78, 0xb4,
	0xfc, 0xaf, 0xac, 0x03, 0x45, 0x40, 0x2a, 0x7e, 0x1f, 0xa6, 0x0f, 0xf1, 0xc4, 0x3b, 0xf2, 0x45,
	0xe8, 0xef, 0x75, 0xd0, 0x1a, 0x4f, 0x75, 0x1b, 0x86, 0xae, 0x54, 0x82, 0xcc, 0xa2, 0x65, 0xa7,
	0x28, 0x7f, 0x37, 0x4b, 0x67, 0xcb, 0x30, 0x9b, 0x87, 0x6b, 0xb9, 0xdf, 0x5a, 0xae, 0x9e, 0x4b,
	0x4c, 0xb4, 0x96, 0xde, 0xd6, 0x60, 0xd2, 0x45, 0x19, 0xd3, 0x08, 0x90, 0xdd, 0x80, 0x6a, 0xea,
	0xa9, 0x6c, 0xd2, 0xa7, 0x8a, 0x99, 0x63, 0xb5, 0xdf, 0xd6, 0x93, 0xbf, 0xae, 0x4e, 0x64, 0x2f,
	0x60, 0x26, 0x71, 0x93, 0x97, 0xb3, 0x13, 0x4d, 0xaa, 0x42, 0x60, 0xdb, 0x00, 0x9b, 0xa6, 0xdb,
	0xa6, 0xf3, 0xe6, 0xd9, 0xd9, 0xad, 0x77, 0x8b, 0x3f, 0xb0, 0x3b, 0x50, 0xcb, 0x5c, 0x4c, 0x8d,
	0x4f, 0x18, 0xaf, 0x0c, 0x30, 0x6a, 0x8f, 0x6f, 0xeb, 0xff, 0x6e, 0x3f, 0x9d, 0x6d, 0x41, 0x45,
	0xf0, 0x63, 0x6a, 0x5b, 0x82, 0xba, 0x6b, 0x97, 0x58, 0x3b, 0xbb, 0xdf, 0x0a, 0xdb, 0xe5, 0xc7,
	0x6e, 0xc2, 0xd2, 0xf8, 0x3e, 0x0e, 0x15, 0x3a, 0xb0, 0x05, 0x98, 0xa0, 0x63, 0x62, 0x83, 0x77,
	0x0e, 0x75, 0xa7, 0xea, 0x56, 0xe9, 0x48, 0x13, 0x7e, 0x0e, 0x0d, 0x63, 0xc2, 0x01, 0xca, 0x7d,
	0x11, 0xc6, 0x69, 0x2f, 0xde, 0x3b, 0x23, 0x67, 0x77, 0x29, 0xef, 0x80, 0xb5, 0x33, 0x2c, 0x7b,
	0x05, 0x97, 0x07, 0xcc, 0x9b, 0xe7, 0xfe, 0x30, 0x9a, 0x7b, 0xb1, 0x68, 0xee, 0x3c, 0xfb, 0xfa,
	0xe0, 0x7e, 0x7f, 0x74, 0xca, 0x2f, 0xf8, 0xd6, 0xf0, 0x6d, 0xfd, 0xe4, 0x9c, 0x67, 0x5d, 0x57,
	0x61, 0x3a, 0x4f, 0x6f, 0x7d, 0x1e, 0x7d, 0xc3, 0x02, 0x80, 0x9e, 0x9b, 0x6a, 0x9a, 0x63, 0x7d,
	0x19, 0x8d, 0xd4, 0x99, 0xec, 0x81, 0xb1, 0x6c, 0x5f, 0x9d, 0x72, 0xdb, 0xb6, 0x31, 0xec, 0xa5,
	0xf8, 0xe6, 0x94, 0x7e, 0x2a, 0x1e, 0xbd, 0x86, 0x66, 0xc8, 0x4d, 0xb3, 0x9e, 0x7e, 0x61, 0x5e,
	0xae, 0xb6, 0xb9, 0x0c, 0x0e, 0xfb, 0xf1, 0xa0, 0xf4, 0x47, 0x68, 0x6f, 0x22, 0x7d, 0xe6, 0x57,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x65, 0x46, 0x7b, 0xc5, 0x06, 0x00, 0x00,
}
