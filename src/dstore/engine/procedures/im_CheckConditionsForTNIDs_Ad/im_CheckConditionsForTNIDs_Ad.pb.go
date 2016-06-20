// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_CheckConditionsForTNIDs_Ad.proto
// DO NOT EDIT!

/*
Package im_CheckConditionsForTNIDs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_CheckConditionsForTNIDs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_CheckConditionsForTNIDs_Ad

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
	TreeNodeIdList       *dstore_values.StringValue `protobuf:"bytes,1,opt,name=tree_node_id_list,json=treeNodeIdList" json:"tree_node_id_list,omitempty"`
	TreeNodeIdListNull   bool                       `protobuf:"varint,1001,opt,name=tree_node_id_list_null,json=treeNodeIdListNull" json:"tree_node_id_list_null,omitempty"`
	ItemConditionIds     *dstore_values.StringValue `protobuf:"bytes,2,opt,name=item_condition_ids,json=itemConditionIds" json:"item_condition_ids,omitempty"`
	ItemConditionIdsNull bool                       `protobuf:"varint,1002,opt,name=item_condition_ids_null,json=itemConditionIdsNull" json:"item_condition_ids_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeIdList() *dstore_values.StringValue {
	if m != nil {
		return m.TreeNodeIdList
	}
	return nil
}

func (m *Parameters) GetItemConditionIds() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionIds
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
	RowId                    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ItemConditionId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	TNIdsFulfillingCondition *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=t_n_ids_fulfilling_condition,json=tNIdsFulfillingCondition" json:"t_n_ids_fulfilling_condition,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Response_Row) GetTNIdsFulfillingCondition() *dstore_values.StringValue {
	if m != nil {
		return m.TNIdsFulfillingCondition
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_CheckConditionsForTNIDs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_CheckConditionsForTNIDs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_CheckConditionsForTNIDs_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xd1, 0x6a, 0x14, 0x31,
	0x14, 0x86, 0xd9, 0x8e, 0xdb, 0x96, 0x23, 0xd8, 0x36, 0x4a, 0x8d, 0xbb, 0x22, 0x52, 0x6f, 0x04,
	0x61, 0x56, 0x2c, 0x14, 0x05, 0x6f, 0xb4, 0xb6, 0x38, 0x68, 0x07, 0x09, 0x2a, 0xa8, 0x17, 0xc3,
	0xb8, 0x39, 0xbb, 0x06, 0x67, 0x92, 0x92, 0x64, 0xed, 0x3b, 0x78, 0xa5, 0x3e, 0x93, 0xaf, 0xe2,
	0x85, 0x3e, 0x85, 0x67, 0x26, 0xb3, 0xbb, 0xee, 0x0c, 0x58, 0xf4, 0x66, 0x97, 0x33, 0x27, 0xdf,
	0x7f, 0x92, 0xff, 0x4f, 0xe0, 0xa1, 0x74, 0xde, 0x58, 0x1c, 0xa1, 0x9e, 0x2a, 0x8d, 0xa3, 0x53,
	0x6b, 0xc6, 0x28, 0x67, 0x16, 0xdd, 0x48, 0x95, 0xd9, 0xe1, 0x07, 0x1c, 0x7f, 0x3c, 0x34, 0x5a,
	0x2a, 0xaf, 0x8c, 0x76, 0xc7, 0xc6, 0xbe, 0x4c, 0x93, 0x27, 0x2e, 0x7b, 0x24, 0x63, 0x5a, 0xe8,
	0x0d, 0xbb, 0x13, 0xe8, 0x38, 0xd0, 0xf1, 0x5f, 0x91, 0xc1, 0xe5, 0x66, 0xd4, 0xa7, 0xbc, 0x98,
	0xa1, 0x0b, 0x0a, 0x83, 0x6b, 0xab, 0xf3, 0xd1, 0x5a, 0x63, 0x9b, 0xd6, 0x70, 0xb5, 0x55, 0xa2,
	0x73, 0xf9, 0x14, 0x9b, 0xe6, 0xad, 0x76, 0xd3, 0xe7, 0x4a, 0x4f, 0x8c, 0x2d, 0xf3, 0x6a, 0x6e,
	0x58, 0xb4, 0xf7, 0x79, 0x0d, 0xe0, 0x45, 0x6e, 0x73, 0xea, 0xa2, 0x75, 0xec, 0x08, 0x76, 0xbc,
	0x45, 0xcc, 0xb4, 0x91, 0x98, 0x29, 0x99, 0x15, 0xca, 0x79, 0xde, 0xbb, 0xd9, 0xbb, 0x7d, 0xf1,
	0xde, 0x20, 0x6e, 0x4e, 0xd2, 0x6c, 0xce, 0x79, 0xab, 0xf4, 0xf4, 0x75, 0x55, 0x88, 0x4b, 0x15,
	0x94, 0x12, 0x93, 0xc8, 0xe7, 0x44, 0xb0, 0x7d, 0xd8, 0xed, 0xc8, 0x64, 0x7a, 0x56, 0x14, 0xfc,
	0xe7, 0x06, 0x89, 0x6d, 0x0a, 0xb6, 0x0a, 0xa4, 0xd4, 0x62, 0x4f, 0x81, 0x29, 0x8f, 0x65, 0x36,
	0x9e, 0x5b, 0x43, 0xa4, 0xe3, 0x6b, 0xe7, 0x0e, 0xdf, 0xae, 0xa8, 0x85, 0x9f, 0x89, 0x74, 0xec,
	0x00, 0xae, 0x76, 0x95, 0xc2, 0xfc, 0x5f, 0x61, 0xfe, 0x95, 0x36, 0x53, 0xed, 0x60, 0xef, 0x47,
	0x04, 0x9b, 0x02, 0xdd, 0x29, 0xa5, 0x82, 0xec, 0x2e, 0xf4, 0x6b, 0xab, 0xdb, 0xc7, 0x6f, 0x82,
	0x0c, 0x31, 0x1c, 0x55, 0xbf, 0x22, 0x2c, 0x64, 0x6f, 0x60, 0xbb, 0x32, 0x39, 0xfb, 0xc3, 0x65,
	0xda, 0x7e, 0x44, 0x70, 0xdc, 0x82, 0xdb, 0x59, 0x9c, 0x50, 0x9d, 0x2c, 0x6b, 0xb1, 0x55, 0xae,
	0x7e, 0x60, 0xf7, 0x61, 0xa3, 0x09, 0x97, 0x47, 0xb5, 0xe2, 0x8d, 0x8e, 0x62, 0x88, 0xfe, 0x24,
	0xfc, 0x8b, 0xf9, 0x72, 0xf6, 0x0c, 0x22, 0x6b, 0xce, 0xf8, 0x85, 0x9a, 0x7a, 0x10, 0xff, 0xc3,
	0x6d, 0x8c, 0xe7, 0x56, 0xc4, 0xc2, 0x9c, 0x89, 0x4a, 0x65, 0xf0, 0xbd, 0x07, 0x11, 0x15, 0x6c,
	0x17, 0xd6, 0xa9, 0x24, 0x57, 0xf9, 0x97, 0x94, 0xdc, 0xe9, 0x8b, 0x3e, 0x95, 0x89, 0xa4, 0x08,
	0x77, 0x3a, 0xc6, 0xf3, 0xaf, 0x69, 0x6d, 0xe0, 0xb0, 0x15, 0xa1, 0xd2, 0x1e, 0xa7, 0x68, 0x43,
	0x86, 0x5b, 0xad, 0x3c, 0xd8, 0x3b, 0xb8, 0x4e, 0x37, 0xa6, 0xce, 0x6d, 0x32, 0x2b, 0x26, 0xaa,
	0x28, 0x28, 0xef, 0xa5, 0x2e, 0xff, 0x96, 0x9e, 0x7b, 0x2f, 0xb8, 0x4f, 0x29, 0xd8, 0xe3, 0x05,
	0xbd, 0x90, 0x7f, 0xfc, 0x0a, 0x86, 0xca, 0xb4, 0xac, 0x58, 0x3e, 0xeb, 0xb7, 0x07, 0xff, 0xf7,
	0xe0, 0xdf, 0xaf, 0xd7, 0x4f, 0x6a, 0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xc3, 0x0c,
	0x54, 0x31, 0x04, 0x00, 0x00,
}