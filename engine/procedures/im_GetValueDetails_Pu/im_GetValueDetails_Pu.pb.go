// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetValueDetails_Pu.proto
// DO NOT EDIT!

/*
Package im_GetValueDetails_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetValueDetails_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetValueDetails_Pu

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
	TreeNodeId               *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull           bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	NodeCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull bool                        `protobuf:"varint,1002,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	Value                    *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ValueNull                bool                        `protobuf:"varint,1003,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
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
	RowId   int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Details *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=details" json:"details,omitempty"`
	Value   *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value" json:"value,omitempty"`
	ValueId *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetDetails() *dstore_values.StringValue {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetValueDetails_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetValueDetails_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetValueDetails_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetValueDetails_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0x55, 0x3a, 0x5f, 0x9a, 0x7c, 0x17, 0xc4, 0x8f, 0x41, 0xd5, 0x90, 0x48, 0x15, 0x6a, 0x37,
	0x88, 0xc5, 0x84, 0xf2, 0xa3, 0xb2, 0x81, 0x05, 0x05, 0xa1, 0x2c, 0x3a, 0x2a, 0x5e, 0x54, 0x82,
	0xcd, 0x68, 0x88, 0x2f, 0x83, 0x45, 0x62, 0x57, 0xb6, 0x43, 0x5f, 0x83, 0x9f, 0x17, 0xe2, 0x35,
	0x90, 0x78, 0x01, 0xe0, 0x25, 0xb8, 0xb6, 0x67, 0x48, 0x33, 0x44, 0x4a, 0x36, 0x49, 0x7c, 0xef,
	0x39, 0xf7, 0x9c, 0xf1, 0xb9, 0x19, 0x78, 0x28, 0xac, 0xd3, 0x06, 0x47, 0xa8, 0x2a, 0xa9, 0x70,
	0x74, 0x66, 0xf4, 0x04, 0xc5, 0xdc, 0xa0, 0x1d, 0xc9, 0x59, 0xf1, 0x12, 0xdd, 0x69, 0x39, 0x9d,
	0xe3, 0x73, 0x74, 0xa5, 0x9c, 0xda, 0xe2, 0x64, 0x9e, 0x11, 0xc0, 0x69, 0xb6, 0x1f, 0x59, 0x59,
	0x64, 0x65, 0x2b, 0xa1, 0x83, 0x1b, 0xf5, 0xe8, 0x8f, 0xbe, 0x6e, 0x23, 0x73, 0x70, 0x6b, 0x59,
	0x0f, 0x8d, 0xd1, 0xa6, 0x6e, 0x0d, 0x97, 0x5b, 0x33, 0xb4, 0xb6, 0xac, 0xb0, 0x6e, 0xee, 0xb7,
	0x9b, 0x24, 0xa3, 0xde, 0x69, 0x33, 0x2b, 0x9d, 0xd4, 0x2a, 0x82, 0xf6, 0x7e, 0x6c, 0x01, 0x9c,
	0x94, 0xa6, 0xa4, 0x2e, 0x1a, 0xcb, 0x9e, 0xc0, 0x65, 0x67, 0x10, 0x0b, 0xa5, 0x05, 0x16, 0x52,
	0xa4, 0x9d, 0xdb, 0x9d, 0x3b, 0x97, 0xee, 0x0f, 0xb3, 0xda, 0x7c, 0xed, 0x4b, 0x2a, 0x87, 0x15,
	0x9a, 0xe0, 0x9e, 0x83, 0x27, 0xe4, 0x84, 0x1f, 0x0b, 0x76, 0x17, 0xae, 0x5f, 0xa4, 0x17, 0x6a,
	0x3e, 0x9d, 0xa6, 0x3f, 0x7b, 0x34, 0xa4, 0xcf, 0xaf, 0x2c, 0x70, 0x39, 0x95, 0xd9, 0x2b, 0xd8,
	0x09, 0xb0, 0xc9, 0x7b, 0x92, 0x9f, 0x90, 0xba, 0xb4, 0x4e, 0x4e, 0xbc, 0xe8, 0xd6, 0x7a, 0xd1,
	0x9b, 0x9e, 0x7a, 0xb4, 0xc4, 0x24, 0xf9, 0xa7, 0x30, 0x5c, 0x3d, 0x32, 0x1a, 0xf9, 0x15, 0x8d,
	0xa4, 0xab, 0xb8, 0xc1, 0xd2, 0x3d, 0xe8, 0x06, 0xb1, 0x34, 0x09, 0x0e, 0x06, 0x2d, 0x07, 0xd6,
	0x19, 0xa9, 0xaa, 0x68, 0x20, 0x02, 0xd9, 0x2e, 0x40, 0xf8, 0x11, 0x05, 0x7e, 0x47, 0x81, 0xff,
	0x43, 0xc9, 0x4f, 0xdc, 0xfb, 0x9e, 0x40, 0x9f, 0xa3, 0x3d, 0xd3, 0xca, 0xa2, 0x1f, 0x1f, 0xc2,
	0xab, 0x6f, 0xf5, 0xef, 0xf8, 0x7a, 0x25, 0x62, 0xb0, 0x2f, 0xfc, 0x27, 0x8f, 0x40, 0xf6, 0x1a,
	0xae, 0xf9, 0xd8, 0x8a, 0x0b, 0xb9, 0xd1, 0xed, 0x24, 0x44, 0xce, 0x5a, 0xe4, 0x76, 0xba, 0xc7,
	0x74, 0x1e, 0x2f, 0xce, 0xfc, 0xea, 0x6c, 0xb9, 0xc0, 0x1e, 0x43, 0xaf, 0x5e, 0x17, 0x7a, 0x5a,
	0x3f, 0x71, 0xf7, 0x9f, 0x89, 0x71, 0x99, 0x8e, 0xe3, 0x37, 0x6f, 0xe0, 0xec, 0x08, 0x12, 0xa3,
	0xcf, 0xd3, 0xff, 0x02, 0xeb, 0x20, 0xdb, 0x60, 0xaf, 0xb3, 0xe6, 0x0a, 0x32, 0xae, 0xcf, 0xb9,
	0x67, 0x0f, 0xbe, 0x75, 0x20, 0xa1, 0x03, 0xdb, 0x81, 0x6d, 0x3a, 0xfa, 0xd4, 0x3f, 0xe5, 0x74,
	0x2b, 0x5d, 0xde, 0xa5, 0x23, 0x45, 0xf9, 0x08, 0x7a, 0x22, 0xf2, 0xd3, 0xcf, 0xf9, 0xda, 0x34,
	0x1a, 0x2c, 0x3b, 0x68, 0x12, 0xfc, 0x92, 0x6f, 0x1a, 0xe1, 0x21, 0xf4, 0x63, 0x84, 0xe4, 0xe1,
	0x6b, 0xbe, 0x7e, 0xf5, 0x7a, 0xa1, 0x38, 0x16, 0xcf, 0x4e, 0x61, 0x28, 0x75, 0xeb, 0xf1, 0x17,
	0x2f, 0x83, 0x37, 0x87, 0x95, 0xb6, 0xe2, 0x43, 0xd3, 0x17, 0x1b, 0xbf, 0x2f, 0xde, 0x6e, 0x87,
	0x7f, 0xe6, 0x83, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xf7, 0x14, 0x35, 0x68, 0x04, 0x00,
	0x00,
}
