// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/do_GetProcedureResultSets_Ad.proto
// DO NOT EDIT!

/*
Package do_GetProcedureResultSets_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/do_GetProcedureResultSets_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package do_GetProcedureResultSets_Ad

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
	ProcedureName                  *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	ProcedureNameNull              bool                        `protobuf:"varint,1001,opt,name=procedure_name_null,json=procedureNameNull" json:"procedure_name_null,omitempty"`
	ProcedureResultConditionId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=procedure_result_condition_id,json=procedureResultConditionId" json:"procedure_result_condition_id,omitempty"`
	ProcedureResultConditionIdNull bool                        `protobuf:"varint,1002,opt,name=procedure_result_condition_id_null,json=procedureResultConditionIdNull" json:"procedure_result_condition_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
	}
	return nil
}

func (m *Parameters) GetProcedureResultConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureResultConditionId
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
	RowId                        int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Condition                    *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=condition" json:"condition,omitempty"`
	ProcedureResultConditionId   *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=procedure_result_condition_id,json=procedureResultConditionId" json:"procedure_result_condition_id,omitempty"`
	DescriptionValidSinceVersion *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=description_valid_since_version,json=descriptionValidSinceVersion" json:"description_valid_since_version,omitempty"`
	Description                  *dstore_values.StringValue    `protobuf:"bytes,10004,opt,name=description" json:"description,omitempty"`
	PrecisionValue               *dstore_values.IntegerValue   `protobuf:"bytes,10005,opt,name=precision_value,json=precisionValue" json:"precision_value,omitempty"`
	ParameterName                *dstore_values.StringValue    `protobuf:"bytes,10006,opt,name=parameter_name,json=parameterName" json:"parameter_name,omitempty"`
	Scale                        *dstore_values.IntegerValue   `protobuf:"bytes,10007,opt,name=scale" json:"scale,omitempty"`
	SortNo                       *dstore_values.IntegerValue   `protobuf:"bytes,10008,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	ResultColumnName             *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=result_column_name,json=resultColumnName" json:"result_column_name,omitempty"`
	Length                       *dstore_values.IntegerValue   `protobuf:"bytes,10010,opt,name=length" json:"length,omitempty"`
	DescriptionLastEdited        *dstore_values.TimestampValue `protobuf:"bytes,10011,opt,name=description_last_edited,json=descriptionLastEdited" json:"description_last_edited,omitempty"`
	IntroducedIndstoreVersion    *dstore_values.StringValue    `protobuf:"bytes,10012,opt,name=introduced_indstore_version,json=introducedIndstoreVersion" json:"introduced_indstore_version,omitempty"`
	DataType                     *dstore_values.StringValue    `protobuf:"bytes,10013,opt,name=data_type,json=dataType" json:"data_type,omitempty"`
	ProcResultCondDescription    *dstore_values.StringValue    `protobuf:"bytes,10014,opt,name=proc_result_cond_description,json=procResultCondDescription" json:"proc_result_cond_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCondition() *dstore_values.StringValue {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *Response_Row) GetProcedureResultConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureResultConditionId
	}
	return nil
}

func (m *Response_Row) GetDescriptionValidSinceVersion() *dstore_values.StringValue {
	if m != nil {
		return m.DescriptionValidSinceVersion
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetPrecisionValue() *dstore_values.IntegerValue {
	if m != nil {
		return m.PrecisionValue
	}
	return nil
}

func (m *Response_Row) GetParameterName() *dstore_values.StringValue {
	if m != nil {
		return m.ParameterName
	}
	return nil
}

func (m *Response_Row) GetScale() *dstore_values.IntegerValue {
	if m != nil {
		return m.Scale
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetResultColumnName() *dstore_values.StringValue {
	if m != nil {
		return m.ResultColumnName
	}
	return nil
}

func (m *Response_Row) GetLength() *dstore_values.IntegerValue {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *Response_Row) GetDescriptionLastEdited() *dstore_values.TimestampValue {
	if m != nil {
		return m.DescriptionLastEdited
	}
	return nil
}

func (m *Response_Row) GetIntroducedIndstoreVersion() *dstore_values.StringValue {
	if m != nil {
		return m.IntroducedIndstoreVersion
	}
	return nil
}

func (m *Response_Row) GetDataType() *dstore_values.StringValue {
	if m != nil {
		return m.DataType
	}
	return nil
}

func (m *Response_Row) GetProcResultCondDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ProcResultCondDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.do_GetProcedureResultSets_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.do_GetProcedureResultSets_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.do_GetProcedureResultSets_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/do_GetProcedureResultSets_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xeb, 0x6e, 0xd3, 0x4a,
	0x10, 0x56, 0x9b, 0x93, 0x34, 0x9d, 0xea, 0xb4, 0x3d, 0x5b, 0x9d, 0x73, 0xdc, 0xf4, 0x02, 0x2a,
	0x7f, 0x10, 0x3f, 0x1c, 0x54, 0x40, 0x2a, 0x42, 0x20, 0xb5, 0x50, 0xa1, 0x00, 0x8d, 0x2a, 0x17,
	0x45, 0xe2, 0x6a, 0xb9, 0xf1, 0x12, 0x2c, 0x1c, 0xaf, 0xb5, 0xbb, 0x6e, 0xc5, 0x5b, 0x70, 0xbf,
	0x83, 0x78, 0x01, 0xde, 0x81, 0xd7, 0x81, 0xa7, 0x60, 0xd6, 0x6b, 0xc7, 0x8e, 0x11, 0xb5, 0x11,
	0x7f, 0x12, 0x8d, 0x67, 0xbe, 0xcf, 0xdf, 0xce, 0x7c, 0xe3, 0x85, 0x0b, 0xae, 0x90, 0x8c, 0xd3,
	0x36, 0x0d, 0x06, 0x5e, 0x40, 0xdb, 0x21, 0x67, 0x7d, 0xea, 0x46, 0x9c, 0x8a, 0xb6, 0xcb, 0xec,
	0xab, 0x54, 0xee, 0xa6, 0x0f, 0x2c, 0x2a, 0x22, 0x5f, 0xee, 0x51, 0x29, 0xec, 0x4d, 0xd7, 0xc4,
	0x3a, 0xc9, 0xc8, 0x29, 0x0d, 0x36, 0x35, 0xd8, 0x3c, 0x0a, 0xd1, 0x5a, 0x48, 0x5e, 0x74, 0xe0,
	0xf8, 0x11, 0x15, 0x9a, 0xa0, 0xb5, 0x38, 0xfe, 0x76, 0xca, 0x39, 0xe3, 0x49, 0x6a, 0x69, 0x3c,
	0x35, 0xa4, 0x42, 0x38, 0x03, 0x9a, 0x24, 0x4f, 0x14, 0x93, 0xd2, 0xf1, 0x82, 0x07, 0x8c, 0x0f,
	0x1d, 0xe9, 0xb1, 0x40, 0x17, 0xad, 0x7d, 0x99, 0x04, 0xd8, 0x75, 0xb8, 0x83, 0x59, 0xca, 0x05,
	0xd9, 0x84, 0xd9, 0xd1, 0xe9, 0xec, 0x00, 0x9f, 0x1a, 0x13, 0xc7, 0x27, 0x4e, 0xce, 0xac, 0xb7,
	0xcc, 0xe4, 0x14, 0x89, 0x32, 0x21, 0xb9, 0x17, 0x0c, 0x7a, 0x2a, 0xb0, 0xfe, 0x1e, 0x21, 0xba,
	0x08, 0x20, 0x6d, 0x58, 0x18, 0xa7, 0xb0, 0x83, 0xc8, 0xf7, 0x8d, 0x6f, 0x53, 0x48, 0xd4, 0xb4,
	0xfe, 0x19, 0x2b, 0xee, 0x62, 0x86, 0xdc, 0x87, 0x95, 0x0c, 0xc0, 0xe3, 0x7e, 0xd8, 0x7d, 0x16,
	0xb8, 0x9e, 0x92, 0x69, 0x7b, 0xae, 0x31, 0x19, 0x4b, 0x58, 0x2a, 0x48, 0xf0, 0x02, 0x49, 0x07,
	0x94, 0x6b, 0x0d, 0xad, 0x70, 0xbc, 0xa1, 0x97, 0x53, 0x7c, 0xc7, 0x25, 0xd7, 0x61, 0xed, 0x48,
	0x7e, 0xad, 0xef, 0xbb, 0xd6, 0xb7, 0xfa, 0x6b, 0x22, 0x25, 0x76, 0xed, 0x2b, 0x40, 0x13, 0x33,
	0x21, 0x0b, 0x04, 0x25, 0xa7, 0xa1, 0x1e, 0x4f, 0xa3, 0xd8, 0xa4, 0x64, 0xd4, 0x7a, 0x52, 0xdb,
	0xea, 0xd7, 0xd2, 0x85, 0xe4, 0x16, 0xcc, 0xab, 0x39, 0xd8, 0xb9, 0x41, 0xe0, 0xf1, 0x6a, 0x08,
	0x36, 0x0b, 0xe0, 0xe2, 0xb8, 0x76, 0x30, 0xee, 0x64, 0xb1, 0x35, 0x37, 0x1c, 0x7f, 0x40, 0x36,
	0x60, 0x2a, 0x99, 0xbf, 0x51, 0x8b, 0x19, 0x57, 0x7f, 0x62, 0xd4, 0xee, 0xd8, 0xd1, 0xff, 0x56,
	0x5a, 0x4e, 0xae, 0x41, 0x8d, 0xb3, 0x43, 0xe3, 0xaf, 0x18, 0xb5, 0x61, 0x56, 0xf7, 0xab, 0x99,
	0x76, 0xc2, 0xb4, 0xd8, 0xa1, 0xa5, 0x48, 0x5a, 0x9f, 0x9b, 0x50, 0xc3, 0x80, 0xfc, 0x07, 0x0d,
	0x0c, 0xd5, 0xf4, 0x9e, 0x74, 0xb1, 0x39, 0x75, 0xab, 0x8e, 0x21, 0x0e, 0xe3, 0x3c, 0x4c, 0x8f,
	0x7a, 0x6f, 0x3c, 0xed, 0x96, 0x9a, 0x2b, 0xab, 0x26, 0x76, 0x99, 0x4f, 0x9e, 0x75, 0xff, 0xcc,
	0x28, 0xfb, 0x70, 0xcc, 0xa5, 0xa2, 0xcf, 0xbd, 0x30, 0x66, 0x44, 0x38, 0x9a, 0x43, 0x78, 0x41,
	0x9f, 0xda, 0x07, 0xb8, 0x1c, 0x4a, 0xf1, 0xf3, 0x72, 0xc5, 0xcb, 0x39, 0x8e, 0x9e, 0xa2, 0xd8,
	0x53, 0x0c, 0x3d, 0x4d, 0x40, 0x2e, 0xc2, 0x4c, 0x2e, 0x6f, 0xbc, 0x28, 0xe7, 0xcb, 0xd7, 0x93,
	0x6d, 0x98, 0x0b, 0x39, 0xed, 0x7b, 0x22, 0x11, 0x18, 0x51, 0xe3, 0x65, 0x85, 0x53, 0xcf, 0x8e,
	0x40, 0x71, 0x4c, 0xb6, 0x70, 0xcd, 0xd3, 0xa5, 0xd7, 0x6b, 0xfe, 0xaa, 0x5b, 0x61, 0xcf, 0x53,
	0x48, 0xbc, 0xe7, 0xeb, 0x50, 0x17, 0x7d, 0xc7, 0xa7, 0xc6, 0xeb, 0x0a, 0x02, 0x74, 0x29, 0x39,
	0x07, 0x53, 0x82, 0x71, 0x69, 0x07, 0xcc, 0x78, 0x53, 0x01, 0xd5, 0x50, 0xc5, 0x5d, 0x46, 0x3a,
	0x40, 0x46, 0xf3, 0xf6, 0xa3, 0x61, 0xa0, 0x25, 0xbf, 0x2d, 0x97, 0x3c, 0xcf, 0x93, 0x21, 0x2b,
	0x54, 0xac, 0xfa, 0x2c, 0x34, 0x7c, 0x74, 0xb6, 0x7c, 0x68, 0xbc, 0xab, 0x22, 0x40, 0xd7, 0x92,
	0x1e, 0xfc, 0x9f, 0x77, 0x86, 0xef, 0x08, 0x69, 0x53, 0xf4, 0x0d, 0x75, 0x8d, 0xf7, 0x9a, 0x66,
	0xa5, 0x40, 0x23, 0x3d, 0xdc, 0x2e, 0xe9, 0x0c, 0x43, 0x4d, 0xf4, 0x6f, 0x0e, 0x7e, 0x03, 0xd1,
	0xdb, 0x31, 0x98, 0xdc, 0x81, 0x25, 0x7c, 0x1f, 0x67, 0x6e, 0x84, 0x9e, 0xc4, 0x8f, 0x82, 0xe6,
	0x18, 0xb9, 0xed, 0x43, 0xf9, 0x09, 0x17, 0x33, 0x7c, 0x27, 0x81, 0xa7, 0x56, 0xdb, 0x80, 0x69,
	0xd7, 0xc1, 0x6f, 0x8d, 0x7c, 0x1c, 0x52, 0xe3, 0x63, 0x39, 0x55, 0x53, 0x55, 0xdf, 0xc4, 0x62,
	0x72, 0x17, 0x96, 0xd5, 0x9a, 0xe4, 0x97, 0xcc, 0xce, 0xbb, 0xf6, 0x53, 0x05, 0x5d, 0x8a, 0x20,
	0x5b, 0xb1, 0x2b, 0x19, 0x7a, 0xeb, 0x1e, 0x1e, 0x9a, 0x15, 0xbe, 0x32, 0xd9, 0x95, 0x7a, 0xfb,
	0xd2, 0x80, 0x09, 0xf7, 0x51, 0x9a, 0x77, 0x7f, 0xf7, 0xd6, 0xdd, 0x6f, 0xc4, 0x17, 0xdb, 0x99,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x0c, 0x5c, 0x43, 0xb5, 0x07, 0x00, 0x00,
}
