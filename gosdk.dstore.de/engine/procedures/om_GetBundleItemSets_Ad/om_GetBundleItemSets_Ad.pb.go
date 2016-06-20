// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetBundleItemSets_Ad.proto
// DO NOT EDIT!

/*
Package om_GetBundleItemSets_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetBundleItemSets_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetBundleItemSets_Ad

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
	ItemSetId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	ItemSetIdNull       bool                        `protobuf:"varint,1001,opt,name=item_set_id_null,json=itemSetIdNull" json:"item_set_id_null,omitempty"`
	ItemConditionId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	ItemConditionIdNull bool                        `protobuf:"varint,1002,opt,name=item_condition_id_null,json=itemConditionIdNull" json:"item_condition_id_null,omitempty"`
	GetUnusedSets       *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=get_unused_sets,json=getUnusedSets" json:"get_unused_sets,omitempty"`
	GetUnusedSetsNull   bool                        `protobuf:"varint,1003,opt,name=get_unused_sets_null,json=getUnusedSetsNull" json:"get_unused_sets_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemSetId
	}
	return nil
}

func (m *Parameters) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Parameters) GetGetUnusedSets() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetUnusedSets
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
	RowId                         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ItemGroupSortNo               *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=item_group_sort_no,json=itemGroupSortNo" json:"item_group_sort_no,omitempty"`
	Operator1                     *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=operator1" json:"operator1,omitempty"`
	Condition1                    *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=condition1" json:"condition1,omitempty"`
	Condition2                    *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=condition2" json:"condition2,omitempty"`
	Operator2                     *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=operator2" json:"operator2,omitempty"`
	ItemConditionGroupDescription *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=item_condition_group_description,json=itemConditionGroupDescription" json:"item_condition_group_description,omitempty"`
	CombinePartsWithANDOperator   *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=combine_parts_with_a_n_d_operator,json=combinePartsWithANDOperator" json:"combine_parts_with_a_n_d_operator,omitempty"`
	RecursiveEvaluation           *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=recursive_evaluation,json=recursiveEvaluation" json:"recursive_evaluation,omitempty"`
	ItemPartSortNo                *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=item_part_sort_no,json=itemPartSortNo" json:"item_part_sort_no,omitempty"`
	ItemSetId                     *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	DomainTreeNodeIds             *dstore_values.StringValue  `protobuf:"bytes,10011,opt,name=domain_tree_node_ids,json=domainTreeNodeIds" json:"domain_tree_node_ids,omitempty"`
	NodeCharacteristicId          *dstore_values.IntegerValue `protobuf:"bytes,10012,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	Quantity                      *dstore_values.IntegerValue `protobuf:"bytes,10013,opt,name=quantity" json:"quantity,omitempty"`
	ItemConditionPartId           *dstore_values.IntegerValue `protobuf:"bytes,10014,opt,name=item_condition_part_id,json=itemConditionPartId" json:"item_condition_part_id,omitempty"`
	ItemConditionId               *dstore_values.IntegerValue `protobuf:"bytes,10015,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	CombineGroupsWithANDOperator  *dstore_values.BooleanValue `protobuf:"bytes,10016,opt,name=combine_groups_with_a_n_d_operator,json=combineGroupsWithANDOperator" json:"combine_groups_with_a_n_d_operator,omitempty"`
	ItemConditionPartDescription  *dstore_values.StringValue  `protobuf:"bytes,10017,opt,name=item_condition_part_description,json=itemConditionPartDescription" json:"item_condition_part_description,omitempty"`
	LevelIds                      *dstore_values.StringValue  `protobuf:"bytes,10018,opt,name=level_ids,json=levelIds" json:"level_ids,omitempty"`
	DistinctItemsOnly             *dstore_values.BooleanValue `protobuf:"bytes,10019,opt,name=distinct_items_only,json=distinctItemsOnly" json:"distinct_items_only,omitempty"`
	ItemConditionGroupId          *dstore_values.IntegerValue `protobuf:"bytes,10020,opt,name=item_condition_group_id,json=itemConditionGroupId" json:"item_condition_group_id,omitempty"`
	ItemConditionDescription      *dstore_values.StringValue  `protobuf:"bytes,10021,opt,name=item_condition_description,json=itemConditionDescription" json:"item_condition_description,omitempty"`
	InheritDepth                  *dstore_values.IntegerValue `protobuf:"bytes,10022,opt,name=inherit_depth,json=inheritDepth" json:"inherit_depth,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetItemGroupSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemGroupSortNo
	}
	return nil
}

func (m *Response_Row) GetOperator1() *dstore_values.StringValue {
	if m != nil {
		return m.Operator1
	}
	return nil
}

func (m *Response_Row) GetCondition1() *dstore_values.StringValue {
	if m != nil {
		return m.Condition1
	}
	return nil
}

func (m *Response_Row) GetCondition2() *dstore_values.StringValue {
	if m != nil {
		return m.Condition2
	}
	return nil
}

func (m *Response_Row) GetOperator2() *dstore_values.StringValue {
	if m != nil {
		return m.Operator2
	}
	return nil
}

func (m *Response_Row) GetItemConditionGroupDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionGroupDescription
	}
	return nil
}

func (m *Response_Row) GetCombinePartsWithANDOperator() *dstore_values.BooleanValue {
	if m != nil {
		return m.CombinePartsWithANDOperator
	}
	return nil
}

func (m *Response_Row) GetRecursiveEvaluation() *dstore_values.IntegerValue {
	if m != nil {
		return m.RecursiveEvaluation
	}
	return nil
}

func (m *Response_Row) GetItemPartSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemPartSortNo
	}
	return nil
}

func (m *Response_Row) GetItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemSetId
	}
	return nil
}

func (m *Response_Row) GetDomainTreeNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.DomainTreeNodeIds
	}
	return nil
}

func (m *Response_Row) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Response_Row) GetItemConditionPartId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionPartId
	}
	return nil
}

func (m *Response_Row) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Response_Row) GetCombineGroupsWithANDOperator() *dstore_values.BooleanValue {
	if m != nil {
		return m.CombineGroupsWithANDOperator
	}
	return nil
}

func (m *Response_Row) GetItemConditionPartDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionPartDescription
	}
	return nil
}

func (m *Response_Row) GetLevelIds() *dstore_values.StringValue {
	if m != nil {
		return m.LevelIds
	}
	return nil
}

func (m *Response_Row) GetDistinctItemsOnly() *dstore_values.BooleanValue {
	if m != nil {
		return m.DistinctItemsOnly
	}
	return nil
}

func (m *Response_Row) GetItemConditionGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionGroupId
	}
	return nil
}

func (m *Response_Row) GetItemConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionDescription
	}
	return nil
}

func (m *Response_Row) GetInheritDepth() *dstore_values.IntegerValue {
	if m != nil {
		return m.InheritDepth
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetBundleItemSets_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetBundleItemSets_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetBundleItemSets_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetBundleItemSets_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 918 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xeb, 0x6e, 0xdc, 0x44,
	0x14, 0x56, 0x1b, 0xd2, 0x26, 0x53, 0x42, 0x9a, 0x49, 0x14, 0xcc, 0xa6, 0x40, 0x09, 0x42, 0xea,
	0x2f, 0x87, 0x6e, 0x11, 0x4a, 0x05, 0x7f, 0xd2, 0x24, 0x0a, 0x2b, 0xb5, 0x4e, 0x71, 0xb8, 0x95,
	0x3f, 0x23, 0xaf, 0x67, 0xba, 0x3b, 0xc2, 0x3b, 0xb3, 0xcc, 0x8c, 0x13, 0xf5, 0x2d, 0xb8, 0xdf,
	0xef, 0xb7, 0xa7, 0xe0, 0x3d, 0x78, 0x86, 0xc2, 0x4b, 0x70, 0xc6, 0xe3, 0x75, 0x6c, 0xef, 0x22,
	0x7b, 0xf9, 0x93, 0x68, 0x76, 0xce, 0xf7, 0x9d, 0xcf, 0xe7, 0x7c, 0xe7, 0xd8, 0xe8, 0x55, 0xaa,
	0x8d, 0x54, 0x6c, 0x87, 0x89, 0x01, 0x17, 0x6c, 0x67, 0xac, 0x64, 0xcc, 0x68, 0xaa, 0x98, 0xde,
	0x91, 0x23, 0x72, 0xc4, 0xcc, 0x9d, 0x54, 0xd0, 0x84, 0xf5, 0x0c, 0x1b, 0x9d, 0x30, 0xa3, 0xc9,
	0x1e, 0xf5, 0x21, 0xc4, 0x48, 0xfc, 0x92, 0xc3, 0xf9, 0x0e, 0xe7, 0xff, 0x47, 0x70, 0x67, 0x3d,
	0xa7, 0x3f, 0x8d, 0x92, 0x94, 0x69, 0x87, 0xed, 0x3c, 0x53, 0xcd, 0xc9, 0x94, 0x92, 0x2a, 0xbf,
	0xda, 0xaa, 0x5e, 0x8d, 0x98, 0xd6, 0xd1, 0x80, 0xe5, 0x97, 0x2f, 0xd6, 0x2f, 0x4d, 0xc4, 0xc5,
	0x43, 0xa9, 0x46, 0x91, 0xe1, 0x52, 0xb8, 0xa0, 0xed, 0xc7, 0x17, 0x11, 0xba, 0x1f, 0xa9, 0x08,
	0x6e, 0x99, 0xd2, 0xf8, 0x35, 0x74, 0x85, 0x83, 0x1e, 0xa2, 0x99, 0x21, 0x9c, 0x7a, 0x17, 0xae,
	0x5f, 0xb8, 0x71, 0xa5, 0xbb, 0xe5, 0xe7, 0xea, 0x73, 0x59, 0x5c, 0x18, 0x36, 0x60, 0xea, 0x1d,
	0x7b, 0x0a, 0x97, 0xb9, 0xd3, 0xdf, 0xa3, 0xf8, 0x06, 0xba, 0x5a, 0x02, 0x13, 0x91, 0x26, 0x89,
	0xf7, 0xf8, 0x32, 0x50, 0x2c, 0x85, 0x2b, 0x45, 0x54, 0x00, 0xbf, 0xe2, 0x23, 0xb4, 0x96, 0x45,
	0xc6, 0x52, 0x50, 0x6e, 0xd5, 0xd8, 0x64, 0x17, 0x9b, 0x93, 0xad, 0x5a, 0xd4, 0xfe, 0x04, 0x04,
	0x29, 0x5f, 0x41, 0x9b, 0x53, 0x44, 0x2e, 0xf1, 0xdf, 0x2e, 0xf1, 0x7a, 0x0d, 0x91, 0xa5, 0xdf,
	0x47, 0xab, 0x03, 0xd0, 0x98, 0x8a, 0x54, 0x33, 0x6a, 0xe5, 0x6a, 0x6f, 0x61, 0x66, 0xf2, 0xbe,
	0x94, 0x09, 0x8b, 0x84, 0x4b, 0xbe, 0x02, 0x98, 0xb7, 0x33, 0x88, 0x6d, 0x17, 0x7e, 0x19, 0x6d,
	0xd4, 0x48, 0x5c, 0xe2, 0x7f, 0x5c, 0xe2, 0xb5, 0x4a, 0xb4, 0x4d, 0xbb, 0xfd, 0xd7, 0x2a, 0x5a,
	0x0a, 0x99, 0x1e, 0x4b, 0xa1, 0x19, 0xc0, 0x17, 0xb3, 0x4e, 0xe6, 0x35, 0xee, 0xf8, 0x55, 0x87,
	0xb8, 0x2e, 0x1f, 0xda, 0xbf, 0xa1, 0x0b, 0xc4, 0x0f, 0xd0, 0x55, 0xdb, 0x43, 0x52, 0x6a, 0x22,
	0xd4, 0x6c, 0x01, 0xc0, 0x7e, 0x0d, 0x5c, 0x6f, 0xf5, 0x3d, 0x38, 0xf7, 0xce, 0xcf, 0xe1, 0xea,
	0xa8, 0xfa, 0x03, 0xde, 0x45, 0x97, 0x73, 0xef, 0x40, 0x21, 0x2c, 0xe3, 0x73, 0x53, 0x8c, 0xce,
	0x59, 0xf7, 0xdc, 0xff, 0x70, 0x12, 0x8e, 0x0f, 0xd1, 0x82, 0x92, 0x67, 0xde, 0x13, 0x19, 0xea,
	0x96, 0xdf, 0xca, 0xe6, 0xfe, 0xa4, 0x08, 0x7e, 0x28, 0xcf, 0x42, 0x8b, 0xef, 0xfc, 0xb9, 0x82,
	0x16, 0xe0, 0x80, 0x37, 0xd1, 0x25, 0x38, 0x5a, 0x37, 0x7c, 0x14, 0x40, 0x5d, 0x16, 0xc3, 0x45,
	0x38, 0x42, 0x9f, 0x7b, 0x08, 0x67, 0x7d, 0x1e, 0x28, 0x99, 0x8e, 0x89, 0x96, 0xca, 0x10, 0x21,
	0xbd, 0x8f, 0x83, 0x96, 0x96, 0x39, 0xb2, 0xb0, 0x13, 0x40, 0x05, 0x12, 0xdf, 0x46, 0xcb, 0x72,
	0xcc, 0x54, 0x04, 0x88, 0x9b, 0xde, 0x27, 0x41, 0xb5, 0xfa, 0x39, 0x83, 0x36, 0x8a, 0x8b, 0x41,
	0x6e, 0xf0, 0x22, 0x1a, 0xa6, 0x03, 0x15, 0x46, 0xbb, 0xe9, 0x7d, 0xda, 0x8c, 0x2d, 0x85, 0x57,
	0xc0, 0x5d, 0xef, 0xb3, 0x79, 0xc0, 0xdd, 0xb2, 0xe8, 0xae, 0xf7, 0xf9, 0x1c, 0xa2, 0xbb, 0x98,
	0xa2, 0xeb, 0xb5, 0x11, 0x71, 0x45, 0xa4, 0x4c, 0xc7, 0x8a, 0x8f, 0x33, 0x1b, 0x7d, 0xd1, 0xcc,
	0xf8, 0x6c, 0x65, 0x90, 0xb2, 0x82, 0x1e, 0x9c, 0x33, 0x40, 0x96, 0x17, 0x62, 0x39, 0xea, 0x43,
	0xd7, 0xc9, 0x38, 0x52, 0xd0, 0xe6, 0x33, 0x6e, 0x86, 0x24, 0x22, 0x82, 0x50, 0x32, 0xd1, 0xe2,
	0x7d, 0x19, 0x34, 0x4f, 0xd9, 0x56, 0x4e, 0x73, 0xdf, 0xb2, 0xbc, 0x0b, 0x24, 0x7b, 0xc1, 0xc1,
	0x71, 0x4e, 0x80, 0x8f, 0xd1, 0x86, 0x62, 0x71, 0xaa, 0x34, 0x3f, 0x65, 0x84, 0x59, 0xb8, 0x1b,
	0x83, 0xaf, 0x5a, 0x18, 0x61, 0xbd, 0x40, 0x1e, 0x16, 0xc0, 0x62, 0x11, 0x59, 0xcd, 0x85, 0xad,
	0xbe, 0x6e, 0xc1, 0xf6, 0x94, 0x85, 0x59, 0x8d, 0xb9, 0xab, 0x5e, 0xaf, 0x2e, 0xce, 0x6f, 0x82,
	0xb9, 0x36, 0xe7, 0x5d, 0xb4, 0x41, 0xe5, 0x08, 0x66, 0x96, 0x18, 0xc5, 0x18, 0x68, 0xa0, 0x0c,
	0x58, 0xb4, 0xf7, 0x6d, 0x73, 0x5f, 0xd6, 0x1c, 0xf0, 0x2d, 0xc0, 0x05, 0x00, 0xeb, 0x51, 0x8d,
	0x43, 0xb4, 0x99, 0x31, 0xc4, 0x43, 0x58, 0xec, 0x31, 0xec, 0x75, 0xae, 0x0d, 0x8f, 0xad, 0xac,
	0xef, 0x5a, 0xc8, 0xda, 0xb0, 0xd8, 0xfd, 0x0a, 0x14, 0x14, 0xee, 0xa2, 0xa5, 0x0f, 0xd3, 0x48,
	0x18, 0x6e, 0x1e, 0x79, 0xdf, 0xb7, 0x60, 0x29, 0xa2, 0xf1, 0x9b, 0x53, 0x2b, 0x3a, 0x2b, 0x36,
	0xa8, 0xf9, 0xa1, 0x4d, 0xd7, 0x2a, 0xb6, 0xb3, 0x05, 0x07, 0x31, 0x6f, 0xcc, 0x7a, 0x7d, 0xfc,
	0x18, 0xfc, 0x8f, 0xf7, 0xc7, 0x43, 0xb4, 0x3d, 0xb1, 0x6d, 0x36, 0x15, 0xb3, 0x7d, 0xfb, 0x53,
	0x0b, 0xdf, 0x5e, 0xcb, 0x79, 0xb2, 0xc9, 0x98, 0x32, 0x6e, 0x1f, 0x3d, 0x3f, 0xab, 0x08, 0xe5,
	0x19, 0xfc, 0xb9, 0xb9, 0xd7, 0xd7, 0xa6, 0x8a, 0x51, 0x1e, 0xc1, 0x5d, 0xb4, 0x9c, 0xb0, 0x53,
	0x96, 0x64, 0xce, 0xf9, 0xa5, 0x99, 0x6d, 0x29, 0x8b, 0xb6, 0x86, 0xb9, 0x8b, 0xd6, 0xa9, 0x6d,
	0xb4, 0x88, 0xa1, 0x2f, 0x90, 0x42, 0x13, 0x29, 0x92, 0x47, 0xde, 0xaf, 0x2d, 0x1e, 0x7b, 0x6d,
	0x02, 0xb4, 0xfb, 0x5d, 0x1f, 0x03, 0x0c, 0x9f, 0xa0, 0xa7, 0x67, 0x2e, 0x1c, 0xe8, 0xd1, 0x6f,
	0x6d, 0xfc, 0x37, 0xbd, 0x68, 0xa0, 0x51, 0x0f, 0x50, 0xa7, 0x46, 0x5a, 0xae, 0xdd, 0xef, 0xcd,
	0x4f, 0xeb, 0x55, 0x68, 0xcb, 0x75, 0xdb, 0x43, 0x2b, 0x5c, 0x0c, 0xc1, 0xea, 0xb6, 0x1f, 0x63,
	0x33, 0xf4, 0xfe, 0x68, 0xa1, 0xf2, 0xc9, 0x1c, 0x72, 0x60, 0x11, 0x77, 0xde, 0x43, 0x5b, 0x5c,
	0xd6, 0x5e, 0x7e, 0xe7, 0xdf, 0x86, 0xef, 0xdf, 0x1e, 0x48, 0x4d, 0x3f, 0x98, 0xdc, 0xd3, 0x39,
	0x3e, 0x1f, 0xfb, 0x97, 0xb2, 0xcf, 0xb4, 0x5b, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x23,
	0x09, 0x02, 0x79, 0x0a, 0x00, 0x00,
}