// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetBinaryCodeIDsForNode_Pu.proto
// DO NOT EDIT!

/*
Package im_GetBinaryCodeIDsForNode_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetBinaryCodeIDsForNode_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetBinaryCodeIDsForNode_Pu

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
	TreeNodeId                 *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull             bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	BinaryPropertyValueId      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=binary_property_value_id,json=binaryPropertyValueId" json:"binary_property_value_id,omitempty"`
	BinaryPropertyValueIdNull  bool                        `protobuf:"varint,1002,opt,name=binary_property_value_id_null,json=binaryPropertyValueIdNull" json:"binary_property_value_id_null,omitempty"`
	IsNodeId                   *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=is_node_id,json=isNodeId" json:"is_node_id,omitempty"`
	IsNodeIdNull               bool                        `protobuf:"varint,1003,opt,name=is_node_id_null,json=isNodeIdNull" json:"is_node_id_null,omitempty"`
	BinaryCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=binary_characteristic_id,json=binaryCharacteristicId" json:"binary_characteristic_id,omitempty"`
	BinaryCharacteristicIdNull bool                        `protobuf:"varint,1004,opt,name=binary_characteristic_id_null,json=binaryCharacteristicIdNull" json:"binary_characteristic_id_null,omitempty"`
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

func (m *Parameters) GetBinaryPropertyValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryPropertyValueId
	}
	return nil
}

func (m *Parameters) GetIsNodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsNodeId
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicId
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
	RowId                 int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	BinaryCodeId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	TreeNodeId            *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	BinaryPropertyValueId *dstore_values.IntegerValue `protobuf:"bytes,20001,opt,name=binary_property_value_id,json=binaryPropertyValueId" json:"binary_property_value_id,omitempty"`
	BinaryPropertyValue   *dstore_values.StringValue  `protobuf:"bytes,20005,opt,name=binary_property_value,json=binaryPropertyValue" json:"binary_property_value,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Response_Row) GetBinaryPropertyValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryPropertyValueId
	}
	return nil
}

func (m *Response_Row) GetBinaryPropertyValue() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryPropertyValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetBinaryCodeIDsForNode_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetBinaryCodeIDsForNode_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetBinaryCodeIDsForNode_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetBinaryCodeIDsForNode_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x6b, 0x13, 0x4f,
	0x14, 0x27, 0xff, 0x34, 0x17, 0xce, 0x3f, 0xb4, 0x3a, 0xa5, 0x25, 0x4d, 0x50, 0xa4, 0x82, 0x88,
	0xc2, 0x46, 0x14, 0xc1, 0x82, 0x17, 0x4c, 0xbd, 0x10, 0xa4, 0x21, 0x2c, 0x5a, 0xd0, 0x07, 0x97,
	0x4d, 0xf6, 0x18, 0x07, 0x93, 0x99, 0x30, 0x33, 0xb1, 0xf4, 0xd5, 0x0f, 0x20, 0x5e, 0x5e, 0xfb,
	0x22, 0xe2, 0x87, 0xf3, 0xf2, 0x21, 0x9c, 0xd9, 0x99, 0x4d, 0xba, 0xdb, 0xb5, 0x31, 0xbe, 0x24,
	0x7b, 0xe6, 0x9c, 0xdf, 0xef, 0xfc, 0xce, 0x65, 0x06, 0x6e, 0x47, 0x52, 0x71, 0x81, 0x2d, 0x64,
	0x43, 0xca, 0xb0, 0x35, 0x11, 0x7c, 0x80, 0xd1, 0x54, 0xa0, 0x6c, 0xd1, 0x71, 0xf0, 0x18, 0x55,
	0x9b, 0xb2, 0x50, 0x1c, 0xee, 0xf2, 0x08, 0x3b, 0x0f, 0xe4, 0x23, 0x2e, 0xba, 0xfa, 0x2b, 0xe8,
	0x4d, 0x3d, 0x1d, 0xa8, 0x38, 0xb9, 0x6a, 0xd1, 0x9e, 0x45, 0x7b, 0xa7, 0x42, 0x1a, 0xeb, 0x2e,
	0xd5, 0xdb, 0x70, 0x34, 0x45, 0x69, 0x19, 0x1a, 0x5b, 0xe9, 0xfc, 0x28, 0x04, 0x17, 0xce, 0xd5,
	0x4c, 0xbb, 0xc6, 0x28, 0x65, 0x38, 0x44, 0xe7, 0xbc, 0x98, 0x75, 0xaa, 0x90, 0xb2, 0x57, 0x5c,
	0x8c, 0x43, 0x45, 0x39, 0xb3, 0x41, 0xdb, 0xef, 0x57, 0x00, 0x7a, 0xa1, 0x08, 0xb5, 0x17, 0x85,
	0x24, 0x77, 0xa0, 0xa6, 0x04, 0x62, 0xc0, 0x8c, 0x20, 0x1a, 0xd5, 0x0b, 0x17, 0x0a, 0x97, 0xff,
	0xbf, 0xde, 0xf4, 0x5c, 0x11, 0x4e, 0x17, 0x65, 0x0a, 0x87, 0x28, 0xf6, 0x8d, 0xe5, 0x83, 0x01,
	0x98, 0x02, 0x3a, 0x11, 0xb9, 0x02, 0x67, 0x8f, 0xc3, 0x03, 0x36, 0x1d, 0x8d, 0xea, 0xdf, 0x2b,
	0x9a, 0xa4, 0xea, 0xaf, 0xce, 0xe3, 0xba, 0xfa, 0x98, 0x3c, 0x85, 0x7a, 0x3f, 0x6e, 0x43, 0xa0,
	0x95, 0x4c, 0x50, 0xa8, 0xc3, 0x20, 0xa6, 0x37, 0x69, 0xff, 0x5b, 0x9c, 0x76, 0xc3, 0x82, 0x7b,
	0x0e, 0x1b, 0x1f, 0x6a, 0x05, 0xf7, 0xe1, 0xdc, 0x9f, 0x58, 0xad, 0x9a, 0x1f, 0x56, 0xcd, 0x56,
	0x2e, 0x3c, 0x16, 0xb6, 0x03, 0x40, 0xe5, 0xac, 0x03, 0xc5, 0x5c, 0x29, 0x7d, 0xce, 0x47, 0x18,
	0x32, 0x2b, 0xa5, 0x4a, 0xa5, 0xab, 0xff, 0x12, 0xac, 0xcd, 0xa1, 0x36, 0xdf, 0x4f, 0x9b, 0xaf,
	0x96, 0xc4, 0xc4, 0x29, 0x9e, 0xcd, 0x6a, 0x1f, 0xbc, 0xd6, 0xcd, 0x1f, 0xe8, 0xde, 0x53, 0xa9,
	0xe8, 0xc0, 0x24, 0x5c, 0x59, 0x5c, 0xfb, 0xa6, 0x05, 0xef, 0xa6, 0xb0, 0x3a, 0x7d, 0x7b, 0x56,
	0xfc, 0x09, 0x5a, 0x2b, 0xe6, 0x97, 0x15, 0xd3, 0xc8, 0xc7, 0x1b, 0x69, 0xdb, 0x5f, 0x4b, 0x50,
	0xf5, 0x51, 0x4e, 0x38, 0x93, 0x48, 0xae, 0x41, 0x29, 0x5e, 0x37, 0xb7, 0x07, 0x0d, 0x2f, 0xbd,
	0xcc, 0x76, 0x15, 0x1f, 0x9a, 0x5f, 0xdf, 0x06, 0x92, 0xe7, 0x70, 0xc6, 0x2c, 0x5a, 0x70, 0x6c,
	0xd3, 0xf4, 0x34, 0x8b, 0x1a, 0xec, 0x65, 0xc0, 0xd9, 0x7d, 0xdc, 0xd3, 0x76, 0x67, 0x6e, 0xfb,
	0x6b, 0xe3, 0xf4, 0x01, 0xb9, 0x05, 0x15, 0xb7, 0xe0, 0x7a, 0x28, 0x86, 0xf1, 0xfc, 0x09, 0x46,
	0xbb, 0xfe, 0x7b, 0xf6, 0xdf, 0x4f, 0xc2, 0xc9, 0x13, 0x28, 0x0a, 0x7e, 0xa0, 0x3b, 0x6b, 0x50,
	0x3b, 0xde, 0x12, 0x37, 0xd2, 0x4b, 0x5a, 0xe1, 0xf9, 0xfc, 0xc0, 0x37, 0x2c, 0x8d, 0x77, 0x45,
	0x28, 0x6a, 0x83, 0x6c, 0x42, 0x59, 0x9b, 0x66, 0x62, 0x1f, 0xba, 0xba, 0x3b, 0x25, 0xbf, 0xa4,
	0xcd, 0x78, 0x08, 0xab, 0xc9, 0x10, 0xdc, 0x0a, 0x7d, 0xec, 0x2e, 0x1e, 0x69, 0xad, 0x3f, 0x17,
	0x10, 0x91, 0xbb, 0x99, 0x6b, 0xf8, 0xa9, 0xbb, 0xdc, 0x3d, 0xbc, 0x09, 0x95, 0x04, 0xfa, 0xf9,
	0x2f, 0xa0, 0x65, 0x66, 0x61, 0xfb, 0xa7, 0x5c, 0xc9, 0x2f, 0x47, 0x85, 0x7f, 0xbe, 0x94, 0x3d,
	0xd8, 0xc8, 0xe5, 0xad, 0x7f, 0x3b, 0xca, 0x2c, 0x96, 0x23, 0x95, 0x4a, 0x50, 0x36, 0xb4, 0x9c,
	0xeb, 0x39, 0x9c, 0xed, 0x97, 0xd0, 0xa4, 0x3c, 0x33, 0xc8, 0xf9, 0xc3, 0xfc, 0xe2, 0xde, 0x90,
	0xcb, 0xe8, 0x4d, 0xe2, 0x8f, 0x96, 0x7e, 0xbb, 0xfb, 0xe5, 0xf8, 0x75, 0xbc, 0xf1, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x51, 0xb9, 0xde, 0xf3, 0xfc, 0x05, 0x00, 0x00,
}
