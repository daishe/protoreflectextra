package protoreflectextra

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type enumType struct {
	desc protoreflect.EnumDescriptor
}

func NewEnumType(ed protoreflect.EnumDescriptor) protoreflect.EnumType {
	return enumType{desc: ed}
}

func (et enumType) New(n protoreflect.EnumNumber) protoreflect.Enum {
	return NewEnum(et.desc, n)
}

func (et enumType) Descriptor() protoreflect.EnumDescriptor {
	return et.desc
}

type enum struct {
	desc protoreflect.EnumDescriptor
	n    protoreflect.EnumNumber
}

func NewEnum(ed protoreflect.EnumDescriptor, n protoreflect.EnumNumber) protoreflect.Enum {
	return enum{desc: ed, n: n}
}

func (e enum) Descriptor() protoreflect.EnumDescriptor {
	return e.desc
}

func (e enum) Type() protoreflect.EnumType {
	return enumType{desc: e.desc}
}

func (e enum) Number() protoreflect.EnumNumber {
	return e.n
}
