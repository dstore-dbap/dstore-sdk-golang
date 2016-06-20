// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_FuzzySearch_Ad.proto
// DO NOT EDIT!

/*
Package im_FuzzySearch_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_FuzzySearch_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_FuzzySearch_Ad

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
	NodeCharacteristicId         *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull     bool                          `protobuf:"varint,1001,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	FuzzyType                    *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=fuzzy_type,json=fuzzyType" json:"fuzzy_type,omitempty"`
	FuzzyTypeNull                bool                          `protobuf:"varint,1002,opt,name=fuzzy_type_null,json=fuzzyTypeNull" json:"fuzzy_type_null,omitempty"`
	IncludeInactiveNodes         *dstore_values.BooleanValue   `protobuf:"bytes,3,opt,name=include_inactive_nodes,json=includeInactiveNodes" json:"include_inactive_nodes,omitempty"`
	IncludeInactiveNodesNull     bool                          `protobuf:"varint,1003,opt,name=include_inactive_nodes_null,json=includeInactiveNodesNull" json:"include_inactive_nodes_null,omitempty"`
	FromDate                     *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                 bool                          `protobuf:"varint,1004,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                       *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                   bool                          `protobuf:"varint,1005,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	LevelId                      *dstore_values.IntegerValue   `protobuf:"bytes,6,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	LevelIdNull                  bool                          `protobuf:"varint,1006,opt,name=level_id_null,json=levelIdNull" json:"level_id_null,omitempty"`
	DomainTreeNodeId             *dstore_values.IntegerValue   `protobuf:"bytes,7,opt,name=domain_tree_node_id,json=domainTreeNodeId" json:"domain_tree_node_id,omitempty"`
	DomainTreeNodeIdNull         bool                          `protobuf:"varint,1007,opt,name=domain_tree_node_id_null,json=domainTreeNodeIdNull" json:"domain_tree_node_id_null,omitempty"`
	FilterByCharacteristicId     *dstore_values.IntegerValue   `protobuf:"bytes,8,opt,name=filter_by_characteristic_id,json=filterByCharacteristicId" json:"filter_by_characteristic_id,omitempty"`
	FilterByCharacteristicIdNull bool                          `protobuf:"varint,1008,opt,name=filter_by_characteristic_id_null,json=filterByCharacteristicIdNull" json:"filter_by_characteristic_id_null,omitempty"`
	FilterByCharacValue          *dstore_values.StringValue    `protobuf:"bytes,9,opt,name=filter_by_charac_value,json=filterByCharacValue" json:"filter_by_charac_value,omitempty"`
	FilterByCharacValueNull      bool                          `protobuf:"varint,1009,opt,name=filter_by_charac_value_null,json=filterByCharacValueNull" json:"filter_by_charac_value_null,omitempty"`
	StartAtRowNo                 *dstore_values.IntegerValue   `protobuf:"bytes,10,opt,name=start_at_row_no,json=startAtRowNo" json:"start_at_row_no,omitempty"`
	StartAtRowNoNull             bool                          `protobuf:"varint,1010,opt,name=start_at_row_no_null,json=startAtRowNoNull" json:"start_at_row_no_null,omitempty"`
	RowCount                     *dstore_values.IntegerValue   `protobuf:"bytes,11,opt,name=row_count,json=rowCount" json:"row_count,omitempty"`
	RowCountNull                 bool                          `protobuf:"varint,1011,opt,name=row_count_null,json=rowCountNull" json:"row_count_null,omitempty"`
	NegateFilterByParams         *dstore_values.BooleanValue   `protobuf:"bytes,12,opt,name=negate_filter_by_params,json=negateFilterByParams" json:"negate_filter_by_params,omitempty"`
	NegateFilterByParamsNull     bool                          `protobuf:"varint,1012,opt,name=negate_filter_by_params_null,json=negateFilterByParamsNull" json:"negate_filter_by_params_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetFuzzyType() *dstore_values.IntegerValue {
	if m != nil {
		return m.FuzzyType
	}
	return nil
}

func (m *Parameters) GetIncludeInactiveNodes() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeInactiveNodes
	}
	return nil
}

func (m *Parameters) GetFromDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Parameters) GetToDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Parameters) GetLevelId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelId
	}
	return nil
}

func (m *Parameters) GetDomainTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.DomainTreeNodeId
	}
	return nil
}

func (m *Parameters) GetFilterByCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilterByCharacteristicId
	}
	return nil
}

func (m *Parameters) GetFilterByCharacValue() *dstore_values.StringValue {
	if m != nil {
		return m.FilterByCharacValue
	}
	return nil
}

func (m *Parameters) GetStartAtRowNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.StartAtRowNo
	}
	return nil
}

func (m *Parameters) GetRowCount() *dstore_values.IntegerValue {
	if m != nil {
		return m.RowCount
	}
	return nil
}

func (m *Parameters) GetNegateFilterByParams() *dstore_values.BooleanValue {
	if m != nil {
		return m.NegateFilterByParams
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
	RowId                      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	NodeDescription            *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Active                     *dstore_values.BooleanValue `protobuf:"bytes,10002,opt,name=active" json:"active,omitempty"`
	PredecessorsLevelNo        *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=predecessors_level_no,json=predecessorsLevelNo" json:"predecessors_level_no,omitempty"`
	PrePredecessorsLevelNo     *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=pre_predecessors_level_no,json=prePredecessorsLevelNo" json:"pre_predecessors_level_no,omitempty"`
	TreeNodeId                 *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                     *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	PrePredecessorsDescription *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=pre_predecessors_description,json=prePredecessorsDescription" json:"pre_predecessors_description,omitempty"`
	PrePredecessorsTreeNodeId  *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=pre_predecessors_tree_node_id,json=prePredecessorsTreeNodeId" json:"pre_predecessors_tree_node_id,omitempty"`
	PredecessorsTreeNodeId     *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=predecessors_tree_node_id,json=predecessorsTreeNodeId" json:"predecessors_tree_node_id,omitempty"`
	PredecessorsDescription    *dstore_values.StringValue  `protobuf:"bytes,10010,opt,name=predecessors_description,json=predecessorsDescription" json:"predecessors_description,omitempty"`
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

func (m *Response_Row) GetPredecessorsLevelNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredecessorsLevelNo
	}
	return nil
}

func (m *Response_Row) GetPrePredecessorsLevelNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PrePredecessorsLevelNo
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

func (m *Response_Row) GetPrePredecessorsDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PrePredecessorsDescription
	}
	return nil
}

func (m *Response_Row) GetPrePredecessorsTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PrePredecessorsTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetPredecessorsTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredecessorsTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetPredecessorsDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PredecessorsDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_FuzzySearch_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_FuzzySearch_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_FuzzySearch_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 968 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0x7b, 0x8f, 0xdb, 0x44,
	0x10, 0x57, 0x39, 0x2e, 0xc9, 0xcd, 0x5d, 0xb9, 0x93, 0xef, 0xc8, 0xf9, 0x92, 0x2b, 0x94, 0x56,
	0x08, 0x84, 0x84, 0x53, 0xf1, 0x16, 0x12, 0x45, 0xbd, 0x96, 0x43, 0x41, 0x34, 0x1c, 0xa6, 0x2a,
	0xa2, 0x12, 0x58, 0x3e, 0x7b, 0x93, 0x5a, 0x72, 0xbc, 0xd6, 0x7a, 0x73, 0x55, 0xee, 0x53, 0xf0,
	0x7e, 0x7d, 0x07, 0xbe, 0x04, 0x5f, 0x82, 0xcf, 0xc0, 0xfb, 0xf9, 0x01, 0x98, 0xdd, 0xd9, 0xbc,
	0x1c, 0xa7, 0xf6, 0x3f, 0x89, 0xec, 0x9d, 0xdf, 0x63, 0x66, 0x77, 0x67, 0x0c, 0xd7, 0xc2, 0x4c,
	0x72, 0xc1, 0x3a, 0x2c, 0x19, 0x44, 0x09, 0xeb, 0xa4, 0x82, 0x07, 0x2c, 0x1c, 0x09, 0x96, 0x75,
	0xa2, 0xa1, 0x77, 0x3c, 0x3a, 0x3f, 0x1f, 0x7f, 0xc0, 0x7c, 0x11, 0xdc, 0xf7, 0x6e, 0x84, 0x0e,
	0x2e, 0x4a, 0x6e, 0x3d, 0x49, 0x08, 0x87, 0x10, 0xce, 0x52, 0x58, 0x6b, 0xd7, 0x50, 0x9e, 0xf9,
	0xf1, 0x88, 0x65, 0x84, 0x6a, 0x1d, 0x2c, 0xea, 0x30, 0x21, 0xb8, 0x30, 0x4b, 0xed, 0xc5, 0xa5,
	0x21, 0xcb, 0x32, 0x7f, 0xc0, 0xcc, 0xe2, 0xd5, 0xfc, 0xa2, 0xf4, 0xa3, 0xa4, 0xcf, 0xc5, 0xd0,
	0x97, 0x11, 0x4f, 0x28, 0xe8, 0xca, 0x4f, 0x9b, 0x00, 0x27, 0xbe, 0xf0, 0x71, 0x95, 0x89, 0xcc,
	0x7a, 0x1f, 0x9a, 0x09, 0x0f, 0x99, 0x17, 0xdc, 0xc7, 0x77, 0x01, 0xbe, 0x8a, 0x32, 0x19, 0x05,
	0x5e, 0x14, 0xda, 0x17, 0x2e, 0x5f, 0x78, 0x76, 0xf3, 0x85, 0xb6, 0x63, 0x52, 0x30, 0x0e, 0xa3,
	0x44, 0xb2, 0x01, 0x13, 0x77, 0xd5, 0x93, 0xbb, 0xa7, 0xa0, 0x37, 0x17, 0x90, 0xdd, 0xd0, 0xba,
	0x0e, 0xed, 0x62, 0x4a, 0x2f, 0x19, 0xc5, 0xb1, 0xfd, 0x73, 0x1d, 0x89, 0x1b, 0xae, 0x5d, 0x84,
	0xed, 0x61, 0x80, 0xf5, 0x3a, 0x40, 0x5f, 0x55, 0xc9, 0x93, 0xe3, 0x94, 0xd9, 0x8f, 0x94, 0xdb,
	0xd8, 0xd0, 0xe1, 0x77, 0x30, 0xda, 0x7a, 0x06, 0xb6, 0x67, 0x58, 0xd2, 0xfb, 0x85, 0xf4, 0x2e,
	0x4e, 0x83, 0xb4, 0x08, 0xe6, 0x1d, 0x25, 0x41, 0x3c, 0x42, 0x9f, 0x51, 0x82, 0x16, 0xa2, 0x33,
	0x0c, 0x47, 0x47, 0x99, 0xbd, 0x56, 0x28, 0x78, 0xca, 0x79, 0xcc, 0xfc, 0xc4, 0xe4, 0x6d, 0xa0,
	0x5d, 0x83, 0xec, 0x29, 0xa0, 0xca, 0xbb, 0x98, 0x92, 0x7c, 0xfc, 0x6a, 0xf2, 0x2e, 0xc2, 0x9a,
	0xbc, 0x37, 0xfa, 0x82, 0x0f, 0xbd, 0xd0, 0x97, 0xcc, 0x7e, 0x54, 0xbb, 0xb8, 0x94, 0x73, 0x21,
	0x23, 0xdc, 0x71, 0xe9, 0x0f, 0x53, 0xf2, 0xd1, 0x50, 0xf1, 0xb7, 0x30, 0xdc, 0x7a, 0x1a, 0x1e,
	0x9b, 0x62, 0x49, 0xee, 0x37, 0x92, 0xdb, 0x9a, 0x84, 0x68, 0x89, 0x57, 0xa0, 0x2e, 0x39, 0x09,
	0xac, 0x57, 0x11, 0xa8, 0x49, 0xae, 0xe9, 0x9f, 0x82, 0x2d, 0x83, 0x23, 0xf2, 0xdf, 0x89, 0x1c,
	0x68, 0xd9, 0x50, 0x37, 0x62, 0x76, 0xc6, 0x62, 0x75, 0x74, 0x6a, 0xe5, 0x7b, 0x56, 0xd7, 0xc1,
	0x78, 0x5a, 0xae, 0xc2, 0xc5, 0x09, 0x8e, 0xb8, 0xff, 0x20, 0xee, 0x4d, 0x13, 0xa0, 0xc9, 0xdf,
	0x81, 0xdd, 0x90, 0x0f, 0xf1, 0x34, 0x7b, 0x52, 0x30, 0xaa, 0xaa, 0xd2, 0xa9, 0x97, 0xeb, 0xec,
	0x10, 0xee, 0x0e, 0xc2, 0x54, 0xa5, 0x51, 0xf0, 0x55, 0xb0, 0x0b, 0xb8, 0x48, 0xfb, 0x4f, 0xd2,
	0xde, 0xcb, 0x83, 0xb4, 0x89, 0x7b, 0xd0, 0xee, 0x47, 0x31, 0x1e, 0x56, 0xef, 0x74, 0x5c, 0x70,
	0x5f, 0x1a, 0xe5, 0x66, 0x6c, 0xc2, 0x1f, 0x8d, 0x97, 0xee, 0xcc, 0xdb, 0x70, 0xf9, 0x21, 0xdc,
	0x64, 0xee, 0x2f, 0x32, 0x77, 0xb8, 0x8a, 0x44, 0x9b, 0x7c, 0x0f, 0x9a, 0x79, 0x22, 0x4f, 0x5b,
	0xb1, 0x37, 0xb4, 0xbf, 0x56, 0xce, 0x5f, 0x26, 0x45, 0x94, 0x0c, 0xc8, 0xde, 0xee, 0x22, 0xb3,
	0x7e, 0x69, 0xbd, 0xb1, 0x9c, 0x35, 0x11, 0x92, 0xa9, 0xbf, 0xc9, 0xd4, 0x7e, 0x01, 0x54, 0xfb,
	0x39, 0x82, 0x6d, 0x3c, 0x4f, 0x42, 0x7a, 0xbe, 0xf4, 0x04, 0x7f, 0x80, 0xe5, 0xb6, 0xa1, 0xbc,
	0x50, 0x5b, 0x1a, 0x73, 0x43, 0xba, 0xfc, 0x41, 0x8f, 0x5b, 0x1d, 0xd8, 0xcb, 0x71, 0x90, 0xf6,
	0x3f, 0xa4, 0xbd, 0x33, 0x1f, 0xac, 0x45, 0x5f, 0x83, 0x0d, 0x15, 0x17, 0xf0, 0x51, 0x22, 0xed,
	0xcd, 0x72, 0xb9, 0x06, 0x46, 0xdf, 0x54, 0xc1, 0xea, 0x1e, 0x4d, 0x91, 0x24, 0xf2, 0xaf, 0xb9,
	0x47, 0x93, 0x10, 0x2d, 0xe0, 0xc2, 0x7e, 0xc2, 0x06, 0xea, 0x3a, 0xcc, 0x6a, 0x93, 0xaa, 0xa6,
	0x9a, 0xd9, 0x5b, 0x15, 0xda, 0x07, 0x61, 0x8f, 0x4d, 0xc9, 0x74, 0x37, 0xce, 0xac, 0x37, 0xe1,
	0x70, 0x05, 0x27, 0x19, 0xf9, 0x6f, 0xd2, 0x37, 0x0b, 0xc0, 0xca, 0xd4, 0x95, 0x1f, 0x1b, 0xd0,
	0x70, 0x59, 0x96, 0xf2, 0x24, 0x63, 0xd6, 0x35, 0x58, 0xd7, 0x73, 0xc3, 0xb4, 0xf1, 0xe9, 0xb6,
	0x9b, 0x49, 0x44, 0x33, 0xe5, 0x2d, 0xf5, 0xeb, 0x52, 0xa0, 0xf5, 0x11, 0xec, 0xa8, 0x89, 0xe1,
	0xcd, 0x8d, 0x0c, 0x6c, 0xbe, 0x6b, 0x08, 0x76, 0x72, 0xe0, 0xfc, 0x60, 0xb9, 0x8d, 0xcf, 0xdd,
	0xd9, 0xb3, 0xbb, 0x3d, 0x5c, 0x7c, 0x81, 0xfb, 0x51, 0x37, 0x93, 0x0a, 0xbb, 0xab, 0x62, 0x7c,
	0x62, 0x89, 0x91, 0xe6, 0xd8, 0x6d, 0xfa, 0x77, 0x27, 0xe1, 0x58, 0x94, 0x35, 0x2c, 0x3c, 0x76,
	0x43, 0x85, 0x7a, 0xde, 0x29, 0x19, 0xa7, 0xce, 0x24, 0x7d, 0x07, 0xcf, 0x82, 0xab, 0x90, 0xad,
	0x1f, 0x6a, 0xb0, 0x86, 0x0f, 0x56, 0x13, 0x6a, 0x6a, 0x63, 0xf1, 0x9e, 0x7e, 0xda, 0xc3, 0x8a,
	0xac, 0xbb, 0xeb, 0xf8, 0x88, 0x17, 0xef, 0x18, 0x76, 0x74, 0x07, 0xc0, 0x26, 0x1c, 0x88, 0x28,
	0xd5, 0x59, 0x7f, 0xd6, 0x2b, 0xbd, 0x2a, 0xdb, 0x0a, 0x74, 0x6b, 0x86, 0xb1, 0x5e, 0x82, 0x1a,
	0xf5, 0x73, 0xfb, 0xf3, 0x5e, 0xf9, 0x09, 0x30, 0xb1, 0xd6, 0x09, 0x3c, 0x9e, 0x0a, 0x16, 0xb2,
	0x00, 0xd3, 0xe5, 0x22, 0xf3, 0xa8, 0x13, 0xe2, 0x1d, 0xf9, 0xa2, 0x57, 0x7e, 0x6a, 0x77, 0xe7,
	0xa1, 0xef, 0x2a, 0x24, 0xde, 0x95, 0x0f, 0xe1, 0x00, 0x5f, 0x7b, 0xc5, 0xac, 0x5f, 0x56, 0x60,
	0x6d, 0x22, 0xf4, 0xa4, 0x80, 0xf8, 0x3a, 0x8e, 0x80, 0xf9, 0xde, 0xfb, 0x55, 0x05, 0x2e, 0x90,
	0xb3, 0xb6, 0xfb, 0x32, 0xd4, 0x27, 0xd0, 0xaf, 0x2b, 0x40, 0x6b, 0x09, 0xc1, 0x3e, 0x86, 0xc3,
	0xa5, 0x7c, 0xe6, 0xf7, 0xea, 0x9b, 0xf2, 0xbd, 0x6a, 0xe5, 0x32, 0x9a, 0xdf, 0xb6, 0x4f, 0xe0,
	0xd2, 0x12, 0xfd, 0x42, 0x9a, 0xdf, 0x56, 0xf0, 0x7a, 0x90, 0x13, 0x98, 0x1b, 0x36, 0xb4, 0x1d,
	0x2b, 0xb8, 0xbf, 0xab, 0xb8, 0x1d, 0x45, 0xc4, 0x77, 0xc1, 0x5e, 0x59, 0x93, 0xef, 0xcb, 0x6b,
	0xb2, 0x9f, 0x16, 0x17, 0xe4, 0xa8, 0x8b, 0x1f, 0x31, 0x3c, 0x77, 0xcf, 0x66, 0x1f, 0xba, 0xf7,
	0x9e, 0xab, 0xfe, 0x09, 0x7c, 0x5a, 0xd3, 0x1f, 0x9c, 0x2f, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0x27, 0xc9, 0x7f, 0x7b, 0x37, 0x0b, 0x00, 0x00,
}