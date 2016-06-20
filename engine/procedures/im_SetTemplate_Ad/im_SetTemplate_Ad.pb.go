// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_SetTemplate_Ad.proto
// DO NOT EDIT!

/*
Package im_SetTemplate_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_SetTemplate_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_SetTemplate_Ad

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
	TargetTemplateId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=target_template_id,json=targetTemplateId" json:"target_template_id,omitempty"`
	TargetTemplateIdNull bool                        `protobuf:"varint,1001,opt,name=target_template_id_null,json=targetTemplateIdNull" json:"target_template_id_null,omitempty"`
	SourceTemplateId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=source_template_id,json=sourceTemplateId" json:"source_template_id,omitempty"`
	SourceTemplateIdNull bool                        `protobuf:"varint,1002,opt,name=source_template_id_null,json=sourceTemplateIdNull" json:"source_template_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTargetTemplateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TargetTemplateId
	}
	return nil
}

func (m *Parameters) GetSourceTemplateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SourceTemplateId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_SetTemplate_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_SetTemplate_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_SetTemplate_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/im_SetTemplate_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x6e, 0xdb, 0x30,
	0x10, 0x84, 0xad, 0xfa, 0x01, 0xf6, 0x50, 0x83, 0x2d, 0x5a, 0x55, 0x46, 0x1f, 0x70, 0x2f, 0xbd,
	0x94, 0x36, 0x6a, 0xa0, 0xe8, 0xad, 0x68, 0x80, 0x1c, 0x74, 0xb0, 0x11, 0x30, 0x41, 0x80, 0xe4,
	0x22, 0x30, 0xe6, 0x46, 0x10, 0x22, 0x89, 0x06, 0x49, 0xc5, 0xbf, 0x11, 0xe4, 0x4f, 0xf2, 0x59,
	0xc9, 0x57, 0x84, 0x12, 0x29, 0x38, 0x92, 0x0e, 0xf1, 0xc5, 0xd6, 0x6a, 0x76, 0x46, 0xc3, 0x9d,
	0x25, 0x5a, 0x70, 0xa5, 0x85, 0x84, 0x39, 0xe4, 0x71, 0x92, 0xc3, 0x7c, 0x2b, 0xc5, 0x06, 0x78,
	0x21, 0x41, 0xcd, 0x93, 0x2c, 0x3a, 0x05, 0x7d, 0x06, 0xd9, 0x36, 0x65, 0x1a, 0xa2, 0xff, 0x9c,
	0x18, 0x50, 0x0b, 0xfc, 0xcd, 0x32, 0x88, 0x65, 0x90, 0x4e, 0x5b, 0xf0, 0xde, 0x49, 0xde, 0xb2,
	0xb4, 0x00, 0x65, 0x59, 0xc1, 0xe7, 0xe6, 0x77, 0x40, 0x4a, 0x21, 0x1d, 0x34, 0x6d, 0x42, 0x19,
	0x28, 0xc5, 0x62, 0x70, 0xe0, 0x8f, 0x36, 0xa8, 0x59, 0x92, 0x5f, 0x0b, 0x99, 0x31, 0x9d, 0x88,
	0xdc, 0x36, 0xcd, 0xee, 0xfb, 0x08, 0x9d, 0x30, 0xc9, 0x0c, 0x0a, 0x52, 0xe1, 0x10, 0x61, 0xcd,
	0x64, 0x0c, 0x3a, 0xd2, 0xb5, 0xad, 0x84, 0xfb, 0xbd, 0xef, 0xbd, 0x9f, 0x6f, 0x7f, 0x4f, 0x89,
	0xb3, 0xef, 0xdc, 0x25, 0xb9, 0x86, 0x18, 0xe4, 0x79, 0x59, 0xd1, 0x89, 0xa5, 0xd5, 0x87, 0x09,
	0x39, 0xfe, 0x83, 0x3e, 0x75, 0xa5, 0xa2, 0xbc, 0x48, 0x53, 0xff, 0x71, 0x64, 0x04, 0xc7, 0xf4,
	0x43, 0x9b, 0xb3, 0x36, 0x60, 0x69, 0x41, 0x89, 0x42, 0x6e, 0xa0, 0x61, 0xa1, 0x7f, 0x80, 0x05,
	0x4b, 0x6b, 0x5a, 0xe8, 0x4a, 0x59, 0x0b, 0x4f, 0xce, 0x42, 0x9b, 0x53, 0x5a, 0x98, 0x3d, 0xf4,
	0xd1, 0x98, 0x82, 0xda, 0x8a, 0x5c, 0x01, 0x5e, 0xa0, 0x41, 0x35, 0x72, 0x37, 0x85, 0x80, 0x34,
	0x43, 0xb4, 0x71, 0x1c, 0x97, 0xbf, 0xd4, 0x36, 0xe2, 0x0b, 0x34, 0x29, 0x87, 0x1d, 0xbd, 0x98,
	0xb6, 0xf1, 0xef, 0x19, 0x32, 0x69, 0x91, 0xdb, 0x99, 0xac, 0x4c, 0x1d, 0xee, 0x6b, 0xfa, 0x2e,
	0x6b, 0xbe, 0xc0, 0x7f, 0xd1, 0xc8, 0x85, 0xec, 0x7b, 0x95, 0xe2, 0xd7, 0x8e, 0xa2, 0x5d, 0x81,
	0x95, 0xfd, 0xa7, 0x75, 0x3b, 0xfe, 0x87, 0x3c, 0x29, 0x76, 0xfe, 0x9b, 0x8a, 0xf5, 0x8b, 0xbc,
	0xb2, 0x89, 0xa4, 0x3e, 0x3e, 0xa1, 0x62, 0x47, 0x4b, 0x66, 0xf0, 0x05, 0x79, 0xe6, 0x19, 0x7f,
	0x44, 0x43, 0x53, 0x95, 0x91, 0xdc, 0xad, 0xcd, 0x40, 0x06, 0x74, 0x60, 0xca, 0x90, 0x1f, 0x51,
	0x34, 0x4d, 0x44, 0x4b, 0x76, 0x7f, 0x25, 0x2e, 0x97, 0xb1, 0x50, 0xfc, 0xa6, 0xc6, 0xf9, 0x41,
	0xb7, 0xe6, 0x6a, 0x58, 0xed, 0xe8, 0xf2, 0x39, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xa6, 0xb5, 0x7a,
	0x6a, 0x03, 0x00, 0x00,
}