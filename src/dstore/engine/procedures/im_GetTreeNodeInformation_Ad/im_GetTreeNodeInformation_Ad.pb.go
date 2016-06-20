// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetTreeNodeInformation_Ad.proto
// DO NOT EDIT!

/*
Package im_GetTreeNodeInformation_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetTreeNodeInformation_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetTreeNodeInformation_Ad

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
	TreeNodeId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull               bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	LanguageId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull               bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	GetAdditionalInformation     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=get_additional_information,json=getAdditionalInformation" json:"get_additional_information,omitempty"`
	GetAdditionalInformationNull bool                        `protobuf:"varint,1003,opt,name=get_additional_information_null,json=getAdditionalInformationNull" json:"get_additional_information_null,omitempty"`
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

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetGetAdditionalInformation() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetAdditionalInformation
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
	RowId                  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Predecessor            *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=predecessor" json:"predecessor,omitempty"`
	LevelNo                *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=level_no,json=levelNo" json:"level_no,omitempty"`
	SortNo                 *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	CountActiveSuccessors  *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=count_active_successors,json=countActiveSuccessors" json:"count_active_successors,omitempty"`
	CountDeletedSuccessors *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=count_deleted_successors,json=countDeletedSuccessors" json:"count_deleted_successors,omitempty"`
	NodeDescription        *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Active                 *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=active" json:"active,omitempty"`
	NodeId                 *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TreeNodeId             *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	CountSuccessors        *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=count_successors,json=countSuccessors" json:"count_successors,omitempty"`
	Deleted                *dstore_values.BooleanValue `protobuf:"bytes,10011,opt,name=deleted" json:"deleted,omitempty"`
	InheritsFrom           *dstore_values.IntegerValue `protobuf:"bytes,10012,opt,name=inherits_from,json=inheritsFrom" json:"inherits_from,omitempty"`
	LevelId                *dstore_values.IntegerValue `protobuf:"bytes,10013,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	SymbolId               *dstore_values.IntegerValue `protobuf:"bytes,10014,opt,name=symbol_id,json=symbolId" json:"symbol_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPredecessor() *dstore_values.IntegerValue {
	if m != nil {
		return m.Predecessor
	}
	return nil
}

func (m *Response_Row) GetLevelNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelNo
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetCountActiveSuccessors() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountActiveSuccessors
	}
	return nil
}

func (m *Response_Row) GetCountDeletedSuccessors() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountDeletedSuccessors
	}
	return nil
}

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

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetCountSuccessors() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountSuccessors
	}
	return nil
}

func (m *Response_Row) GetDeleted() *dstore_values.BooleanValue {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *Response_Row) GetInheritsFrom() *dstore_values.IntegerValue {
	if m != nil {
		return m.InheritsFrom
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetTreeNodeInformation_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetTreeNodeInformation_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetTreeNodeInformation_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x95, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xc7, 0x03, 0x95, 0xb6, 0x1c, 0x50, 0xea, 0x18, 0x71, 0x2d, 0x46, 0x09, 0xde, 0x28, 0x17,
	0x8b, 0x51, 0x51, 0x8c, 0x4a, 0x52, 0x83, 0x10, 0x4c, 0xd8, 0x98, 0x95, 0x90, 0xe8, 0xcd, 0x66,
	0xe9, 0x1e, 0xea, 0x26, 0xdb, 0x99, 0x66, 0x66, 0x0a, 0xf1, 0xda, 0x17, 0xf0, 0xfb, 0x5b, 0xaf,
	0x7d, 0x14, 0x5f, 0x45, 0x7d, 0x09, 0xcf, 0xec, 0x6c, 0xe9, 0xb6, 0x8a, 0xac, 0xde, 0x40, 0xa6,
	0x73, 0x7e, 0xff, 0xf3, 0x31, 0xff, 0x99, 0x85, 0x9b, 0x91, 0xd2, 0x42, 0xe2, 0x02, 0xf2, 0x56,
	0xcc, 0x71, 0xa1, 0x23, 0x45, 0x13, 0xa3, 0xae, 0x44, 0xb5, 0x10, 0xb7, 0x83, 0x35, 0xd4, 0x9b,
	0x12, 0xd1, 0x13, 0x11, 0xae, 0xf3, 0x1d, 0x21, 0xdb, 0xa1, 0x8e, 0x05, 0x0f, 0x1a, 0x91, 0x4b,
	0x71, 0x5a, 0xb0, 0x79, 0x0b, 0xbb, 0x16, 0x76, 0xff, 0x46, 0xd4, 0x4f, 0x64, 0x89, 0x76, 0xc3,
	0xa4, 0x8b, 0xca, 0x0a, 0xd4, 0x4f, 0x0f, 0x66, 0x47, 0x29, 0x85, 0xcc, 0xb6, 0x66, 0x06, 0xb7,
	0xda, 0xa8, 0x54, 0xd8, 0xc2, 0x6c, 0xf3, 0xfc, 0xf0, 0xa6, 0x0e, 0xe3, 0x7e, 0x3a, 0x1b, 0x34,
	0xf7, 0xb4, 0x04, 0x70, 0x3f, 0x94, 0x21, 0xed, 0xa2, 0x54, 0xec, 0x36, 0x4c, 0x6a, 0x2a, 0x2d,
	0xe0, 0x54, 0x5b, 0x10, 0x47, 0xce, 0xc8, 0xec, 0xc8, 0x85, 0x89, 0xcb, 0x33, 0x6e, 0xd6, 0x43,
	0x56, 0x57, 0xcc, 0x35, 0xb6, 0x50, 0x6e, 0x99, 0x95, 0x0f, 0xba, 0xd7, 0x4b, 0xc4, 0xe6, 0xe1,
	0x78, 0x1e, 0x0f, 0x78, 0x37, 0x49, 0x9c, 0xef, 0x15, 0x12, 0xa9, 0xfa, 0xc7, 0xfa, 0x71, 0x1e,
	0xfd, 0xcc, 0x6e, 0xc1, 0x44, 0x12, 0xf2, 0x56, 0x97, 0x0a, 0x36, 0x99, 0x46, 0x0b, 0x64, 0xea,
	0xc5, 0x53, 0xa6, 0x8b, 0x50, 0xcb, 0xd1, 0x36, 0xd1, 0x8f, 0x2c, 0x51, 0x3f, 0x2c, 0x4d, 0xf4,
	0x10, 0xea, 0x2d, 0xd4, 0x41, 0x18, 0x45, 0xb1, 0x69, 0x3c, 0x4c, 0x82, 0xdc, 0x18, 0x9c, 0xd2,
	0x1f, 0xf3, 0x6e, 0x0b, 0x91, 0x60, 0xc8, 0x6d, 0x5e, 0x87, 0xf0, 0xc6, 0x3e, 0x9d, 0x3b, 0x32,
	0xb6, 0x0a, 0xe7, 0x0e, 0x96, 0xb6, 0x45, 0xfd, 0xb4, 0x45, 0x9d, 0x39, 0x48, 0xc3, 0x94, 0x38,
	0xf7, 0x6d, 0x1c, 0xaa, 0x3e, 0xaa, 0x8e, 0xe0, 0x0a, 0xd9, 0x25, 0x18, 0x4b, 0xcf, 0x38, 0x1b,
	0x7e, 0xdd, 0x1d, 0x34, 0x90, 0x3d, 0xff, 0xbb, 0xe6, 0xaf, 0x6f, 0x03, 0xa9, 0xc3, 0x9a, 0x39,
	0xdd, 0x81, 0xbe, 0x46, 0x67, 0x4b, 0x04, 0xbb, 0x43, 0xf0, 0xb0, 0x09, 0x36, 0x68, 0x9d, 0x2b,
	0xc6, 0x9f, 0x6a, 0x0f, 0xfe, 0xc0, 0x96, 0xa0, 0x92, 0xb9, 0x8a, 0x26, 0x65, 0x14, 0xcf, 0xfe,
	0xa6, 0x68, 0x3d, 0xb7, 0x61, 0xff, 0xfb, 0xbd, 0x70, 0x76, 0x0f, 0x4a, 0x52, 0xec, 0x39, 0x47,
	0x52, 0x6a, 0xc9, 0x2d, 0x7e, 0x0b, 0xdc, 0xde, 0x24, 0x5c, 0x5f, 0xec, 0xf9, 0x46, 0xa4, 0xfe,
	0xb5, 0x02, 0x25, 0x5a, 0xb0, 0x69, 0x28, 0xd3, 0xd2, 0xd8, 0xe5, 0x99, 0x47, 0xc3, 0x19, 0xf3,
	0xc7, 0x68, 0x49, 0x6e, 0x58, 0x86, 0x89, 0x8e, 0xc4, 0x08, 0x9b, 0x94, 0x9b, 0x06, 0xf7, 0xdc,
	0x3b, 0xdc, 0x4c, 0x79, 0x80, 0x5d, 0x87, 0x6a, 0x82, 0xbb, 0x98, 0x90, 0x71, 0x9d, 0x17, 0x05,
	0xe0, 0x4a, 0x1a, 0xed, 0x09, 0xb6, 0x08, 0x15, 0xe2, 0xb5, 0xe1, 0x5e, 0x16, 0xe0, 0xca, 0x26,
	0x98, 0xb0, 0x4d, 0x38, 0xd5, 0x14, 0x5d, 0x4e, 0xce, 0x69, 0xea, 0x78, 0x17, 0x03, 0xd5, 0x6d,
	0xda, 0x4a, 0x94, 0xf3, 0xaa, 0x80, 0xcc, 0xc9, 0x14, 0x6e, 0xa4, 0xec, 0x83, 0x7d, 0x94, 0x6d,
	0x81, 0x63, 0x55, 0x23, 0x4c, 0xe8, 0x3a, 0x47, 0x79, 0xd9, 0xd7, 0x05, 0x64, 0xa7, 0x53, 0x7a,
	0xc5, 0xc2, 0x39, 0xdd, 0x55, 0xa8, 0xa5, 0x17, 0x3a, 0x42, 0xd5, 0x94, 0x71, 0x27, 0xb5, 0xd7,
	0x1b, 0x6f, 0xd0, 0x9c, 0x99, 0x9e, 0xd2, 0x32, 0xe6, 0x2d, 0x2b, 0x37, 0x65, 0xa0, 0x95, 0x3e,
	0xc3, 0xae, 0x42, 0xd9, 0xf6, 0xeb, 0xbc, 0xf5, 0x0e, 0xbf, 0x75, 0x59, 0xac, 0x19, 0x71, 0xef,
	0x35, 0x7a, 0x57, 0x64, 0xc4, 0xdc, 0x3e, 0x45, 0xcb, 0x43, 0x2f, 0xd9, 0x7b, 0xef, 0xdf, 0x9e,
	0xb2, 0x35, 0xa8, 0xd9, 0x61, 0xe6, 0x86, 0xf8, 0xa1, 0x80, 0xc6, 0x54, 0x4a, 0xe5, 0xa6, 0x77,
	0x0d, 0x2a, 0xd9, 0x79, 0x38, 0x1f, 0x0b, 0xb4, 0xdd, 0x0b, 0x66, 0x0d, 0x38, 0x1a, 0xf3, 0xc7,
	0x28, 0x63, 0xad, 0x82, 0x1d, 0x29, 0xda, 0xce, 0xa7, 0x02, 0xd9, 0x27, 0x7b, 0xc8, 0x2a, 0x11,
	0x7d, 0x5b, 0x53, 0xff, 0x9f, 0x0b, 0xdb, 0x9a, 0x9a, 0xbf, 0x01, 0xe3, 0xea, 0x49, 0x7b, 0x5b,
	0xa4, 0xe4, 0x97, 0x02, 0x64, 0xd5, 0x86, 0xaf, 0x47, 0x77, 0x36, 0x61, 0x26, 0x16, 0x43, 0xb7,
	0xbd, 0xff, 0xc1, 0x7c, 0xb4, 0xf8, 0x5f, 0x9f, 0xd2, 0xed, 0x72, 0xfa, 0xb5, 0xba, 0xf2, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0x43, 0xbd, 0x92, 0x8a, 0x07, 0x00, 0x00,
}