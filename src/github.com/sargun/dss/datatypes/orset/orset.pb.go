// Code generated by protoc-gen-go.
// source: orset.proto
// DO NOT EDIT!

/*
Package dss_datatypes_orset is a generated protocol buffer package.

It is generated from these files:
	orset.proto

It has these top-level messages:
	OrSetMember
	OrSet
	DelRq
	AddRq
	NewRq
	MergeRq
*/
package dss_datatypes_orset

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type OrSetMember struct {
	Value            []byte `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OrSetMember) Reset()         { *m = OrSetMember{} }
func (m *OrSetMember) String() string { return proto.CompactTextString(m) }
func (*OrSetMember) ProtoMessage()    {}

func (m *OrSetMember) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type OrSet struct {
	AddItems         []*OrSetMember `protobuf:"bytes,1,rep,name=add_items" json:"add_items,omitempty"`
	DelItems         []*OrSetMember `protobuf:"bytes,2,rep,name=del_items" json:"del_items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *OrSet) Reset()         { *m = OrSet{} }
func (m *OrSet) String() string { return proto.CompactTextString(m) }
func (*OrSet) ProtoMessage()    {}

func (m *OrSet) GetAddItems() []*OrSetMember {
	if m != nil {
		return m.AddItems
	}
	return nil
}

func (m *OrSet) GetDelItems() []*OrSetMember {
	if m != nil {
		return m.DelItems
	}
	return nil
}

type DelRq struct {
	Orset            *OrSet       `protobuf:"bytes,1,req,name=orset" json:"orset,omitempty"`
	Item             *OrSetMember `protobuf:"bytes,2,req,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DelRq) Reset()         { *m = DelRq{} }
func (m *DelRq) String() string { return proto.CompactTextString(m) }
func (*DelRq) ProtoMessage()    {}

func (m *DelRq) GetOrset() *OrSet {
	if m != nil {
		return m.Orset
	}
	return nil
}

func (m *DelRq) GetItem() *OrSetMember {
	if m != nil {
		return m.Item
	}
	return nil
}

type AddRq struct {
	Orset            *OrSet       `protobuf:"bytes,1,req,name=orset" json:"orset,omitempty"`
	Item             *OrSetMember `protobuf:"bytes,2,req,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AddRq) Reset()         { *m = AddRq{} }
func (m *AddRq) String() string { return proto.CompactTextString(m) }
func (*AddRq) ProtoMessage()    {}

func (m *AddRq) GetOrset() *OrSet {
	if m != nil {
		return m.Orset
	}
	return nil
}

func (m *AddRq) GetItem() *OrSetMember {
	if m != nil {
		return m.Item
	}
	return nil
}

type NewRq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *NewRq) Reset()         { *m = NewRq{} }
func (m *NewRq) String() string { return proto.CompactTextString(m) }
func (*NewRq) ProtoMessage()    {}

type MergeRq struct {
	Orset            []*OrSet `protobuf:"bytes,1,rep,name=orset" json:"orset,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MergeRq) Reset()         { *m = MergeRq{} }
func (m *MergeRq) String() string { return proto.CompactTextString(m) }
func (*MergeRq) ProtoMessage()    {}

func (m *MergeRq) GetOrset() []*OrSet {
	if m != nil {
		return m.Orset
	}
	return nil
}

func init() {
}