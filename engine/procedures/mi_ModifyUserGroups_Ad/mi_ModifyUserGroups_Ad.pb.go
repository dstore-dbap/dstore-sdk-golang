// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyUserGroups_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyUserGroups_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyUserGroups_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyUserGroups_Ad

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
	UserGroupId          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	UserGroupIdNull      bool                        `protobuf:"varint,1001,opt,name=user_group_id_null,json=userGroupIdNull" json:"user_group_id_null,omitempty"`
	GroupName            *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	GroupNameNull        bool                        `protobuf:"varint,1002,opt,name=group_name_null,json=groupNameNull" json:"group_name_null,omitempty"`
	GroupDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=group_description,json=groupDescription" json:"group_description,omitempty"`
	GroupDescriptionNull bool                        `protobuf:"varint,1003,opt,name=group_description_null,json=groupDescriptionNull" json:"group_description_null,omitempty"`
	DeleteGroup          *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_group,json=deleteGroup" json:"delete_group,omitempty"`
	DeleteGroupNull      bool                        `protobuf:"varint,1004,opt,name=delete_group_null,json=deleteGroupNull" json:"delete_group_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUserGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserGroupId
	}
	return nil
}

func (m *Parameters) GetGroupName() *dstore_values.StringValue {
	if m != nil {
		return m.GroupName
	}
	return nil
}

func (m *Parameters) GetGroupDescription() *dstore_values.StringValue {
	if m != nil {
		return m.GroupDescription
	}
	return nil
}

func (m *Parameters) GetDeleteGroup() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteGroup
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyUserGroups_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyUserGroups_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyUserGroups_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifyUserGroups_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xed, 0x6a, 0x13, 0x41,
	0x14, 0xa5, 0x5d, 0xd3, 0xd6, 0x5b, 0x4b, 0xda, 0x51, 0xca, 0x9a, 0xa0, 0x48, 0x15, 0x14, 0x94,
	0x8d, 0x54, 0x0a, 0xf5, 0x8f, 0xa2, 0x54, 0x4a, 0x7e, 0x24, 0xc8, 0x80, 0x8a, 0xfe, 0x59, 0xb6,
	0x9d, 0xdb, 0x65, 0x70, 0x33, 0x13, 0x66, 0x76, 0x2d, 0xbe, 0x85, 0x4f, 0xe3, 0x53, 0xf8, 0x22,
	0x7e, 0x3c, 0x84, 0x33, 0x7b, 0x27, 0x4d, 0x76, 0x5b, 0x68, 0xfe, 0x24, 0x7b, 0xf7, 0x9e, 0x73,
	0xcf, 0x9d, 0xb3, 0x67, 0xe0, 0x40, 0xd8, 0x52, 0x1b, 0x1c, 0xa0, 0xca, 0xa5, 0xc2, 0xc1, 0xd4,
	0xe8, 0x53, 0x14, 0x95, 0x41, 0x3b, 0x98, 0xc8, 0x74, 0xa4, 0x85, 0x3c, 0xfb, 0xfe, 0xc1, 0xa2,
	0x39, 0x36, 0xba, 0x9a, 0xda, 0xf4, 0x8d, 0x48, 0x1c, 0xa2, 0xd4, 0xec, 0x11, 0xd1, 0x12, 0xa2,
	0x25, 0x57, 0x63, 0x7b, 0xb7, 0xc3, 0xf0, 0x6f, 0x59, 0x51, 0xa1, 0x25, 0x6a, 0xef, 0x6e, 0x53,
	0x11, 0x8d, 0xd1, 0x26, 0xb4, 0xfa, 0xcd, 0xd6, 0x04, 0xad, 0xcd, 0x72, 0x0c, 0xcd, 0x87, 0xed,
	0x66, 0x99, 0x49, 0x75, 0xa6, 0xcd, 0x24, 0x2b, 0xa5, 0x56, 0x04, 0xda, 0xfb, 0x15, 0x01, 0xbc,
	0xcf, 0x4c, 0xe6, 0xba, 0x68, 0x2c, 0x7b, 0x0d, 0x5b, 0x95, 0xdb, 0x28, 0xcd, 0xfd, 0x4a, 0xa9,
	0x14, 0xf1, 0xca, 0x83, 0x95, 0x27, 0x9b, 0xfb, 0xfd, 0x24, 0xac, 0x1f, 0x16, 0x93, 0xaa, 0xc4,
	0x1c, 0xcd, 0x47, 0x5f, 0xf1, 0xcd, 0x6a, 0x76, 0x86, 0xa1, 0x60, 0xcf, 0x80, 0x35, 0x06, 0xa4,
	0xaa, 0x2a, 0x8a, 0xf8, 0xf7, 0xba, 0x1b, 0xb3, 0xc1, 0xbb, 0x0b, 0xc8, 0xb1, 0x7b, 0xcf, 0x5e,
	0x02, 0x10, 0x50, 0xb9, 0x05, 0xe2, 0xd5, 0x5a, 0xab, 0xd7, 0xd2, 0xb2, 0xa5, 0x91, 0x2a, 0x27,
	0xa9, 0x9b, 0x35, 0x7a, 0xec, 0xc0, 0xec, 0x31, 0x74, 0xe7, 0x54, 0x52, 0xf9, 0x43, 0x2a, 0x5b,
	0x17, 0xa0, 0x5a, 0xe3, 0x18, 0x76, 0x08, 0x28, 0xd0, 0x9e, 0x1a, 0x39, 0xf5, 0x87, 0x8f, 0xa3,
	0x6b, 0xa5, 0xb6, 0x6b, 0xd2, 0xd1, 0x9c, 0xc3, 0x0e, 0x60, 0xf7, 0xd2, 0x20, 0x12, 0xfe, 0x4b,
	0xc2, 0x77, 0xda, 0x94, 0x5a, 0xff, 0x15, 0xdc, 0x12, 0x58, 0x38, 0x7b, 0xc9, 0x93, 0xf8, 0xc6,
	0x95, 0x8e, 0x9e, 0x68, 0x5d, 0x60, 0xa6, 0x82, 0xa3, 0x44, 0xa8, 0x9d, 0x62, 0x4f, 0x61, 0x67,
	0x91, 0x4f, 0x8a, 0xff, 0x82, 0xa1, 0x0b, 0x40, 0x2f, 0xb6, 0xf7, 0x73, 0x15, 0x36, 0x38, 0xda,
	0xa9, 0x56, 0x16, 0xd9, 0x73, 0xe8, 0xd4, 0x61, 0x09, 0x1f, 0xf1, 0xe2, 0xb4, 0x21, 0x83, 0x14,
	0xa4, 0x77, 0xfe, 0x97, 0x13, 0x90, 0x7d, 0x86, 0x6d, 0x1f, 0x93, 0x74, 0x21, 0x27, 0xee, 0xab,
	0x44, 0x8e, 0x9c, 0xb4, 0xc8, 0xed, 0x34, 0x8d, 0x5c, 0x3d, 0x9c, 0xd7, 0xbc, 0x3b, 0x69, 0xbe,
	0x60, 0x87, 0xb0, 0x1e, 0xe2, 0xe9, 0xcc, 0xf7, 0x13, 0xef, 0x5f, 0x9a, 0x48, 0xe1, 0x1d, 0xd1,
	0x3f, 0x9f, 0xc1, 0xd9, 0x11, 0x44, 0x46, 0x9f, 0x3b, 0xdf, 0x3c, 0x6b, 0x3f, 0x59, 0xe6, 0x22,
	0x25, 0x33, 0x0f, 0x12, 0xae, 0xcf, 0xb9, 0xa7, 0xf7, 0xee, 0x41, 0xe4, 0x9e, 0xd9, 0x2e, 0xac,
	0xb9, 0xca, 0x27, 0xfb, 0xc7, 0xd8, 0xb9, 0xd2, 0xe1, 0x1d, 0x57, 0x0e, 0xc5, 0xdb, 0x4f, 0xd0,
	0x97, 0xba, 0x35, 0x7b, 0x7e, 0xb7, 0xbf, 0x1c, 0xe6, 0xda, 0x8a, 0xaf, 0xb3, 0xbe, 0x58, 0xfe,
	0xfa, 0x9f, 0xac, 0xd5, 0xf7, 0xec, 0xc5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xb1, 0x82,
	0xfe, 0x38, 0x04, 0x00, 0x00,
}
