package protoreflectextra

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type VirtualMapDescriptor interface {
	// Parent returns the parent field descriptor that this map belongs to.
	Parent() protoreflect.FieldDescriptor

	// MapKey returns the field descriptor for the key in the map entry. It returns nil if IsMap reports false.
	MapKey() protoreflect.FieldDescriptor

	// MapValue returns the field descriptor for the value in the map entry. It returns nil if IsMap reports false.
	MapValue() protoreflect.FieldDescriptor
}

type mapDescriptor struct{ field protoreflect.FieldDescriptor }

func MapDescriptorOfField(field protoreflect.FieldDescriptor) VirtualMapDescriptor {
	if !field.IsMap() {
		panic("protoreflectextra: cannot use provided field descriptor as map descriptor: given field descriptor do not describes a map field")
	}
	return mapDescriptor{field}
}

func (md mapDescriptor) Parent() protoreflect.FieldDescriptor {
	return md.field
}

func (md mapDescriptor) MapKey() protoreflect.FieldDescriptor {
	return md.field.MapKey()
}

func (md mapDescriptor) MapValue() protoreflect.FieldDescriptor {
	return md.field.MapValue()
}

// MapType encapsulates an VirtualMapDescriptor with a map implementation.
type MapType interface {
	// New returns a new, empty and mutable instance of empty map.
	New() Map

	// Zero returns a new, read-only instance of map.
	Zero() Map

	// Descriptor returns the map descriptor.
	//
	// Invariants:
	//   t.Descriptor() == t.New().Descriptor()
	//   t.Descriptor() == t.Zero().Descriptor()
	Descriptor() VirtualMapDescriptor
}

type mapType struct{ new func(init bool) Map }

func NewMapType(mt protoreflect.MessageType, mapDesc VirtualMapDescriptor) MapType {
	return mapType{
		new: func(init bool) Map {
			if init {
				return NewMap(mt, mapDesc, mt.Zero().NewField(mapDesc.Parent()).Map())
			}
			return NewMap(mt, mapDesc, mt.Zero().Get(mapDesc.Parent()).Map())
		},
	}
}

func (mt mapType) New() Map {
	return mt.new(true)
}

func (mt mapType) Zero() Map {
	return mt.new(false)
}

func (mt mapType) Descriptor() VirtualMapDescriptor {
	return mt.new(false).Descriptor()
}

// Map is an extended reflection interface for an unordered, associative map. It provides type information and ability to mutate the map or retrieve its children. The entry protoreflect.MapKey type is determined by [VirtualMapDescriptor.MapKey].Kind. The entry protoreflect.Value type is determined by [VirtualMapDescriptor.MapValue].Kind. Providing a protoreflect.MapKey or protoreflect.Value that is invalid or of an incorrect type panics.
type Map interface {
	// Descriptor returns map descriptor, which contains only the protobuf type information for the map.
	Descriptor() VirtualMapDescriptor

	// Type returns the map type, which encapsulates both map implementation and protobuf type information. If the map implementation is not needed, it is recommended that the map descriptor be used instead.
	Type() MapType

	// New returns a new, empty and mutable instance of empty map.
	New() Map

	// Zero returns a new, read-only instance of map.
	Zero() Map

	// Len reports the number of elements in the map.
	Len() int

	// Range iterates over every map entry in an undefined order, calling f for each key and value encountered. Range calls f Len times unless f returns false, which stops iteration. While iterating, mutating operations may only be performed on the current map key.
	Range(f func(protoreflect.MapKey, protoreflect.Value) bool)

	// Has reports whether an entry with the given key is in the map.
	Has(k protoreflect.MapKey) bool

	// Clear clears the entry associated with they given key. The operation does nothing if there is no entry associated with the key.
	//
	// Clear is a mutating operation and unsafe for concurrent use.
	Clear(k protoreflect.MapKey)

	// Get retrieves the value for an entry with the given key. It returns an invalid value for non-existent entries.
	Get(k protoreflect.MapKey) protoreflect.Value

	// Set stores the value for an entry with the given key. It panics when given a key or value that is invalid or the wrong type. When setting a composite type, it is unspecified whether the set value aliases the source's memory in any way.
	//
	// Set is a mutating operation and unsafe for concurrent use.
	Set(k protoreflect.MapKey, v protoreflect.Value)

	// Mutable retrieves a mutable reference to the entry for the given key. If no entry exists for the key, it creates a new, empty, mutable value and stores it as the entry for the key. It panics if the map value is not a message.
	Mutable(k protoreflect.MapKey) protoreflect.Value

	// NewKey returns a new value assignable as a map key (zero scalar value).
	NewKey() protoreflect.MapKey

	// NewValue returns a new value assignable as a map value. For enums, this returns the first enum value. For other scalars, this returns the zero value. For messages, this returns a new, empty, mutable value.
	NewValue() protoreflect.Value

	// IsMutable reports whether the map is mutable.
	//
	// An immutable map is an empty, read-only value, for which all content access or mutation operations will panic.
	//
	// An immutable map often corresponds to a nil Go map value, but the details are implementation dependent. Mutability is not part of the protobuf data model, and may not be preserved in marshaling or other operations.
	IsMutable() bool

	// Value returns underlying protoreflect.Value of the map.
	Value() protoreflect.Value
}

type _map struct {
	prMap protoreflect.Map
	desc  VirtualMapDescriptor
	new   func(init bool) Map
}

// NewMap return an extended reflection interface for a protoreflect.Map that combines it with type information.
func NewMap(mt protoreflect.MessageType, mapDesc VirtualMapDescriptor, prMap protoreflect.Map) Map {
	return _map{
		prMap: prMap,
		desc:  mapDesc,
		new: func(init bool) Map {
			if init {
				return NewMap(mt, mapDesc, mt.Zero().NewField(mapDesc.Parent()).Map())
			}
			return NewMap(mt, mapDesc, mt.Zero().Get(mapDesc.Parent()).Map())
		},
	}
}

func (m _map) Descriptor() VirtualMapDescriptor {
	return m.desc
}

func (m _map) Type() MapType {
	return mapType{m.new}
}

func (m _map) New() Map {
	return m.new(true)
}

func (m _map) Zero() Map {
	return m.new(false)
}

func (m _map) Len() int {
	return m.prMap.Len()
}

func (m _map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	m.prMap.Range(f)
}

func (m _map) Has(k protoreflect.MapKey) bool {
	return m.prMap.Has(k)
}

func (m _map) Clear(k protoreflect.MapKey) {
	m.prMap.Clear(k)
}

func (m _map) Get(k protoreflect.MapKey) protoreflect.Value {
	return m.prMap.Get(k)
}

func (m _map) Set(k protoreflect.MapKey, v protoreflect.Value) {
	m.prMap.Set(k, v)
}

func (m _map) Mutable(k protoreflect.MapKey) protoreflect.Value {
	return m.prMap.Mutable(k)
}

func (m _map) NewKey() protoreflect.MapKey {
	return m.desc.MapKey().Default().MapKey()
}

func (m _map) NewValue() protoreflect.Value {
	return m.prMap.NewValue()
}

func (m _map) IsMutable() bool {
	return m.prMap.IsValid()
}

func (m _map) Value() protoreflect.Value {
	return protoreflect.ValueOfMap(m.prMap)
}
