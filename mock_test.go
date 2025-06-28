package protoreflectextra_test

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type mockFieldDescriptor_Kind struct {
	protoreflect.FieldDescriptor

	MockKind protoreflect.Kind
}

func (mock mockFieldDescriptor_Kind) Kind() protoreflect.Kind { return mock.MockKind }

type mockFieldDescriptor_IsList struct {
	protoreflect.FieldDescriptor

	MockIsList bool
}

func (mock mockFieldDescriptor_IsList) IsList() bool { return mock.MockIsList }

type mockFieldDescriptor_Enum struct {
	protoreflect.FieldDescriptor

	MockEnum protoreflect.EnumDescriptor
}

func (mock mockFieldDescriptor_Enum) Enum() protoreflect.EnumDescriptor { return mock.MockEnum }
