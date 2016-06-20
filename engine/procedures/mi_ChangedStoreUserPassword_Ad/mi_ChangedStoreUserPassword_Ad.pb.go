// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ChangedStoreUserPassword_Ad.proto
// DO NOT EDIT!

/*
Package mi_ChangedStoreUserPassword_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ChangedStoreUserPassword_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ChangedStoreUserPassword_Ad

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
	UserName                          *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserNameNull                      bool                        `protobuf:"varint,1001,opt,name=user_name_null,json=userNameNull" json:"user_name_null,omitempty"`
	DBLogin                           *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=d_b_login,json=dBLogin" json:"d_b_login,omitempty"`
	DBLoginNull                       bool                        `protobuf:"varint,1002,opt,name=d_b_login_null,json=dBLoginNull" json:"d_b_login_null,omitempty"`
	CallerPassword                    *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=caller_password,json=callerPassword" json:"caller_password,omitempty"`
	CallerPasswordNull                bool                        `protobuf:"varint,1003,opt,name=caller_password_null,json=callerPasswordNull" json:"caller_password_null,omitempty"`
	NewPassword                       *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=new_password,json=newPassword" json:"new_password,omitempty"`
	NewPasswordNull                   bool                        `protobuf:"varint,1004,opt,name=new_password_null,json=newPasswordNull" json:"new_password_null,omitempty"`
	EncryptedPassword                 *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=encrypted_password,json=encryptedPassword" json:"encrypted_password,omitempty"`
	EncryptedPasswordNull             bool                        `protobuf:"varint,1005,opt,name=encrypted_password_null,json=encryptedPasswordNull" json:"encrypted_password_null,omitempty"`
	CreatedstoreUserIfNotExists       *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=createdstore_user_if_not_exists,json=createdstoreUserIfNotExists" json:"createdstore_user_if_not_exists,omitempty"`
	CreatedstoreUserIfNotExistsNull   bool                        `protobuf:"varint,1006,opt,name=createdstore_user_if_not_exists_null,json=createdstoreUserIfNotExistsNull" json:"createdstore_user_if_not_exists_null,omitempty"`
	IgnoreErrorNewPasswdSameAsOld     *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=ignore_error_new_passwd_same_as_old,json=ignoreErrorNewPasswdSameAsOld" json:"ignore_error_new_passwd_same_as_old,omitempty"`
	IgnoreErrorNewPasswdSameAsOldNull bool                        `protobuf:"varint,1007,opt,name=ignore_error_new_passwd_same_as_old_null,json=ignoreErrorNewPasswdSameAsOldNull" json:"ignore_error_new_passwd_same_as_old_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Parameters) GetDBLogin() *dstore_values.StringValue {
	if m != nil {
		return m.DBLogin
	}
	return nil
}

func (m *Parameters) GetCallerPassword() *dstore_values.StringValue {
	if m != nil {
		return m.CallerPassword
	}
	return nil
}

func (m *Parameters) GetNewPassword() *dstore_values.StringValue {
	if m != nil {
		return m.NewPassword
	}
	return nil
}

func (m *Parameters) GetEncryptedPassword() *dstore_values.StringValue {
	if m != nil {
		return m.EncryptedPassword
	}
	return nil
}

func (m *Parameters) GetCreatedstoreUserIfNotExists() *dstore_values.BooleanValue {
	if m != nil {
		return m.CreatedstoreUserIfNotExists
	}
	return nil
}

func (m *Parameters) GetIgnoreErrorNewPasswdSameAsOld() *dstore_values.BooleanValue {
	if m != nil {
		return m.IgnoreErrorNewPasswdSameAsOld
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ChangedStoreUserPassword_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ChangedStoreUserPassword_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ChangedStoreUserPassword_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ChangedStoreUserPassword_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xeb, 0x4e, 0x13, 0x4f,
	0x14, 0x0f, 0xf4, 0x5f, 0x0a, 0x07, 0x02, 0x7f, 0xc6, 0xdb, 0xda, 0x06, 0x51, 0xc0, 0x84, 0x44,
	0xb3, 0x78, 0x49, 0xc4, 0x98, 0x90, 0x08, 0x84, 0x0f, 0x4d, 0x60, 0x25, 0x8b, 0x9a, 0xe8, 0x97,
	0xc9, 0xb4, 0x33, 0x5d, 0x37, 0x6e, 0x67, 0x9a, 0x99, 0xad, 0xd5, 0xb7, 0xf0, 0x21, 0x7c, 0x0e,
	0xdf, 0xc7, 0xfb, 0x2b, 0x38, 0x97, 0xed, 0xb6, 0xdd, 0x26, 0x5d, 0xf9, 0xd4, 0x9e, 0x39, 0xbf,
	0x5b, 0x67, 0x7a, 0x0e, 0x1c, 0x50, 0x95, 0x0a, 0xc9, 0xf6, 0x18, 0x8f, 0x62, 0xce, 0xf6, 0x7a,
	0x52, 0xb4, 0x19, 0xed, 0x4b, 0xa6, 0xf6, 0xba, 0x31, 0x3e, 0x7e, 0x47, 0x78, 0xc4, 0xe8, 0x85,
	0x41, 0xbc, 0x52, 0x4c, 0x9e, 0x13, 0xa5, 0x06, 0x42, 0x52, 0x7c, 0x48, 0x7d, 0x8d, 0x4c, 0x05,
	0xba, 0xef, 0xe8, 0xbe, 0xa3, 0xfb, 0xb3, 0x39, 0xf5, 0x2b, 0x99, 0xd9, 0x07, 0x92, 0xf4, 0x99,
	0x72, 0x12, 0xf5, 0x9b, 0x93, 0x09, 0x98, 0x94, 0x42, 0x66, 0xad, 0xc6, 0x64, 0xab, 0xcb, 0x94,
	0x22, 0x11, 0xcb, 0x9a, 0xdb, 0xc5, 0x66, 0x4a, 0x62, 0xde, 0x11, 0xb2, 0x4b, 0xd2, 0x58, 0x70,
	0x07, 0xda, 0xfa, 0x52, 0x03, 0x38, 0x27, 0x92, 0xe8, 0x2e, 0x93, 0x0a, 0xed, 0xc3, 0x52, 0x5f,
	0x67, 0xc2, 0x5c, 0x1f, 0x78, 0x73, 0xb7, 0xe7, 0x76, 0x97, 0x1f, 0xd5, 0xfd, 0xec, 0x27, 0x64,
	0xa1, 0x54, 0x2a, 0x63, 0x1e, 0xbd, 0x36, 0x45, 0xb8, 0x68, 0xc0, 0x81, 0xc6, 0xa2, 0xbb, 0xb0,
	0x9a, 0x13, 0x31, 0xef, 0x27, 0x89, 0xf7, 0xad, 0xa6, 0xe9, 0x8b, 0xe1, 0xca, 0x10, 0x12, 0xe8,
	0x43, 0xf4, 0x04, 0x96, 0x28, 0x6e, 0xe1, 0x44, 0xe8, 0x48, 0xde, 0x7c, 0xa9, 0x7e, 0x8d, 0x1e,
	0x9d, 0x1a, 0x28, 0xda, 0x81, 0xd5, 0x9c, 0xe7, 0xe4, 0xbf, 0x3b, 0xf9, 0xe5, 0x0c, 0x61, 0xd5,
	0x8f, 0x61, 0xad, 0x4d, 0x92, 0x44, 0xc7, 0xe8, 0x65, 0x97, 0xea, 0x55, 0x4a, 0x3d, 0x56, 0x1d,
	0x65, 0xf8, 0x0c, 0xe8, 0x21, 0x5c, 0x2d, 0x88, 0x38, 0xc3, 0x1f, 0xce, 0x10, 0x4d, 0xc2, 0xad,
	0xef, 0x01, 0xac, 0x70, 0x36, 0x18, 0x99, 0xfe, 0x57, 0x6a, 0xba, 0xac, 0xf1, 0xb9, 0xe3, 0x3d,
	0x58, 0x1f, 0xa7, 0x3b, 0xbb, 0x9f, 0xce, 0x6e, 0x6d, 0x0c, 0x68, 0xbd, 0x9a, 0x80, 0x18, 0x6f,
	0xcb, 0x4f, 0xbd, 0x94, 0xd1, 0x91, 0x63, 0xb5, 0xd4, 0x71, 0x3d, 0x67, 0xe5, 0xbe, 0xfb, 0x70,
	0x63, 0x5a, 0xca, 0xb9, 0xff, 0x72, 0xee, 0xd7, 0xa6, 0x48, 0x36, 0x03, 0x81, 0xcd, 0xb6, 0x64,
	0x44, 0x1f, 0x5b, 0x3b, 0x6c, 0x5f, 0x3e, 0xee, 0x60, 0x2e, 0x52, 0xcc, 0x3e, 0xc6, 0x2a, 0x55,
	0xde, 0x82, 0x0d, 0xd4, 0x28, 0x04, 0x6a, 0x09, 0x91, 0x30, 0xc2, 0x5d, 0xa2, 0xc6, 0xb8, 0x86,
	0x99, 0x84, 0x66, 0x27, 0x10, 0xe9, 0x89, 0xe5, 0xa3, 0x00, 0x76, 0x4a, 0x2c, 0x5c, 0xd0, 0xdf,
	0x2e, 0xe8, 0xe6, 0x0c, 0x2d, 0x1b, 0x39, 0x82, 0xed, 0x38, 0xe2, 0x46, 0xc9, 0xce, 0x0f, 0xce,
	0x2f, 0x9c, 0x62, 0x65, 0xfe, 0xb2, 0x44, 0x61, 0x91, 0x50, 0xaf, 0x56, 0x1e, 0x7b, 0xc3, 0xe9,
	0x9c, 0x18, 0x99, 0x20, 0x7b, 0x1c, 0x7a, 0xa1, 0x35, 0x0e, 0xd5, 0x8b, 0x84, 0xa2, 0x97, 0xb0,
	0xfb, 0x0f, 0x46, 0x2e, 0xfc, 0x1f, 0x17, 0xfe, 0xce, 0x4c, 0x45, 0x13, 0x7f, 0xeb, 0xeb, 0x3c,
	0x2c, 0x86, 0x4c, 0xf5, 0x04, 0x57, 0x0c, 0x3d, 0x80, 0xaa, 0xd5, 0x2e, 0x0e, 0x68, 0xb6, 0x63,
	0xdc, 0x82, 0xb0, 0x6a, 0xa1, 0x03, 0xa2, 0x37, 0xf0, 0xbf, 0x19, 0x7f, 0x3c, 0x36, 0xff, 0x7a,
	0xfa, 0x2a, 0x9a, 0xec, 0x17, 0xc8, 0xc5, 0x2d, 0x71, 0xa6, 0xeb, 0xe6, 0xa8, 0x0e, 0xd7, 0xba,
	0x93, 0x07, 0xe8, 0x29, 0xd4, 0xb2, 0xb5, 0xa3, 0x67, 0xcd, 0x28, 0xde, 0x9a, 0x52, 0x74, 0x4b,
	0xe9, 0xcc, 0x7d, 0x86, 0x43, 0x38, 0x3a, 0x85, 0x8a, 0x14, 0x03, 0x3d, 0x2c, 0x86, 0xf5, 0xcc,
	0xbf, 0xcc, 0xa2, 0xf4, 0x87, 0x77, 0xe1, 0x87, 0x62, 0x10, 0x1a, 0x99, 0xfa, 0x06, 0x54, 0xf4,
	0x77, 0x74, 0x1d, 0x16, 0x74, 0x85, 0x63, 0xea, 0x7d, 0x0e, 0xf4, 0xed, 0x54, 0xc3, 0xaa, 0x2e,
	0x9b, 0xf4, 0x08, 0x43, 0x23, 0x16, 0x05, 0x8f, 0xd1, 0x2e, 0x7f, 0xfb, 0x3c, 0x12, 0x8a, 0xbe,
	0x1f, 0xf6, 0xe9, 0xe5, 0xd7, 0x7d, 0x6b, 0xc1, 0xee, 0xd3, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x4b, 0xf8, 0xb7, 0x3a, 0x30, 0x06, 0x00, 0x00,
}