// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetSearchItemLacks_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetSearchItemLacks_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetSearchItemLacks_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetSearchItemLacks_Ad

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
	TableId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	TableIdNull               bool                        `protobuf:"varint,1001,opt,name=table_id_null,json=tableIdNull" json:"table_id_null,omitempty"`
	CharacteristicId          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=characteristic_id,json=characteristicId" json:"characteristic_id,omitempty"`
	CharacteristicIdNull      bool                        `protobuf:"varint,1002,opt,name=characteristic_id_null,json=characteristicIdNull" json:"characteristic_id_null,omitempty"`
	SearchValue               *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=search_value,json=searchValue" json:"search_value,omitempty"`
	SearchValueNull           bool                        `protobuf:"varint,1003,opt,name=search_value_null,json=searchValueNull" json:"search_value_null,omitempty"`
	MinimalRequestCounter     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=minimal_request_counter,json=minimalRequestCounter" json:"minimal_request_counter,omitempty"`
	MinimalRequestCounterNull bool                        `protobuf:"varint,1004,opt,name=minimal_request_counter_null,json=minimalRequestCounterNull" json:"minimal_request_counter_null,omitempty"`
	Rowcount                  *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=rowcount" json:"rowcount,omitempty"`
	RowcountNull              bool                        `protobuf:"varint,1005,opt,name=rowcount_null,json=rowcountNull" json:"rowcount_null,omitempty"`
	MaxRequestCounter         *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=max_request_counter,json=maxRequestCounter" json:"max_request_counter,omitempty"`
	MaxRequestCounterNull     bool                        `protobuf:"varint,1006,opt,name=max_request_counter_null,json=maxRequestCounterNull" json:"max_request_counter_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTableId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableId
	}
	return nil
}

func (m *Parameters) GetCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CharacteristicId
	}
	return nil
}

func (m *Parameters) GetSearchValue() *dstore_values.StringValue {
	if m != nil {
		return m.SearchValue
	}
	return nil
}

func (m *Parameters) GetMinimalRequestCounter() *dstore_values.IntegerValue {
	if m != nil {
		return m.MinimalRequestCounter
	}
	return nil
}

func (m *Parameters) GetRowcount() *dstore_values.IntegerValue {
	if m != nil {
		return m.Rowcount
	}
	return nil
}

func (m *Parameters) GetMaxRequestCounter() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxRequestCounter
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
	RequestCounter   *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=request_counter,json=requestCounter" json:"request_counter,omitempty"`
	SearchValue      *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=search_value,json=searchValue" json:"search_value,omitempty"`
	CharacteristicId *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=characteristic_id,json=characteristicId" json:"characteristic_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetRequestCounter() *dstore_values.IntegerValue {
	if m != nil {
		return m.RequestCounter
	}
	return nil
}

func (m *Response_Row) GetSearchValue() *dstore_values.StringValue {
	if m != nil {
		return m.SearchValue
	}
	return nil
}

func (m *Response_Row) GetCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CharacteristicId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetSearchItemLacks_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetSearchItemLacks_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetSearchItemLacks_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetSearchItemLacks_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xeb, 0x6e, 0xd3, 0x4c,
	0x10, 0x55, 0xbf, 0x34, 0x4d, 0x34, 0x69, 0xbf, 0x34, 0x2e, 0x2d, 0x6e, 0x82, 0x10, 0x6a, 0x11,
	0x42, 0x42, 0x72, 0x10, 0xb7, 0x22, 0x24, 0x10, 0x17, 0x15, 0x88, 0xa0, 0x11, 0x72, 0x25, 0xa4,
	0xf2, 0xc7, 0xda, 0xda, 0x4b, 0xba, 0x6a, 0xec, 0x2d, 0xbb, 0x1b, 0xca, 0x3b, 0xf0, 0x87, 0xcb,
	0xa3, 0xf1, 0x14, 0xdc, 0x5f, 0x81, 0xf1, 0x8e, 0xdd, 0x26, 0x4e, 0xaa, 0x50, 0xfe, 0x24, 0x9e,
	0x9d, 0x73, 0xe6, 0xcc, 0x5e, 0xce, 0xc0, 0x46, 0xa4, 0x8d, 0x54, 0xbc, 0xcd, 0x93, 0x9e, 0x48,
	0x78, 0xfb, 0x40, 0xc9, 0x90, 0x47, 0x03, 0xc5, 0x75, 0x3b, 0x16, 0xc1, 0x13, 0x6e, 0xb6, 0x39,
	0x53, 0xe1, 0x5e, 0xc7, 0xf0, 0xf8, 0x39, 0x0b, 0xf7, 0x75, 0xf0, 0x20, 0xf2, 0x10, 0x63, 0xa4,
	0x73, 0x89, 0x88, 0x1e, 0x11, 0xbd, 0x93, 0xd0, 0xcd, 0xa5, 0x4c, 0xe0, 0x2d, 0xeb, 0x0f, 0xb8,
	0x26, 0x72, 0x73, 0x75, 0x54, 0x95, 0x2b, 0x25, 0x55, 0x96, 0x6a, 0x8d, 0xa6, 0x62, 0xae, 0x35,
	0xeb, 0xf1, 0x2c, 0xb9, 0x5e, 0x4c, 0x1a, 0x26, 0x92, 0xd7, 0x52, 0xc5, 0xcc, 0x08, 0x99, 0x10,
	0x68, 0xed, 0x4b, 0x19, 0xe0, 0x05, 0x53, 0x0c, 0xb3, 0x5c, 0x69, 0xe7, 0x16, 0x54, 0x0d, 0xdb,
	0xed, 0xf3, 0x40, 0x44, 0xee, 0xcc, 0x85, 0x99, 0xcb, 0xb5, 0x6b, 0x2d, 0x2f, 0xeb, 0x3d, 0xeb,
	0x49, 0x24, 0x86, 0xf7, 0xb8, 0x7a, 0x99, 0x46, 0x7e, 0xc5, 0x82, 0x3b, 0x91, 0xb3, 0x0e, 0x0b,
	0x39, 0x2f, 0x48, 0x06, 0xfd, 0xbe, 0xfb, 0xb5, 0x82, 0xec, 0xaa, 0x5f, 0xcb, 0x00, 0x5d, 0x5c,
	0x73, 0x9e, 0x42, 0x23, 0xdc, 0x43, 0xad, 0x10, 0xa5, 0x84, 0x36, 0x22, 0x4c, 0x55, 0xfe, 0x9b,
	0xae, 0xb2, 0x38, 0xca, 0x42, 0xb9, 0x9b, 0xb0, 0x32, 0x56, 0x89, 0x74, 0xbf, 0x91, 0xee, 0x99,
	0x22, 0xc5, 0x36, 0x70, 0x17, 0xe6, 0xb5, 0x3d, 0xf4, 0xc0, 0xca, 0xb8, 0x25, 0xab, 0xdd, 0x2c,
	0x68, 0x6b, 0xa3, 0x44, 0xd2, 0x23, 0xe9, 0x1a, 0xe1, 0x6d, 0xe0, 0x5c, 0x81, 0xc6, 0x30, 0x9d,
	0x04, 0xbf, 0x93, 0x60, 0x7d, 0x08, 0x68, 0xb5, 0xb6, 0xe1, 0x6c, 0x2c, 0x12, 0x11, 0xb3, 0x7e,
	0xa0, 0xf8, 0x1b, 0xac, 0x6b, 0x82, 0x50, 0x0e, 0x70, 0x5b, 0xca, 0x9d, 0x9d, 0xbe, 0xe5, 0xe5,
	0x8c, 0xeb, 0x13, 0xf5, 0x11, 0x31, 0x9d, 0xfb, 0x70, 0xee, 0x84, 0xa2, 0xd4, 0xcc, 0x0f, 0x6a,
	0x66, 0x75, 0x22, 0xdb, 0xb6, 0xb5, 0x01, 0x55, 0x25, 0x0f, 0x2d, 0xc9, 0x2d, 0x4f, 0xef, 0xe3,
	0x08, 0xec, 0x5c, 0x84, 0x85, 0xfc, 0x9b, 0xb4, 0x7e, 0x92, 0xd6, 0x7c, 0xbe, 0x6a, 0xcb, 0x3f,
	0x83, 0xa5, 0x98, 0xbd, 0x1b, 0xdb, 0xf1, 0xdc, 0x74, 0xa5, 0x06, 0xf2, 0x0a, 0xbb, 0xbd, 0x0d,
	0xee, 0x84, 0x62, 0xa4, 0xfe, 0x8b, 0xd4, 0x97, 0xc7, 0x58, 0x69, 0x1b, 0x6b, 0xef, 0x67, 0xa1,
	0xea, 0x73, 0x7d, 0x20, 0x13, 0xcd, 0x9d, 0xab, 0x50, 0xb6, 0x9e, 0xc9, 0x1e, 0xf4, 0xd1, 0x75,
	0x67, 0x66, 0x24, 0x3f, 0x6d, 0xa6, 0xbf, 0x3e, 0x01, 0x9d, 0x1d, 0x58, 0x4c, 0xdd, 0x12, 0x0c,
	0xd9, 0x05, 0xdf, 0x69, 0x09, 0xc9, 0x5e, 0x81, 0x5c, 0x34, 0xd5, 0x16, 0xc6, 0x9d, 0xe3, 0xd8,
	0xaf, 0xc7, 0xa3, 0x0b, 0xb8, 0xa7, 0x4a, 0xe6, 0x52, 0x7c, 0x7d, 0x69, 0xc5, 0xf3, 0x63, 0x15,
	0xc9, 0xc3, 0x5b, 0xf4, 0xef, 0xe7, 0x70, 0xe7, 0x31, 0x94, 0xf0, 0xa8, 0xf1, 0xf1, 0xa4, 0xac,
	0x1b, 0xde, 0xdf, 0x4d, 0x14, 0x2f, 0x3f, 0x05, 0xcf, 0x97, 0x87, 0x7e, 0x5a, 0xa0, 0xf9, 0x7b,
	0x06, 0x4a, 0x18, 0x38, 0x2b, 0x30, 0x87, 0x61, 0x6a, 0xc1, 0x0f, 0x5d, 0x3c, 0x98, 0xb2, 0x5f,
	0xc6, 0x10, 0xbd, 0xb5, 0x09, 0xf5, 0xe2, 0xf5, 0x7d, 0xec, 0x4e, 0xbf, 0xbf, 0xff, 0xd5, 0xe8,
	0xe5, 0xdd, 0x2b, 0x78, 0xed, 0x53, 0xf7, 0x74, 0x66, 0xeb, 0x4c, 0x1a, 0x16, 0x9f, 0xbb, 0xff,
	0x30, 0x2d, 0x1e, 0xee, 0x40, 0x4b, 0xc8, 0xc2, 0x81, 0x1d, 0xcf, 0xee, 0x57, 0x77, 0x7a, 0x52,
	0x47, 0xfb, 0x79, 0x3e, 0x3a, 0xcd, 0x78, 0xdf, 0x9d, 0xb3, 0x53, 0xf4, 0xfa, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x72, 0x1d, 0x99, 0xe7, 0x1a, 0x06, 0x00, 0x00,
}