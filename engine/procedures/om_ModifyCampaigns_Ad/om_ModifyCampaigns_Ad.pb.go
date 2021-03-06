// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampaigns_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampaigns_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampaigns_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampaigns_Ad

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
	CampaignName            *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=campaign_name,json=campaignName" json:"campaign_name,omitempty"`
	CampaignNameNull        bool                        `protobuf:"varint,1001,opt,name=campaign_name_null,json=campaignNameNull" json:"campaign_name_null,omitempty"`
	CampaignDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=campaign_description,json=campaignDescription" json:"campaign_description,omitempty"`
	CampaignDescriptionNull bool                        `protobuf:"varint,1002,opt,name=campaign_description_null,json=campaignDescriptionNull" json:"campaign_description_null,omitempty"`
	CampaignTypeId          *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignTypeIdNull      bool                        `protobuf:"varint,1003,opt,name=campaign_type_id_null,json=campaignTypeIdNull" json:"campaign_type_id_null,omitempty"`
	Active                  *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=active" json:"active,omitempty"`
	ActiveNull              bool                        `protobuf:"varint,1004,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
	DeleteCampaign          *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_campaign,json=deleteCampaign" json:"delete_campaign,omitempty"`
	DeleteCampaignNull      bool                        `protobuf:"varint,1005,opt,name=delete_campaign_null,json=deleteCampaignNull" json:"delete_campaign_null,omitempty"`
	ForceDelete             *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=force_delete,json=forceDelete" json:"force_delete,omitempty"`
	ForceDeleteNull         bool                        `protobuf:"varint,1006,opt,name=force_delete_null,json=forceDeleteNull" json:"force_delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCampaignName() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignName
	}
	return nil
}

func (m *Parameters) GetCampaignDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignDescription
	}
	return nil
}

func (m *Parameters) GetCampaignTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignTypeId
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Parameters) GetDeleteCampaign() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCampaign
	}
	return nil
}

func (m *Parameters) GetForceDelete() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForceDelete
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	CampaignId      *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
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

func (m *Response) GetCampaignId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampaigns_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampaigns_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampaigns_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyCampaigns_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x8d, 0x49, 0xcb, 0x49, 0xb5, 0x75, 0x5a, 0x75, 0x9b, 0xa0, 0x94, 0xf6, 0x46, 0x10,
	0x37, 0xb6, 0x15, 0x14, 0x14, 0x45, 0x9b, 0x5e, 0xe4, 0x22, 0x41, 0x16, 0x29, 0xe8, 0xcd, 0xb2,
	0xcd, 0x9c, 0x2e, 0x83, 0xc9, 0xcc, 0x32, 0xb3, 0x69, 0xe9, 0x5b, 0xf8, 0x48, 0xbe, 0x8e, 0xbf,
	0x37, 0xbe, 0x80, 0xb3, 0x7b, 0x66, 0x93, 0xec, 0x1a, 0x48, 0x6f, 0x92, 0xcc, 0x9c, 0xef, 0x6f,
	0x32, 0xe7, 0x0c, 0x3c, 0xe7, 0x26, 0x55, 0x1a, 0x3b, 0x28, 0x63, 0x21, 0xb1, 0x93, 0x68, 0x35,
	0x44, 0x3e, 0xd1, 0x68, 0x3a, 0x6a, 0x1c, 0xf6, 0x15, 0x17, 0x17, 0xd7, 0x27, 0xd1, 0x38, 0x89,
	0x44, 0x2c, 0x4d, 0xf8, 0x8e, 0xfb, 0x16, 0x90, 0x2a, 0x76, 0x40, 0x2c, 0x9f, 0x58, 0xfe, 0x42,
	0x68, 0x6b, 0xdb, 0x49, 0x5f, 0x46, 0xa3, 0x09, 0x1a, 0x62, 0xb6, 0x76, 0xcb, 0x7e, 0xa8, 0xb5,
	0xd2, 0xae, 0xd4, 0x2e, 0x97, 0xc6, 0x68, 0x4c, 0x14, 0xa3, 0x2b, 0x1e, 0x54, 0x8b, 0x69, 0x24,
	0xe4, 0x85, 0xd2, 0xe3, 0x28, 0x15, 0x4a, 0x12, 0x68, 0xff, 0x5b, 0x1d, 0xe0, 0x43, 0xa4, 0x23,
	0x5b, 0x45, 0x6d, 0xd8, 0x5b, 0xb8, 0x3d, 0x74, 0x81, 0x42, 0x69, 0x37, 0xbd, 0x95, 0xbd, 0x95,
	0xc7, 0xcd, 0xa3, 0x96, 0xef, 0xd2, 0xbb, 0x60, 0x26, 0xd5, 0x42, 0xc6, 0x67, 0xd9, 0x22, 0xd8,
	0x28, 0x08, 0x03, 0x8b, 0x67, 0x4f, 0x81, 0x95, 0x04, 0x42, 0x39, 0x19, 0x8d, 0xbc, 0xef, 0x6b,
	0x56, 0x66, 0x3d, 0xd8, 0x9a, 0x87, 0x0e, 0x6c, 0x81, 0xf5, 0x61, 0x67, 0x0a, 0xe7, 0x68, 0x86,
	0x5a, 0x24, 0x59, 0x38, 0x6f, 0x75, 0xa9, 0xed, 0x76, 0xc1, 0xeb, 0xce, 0x68, 0xec, 0x15, 0xec,
	0x2e, 0x92, 0xa3, 0x10, 0x3f, 0x28, 0xc4, 0x83, 0x05, 0xc4, 0x3c, 0xcb, 0x29, 0x4c, 0xf3, 0x85,
	0xe9, 0x75, 0x82, 0xa1, 0xe0, 0x5e, 0x2d, 0xcf, 0xd1, 0xae, 0xe4, 0x10, 0x32, 0xc5, 0x18, 0x35,
	0x05, 0xb9, 0x53, 0x90, 0x3e, 0x5a, 0x4e, 0x8f, 0xb3, 0x23, 0xb8, 0x57, 0x95, 0x21, 0xff, 0x9f,
	0xe4, 0xcf, 0xca, 0xf8, 0xdc, 0xfa, 0x18, 0x1a, 0xd1, 0x30, 0x15, 0x97, 0xe8, 0xdd, 0x5a, 0x6e,
	0xe8, 0xa0, 0x6c, 0x0f, 0x9a, 0xf4, 0x8b, 0xe4, 0x7f, 0x91, 0x3c, 0xd0, 0x5e, 0x2e, 0xdb, 0x85,
	0x4d, 0x8e, 0x23, 0x7b, 0xb3, 0x61, 0xe1, 0xe9, 0xd5, 0x17, 0xea, 0x9f, 0x2b, 0x35, 0xc2, 0x48,
	0xba, 0x03, 0x11, 0xa7, 0x68, 0x4c, 0x76, 0x08, 0x3b, 0x15, 0x15, 0x32, 0xfc, 0xed, 0xce, 0x53,
	0x86, 0xe7, 0xc6, 0x6f, 0x60, 0xc3, 0x36, 0xda, 0x10, 0x43, 0xaa, 0x79, 0x8d, 0xe5, 0xa7, 0x6a,
	0xe6, 0x84, 0x6e, 0x8e, 0x67, 0x4f, 0xe0, 0xee, 0x3c, 0x9f, 0xfc, 0xfe, 0x90, 0xdf, 0xe6, 0x1c,
	0x30, 0x33, 0xdb, 0xff, 0xbb, 0x0a, 0xeb, 0x01, 0x9a, 0x44, 0x49, 0x83, 0xec, 0x19, 0xd4, 0xf3,
	0x01, 0xa9, 0x36, 0xae, 0x1b, 0x3b, 0x1a, 0x9e, 0xd3, 0xec, 0x33, 0x20, 0x20, 0xfb, 0x04, 0x5b,
	0xd9, 0x68, 0x84, 0x73, 0xb3, 0x61, 0xdb, 0xaf, 0x66, 0xc9, 0x7e, 0x85, 0x5c, 0x9d, 0xa0, 0xbe,
	0x5d, 0xf7, 0x66, 0xeb, 0x60, 0x73, 0x5c, 0xde, 0x60, 0x2f, 0x61, 0xcd, 0x8d, 0xa4, 0x6d, 0xa4,
	0x4c, 0xf1, 0xd1, 0x7f, 0x8a, 0x34, 0xb0, 0x7d, 0xfa, 0x0e, 0x0a, 0x38, 0x3b, 0x81, 0x9a, 0x56,
	0x57, 0xb6, 0x1b, 0x32, 0xd6, 0xa1, 0x7f, 0x83, 0xb7, 0xc3, 0x2f, 0xfe, 0x02, 0x3f, 0x50, 0x57,
	0x41, 0xc6, 0x66, 0xaf, 0xa1, 0x39, 0xbd, 0x31, 0xdb, 0xcb, 0xb8, 0xfc, 0x12, 0xa0, 0xc0, 0xf7,
	0x78, 0xeb, 0x21, 0xd4, 0xac, 0x12, 0xbb, 0x0f, 0x0d, 0xab, 0x95, 0xf1, 0xbf, 0x0e, 0xac, 0x40,
	0x3d, 0xa8, 0xdb, 0x65, 0x8f, 0xbf, 0x3f, 0x83, 0xb6, 0x50, 0x95, 0x60, 0xb3, 0xa7, 0xf0, 0xf3,
	0x8b, 0x58, 0x19, 0xfe, 0xa5, 0xa8, 0xf3, 0x1b, 0xbf, 0x96, 0xe7, 0x8d, 0xfc, 0x5d, 0x3a, 0xfe,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xe5, 0x9d, 0x0d, 0x66, 0x05, 0x00, 0x00,
}
