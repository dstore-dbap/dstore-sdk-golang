// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetTreeNodePWStatistics_Ad.proto
// DO NOT EDIT!

/*
Package st_GetTreeNodePWStatistics_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetTreeNodePWStatistics_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetTreeNodePWStatistics_Ad

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
	FromWeek                       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=from_week,json=fromWeek" json:"from_week,omitempty"`
	FromWeekNull                   bool                        `protobuf:"varint,1001,opt,name=from_week_null,json=fromWeekNull" json:"from_week_null,omitempty"`
	FromYear                       *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=from_year,json=fromYear" json:"from_year,omitempty"`
	FromYearNull                   bool                        `protobuf:"varint,1002,opt,name=from_year_null,json=fromYearNull" json:"from_year_null,omitempty"`
	ToWeek                         *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=to_week,json=toWeek" json:"to_week,omitempty"`
	ToWeekNull                     bool                        `protobuf:"varint,1003,opt,name=to_week_null,json=toWeekNull" json:"to_week_null,omitempty"`
	ToYear                         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=to_year,json=toYear" json:"to_year,omitempty"`
	ToYearNull                     bool                        `protobuf:"varint,1004,opt,name=to_year_null,json=toYearNull" json:"to_year_null,omitempty"`
	BasicCharacteristicNumbers     *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=basic_characteristic_numbers,json=basicCharacteristicNumbers" json:"basic_characteristic_numbers,omitempty"`
	BasicCharacteristicNumbersNull bool                        `protobuf:"varint,1005,opt,name=basic_characteristic_numbers_null,json=basicCharacteristicNumbersNull" json:"basic_characteristic_numbers_null,omitempty"`
	HTreeNodeIds                   *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=h_tree_node_ids,json=hTreeNodeIds" json:"h_tree_node_ids,omitempty"`
	HTreeNodeIdsNull               bool                        `protobuf:"varint,1006,opt,name=h_tree_node_ids_null,json=hTreeNodeIdsNull" json:"h_tree_node_ids_null,omitempty"`
	SummarizeWeeks                 *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=summarize_weeks,json=summarizeWeeks" json:"summarize_weeks,omitempty"`
	SummarizeWeeksNull             bool                        `protobuf:"varint,1007,opt,name=summarize_weeks_null,json=summarizeWeeksNull" json:"summarize_weeks_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetFromWeek() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromWeek
	}
	return nil
}

func (m *Parameters) GetFromYear() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromYear
	}
	return nil
}

func (m *Parameters) GetToWeek() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToWeek
	}
	return nil
}

func (m *Parameters) GetToYear() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToYear
	}
	return nil
}

func (m *Parameters) GetBasicCharacteristicNumbers() *dstore_values.StringValue {
	if m != nil {
		return m.BasicCharacteristicNumbers
	}
	return nil
}

func (m *Parameters) GetHTreeNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.HTreeNodeIds
	}
	return nil
}

func (m *Parameters) GetSummarizeWeeks() *dstore_values.BooleanValue {
	if m != nil {
		return m.SummarizeWeeks
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TotalValue                *dstore_values.DecimalValue `protobuf:"bytes,10001,opt,name=total_value,json=totalValue" json:"total_value,omitempty"`
	Year                      *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=year" json:"year,omitempty"`
	HTreeNodeId               *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	DirectValue               *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=direct_value,json=directValue" json:"direct_value,omitempty"`
	Week                      *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=week" json:"week,omitempty"`
	BasicCharacteristicNumber *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=basic_characteristic_number,json=basicCharacteristicNumber" json:"basic_characteristic_number,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTotalValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalValue
	}
	return nil
}

func (m *Response_Row) GetYear() *dstore_values.IntegerValue {
	if m != nil {
		return m.Year
	}
	return nil
}

func (m *Response_Row) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetDirectValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.DirectValue
	}
	return nil
}

func (m *Response_Row) GetWeek() *dstore_values.IntegerValue {
	if m != nil {
		return m.Week
	}
	return nil
}

func (m *Response_Row) GetBasicCharacteristicNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicCharacteristicNumber
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetTreeNodePWStatistics_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetTreeNodePWStatistics_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetTreeNodePWStatistics_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetTreeNodePWStatistics_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0xa6, 0x49, 0xea, 0x24, 0xb4, 0x65, 0x2c, 0xb2, 0x4d, 0xa4, 0xd8, 0x8a, 0x20, 0x08,
	0x1b, 0x6f, 0x3f, 0x14, 0x94, 0x52, 0x2f, 0x48, 0x95, 0x86, 0xb2, 0x8a, 0xa5, 0x22, 0x2e, 0x9b,
	0xdd, 0xd3, 0x74, 0x69, 0x76, 0xa7, 0xcc, 0x4c, 0x2c, 0xfa, 0x14, 0xde, 0xfd, 0xef, 0xab, 0xf8,
	0x34, 0xde, 0x5f, 0xc1, 0x33, 0x73, 0x36, 0xdb, 0x64, 0xed, 0x25, 0xfd, 0xd3, 0x66, 0x72, 0xce,
	0x77, 0xc9, 0x9c, 0xf3, 0xed, 0xb2, 0xdb, 0x91, 0xd2, 0x42, 0x42, 0x0b, 0xd2, 0x6e, 0x9c, 0x42,
	0x6b, 0x57, 0x8a, 0x10, 0xa2, 0xbe, 0x04, 0xd5, 0x52, 0xda, 0x7f, 0x08, 0xfa, 0xa9, 0x04, 0x68,
	0x8b, 0x08, 0xd6, 0x37, 0x9e, 0xe8, 0x40, 0xc7, 0x4a, 0xc7, 0xa1, 0xf2, 0x57, 0x22, 0x17, 0x1b,
	0xb5, 0xe0, 0x97, 0x09, 0xed, 0x12, 0xda, 0x3d, 0x12, 0xd2, 0x38, 0x93, 0x49, 0xbd, 0x0a, 0x7a,
	0x7d, 0x50, 0xc4, 0xd0, 0x98, 0x1f, 0xd5, 0x07, 0x29, 0x85, 0xcc, 0x4a, 0xcd, 0xd1, 0x52, 0x02,
	0x4a, 0x05, 0x5d, 0xc8, 0x8a, 0x17, 0x8a, 0x45, 0x1d, 0xc4, 0xe9, 0x96, 0x90, 0x09, 0x2a, 0x8a,
	0x94, 0x9a, 0x96, 0xbe, 0x54, 0x18, 0x5b, 0x0f, 0x64, 0x80, 0x55, 0x90, 0x8a, 0xdf, 0x64, 0xa7,
	0xb7, 0xa4, 0x48, 0xfc, 0x3d, 0x80, 0x1d, 0x67, 0xe2, 0xfc, 0xc4, 0xa5, 0xda, 0xb5, 0xa6, 0x9b,
	0xfd, 0x82, 0xcc, 0x54, 0x9c, 0x6a, 0xe8, 0x82, 0x7c, 0x66, 0x4e, 0xde, 0x94, 0xe9, 0xde, 0xc0,
	0x66, 0x7e, 0x91, 0x4d, 0xe7, 0x48, 0x3f, 0xed, 0xf7, 0x7a, 0xce, 0xf7, 0x2a, 0xe2, 0xa7, 0xbc,
	0xfa, 0xa0, 0xa5, 0x8d, 0x5f, 0xe6, 0x02, 0xaf, 0x21, 0x90, 0xce, 0xa9, 0x31, 0x05, 0x36, 0xb1,
	0x39, 0x17, 0x30, 0x48, 0x12, 0xf8, 0x31, 0x24, 0x60, 0x5a, 0xac, 0xc0, 0x0d, 0x56, 0xd5, 0x82,
	0xfc, 0x97, 0x8e, 0xa7, 0xaf, 0x68, 0x61, 0xdd, 0x2f, 0xb2, 0x7a, 0x86, 0x22, 0xea, 0x9f, 0x44,
	0xcd, 0xa8, 0x3c, 0x44, 0x6c, 0x7d, 0x4f, 0x8e, 0x45, 0x6c, 0x5d, 0x13, 0xf1, 0xbe, 0xe7, 0x5f,
	0x39, 0x71, 0xee, 0xf8, 0x05, 0x3b, 0xd7, 0x09, 0x54, 0x1c, 0xfa, 0xe1, 0x36, 0x0e, 0x22, 0xc4,
	0x39, 0xd8, 0x85, 0xc0, 0xfe, 0xa4, 0x83, 0x33, 0x71, 0xca, 0x56, 0xad, 0x51, 0x50, 0x53, 0x5a,
	0xc6, 0x69, 0x97, 0xc4, 0x1a, 0x16, 0x7f, 0x6f, 0x04, 0xde, 0x26, 0x34, 0x7f, 0xc4, 0x16, 0x8f,
	0x62, 0x27, 0x57, 0xbf, 0xc9, 0xd5, 0xc2, 0xe1, 0x3c, 0xd6, 0xe9, 0x0a, 0x9b, 0xd9, 0xf6, 0x35,
	0xee, 0xae, 0x9f, 0xe2, 0xf2, 0xfa, 0x71, 0xa4, 0x9c, 0xca, 0xb1, 0xe6, 0xea, 0xdb, 0x83, 0x6d,
	0x5f, 0x8d, 0x14, 0x6f, 0xb1, 0xb9, 0x02, 0x05, 0x39, 0xf8, 0x43, 0x0e, 0x66, 0x87, 0x9b, 0xad,
	0xe6, 0x7d, 0x36, 0xa3, 0xfa, 0x49, 0x12, 0xc8, 0xf8, 0x0d, 0xd8, 0x01, 0x29, 0xa7, 0x7a, 0xe0,
	0xf5, 0x77, 0x84, 0xe8, 0x41, 0x90, 0x92, 0xe8, 0x74, 0x8e, 0x31, 0xf3, 0x53, 0xfc, 0x2a, 0x9b,
	0x2b, 0xb0, 0x90, 0xec, 0x5f, 0x92, 0xe5, 0xa3, 0xed, 0x46, 0x78, 0xe9, 0x5b, 0x99, 0x4d, 0x79,
	0xa0, 0x76, 0x45, 0xaa, 0x80, 0x5f, 0x61, 0x65, 0x9b, 0xbb, 0x2c, 0x13, 0xf9, 0xef, 0xcd, 0x52,
	0x4d, 0x99, 0x7c, 0x60, 0xfe, 0x7a, 0xd4, 0xc8, 0x37, 0xd9, 0xac, 0x49, 0x9c, 0x3f, 0x14, 0x39,
	0xdc, 0xf7, 0x12, 0x82, 0xdd, 0x02, 0xb8, 0x18, 0xcc, 0x35, 0x3c, 0xaf, 0xee, 0x9f, 0xbd, 0x99,
	0x64, 0xf4, 0x0b, 0xcc, 0x50, 0x35, 0x4b, 0x3a, 0xae, 0xb8, 0x61, 0x5c, 0xf8, 0x8f, 0x91, 0x9e,
	0x03, 0x6b, 0xf4, 0xdf, 0x1b, 0xb4, 0xf3, 0xc7, 0xac, 0x24, 0xc5, 0x1e, 0xee, 0xaf, 0x41, 0xdd,
	0x72, 0x4f, 0xf0, 0x68, 0x72, 0x07, 0x57, 0xe1, 0x7a, 0x62, 0xcf, 0x33, 0x2c, 0x8d, 0xaf, 0x25,
	0x56, 0xc2, 0x03, 0x3f, 0xcb, 0x2a, 0x78, 0xc4, 0x51, 0x3a, 0x6f, 0xdb, 0x78, 0x3b, 0x65, 0xaf,
	0x8c, 0xc7, 0xd5, 0x88, 0xdf, 0x61, 0x35, 0x2d, 0x74, 0xd0, 0xf3, 0xed, 0x80, 0x9c, 0x77, 0xed,
	0x03, 0xc7, 0x16, 0x41, 0x18, 0x27, 0x41, 0x8f, 0xc6, 0xc6, 0x2c, 0xc0, 0x7e, 0xc6, 0x2b, 0x9f,
	0xb4, 0x61, 0x7b, 0xdf, 0x3e, 0x3e, 0x6d, 0xb6, 0x13, 0xd7, 0x73, 0x7a, 0x74, 0xb7, 0x9c, 0x0f,
	0x63, 0x60, 0x6b, 0x43, 0x2b, 0xc7, 0x97, 0x59, 0x3d, 0x8a, 0x25, 0x84, 0x3a, 0x33, 0xfd, 0x71,
	0x0c, 0xd3, 0x35, 0x42, 0xe4, 0xae, 0xed, 0xb3, 0xe7, 0xd3, 0x38, 0xae, 0x4d, 0x27, 0xc6, 0xbf,
	0x79, 0x44, 0x40, 0x9d, 0xcf, 0x63, 0x10, 0xcd, 0x1f, 0x9a, 0xdb, 0xbb, 0x2f, 0x59, 0x33, 0x16,
	0x85, 0x41, 0xef, 0xbf, 0xc1, 0x9e, 0x2f, 0x77, 0x85, 0x8a, 0x76, 0x06, 0xf5, 0xe8, 0xc4, 0x2f,
	0xb9, 0x4e, 0xc5, 0xbe, 0x46, 0xae, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xed, 0x9c, 0xb7, 0x99,
	0x25, 0x07, 0x00, 0x00,
}
