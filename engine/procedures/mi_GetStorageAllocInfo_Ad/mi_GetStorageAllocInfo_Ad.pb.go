// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetStorageAllocInfo_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetStorageAllocInfo_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetStorageAllocInfo_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetStorageAllocInfo_Ad

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
	GetInfoForADatabase        *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=get_info_for_a_database,json=getInfoForADatabase" json:"get_info_for_a_database,omitempty"`
	GetInfoForADatabaseNull    bool                        `protobuf:"varint,1001,opt,name=get_info_for_a_database_null,json=getInfoForADatabaseNull" json:"get_info_for_a_database_null,omitempty"`
	GetStorageAllocInfoFor     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=get_storage_alloc_info_for,json=getStorageAllocInfoFor" json:"get_storage_alloc_info_for,omitempty"`
	GetStorageAllocInfoForNull bool                        `protobuf:"varint,1002,opt,name=get_storage_alloc_info_for_null,json=getStorageAllocInfoForNull" json:"get_storage_alloc_info_for_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetGetInfoForADatabase() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetInfoForADatabase
	}
	return nil
}

func (m *Parameters) GetGetStorageAllocInfoFor() *dstore_values.StringValue {
	if m != nil {
		return m.GetStorageAllocInfoFor
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
	RowId              int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TableName          *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	IndexSizeMb        *dstore_values.DecimalValue `protobuf:"bytes,10002,opt,name=index_size_mb,json=indexSizeMb" json:"index_size_mb,omitempty"`
	NumberOfIndexes    *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=number_of_indexes,json=numberOfIndexes" json:"number_of_indexes,omitempty"`
	NumberOfRows       *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=number_of_rows,json=numberOfRows" json:"number_of_rows,omitempty"`
	TableSizeMb        *dstore_values.DecimalValue `protobuf:"bytes,10005,opt,name=table_size_mb,json=tableSizeMb" json:"table_size_mb,omitempty"`
	DataSizeMb         *dstore_values.DecimalValue `protobuf:"bytes,10006,opt,name=data_size_mb,json=dataSizeMb" json:"data_size_mb,omitempty"`
	Mballocated        *dstore_values.DecimalValue `protobuf:"bytes,20001,opt,name=mballocated" json:"mballocated,omitempty"`
	SegmentName        *dstore_values.StringValue  `protobuf:"bytes,20002,opt,name=segment_name,json=segmentName" json:"segment_name,omitempty"`
	Mbused             *dstore_values.DecimalValue `protobuf:"bytes,20003,opt,name=mbused" json:"mbused,omitempty"`
	FreeSpaceInPercent *dstore_values.DecimalValue `protobuf:"bytes,20004,opt,name=free_space_in_percent,json=freeSpaceInPercent" json:"free_space_in_percent,omitempty"`
	DBName             *dstore_values.StringValue  `protobuf:"bytes,20005,opt,name=d_b_name,json=dBName" json:"d_b_name,omitempty"`
	Mbfree             *dstore_values.DecimalValue `protobuf:"bytes,20006,opt,name=mbfree" json:"mbfree,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTableName() *dstore_values.StringValue {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *Response_Row) GetIndexSizeMb() *dstore_values.DecimalValue {
	if m != nil {
		return m.IndexSizeMb
	}
	return nil
}

func (m *Response_Row) GetNumberOfIndexes() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfIndexes
	}
	return nil
}

func (m *Response_Row) GetNumberOfRows() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfRows
	}
	return nil
}

func (m *Response_Row) GetTableSizeMb() *dstore_values.DecimalValue {
	if m != nil {
		return m.TableSizeMb
	}
	return nil
}

func (m *Response_Row) GetDataSizeMb() *dstore_values.DecimalValue {
	if m != nil {
		return m.DataSizeMb
	}
	return nil
}

func (m *Response_Row) GetMballocated() *dstore_values.DecimalValue {
	if m != nil {
		return m.Mballocated
	}
	return nil
}

func (m *Response_Row) GetSegmentName() *dstore_values.StringValue {
	if m != nil {
		return m.SegmentName
	}
	return nil
}

func (m *Response_Row) GetMbused() *dstore_values.DecimalValue {
	if m != nil {
		return m.Mbused
	}
	return nil
}

func (m *Response_Row) GetFreeSpaceInPercent() *dstore_values.DecimalValue {
	if m != nil {
		return m.FreeSpaceInPercent
	}
	return nil
}

func (m *Response_Row) GetDBName() *dstore_values.StringValue {
	if m != nil {
		return m.DBName
	}
	return nil
}

func (m *Response_Row) GetMbfree() *dstore_values.DecimalValue {
	if m != nil {
		return m.Mbfree
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetStorageAllocInfo_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetStorageAllocInfo_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetStorageAllocInfo_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetStorageAllocInfo_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xd3, 0x5a,
	0x10, 0x56, 0x9a, 0xdb, 0x34, 0x77, 0xd2, 0x7b, 0x7b, 0xef, 0xa9, 0x68, 0x4d, 0x8a, 0x00, 0x95,
	0x05, 0xac, 0x5c, 0x04, 0x2a, 0xaa, 0x40, 0x2a, 0xb4, 0x2a, 0x94, 0x2c, 0x1a, 0x8a, 0x2b, 0x55,
	0xa2, 0x9b, 0xa3, 0xe3, 0x78, 0x62, 0x59, 0xd8, 0x3e, 0xd1, 0x39, 0x0e, 0x45, 0x3c, 0x05, 0xff,
	0xab, 0x6e, 0xf8, 0x29, 0x6f, 0xc1, 0x86, 0xb7, 0x81, 0xa7, 0x60, 0x8e, 0x8f, 0xd3, 0xb4, 0x69,
	0xaa, 0x9a, 0x4d, 0xa2, 0xc9, 0xcc, 0xf7, 0xcd, 0x7c, 0xf3, 0x8d, 0x63, 0x58, 0x09, 0x74, 0x26,
	0x15, 0x2e, 0x61, 0x1a, 0x46, 0x29, 0x2e, 0xf5, 0x94, 0xec, 0x60, 0xd0, 0x57, 0xa8, 0x97, 0x92,
	0x88, 0x6f, 0x62, 0xb6, 0x43, 0x59, 0x11, 0xe2, 0x5a, 0x1c, 0xcb, 0x4e, 0x2b, 0xed, 0x4a, 0xbe,
	0x16, 0xb8, 0x54, 0x94, 0x49, 0x76, 0xdd, 0x22, 0x5d, 0x8b, 0x74, 0xcf, 0x2c, 0x6f, 0xce, 0x16,
	0x2d, 0x5e, 0x88, 0xb8, 0x8f, 0xda, 0xa2, 0x9b, 0x17, 0x4f, 0xf6, 0x45, 0xa5, 0xa4, 0x2a, 0x52,
	0x0b, 0x27, 0x53, 0x09, 0x6a, 0x4d, 0x94, 0x45, 0xf2, 0xda, 0x68, 0x32, 0x13, 0x11, 0xf5, 0x51,
	0x89, 0xc8, 0x22, 0x99, 0xda, 0xa2, 0xc5, 0xef, 0x13, 0x00, 0xdb, 0x42, 0x09, 0xca, 0xa2, 0xd2,
	0xec, 0x29, 0xcc, 0x87, 0x98, 0x71, 0x53, 0xc7, 0xa9, 0x94, 0x0b, 0x1e, 0x88, 0x4c, 0xf8, 0x42,
	0xa3, 0x53, 0xb9, 0x5a, 0xb9, 0xd1, 0xb8, 0xb5, 0xe0, 0x16, 0x5a, 0x8a, 0x11, 0x7d, 0x29, 0x63,
	0x14, 0xe9, 0xae, 0x89, 0xbc, 0x59, 0xc2, 0x1a, 0x29, 0x8f, 0xa4, 0x5a, 0xdb, 0x28, 0x70, 0x6c,
	0x15, 0x2e, 0x9d, 0x41, 0xc9, 0xd3, 0x7e, 0x1c, 0x3b, 0x3f, 0xa7, 0x88, 0xb8, 0xee, 0xcd, 0x8f,
	0xc1, 0xb6, 0x29, 0xcf, 0x76, 0xa1, 0x69, 0xf0, 0xda, 0xae, 0x8b, 0x0b, 0xb3, 0xaf, 0x23, 0x36,
	0x67, 0x22, 0x9f, 0xaa, 0x39, 0x32, 0x95, 0xce, 0x54, 0x94, 0x86, 0x76, 0xa8, 0xb9, 0xf0, 0xf4,
	0xae, 0xa9, 0x09, 0xdb, 0x80, 0x2b, 0x67, 0xf3, 0xda, 0xd1, 0x7e, 0xd9, 0xd1, 0x9a, 0xe3, 0x19,
	0xcc, 0x74, 0x8b, 0x3f, 0xea, 0x50, 0xf7, 0x50, 0xf7, 0x64, 0x4a, 0x52, 0x6f, 0xc2, 0x64, 0xee,
	0x4e, 0xb1, 0xab, 0xa3, 0xa9, 0x0a, 0xdf, 0xad, 0x73, 0x0f, 0xcd, 0xa7, 0x67, 0x0b, 0xd9, 0x33,
	0xf8, 0xcf, 0xf8, 0xc2, 0x8f, 0x19, 0x43, 0x92, 0xaa, 0x04, 0x76, 0x47, 0xc0, 0xa3, 0xf6, 0x6d,
	0x51, 0xdc, 0x1a, 0xc6, 0xde, 0x4c, 0x72, 0xf2, 0x07, 0xb6, 0x02, 0x53, 0xc5, 0x3d, 0x38, 0xd5,
	0x9c, 0xf1, 0xf2, 0x29, 0x46, 0x7b, 0x2d, 0x5b, 0xf6, 0xdb, 0x1b, 0x94, 0xb3, 0x4d, 0xa8, 0x2a,
	0xb9, 0xef, 0xfc, 0x95, 0xa3, 0x96, 0xdd, 0x92, 0xc7, 0xeb, 0x0e, 0xd6, 0xe0, 0x7a, 0x72, 0xdf,
	0x33, 0x0c, 0xcd, 0xc3, 0x1a, 0x54, 0x29, 0x60, 0x73, 0x50, 0xa3, 0x90, 0x47, 0x81, 0xf3, 0xba,
	0x4d, 0x9b, 0x99, 0xf4, 0x26, 0x29, 0x6c, 0x05, 0xec, 0x2e, 0x00, 0x19, 0x1d, 0xd3, 0x21, 0xd0,
	0xfd, 0x39, 0x6f, 0xda, 0xe7, 0x7a, 0xf9, 0x77, 0x5e, 0xde, 0xa6, 0x6a, 0xf6, 0x00, 0xfe, 0x89,
	0xd2, 0x00, 0x5f, 0x72, 0x1d, 0xbd, 0x42, 0x9e, 0xf8, 0xce, 0xdb, 0xf6, 0xd8, 0x03, 0x0d, 0xb0,
	0x13, 0x25, 0x22, 0xb6, 0xf8, 0x46, 0x0e, 0xd9, 0x21, 0xc4, 0x96, 0xcf, 0x1e, 0xc3, 0xff, 0x69,
	0x3f, 0xf1, 0x51, 0x71, 0xd9, 0xe5, 0x79, 0x02, 0xb5, 0xf3, 0x6e, 0x3c, 0x4b, 0x94, 0x66, 0x18,
	0xa2, 0xb2, 0x2c, 0x33, 0x16, 0xf6, 0xa4, 0xdb, 0xb2, 0x20, 0xb6, 0x0e, 0xff, 0x0e, 0x99, 0x48,
	0x9a, 0x76, 0xde, 0x97, 0xa0, 0x99, 0x1e, 0xd0, 0xd0, 0x8a, 0xb4, 0xd1, 0x63, 0x77, 0x31, 0xd0,
	0xf3, 0xa1, 0x8c, 0x9e, 0x1c, 0x52, 0xe8, 0x59, 0x85, 0x69, 0xf3, 0x64, 0x1d, 0x11, 0x7c, 0x2c,
	0x41, 0x00, 0x06, 0x51, 0xe0, 0xef, 0x43, 0x23, 0xf1, 0xf3, 0xc7, 0x40, 0x64, 0x18, 0x38, 0x9f,
	0x0e, 0x2a, 0x25, 0x06, 0x38, 0x86, 0x20, 0x82, 0x69, 0x8d, 0x61, 0x82, 0x69, 0x66, 0x0d, 0xfd,
	0x7c, 0x50, 0x39, 0xd7, 0xd1, 0x46, 0x81, 0xc8, 0x3d, 0x5d, 0x86, 0x5a, 0xe2, 0xf7, 0x35, 0x35,
	0xff, 0x52, 0xa6, 0x79, 0x51, 0xcc, 0xb6, 0xe1, 0x42, 0x57, 0x21, 0x6d, 0xae, 0x27, 0x3a, 0x48,
	0x4e, 0xf2, 0x1e, 0xaa, 0x0e, 0x71, 0x3a, 0x5f, 0xcb, 0xb0, 0x30, 0x83, 0xdd, 0x31, 0xd0, 0x56,
	0xba, 0x6d, 0x81, 0xec, 0x0e, 0xd4, 0x03, 0xee, 0x5b, 0x15, 0x87, 0x25, 0x54, 0xd4, 0x82, 0xf5,
	0xa1, 0x00, 0xc3, 0xe7, 0x7c, 0x2b, 0x29, 0xc0, 0x14, 0xaf, 0xef, 0xc1, 0x42, 0x24, 0x47, 0x9e,
	0xb3, 0xe1, 0xeb, 0x65, 0xef, 0x5e, 0x28, 0x75, 0xf0, 0x7c, 0x90, 0x0f, 0xfe, 0xe8, 0x0d, 0xe4,
	0xd7, 0xf2, 0xff, 0xf9, 0xdb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xd5, 0xb1, 0x0e, 0xbe,
	0x06, 0x00, 0x00,
}