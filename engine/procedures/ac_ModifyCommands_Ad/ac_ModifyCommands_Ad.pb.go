// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/ac_ModifyCommands_Ad.proto
// DO NOT EDIT!

/*
Package ac_ModifyCommands_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/ac_ModifyCommands_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package ac_ModifyCommands_Ad

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
	CommandId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=command_id,json=commandId" json:"command_id,omitempty"`
	CommandIdNull          bool                        `protobuf:"varint,1001,opt,name=command_id_null,json=commandIdNull" json:"command_id_null,omitempty"`
	Command                *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	CommandNull            bool                        `protobuf:"varint,1002,opt,name=command_null,json=commandNull" json:"command_null,omitempty"`
	CommandDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=command_description,json=commandDescription" json:"command_description,omitempty"`
	CommandDescriptionNull bool                        `protobuf:"varint,1003,opt,name=command_description_null,json=commandDescriptionNull" json:"command_description_null,omitempty"`
	DeleteCommand          *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_command,json=deleteCommand" json:"delete_command,omitempty"`
	DeleteCommandNull      bool                        `protobuf:"varint,1004,opt,name=delete_command_null,json=deleteCommandNull" json:"delete_command_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommandId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommandId
	}
	return nil
}

func (m *Parameters) GetCommand() *dstore_values.StringValue {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Parameters) GetCommandDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CommandDescription
	}
	return nil
}

func (m *Parameters) GetDeleteCommand() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCommand
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.ac_ModifyCommands_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.ac_ModifyCommands_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.ac_ModifyCommands_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/ac_ModifyCommands_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xd6, 0x16, 0xba, 0x8e, 0x33, 0xc6, 0xc0, 0x95, 0xa6, 0xd0, 0x0a, 0x84, 0xca, 0x05, 0x5c,
	0xa5, 0x13, 0x43, 0x08, 0xb8, 0x63, 0xc0, 0x45, 0x85, 0x3a, 0xa1, 0x48, 0x20, 0xc1, 0x4d, 0xe4,
	0xd5, 0x67, 0x91, 0x45, 0x62, 0x57, 0x76, 0xca, 0xc4, 0x5b, 0xf0, 0x30, 0x3c, 0x04, 0xaf, 0xc2,
	0xcf, 0x43, 0x60, 0xe7, 0x38, 0x2b, 0xc9, 0x2a, 0x95, 0x9b, 0x36, 0xf6, 0xf9, 0x7e, 0x8e, 0x8f,
	0x3f, 0xc3, 0xb1, 0xb0, 0x95, 0x36, 0x38, 0x41, 0x95, 0x4b, 0x85, 0x93, 0x85, 0xd1, 0x73, 0x14,
	0x4b, 0x83, 0x76, 0xc2, 0xe7, 0xd9, 0x4c, 0x0b, 0x79, 0xfe, 0xf5, 0x95, 0x2e, 0x4b, 0xae, 0x84,
	0xcd, 0x5e, 0x8a, 0xc4, 0xd5, 0x2b, 0xcd, 0xc6, 0x44, 0x4a, 0x88, 0x94, 0xac, 0x43, 0x0e, 0x07,
	0x41, 0xf8, 0x0b, 0x2f, 0x96, 0x68, 0x89, 0x38, 0xbc, 0xd3, 0x76, 0x43, 0x63, 0xb4, 0x09, 0xa5,
	0x51, 0xbb, 0x54, 0xa2, 0xb5, 0x3c, 0xc7, 0x50, 0x7c, 0xd0, 0x2d, 0x56, 0x5c, 0xaa, 0x73, 0x6d,
	0x4a, 0x5e, 0x49, 0xad, 0x08, 0x34, 0xfe, 0x11, 0x01, 0xbc, 0xe3, 0x86, 0xbb, 0x2a, 0x1a, 0xcb,
	0x5e, 0x00, 0xcc, 0xa9, 0x9f, 0x4c, 0x8a, 0x78, 0xeb, 0xfe, 0xd6, 0xa3, 0xbd, 0xc7, 0xa3, 0x24,
	0x74, 0x1e, 0xba, 0x92, 0xaa, 0xc2, 0x1c, 0xcd, 0x07, 0xbf, 0x4a, 0xaf, 0x07, 0xf8, 0x54, 0xb0,
	0x87, 0x70, 0xb0, 0xe2, 0x66, 0x6a, 0x59, 0x14, 0xf1, 0xcf, 0xbe, 0x53, 0xd8, 0x4d, 0xf7, 0x2f,
	0x41, 0xa7, 0x6e, 0x97, 0x3d, 0x81, 0x7e, 0xd8, 0x88, 0xb7, 0x6b, 0x87, 0x61, 0xc7, 0xc1, 0x56,
	0x46, 0xaa, 0x9c, 0x0c, 0x1a, 0x28, 0x1b, 0xc3, 0x8d, 0x46, 0xbe, 0xd6, 0xfe, 0x45, 0xda, 0x7b,
	0x61, 0xb3, 0x56, 0x7e, 0x0b, 0x83, 0x06, 0x23, 0xd0, 0xce, 0x8d, 0x5c, 0xf8, 0xa3, 0xc6, 0xd1,
	0x46, 0x17, 0x16, 0x68, 0xaf, 0x57, 0x2c, 0xf6, 0x1c, 0xe2, 0x35, 0x62, 0x64, 0xfe, 0x9b, 0xcc,
	0x0f, 0xaf, 0xd2, 0xea, 0x3e, 0x4e, 0xe0, 0xa6, 0xc0, 0xc2, 0x8d, 0x34, 0x6b, 0x0e, 0x7a, 0x6d,
	0xed, 0x28, 0xcf, 0xb4, 0x2e, 0x90, 0x2b, 0xea, 0x61, 0x9f, 0x28, 0x21, 0x0f, 0x6c, 0x02, 0x83,
	0xb6, 0x06, 0x39, 0xff, 0x21, 0xe7, 0xdb, 0x2d, 0xb0, 0x37, 0x1d, 0x7f, 0xdf, 0x86, 0xdd, 0x14,
	0xed, 0x42, 0x2b, 0x8b, 0xec, 0x08, 0x7a, 0x75, 0x50, 0xc2, 0x1d, 0x5e, 0x9e, 0x3d, 0xa4, 0x8f,
	0x42, 0xf4, 0xc6, 0xff, 0xa6, 0x04, 0x64, 0x1f, 0xe1, 0x96, 0x8f, 0x48, 0xf6, 0x4f, 0x46, 0xdc,
	0xf5, 0x44, 0x8e, 0x9c, 0x74, 0xc8, 0xdd, 0x24, 0xcd, 0xdc, 0x7a, 0xba, 0x5a, 0xa7, 0x07, 0x65,
	0x7b, 0x83, 0x3d, 0x83, 0x7e, 0x88, 0xa6, 0xbb, 0x0a, 0xaf, 0x78, 0xef, 0x8a, 0x22, 0x05, 0x77,
	0x46, 0xff, 0x69, 0x03, 0x77, 0x83, 0x8c, 0x8c, 0xbe, 0x70, 0xd3, 0xf3, 0xac, 0xa3, 0x64, 0xf3,
	0x13, 0x4a, 0x9a, 0x09, 0x24, 0xa9, 0xbe, 0x48, 0x3d, 0x79, 0x78, 0x17, 0x22, 0xf7, 0xcd, 0x0e,
	0x61, 0xc7, 0xad, 0x7c, 0xac, 0xbf, 0x9d, 0xba, 0x99, 0xf4, 0xd2, 0x9e, 0x5b, 0x4e, 0xc5, 0xc9,
	0x7b, 0x18, 0x49, 0xdd, 0x51, 0x5e, 0xbd, 0xe8, 0x4f, 0x4f, 0x73, 0x6d, 0xc5, 0xe7, 0xa6, 0x2e,
	0xfe, 0xf7, 0xd1, 0x9f, 0xed, 0xd4, 0xef, 0xeb, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73,
	0x7b, 0xee, 0xb9, 0x2c, 0x04, 0x00, 0x00,
}
