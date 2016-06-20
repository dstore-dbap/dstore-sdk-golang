// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_TraverseTreeView_Pu.proto
// DO NOT EDIT!

/*
Package im_TraverseTreeView_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_TraverseTreeView_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_TraverseTreeView_Pu

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
	TreeNodeId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull       bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	Next                 *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=next" json:"next,omitempty"`
	NextNull             bool                        `protobuf:"varint,1002,opt,name=next_null,json=nextNull" json:"next_null,omitempty"`
	LevelNo              *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=level_no,json=levelNo" json:"level_no,omitempty"`
	LevelNoNull          bool                        `protobuf:"varint,1003,opt,name=level_no_null,json=levelNoNull" json:"level_no_null,omitempty"`
	DomainTreeNodeId     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=domain_tree_node_id,json=domainTreeNodeId" json:"domain_tree_node_id,omitempty"`
	DomainTreeNodeIdNull bool                        `protobuf:"varint,1004,opt,name=domain_tree_node_id_null,json=domainTreeNodeIdNull" json:"domain_tree_node_id_null,omitempty"`
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

func (m *Parameters) GetNext() *dstore_values.BooleanValue {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *Parameters) GetLevelNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelNo
	}
	return nil
}

func (m *Parameters) GetDomainTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.DomainTreeNodeId
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
	RowId            int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ResultTreeNodeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=result_tree_node_id,json=resultTreeNodeId" json:"result_tree_node_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetResultTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ResultTreeNodeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_TraverseTreeView_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_TraverseTreeView_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_TraverseTreeView_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x71, 0xd3, 0x84, 0x29, 0x3f, 0x65, 0x83, 0x90, 0x49, 0x10, 0x42, 0x2d, 0x07, 0xc4,
	0x61, 0x8d, 0x8a, 0xf8, 0xb9, 0x70, 0x41, 0x70, 0x28, 0xa8, 0x56, 0xb5, 0x8a, 0x2a, 0xc1, 0xc5,
	0x72, 0xf1, 0x10, 0x19, 0xd9, 0xbb, 0xd5, 0xda, 0x4e, 0xb8, 0xf3, 0x02, 0xf0, 0x16, 0xbc, 0x1a,
	0x3f, 0x0f, 0xc1, 0xac, 0xc7, 0x91, 0x6b, 0x13, 0xa9, 0xed, 0x25, 0xce, 0xec, 0x7c, 0xdf, 0x37,
	0xa3, 0x6f, 0x66, 0x17, 0x9e, 0x25, 0x45, 0x69, 0x2c, 0x06, 0xa8, 0xe7, 0xa9, 0xc6, 0xe0, 0xd4,
	0x9a, 0x4f, 0x98, 0x54, 0x16, 0x8b, 0x20, 0xcd, 0xa3, 0x99, 0x8d, 0x17, 0x68, 0x0b, 0x9c, 0x59,
	0xc4, 0xe3, 0x14, 0x97, 0xd1, 0x51, 0x25, 0x09, 0x51, 0x1a, 0xf1, 0x90, 0x69, 0x92, 0x69, 0x72,
	0x3d, 0x76, 0x32, 0x6e, 0xc4, 0x17, 0x71, 0x56, 0x61, 0xc1, 0xd4, 0xc9, 0xdd, 0x6e, 0x45, 0xb4,
	0xd6, 0xd8, 0x26, 0x35, 0xed, 0xa6, 0x72, 0x2c, 0x8a, 0x78, 0x8e, 0x4d, 0x72, 0xaf, 0x9f, 0x2c,
	0xe3, 0x54, 0x7f, 0x36, 0x36, 0x8f, 0xcb, 0xd4, 0x68, 0x06, 0xed, 0xfe, 0xf4, 0x00, 0x8e, 0x62,
	0x1b, 0x53, 0x96, 0x9a, 0x11, 0xaf, 0xe0, 0x5a, 0x49, 0xfd, 0x44, 0xda, 0x24, 0x18, 0xa5, 0x89,
	0xbf, 0xf1, 0x60, 0xe3, 0xd1, 0xf6, 0xfe, 0x54, 0x36, 0xdd, 0x37, 0x7d, 0xa5, 0xba, 0xc4, 0x39,
	0xda, 0x63, 0x17, 0x29, 0x70, 0x84, 0x90, 0xf0, 0x07, 0x89, 0x78, 0x0c, 0xb7, 0xce, 0xd2, 0x23,
	0x5d, 0x65, 0x99, 0xff, 0x6b, 0x48, 0x22, 0x23, 0x75, 0xa3, 0xc5, 0x85, 0x74, 0x2c, 0x02, 0xd8,
	0xd4, 0xf8, 0xb5, 0xf4, 0xaf, 0xac, 0x2d, 0x71, 0x62, 0x4c, 0x86, 0xb1, 0xe6, 0x12, 0x35, 0x50,
	0xdc, 0x83, 0xab, 0xee, 0xcb, 0xa2, 0xbf, 0x59, 0x74, 0xe4, 0x4e, 0x6a, 0xb9, 0xe7, 0x30, 0xca,
	0x70, 0x81, 0x19, 0xd5, 0xf6, 0xbd, 0xf3, 0xbb, 0x1e, 0xd6, 0xe0, 0xd0, 0x88, 0x3d, 0xb8, 0xbe,
	0xe2, 0xb1, 0xf2, 0x1f, 0x56, 0xde, 0x6e, 0x00, 0xb5, 0xf8, 0x3b, 0x18, 0x27, 0x26, 0x27, 0xfb,
	0xa2, 0x8e, 0x3b, 0x9b, 0xe7, 0xd7, 0xd9, 0x61, 0xde, 0xac, 0xf5, 0xe8, 0x05, 0xf8, 0x6b, 0xb4,
	0xb8, 0xf6, 0x5f, 0xae, 0x7d, 0xbb, 0x4f, 0x72, 0x4d, 0xec, 0x7e, 0xf3, 0x60, 0xa4, 0xb0, 0x38,
	0x35, 0xba, 0x40, 0xf1, 0x04, 0x06, 0xf5, 0x22, 0x34, 0x13, 0x9a, 0xc8, 0xee, 0x7e, 0xf1, 0x92,
	0xbc, 0x75, 0xbf, 0x8a, 0x81, 0xe2, 0x03, 0xec, 0xb8, 0x15, 0x88, 0xce, 0xec, 0x00, 0x79, 0xef,
	0x11, 0x59, 0xf6, 0xc8, 0xfd, 0x4d, 0x39, 0xa4, 0xf8, 0xa0, 0x8d, 0xd5, 0xcd, 0xbc, 0x7b, 0x20,
	0x5e, 0xc2, 0xb0, 0x59, 0x3d, 0xb2, 0xde, 0x29, 0xde, 0xff, 0x4f, 0x91, 0x17, 0xf3, 0x90, 0xbf,
	0x6a, 0x05, 0x17, 0x6f, 0xc0, 0xb3, 0x66, 0x49, 0x46, 0x3a, 0xd6, 0xbe, 0xbc, 0xc8, 0x25, 0x91,
	0x2b, 0x0f, 0xa4, 0x32, 0x4b, 0xe5, 0xe8, 0x93, 0x2f, 0xe0, 0xd1, 0x7f, 0x71, 0x07, 0xb6, 0x28,
	0x72, 0x83, 0xf9, 0x1e, 0x92, 0x2b, 0x03, 0x35, 0xa0, 0x90, 0x1c, 0x7f, 0x0f, 0x63, 0xba, 0x9f,
	0x55, 0x56, 0x76, 0xa7, 0xf7, 0x23, 0xbc, 0xc0, 0xf8, 0x98, 0xd8, 0x4e, 0xe2, 0x75, 0x08, 0xd3,
	0xd4, 0xf4, 0x1a, 0x6d, 0x1f, 0x81, 0x8f, 0xc1, 0x25, 0x9f, 0x87, 0x93, 0xad, 0xfa, 0x1e, 0x3e,
	0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x79, 0x05, 0x17, 0xbe, 0x58, 0x04, 0x00, 0x00,
}