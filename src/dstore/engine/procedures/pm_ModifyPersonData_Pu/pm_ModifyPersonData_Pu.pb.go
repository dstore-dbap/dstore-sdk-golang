// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonData_Pu.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonData_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonData_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonData_Pu

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
	CharacteristicIdList             *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=characteristic_id_list,json=characteristicIdList" json:"characteristic_id_list,omitempty"`
	CharacteristicIdListNull         bool                        `protobuf:"varint,1001,opt,name=characteristic_id_list_null,json=characteristicIdListNull" json:"characteristic_id_list_null,omitempty"`
	ValueList                        *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=value_list,json=valueList" json:"value_list,omitempty"`
	ValueListNull                    bool                        `protobuf:"varint,1002,opt,name=value_list_null,json=valueListNull" json:"value_list_null,omitempty"`
	IdentificationValues             *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=identification_values,json=identificationValues" json:"identification_values,omitempty"`
	IdentificationValuesNull         bool                        `protobuf:"varint,1003,opt,name=identification_values_null,json=identificationValuesNull" json:"identification_values_null,omitempty"`
	PersonId                         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                     bool                        `protobuf:"varint,1004,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	PersonTypeId                     *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull                 bool                        `protobuf:"varint,1005,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	PersonGrantAccessIds             *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=person_grant_access_ids,json=personGrantAccessIds" json:"person_grant_access_ids,omitempty"`
	PersonGrantAccessIdsNull         bool                        `protobuf:"varint,1006,opt,name=person_grant_access_ids_null,json=personGrantAccessIdsNull" json:"person_grant_access_ids_null,omitempty"`
	PersonCharacCategoryId           *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=person_charac_category_id,json=personCharacCategoryId" json:"person_charac_category_id,omitempty"`
	PersonCharacCategoryIdNull       bool                        `protobuf:"varint,1007,opt,name=person_charac_category_id_null,json=personCharacCategoryIdNull" json:"person_charac_category_id_null,omitempty"`
	DeleteCharacCategoryId           *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=delete_charac_category_id,json=deleteCharacCategoryId" json:"delete_charac_category_id,omitempty"`
	DeleteCharacCategoryIdNull       bool                        `protobuf:"varint,1008,opt,name=delete_charac_category_id_null,json=deleteCharacCategoryIdNull" json:"delete_charac_category_id_null,omitempty"`
	ResultInErrorIdList              *dstore_values.BooleanValue `protobuf:"bytes,9,opt,name=result_in_error_id_list,json=resultInErrorIdList" json:"result_in_error_id_list,omitempty"`
	ResultInErrorIdListNull          bool                        `protobuf:"varint,1009,opt,name=result_in_error_id_list_null,json=resultInErrorIdListNull" json:"result_in_error_id_list_null,omitempty"`
	ValueIdsForPredefinedCharacs     *dstore_values.BooleanValue `protobuf:"bytes,10,opt,name=value_ids_for_predefined_characs,json=valueIdsForPredefinedCharacs" json:"value_ids_for_predefined_characs,omitempty"`
	ValueIdsForPredefinedCharacsNull bool                        `protobuf:"varint,1010,opt,name=value_ids_for_predefined_characs_null,json=valueIdsForPredefinedCharacsNull" json:"value_ids_for_predefined_characs_null,omitempty"`
	ChangeAllOrNothing               *dstore_values.BooleanValue `protobuf:"bytes,11,opt,name=change_all_or_nothing,json=changeAllOrNothing" json:"change_all_or_nothing,omitempty"`
	ChangeAllOrNothingNull           bool                        `protobuf:"varint,1011,opt,name=change_all_or_nothing_null,json=changeAllOrNothingNull" json:"change_all_or_nothing_null,omitempty"`
	CaseSensitive                    *dstore_values.BooleanValue `protobuf:"bytes,12,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull                bool                        `protobuf:"varint,1012,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
	Country                          *dstore_values.StringValue  `protobuf:"bytes,13,opt,name=country" json:"country,omitempty"`
	CountryNull                      bool                        `protobuf:"varint,1013,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
	SeparatorInValueList             *dstore_values.StringValue  `protobuf:"bytes,14,opt,name=separator_in_value_list,json=separatorInValueList" json:"separator_in_value_list,omitempty"`
	SeparatorInValueListNull         bool                        `protobuf:"varint,1014,opt,name=separator_in_value_list_null,json=separatorInValueListNull" json:"separator_in_value_list_null,omitempty"`
	SeparatorForIdentValues          *dstore_values.StringValue  `protobuf:"bytes,15,opt,name=separator_for_ident_values,json=separatorForIdentValues" json:"separator_for_ident_values,omitempty"`
	SeparatorForIdentValuesNull      bool                        `protobuf:"varint,1015,opt,name=separator_for_ident_values_null,json=separatorForIdentValuesNull" json:"separator_for_ident_values_null,omitempty"`
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

func (m *Parameters) GetValueList() *dstore_values.StringValue {
	if m != nil {
		return m.ValueList
	}
	return nil
}

func (m *Parameters) GetIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.IdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetPersonGrantAccessIds() *dstore_values.BooleanValue {
	if m != nil {
		return m.PersonGrantAccessIds
	}
	return nil
}

func (m *Parameters) GetPersonCharacCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacCategoryId
	}
	return nil
}

func (m *Parameters) GetDeleteCharacCategoryId() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCharacCategoryId
	}
	return nil
}

func (m *Parameters) GetResultInErrorIdList() *dstore_values.BooleanValue {
	if m != nil {
		return m.ResultInErrorIdList
	}
	return nil
}

func (m *Parameters) GetValueIdsForPredefinedCharacs() *dstore_values.BooleanValue {
	if m != nil {
		return m.ValueIdsForPredefinedCharacs
	}
	return nil
}

func (m *Parameters) GetChangeAllOrNothing() *dstore_values.BooleanValue {
	if m != nil {
		return m.ChangeAllOrNothing
	}
	return nil
}

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Parameters) GetSeparatorInValueList() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInValueList
	}
	return nil
}

func (m *Parameters) GetSeparatorForIdentValues() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorForIdentValues
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ErrorIdList     *dstore_values.StringValue                       `protobuf:"bytes,101,opt,name=error_id_list,json=errorIdList" json:"error_id_list,omitempty"`
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

func (m *Response) GetErrorIdList() *dstore_values.StringValue {
	if m != nil {
		return m.ErrorIdList
	}
	return nil
}

type Response_Row struct {
	RowId                  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PersonCharacteristicId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	ResultCode             *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=result_code,json=resultCode" json:"result_code,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetResultCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ResultCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonData_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonData_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonData_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 971 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xeb, 0x6f, 0x1b, 0x45,
	0x10, 0x57, 0x48, 0xf3, 0x1a, 0xe7, 0xc5, 0x85, 0x3a, 0x57, 0x3b, 0x2a, 0x51, 0xa0, 0x82, 0x4f,
	0x67, 0x54, 0x40, 0x02, 0x21, 0x40, 0x69, 0xda, 0x22, 0x4b, 0xc4, 0x35, 0x07, 0x0a, 0x82, 0x2f,
	0xa7, 0xeb, 0xdd, 0xda, 0x5d, 0xe9, 0xb2, 0x6b, 0xed, 0xae, 0x5b, 0xe5, 0x9f, 0xe0, 0xf5, 0xb7,
	0xf0, 0x4f, 0xf1, 0x7e, 0xc3, 0x57, 0x66, 0x77, 0xd6, 0x76, 0xce, 0x3d, 0xf7, 0x9a, 0x4f, 0xf6,
	0xde, 0xce, 0xef, 0xb1, 0x73, 0x7b, 0x33, 0x03, 0x6f, 0xe7, 0xda, 0x48, 0xc5, 0x3a, 0x4c, 0x0c,
	0xb9, 0x60, 0x9d, 0x91, 0x92, 0x19, 0xcb, 0xc7, 0x8a, 0xe9, 0xce, 0xe8, 0x3c, 0x39, 0x95, 0x39,
	0x1f, 0x5c, 0xf4, 0x99, 0xd2, 0x52, 0xdc, 0x4d, 0x4d, 0x9a, 0xf4, 0xc7, 0x11, 0x46, 0x18, 0x19,
	0xbc, 0x4a, 0xb0, 0x88, 0x60, 0x51, 0x75, 0x6c, 0x6b, 0xcf, 0x93, 0x3f, 0x4e, 0x8b, 0x31, 0xd3,
	0x04, 0x6d, 0xdd, 0x28, 0x2b, 0x32, 0xa5, 0xa4, 0xf2, 0x5b, 0xed, 0xf2, 0xd6, 0x39, 0xd3, 0x3a,
	0x1d, 0x32, 0xbf, 0xf9, 0xca, 0xfc, 0xa6, 0x49, 0xb9, 0x18, 0x48, 0x75, 0x9e, 0x1a, 0x2e, 0x05,
	0x05, 0x1d, 0x7d, 0xbd, 0x0b, 0xd0, 0x4f, 0x55, 0x8a, 0xbb, 0xe8, 0x24, 0xe8, 0x43, 0x33, 0x7b,
	0x84, 0xcb, 0x0c, 0x57, 0x5c, 0x1b, 0x9e, 0x25, 0x3c, 0x4f, 0x0a, 0xfc, 0x17, 0x2e, 0x1d, 0x2e,
	0xbd, 0xde, 0xb8, 0xdd, 0x8a, 0xfc, 0x39, 0xbc, 0x43, 0x6d, 0x14, 0x17, 0xc3, 0x33, 0xbb, 0x88,
	0x5f, 0x2a, 0x23, 0xbb, 0xf9, 0xc7, 0xf8, 0x1b, 0x7c, 0x00, 0xed, 0x6a, 0xc6, 0x44, 0x8c, 0x8b,
	0x22, 0xfc, 0x61, 0x0d, 0x79, 0xd7, 0xe3, 0xb0, 0x0a, 0xdb, 0xc3, 0x80, 0xe0, 0x5d, 0x00, 0xa7,
	0x45, 0x2e, 0x5e, 0xa8, 0x75, 0xb1, 0xe1, 0x9e, 0x39, 0xe9, 0xd7, 0x60, 0x67, 0x06, 0x25, 0xb9,
	0x1f, 0x49, 0x6e, 0x6b, 0x1a, 0xe4, 0x34, 0x1e, 0xc0, 0x75, 0x9e, 0x33, 0x61, 0xf8, 0x80, 0x67,
	0x2e, 0x39, 0x09, 0x11, 0x87, 0xcb, 0xf5, 0x87, 0x2e, 0x03, 0xdd, 0x43, 0x1d, 0xbc, 0x0f, 0xad,
	0x4a, 0x42, 0x32, 0xf1, 0x93, 0x3f, 0x73, 0x15, 0xd4, 0xf9, 0x79, 0x07, 0x36, 0x46, 0xee, 0x5e,
	0x60, 0xae, 0xc2, 0x6b, 0xce, 0x43, 0x7b, 0xce, 0x03, 0x17, 0x86, 0x0d, 0x99, 0x22, 0x13, 0xeb,
	0x14, 0xdd, 0xcd, 0x83, 0x5b, 0xb0, 0x3d, 0x45, 0x92, 0xd8, 0xcf, 0x24, 0xb6, 0x39, 0x09, 0x71,
	0x02, 0xc7, 0xd3, 0x30, 0x73, 0x31, 0x62, 0x56, 0x65, 0xa5, 0x5e, 0xc5, 0x53, 0x7c, 0x86, 0x08,
	0x54, 0x8a, 0x60, 0xaf, 0x4c, 0x41, 0x72, 0xbf, 0x90, 0xdc, 0xee, 0xe5, 0x58, 0x27, 0x19, 0xc3,
	0xbe, 0x8f, 0x1f, 0xaa, 0x54, 0x98, 0x24, 0xcd, 0x32, 0xbc, 0xad, 0x88, 0xd3, 0xe1, 0x6a, 0xa5,
	0xf6, 0x43, 0x29, 0x0b, 0x96, 0x0a, 0x9f, 0x66, 0xc2, 0x7e, 0x64, 0xa1, 0xc7, 0x0e, 0xd9, 0xcd,
	0x75, 0xf0, 0x21, 0x1c, 0x2c, 0xe0, 0x24, 0x33, 0xbf, 0xfa, 0x44, 0x57, 0x81, 0x9d, 0xa9, 0x33,
	0xb8, 0xe1, 0x09, 0xe8, 0xfe, 0x25, 0xf8, 0x2a, 0xd8, 0x50, 0xaa, 0x0b, 0x9b, 0x92, 0xb5, 0xfa,
	0x94, 0x34, 0x09, 0x7d, 0xe2, 0xc0, 0x27, 0x1e, 0x8b, 0xc9, 0x39, 0x81, 0x9b, 0x0b, 0x79, 0xc9,
	0xda, 0x6f, 0x64, 0xad, 0x55, 0x4d, 0x30, 0x31, 0x97, 0xb3, 0x02, 0xbf, 0xcb, 0x2a, 0x73, 0xeb,
	0xf5, 0x39, 0x6b, 0x12, 0xba, 0xca, 0xdc, 0x42, 0x5e, 0x32, 0xf7, 0xbb, 0x37, 0x57, 0x4d, 0xe0,
	0xcc, 0x7d, 0x02, 0xfb, 0x58, 0xf3, 0xc6, 0x85, 0x49, 0xb8, 0x48, 0x5c, 0x49, 0x9a, 0x56, 0x8a,
	0x8d, 0x7a, 0x6b, 0x7b, 0x84, 0xed, 0x8a, 0x7b, 0x16, 0x39, 0xad, 0x14, 0x07, 0x0b, 0x28, 0xc9,
	0xd5, 0x1f, 0xe4, 0x6a, 0xbf, 0x02, 0xeb, 0x2c, 0x65, 0x70, 0x48, 0x9f, 0xbb, 0x7d, 0xff, 0x58,
	0xe7, 0x92, 0x91, 0x62, 0x39, 0x1b, 0x60, 0xed, 0xcb, 0xfd, 0x49, 0x75, 0x08, 0xf5, 0xde, 0x0e,
	0xdc, 0x43, 0xbc, 0x21, 0xf7, 0xa5, 0xea, 0x4f, 0x19, 0x28, 0x09, 0xb6, 0x40, 0xde, 0xaa, 0x13,
	0x21, 0xb7, 0x7f, 0x92, 0xdb, 0xc3, 0x67, 0xb1, 0x39, 0xdb, 0x3d, 0xb8, 0x8e, 0x40, 0x31, 0x64,
	0x49, 0x5a, 0x14, 0x09, 0x32, 0x0a, 0x69, 0x1e, 0x61, 0x7d, 0x09, 0x1b, 0xf5, 0x5e, 0x03, 0x42,
	0x1e, 0x17, 0xc5, 0x03, 0xd5, 0x23, 0x58, 0xf0, 0x1e, 0xb4, 0x2a, 0xf9, 0xc8, 0xd6, 0x5f, 0x64,
	0xab, 0xf9, 0x34, 0xd0, 0x99, 0xb9, 0x03, 0xdb, 0x59, 0xaa, 0x59, 0xa2, 0x99, 0xd0, 0xdc, 0xf0,
	0xc7, 0x2c, 0xdc, 0xac, 0x77, 0xb1, 0x65, 0x21, 0x9f, 0x4e, 0x10, 0x41, 0x07, 0xf6, 0xca, 0x1c,
	0xa4, 0xfc, 0x37, 0x29, 0xbf, 0x58, 0x0a, 0x76, 0xa2, 0x6f, 0xc1, 0x5a, 0x26, 0xc7, 0xc2, 0xa8,
	0x8b, 0x70, 0xab, 0xb6, 0xe0, 0x4e, 0x42, 0x83, 0x23, 0xd8, 0xf4, 0x7f, 0x89, 0xff, 0x1f, 0xe2,
	0x6f, 0xf8, 0x87, 0x93, 0x5b, 0xaa, 0xd9, 0x08, 0x73, 0x6d, 0xec, 0x65, 0xf2, 0x55, 0x98, 0x6e,
	0xe9, 0x76, 0x7d, 0x69, 0x9f, 0x42, 0xbb, 0x74, 0x50, 0x77, 0x4b, 0xb1, 0xe6, 0x2c, 0xa0, 0x24,
	0x1b, 0xff, 0xfa, 0x9a, 0x53, 0x05, 0x76, 0x9e, 0x3e, 0x87, 0xd6, 0x8c, 0x60, 0xe0, 0x2e, 0x39,
	0xb6, 0x81, 0x49, 0xc7, 0xd9, 0xa9, 0xb5, 0x35, 0x3b, 0xd1, 0x7d, 0x7b, 0xfd, 0x11, 0xeb, 0x9b,
	0xce, 0x3d, 0x78, 0x79, 0x31, 0x31, 0x99, 0xfb, 0x8f, 0xcc, 0xb5, 0x17, 0x50, 0x58, 0x7f, 0x47,
	0x5f, 0x5d, 0x83, 0xf5, 0x98, 0xe9, 0x91, 0x14, 0x9a, 0x05, 0x6f, 0xc0, 0x8a, 0xfb, 0x12, 0xe7,
	0xdb, 0xbf, 0x1f, 0x63, 0x68, 0x16, 0x71, 0x1f, 0x62, 0x4c, 0x81, 0xc1, 0x17, 0xb0, 0x6b, 0x27,
	0x8d, 0xe4, 0xd2, 0xa8, 0x81, 0x5d, 0x7b, 0x19, 0xc1, 0xd1, 0x1c, 0x78, 0x7e, 0x20, 0x39, 0xc5,
	0x75, 0x77, 0xb6, 0x8e, 0x77, 0xce, 0xcb, 0x0f, 0xb0, 0x2d, 0xae, 0xf9, 0x09, 0x07, 0x1b, 0xb3,
	0x65, 0xbc, 0xf9, 0x14, 0x23, 0xcd, 0x3f, 0xa7, 0xf4, 0x1b, 0x4f, 0xc2, 0x83, 0xbb, 0xb0, 0xac,
	0xe4, 0x13, 0x6c, 0xa5, 0x16, 0x75, 0x3b, 0x7a, 0x9e, 0x59, 0x2c, 0x9a, 0xe4, 0x20, 0x8a, 0xe5,
	0x93, 0xd8, 0xc2, 0xb1, 0x40, 0x6d, 0x95, 0x2b, 0x1d, 0xab, 0x7d, 0x59, 0x0d, 0x36, 0x2b, 0x52,
	0xad, 0xef, 0x97, 0x60, 0x19, 0xc9, 0x82, 0x26, 0xac, 0x22, 0x9d, 0xad, 0xe2, 0xdf, 0xf4, 0x90,
	0x61, 0x25, 0x5e, 0xc1, 0x25, 0x16, 0xe6, 0x33, 0x08, 0x4b, 0x5d, 0xe3, 0xd2, 0xc4, 0x14, 0x7e,
	0xdb, 0xbb, 0x62, 0x37, 0x9a, 0x8d, 0x52, 0x38, 0x8d, 0x34, 0x7c, 0x61, 0xcd, 0x64, 0xce, 0xc2,
	0xef, 0x9e, 0x83, 0x0a, 0x08, 0x70, 0x82, 0xf1, 0x77, 0x7a, 0xd0, 0xe6, 0x72, 0x3e, 0x67, 0xd3,
	0xb1, 0xf7, 0xcb, 0xce, 0x15, 0x07, 0xe2, 0x87, 0xab, 0x6e, 0xf2, 0x7c, 0xf3, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x52, 0xa0, 0x4d, 0xd3, 0x4a, 0x0b, 0x00, 0x00,
}