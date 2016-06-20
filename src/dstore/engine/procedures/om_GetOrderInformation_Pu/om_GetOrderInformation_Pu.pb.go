// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderInformation_Pu.proto
// DO NOT EDIT!

/*
Package om_GetOrderInformation_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderInformation_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderInformation_Pu

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
	OrderId                        *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	OrderIdNull                    bool                        `protobuf:"varint,1001,opt,name=order_id_null,json=orderIdNull" json:"order_id_null,omitempty"`
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1003,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	InformationTypeId              *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull          bool                        `protobuf:"varint,1004,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
	OrderCode                      *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=order_code,json=orderCode" json:"order_code,omitempty"`
	OrderCodeNull                  bool                        `protobuf:"varint,1005,opt,name=order_code_null,json=orderCodeNull" json:"order_code_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

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

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Parameters) GetOrderCode() *dstore_values.StringValue {
	if m != nil {
		return m.OrderCode
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
	RowId                  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	InformationTypeId      *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	PersonCharacteristicId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	OrderId                *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	InformationType        *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	Information            *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=information" json:"information,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Response_Row) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *Response_Row) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func (m *Response_Row) GetInformation() *dstore_values.StringValue {
	if m != nil {
		return m.Information
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderInformation_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderInformation_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderInformation_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x49, 0xd3, 0x86, 0x09, 0x25, 0x61, 0x23, 0x2a, 0x93, 0xa0, 0x0a, 0xda, 0x43, 0x39,
	0x39, 0xa8, 0x08, 0x28, 0x07, 0x0e, 0x50, 0x41, 0x15, 0xa0, 0xa1, 0x58, 0xa8, 0x12, 0x08, 0x29,
	0x32, 0xf6, 0x36, 0xac, 0x94, 0x78, 0xa3, 0x5d, 0x87, 0x8a, 0x33, 0x2f, 0xc0, 0xef, 0x9d, 0xd7,
	0xe3, 0xef, 0xc0, 0x1b, 0x30, 0xeb, 0xd9, 0x24, 0xb6, 0xdb, 0xe2, 0x96, 0x4b, 0xab, 0xf5, 0x7e,
	0x3f, 0xb3, 0x3b, 0xdf, 0x4e, 0x60, 0x33, 0xd4, 0xb1, 0x54, 0xbc, 0xcd, 0xa3, 0xbe, 0x88, 0x78,
	0x7b, 0xa4, 0x64, 0xc0, 0xc3, 0xb1, 0xe2, 0xba, 0x2d, 0x87, 0xbd, 0x6d, 0x1e, 0x3f, 0x55, 0x21,
	0x57, 0x9d, 0x68, 0x5f, 0xaa, 0xa1, 0x1f, 0x0b, 0x19, 0xf5, 0x76, 0xc7, 0x2e, 0x82, 0x62, 0xc9,
	0xd6, 0x89, 0xe9, 0x12, 0xd3, 0x3d, 0x16, 0xde, 0x6c, 0x58, 0x8b, 0xb7, 0xfe, 0x60, 0xcc, 0x35,
	0xb1, 0x9b, 0x97, 0xb2, 0xbe, 0x5c, 0x29, 0xa9, 0xec, 0x56, 0x2b, 0xbb, 0x35, 0xe4, 0x5a, 0xfb,
	0x7d, 0x6e, 0x37, 0xd7, 0xf2, 0x9b, 0xb1, 0x2f, 0x66, 0x76, 0x04, 0x5a, 0xfd, 0x33, 0x0f, 0xb0,
	0xeb, 0x2b, 0x1f, 0x77, 0xb9, 0xd2, 0xec, 0x16, 0x54, 0xa4, 0xa9, 0xab, 0x27, 0x42, 0x67, 0xee,
	0xca, 0xdc, 0xb5, 0xea, 0x46, 0xcb, 0xb5, 0xc5, 0xdb, 0x9a, 0x44, 0x14, 0xf3, 0x3e, 0x57, 0x7b,
	0x66, 0xe5, 0x2d, 0x26, 0xe0, 0x4e, 0xc8, 0xd6, 0x60, 0x69, 0xc2, 0xeb, 0x45, 0xe3, 0xc1, 0xc0,
	0xf9, 0xbe, 0x88, 0xec, 0x8a, 0x57, 0xb5, 0x80, 0x2e, 0x7e, 0x63, 0xaf, 0xe0, 0xf2, 0x08, 0x4d,
	0xf0, 0xa8, 0x22, 0xe4, 0x51, 0x2c, 0xf6, 0x45, 0x40, 0x27, 0x27, 0x69, 0xe7, 0x4c, 0x62, 0xd8,
	0xcc, 0x19, 0xea, 0x58, 0x89, 0xa8, 0x4f, 0x7e, 0x4d, 0xe2, 0x77, 0x32, 0xf4, 0x64, 0x4b, 0xb3,
	0x47, 0x70, 0xf5, 0x5f, 0xea, 0x54, 0xd6, 0x0f, 0x2a, 0x6b, 0xe5, 0x78, 0x9d, 0xa4, 0xd2, 0x7b,
	0x70, 0xde, 0x6a, 0xc5, 0xef, 0x46, 0xdc, 0x5c, 0x46, 0xa9, 0xf8, 0x32, 0xce, 0x11, 0xe5, 0x39,
	0x32, 0xf0, 0x46, 0x5c, 0x68, 0x64, 0x25, 0xa8, 0x80, 0x9f, 0x54, 0x40, 0x3d, 0x8d, 0x4d, 0x2c,
	0x1f, 0x43, 0x23, 0xd5, 0x9d, 0xa9, 0xef, 0x7c, 0xb1, 0xef, 0x85, 0x14, 0xcf, 0x9a, 0x6f, 0x82,
	0x73, 0x84, 0x18, 0x55, 0xf0, 0x8b, 0x2a, 0xb8, 0x78, 0x88, 0x95, 0x94, 0x71, 0x07, 0x80, 0x1a,
	0x19, 0xc8, 0x90, 0x3b, 0xe5, 0xc2, 0x8e, 0x9c, 0x4d, 0xd0, 0x5b, 0x08, 0x66, 0xeb, 0x50, 0x9b,
	0x51, 0xc9, 0xeb, 0x37, 0x79, 0x2d, 0x4d, 0x41, 0xc6, 0x63, 0xf5, 0x5b, 0x19, 0x2a, 0x1e, 0xd7,
	0x23, 0x19, 0x69, 0xce, 0xae, 0x43, 0x39, 0x49, 0xb4, 0x8d, 0xdb, 0xd4, 0xcb, 0xbe, 0x15, 0x4a,
	0xfb, 0x03, 0xf3, 0xd7, 0x23, 0x20, 0x7b, 0x01, 0x75, 0x93, 0xe5, 0x5e, 0xea, 0x00, 0x18, 0x9d,
	0x12, 0x92, 0xdd, 0x1c, 0x39, 0x1f, 0xf9, 0x1d, 0x5c, 0xa7, 0x5e, 0x9c, 0x57, 0x1b, 0x66, 0x3f,
	0xe0, 0xbd, 0x2d, 0xda, 0x37, 0x84, 0x0d, 0x37, 0x8a, 0x2b, 0x87, 0x14, 0xe9, 0x85, 0xed, 0xd0,
	0x7f, 0x6f, 0x02, 0x67, 0xdb, 0x50, 0x52, 0xf2, 0x00, 0xdb, 0x65, 0x58, 0x37, 0xdd, 0x13, 0x3e,
	0x78, 0x77, 0x72, 0x0d, 0xae, 0x27, 0x0f, 0x3c, 0xa3, 0xd0, 0x7c, 0x5f, 0x82, 0x12, 0x2e, 0xd8,
	0x32, 0x2c, 0xe0, 0xd2, 0x44, 0xe0, 0x43, 0x17, 0x6f, 0xa6, 0xec, 0x95, 0x71, 0x89, 0xad, 0x7d,
	0x72, 0x74, 0x4e, 0x3e, 0x76, 0xff, 0x2b, 0x28, 0x7b, 0xe0, 0xd8, 0x94, 0x06, 0x6f, 0x70, 0x0a,
	0x04, 0x38, 0x04, 0x84, 0x8e, 0x45, 0x60, 0x24, 0x3f, 0x9d, 0x40, 0x72, 0x99, 0xd8, 0x5b, 0x19,
	0x32, 0xea, 0xde, 0x4e, 0xcd, 0x91, 0xcf, 0xdd, 0x53, 0x0c, 0x92, 0x87, 0x50, 0xcf, 0x1f, 0xcf,
	0xf9, 0xd2, 0x2d, 0x8c, 0x61, 0x2d, 0x77, 0x34, 0x76, 0x17, 0xaa, 0xe9, 0x7c, 0x7c, 0x2d, 0x96,
	0x48, 0xe3, 0xef, 0x3f, 0x83, 0x96, 0x90, 0xb9, 0x2e, 0xce, 0x06, 0xfe, 0xcb, 0x8d, 0xd3, 0xff,
	0x14, 0xbc, 0x5e, 0x48, 0x06, 0xee, 0x8d, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x19, 0xa0,
	0xb9, 0x47, 0x06, 0x00, 0x00,
}