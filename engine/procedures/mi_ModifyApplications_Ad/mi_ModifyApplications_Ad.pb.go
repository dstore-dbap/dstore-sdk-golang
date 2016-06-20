// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyApplications_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyApplications_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyApplications_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyApplications_Ad

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
	ApplicationName     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=application_name,json=applicationName" json:"application_name,omitempty"`
	ApplicationNameNull bool                        `protobuf:"varint,1001,opt,name=application_name_null,json=applicationNameNull" json:"application_name_null,omitempty"`
	Delete              *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=delete" json:"delete,omitempty"`
	DeleteNull          bool                        `protobuf:"varint,1002,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetApplicationName() *dstore_values.StringValue {
	if m != nil {
		return m.ApplicationName
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
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ApplicationId   *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
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

func (m *Response) GetApplicationId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyApplications_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyApplications_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyApplications_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifyApplications_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0x26, 0x8d, 0x49, 0xcb, 0x14, 0x6d, 0xd9, 0xa2, 0x9c, 0x09, 0x4a, 0xa9, 0x20, 0x3e, 0x6d,
	0xc4, 0x0a, 0x8a, 0x6f, 0x2d, 0x54, 0xc8, 0x43, 0x82, 0xec, 0x83, 0x50, 0x5f, 0x8e, 0x6d, 0x77,
	0x7a, 0x2c, 0xde, 0xed, 0x1e, 0xbb, 0x17, 0x8b, 0xff, 0xc2, 0x7f, 0xe8, 0xb3, 0xfe, 0x06, 0x1f,
	0x9c, 0xbb, 0xdd, 0x90, 0xcb, 0xaa, 0x60, 0x5f, 0x92, 0x9b, 0xf9, 0xe6, 0xfb, 0x66, 0xf6, 0x9b,
	0x5d, 0x78, 0xa3, 0x7c, 0x63, 0x1d, 0xce, 0xd0, 0x14, 0xda, 0xe0, 0xac, 0x76, 0xf6, 0x1a, 0xd5,
	0xca, 0xa1, 0x9f, 0x55, 0x3a, 0x5f, 0x58, 0xa5, 0x6f, 0xbe, 0x9e, 0xd5, 0x75, 0xa9, 0xaf, 0x65,
	0xa3, 0xad, 0xf1, 0xf9, 0x99, 0xe2, 0x54, 0xd3, 0x58, 0xf6, 0x3c, 0x10, 0x79, 0x20, 0xf2, 0x7f,
	0x55, 0x4f, 0x8e, 0x62, 0x83, 0x2f, 0xb2, 0x5c, 0xa1, 0x0f, 0xe4, 0xc9, 0xe3, 0xed, 0xae, 0xe8,
	0x9c, 0x75, 0x11, 0x9a, 0x6e, 0x43, 0x15, 0x7a, 0x2f, 0x0b, 0x8c, 0xe0, 0xb3, 0x14, 0x6c, 0xa4,
	0x36, 0x37, 0xd6, 0x55, 0x5d, 0xbf, 0x50, 0x74, 0xf2, 0x7d, 0x00, 0xf0, 0x41, 0x3a, 0x49, 0x28,
	0x3a, 0xcf, 0x2e, 0xe0, 0x50, 0x6e, 0x66, 0xca, 0x0d, 0xe5, 0xb3, 0xc1, 0xf1, 0xe0, 0xc5, 0xfe,
	0xab, 0x09, 0x8f, 0x67, 0x88, 0xb3, 0xf9, 0xc6, 0x69, 0x53, 0x7c, 0x6c, 0x03, 0x71, 0xd0, 0xe3,
	0x2c, 0x89, 0xc2, 0x4e, 0xe1, 0x61, 0x2a, 0x93, 0x9b, 0x55, 0x59, 0x66, 0x3f, 0x76, 0x49, 0x6c,
	0x4f, 0x1c, 0x25, 0x84, 0x25, 0x61, 0x44, 0x1a, 0x2b, 0x2c, 0x69, 0x8e, 0x6c, 0xa7, 0xeb, 0x38,
	0x4d, 0x3a, 0x5e, 0x59, 0x5b, 0xa2, 0x34, 0xa1, 0x65, 0x2c, 0x65, 0xc7, 0xb0, 0x1f, 0xbe, 0x82,
	0xfe, 0xcf, 0xa0, 0x0f, 0x21, 0xd7, 0xca, 0x9e, 0xfc, 0xda, 0x81, 0x3d, 0x81, 0xbe, 0x26, 0x8b,
	0x91, 0xbd, 0x84, 0x51, 0xe7, 0x5f, 0x7a, 0xa8, 0xb8, 0x98, 0xe0, 0xed, 0x45, 0xfb, 0x2b, 0x42,
	0x21, 0xbb, 0x84, 0xc3, 0xd6, 0xb9, 0xbc, 0x67, 0x1d, 0xcd, 0x37, 0x24, 0x32, 0x4f, 0xc8, 0xa9,
	0xc1, 0x0b, 0x8a, 0xe7, 0x9b, 0x58, 0x1c, 0x54, 0xdb, 0x09, 0xf6, 0x16, 0x76, 0xe3, 0xc6, 0xb2,
	0x61, 0xa7, 0xf8, 0xf4, 0x0f, 0xc5, 0xb0, 0xcf, 0x45, 0xf8, 0x17, 0xeb, 0x72, 0xf6, 0x1e, 0x86,
	0xce, 0xde, 0x66, 0xf7, 0x3a, 0xd6, 0x6b, 0xfe, 0x7f, 0xb7, 0x8b, 0xaf, 0x5d, 0xe0, 0xc2, 0xde,
	0x8a, 0x56, 0x80, 0x9d, 0xc3, 0x83, 0xfe, 0x9e, 0xb4, 0xca, 0xf0, 0xaf, 0xd6, 0x6b, 0xd3, 0x60,
	0x81, 0x2e, 0x58, 0x7f, 0xbf, 0x47, 0x99, 0xab, 0xc9, 0x13, 0x18, 0x92, 0x1e, 0x7b, 0x04, 0x63,
	0x52, 0x6c, 0x25, 0xbe, 0x2d, 0x49, 0x63, 0x24, 0x46, 0x14, 0xce, 0xd5, 0xf9, 0x25, 0x4c, 0xb5,
	0x4d, 0x26, 0xdc, 0x3c, 0x9c, 0x4f, 0xef, 0x0a, 0xeb, 0xd5, 0xe7, 0x35, 0xae, 0xee, 0xf2, 0xb6,
	0xae, 0xc6, 0xdd, 0x15, 0x3e, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x04, 0xf3, 0xf7, 0x97,
	0x03, 0x00, 0x00,
}