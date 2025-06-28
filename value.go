package protoreflectextra

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewMessageListField(m protoreflect.Message, field protoreflect.FieldDescriptor) List {
	return NewList(m.Type(), ListDescriptorOfField(field), m.NewField(field).List())
}

func GetMessageListField(m protoreflect.Message, field protoreflect.FieldDescriptor) List {
	return NewList(m.Type(), ListDescriptorOfField(field), m.Get(field).List())
}

func MutableMessageListField(m protoreflect.Message, field protoreflect.FieldDescriptor) List {
	return NewList(m.Type(), ListDescriptorOfField(field), m.Mutable(field).List())
}

func NewMessageMapField(m protoreflect.Message, field protoreflect.FieldDescriptor) Map {
	return NewMap(m.Type(), MapDescriptorOfField(field), m.NewField(field).Map())
}

func GetMessageMapField(m protoreflect.Message, field protoreflect.FieldDescriptor) Map {
	return NewMap(m.Type(), MapDescriptorOfField(field), m.Get(field).Map())
}

func MutableMessageMapField(m protoreflect.Message, field protoreflect.FieldDescriptor) Map {
	return NewMap(m.Type(), MapDescriptorOfField(field), m.Mutable(field).Map())
}

func NewMessageEnumField(m protoreflect.Message, field protoreflect.FieldDescriptor) protoreflect.Enum {
	return NewEnum(field.Enum(), m.NewField(field).Enum())
}

func GetMessageEnumField(m protoreflect.Message, field protoreflect.FieldDescriptor) protoreflect.Enum {
	return NewEnum(field.Enum(), m.Get(field).Enum())
}

func NewListEnumElement(l List) protoreflect.Enum {
	return NewEnum(l.Descriptor().ListElement().Enum(), l.NewElement().Enum())
}

func GetListEnumElement(l List, i int) protoreflect.Enum {
	return NewEnum(l.Descriptor().ListElement().Enum(), l.Get(i).Enum())
}

func NewMapEnumValue(m Map) protoreflect.Enum {
	return NewEnum(m.Descriptor().MapValue().Enum(), m.NewValue().Enum())
}

func GetMapEnumValue(m Map, k protoreflect.MapKey) protoreflect.Enum {
	return NewEnum(m.Descriptor().MapValue().Enum(), m.Get(k).Enum())
}
