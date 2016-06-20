// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyOrderStates_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyOrderStates_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyOrderStates_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyOrderStates_Ad

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
	OrderStateId                  *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=order_state_id,json=orderStateId" json:"order_state_id,omitempty"`
	OrderStateIdNull              bool                        `protobuf:"varint,1001,opt,name=order_state_id_null,json=orderStateIdNull" json:"order_state_id_null,omitempty"`
	OrderState                    *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=order_state,json=orderState" json:"order_state,omitempty"`
	OrderStateNull                bool                        `protobuf:"varint,1002,opt,name=order_state_null,json=orderStateNull" json:"order_state_null,omitempty"`
	PublicDescription             *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=public_description,json=publicDescription" json:"public_description,omitempty"`
	PublicDescriptionNull         bool                        `protobuf:"varint,1003,opt,name=public_description_null,json=publicDescriptionNull" json:"public_description_null,omitempty"`
	InsertUpdateOrDelete          *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=insert_update_or_delete,json=insertUpdateOrDelete" json:"insert_update_or_delete,omitempty"`
	InsertUpdateOrDeleteNull      bool                        `protobuf:"varint,1004,opt,name=insert_update_or_delete_null,json=insertUpdateOrDeleteNull" json:"insert_update_or_delete_null,omitempty"`
	Active                        *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=active" json:"active,omitempty"`
	ActiveNull                    bool                        `protobuf:"varint,1005,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
	ChangeOrderStateTriggerId     *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=change_order_state_trigger_id,json=changeOrderStateTriggerId" json:"change_order_state_trigger_id,omitempty"`
	ChangeOrderStateTriggerIdNull bool                        `protobuf:"varint,1006,opt,name=change_order_state_trigger_id_null,json=changeOrderStateTriggerIdNull" json:"change_order_state_trigger_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderStateId
	}
	return nil
}

func (m *Parameters) GetOrderState() *dstore_values.StringValue {
	if m != nil {
		return m.OrderState
	}
	return nil
}

func (m *Parameters) GetPublicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PublicDescription
	}
	return nil
}

func (m *Parameters) GetInsertUpdateOrDelete() *dstore_values.IntegerValue {
	if m != nil {
		return m.InsertUpdateOrDelete
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Parameters) GetChangeOrderStateTriggerId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ChangeOrderStateTriggerId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyOrderStates_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyOrderStates_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyOrderStates_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyOrderStates_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x4b, 0x6f, 0x13, 0x31,
	0x10, 0x56, 0x1b, 0x92, 0x56, 0x53, 0x54, 0x82, 0x0b, 0x74, 0x9b, 0x10, 0x54, 0x05, 0x21, 0xc1,
	0x65, 0x83, 0x88, 0xc4, 0x43, 0x1c, 0x50, 0x51, 0x7b, 0x08, 0x52, 0x5a, 0xb4, 0x3c, 0x04, 0x48,
	0x68, 0xb5, 0x89, 0xdd, 0xc5, 0x62, 0x63, 0x47, 0xb6, 0xd3, 0x8a, 0x2b, 0xbf, 0x80, 0x7f, 0xc3,
	0x6f, 0xe2, 0xf9, 0x1b, 0xb0, 0x3d, 0x9b, 0x6e, 0x1e, 0x0d, 0x81, 0x4b, 0xb2, 0xb3, 0x33, 0xdf,
	0xc3, 0xb3, 0x33, 0x86, 0xfb, 0x54, 0x1b, 0xa9, 0x58, 0x8b, 0x89, 0x94, 0x0b, 0xd6, 0x1a, 0x2a,
	0xd9, 0x67, 0x74, 0xa4, 0x98, 0x6e, 0xc9, 0x41, 0xdc, 0x95, 0x94, 0x1f, 0x7f, 0x3a, 0x52, 0x94,
	0xa9, 0x17, 0x26, 0x31, 0x4c, 0xc7, 0x7b, 0x34, 0xb4, 0x25, 0x46, 0x92, 0x5b, 0x88, 0x0b, 0x11,
	0x17, 0x2e, 0x28, 0xae, 0x6d, 0xe5, 0xf4, 0x27, 0x49, 0x36, 0x62, 0x1a, 0xb1, 0xb5, 0x9d, 0x69,
	0x4d, 0xa6, 0x94, 0x54, 0x79, 0xaa, 0x3e, 0x9d, 0x1a, 0x30, 0xad, 0x93, 0x94, 0xe5, 0xc9, 0x9b,
	0xb3, 0x49, 0x93, 0x70, 0x71, 0x2c, 0xd5, 0x20, 0x31, 0x5c, 0x0a, 0x2c, 0x6a, 0x7e, 0xae, 0x00,
	0x3c, 0x4f, 0x54, 0x62, 0xb3, 0x4c, 0x69, 0xb2, 0x07, 0x9b, 0xd2, 0x59, 0x8a, 0xb5, 0xf3, 0x14,
	0x73, 0x1a, 0xac, 0xec, 0xae, 0xdc, 0xde, 0xb8, 0x57, 0x0f, 0xf3, 0x03, 0xe4, 0xce, 0xb8, 0x30,
	0x2c, 0x65, 0xea, 0xb5, 0x8b, 0xa2, 0x8b, 0xf2, 0xec, 0x14, 0x1d, 0x4a, 0x42, 0xd8, 0x9a, 0xa6,
	0x88, 0xc5, 0x28, 0xcb, 0x82, 0x6f, 0x6b, 0x96, 0x68, 0x3d, 0xaa, 0x4e, 0xd6, 0x1e, 0xda, 0x04,
	0x79, 0x0c, 0x1b, 0x13, 0xf5, 0xc1, 0xaa, 0xd7, 0xab, 0xcd, 0xe8, 0x69, 0xa3, 0xb8, 0x48, 0x51,
	0x0e, 0x0a, 0x0a, 0x72, 0x07, 0xaa, 0x93, 0x62, 0x5e, 0xe9, 0x3b, 0x2a, 0x6d, 0x16, 0x65, 0x5e,
	0xa7, 0x03, 0x64, 0x38, 0xea, 0x65, 0xbc, 0x1f, 0x53, 0xa6, 0xfb, 0x8a, 0x0f, 0x5d, 0x17, 0x82,
	0xd2, 0x52, 0xb9, 0xcb, 0x88, 0xda, 0x2f, 0x40, 0xe4, 0x01, 0x6c, 0xcf, 0x53, 0xa1, 0xf8, 0x0f,
	0x14, 0xbf, 0x3a, 0x07, 0xf2, 0x1e, 0x22, 0xd8, 0xe6, 0x42, 0x33, 0x65, 0xe2, 0xd1, 0x90, 0x3a,
	0xc3, 0x52, 0x59, 0x8a, 0xcc, 0xf6, 0x3e, 0xb8, 0xb0, 0xbc, 0xcf, 0x57, 0x10, 0xfb, 0xca, 0x43,
	0x8f, 0xd4, 0xbe, 0x07, 0x92, 0x27, 0x70, 0x7d, 0x01, 0x27, 0x3a, 0xfa, 0x89, 0x8e, 0x82, 0xf3,
	0xc0, 0xde, 0x54, 0x1b, 0x2a, 0x49, 0xdf, 0xf0, 0x13, 0x16, 0x94, 0xcf, 0xf5, 0xd0, 0x93, 0x32,
	0x63, 0x89, 0x40, 0x0f, 0x79, 0x29, 0xd9, 0x85, 0x0d, 0x7c, 0x42, 0x91, 0x5f, 0x28, 0x02, 0xf8,
	0xce, 0xd3, 0xbe, 0x87, 0x46, 0xff, 0x43, 0x22, 0x52, 0x67, 0xa8, 0xf8, 0x42, 0xb6, 0xb1, 0xa9,
	0x3d, 0x90, 0x9b, 0xac, 0xca, 0xf2, 0x13, 0xef, 0x20, 0x43, 0xb1, 0x25, 0x2f, 0x11, 0x6e, 0xc7,
	0xec, 0x19, 0x34, 0xff, 0x4a, 0x8f, 0xbe, 0x7e, 0xa3, 0xaf, 0xc6, 0x42, 0x1e, 0x67, 0xb5, 0xf9,
	0x75, 0x15, 0xd6, 0x23, 0xa6, 0x87, 0xd2, 0xb6, 0x88, 0xdc, 0x85, 0xb2, 0x5f, 0xb1, 0x7c, 0xf2,
	0xcf, 0x46, 0x23, 0x5f, 0x5d, 0x5c, 0xbf, 0x03, 0xf7, 0x1b, 0x61, 0x21, 0x79, 0x0b, 0x55, 0xb7,
	0x5c, 0xf1, 0xc4, 0x76, 0xd9, 0x31, 0x2e, 0x59, 0x70, 0x38, 0x03, 0x9e, 0xdd, 0xc1, 0xae, 0x8d,
	0x3b, 0x45, 0x1c, 0x5d, 0x1a, 0x4c, 0xbf, 0x20, 0x0f, 0x61, 0x2d, 0x5f, 0x6a, 0x3b, 0xa9, 0x8e,
	0xf1, 0xc6, 0x1c, 0x23, 0xae, 0x7c, 0x17, 0xff, 0xa3, 0x71, 0x39, 0x39, 0x80, 0x92, 0x92, 0xa7,
	0x76, 0xac, 0x1c, 0xaa, 0x1d, 0xfe, 0xd3, 0xfd, 0x13, 0x8e, 0x9b, 0x10, 0x46, 0xf2, 0x34, 0x72,
	0xf8, 0x5a, 0x03, 0x4a, 0xf6, 0x99, 0x5c, 0x83, 0x8a, 0x8d, 0xdc, 0x57, 0xfb, 0x72, 0x68, 0xdb,
	0x52, 0x8e, 0xca, 0x36, 0xec, 0xd0, 0xa7, 0x6f, 0xa0, 0xce, 0xe5, 0x0c, 0x79, 0x71, 0x29, 0xbe,
	0x7b, 0x94, 0x4a, 0x4d, 0x3f, 0x8e, 0xf3, 0xf4, 0x3f, 0xee, 0xcd, 0x5e, 0xc5, 0xdf, 0x4f, 0xed,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x1c, 0x1a, 0xe2, 0x72, 0x05, 0x00, 0x00,
}