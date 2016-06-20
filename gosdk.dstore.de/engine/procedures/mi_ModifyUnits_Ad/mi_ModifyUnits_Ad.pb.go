// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyUnits_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyUnits_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyUnits_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyUnits_Ad

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
	UnitId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=unit_id,json=unitId" json:"unit_id,omitempty"`
	UnitIdNull          bool                        `protobuf:"varint,1001,opt,name=unit_id_null,json=unitIdNull" json:"unit_id_null,omitempty"`
	UnitSymbol          *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=unit_symbol,json=unitSymbol" json:"unit_symbol,omitempty"`
	UnitSymbolNull      bool                        `protobuf:"varint,1002,opt,name=unit_symbol_null,json=unitSymbolNull" json:"unit_symbol_null,omitempty"`
	UnitDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unit_description,json=unitDescription" json:"unit_description,omitempty"`
	UnitDescriptionNull bool                        `protobuf:"varint,1003,opt,name=unit_description_null,json=unitDescriptionNull" json:"unit_description_null,omitempty"`
	TriangleConvert     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=triangle_convert,json=triangleConvert" json:"triangle_convert,omitempty"`
	TriangleConvertNull bool                        `protobuf:"varint,1004,opt,name=triangle_convert_null,json=triangleConvertNull" json:"triangle_convert_null,omitempty"`
	DeleteUnit          *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_unit,json=deleteUnit" json:"delete_unit,omitempty"`
	DeleteUnitNull      bool                        `protobuf:"varint,1005,opt,name=delete_unit_null,json=deleteUnitNull" json:"delete_unit_null,omitempty"`
	Active              *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=active" json:"active,omitempty"`
	ActiveNull          bool                        `protobuf:"varint,1006,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitId
	}
	return nil
}

func (m *Parameters) GetUnitSymbol() *dstore_values.StringValue {
	if m != nil {
		return m.UnitSymbol
	}
	return nil
}

func (m *Parameters) GetUnitDescription() *dstore_values.StringValue {
	if m != nil {
		return m.UnitDescription
	}
	return nil
}

func (m *Parameters) GetTriangleConvert() *dstore_values.BooleanValue {
	if m != nil {
		return m.TriangleConvert
	}
	return nil
}

func (m *Parameters) GetDeleteUnit() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteUnit
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyUnits_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyUnits_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyUnits_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_ModifyUnits_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xd6, 0x56, 0xda, 0x4d, 0xa7, 0x88, 0x55, 0x9e, 0x40, 0xa1, 0x15, 0x30, 0xc6, 0x0d, 0x5c,
	0x90, 0x4e, 0x94, 0x0b, 0x24, 0x90, 0x10, 0x3f, 0x43, 0xea, 0x45, 0x27, 0x14, 0x04, 0x12, 0xdc,
	0x44, 0x69, 0x7d, 0x16, 0x59, 0xa4, 0x76, 0x65, 0xa7, 0x9d, 0xf6, 0x02, 0x5c, 0xf3, 0x2a, 0x3c,
	0x16, 0xbf, 0xcf, 0x80, 0xed, 0xe3, 0x92, 0x36, 0x43, 0xea, 0x6e, 0xda, 0x1c, 0x9f, 0xef, 0xc7,
	0x71, 0xbe, 0x63, 0x38, 0xe2, 0xa6, 0x54, 0x1a, 0xfb, 0x28, 0x73, 0x21, 0xb1, 0x3f, 0xd3, 0x6a,
	0x82, 0x7c, 0xae, 0xd1, 0xf4, 0xa7, 0x22, 0x1d, 0x29, 0x2e, 0x4e, 0xcf, 0xdf, 0x4b, 0x51, 0x9a,
	0xf4, 0x05, 0x8f, 0x6d, 0xb3, 0x54, 0xec, 0x0e, 0x31, 0x62, 0x62, 0xc4, 0x17, 0x60, 0xdd, 0xfd,
	0x20, 0xb9, 0xc8, 0x8a, 0x39, 0x1a, 0x62, 0x75, 0x6f, 0xae, 0xfb, 0xa0, 0xd6, 0x4a, 0x87, 0x56,
	0x6f, 0xbd, 0x35, 0x45, 0x63, 0xb2, 0x1c, 0x43, 0xf3, 0x5e, 0xbd, 0x59, 0x66, 0x42, 0x9e, 0x2a,
	0x3d, 0xcd, 0x4a, 0xa1, 0x24, 0x81, 0x0e, 0xbf, 0x34, 0x01, 0xde, 0x66, 0x3a, 0xb3, 0x5d, 0xd4,
	0x86, 0x3d, 0x86, 0x9d, 0xb9, 0xdd, 0x4c, 0x2a, 0x78, 0xb4, 0x75, 0xb0, 0x75, 0xbf, 0xfd, 0xa8,
	0x17, 0x87, 0x3d, 0x87, 0x2d, 0x09, 0x59, 0x62, 0x8e, 0xfa, 0x83, 0xab, 0x92, 0x96, 0xc3, 0x0e,
	0x39, 0xbb, 0x0b, 0x57, 0x03, 0x2b, 0x95, 0xf3, 0xa2, 0x88, 0xbe, 0xef, 0x58, 0xee, 0x6e, 0x02,
	0xd4, 0x3e, 0xb1, 0x4b, 0xec, 0x29, 0xb4, 0x3d, 0xc4, 0x9c, 0x4f, 0xc7, 0xaa, 0x88, 0xb6, 0xbd,
	0x78, 0xb7, 0x26, 0x6e, 0x4a, 0x2d, 0x64, 0x4e, 0xda, 0x9e, 0xfc, 0xce, 0xa3, 0xd9, 0x03, 0xe8,
	0xac, 0x90, 0xc9, 0xe3, 0x07, 0x79, 0x5c, 0xab, 0x60, 0xde, 0xe7, 0x38, 0x40, 0x39, 0x9a, 0x89,
	0x16, 0x33, 0xf7, 0xa6, 0x51, 0x63, 0xa3, 0xd9, 0x9e, 0xe3, 0xbc, 0xae, 0x28, 0x6c, 0x00, 0xd7,
	0xeb, 0x32, 0x64, 0xfb, 0x93, 0x6c, 0xf7, 0x6b, 0x04, 0xef, 0xfd, 0x06, 0x3a, 0x56, 0x33, 0x93,
	0x79, 0x81, 0xe9, 0x44, 0xc9, 0x05, 0xea, 0x32, 0xba, 0xf2, 0xdf, 0x53, 0x1c, 0x2b, 0x55, 0x60,
	0x26, 0x83, 0xf9, 0x92, 0xf4, 0x8a, 0x38, 0xce, 0xbc, 0xae, 0x43, 0xe6, 0xbf, 0x82, 0x79, 0x8d,
	0xe0, 0xcd, 0x9f, 0x41, 0x9b, 0x63, 0x61, 0xbf, 0x62, 0xea, 0xb6, 0x16, 0x35, 0x37, 0xfb, 0x02,
	0xe1, 0x5d, 0xf8, 0xdc, 0x09, 0xaf, 0xb0, 0xc9, 0xed, 0x77, 0x38, 0xe1, 0x0a, 0xe6, 0x8d, 0x06,
	0xd0, 0xca, 0x26, 0xa5, 0x58, 0x60, 0xd4, 0xba, 0x44, 0x42, 0x08, 0xca, 0x0e, 0xa0, 0x4d, 0x4f,
	0x24, 0xfd, 0x27, 0x04, 0x84, 0xd6, 0x9c, 0xec, 0xe1, 0xb7, 0x6d, 0xd8, 0x4d, 0xd0, 0xcc, 0x94,
	0x34, 0xc8, 0x8e, 0xa0, 0xe9, 0x63, 0x1e, 0x42, 0xf8, 0xef, 0xd3, 0x85, 0xc1, 0xa1, 0x11, 0x38,
	0x76, 0xbf, 0x09, 0x01, 0xd9, 0x47, 0xe8, 0xb8, 0x80, 0xa7, 0x2b, 0x09, 0xb7, 0x21, 0x6b, 0x58,
	0x72, 0x5c, 0x23, 0xd7, 0xe7, 0x60, 0x64, 0xeb, 0x61, 0x55, 0x27, 0x7b, 0xd3, 0xf5, 0x05, 0xf6,
	0x04, 0x76, 0xc2, 0x60, 0xd9, 0x24, 0x39, 0xc5, 0xdb, 0x17, 0x14, 0x69, 0xec, 0x46, 0xf4, 0x9f,
	0x2c, 0xe1, 0xec, 0x39, 0x34, 0xb4, 0x3a, 0xb3, 0x19, 0x70, 0xac, 0x87, 0xf1, 0x86, 0xe9, 0x8f,
	0x97, 0xaf, 0x1f, 0x27, 0xea, 0x2c, 0x71, 0xcc, 0xee, 0x2d, 0x68, 0xd8, 0x67, 0x76, 0x03, 0x5a,
	0xb6, 0x72, 0x43, 0xf9, 0xf5, 0xc4, 0x1e, 0x48, 0x33, 0x69, 0xda, 0x72, 0xc8, 0x5f, 0x26, 0xd0,
	0x13, 0xaa, 0x26, 0x5b, 0x5d, 0x43, 0x9f, 0x06, 0xb9, 0x32, 0xfc, 0xf3, 0xb2, 0xcf, 0x2f, 0x75,
	0x53, 0x8d, 0x5b, 0xfe, 0x5e, 0x18, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x44, 0x33, 0x63, 0xd2,
	0xde, 0x04, 0x00, 0x00,
}