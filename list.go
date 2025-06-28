package protoreflectextra

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type VirtualListDescriptor interface {
	// Parent returns the parent field descriptor that this list belongs to.
	Parent() protoreflect.FieldDescriptor

	// ListElement returns the field descriptor for the value of an item in the list.
	ListElement() protoreflect.FieldDescriptor
}

type listDescriptor struct{ field protoreflect.FieldDescriptor }

func ListDescriptorOfField(field protoreflect.FieldDescriptor) VirtualListDescriptor {
	if !field.IsList() {
		panic("protoreflectextra: cannot use provided field descriptor as list descriptor: given field descriptor do not describes a list (repeated) field")
	}
	return listDescriptor{field}
}

func (ld listDescriptor) Parent() protoreflect.FieldDescriptor {
	return ld.field
}

func (ld listDescriptor) ListElement() protoreflect.FieldDescriptor {
	return listElementDescriptor{ld.field}
}

type listElementDescriptor struct{ protoreflect.FieldDescriptor }

func (led listElementDescriptor) Parent() protoreflect.Descriptor {
	return led.FieldDescriptor
}

func (led listElementDescriptor) Name() protoreflect.Name {
	return led.FieldDescriptor.Name() + ".element"
}

func (led listElementDescriptor) FullName() protoreflect.FullName {
	return led.FieldDescriptor.FullName() + ".element"
}

func (led listElementDescriptor) Index() int                                        { return 0 }
func (led listElementDescriptor) Options() protoreflect.ProtoMessage                { return nil }
func (led listElementDescriptor) Cardinality() protoreflect.Cardinality             { return protoreflect.Optional }
func (led listElementDescriptor) HasJSONName() bool                                 { return false }
func (led listElementDescriptor) JSONName() string                                  { return "element" }
func (led listElementDescriptor) TextName() string                                  { return "element" }
func (led listElementDescriptor) HasPresence() bool                                 { return true }
func (led listElementDescriptor) IsExtension() bool                                 { return false }
func (led listElementDescriptor) HasOptionalKeyword() bool                          { return false }
func (led listElementDescriptor) IsWeak() bool                                      { return false }
func (led listElementDescriptor) IsPacked() bool                                    { return false }
func (led listElementDescriptor) IsList() bool                                      { return false }
func (led listElementDescriptor) IsMap() bool                                       { return false }
func (led listElementDescriptor) MapKey() protoreflect.FieldDescriptor              { return nil }
func (led listElementDescriptor) MapValue() protoreflect.FieldDescriptor            { return nil }
func (led listElementDescriptor) ContainingOneof() protoreflect.OneofDescriptor     { return nil }
func (led listElementDescriptor) ContainingMessage() protoreflect.MessageDescriptor { return nil }

func (led listElementDescriptor) Default() protoreflect.Value {
	//nolint:exhaustive // missing kinds should be handled by the default block
	switch led.Kind() {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(false)
	case protoreflect.EnumKind:
		if evd := led.DefaultEnumValue(); evd != nil {
			return protoreflect.ValueOfEnum(evd.Number())
		}
		return protoreflect.ValueOfEnum(0)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(0)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(0)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(0)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(0)
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(0)
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(0)
	case protoreflect.StringKind:
		return protoreflect.ValueOfString("")
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes(nil)
	default:
		return protoreflect.Value{}
	}
}

func (led listElementDescriptor) DefaultEnumValue() protoreflect.EnumValueDescriptor {
	if led.Kind() != protoreflect.EnumKind {
		return nil
	}
	enum := led.Enum()
	if enum == nil {
		return nil
	}
	values := enum.Values()
	if values.Len() == 0 {
		return nil
	}
	return values.Get(0)
}

// ListType encapsulates an VirtualListDescriptor with a list implementation.
type ListType interface {
	// New returns a new, empty and mutable instance of empty list.
	New() List

	// Zero returns a new, read-only instance of list.
	Zero() List

	// Descriptor returns the list descriptor.
	//
	// Invariants:
	//   t.Descriptor() == t.New().Descriptor()
	//   t.Descriptor() == t.Zero().Descriptor()
	Descriptor() VirtualListDescriptor
}

type listType struct{ new func(init bool) List }

func NewListType(mt protoreflect.MessageType, listDesc VirtualListDescriptor) ListType {
	return listType{
		new: func(init bool) List {
			if init {
				return NewList(mt, listDesc, mt.Zero().NewField(listDesc.Parent()).List())
			}
			return NewList(mt, listDesc, mt.Zero().Get(listDesc.Parent()).List())
		},
	}
}

func (lt listType) New() List {
	return lt.new(true)
}

func (lt listType) Zero() List {
	return lt.new(false)
}

func (lt listType) Descriptor() VirtualListDescriptor {
	return lt.new(false).Descriptor()
}

// List is an extended reflection interface for a zero-indexed, ordered list. It provides type information and ability to mutate the list or retrieve its children. The element protoreflect.Value type is determined by [VirtualListDescriptor.ListElement].Kind. Providing a protoreflect.Value that is invalid or of an incorrect type panics.
type List interface {
	// Descriptor returns list descriptor, which contains only the protobuf type information for the list.
	Descriptor() VirtualListDescriptor

	// Type returns the list type, which encapsulates both list implementation and protobuf type information. If the list implementation is not needed, it is recommended that the list descriptor be used instead.
	Type() ListType

	// New returns a new, empty and mutable instance of empty list.
	New() List

	// Zero returns a new, read-only instance of list.
	Zero() List

	// Len reports the number of entries in the List.
	// Get, Set, and Truncate panic with out of bound indexes.
	Len() int

	// Range iterates over every item in order from first (zero'th element) to last, calling f for each index and value encountered. Range calls f Len times unless f returns false, which stops iteration. While iterating, mutating operations may only be performed on the current list index.
	Range(f func(int, protoreflect.Value) bool)

	// Get retrieves the value at the given index. It never returns an invalid value.
	Get(i int) protoreflect.Value

	// Set stores a value for the given index. When setting a composite type, it is unspecified whether the set value aliases the source's memory in any way.
	//
	// Set is a mutating operation and unsafe for concurrent use.
	Set(i int, v protoreflect.Value)

	// Append appends the provided value to the end of the list. When appending a composite type, it is unspecified whether the appended value aliases the source's memory in any way.
	//
	// Append is a mutating operation and unsafe for concurrent use.
	Append(v protoreflect.Value)

	// AppendMutable appends a new, empty, mutable message value to the end of the list and returns it. It panics if the list does not contain a message type.
	AppendMutable() protoreflect.Value

	// Truncate truncates the list to a smaller length.
	//
	// Truncate is a mutating operation and unsafe for concurrent use.
	Truncate(len int)

	// NewElement returns a new value of a list element. For enums, this returns the first enum value. For other scalars, this returns the zero value. For messages, this returns a new, empty, mutable value.
	NewElement() protoreflect.Value

	// IsMutable reports whether the list is mutable.
	//
	// An immutable list is an empty, read-only value, for which all content access or mutation operations will panic.
	//
	// Mutability is not part of the protobuf data model, and may not be preserved in marshaling or other operations.
	IsMutable() bool

	// Value returns underlying protoreflect.Value of the list.
	Value() protoreflect.Value
}

type list struct {
	prList protoreflect.List
	desc   VirtualListDescriptor
	new    func(init bool) List
}

// NewList return an extended reflection interface for a protoreflect.List that combines it with type information.
func NewList(mt protoreflect.MessageType, listDesc VirtualListDescriptor, prList protoreflect.List) List {
	return list{
		prList: prList,
		desc:   listDesc,
		new: func(init bool) List {
			if init {
				return NewList(mt, listDesc, mt.Zero().NewField(listDesc.Parent()).List())
			}
			return NewList(mt, listDesc, mt.Zero().Get(listDesc.Parent()).List())
		},
	}
}

func (l list) Descriptor() VirtualListDescriptor {
	return l.desc
}

func (l list) Type() ListType {
	return listType{l.new}
}

func (l list) New() List {
	return l.new(true)
}

func (l list) Zero() List {
	return l.new(false)
}

func (l list) Len() int {
	return l.prList.Len()
}

func (l list) Range(yield func(int, protoreflect.Value) bool) {
	for i, len := 0, l.prList.Len(); i < len; i++ {
		if !yield(i, l.prList.Get(i)) {
			break
		}
	}
}

func (l list) Get(i int) protoreflect.Value {
	return l.prList.Get(i)
}

func (l list) Set(i int, v protoreflect.Value) {
	l.prList.Set(i, v)
}

func (l list) Append(v protoreflect.Value) {
	l.prList.Append(v)
}

func (l list) AppendMutable() protoreflect.Value {
	return l.prList.AppendMutable()
}

func (l list) Truncate(len int) {
	l.prList.Truncate(len)
}

func (l list) NewElement() protoreflect.Value {
	return l.prList.NewElement()
}

func (l list) IsMutable() bool {
	return l.prList.IsValid()
}

func (l list) Value() protoreflect.Value {
	return protoreflect.ValueOfList(l.prList)
}
