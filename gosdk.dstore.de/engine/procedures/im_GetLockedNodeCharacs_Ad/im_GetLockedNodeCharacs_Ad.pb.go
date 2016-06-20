// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetLockedNodeCharacs_Ad.proto
// DO NOT EDIT!

/*
Package im_GetLockedNodeCharacs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetLockedNodeCharacs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetLockedNodeCharacs_Ad

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
	UserId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull               bool                        `protobuf:"varint,1001,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	NodeCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull bool                        `protobuf:"varint,1002,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	LockStatus               *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=lock_status,json=lockStatus" json:"lock_status,omitempty"`
	LockStatusNull           bool                        `protobuf:"varint,1003,opt,name=lock_status_null,json=lockStatusNull" json:"lock_status_null,omitempty"`
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

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetLockStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.LockStatus
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
	UserName                  *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	CharacteristicDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	UserId                    *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	LockStatus                *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=lock_status,json=lockStatus" json:"lock_status,omitempty"`
	NodeCharacteristicId      *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Response_Row) GetLockStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.LockStatus
	}
	return nil
}

func (m *Response_Row) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetLockedNodeCharacs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetLockedNodeCharacs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetLockedNodeCharacs_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetLockedNodeCharacs_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x6f, 0xd4, 0x3c,
	0x10, 0x55, 0x77, 0xbf, 0xbd, 0x7c, 0x53, 0x04, 0x95, 0x41, 0x55, 0x9a, 0x95, 0x10, 0x94, 0x97,
	0xf2, 0x92, 0x45, 0xdc, 0x54, 0xa4, 0x82, 0xc4, 0x4d, 0x50, 0x89, 0x46, 0x60, 0x24, 0x24, 0xca,
	0x43, 0x14, 0x62, 0x13, 0xac, 0xee, 0xda, 0x2b, 0xdb, 0xa1, 0x7f, 0x83, 0xeb, 0x13, 0xbf, 0x8e,
	0x47, 0xe0, 0x37, 0x20, 0x31, 0x8e, 0xb3, 0xa4, 0x09, 0xad, 0xb6, 0xe5, 0x65, 0x57, 0x93, 0x39,
	0x67, 0x66, 0x3c, 0xe7, 0xd8, 0x70, 0x8b, 0x19, 0xab, 0x34, 0x1f, 0x73, 0x99, 0x0b, 0xc9, 0xc7,
	0x33, 0xad, 0x32, 0xce, 0x0a, 0xcd, 0xcd, 0x58, 0x4c, 0x93, 0x47, 0xdc, 0x3e, 0x51, 0xd9, 0x1e,
	0x67, 0xb1, 0x62, 0xfc, 0xfe, 0xdb, 0x54, 0xa7, 0x99, 0x49, 0xee, 0xb2, 0x08, 0x51, 0x56, 0x91,
	0x0d, 0x4f, 0x8d, 0x3c, 0x35, 0x3a, 0x1a, 0x1f, 0x9e, 0xad, 0x9a, 0xbc, 0x4b, 0x27, 0x05, 0x37,
	0x9e, 0x1e, 0xae, 0x35, 0x3b, 0x73, 0xad, 0x95, 0xae, 0x52, 0xa3, 0x66, 0x6a, 0xca, 0x8d, 0x49,
	0x73, 0x5e, 0x25, 0x2f, 0xb5, 0x93, 0x36, 0x15, 0xf2, 0x8d, 0xd2, 0xd3, 0xd4, 0x0a, 0x25, 0x3d,
	0x68, 0xfd, 0x5b, 0x07, 0xe0, 0x29, 0xf6, 0xc7, 0x2c, 0xd7, 0x86, 0x5c, 0x87, 0x41, 0x61, 0xb8,
	0x4e, 0x04, 0x0b, 0x96, 0x2e, 0x2c, 0x6d, 0x2c, 0x5f, 0x1d, 0x45, 0xd5, 0xf0, 0xd5, 0x48, 0x42,
	0x5a, 0x9e, 0x73, 0xfd, 0xc2, 0x45, 0xb4, 0xef, 0xb0, 0xdb, 0x8c, 0x5c, 0x84, 0x53, 0x15, 0x2b,
	0x91, 0xc5, 0x64, 0x12, 0x7c, 0x1f, 0x20, 0x77, 0x48, 0xc1, 0xa7, 0x63, 0xfc, 0x44, 0x9e, 0xc1,
	0xaa, 0xc4, 0xb3, 0x26, 0x59, 0x79, 0x58, 0xec, 0x25, 0x8c, 0x15, 0x99, 0xeb, 0xd3, 0x59, 0xdc,
	0xe7, 0x9c, 0xfc, 0xb3, 0xa6, 0x39, 0x13, 0xbb, 0xde, 0x81, 0xd1, 0xe1, 0x25, 0xfd, 0x10, 0x3f,
	0xfc, 0x10, 0xc1, 0x61, 0xdc, 0x72, 0xa4, 0x2d, 0x58, 0x9e, 0xa0, 0x08, 0x89, 0xb1, 0xa9, 0x2d,
	0x4c, 0xd0, 0x5d, 0x3c, 0x07, 0x38, 0xfc, 0xf3, 0x12, 0x4e, 0x2e, 0xc3, 0xca, 0x01, 0xb6, 0x6f,
	0xf9, 0xd3, 0xb7, 0x3c, 0x5d, 0xc3, 0x5c, 0xa3, 0xf5, 0xaf, 0x3d, 0x18, 0x52, 0x6e, 0x66, 0x4a,
	0x1a, 0x4e, 0xae, 0x40, 0xaf, 0x54, 0xb0, 0xda, 0x6f, 0x18, 0x35, 0xcd, 0xe1, 0xd5, 0x7d, 0xe8,
	0x7e, 0xa9, 0x07, 0x92, 0x97, 0xb0, 0xe2, 0xb4, 0x4b, 0x0e, 0x88, 0x87, 0x4b, 0xeb, 0x22, 0x39,
	0x6a, 0x91, 0xdb, 0x12, 0xef, 0x60, 0xbc, 0x5d, 0xc7, 0xf4, 0xcc, 0xb4, 0xf9, 0x81, 0x6c, 0xc2,
	0xa0, 0xf2, 0x0c, 0x1e, 0xdf, 0x55, 0x3c, 0xff, 0x57, 0x45, 0xef, 0xa8, 0x1d, 0xff, 0x4f, 0xe7,
	0x70, 0xf2, 0x18, 0xba, 0x5a, 0xed, 0x07, 0xff, 0x95, 0xac, 0x9b, 0xd1, 0x71, 0x1d, 0x1e, 0xcd,
	0xf7, 0x10, 0x51, 0xb5, 0x4f, 0x5d, 0x89, 0xf0, 0x57, 0x07, 0xba, 0x18, 0x90, 0x55, 0xe8, 0x63,
	0xe8, 0x1c, 0xf1, 0x3e, 0xc6, 0xd5, 0xf4, 0x68, 0x0f, 0x43, 0x94, 0x79, 0x13, 0xfe, 0x2f, 0xcd,
	0x25, 0xd1, 0xa3, 0xc1, 0x87, 0xb8, 0xb9, 0xb5, 0x4a, 0x25, 0x63, 0xb5, 0x90, 0xb9, 0x17, 0x69,
	0xe8, 0xd0, 0x31, 0x82, 0xc9, 0x2e, 0x84, 0x2d, 0x6f, 0x30, 0x6e, 0x32, 0x2d, 0x66, 0xe5, 0x0a,
	0x3f, 0x2e, 0x2e, 0xb5, 0xd6, 0xa4, 0x3f, 0xa8, 0xd9, 0xe4, 0x46, 0x7d, 0x51, 0x3e, 0xc5, 0xc7,
	0xbf, 0x29, 0xb7, 0x9b, 0x9e, 0xfb, 0x1c, 0x9f, 0xcc, 0x74, 0xf4, 0xc8, 0x5b, 0xf4, 0x25, 0xfe,
	0xc7, 0x6b, 0x74, 0xef, 0x15, 0x8c, 0x84, 0x6a, 0x09, 0x58, 0xbf, 0x6e, 0xbb, 0x5b, 0xb9, 0x32,
	0x6c, 0x6f, 0x9e, 0x67, 0x27, 0x7b, 0x00, 0x5f, 0xf7, 0xcb, 0x57, 0xe6, 0xda, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x98, 0x3b, 0xf1, 0x61, 0x3e, 0x05, 0x00, 0x00,
}