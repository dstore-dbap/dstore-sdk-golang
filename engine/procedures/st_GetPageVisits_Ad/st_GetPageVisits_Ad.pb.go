// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetPageVisits_Ad.proto
// DO NOT EDIT!

/*
Package st_GetPageVisits_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetPageVisits_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetPageVisits_Ad

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
	ComputeSum                 *dstore_values.BooleanValue   `protobuf:"bytes,1,opt,name=compute_sum,json=computeSum" json:"compute_sum,omitempty"`
	ComputeSumNull             bool                          `protobuf:"varint,1001,opt,name=compute_sum_null,json=computeSumNull" json:"compute_sum_null,omitempty"`
	Day                        *dstore_values.TimestampValue `protobuf:"bytes,2,opt,name=day" json:"day,omitempty"`
	DayNull                    bool                          `protobuf:"varint,1002,opt,name=day_null,json=dayNull" json:"day_null,omitempty"`
	PageNo                     *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=page_no,json=pageNo" json:"page_no,omitempty"`
	PageNoNull                 bool                          `protobuf:"varint,1003,opt,name=page_no_null,json=pageNoNull" json:"page_no_null,omitempty"`
	FromDay                    *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=from_day,json=fromDay" json:"from_day,omitempty"`
	FromDayNull                bool                          `protobuf:"varint,1004,opt,name=from_day_null,json=fromDayNull" json:"from_day_null,omitempty"`
	ToDay                      *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=to_day,json=toDay" json:"to_day,omitempty"`
	ToDayNull                  bool                          `protobuf:"varint,1005,opt,name=to_day_null,json=toDayNull" json:"to_day_null,omitempty"`
	PageCategoryId             *dstore_values.IntegerValue   `protobuf:"bytes,6,opt,name=page_category_id,json=pageCategoryId" json:"page_category_id,omitempty"`
	PageCategoryIdNull         bool                          `protobuf:"varint,1006,opt,name=page_category_id_null,json=pageCategoryIdNull" json:"page_category_id_null,omitempty"`
	GroupByCategory            *dstore_values.BooleanValue   `protobuf:"bytes,7,opt,name=group_by_category,json=groupByCategory" json:"group_by_category,omitempty"`
	GroupByCategoryNull        bool                          `protobuf:"varint,1007,opt,name=group_by_category_null,json=groupByCategoryNull" json:"group_by_category_null,omitempty"`
	OrderByDay                 *dstore_values.BooleanValue   `protobuf:"bytes,8,opt,name=order_by_day,json=orderByDay" json:"order_by_day,omitempty"`
	OrderByDayNull             bool                          `protobuf:"varint,1008,opt,name=order_by_day_null,json=orderByDayNull" json:"order_by_day_null,omitempty"`
	GetSumOfAllPagesPerDay     *dstore_values.BooleanValue   `protobuf:"bytes,9,opt,name=get_sum_of_all_pages_per_day,json=getSumOfAllPagesPerDay" json:"get_sum_of_all_pages_per_day,omitempty"`
	GetSumOfAllPagesPerDayNull bool                          `protobuf:"varint,1009,opt,name=get_sum_of_all_pages_per_day_null,json=getSumOfAllPagesPerDayNull" json:"get_sum_of_all_pages_per_day_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetComputeSum() *dstore_values.BooleanValue {
	if m != nil {
		return m.ComputeSum
	}
	return nil
}

func (m *Parameters) GetDay() *dstore_values.TimestampValue {
	if m != nil {
		return m.Day
	}
	return nil
}

func (m *Parameters) GetPageNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageNo
	}
	return nil
}

func (m *Parameters) GetFromDay() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDay
	}
	return nil
}

func (m *Parameters) GetToDay() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDay
	}
	return nil
}

func (m *Parameters) GetPageCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageCategoryId
	}
	return nil
}

func (m *Parameters) GetGroupByCategory() *dstore_values.BooleanValue {
	if m != nil {
		return m.GroupByCategory
	}
	return nil
}

func (m *Parameters) GetOrderByDay() *dstore_values.BooleanValue {
	if m != nil {
		return m.OrderByDay
	}
	return nil
}

func (m *Parameters) GetGetSumOfAllPagesPerDay() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetSumOfAllPagesPerDay
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
	RowId                   int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Counter                 *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=counter" json:"counter,omitempty"`
	PageNo                  *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=page_no,json=pageNo" json:"page_no,omitempty"`
	Page                    *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=page" json:"page,omitempty"`
	Day                     *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=day" json:"day,omitempty"`
	PageCategoryId          *dstore_values.IntegerValue `protobuf:"bytes,20002,opt,name=page_category_id,json=pageCategoryId" json:"page_category_id,omitempty"`
	PageCategoryDescription *dstore_values.StringValue  `protobuf:"bytes,20003,opt,name=page_category_description,json=pageCategoryDescription" json:"page_category_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCounter() *dstore_values.IntegerValue {
	if m != nil {
		return m.Counter
	}
	return nil
}

func (m *Response_Row) GetPageNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageNo
	}
	return nil
}

func (m *Response_Row) GetPage() *dstore_values.StringValue {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *Response_Row) GetDay() *dstore_values.StringValue {
	if m != nil {
		return m.Day
	}
	return nil
}

func (m *Response_Row) GetPageCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageCategoryId
	}
	return nil
}

func (m *Response_Row) GetPageCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PageCategoryDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetPageVisits_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetPageVisits_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetPageVisits_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/st_GetPageVisits_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x4b, 0x6f, 0xd3, 0x4a,
	0x14, 0x56, 0x6e, 0x9a, 0x38, 0x3d, 0xe9, 0xed, 0xc3, 0xd5, 0xed, 0x75, 0xd3, 0x7b, 0x81, 0xb6,
	0x1b, 0x40, 0xc2, 0x41, 0x85, 0x22, 0x16, 0xb0, 0x68, 0x69, 0xa9, 0xba, 0x68, 0xa8, 0x0c, 0xaa,
	0x80, 0x8d, 0xe5, 0xc6, 0x93, 0xc8, 0xc2, 0xf6, 0x58, 0x33, 0x36, 0x55, 0x7e, 0x05, 0xcf, 0x65,
	0x57, 0x20, 0xf1, 0x9f, 0xf8, 0x09, 0xbc, 0xe1, 0x1f, 0x70, 0x66, 0x8e, 0x43, 0x1e, 0x8d, 0x48,
	0xd8, 0x34, 0x1d, 0x9f, 0xef, 0x71, 0x66, 0x3c, 0xe7, 0x33, 0x6c, 0xf8, 0x32, 0xe5, 0x82, 0xd5,
	0x59, 0xdc, 0x0e, 0x62, 0x56, 0x4f, 0x04, 0x6f, 0x32, 0x3f, 0x13, 0x4c, 0xd6, 0x65, 0xea, 0xee,
	0xb1, 0xf4, 0xd0, 0x6b, 0xb3, 0xa3, 0x40, 0x06, 0xa9, 0x74, 0xb7, 0x7c, 0x1b, 0xcb, 0x29, 0x37,
	0x57, 0x89, 0x63, 0x13, 0xc7, 0x1e, 0x01, 0xac, 0x2d, 0xe6, 0xb2, 0x4f, 0xbd, 0x30, 0x63, 0x92,
	0x78, 0xb5, 0xe5, 0x41, 0x2f, 0x26, 0x04, 0x17, 0x79, 0x69, 0x65, 0xb0, 0x14, 0x31, 0x29, 0x51,
	0x2f, 0x2f, 0xae, 0x0f, 0x17, 0x53, 0x2f, 0x88, 0x5b, 0x5c, 0x44, 0x5e, 0x1a, 0xf0, 0x98, 0x40,
	0x6b, 0xef, 0x0d, 0x80, 0x43, 0x4f, 0x78, 0x58, 0x65, 0x42, 0x9a, 0xb7, 0xa0, 0xda, 0xe4, 0x51,
	0x92, 0xa5, 0xcc, 0x95, 0x59, 0x64, 0x15, 0x2e, 0x14, 0x2e, 0x56, 0x37, 0x56, 0xec, 0xbc, 0xf3,
	0xbc, 0xad, 0x63, 0xce, 0x43, 0xe6, 0xc5, 0x47, 0x6a, 0xe5, 0x40, 0x8e, 0xbf, 0x9f, 0x45, 0xe6,
	0x25, 0x98, 0xef, 0x63, 0xbb, 0x71, 0x16, 0x86, 0xd6, 0x07, 0x03, 0x35, 0x2a, 0xce, 0x6c, 0x0f,
	0xd6, 0xc0, 0xc7, 0x66, 0x1d, 0x8a, 0xbe, 0xd7, 0xb1, 0xfe, 0xd2, 0x06, 0xff, 0x0f, 0x19, 0xa4,
	0x01, 0xee, 0x24, 0xf5, 0xa2, 0x84, 0x2c, 0x14, 0xd2, 0xac, 0x41, 0x05, 0x7f, 0x48, 0xf3, 0x23,
	0x69, 0x1a, 0xf8, 0x40, 0x8b, 0x5d, 0x07, 0x23, 0xc1, 0x7d, 0xbb, 0x31, 0xb7, 0x8a, 0x23, 0x3b,
	0x0e, 0xe2, 0x94, 0xb5, 0x99, 0x20, 0xb9, 0xb2, 0xc2, 0x36, 0xf0, 0x7d, 0xc0, 0x4c, 0xce, 0x22,
	0xd5, 0x4f, 0xa4, 0x0a, 0x54, 0xd6, 0xc2, 0x37, 0xa1, 0xd2, 0x12, 0x3c, 0x72, 0x55, 0xab, 0x53,
	0x93, 0xb4, 0x6a, 0x28, 0xf8, 0x0e, 0xb6, 0xbb, 0x0e, 0x7f, 0x77, 0x99, 0xa4, 0xfe, 0x99, 0xd4,
	0xab, 0x39, 0x20, 0xef, 0xbb, 0x9c, 0x72, 0x2d, 0x5e, 0x9a, 0x44, 0xbc, 0x94, 0x72, 0x25, 0x7d,
	0x1e, 0xaa, 0xc4, 0x22, 0xe1, 0x2f, 0x24, 0x3c, 0xad, 0x8b, 0x5a, 0x76, 0x17, 0xe6, 0xf5, 0xc6,
	0x9a, 0x1e, 0xee, 0x9a, 0x8b, 0x8e, 0x1b, 0xf8, 0x56, 0x79, 0xfc, 0xb9, 0xcc, 0x2a, 0xd2, 0x9d,
	0x9c, 0xb3, 0xef, 0x9b, 0x1b, 0xf0, 0xcf, 0xb0, 0x0c, 0x39, 0x7e, 0x25, 0x47, 0x73, 0x10, 0xaf,
	0xad, 0xf7, 0x60, 0xa1, 0x2d, 0x78, 0x96, 0xb8, 0xc7, 0x9d, 0x5f, 0x3c, 0xcb, 0x18, 0x7f, 0x8b,
	0xe6, 0x34, 0x6b, 0xbb, 0xd3, 0x95, 0xc3, 0xa3, 0x59, 0x3a, 0x23, 0x44, 0xee, 0xdf, 0xc8, 0x7d,
	0x71, 0x88, 0xa1, 0xed, 0x6f, 0xc3, 0x0c, 0x17, 0x3e, 0x13, 0x8a, 0xa5, 0x8e, 0xb5, 0x32, 0xc1,
	0xfd, 0xd5, 0x84, 0xed, 0x8e, 0x3a, 0xd9, 0xcb, 0xb0, 0xd0, 0x4f, 0x27, 0xbf, 0xef, 0xf9, 0x05,
	0xee, 0xe1, 0xb4, 0xd5, 0x23, 0xf8, 0xaf, 0xcd, 0x52, 0x7d, 0xcf, 0x79, 0xcb, 0xf5, 0xc2, 0xd0,
	0x55, 0xc7, 0x21, 0xdd, 0x04, 0xe9, 0xca, 0x7a, 0x7a, 0xbc, 0xf5, 0x12, 0x0a, 0xe0, 0x3c, 0xdc,
	0x6b, 0x6d, 0x85, 0xa1, 0xca, 0x01, 0x79, 0xc8, 0x84, 0x6a, 0xe3, 0x2e, 0xac, 0xfe, 0x4e, 0x9a,
	0xda, 0xfa, 0x41, 0x6d, 0xd5, 0x46, 0x6b, 0xa8, 0x16, 0xd7, 0xde, 0x95, 0xa0, 0xe2, 0x30, 0x99,
	0xf0, 0x58, 0x32, 0xf3, 0x2a, 0x94, 0x74, 0x72, 0xe4, 0x33, 0x5d, 0xb3, 0x07, 0xd3, 0x88, 0x52,
	0x65, 0x57, 0xfd, 0x75, 0x08, 0x88, 0x3b, 0x9c, 0x57, 0x99, 0xe1, 0xf6, 0x85, 0x06, 0xce, 0x6b,
	0x11, 0xc9, 0xf6, 0x10, 0x79, 0x38, 0x5a, 0x0e, 0x70, 0xbd, 0xdf, 0x5b, 0x3b, 0x73, 0xd1, 0xe0,
	0x03, 0x9c, 0x2b, 0x23, 0xcf, 0x2a, 0x1c, 0x58, 0xa5, 0x78, 0xee, 0x8c, 0x22, 0x25, 0xd9, 0x01,
	0xfd, 0x3a, 0x5d, 0xb8, 0xb9, 0x05, 0x45, 0xc1, 0x4f, 0x70, 0x18, 0x15, 0xab, 0x6e, 0x8f, 0x8d,
	0x54, 0xbb, 0x7b, 0x00, 0xb6, 0xc3, 0x4f, 0x1c, 0xc5, 0xad, 0x3d, 0x2b, 0x42, 0x11, 0x17, 0xe6,
	0x12, 0x94, 0x71, 0xa9, 0x86, 0xe3, 0x79, 0x03, 0xcf, 0xa4, 0xe4, 0x94, 0x70, 0x89, 0xf7, 0xfe,
	0x06, 0x18, 0x4d, 0x9e, 0xe1, 0x64, 0x08, 0xeb, 0x45, 0x63, 0xfc, 0xd8, 0x74, 0xc1, 0xe6, 0x66,
	0x2f, 0x85, 0x5e, 0x36, 0x26, 0x8f, 0xa1, 0x3a, 0x4c, 0xa9, 0xff, 0xac, 0x57, 0x8d, 0xc1, 0x17,
	0x93, 0x73, 0x64, 0x2a, 0x82, 0xb8, 0x4d, 0x14, 0x0d, 0x34, 0xaf, 0x50, 0x74, 0xbe, 0x1e, 0x8f,
	0xd7, 0xc1, 0xb9, 0x37, 0x22, 0x0d, 0xde, 0x9c, 0x16, 0xfe, 0x3c, 0x0f, 0x1e, 0xc2, 0xf2, 0xa0,
	0x90, 0xcf, 0x64, 0x53, 0x04, 0x89, 0xbe, 0x18, 0x6f, 0x4f, 0x0b, 0x63, 0xdb, 0xf9, 0xb7, 0x5f,
	0x70, 0xa7, 0x47, 0xde, 0x7e, 0x00, 0x2b, 0x01, 0x1f, 0x7a, 0x97, 0xbd, 0x4f, 0xea, 0xe3, 0xcd,
	0x36, 0x97, 0xfe, 0x93, 0x6e, 0xdd, 0x9f, 0xf0, 0xab, 0x7b, 0x5c, 0xd6, 0x5f, 0xb8, 0x6b, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x07, 0xa9, 0x63, 0xac, 0x07, 0x00, 0x00,
}
