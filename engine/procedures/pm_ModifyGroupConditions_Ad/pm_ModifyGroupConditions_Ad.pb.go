// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyGroupConditions_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyGroupConditions_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyGroupConditions_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyGroupConditions_Ad

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
	ConditionId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
	ConditionIdNull          bool                        `protobuf:"varint,1001,opt,name=condition_id_null,json=conditionIdNull" json:"condition_id_null,omitempty"`
	ConditionDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=condition_description,json=conditionDescription" json:"condition_description,omitempty"`
	ConditionDescriptionNull bool                        `protobuf:"varint,1002,opt,name=condition_description_null,json=conditionDescriptionNull" json:"condition_description_null,omitempty"`
	DeleteGroupCondition     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete_group_condition,json=deleteGroupCondition" json:"delete_group_condition,omitempty"`
	DeleteGroupConditionNull bool                        `protobuf:"varint,1003,opt,name=delete_group_condition_null,json=deleteGroupConditionNull" json:"delete_group_condition_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
	}
	return nil
}

func (m *Parameters) GetConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ConditionDescription
	}
	return nil
}

func (m *Parameters) GetDeleteGroupCondition() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteGroupCondition
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyGroupConditions_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyGroupConditions_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyGroupConditions_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyGroupConditions_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xa6, 0x8d, 0xdb, 0x96, 0x53, 0xa1, 0x3a, 0x6a, 0x89, 0x59, 0x14, 0xa9, 0x37, 0x8a, 0x30,
	0x2b, 0x7a, 0xa1, 0x08, 0x15, 0xfc, 0x43, 0xf6, 0x62, 0xab, 0xce, 0x85, 0xa0, 0x08, 0x21, 0xdd,
	0x39, 0x0d, 0x83, 0xd9, 0x99, 0x30, 0x93, 0x75, 0xf1, 0x2d, 0x7c, 0x00, 0xdf, 0xc4, 0x27, 0x52,
	0x5f, 0xc2, 0x99, 0xcc, 0x24, 0xbb, 0x89, 0x41, 0xec, 0xcd, 0x6e, 0x4e, 0xce, 0xf9, 0xce, 0xf9,
	0xf2, 0x9d, 0xef, 0xc0, 0x13, 0x6e, 0x2a, 0xa5, 0x71, 0x82, 0x32, 0x17, 0x12, 0x27, 0xa5, 0x56,
	0x73, 0xe4, 0x4b, 0x8d, 0x66, 0x52, 0x2e, 0xd2, 0x99, 0xe2, 0xe2, 0xec, 0xeb, 0x6b, 0xad, 0x96,
	0xe5, 0x0b, 0x25, 0xb9, 0xa8, 0x84, 0x92, 0x26, 0x7d, 0xc6, 0xa9, 0x2d, 0xab, 0x14, 0xb9, 0xeb,
	0xb1, 0xd4, 0x63, 0xe9, 0x3f, 0x00, 0xc9, 0x95, 0x30, 0xe6, 0x4b, 0x56, 0x2c, 0xd1, 0x78, 0x7c,
	0x72, 0xbd, 0x3b, 0x1b, 0xb5, 0x56, 0x3a, 0xa4, 0xc6, 0xdd, 0xd4, 0x02, 0x8d, 0xc9, 0x72, 0x0c,
	0xc9, 0xdb, 0xfd, 0x64, 0x95, 0x09, 0x79, 0xa6, 0xf4, 0x22, 0x73, 0xf3, 0x7c, 0xd1, 0xd1, 0xf7,
	0x08, 0xe0, 0x6d, 0xa6, 0x33, 0x9b, 0x45, 0x6d, 0xc8, 0x53, 0xb8, 0x38, 0x6f, 0x18, 0xa5, 0x82,
	0xc7, 0x5b, 0xb7, 0xb6, 0xee, 0xec, 0x3f, 0x18, 0xd3, 0xf0, 0x09, 0x81, 0x97, 0x90, 0x15, 0xe6,
	0xa8, 0xdf, 0xbb, 0x88, 0xed, 0xb7, 0x80, 0x29, 0x27, 0xf7, 0xe0, 0xf2, 0x26, 0x3e, 0x95, 0xcb,
	0xa2, 0x88, 0x7f, 0xee, 0xda, 0x2e, 0x7b, 0xec, 0x60, 0xa3, 0xf0, 0xc4, 0xbe, 0x27, 0x6f, 0xe0,
	0xda, 0xba, 0x98, 0xa3, 0x99, 0x6b, 0x51, 0xba, 0xe7, 0x78, 0xbb, 0x9e, 0x9a, 0xf4, 0xa6, 0x9a,
	0x4a, 0x0b, 0x99, 0xfb, 0xa1, 0x57, 0x5b, 0xe0, 0xcb, 0x35, 0x8e, 0x1c, 0x43, 0x32, 0xd8, 0xd0,
	0xd3, 0xf8, 0xe5, 0x69, 0xc4, 0x43, 0xd0, 0x9a, 0xcf, 0x3b, 0x38, 0xe4, 0x58, 0x58, 0x21, 0xd2,
	0xdc, 0xad, 0x26, 0x6d, 0x0b, 0xe3, 0x68, 0x50, 0x86, 0x53, 0xa5, 0x0a, 0xcc, 0x64, 0x60, 0xe4,
	0xa1, 0xdd, 0xa5, 0x5a, 0x3d, 0xc7, 0xc3, 0x2d, 0x3d, 0xa5, 0xdf, 0x81, 0xd2, 0x10, 0xd6, 0x51,
	0x3a, 0xfa, 0xb1, 0x0d, 0x7b, 0x0c, 0x4d, 0x69, 0xfd, 0x81, 0xe4, 0x3e, 0x8c, 0xea, 0xe5, 0x87,
	0xad, 0xb4, 0xfa, 0x04, 0x63, 0x79, 0x63, 0xbc, 0x72, 0xbf, 0xcc, 0x17, 0x92, 0x0f, 0x70, 0xc9,
	0xad, 0x3d, 0xdd, 0xd8, 0xbb, 0x15, 0x37, 0xb2, 0x60, 0xda, 0x03, 0xf7, 0xdd, 0x31, 0xb3, 0xf1,
	0x74, 0x1d, 0xb3, 0x83, 0x45, 0xf7, 0x05, 0x79, 0x0c, 0xbb, 0xc1, 0x6e, 0x56, 0x1d, 0xd7, 0xf1,
	0xe6, 0x5f, 0x1d, 0xbd, 0x19, 0x67, 0xfe, 0x9f, 0x35, 0xe5, 0x64, 0x0a, 0x91, 0x56, 0xab, 0xf8,
	0x42, 0x8d, 0x7a, 0x44, 0xff, 0xfb, 0x3a, 0x68, 0x23, 0x04, 0x65, 0x6a, 0xc5, 0x5c, 0x8f, 0xe4,
	0x06, 0x44, 0xf6, 0x99, 0x1c, 0xc2, 0x8e, 0x8d, 0x9c, 0x5f, 0xbf, 0x9d, 0x58, 0x69, 0x46, 0x6c,
	0x64, 0xc3, 0x29, 0x7f, 0xfe, 0x09, 0xc6, 0x42, 0xf5, 0x07, 0xb4, 0xa7, 0xfb, 0xf1, 0x38, 0x57,
	0x86, 0x7f, 0x6e, 0xf2, 0xfc, 0x9c, 0xd7, 0x7d, 0xba, 0x53, 0x5f, 0xd0, 0xc3, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0xee, 0x6f, 0x32, 0x1c, 0x04, 0x00, 0x00,
}