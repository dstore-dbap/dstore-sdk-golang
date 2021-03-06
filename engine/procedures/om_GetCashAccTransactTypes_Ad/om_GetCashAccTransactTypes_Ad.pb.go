// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCashAccTransactTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCashAccTransactTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCashAccTransactTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCashAccTransactTypes_Ad

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
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	TransactionType           *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=transaction_type,json=transactionType" json:"transaction_type,omitempty"`
	AccountHolderTransactions *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=account_holder_transactions,json=accountHolderTransactions" json:"account_holder_transactions,omitempty"`
	TransactionTypeId         *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=transaction_type_id,json=transactionTypeId" json:"transaction_type_id,omitempty"`
	PositiveTransactionValues *dstore_values.BooleanValue `protobuf:"bytes,10004,opt,name=positive_transaction_values,json=positiveTransactionValues" json:"positive_transaction_values,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTransactionType() *dstore_values.StringValue {
	if m != nil {
		return m.TransactionType
	}
	return nil
}

func (m *Response_Row) GetAccountHolderTransactions() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccountHolderTransactions
	}
	return nil
}

func (m *Response_Row) GetTransactionTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TransactionTypeId
	}
	return nil
}

func (m *Response_Row) GetPositiveTransactionValues() *dstore_values.BooleanValue {
	if m != nil {
		return m.PositiveTransactionValues
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCashAccTransactTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCashAccTransactTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCashAccTransactTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetCashAccTransactTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0x6f, 0x6b, 0x14, 0x31,
	0x10, 0xc6, 0x69, 0xcf, 0xab, 0x12, 0x85, 0xd6, 0x14, 0x64, 0x7b, 0x07, 0x52, 0xea, 0x1b, 0x41,
	0xd8, 0x13, 0x7d, 0xa3, 0x20, 0x48, 0x15, 0xff, 0x14, 0x6d, 0x91, 0x70, 0x08, 0x8a, 0x18, 0xd2,
	0xcd, 0xb8, 0x0d, 0xde, 0x65, 0x8e, 0x4c, 0xae, 0xc5, 0x6f, 0x51, 0xff, 0x7c, 0x4f, 0xbf, 0x82,
	0xd9, 0x4d, 0x96, 0xee, 0xae, 0x22, 0xbd, 0x37, 0x77, 0xcc, 0x3e, 0xf3, 0xfc, 0xe6, 0x99, 0x84,
	0xb0, 0x27, 0x9a, 0x3c, 0x3a, 0x98, 0x80, 0x2d, 0x8d, 0x85, 0xc9, 0xc2, 0x61, 0x01, 0x7a, 0xe9,
	0x80, 0x26, 0x38, 0x97, 0xaf, 0xc0, 0x3f, 0x57, 0x74, 0xb2, 0x5f, 0x14, 0x53, 0xa7, 0x2c, 0xa9,
	0xc2, 0x4f, 0xbf, 0x2d, 0x80, 0xe4, 0xbe, 0xce, 0x43, 0xa3, 0x47, 0x7e, 0x2f, 0xba, 0xf3, 0xe8,
	0xce, 0xff, 0x6b, 0x19, 0x6d, 0xa7, 0x51, 0xa7, 0x6a, 0xb6, 0x04, 0x8a, 0x84, 0xd1, 0x4e, 0x77,
	0x3e, 0x38, 0x87, 0x2e, 0x49, 0xe3, 0xae, 0x34, 0x07, 0x22, 0x55, 0x42, 0x12, 0xef, 0xf4, 0x45,
	0xaf, 0x8c, 0xfd, 0x82, 0x6e, 0xae, 0xbc, 0x41, 0x1b, 0x9b, 0xf6, 0x6e, 0x30, 0xf6, 0x4e, 0x39,
	0x15, 0x44, 0x70, 0xb4, 0x77, 0x3e, 0x64, 0xd7, 0x04, 0xd0, 0x02, 0x2d, 0x01, 0xbf, 0xcf, 0x86,
	0xf5, 0xac, 0x6c, 0x6d, 0x77, 0xed, 0xee, 0xf5, 0x07, 0xa3, 0xbc, 0xbb, 0x49, 0xcc, 0xf1, 0xa2,
	0xfa, 0x15, 0xb1, 0x91, 0x7f, 0x60, 0x5b, 0xd5, 0x14, 0xd9, 0x1a, 0x93, 0xad, 0xef, 0x0e, 0x82,
	0x39, 0xef, 0x99, 0xfb, 0x61, 0x0e, 0x43, 0x7d, 0x70, 0x51, 0x8b, 0xcd, 0x79, 0xf7, 0x03, 0x7f,
	0xc4, 0xae, 0xa6, 0xed, 0xb2, 0x41, 0x4d, 0xbc, 0xfd, 0x17, 0x31, 0xee, 0x7e, 0x18, 0xff, 0x45,
	0xd3, 0xce, 0xdf, 0xb0, 0x81, 0xc3, 0xb3, 0xec, 0x4a, 0xed, 0x7a, 0x9c, 0xaf, 0x70, 0x1d, 0x79,
	0x73, 0x14, 0xb9, 0xc0, 0x33, 0x51, 0x51, 0x46, 0xbf, 0xd7, 0xd9, 0x20, 0x14, 0xfc, 0x16, 0xdb,
	0x08, 0xa5, 0x34, 0x3a, 0x3b, 0x3f, 0x0a, 0xa7, 0x33, 0x14, 0xc3, 0x50, 0x1e, 0x68, 0xfe, 0x92,
	0x6d, 0xf9, 0x44, 0x09, 0xa9, 0xa5, 0x0f, 0xa4, 0xec, 0xfb, 0x51, 0xf7, 0xfc, 0xd2, 0xe5, 0x92,
	0x77, 0xc6, 0x96, 0xef, 0xab, 0x42, 0x6c, 0xb6, 0x4c, 0xd5, 0x74, 0xfe, 0x89, 0x8d, 0x55, 0x51,
	0xe0, 0xd2, 0x7a, 0x79, 0x82, 0x33, 0x0d, 0x4e, 0xb6, 0x3a, 0x28, 0xfb, 0x11, 0x91, 0xe3, 0x1e,
	0xd2, 0x58, 0x0f, 0x25, 0xb8, 0xc8, 0xdc, 0x49, 0x80, 0xd7, 0xb5, 0x7f, 0xda, 0xb2, 0xf3, 0xb7,
	0x6c, 0xbb, 0x9f, 0xb2, 0x5a, 0xe5, 0xe7, 0x25, 0xa8, 0x37, 0x7b, 0x49, 0xc3, 0xce, 0x21, 0xeb,
	0x02, 0xc9, 0x78, 0x73, 0x0a, 0xed, 0x94, 0x32, 0xda, 0xb3, 0x5f, 0xff, 0xa6, 0x1e, 0x23, 0xce,
	0x40, 0xd9, 0x94, 0xb5, 0x01, 0xb4, 0x62, 0xd6, 0x0a, 0x3d, 0xfb, 0xcc, 0xc6, 0x06, 0x7b, 0xb7,
	0x76, 0xf1, 0x04, 0x3f, 0x3e, 0x2d, 0x91, 0xf4, 0xd7, 0x46, 0xd7, 0x2b, 0xbf, 0xd2, 0xe3, 0x8d,
	0xfa, 0x1d, 0x3c, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xc9, 0x97, 0x26, 0xe6, 0x03, 0x00,
	0x00,
}
