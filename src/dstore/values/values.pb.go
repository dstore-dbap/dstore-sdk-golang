// Code generated by protoc-gen-go.
// source: dstore/values.proto
// DO NOT EDIT!

/*
Package values is a generated protocol buffer package.

It is generated from these files:
	dstore/values.proto

It has these top-level messages:
	IntegerValue
	StringValue
	BytesValue
	DoubleValue
	BooleanValue
	DecimalValue
	TimestampValue
	LongValue
	Value
*/
package values

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IntegerValue struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *IntegerValue) Reset()                    { *m = IntegerValue{} }
func (m *IntegerValue) String() string            { return proto.CompactTextString(m) }
func (*IntegerValue) ProtoMessage()               {}
func (*IntegerValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StringValue struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringValue) Reset()                    { *m = StringValue{} }
func (m *StringValue) String() string            { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()               {}
func (*StringValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BytesValue struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesValue) Reset()                    { *m = BytesValue{} }
func (m *BytesValue) String() string            { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()               {}
func (*BytesValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DoubleValue struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
}

func (m *DoubleValue) Reset()                    { *m = DoubleValue{} }
func (m *DoubleValue) String() string            { return proto.CompactTextString(m) }
func (*DoubleValue) ProtoMessage()               {}
func (*DoubleValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type BooleanValue struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *BooleanValue) Reset()                    { *m = BooleanValue{} }
func (m *BooleanValue) String() string            { return proto.CompactTextString(m) }
func (*BooleanValue) ProtoMessage()               {}
func (*BooleanValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DecimalValue struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DecimalValue) Reset()                    { *m = DecimalValue{} }
func (m *DecimalValue) String() string            { return proto.CompactTextString(m) }
func (*DecimalValue) ProtoMessage()               {}
func (*DecimalValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type TimestampValue struct {
	Value *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *TimestampValue) Reset()                    { *m = TimestampValue{} }
func (m *TimestampValue) String() string            { return proto.CompactTextString(m) }
func (*TimestampValue) ProtoMessage()               {}
func (*TimestampValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TimestampValue) GetValue() *google_protobuf.Timestamp {
	if m != nil {
		return m.Value
	}
	return nil
}

type LongValue struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *LongValue) Reset()                    { *m = LongValue{} }
func (m *LongValue) String() string            { return proto.CompactTextString(m) }
func (*LongValue) ProtoMessage()               {}
func (*LongValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_IntegerValue
	//	*Value_StringValue
	//	*Value_ByteValue
	//	*Value_DoubleValue
	//	*Value_BooleanValue
	//	*Value_DecimalValue
	//	*Value_TimestampValue
	//	*Value_LongValue
	Value isValue_Value `protobuf_oneof:"value"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type isValue_Value interface {
	isValue_Value()
}

type Value_IntegerValue struct {
	IntegerValue *IntegerValue `protobuf:"bytes,10,opt,name=integer_value,json=integerValue,oneof"`
}
type Value_StringValue struct {
	StringValue *StringValue `protobuf:"bytes,11,opt,name=string_value,json=stringValue,oneof"`
}
type Value_ByteValue struct {
	ByteValue *BytesValue `protobuf:"bytes,12,opt,name=byte_value,json=byteValue,oneof"`
}
type Value_DoubleValue struct {
	DoubleValue *DoubleValue `protobuf:"bytes,13,opt,name=double_value,json=doubleValue,oneof"`
}
type Value_BooleanValue struct {
	BooleanValue *BooleanValue `protobuf:"bytes,14,opt,name=boolean_value,json=booleanValue,oneof"`
}
type Value_DecimalValue struct {
	DecimalValue *DecimalValue `protobuf:"bytes,15,opt,name=decimal_value,json=decimalValue,oneof"`
}
type Value_TimestampValue struct {
	TimestampValue *TimestampValue `protobuf:"bytes,16,opt,name=timestamp_value,json=timestampValue,oneof"`
}
type Value_LongValue struct {
	LongValue *LongValue `protobuf:"bytes,17,opt,name=long_value,json=longValue,oneof"`
}

func (*Value_IntegerValue) isValue_Value()   {}
func (*Value_StringValue) isValue_Value()    {}
func (*Value_ByteValue) isValue_Value()      {}
func (*Value_DoubleValue) isValue_Value()    {}
func (*Value_BooleanValue) isValue_Value()   {}
func (*Value_DecimalValue) isValue_Value()   {}
func (*Value_TimestampValue) isValue_Value() {}
func (*Value_LongValue) isValue_Value()      {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetIntegerValue() *IntegerValue {
	if x, ok := m.GetValue().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return nil
}

func (m *Value) GetStringValue() *StringValue {
	if x, ok := m.GetValue().(*Value_StringValue); ok {
		return x.StringValue
	}
	return nil
}

func (m *Value) GetByteValue() *BytesValue {
	if x, ok := m.GetValue().(*Value_ByteValue); ok {
		return x.ByteValue
	}
	return nil
}

func (m *Value) GetDoubleValue() *DoubleValue {
	if x, ok := m.GetValue().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return nil
}

func (m *Value) GetBooleanValue() *BooleanValue {
	if x, ok := m.GetValue().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return nil
}

func (m *Value) GetDecimalValue() *DecimalValue {
	if x, ok := m.GetValue().(*Value_DecimalValue); ok {
		return x.DecimalValue
	}
	return nil
}

func (m *Value) GetTimestampValue() *TimestampValue {
	if x, ok := m.GetValue().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Value) GetLongValue() *LongValue {
	if x, ok := m.GetValue().(*Value_LongValue); ok {
		return x.LongValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_IntegerValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_ByteValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_DecimalValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_LongValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_IntegerValue:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IntegerValue); err != nil {
			return err
		}
	case *Value_StringValue:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StringValue); err != nil {
			return err
		}
	case *Value_ByteValue:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ByteValue); err != nil {
			return err
		}
	case *Value_DoubleValue:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DoubleValue); err != nil {
			return err
		}
	case *Value_BooleanValue:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BooleanValue); err != nil {
			return err
		}
	case *Value_DecimalValue:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DecimalValue); err != nil {
			return err
		}
	case *Value_TimestampValue:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampValue); err != nil {
			return err
		}
	case *Value_LongValue:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LongValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.Value has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 10: // value.integer_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IntegerValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_IntegerValue{msg}
		return true, err
	case 11: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_StringValue{msg}
		return true, err
	case 12: // value.byte_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BytesValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_ByteValue{msg}
		return true, err
	case 13: // value.double_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DoubleValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_DoubleValue{msg}
		return true, err
	case 14: // value.boolean_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BooleanValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_BooleanValue{msg}
		return true, err
	case 15: // value.decimal_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DecimalValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_DecimalValue{msg}
		return true, err
	case 16: // value.timestamp_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TimestampValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_TimestampValue{msg}
		return true, err
	case 17: // value.long_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LongValue)
		err := b.DecodeMessage(msg)
		m.Value = &Value_LongValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_IntegerValue:
		s := proto.Size(x.IntegerValue)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_StringValue:
		s := proto.Size(x.StringValue)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ByteValue:
		s := proto.Size(x.ByteValue)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_DoubleValue:
		s := proto.Size(x.DoubleValue)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_BooleanValue:
		s := proto.Size(x.BooleanValue)
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_DecimalValue:
		s := proto.Size(x.DecimalValue)
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_TimestampValue:
		s := proto.Size(x.TimestampValue)
		n += proto.SizeVarint(16<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_LongValue:
		s := proto.Size(x.LongValue)
		n += proto.SizeVarint(17<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*IntegerValue)(nil), "dstore.values.integerValue")
	proto.RegisterType((*StringValue)(nil), "dstore.values.stringValue")
	proto.RegisterType((*BytesValue)(nil), "dstore.values.bytesValue")
	proto.RegisterType((*DoubleValue)(nil), "dstore.values.doubleValue")
	proto.RegisterType((*BooleanValue)(nil), "dstore.values.booleanValue")
	proto.RegisterType((*DecimalValue)(nil), "dstore.values.decimalValue")
	proto.RegisterType((*TimestampValue)(nil), "dstore.values.timestampValue")
	proto.RegisterType((*LongValue)(nil), "dstore.values.longValue")
	proto.RegisterType((*Value)(nil), "dstore.values.Value")
}

var fileDescriptor0 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xbf, 0xea, 0xa3, 0x40, 0x6e, 0x93, 0x16, 0x0a, 0x43, 0x29, 0x42, 0x40, 0x61, 0x60,
	0x72, 0x11, 0x4c, 0xb0, 0x20, 0x65, 0xea, 0x5c, 0x21, 0x06, 0x16, 0x94, 0x50, 0x13, 0x45, 0x72,
	0xeb, 0x2a, 0x71, 0x90, 0x78, 0x1a, 0x5e, 0x15, 0xff, 0xb9, 0x0e, 0x76, 0x64, 0x31, 0x3a, 0x39,
	0xf7, 0x77, 0x9d, 0x73, 0x4e, 0xe0, 0x68, 0x55, 0x0b, 0x5e, 0xd1, 0xf9, 0x67, 0xc6, 0x1a, 0x5a,
	0x93, 0x6d, 0xc5, 0x05, 0x1f, 0x27, 0xe6, 0x21, 0x31, 0x0f, 0xa7, 0xe7, 0x05, 0xe7, 0x05, 0xa3,
	0x73, 0xfd, 0x32, 0x6f, 0x3e, 0xe6, 0xa2, 0x5c, 0xd3, 0x5a, 0x64, 0xeb, 0xad, 0xd1, 0xcf, 0xae,
	0x21, 0x2e, 0x37, 0x82, 0x16, 0xb4, 0x7a, 0x51, 0x13, 0xe3, 0x63, 0xe8, 0xeb, 0xd1, 0x49, 0xef,
	0xa2, 0x77, 0xd3, 0x5f, 0x9a, 0xc3, 0xec, 0x0a, 0x06, 0xb5, 0xa8, 0xca, 0x4d, 0x11, 0x10, 0x45,
	0x56, 0x34, 0x03, 0xc8, 0xbf, 0x04, 0xad, 0x03, 0x9a, 0xd8, 0x01, 0xad, 0x78, 0x93, 0x33, 0x1a,
	0x10, 0xf5, 0xac, 0x48, 0xde, 0x29, 0xe7, 0x9c, 0xd1, 0x6c, 0x13, 0x50, 0xed, 0x3b, 0xaa, 0x15,
	0x7d, 0x2f, 0xd7, 0x19, 0xfb, 0xeb, 0x52, 0x29, 0x0c, 0xdb, 0x4f, 0x36, 0xba, 0x5b, 0x57, 0x37,
	0xb8, 0x9b, 0x12, 0x63, 0x11, 0xb1, 0x16, 0x91, 0x67, 0xab, 0xb7, 0x8c, 0x4b, 0x88, 0x18, 0x0f,
	0x7e, 0xfb, 0x7f, 0x2b, 0xf9, 0xde, 0x81, 0xbe, 0x79, 0x9f, 0x42, 0x82, 0x86, 0xbe, 0x19, 0x1d,
	0xe8, 0x35, 0xa7, 0xc4, 0x0b, 0x86, 0xb8, 0xa6, 0x2f, 0xfe, 0x2d, 0xfd, 0x10, 0x9e, 0x20, 0x36,
	0x76, 0x23, 0x62, 0x80, 0x37, 0xf5, 0x11, 0x4e, 0x22, 0x92, 0xe0, 0x05, 0xf4, 0x68, 0xa2, 0xc0,
	0xf1, 0x58, 0x8f, 0x9f, 0x74, 0xc6, 0x7f, 0xb3, 0x92, 0xd3, 0x91, 0x3a, 0xb5, 0xcb, 0x4d, 0x44,
	0x38, 0x9d, 0x04, 0x97, 0x3b, 0x29, 0xaa, 0xe5, 0x6e, 0xa8, 0xd2, 0x01, 0x8c, 0x0f, 0x09, 0xc3,
	0xa0, 0x03, 0x6e, 0xc4, 0xca, 0x01, 0x2f, 0x72, 0xc9, 0xc0, 0x70, 0x91, 0x31, 0x0a, 0x32, 0xdc,
	0x02, 0x28, 0x86, 0x57, 0x88, 0x05, 0x8c, 0xda, 0xe8, 0x91, 0x72, 0xa0, 0x29, 0x67, 0x1d, 0x8a,
	0x5f, 0x10, 0xc9, 0xe9, 0x56, 0xe6, 0x01, 0x40, 0x15, 0x00, 0x21, 0x87, 0x1a, 0x32, 0xe9, 0x40,
	0xda, 0x86, 0x28, 0x37, 0xdb, 0x43, 0xba, 0x87, 0x75, 0x49, 0xa7, 0x10, 0x95, 0x1c, 0x67, 0x5e,
	0x13, 0xef, 0xd7, 0xcd, 0x77, 0x75, 0xf7, 0xee, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xf4,
	0x4b, 0x86, 0xd2, 0x03, 0x00, 0x00,
}