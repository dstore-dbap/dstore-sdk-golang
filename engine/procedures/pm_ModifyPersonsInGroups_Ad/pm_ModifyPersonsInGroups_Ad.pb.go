// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonsInGroups_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonsInGroups_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonsInGroups_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonsInGroups_Ad

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
	PersonId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull bool                        `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	GroupId      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	GroupIdNull  bool                        `protobuf:"varint,1002,opt,name=group_id_null,json=groupIdNull" json:"group_id_null,omitempty"`
	Revoke       *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=revoke" json:"revoke,omitempty"`
	RevokeNull   bool                        `protobuf:"varint,1003,opt,name=revoke_null,json=revokeNull" json:"revoke_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *Parameters) GetRevoke() *dstore_values.BooleanValue {
	if m != nil {
		return m.Revoke
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonsInGroups_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonsInGroups_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonsInGroups_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyPersonsInGroups_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x66, 0x5b, 0xbb, 0xad, 0xa9, 0x7f, 0x44, 0x90, 0xda, 0xa2, 0x2c, 0xbb, 0x08, 0x7a, 0x93,
	0x8a, 0x0b, 0x2a, 0x82, 0x17, 0x0a, 0x22, 0x73, 0xd1, 0x65, 0xc9, 0x85, 0xa0, 0x08, 0xc3, 0xac,
	0x39, 0x3b, 0x0c, 0xdb, 0xe6, 0x0c, 0xc9, 0x74, 0x8b, 0x6f, 0xe1, 0x2b, 0xf8, 0x1c, 0x3e, 0x91,
	0xfa, 0x12, 0x9e, 0xcc, 0xc9, 0x54, 0x67, 0x14, 0x75, 0x6f, 0x66, 0x72, 0x72, 0xbe, 0xef, 0xcb,
	0xc9, 0xf9, 0x4e, 0xc4, 0x33, 0xe3, 0x2b, 0x74, 0x30, 0x07, 0x9b, 0x17, 0x16, 0xe6, 0xa5, 0xc3,
	0x0f, 0x60, 0xd6, 0x0e, 0xfc, 0xbc, 0x5c, 0xa5, 0x0b, 0x34, 0xc5, 0xe9, 0xc7, 0x63, 0x70, 0x1e,
	0xad, 0x4f, 0xec, 0x6b, 0x87, 0xeb, 0xd2, 0xa7, 0x2f, 0x8c, 0x22, 0x58, 0x85, 0xf2, 0x01, 0x73,
	0x15, 0x73, 0xd5, 0x5f, 0x08, 0xd3, 0x9b, 0xf1, 0x98, 0xf3, 0x6c, 0xb9, 0x06, 0xcf, 0xfc, 0xe9,
	0xed, 0xf6, 0xd9, 0xe0, 0x1c, 0xba, 0x98, 0x9a, 0xb5, 0x53, 0x2b, 0xf0, 0x3e, 0xcb, 0x21, 0x26,
	0x0f, 0xba, 0xc9, 0x2a, 0x2b, 0xec, 0x29, 0xba, 0x55, 0x56, 0x15, 0x68, 0x19, 0xb4, 0xff, 0xb9,
	0x27, 0xc4, 0x71, 0xe6, 0x32, 0xca, 0x52, 0x39, 0xf2, 0xa9, 0xb8, 0x5c, 0xd6, 0x65, 0xa5, 0x85,
	0x99, 0xec, 0xec, 0xed, 0xdc, 0x1f, 0x3f, 0x9a, 0xa9, 0x58, 0x7f, 0x2c, 0xaa, 0xb0, 0x15, 0xe4,
	0xe0, 0xde, 0x84, 0x48, 0x8f, 0x18, 0x9d, 0x18, 0x79, 0x4f, 0x5c, 0xdb, 0x32, 0x53, 0xbb, 0x5e,
	0x2e, 0x27, 0x5f, 0x87, 0xc4, 0x1f, 0xe9, 0x2b, 0x0d, 0xe4, 0x88, 0x36, 0xe5, 0x63, 0x31, 0xca,
	0xc3, 0x75, 0x83, 0x7e, 0xef, 0xdf, 0xfa, 0xc3, 0x1a, 0x4c, 0xf2, 0x07, 0xe2, 0x6a, 0xc3, 0x63,
	0xf5, 0x6f, 0xac, 0x3e, 0x8e, 0x80, 0x5a, 0xfc, 0x50, 0xec, 0x3a, 0x38, 0xc7, 0x33, 0x98, 0xf4,
	0xff, 0x28, 0x7d, 0x82, 0xb8, 0x84, 0xcc, 0xb2, 0x74, 0x84, 0xca, 0x3d, 0x31, 0xe6, 0x15, 0xeb,
	0x7e, 0x67, 0x5d, 0xc1, 0x7b, 0x41, 0x76, 0xff, 0x4b, 0x4f, 0x8c, 0x34, 0xf8, 0x92, 0xcc, 0x02,
	0xf9, 0x50, 0x0c, 0x6a, 0x07, 0x62, 0x77, 0xa6, 0xaa, 0xed, 0x2e, 0xbb, 0xf3, 0x2a, 0x7c, 0x35,
	0x03, 0xe5, 0x5b, 0x71, 0x23, 0xf4, 0x3e, 0xfd, 0xa5, 0xf9, 0x74, 0xf5, 0x3e, 0x91, 0x55, 0x87,
	0xdc, 0xb5, 0x68, 0x41, 0x71, 0xf2, 0x33, 0xd6, 0xd7, 0x57, 0xed, 0x0d, 0xb2, 0x6b, 0x18, 0x3d,
	0xa7, 0x1b, 0x07, 0xc5, 0xbb, 0xbf, 0x29, 0xf2, 0x44, 0x2c, 0xf8, 0xaf, 0x1b, 0xb8, 0x4c, 0x44,
	0xdf, 0xe1, 0x66, 0x72, 0xa9, 0x66, 0x3d, 0x51, 0xff, 0x3d, 0xa2, 0xaa, 0x69, 0x84, 0xd2, 0xb8,
	0xd1, 0x41, 0x63, 0x7a, 0x47, 0xf4, 0x69, 0x2d, 0x6f, 0x51, 0xf3, 0x71, 0x13, 0x7c, 0xfd, 0x74,
	0x44, 0xad, 0x19, 0xe8, 0x01, 0x85, 0x89, 0x79, 0xf9, 0x5e, 0xcc, 0x0a, 0xec, 0x1e, 0xb0, 0x7d,
	0x3f, 0xef, 0x9e, 0xe7, 0xe8, 0xcd, 0x59, 0x93, 0x37, 0x17, 0x7c, 0x62, 0x27, 0xbb, 0xf5, 0x18,
	0x1f, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x63, 0x5d, 0x83, 0xc5, 0xa1, 0x03, 0x00, 0x00,
}
