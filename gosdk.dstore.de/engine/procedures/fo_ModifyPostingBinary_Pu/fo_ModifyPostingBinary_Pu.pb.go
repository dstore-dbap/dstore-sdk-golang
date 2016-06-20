// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_ModifyPostingBinary_Pu.proto
// DO NOT EDIT!

/*
Package fo_ModifyPostingBinary_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_ModifyPostingBinary_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_ModifyPostingBinary_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PostingId                      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=posting_id,json=postingId" json:"posting_id,omitempty"`
	PostingIdNull                  bool                        `protobuf:"varint,1004,opt,name=posting_id_null,json=postingIdNull" json:"posting_id_null,omitempty"`
	SortNoOfBinaryToModify         *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=sort_no_of_binary_to_modify,json=sortNoOfBinaryToModify" json:"sort_no_of_binary_to_modify,omitempty"`
	SortNoOfBinaryToModifyNull     bool                        `protobuf:"varint,1005,opt,name=sort_no_of_binary_to_modify_null,json=sortNoOfBinaryToModifyNull" json:"sort_no_of_binary_to_modify_null,omitempty"`
	NewDescription                 *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=new_description,json=newDescription" json:"new_description,omitempty"`
	NewDescriptionNull             bool                        `protobuf:"varint,1006,opt,name=new_description_null,json=newDescriptionNull" json:"new_description_null,omitempty"`
	MoveSortNo                     *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=move_sort_no,json=moveSortNo" json:"move_sort_no,omitempty"`
	MoveSortNoNull                 bool                        `protobuf:"varint,1007,opt,name=move_sort_no_null,json=moveSortNoNull" json:"move_sort_no_null,omitempty"`
	PostingBinaryIdentifier        *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=posting_binary_identifier,json=postingBinaryIdentifier" json:"posting_binary_identifier,omitempty"`
	PostingBinaryIdentifierNull    bool                        `protobuf:"varint,1008,opt,name=posting_binary_identifier_null,json=postingBinaryIdentifierNull" json:"posting_binary_identifier_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,9,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1009,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingId
	}
	return nil
}

func (m *Parameters) GetSortNoOfBinaryToModify() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNoOfBinaryToModify
	}
	return nil
}

func (m *Parameters) GetNewDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NewDescription
	}
	return nil
}

func (m *Parameters) GetMoveSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.MoveSortNo
	}
	return nil
}

func (m *Parameters) GetPostingBinaryIdentifier() *dstore_values.StringValue {
	if m != nil {
		return m.PostingBinaryIdentifier
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_ModifyPostingBinary_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_ModifyPostingBinary_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_ModifyPostingBinary_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_ModifyPostingBinary_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x4e, 0x14, 0x31,
	0x14, 0x0e, 0xac, 0x0b, 0xbb, 0x47, 0x5c, 0xb0, 0x10, 0x18, 0x76, 0x91, 0x20, 0xc6, 0x60, 0xfc,
	0x31, 0x78, 0x89, 0x91, 0x68, 0x8c, 0x11, 0x21, 0x66, 0x4d, 0x40, 0x1c, 0x09, 0x51, 0x62, 0xd2,
	0x0c, 0x4c, 0x77, 0xd3, 0xc8, 0xb6, 0x63, 0x3b, 0x0b, 0xe1, 0x2d, 0x7c, 0x08, 0x5f, 0xc2, 0x47,
	0xf2, 0xae, 0x6f, 0x60, 0xdb, 0x33, 0x7b, 0x1b, 0x59, 0x46, 0xff, 0xec, 0x6c, 0xe7, 0x9c, 0xef,
	0xd2, 0x4e, 0xfb, 0x15, 0xd6, 0x22, 0x9d, 0x48, 0xc5, 0x56, 0x99, 0x68, 0x72, 0xc1, 0x56, 0x63,
	0x25, 0x0f, 0x59, 0xd4, 0x56, 0x4c, 0xaf, 0x36, 0x24, 0xdd, 0x92, 0x11, 0x6f, 0x9c, 0xee, 0x48,
	0x9d, 0x70, 0xd1, 0x5c, 0xe7, 0x22, 0x54, 0xa7, 0x74, 0xa7, 0xed, 0x9b, 0xa6, 0x44, 0x92, 0x15,
	0x44, 0xfa, 0x88, 0xf4, 0x87, 0xb6, 0x57, 0xa7, 0x53, 0x89, 0xe3, 0xf0, 0xa8, 0xcd, 0x34, 0xa2,
	0xab, 0xf3, 0x83, 0xba, 0x4c, 0x29, 0xa9, 0xd2, 0x52, 0x6d, 0xb0, 0xd4, 0x62, 0x5a, 0x87, 0x4d,
	0x96, 0x16, 0xaf, 0x65, 0x8b, 0x49, 0xc8, 0x45, 0x43, 0xaa, 0x56, 0x98, 0x70, 0x29, 0xb0, 0x69,
	0xf9, 0x63, 0x19, 0x60, 0x27, 0x54, 0xa1, 0xa9, 0x32, 0xa5, 0xc9, 0x5b, 0x58, 0x88, 0xcd, 0x53,
	0x0a, 0xca, 0x23, 0x26, 0x12, 0xde, 0xe0, 0x87, 0xae, 0x9b, 0xa2, 0x23, 0x6f, 0x64, 0x69, 0xe4,
	0xc6, 0xc5, 0x3b, 0x55, 0x3f, 0x9d, 0x50, 0xea, 0x53, 0x27, 0xca, 0x4c, 0x63, 0xcf, 0x0e, 0x82,
	0x2a, 0xe2, 0xeb, 0x03, 0x70, 0x57, 0xd2, 0xe4, 0x39, 0x5c, 0x3d, 0x8f, 0x9d, 0x8a, 0xf6, 0xd1,
	0x91, 0xf7, 0x79, 0xdc, 0x68, 0x94, 0x82, 0xc5, 0xe1, 0x3c, 0xdb, 0xa6, 0x8d, 0x3c, 0x81, 0x4a,
	0xca, 0x95, 0x9c, 0xc6, 0xcc, 0x10, 0x7a, 0xa3, 0xce, 0x5b, 0x2d, 0xe3, 0x8d, 0x8b, 0x84, 0x35,
	0x99, 0x42, 0x73, 0x13, 0x08, 0xd9, 0x35, 0x88, 0x7a, 0x44, 0x7c, 0x98, 0x1e, 0xa4, 0x40, 0x03,
	0x5f, 0xd0, 0xc0, 0x54, 0x7f, 0xaf, 0x93, 0xbc, 0x0f, 0xe5, 0xb6, 0xe0, 0xef, 0xdb, 0x4e, 0xad,
	0x90, 0xbb, 0x12, 0x25, 0x6c, 0x36, 0x42, 0xd7, 0xa1, 0xd2, 0x05, 0xa2, 0xc6, 0x57, 0xd4, 0x98,
	0xe8, 0xb4, 0x38, 0xfe, 0x07, 0x00, 0x31, 0xee, 0x08, 0x2b, 0x70, 0x21, 0x7f, 0x3a, 0xe5, 0xb4,
	0xdd, 0x48, 0xac, 0xc0, 0x64, 0x0f, 0x8b, 0x1a, 0xdf, 0x50, 0xe3, 0x52, 0xb7, 0xc9, 0x89, 0xbc,
	0x86, 0x9a, 0x96, 0x2a, 0xa1, 0x42, 0x52, 0xd9, 0xa0, 0x07, 0xb8, 0xf5, 0x12, 0x49, 0x5b, 0x6e,
	0x47, 0x7a, 0xc5, 0x7c, 0xd5, 0x59, 0x8b, 0xdf, 0x96, 0x2f, 0x1a, 0xb8, 0x6f, 0x77, 0x25, 0x6e,
	0x66, 0xb2, 0x09, 0x4b, 0xe7, 0x30, 0xa3, 0xa7, 0xef, 0xe8, 0xa9, 0x7a, 0x36, 0x85, 0x33, 0xf8,
	0x14, 0x26, 0x05, 0x3b, 0xa1, 0x11, 0xd3, 0x87, 0x8a, 0xc7, 0xf6, 0xab, 0x7b, 0x63, 0xb9, 0x6b,
	0x5d, 0x31, 0x90, 0x8d, 0x1e, 0x82, 0xdc, 0x86, 0x99, 0x0c, 0x09, 0xea, 0xff, 0x40, 0x7d, 0x32,
	0xd8, 0xee, 0x74, 0x1f, 0xc1, 0x44, 0x4b, 0x1e, 0x33, 0x9a, 0xce, 0xc1, 0x1b, 0xcf, 0x5f, 0x09,
	0xb0, 0x80, 0x57, 0x6e, 0x2a, 0xe4, 0x26, 0x5c, 0xee, 0x87, 0xa3, 0xdc, 0x4f, 0x94, 0xab, 0xf4,
	0xfa, 0x9c, 0xd4, 0x1e, 0xcc, 0x77, 0x3e, 0x56, 0xba, 0x4c, 0x9d, 0xf3, 0xc0, 0x94, 0x57, 0xca,
	0x9d, 0xec, 0x5c, 0xdc, 0x9f, 0x1b, 0xf5, 0x2e, 0x94, 0x6c, 0xc0, 0xe2, 0x50, 0x5e, 0x34, 0xf4,
	0x0b, 0x0d, 0xd5, 0x86, 0x30, 0x38, 0x77, 0x2f, 0x61, 0x4e, 0xb3, 0xd8, 0x64, 0x82, 0xd1, 0xa7,
	0x3c, 0x3d, 0xab, 0xf6, 0x88, 0x6a, 0xaf, 0x9c, 0xeb, 0x6d, 0xa6, 0x0b, 0xad, 0xe3, 0xd9, 0x35,
	0xaf, 0x35, 0x79, 0x0c, 0x0b, 0x43, 0x28, 0xd1, 0xd6, 0x6f, 0xb4, 0xe5, 0x9d, 0x05, 0xb6, 0x9e,
	0x96, 0x3f, 0x8d, 0x42, 0x29, 0x60, 0x3a, 0x96, 0x42, 0x33, 0x72, 0x0b, 0x8a, 0x2e, 0x04, 0xb3,
	0x69, 0x94, 0xc6, 0x2b, 0x06, 0xe4, 0xa6, 0xfd, 0x0d, 0xb0, 0x91, 0xbc, 0x81, 0x29, 0x1b, 0x7f,
	0xb4, 0x2f, 0xff, 0x4c, 0x5c, 0x14, 0x0c, 0xd8, 0xcf, 0x80, 0xb3, 0x29, 0xb9, 0x65, 0xc6, 0xf5,
	0xde, 0x38, 0x98, 0x6c, 0x0d, 0xbe, 0x20, 0x6b, 0x30, 0x9e, 0xc6, 0xae, 0x89, 0x04, 0xcb, 0xb8,
	0xf8, 0x17, 0x23, 0x86, 0xf2, 0x16, 0x3e, 0x83, 0x4e, 0x3b, 0x79, 0x06, 0x05, 0x25, 0x4f, 0xcc,
	0x39, 0xb7, 0xa8, 0x7b, 0xfe, 0x3f, 0xde, 0x11, 0x7e, 0x67, 0x19, 0xfc, 0x40, 0x9e, 0x04, 0x96,
	0xa1, 0x7a, 0x05, 0x0a, 0xe6, 0x3f, 0x99, 0x85, 0x31, 0x33, 0xb2, 0xd1, 0xf1, 0x61, 0xdb, 0x2c,
	0x4c, 0x31, 0x28, 0x9a, 0x61, 0x3d, 0x5a, 0xdf, 0x87, 0x1a, 0x97, 0x19, 0xfa, 0xde, 0xe5, 0xb5,
	0xff, 0xb0, 0x29, 0x75, 0xf4, 0xae, 0x53, 0x8f, 0xfe, 0xeb, 0x7e, 0x3b, 0x18, 0x73, 0xb7, 0xc8,
	0xdd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0x72, 0x36, 0x22, 0x1c, 0x07, 0x00, 0x00,
}