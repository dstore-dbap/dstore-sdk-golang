// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyApplicationParts_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyApplicationParts_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyApplicationParts_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyApplicationParts_Ad

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
	ApplicationId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ApplicationIdNull   bool                        `protobuf:"varint,1001,opt,name=application_id_null,json=applicationIdNull" json:"application_id_null,omitempty"`
	UserId              *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull          bool                        `protobuf:"varint,1002,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	ApplicationPart     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=application_part,json=applicationPart" json:"application_part,omitempty"`
	ApplicationPartNull bool                        `protobuf:"varint,1003,opt,name=application_part_null,json=applicationPartNull" json:"application_part_null,omitempty"`
	Delete              *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull          bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetApplicationId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Parameters) GetApplicationPart() *dstore_values.StringValue {
	if m != nil {
		return m.ApplicationPart
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
	}
	return nil
}

type Response struct {
	Error             *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation   []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message           []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row               []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ApplicationPartId *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=application_part_id,json=applicationPartId" json:"application_part_id,omitempty"`
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

func (m *Response) GetApplicationPartId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationPartId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyApplicationParts_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyApplicationParts_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyApplicationParts_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0xae, 0x49, 0xca, 0xa9, 0xda, 0x76, 0x82, 0x12, 0x13, 0x94, 0x5a, 0x6f, 0xc4, 0x8b,
	0x8d, 0x18, 0x85, 0x82, 0x57, 0x2d, 0xf4, 0x22, 0x4a, 0x8a, 0x2c, 0x22, 0xe8, 0xcd, 0x32, 0xed,
	0x9c, 0x86, 0x81, 0xcd, 0xcc, 0x32, 0xb3, 0xb1, 0xf8, 0x08, 0xde, 0xf9, 0x22, 0x3e, 0x98, 0x3f,
	0x0f, 0xe1, 0x99, 0x3d, 0x53, 0x93, 0x5d, 0xc5, 0x8a, 0x37, 0xc9, 0xce, 0x9e, 0xf3, 0x7d, 0xdf,
	0xf9, 0xf9, 0x76, 0xe0, 0x85, 0xf2, 0x95, 0x75, 0x38, 0x46, 0x33, 0xd7, 0x06, 0xc7, 0xa5, 0xb3,
	0x67, 0xa8, 0x96, 0x0e, 0xfd, 0x78, 0xa1, 0xf3, 0x99, 0x55, 0xfa, 0xfc, 0xe3, 0x61, 0x59, 0x16,
	0xfa, 0x4c, 0x56, 0xda, 0x9a, 0xd7, 0xd2, 0x55, 0x3e, 0x3f, 0x54, 0x29, 0xe5, 0x55, 0x56, 0x3c,
	0x66, 0x70, 0xca, 0xe0, 0xf4, 0x6f, 0x88, 0x61, 0x3f, 0x0a, 0x7d, 0x90, 0xc5, 0x12, 0x3d, 0x13,
	0x0c, 0xef, 0x36, 0xd5, 0xd1, 0x39, 0xeb, 0x62, 0x68, 0xd4, 0x0c, 0x2d, 0xd0, 0x7b, 0x39, 0xc7,
	0x18, 0x7c, 0xd8, 0x0e, 0x56, 0x52, 0x9b, 0x73, 0xeb, 0x16, 0xb5, 0x26, 0x27, 0xed, 0x7f, 0x49,
	0x00, 0x48, 0x5e, 0x52, 0x14, 0x9d, 0x17, 0x47, 0x70, 0x4b, 0xae, 0xea, 0xca, 0xb5, 0x1a, 0x6c,
	0xec, 0x6d, 0x3c, 0xda, 0x7a, 0x3a, 0x4a, 0x63, 0x17, 0xb1, 0x32, 0x6d, 0x2a, 0x9c, 0xa3, 0x7b,
	0x1b, 0x4e, 0xd9, 0xcd, 0x35, 0xc8, 0x54, 0x89, 0x31, 0xf4, 0x9b, 0x1c, 0xb9, 0x59, 0x16, 0xc5,
	0xe0, 0x6b, 0x8f, 0x98, 0x36, 0xb3, 0xdd, 0x46, 0xf2, 0x09, 0x45, 0xc4, 0x33, 0xe8, 0x2d, 0x3d,
	0xba, 0xa0, 0x76, 0xed, 0x6a, 0xb5, 0x6e, 0xc8, 0x25, 0x99, 0x07, 0x70, 0x23, 0xa2, 0x98, 0xff,
	0x1b, 0xf3, 0x03, 0x87, 0x6b, 0xe2, 0x63, 0xd8, 0x59, 0xaf, 0xa4, 0xa4, 0x31, 0x0f, 0x92, 0x5a,
	0x61, 0xd8, 0x52, 0xf0, 0x95, 0xd3, 0x66, 0xce, 0x02, 0xdb, 0xb2, 0xb9, 0x19, 0x31, 0x81, 0xdb,
	0x6d, 0x1a, 0x96, 0xfc, 0xce, 0x92, 0xfd, 0x16, 0xa0, 0xd6, 0x9e, 0x40, 0x57, 0x61, 0x41, 0x53,
	0x1d, 0x5c, 0xff, 0x63, 0x4f, 0xa7, 0xd6, 0x16, 0x28, 0x4d, 0xec, 0x89, 0x53, 0xc5, 0x1e, 0x6c,
	0xf1, 0x13, 0xf3, 0xff, 0x88, 0x2d, 0xf1, 0xbb, 0x40, 0xbb, 0xff, 0x29, 0x81, 0xcd, 0x0c, 0x7d,
	0x69, 0x8d, 0x47, 0xf1, 0x04, 0x3a, 0xb5, 0x1b, 0xe2, 0x92, 0x7e, 0x35, 0x15, 0xad, 0xc6, 0x4e,
	0x39, 0x0e, 0xbf, 0x19, 0x27, 0x8a, 0x77, 0xb0, 0x13, 0x7c, 0x90, 0xaf, 0x19, 0x81, 0x66, 0x9e,
	0x10, 0x38, 0x6d, 0x81, 0xdb, 0x76, 0x99, 0xd1, 0x79, 0xba, 0x3a, 0x67, 0xdb, 0x8b, 0xe6, 0x0b,
	0x71, 0x00, 0xbd, 0xe8, 0x3f, 0x9a, 0x71, 0x60, 0xbc, 0xff, 0x1b, 0x23, 0xbb, 0x73, 0xc6, 0xff,
	0xd9, 0x65, 0xba, 0x78, 0x09, 0x89, 0xb3, 0x17, 0x34, 0xa7, 0x80, 0x3a, 0x48, 0xff, 0xfd, 0x7b,
	0x49, 0x2f, 0x27, 0x91, 0x66, 0xf6, 0x22, 0x0b, 0x24, 0xe2, 0x55, 0xd3, 0x7c, 0xf5, 0xae, 0xc8,
	0x57, 0x78, 0xb5, 0xaf, 0x76, 0x5b, 0x5b, 0x9c, 0xaa, 0xe1, 0x3d, 0x48, 0x88, 0x58, 0xdc, 0x81,
	0x2e, 0x51, 0x07, 0x9a, 0xcf, 0x27, 0xc4, 0xd3, 0xc9, 0x3a, 0x74, 0x9c, 0xaa, 0xa3, 0x37, 0x30,
	0xd2, 0xb6, 0x55, 0xee, 0xea, 0x6e, 0x78, 0xff, 0xfc, 0xbf, 0x6e, 0x8d, 0xd3, 0x6e, 0xfd, 0x61,
	0x4e, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x60, 0x32, 0x59, 0xc8, 0x75, 0x04, 0x00, 0x00,
}