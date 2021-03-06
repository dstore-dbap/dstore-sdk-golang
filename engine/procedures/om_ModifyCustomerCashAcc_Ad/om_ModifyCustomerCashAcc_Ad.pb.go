// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCustomerCashAcc_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCustomerCashAcc_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCustomerCashAcc_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCustomerCashAcc_Ad

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
	PersonId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull          bool                        `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	CashAccountTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=cash_account_type_id,json=cashAccountTypeId" json:"cash_account_type_id,omitempty"`
	CashAccountTypeIdNull bool                        `protobuf:"varint,1002,opt,name=cash_account_type_id_null,json=cashAccountTypeIdNull" json:"cash_account_type_id_null,omitempty"`
	CurrencyId            *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencyIdNull        bool                        `protobuf:"varint,1003,opt,name=currency_id_null,json=currencyIdNull" json:"currency_id_null,omitempty"`
	AccountStatus         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=account_status,json=accountStatus" json:"account_status,omitempty"`
	AccountStatusNull     bool                        `protobuf:"varint,1004,opt,name=account_status_null,json=accountStatusNull" json:"account_status_null,omitempty"`
	MinAccountBalance     *dstore_values.DecimalValue `protobuf:"bytes,5,opt,name=min_account_balance,json=minAccountBalance" json:"min_account_balance,omitempty"`
	MinAccountBalanceNull bool                        `protobuf:"varint,1005,opt,name=min_account_balance_null,json=minAccountBalanceNull" json:"min_account_balance_null,omitempty"`
	DeleteAccount         *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=delete_account,json=deleteAccount" json:"delete_account,omitempty"`
	DeleteAccountNull     bool                        `protobuf:"varint,1006,opt,name=delete_account_null,json=deleteAccountNull" json:"delete_account_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetCashAccountTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CashAccountTypeId
	}
	return nil
}

func (m *Parameters) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func (m *Parameters) GetAccountStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccountStatus
	}
	return nil
}

func (m *Parameters) GetMinAccountBalance() *dstore_values.DecimalValue {
	if m != nil {
		return m.MinAccountBalance
	}
	return nil
}

func (m *Parameters) GetDeleteAccount() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteAccount
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCustomerCashAcc_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCustomerCashAcc_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCustomerCashAcc_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyCustomerCashAcc_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0xd5, 0xe6, 0x73, 0x9a, 0x6f, 0x0b, 0xa1, 0x75, 0x00, 0xb9, 0x89, 0x40, 0xa8, 0x08,
	0x89, 0x9e, 0x38, 0x08, 0x0e, 0xf8, 0x11, 0x1c, 0x34, 0x15, 0x07, 0x11, 0xa4, 0x42, 0x06, 0x21,
	0x81, 0x90, 0xac, 0xcd, 0xee, 0x34, 0x58, 0xd8, 0xbb, 0xd1, 0xae, 0x4d, 0x95, 0xbb, 0xe0, 0x7e,
	0xb8, 0x09, 0x6e, 0x83, 0xdf, 0x6b, 0x60, 0xd6, 0xbb, 0x6e, 0xea, 0x24, 0x22, 0x70, 0xd2, 0x74,
	0x3d, 0xef, 0x3b, 0xcf, 0xec, 0x68, 0x66, 0xc9, 0x23, 0xae, 0x73, 0xa9, 0xa0, 0x0f, 0x62, 0x92,
	0x08, 0xe8, 0x4f, 0x95, 0x64, 0xc0, 0x0b, 0x05, 0xba, 0x2f, 0xb3, 0x78, 0x24, 0x79, 0x72, 0x32,
	0x3b, 0x2a, 0x50, 0x92, 0x81, 0x3a, 0xa2, 0xfa, 0xfd, 0x21, 0x63, 0xf1, 0x21, 0x0f, 0x51, 0x96,
	0x4b, 0xff, 0xc0, 0x7a, 0x43, 0xeb, 0x0d, 0xff, 0x60, 0xe8, 0x76, 0x1c, 0xe6, 0x23, 0x4d, 0x0b,
	0xd0, 0xd6, 0xdf, 0xdd, 0xab, 0xb3, 0x41, 0x29, 0xa9, 0x5c, 0xa8, 0x57, 0x0f, 0x65, 0xa0, 0x35,
	0x9d, 0x80, 0x0b, 0xde, 0x5c, 0x0c, 0xe6, 0x34, 0x11, 0x27, 0x52, 0x65, 0x34, 0x4f, 0xa4, 0xb0,
	0xa2, 0xfd, 0x2f, 0x1e, 0x21, 0x2f, 0xa8, 0xa2, 0x18, 0x05, 0xa5, 0xfd, 0x07, 0xe4, 0xff, 0x29,
	0xfe, 0x4a, 0x11, 0x27, 0x3c, 0xd8, 0xb8, 0xb1, 0x71, 0x7b, 0xfb, 0x6e, 0x2f, 0x74, 0xf5, 0xbb,
	0xa2, 0x12, 0x91, 0xc3, 0x04, 0xd4, 0x6b, 0x73, 0x8a, 0x5a, 0x56, 0x3d, 0xe4, 0xfe, 0x2d, 0xd2,
	0x3e, 0x73, 0xc6, 0xa2, 0x48, 0xd3, 0xe0, 0xeb, 0x16, 0xfa, 0x5b, 0xd1, 0x85, 0x4a, 0x72, 0x8c,
	0x1f, 0xfd, 0xe7, 0xe4, 0x32, 0xc3, 0xfb, 0xc6, 0x94, 0x31, 0x59, 0x88, 0x3c, 0xce, 0x67, 0x53,
	0x30, 0xac, 0xcd, 0xf5, 0xac, 0x5d, 0x66, 0x1b, 0x65, 0x7c, 0xaf, 0xd0, 0x86, 0xd0, 0x87, 0x64,
	0x6f, 0x55, 0x36, 0xcb, 0xff, 0x66, 0xf9, 0x57, 0x96, 0x6c, 0x65, 0x21, 0x8f, 0xc9, 0x36, 0x2b,
	0x94, 0x02, 0xc1, 0x66, 0x86, 0xdf, 0x58, 0xcf, 0x27, 0x95, 0x1e, 0xc1, 0x07, 0x64, 0xe7, 0x9c,
	0xdb, 0xf2, 0xbe, 0x5b, 0x5e, 0x7b, 0x2e, 0x2b, 0x41, 0x03, 0xd2, 0xae, 0xca, 0xd3, 0x39, 0xcd,
	0x0b, 0x1d, 0xfc, 0xb7, 0x9e, 0x75, 0xd1, 0x59, 0x5e, 0x96, 0x0e, 0xbf, 0x4f, 0x3a, 0xf5, 0x1c,
	0x96, 0xf8, 0xc3, 0x12, 0x77, 0x6b, 0xe2, 0x12, 0xfa, 0x8c, 0x74, 0xb2, 0x44, 0x9c, 0xf5, 0x65,
	0x4c, 0x53, 0x2a, 0x18, 0x04, 0xde, 0x4a, 0x32, 0x07, 0x96, 0x64, 0x34, 0x75, 0x5d, 0x46, 0x9f,
	0xeb, 0xd6, 0xc0, 0xba, 0x70, 0x28, 0x82, 0x15, 0xc9, 0x6c, 0x09, 0x3f, 0x5d, 0x93, 0x97, 0x5c,
	0xd5, 0xdd, 0x39, 0xa4, 0x38, 0x5a, 0x95, 0x39, 0x68, 0xae, 0xac, 0x60, 0x2c, 0x65, 0x0a, 0x54,
	0xb8, 0xbb, 0x5b, 0x8b, 0x4b, 0x67, 0xee, 0x5e, 0xcf, 0x61, 0xc1, 0xbf, 0xdc, 0xdd, 0x6b, 0x62,
	0x03, 0xdd, 0xff, 0xbc, 0x49, 0x5a, 0x11, 0xe8, 0xa9, 0x14, 0x1a, 0xfc, 0x3b, 0xc4, 0x2b, 0x17,
	0xc6, 0x0d, 0x73, 0x37, 0xac, 0x2f, 0xa3, 0x5d, 0xa6, 0xa7, 0xe6, 0x6f, 0x64, 0x85, 0xfe, 0x1b,
	0xb2, 0x63, 0x56, 0x25, 0x3e, 0xb7, 0x2b, 0x38, 0x9d, 0x0d, 0x34, 0x87, 0x0b, 0xe6, 0xc5, 0x8d,
	0x1a, 0xe1, 0x79, 0x38, 0x3f, 0x47, 0x97, 0xb2, 0xfa, 0x07, 0x6c, 0xe4, 0x96, 0x5b, 0x51, 0x9c,
	0x37, 0x93, 0xf1, 0xfa, 0x52, 0x46, 0xbb, 0xc0, 0x23, 0xfb, 0x1b, 0x55, 0x72, 0x7f, 0x48, 0x1a,
	0x4a, 0x9e, 0xe2, 0xe4, 0x18, 0xd7, 0xfd, 0xf0, 0xaf, 0x5f, 0x94, 0xb0, 0x6a, 0x44, 0x18, 0xc9,
	0xd3, 0xc8, 0xe4, 0xe8, 0x5e, 0x23, 0x0d, 0xfc, 0xdf, 0xbf, 0x4a, 0x9a, 0x78, 0x32, 0xa3, 0xff,
	0xe9, 0x18, 0x5b, 0xe3, 0x45, 0x1e, 0x1e, 0x87, 0x7c, 0xf0, 0x8e, 0xf4, 0x12, 0xb9, 0x00, 0x98,
	0x3f, 0x77, 0x6f, 0x9f, 0x4c, 0xa4, 0xe6, 0x1f, 0xaa, 0x38, 0xff, 0xc7, 0x17, 0x71, 0xdc, 0x2c,
	0x5f, 0x9d, 0x7b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x8f, 0x96, 0x1b, 0x50, 0x05, 0x00,
	0x00,
}
