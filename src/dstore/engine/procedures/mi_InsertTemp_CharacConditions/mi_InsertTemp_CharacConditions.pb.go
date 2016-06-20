// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_InsertTemp_CharacConditions.proto
// DO NOT EDIT!

/*
Package mi_InsertTemp_CharacConditions is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_InsertTemp_CharacConditions.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_InsertTemp_CharacConditions

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
	CharacteristicIdList          *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=characteristic_id_list,json=characteristicIdList" json:"characteristic_id_list,omitempty"`
	CharacteristicIdListNull      bool                        `protobuf:"varint,1001,opt,name=characteristic_id_list_null,json=characteristicIdListNull" json:"characteristic_id_list_null,omitempty"`
	BasicFieldTypeIdList          *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=basic_field_type_id_list,json=basicFieldTypeIdList" json:"basic_field_type_id_list,omitempty"`
	BasicFieldTypeIdListNull      bool                        `protobuf:"varint,1002,opt,name=basic_field_type_id_list_null,json=basicFieldTypeIdListNull" json:"basic_field_type_id_list_null,omitempty"`
	Operator1List                 *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=operator1_list,json=operator1List" json:"operator1_list,omitempty"`
	Operator1ListNull             bool                        `protobuf:"varint,1003,opt,name=operator1_list_null,json=operator1ListNull" json:"operator1_list_null,omitempty"`
	Condition1List                *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=condition1_list,json=condition1List" json:"condition1_list,omitempty"`
	Condition1ListNull            bool                        `protobuf:"varint,1004,opt,name=condition1_list_null,json=condition1ListNull" json:"condition1_list_null,omitempty"`
	Operator2List                 *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=operator2_list,json=operator2List" json:"operator2_list,omitempty"`
	Operator2ListNull             bool                        `protobuf:"varint,1005,opt,name=operator2_list_null,json=operator2ListNull" json:"operator2_list_null,omitempty"`
	Condition2List                *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=condition2_list,json=condition2List" json:"condition2_list,omitempty"`
	Condition2ListNull            bool                        `protobuf:"varint,1006,opt,name=condition2_list_null,json=condition2ListNull" json:"condition2_list_null,omitempty"`
	EstimatedRowsAffectedList     *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=estimated_rows_affected_list,json=estimatedRowsAffectedList" json:"estimated_rows_affected_list,omitempty"`
	EstimatedRowsAffectedListNull bool                        `protobuf:"varint,1007,opt,name=estimated_rows_affected_list_null,json=estimatedRowsAffectedListNull" json:"estimated_rows_affected_list_null,omitempty"`
	Delete                        *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=delete" json:"delete,omitempty"`
	DeleteNull                    bool                        `protobuf:"varint,1008,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
	Separator                     *dstore_values.StringValue  `protobuf:"bytes,9,opt,name=separator" json:"separator,omitempty"`
	SeparatorNull                 bool                        `protobuf:"varint,1009,opt,name=separator_null,json=separatorNull" json:"separator_null,omitempty"`
	CheckByteLengthForStrings     *dstore_values.BooleanValue `protobuf:"bytes,10,opt,name=check_byte_length_for_strings,json=checkByteLengthForStrings" json:"check_byte_length_for_strings,omitempty"`
	CheckByteLengthForStringsNull bool                        `protobuf:"varint,1010,opt,name=check_byte_length_for_strings_null,json=checkByteLengthForStringsNull" json:"check_byte_length_for_strings_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCharacteristicIdList() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicIdList
	}
	return nil
}

func (m *Parameters) GetBasicFieldTypeIdList() *dstore_values.StringValue {
	if m != nil {
		return m.BasicFieldTypeIdList
	}
	return nil
}

func (m *Parameters) GetOperator1List() *dstore_values.StringValue {
	if m != nil {
		return m.Operator1List
	}
	return nil
}

func (m *Parameters) GetCondition1List() *dstore_values.StringValue {
	if m != nil {
		return m.Condition1List
	}
	return nil
}

func (m *Parameters) GetOperator2List() *dstore_values.StringValue {
	if m != nil {
		return m.Operator2List
	}
	return nil
}

func (m *Parameters) GetCondition2List() *dstore_values.StringValue {
	if m != nil {
		return m.Condition2List
	}
	return nil
}

func (m *Parameters) GetEstimatedRowsAffectedList() *dstore_values.StringValue {
	if m != nil {
		return m.EstimatedRowsAffectedList
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Parameters) GetSeparator() *dstore_values.StringValue {
	if m != nil {
		return m.Separator
	}
	return nil
}

func (m *Parameters) GetCheckByteLengthForStrings() *dstore_values.BooleanValue {
	if m != nil {
		return m.CheckByteLengthForStrings
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_InsertTemp_CharacConditions.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_InsertTemp_CharacConditions.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_InsertTemp_CharacConditions.Response.Row")
}

var fileDescriptor0 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x6d, 0x4f, 0xd3, 0x5e,
	0x14, 0x0f, 0xec, 0xbf, 0x01, 0x87, 0x00, 0x7f, 0x0b, 0x21, 0x65, 0x73, 0x06, 0x31, 0x31, 0xbe,
	0x30, 0x9d, 0x8c, 0x17, 0x1a, 0x13, 0x8d, 0x40, 0x24, 0x99, 0x01, 0x42, 0x2a, 0x21, 0x51, 0x63,
	0x9a, 0xae, 0x3d, 0x83, 0xc6, 0xae, 0x77, 0xb9, 0xb7, 0x68, 0xf8, 0x16, 0xbe, 0xf4, 0xd3, 0xf8,
	0x7d, 0x7c, 0x7e, 0xf8, 0x04, 0x9e, 0xde, 0x73, 0xd9, 0x68, 0x19, 0x4c, 0x7c, 0xb3, 0xf5, 0xf4,
	0x9e, 0xdf, 0xc3, 0x69, 0x6f, 0x7f, 0x17, 0x1e, 0x85, 0x2a, 0x15, 0x12, 0x1b, 0x98, 0x1c, 0x46,
	0x09, 0x36, 0x7a, 0x52, 0x04, 0x18, 0x1e, 0x4b, 0x54, 0x8d, 0x6e, 0xe4, 0xb5, 0x12, 0x85, 0x32,
	0xdd, 0xc7, 0x6e, 0xcf, 0xdb, 0x3c, 0xf2, 0xa5, 0x1f, 0x6c, 0x8a, 0x24, 0x8c, 0xd2, 0x48, 0x24,
	0xca, 0xa1, 0xce, 0x54, 0x58, 0x77, 0x19, 0xee, 0x30, 0xdc, 0xb9, 0x1c, 0x53, 0x9d, 0x37, 0x62,
	0x6f, 0xfd, 0xf8, 0x18, 0x0d, 0x45, 0x75, 0x29, 0xef, 0x00, 0xa5, 0x14, 0xd2, 0x2c, 0xd5, 0xf2,
	0x4b, 0x5d, 0x54, 0xca, 0x3f, 0x44, 0xb3, 0x78, 0xab, 0xb8, 0x98, 0xfa, 0x51, 0xd2, 0x11, 0xb2,
	0xeb, 0x67, 0x62, 0xdc, 0xb4, 0xf2, 0x01, 0x00, 0xf6, 0xc8, 0x05, 0xad, 0xa2, 0x54, 0xd6, 0x1e,
	0x2c, 0x06, 0xda, 0x14, 0x55, 0x91, 0x4a, 0xa3, 0xc0, 0x8b, 0x42, 0x2f, 0xa6, 0x2b, 0x7b, 0x6c,
	0x79, 0xec, 0xce, 0x74, 0xb3, 0xea, 0x98, 0x79, 0x8c, 0x43, 0x95, 0xca, 0x28, 0x39, 0x3c, 0xc8,
	0x0a, 0x77, 0x21, 0x8f, 0x6c, 0x85, 0xdb, 0xf4, 0x6f, 0x3d, 0x86, 0xda, 0x70, 0x46, 0x2f, 0x39,
	0x8e, 0x63, 0xfb, 0xd3, 0x04, 0xf1, 0x4e, 0xba, 0xf6, 0x30, 0xec, 0x2e, 0x35, 0x58, 0x2e, 0xd8,
	0x6d, 0x5f, 0x11, 0xac, 0x13, 0x61, 0x1c, 0x7a, 0xe9, 0x49, 0x0f, 0xfb, 0x9e, 0xc6, 0x47, 0x7b,
	0xd2, 0xd8, 0xad, 0x0c, 0xba, 0x4f, 0x48, 0xe3, 0xe9, 0x09, 0xd4, 0x2f, 0xe2, 0x64, 0x57, 0x9f,
	0x8d, 0xab, 0x61, 0x68, 0xed, 0x6a, 0x1d, 0x66, 0x45, 0x0f, 0xa5, 0x4f, 0xba, 0xab, 0xec, 0xa5,
	0x34, 0xd2, 0xcb, 0x4c, 0x1f, 0xa1, 0x4d, 0x34, 0x60, 0x3e, 0x4f, 0xc1, 0xd2, 0x5f, 0x58, 0xfa,
	0x5a, 0xae, 0x59, 0x6b, 0x6e, 0xc2, 0x5c, 0x70, 0xba, 0x55, 0x8c, 0xe8, 0x7f, 0x23, 0x45, 0x67,
	0x07, 0x10, 0xad, 0xba, 0x0a, 0x0b, 0x05, 0x12, 0x96, 0xfd, 0xca, 0xb2, 0x56, 0xbe, 0xbd, 0x38,
	0x6b, 0x93, 0x65, 0xcb, 0x7f, 0x3f, 0x6b, 0xb3, 0x38, 0x6b, 0xf3, 0x8c, 0xe8, 0xb7, 0xc2, 0xac,
	0xcd, 0xa1, 0xb3, 0x1a, 0xd1, 0xca, 0x15, 0x66, 0x6d, 0x9e, 0x9b, 0xf5, 0xac, 0xec, 0xf7, 0xe2,
	0xac, 0x03, 0xdd, 0x57, 0x70, 0x1d, 0x69, 0x03, 0xd2, 0x27, 0x82, 0xa1, 0x27, 0xc5, 0x3b, 0xe5,
	0xf9, 0x9d, 0x0e, 0x06, 0x59, 0xa5, 0x4d, 0x4c, 0x8c, 0x34, 0xb1, 0xd4, 0xc7, 0xbb, 0x04, 0x5f,
	0x37, 0x68, 0xed, 0xa7, 0x05, 0x37, 0x2f, 0x23, 0x67, 0x73, 0x3f, 0xd8, 0x5c, 0xfd, 0x42, 0x1a,
	0xed, 0x73, 0x0d, 0x2a, 0x21, 0xc6, 0xf4, 0xcd, 0xda, 0x93, 0xda, 0x51, 0xad, 0xe0, 0xa8, 0x2d,
	0x44, 0x8c, 0x7e, 0xc2, 0x96, 0x4c, 0xab, 0xb5, 0x0c, 0xd3, 0x7c, 0xc5, 0x4a, 0x3f, 0x59, 0x09,
	0xf8, 0x9e, 0xa6, 0x7d, 0x00, 0x53, 0x0a, 0x7b, 0xbe, 0x7e, 0x19, 0xf6, 0xd4, 0xc8, 0x59, 0x07,
	0xcd, 0xd6, 0x6d, 0x98, 0xed, 0x17, 0x4c, 0xff, 0x8b, 0xe9, 0x67, 0xfa, 0xb7, 0xb5, 0xc2, 0x6b,
	0xa8, 0x07, 0x47, 0x18, 0xbc, 0xf1, 0xda, 0x27, 0xe4, 0x23, 0xa6, 0x6c, 0x4a, 0x8f, 0x3c, 0x0a,
	0x25, 0x8f, 0x79, 0x95, 0x0d, 0xa3, 0xe7, 0x59, 0xd2, 0x0c, 0x1b, 0x44, 0xb0, 0xad, 0xf1, 0x5b,
	0x42, 0x3e, 0x67, 0xb4, 0xf5, 0x0c, 0x56, 0x2e, 0xa5, 0x67, 0x6b, 0xbf, 0xcd, 0x33, 0xbe, 0x90,
	0x27, 0xb3, 0xba, 0xf2, 0x71, 0x1c, 0x26, 0x5d, 0x54, 0x3d, 0x0a, 0x66, 0xb4, 0xee, 0x41, 0x59,
	0x07, 0x6f, 0x31, 0x07, 0x4d, 0xae, 0x73, 0x28, 0x3f, 0xcd, 0x7e, 0x5d, 0x6e, 0xb4, 0x5e, 0xc0,
	0xff, 0x59, 0xe4, 0x7a, 0x67, 0x32, 0x97, 0x02, 0xab, 0x44, 0x60, 0xa7, 0x00, 0x2e, 0x26, 0xf3,
	0x0e, 0xd5, 0xad, 0x41, 0xed, 0xce, 0x75, 0xf3, 0x37, 0xe8, 0x35, 0x4d, 0x98, 0xa8, 0xa7, 0xd8,
	0xc9, 0x18, 0x6f, 0x9c, 0x63, 0xe4, 0x83, 0x60, 0x87, 0xff, 0xdd, 0xd3, 0x76, 0x6b, 0x1b, 0x4a,
	0xb4, 0xf1, 0x28, 0x37, 0x32, 0xd4, 0x43, 0xe7, 0x2a, 0x87, 0x93, 0x73, 0xfa, 0x2c, 0x1c, 0xda,
	0x91, 0x6e, 0x46, 0x53, 0xad, 0x43, 0x89, 0xae, 0xad, 0x45, 0xa8, 0x50, 0x45, 0x09, 0x6a, 0xbf,
	0xdf, 0xa5, 0xa7, 0x53, 0x76, 0xcb, 0x54, 0xb6, 0xc2, 0x8d, 0x03, 0xa8, 0x45, 0xa2, 0xa0, 0x31,
	0x38, 0x3f, 0x5f, 0xde, 0xff, 0xc7, 0x93, 0xb5, 0x5d, 0xd1, 0x47, 0xd7, 0xda, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0xe8, 0xa6, 0xf1, 0x9b, 0x07, 0x00, 0x00,
}