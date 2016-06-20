// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetPageTreeNodes_Ad.proto
// DO NOT EDIT!

/*
Package im_GetPageTreeNodes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetPageTreeNodes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetPageTreeNodes_Ad

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
	IncludeInactive     *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=include_inactive,json=includeInactive" json:"include_inactive,omitempty"`
	IncludeInactiveNull bool                        `protobuf:"varint,1001,opt,name=include_inactive_null,json=includeInactiveNull" json:"include_inactive_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetIncludeInactive() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeInactive
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
	RowId           int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	NodeDescription *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Active          *dstore_values.BooleanValue `protobuf:"bytes,10002,opt,name=active" json:"active,omitempty"`
	TreeNodeId      *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Deleted         *dstore_values.BooleanValue `protobuf:"bytes,10005,opt,name=deleted" json:"deleted,omitempty"`
	LevelId         *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	SymbolId        *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=symbol_id,json=symbolId" json:"symbol_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Response_Row) GetActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.Active
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

func (m *Response_Row) GetDeleted() *dstore_values.BooleanValue {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *Response_Row) GetLevelId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelId
	}
	return nil
}

func (m *Response_Row) GetSymbolId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SymbolId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetPageTreeNodes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetPageTreeNodes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetPageTreeNodes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetPageTreeNodes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0xa9, 0xb5, 0x4d, 0x3d, 0x0a, 0xbb, 0xcc, 0xa2, 0xc4, 0x16, 0x64, 0x59, 0xbd, 0xf0,
	0x2a, 0x95, 0x5d, 0x57, 0xd7, 0x1b, 0x41, 0x59, 0x95, 0x5e, 0x6c, 0x58, 0x06, 0x51, 0xf4, 0x26,
	0xa4, 0x9d, 0x63, 0x18, 0x4c, 0x66, 0xca, 0xcc, 0xb4, 0x8b, 0x8f, 0xe0, 0x9d, 0xdf, 0xbe, 0x8d,
	0xcf, 0xe1, 0x2b, 0xf8, 0x16, 0x4e, 0x32, 0x13, 0x6a, 0xe2, 0xc2, 0xe6, 0xa6, 0xe5, 0xf4, 0xfc,
	0x7f, 0xe7, 0xe3, 0x7f, 0xd2, 0xc0, 0x21, 0xd3, 0x46, 0x2a, 0x9c, 0xa2, 0xc8, 0xb8, 0xc0, 0xe9,
	0x52, 0xc9, 0x05, 0xb2, 0x95, 0x42, 0x3d, 0xe5, 0x45, 0xf2, 0x02, 0xcd, 0x69, 0x9a, 0xe1, 0x4b,
	0x85, 0x18, 0x4b, 0x86, 0x3a, 0x79, 0xc2, 0x22, 0xab, 0x30, 0x92, 0xdc, 0x71, 0x58, 0xe4, 0xb0,
	0xe8, 0x7c, 0xed, 0x78, 0xc7, 0x17, 0x5f, 0xa7, 0xf9, 0x0a, 0xb5, 0x43, 0xc7, 0x37, 0x9b, 0x1d,
	0x51, 0x29, 0xa9, 0x7c, 0x6a, 0xd2, 0x4c, 0x15, 0xa8, 0xb5, 0x2d, 0xe9, 0x93, 0xb7, 0xdb, 0x49,
	0x93, 0x72, 0xf1, 0x4e, 0xaa, 0x22, 0x35, 0x5c, 0x0a, 0x27, 0xda, 0xfb, 0xd8, 0x03, 0x38, 0x4d,
	0x55, 0x6a, 0xb3, 0xa8, 0x34, 0x79, 0x0e, 0xdb, 0x5c, 0x2c, 0xf2, 0x15, 0xc3, 0x84, 0x8b, 0x74,
	0x61, 0xf8, 0x1a, 0xc3, 0xde, 0x6e, 0xef, 0xee, 0xd5, 0xfd, 0x49, 0xe4, 0x37, 0xf0, 0xb3, 0xcd,
	0xa5, 0xcc, 0x31, 0x15, 0xaf, 0xca, 0x88, 0x6e, 0x79, 0x68, 0xe6, 0x19, 0x72, 0x00, 0xd7, 0xdb,
	0x75, 0x12, 0xb1, 0xca, 0xf3, 0xf0, 0x4f, 0x60, 0xab, 0x8d, 0xe8, 0x4e, 0x0b, 0x88, 0x6d, 0x6e,
	0xef, 0xf7, 0x00, 0x46, 0x14, 0xf5, 0x52, 0x0a, 0x8d, 0xe4, 0x1e, 0x0c, 0xaa, 0x4d, 0x7d, 0xfb,
	0x71, 0xd4, 0x34, 0xd0, 0xb9, 0xf0, 0xac, 0xfc, 0xa4, 0x4e, 0x48, 0xde, 0xc0, 0x76, 0xb9, 0x63,
	0xf2, 0xcf, 0x92, 0xe1, 0xa5, 0xdd, 0xbe, 0x85, 0xa3, 0x16, 0xdc, 0xb6, 0xe2, 0xc4, 0xc6, 0xb3,
	0x4d, 0x4c, 0xb7, 0x8a, 0xe6, 0x0f, 0xe4, 0x08, 0x02, 0xef, 0x6d, 0xd8, 0xaf, 0x2a, 0xde, 0xfa,
	0xaf, 0xa2, 0x73, 0xfe, 0xc4, 0x7d, 0xd3, 0x5a, 0x4e, 0x8e, 0xa1, 0xaf, 0xe4, 0x59, 0x78, 0xb9,
	0xa2, 0xf6, 0xa3, 0x2e, 0x4f, 0x41, 0x54, 0x7b, 0x10, 0x51, 0x79, 0x46, 0x4b, 0x7c, 0xfc, 0xab,
	0x0f, 0x7d, 0x1b, 0x90, 0x1b, 0x30, 0xb4, 0x61, 0xc2, 0x59, 0xf8, 0x29, 0xb6, 0xb6, 0x0c, 0xe8,
	0xc0, 0x86, 0x33, 0x56, 0x9e, 0x4d, 0x58, 0x3a, 0xb1, 0x05, 0x16, 0x8a, 0x2f, 0xab, 0xd5, 0x3f,
	0xc7, 0x4d, 0xe3, 0xfc, 0xdd, 0xb4, 0x51, 0x5c, 0x64, 0xfe, 0x6c, 0x25, 0x74, 0xbc, 0x61, 0xc8,
	0x7d, 0x18, 0xfa, 0xa3, 0x7f, 0x89, 0x2f, 0xbe, 0xba, 0xd7, 0x92, 0xc7, 0x70, 0xcd, 0xd8, 0xf9,
	0x93, 0x6a, 0x04, 0x3b, 0xdb, 0xd7, 0xf3, 0x59, 0x2e, 0x0c, 0x66, 0xa8, 0x1c, 0x0b, 0xc6, 0x6f,
	0x6c, 0xa7, 0x3f, 0x84, 0xa0, 0x46, 0xbf, 0x75, 0x40, 0x87, 0xc2, 0x61, 0x0f, 0x20, 0x60, 0x98,
	0xdb, 0xe7, 0x96, 0x85, 0xdf, 0x3b, 0x4c, 0x5b, 0x8b, 0xc9, 0x43, 0x18, 0xe5, 0xb8, 0xc6, 0xbc,
	0xec, 0xf7, 0xa3, 0x43, 0xbf, 0xa0, 0x52, 0xdb, 0x86, 0x8f, 0xe0, 0x8a, 0xfe, 0x50, 0xcc, 0x65,
	0x45, 0xfe, 0xec, 0x40, 0x8e, 0x9c, 0x7c, 0xc6, 0x9e, 0xbe, 0x86, 0x09, 0x97, 0xad, 0xeb, 0x6f,
	0x5e, 0x1d, 0x6f, 0x8f, 0x32, 0xa9, 0xd9, 0xfb, 0x3a, 0xcf, 0xba, 0xbf, 0x5d, 0xe6, 0xc3, 0xea,
	0x6f, 0x7c, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x4e, 0x9f, 0x8f, 0x97, 0x04, 0x00, 0x00,
}