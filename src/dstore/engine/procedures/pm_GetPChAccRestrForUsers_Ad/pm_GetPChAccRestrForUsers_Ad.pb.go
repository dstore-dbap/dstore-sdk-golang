// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPChAccRestrForUsers_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPChAccRestrForUsers_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPChAccRestrForUsers_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPChAccRestrForUsers_Ad

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
	UserId                      *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull                  bool                        `protobuf:"varint,1001,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	PersonCharacteristicId      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	PersonCharacteristicIdNull  bool                        `protobuf:"varint,1002,opt,name=person_characteristic_id_null,json=personCharacteristicIdNull" json:"person_characteristic_id_null,omitempty"`
	EffectiveRestrForUserId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=effective_restr_for_user_id,json=effectiveRestrForUserId" json:"effective_restr_for_user_id,omitempty"`
	EffectiveRestrForUserIdNull bool                        `protobuf:"varint,1003,opt,name=effective_restr_for_user_id_null,json=effectiveRestrForUserIdNull" json:"effective_restr_for_user_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Parameters) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Parameters) GetEffectiveRestrForUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.EffectiveRestrForUserId
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
	ReadAccessRestrictionPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=read_access_restriction_pattern,json=readAccessRestrictionPattern" json:"read_access_restriction_pattern,omitempty"`
	AccessRestriction            *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=access_restriction,json=accessRestriction" json:"access_restriction,omitempty"`
	PersonCharacteristicId       *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	RestrictionForUserId         *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=restriction_for_user_id,json=restrictionForUserId" json:"restriction_for_user_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetReadAccessRestrictionPattern() *dstore_values.StringValue {
	if m != nil {
		return m.ReadAccessRestrictionPattern
	}
	return nil
}

func (m *Response_Row) GetAccessRestriction() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessRestriction
	}
	return nil
}

func (m *Response_Row) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetRestrictionForUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RestrictionForUserId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPChAccRestrForUsers_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPChAccRestrForUsers_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPChAccRestrForUsers_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0x1b, 0x92, 0x54, 0x03, 0x12, 0x60, 0x50, 0xbb, 0x24, 0x7c, 0x94, 0x72, 0x41, 0x1c,
	0x36, 0x88, 0x0f, 0xa9, 0x12, 0xa7, 0xb4, 0xa2, 0xa8, 0xa0, 0x46, 0x91, 0xa1, 0x95, 0xe0, 0x62,
	0xb9, 0x5e, 0x27, 0xb5, 0x94, 0xd8, 0x91, 0xed, 0xb4, 0x7f, 0x83, 0xaf, 0x1b, 0x47, 0x7e, 0x1d,
	0xf0, 0x13, 0xb8, 0x30, 0x8e, 0x37, 0x69, 0x36, 0xa5, 0x49, 0xd5, 0x4b, 0xa2, 0xd9, 0x99, 0xf7,
	0xe6, 0x79, 0xde, 0xd8, 0xf0, 0x2a, 0x73, 0xde, 0x58, 0xd9, 0x90, 0xba, 0xab, 0xb4, 0x6c, 0x0c,
	0xac, 0x11, 0x32, 0x1b, 0x5a, 0xe9, 0x1a, 0x83, 0x3e, 0x7b, 0x23, 0x7d, 0x7b, 0xfb, 0xa8, 0x29,
	0x04, 0x95, 0xce, 0xdb, 0x1d, 0x63, 0xf7, 0x9d, 0xb4, 0x8e, 0x35, 0xb3, 0x14, 0xeb, 0xbc, 0x21,
	0x4f, 0x22, 0x38, 0x8d, 0xe0, 0x74, 0x1e, 0xa2, 0x76, 0x2b, 0x6f, 0x74, 0xcc, 0x7b, 0x43, 0xe9,
	0x22, 0x41, 0xed, 0x4e, 0xb1, 0xbb, 0xb4, 0xd6, 0xd8, 0x3c, 0x55, 0x2f, 0xa6, 0xfa, 0xd2, 0x39,
	0xde, 0x95, 0x79, 0xf2, 0xd1, 0x6c, 0xd2, 0x73, 0xa5, 0x3b, 0xc6, 0xf6, 0xb9, 0x57, 0x46, 0xc7,
	0xa2, 0x8d, 0x9f, 0x25, 0x80, 0x36, 0xb7, 0x1c, 0xb3, 0xa8, 0x81, 0xbc, 0x80, 0xea, 0x10, 0xc5,
	0x30, 0x95, 0x25, 0x4b, 0xeb, 0x4b, 0x8f, 0xaf, 0x3e, 0xab, 0xa7, 0xb9, 0xfc, 0x5c, 0x92, 0xd2,
	0x5e, 0x76, 0xa5, 0x3d, 0x08, 0x11, 0xad, 0x84, 0xda, 0xdd, 0x8c, 0x3c, 0x84, 0x6b, 0x39, 0x8a,
	0xe9, 0x61, 0xaf, 0x97, 0xfc, 0xaa, 0x22, 0x76, 0x85, 0x42, 0x4c, 0xb7, 0xf0, 0x13, 0xd9, 0x87,
	0x64, 0x80, 0x0d, 0x8c, 0x66, 0xe2, 0x08, 0xdb, 0x09, 0xec, 0xa6, 0x9c, 0x57, 0x22, 0x74, 0x5a,
	0x5e, 0xdc, 0x69, 0x35, 0x82, 0xb7, 0x0b, 0x58, 0xec, 0xbc, 0x05, 0xf7, 0xce, 0xa3, 0x8d, 0x52,
	0x7e, 0x47, 0x29, 0xb5, 0xff, 0xe3, 0x47, 0xd2, 0x3e, 0x42, 0x5d, 0x76, 0x3a, 0x52, 0x78, 0x75,
	0x2c, 0x99, 0x0d, 0x96, 0x30, 0x9c, 0x12, 0x1b, 0xcf, 0xa1, 0xb4, 0x58, 0xdd, 0xda, 0x04, 0x3f,
	0xed, 0x28, 0xca, 0xdb, 0x81, 0xf5, 0x39, 0xd4, 0x51, 0xe1, 0x9f, 0xa8, 0xb0, 0x7e, 0x0e, 0x47,
	0x90, 0xb8, 0xf1, 0xa3, 0x0c, 0x2b, 0xf8, 0x79, 0x60, 0xb4, 0x93, 0xe4, 0x29, 0x94, 0x47, 0x3b,
	0x90, 0x3b, 0x54, 0x4b, 0x8b, 0x0b, 0x16, 0xf7, 0xe3, 0x75, 0xf8, 0xa5, 0xb1, 0x10, 0x4f, 0x78,
	0x23, 0xb8, 0xcf, 0xa6, 0xec, 0xc7, 0xa1, 0x97, 0x10, 0x9c, 0xce, 0x80, 0x67, 0x97, 0x64, 0x0f,
	0xe3, 0xdd, 0xd3, 0x98, 0x5e, 0xef, 0x17, 0x3f, 0x90, 0x4d, 0xa8, 0xe6, 0x5b, 0x87, 0x83, 0x0a,
	0x8c, 0xf7, 0xcf, 0x30, 0xc6, 0x9d, 0xdc, 0x8b, 0xff, 0x74, 0x5c, 0x4e, 0xde, 0x42, 0xc9, 0x9a,
	0x93, 0xe4, 0xca, 0x08, 0xb5, 0x99, 0x5e, 0xfc, 0x96, 0xa4, 0xe3, 0x49, 0xa4, 0xd4, 0x9c, 0xd0,
	0x40, 0x52, 0xfb, 0xbb, 0x0c, 0x25, 0x0c, 0xc8, 0x2a, 0x54, 0x30, 0x0c, 0xae, 0x7d, 0x6e, 0xe1,
	0x70, 0xca, 0xb4, 0x8c, 0x21, 0xfa, 0x70, 0x08, 0x0f, 0xac, 0xe4, 0x19, 0xe3, 0x42, 0x60, 0xf7,
	0xe8, 0x84, 0x12, 0xe1, 0x00, 0x6c, 0xc0, 0x3d, 0xee, 0x83, 0x4e, 0xbe, 0xb4, 0x8a, 0xd3, 0xcc,
	0x7d, 0x0e, 0x85, 0xba, 0x1b, 0x6d, 0xbe, 0x1b, 0x38, 0x9a, 0x23, 0x0a, 0x7a, 0xca, 0xd0, 0x8e,
	0x04, 0xe4, 0x1d, 0x90, 0xb3, 0xf4, 0xc9, 0xd7, 0xd6, 0xe2, 0xf5, 0xb9, 0xc9, 0x67, 0x39, 0xc9,
	0xc1, 0x9c, 0xeb, 0xf2, 0xad, 0x75, 0xf9, 0xfb, 0xf2, 0x1e, 0xd6, 0xa6, 0x0f, 0x3f, 0xbd, 0xe7,
	0xdf, 0x2f, 0x40, 0x7b, 0x7b, 0x0a, 0x3c, 0xd9, 0xd0, 0xad, 0x0f, 0x50, 0x57, 0x66, 0xd6, 0xc0,
	0xc9, 0x1b, 0xf9, 0xe9, 0xe5, 0xa5, 0x5e, 0xcf, 0xc3, 0xca, 0xe8, 0x81, 0x7a, 0xfe, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x7d, 0x53, 0xdc, 0xe0, 0x7d, 0x05, 0x00, 0x00,
}