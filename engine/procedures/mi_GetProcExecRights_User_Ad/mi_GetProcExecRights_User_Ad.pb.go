// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetProcExecRights_User_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetProcExecRights_User_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetProcExecRights_User_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetProcExecRights_User_Ad

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
	UserId                  *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull              bool                        `protobuf:"varint,1001,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	ProcedureId             *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=procedure_id,json=procedureId" json:"procedure_id,omitempty"`
	ProcedureIdNull         bool                        `protobuf:"varint,1002,opt,name=procedure_id_null,json=procedureIdNull" json:"procedure_id_null,omitempty"`
	ProcedureCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=procedure_category_id,json=procedureCategoryId" json:"procedure_category_id,omitempty"`
	ProcedureCategoryIdNull bool                        `protobuf:"varint,1003,opt,name=procedure_category_id_null,json=procedureCategoryIdNull" json:"procedure_category_id_null,omitempty"`
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

func (m *Parameters) GetProcedureId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureId
	}
	return nil
}

func (m *Parameters) GetProcedureCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureCategoryId
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
	RowId          int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ExecutionRight *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=execution_right,json=executionRight" json:"execution_right,omitempty"`
	ProcedureId    *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=procedure_id,json=procedureId" json:"procedure_id,omitempty"`
	UserName       *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	ProcedureName  *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	UserId         *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetExecutionRight() *dstore_values.IntegerValue {
	if m != nil {
		return m.ExecutionRight
	}
	return nil
}

func (m *Response_Row) GetProcedureId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureId
	}
	return nil
}

func (m *Response_Row) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Response_Row) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetProcExecRights_User_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetProcExecRights_User_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetProcExecRights_User_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetProcExecRights_User_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0x6d, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xe9, 0x5d, 0xef, 0xc1, 0x69, 0xed, 0xe9, 0x16, 0x35, 0xe6, 0x40, 0xb4, 0xbe, 0x11,
	0x85, 0x9c, 0xf8, 0x00, 0x07, 0x4a, 0xc5, 0xca, 0x21, 0x15, 0x1a, 0xcb, 0x82, 0x82, 0x82, 0x84,
	0x78, 0x59, 0x63, 0xf0, 0x6e, 0xb7, 0xec, 0x26, 0x56, 0xbf, 0x85, 0x8f, 0x5f, 0xce, 0x77, 0xbe,
	0x54, 0xbf, 0x81, 0xaf, 0x9c, 0xcd, 0xee, 0x25, 0x97, 0x58, 0x7a, 0xbd, 0x37, 0x09, 0xb3, 0x33,
	0xbf, 0xff, 0x0c, 0x33, 0xb3, 0x0b, 0xf7, 0x22, 0x95, 0x0a, 0xc9, 0x06, 0x8c, 0xc7, 0x09, 0x67,
	0x83, 0x03, 0x29, 0xc6, 0x2c, 0xca, 0x24, 0x53, 0x83, 0x69, 0x12, 0x3c, 0x66, 0xe9, 0x3e, 0x1e,
	0x8c, 0x3e, 0xb0, 0x31, 0x4d, 0xe2, 0xb7, 0xa9, 0x0a, 0x9e, 0x29, 0x26, 0x83, 0x87, 0x91, 0x87,
	0x71, 0xa9, 0x20, 0xd7, 0x0d, 0xec, 0x19, 0xd8, 0x3b, 0x8e, 0x70, 0x37, 0x6d, 0xa2, 0xf7, 0xe1,
	0x24, 0x63, 0xca, 0x08, 0xb8, 0x17, 0xab, 0xd9, 0x99, 0x94, 0x42, 0x5a, 0x57, 0xbf, 0xea, 0x9a,
	0x32, 0xa5, 0xc2, 0x98, 0x59, 0xe7, 0xd5, 0xba, 0x33, 0x0d, 0x13, 0xfe, 0x46, 0xc8, 0x69, 0x98,
	0x26, 0x82, 0x9b, 0xa0, 0xad, 0x9f, 0x0d, 0x80, 0xfd, 0x50, 0x86, 0xe8, 0x65, 0x52, 0x91, 0x3b,
	0xd0, 0xc9, 0x74, 0x2d, 0x49, 0xe4, 0xac, 0x5c, 0x5e, 0xb9, 0xb6, 0x76, 0xab, 0xef, 0xd9, 0xf2,
	0x6d, 0x49, 0x09, 0x4f, 0x59, 0xcc, 0xe4, 0x73, 0x6d, 0xd1, 0xb6, 0x8e, 0xdd, 0x8d, 0xc8, 0x15,
	0x58, 0xb7, 0x54, 0xc0, 0xb3, 0xc9, 0xc4, 0xf9, 0xd5, 0x41, 0xb6, 0x4b, 0xc1, 0xb8, 0x7d, 0x3c,
	0x22, 0xdb, 0xb0, 0x5e, 0xb4, 0x4d, 0xab, 0x37, 0x16, 0xab, 0xaf, 0x15, 0x00, 0xa6, 0xb8, 0x01,
	0x67, 0xe7, 0x79, 0x93, 0xe7, 0xb7, 0xc9, 0xd3, 0x9b, 0x0b, 0xcc, 0x93, 0x3d, 0x85, 0x73, 0x65,
	0xf0, 0x38, 0x44, 0x4d, 0x21, 0x3f, 0xea, 0xac, 0xcd, 0xc5, 0x59, 0x37, 0x0b, 0xf2, 0x91, 0x05,
	0x31, 0xfb, 0x7d, 0x70, 0x8f, 0x14, 0x34, 0x65, 0xfc, 0x31, 0x65, 0x5c, 0x38, 0x82, 0xd4, 0xe5,
	0x6c, 0xfd, 0x5d, 0x85, 0x2e, 0x65, 0xea, 0x40, 0x70, 0xc5, 0xc8, 0x4d, 0x68, 0xe5, 0x13, 0xb4,
	0xfd, 0x75, 0xbd, 0xea, 0x7a, 0x98, 0xe9, 0x8e, 0xf4, 0x97, 0x9a, 0x40, 0xf2, 0x02, 0xce, 0xe8,
	0xd9, 0x05, 0x73, 0xc3, 0xc3, 0xf6, 0x35, 0x11, 0xf6, 0x6a, 0x70, 0x7d, 0xc4, 0x7b, 0x68, 0xef,
	0x96, 0x36, 0xed, 0x4d, 0xab, 0x07, 0x64, 0x08, 0x1d, 0xbb, 0x33, 0xd8, 0x1a, 0xad, 0x78, 0xe9,
	0x3f, 0x45, 0xb3, 0x51, 0x7b, 0xe6, 0x4f, 0x67, 0xe1, 0xe4, 0x09, 0x34, 0xa5, 0x38, 0x74, 0x56,
	0x73, 0x6a, 0xe8, 0x9d, 0x7c, 0xc7, 0xbd, 0x59, 0x27, 0x3c, 0x2a, 0x0e, 0xa9, 0x16, 0x71, 0x7f,
	0x34, 0xa0, 0x89, 0x06, 0x39, 0x0f, 0x6d, 0x34, 0xf5, 0x9c, 0x3e, 0xf9, 0xd8, 0x9c, 0x16, 0x6d,
	0xa1, 0x89, 0xdd, 0x1f, 0x41, 0x8f, 0xa1, 0x4e, 0xa6, 0x4b, 0x0e, 0xa4, 0x56, 0x73, 0x3e, 0xfb,
	0x8b, 0x27, 0xb9, 0x51, 0x40, 0x79, 0x05, 0xe4, 0x41, 0x6d, 0x05, 0xbf, 0xf8, 0x4b, 0xee, 0xe0,
	0x10, 0x4e, 0xe5, 0x6b, 0xce, 0xf1, 0xb6, 0x38, 0x5f, 0xfd, 0xea, 0xfc, 0x2c, 0xad, 0x52, 0x99,
	0xf0, 0xd8, 0xc0, 0x5d, 0x1d, 0xed, 0x63, 0x30, 0xd9, 0x81, 0x8d, 0x32, 0x75, 0x8e, 0x7f, 0x5b,
	0x8c, 0x9f, 0x2e, 0x90, 0x5c, 0xe3, 0x6e, 0x79, 0x35, 0xbf, 0xfb, 0x27, 0xbe, 0x9b, 0x3b, 0xaf,
	0xa0, 0x9f, 0x88, 0xda, 0x7c, 0xca, 0x07, 0xec, 0xe5, 0x76, 0x2c, 0x54, 0xf4, 0x6e, 0xe6, 0x8f,
	0x96, 0x7d, 0xe3, 0x5e, 0xb7, 0xf3, 0x67, 0xe4, 0xf6, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0xb7, 0xff, 0xc5, 0x23, 0x05, 0x00, 0x00,
}