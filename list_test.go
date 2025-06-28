package protoreflectextra_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/daishe/protoreflectextra"
	protoreflectextrav1 "github.com/daishe/protoreflectextra/internal/testtypes/protoreflectextra/v1"
)

func TestListDescriptor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
	}{
		{
			name:      "strings",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings",
		},
		{
			name:      "enums",
			base:      &protoreflectextrav1.Message{},
			fieldName: "enums",
		},
		{
			name:      "messages",
			base:      &protoreflectextrav1.Message{},
			fieldName: "messages",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base)
			field := base.ProtoReflect().Descriptor().Fields().ByName(test.fieldName)
			ld := protoreflectextra.ListDescriptorOfField(field)
			lt := protoreflectextra.NewListType(base.ProtoReflect().Type(), ld)
			li := protoreflectextra.GetMessageListField(base.ProtoReflect(), field)
			RequireVirtualListDescriptorsEqual(t, ld, lt.Descriptor(), "VirtualListDescriptor from field and VirtualListDescriptor from list type differs")
			RequireVirtualListDescriptorsEqual(t, ld, li.Descriptor(), "VirtualListDescriptor from field and VirtualListDescriptor from list differs")
			RequireVirtualListDescriptorsEqual(t, ld, li.Type().Descriptor(), "VirtualListDescriptor from field and VirtualListDescriptor from list type from list differs")
		})
	}
}

func RequireVirtualListDescriptorsEqual(t *testing.T, want, got protoreflectextra.VirtualListDescriptor, msgAndArgs ...any) {
	require.Equal(t, want.Parent(), got.Parent(), msgAndArgs...)
	require.Equal(t, want.ListElement(), got.ListElement(), msgAndArgs...)
}

func TestListElementDescriptor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
	}{
		{
			name:      "strings",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings",
		},
		{
			name:      "enums",
			base:      &protoreflectextrav1.Message{},
			fieldName: "enums",
		},
		{
			name:      "messages",
			base:      &protoreflectextrav1.Message{},
			fieldName: "messages",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			elem := protoreflectextra.ListDescriptorOfField(field).ListElement()

			require.Equal(t, field.ParentFile(), elem.ParentFile())
			require.Equal(t, field, elem.Parent())
			require.Equal(t, 0, elem.Index())
			require.Equal(t, field.Syntax(), elem.Syntax())
			require.Equal(t, field.Name()+".element", elem.Name())
			require.Equal(t, field.FullName()+".element", elem.FullName())
			require.Equal(t, field.IsPlaceholder(), elem.IsPlaceholder())
			require.Equal(t, nil, elem.Options())
			require.Equal(t, field.Number(), elem.Number())
			require.Equal(t, protoreflect.Optional, elem.Cardinality())
			require.Equal(t, field.Kind(), elem.Kind())
			require.Equal(t, false, elem.HasJSONName())
			require.Equal(t, "element", elem.JSONName())
			require.Equal(t, "element", elem.TextName())
			require.Equal(t, true, elem.HasPresence())
			require.Equal(t, false, elem.IsExtension())
			require.Equal(t, false, elem.HasOptionalKeyword())
			require.Equal(t, false, elem.IsWeak())
			require.Equal(t, false, elem.IsPacked())
			require.Equal(t, false, elem.IsList())
			require.Equal(t, false, elem.IsMap())
			require.Equal(t, nil, elem.MapKey())
			require.Equal(t, nil, elem.MapValue())
			require.Equal(t, false, elem.HasDefault())
			require.Equal(t, nil, elem.ContainingOneof())
			require.Equal(t, nil, elem.ContainingMessage())
			require.Equal(t, field.Enum(), elem.Enum())
			require.Equal(t, field.Message(), elem.Message())

			if elem.Kind() == protoreflect.EnumKind {
				require.NotNil(t, elem.DefaultEnumValue())
				require.NotNil(t, elem.Enum())
			} else {
				require.Nil(t, elem.DefaultEnumValue())
				require.Nil(t, elem.Enum())
			}
			if elem.Kind() == protoreflect.MessageKind {
				require.NotNil(t, elem.Message())
				require.Equal(t, protoreflect.Value{}.Interface(), elem.Default().Interface())
			} else {
				require.Nil(t, elem.Message())
				require.Equal(t, base.Get(field).List().NewElement().Interface(), elem.Default().Interface())
			}
		})
	}
}

func TestListElementDescriptorWithMock(t *testing.T) {
	t.Parallel()
	base := (&protoreflectextrav1.Message{}).ProtoReflect()
	baseField := base.Descriptor().Fields().ByName("string")
	listField := protoreflect.FieldDescriptor(mockFieldDescriptor_IsList{FieldDescriptor: baseField, MockIsList: true})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Kind(0)})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, nil, elem.Default().Interface())
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.BoolKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, false, elem.Default().Interface())
	})

	t.Run("enum/withoutEnumDescriptor", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.EnumKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, protoreflect.EnumNumber(0), elem.Default().Interface())
	})

	t.Run("enum/withEnumDescriptor", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.EnumKind})
		mockField = protoreflect.FieldDescriptor(mockFieldDescriptor_Enum{FieldDescriptor: mockField, MockEnum: protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED.Descriptor()})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, protoreflect.EnumNumber(0), elem.Default().Interface())
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Int32Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, int32(0), elem.Default().Interface())
	})

	t.Run("sint32", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Sint32Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, int32(0), elem.Default().Interface())
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Uint32Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, uint32(0), elem.Default().Interface())
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Int64Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, int64(0), elem.Default().Interface())
	})

	t.Run("sint64", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Sint64Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, int64(0), elem.Default().Interface())
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Uint64Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, uint64(0), elem.Default().Interface())
	})

	t.Run("sfixed32", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Sfixed32Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, int32(0), elem.Default().Interface())
	})

	t.Run("fixed32", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Fixed32Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, uint32(0), elem.Default().Interface())
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.FloatKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, float32(0), elem.Default().Interface())
	})

	t.Run("sfixed64", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Sfixed64Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, int64(0), elem.Default().Interface())
	})

	t.Run("fixed64", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.Fixed64Kind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, uint64(0), elem.Default().Interface())
	})

	t.Run("double", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.DoubleKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, float64(0), elem.Default().Interface())
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.StringKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, "", elem.Default().Interface())
	})

	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.BytesKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, []byte(nil), elem.Default().Interface())
	})

	t.Run("message", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.MessageKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, nil, elem.Default().Interface())
	})

	t.Run("group", func(t *testing.T) {
		t.Parallel()
		mockField := protoreflect.FieldDescriptor(mockFieldDescriptor_Kind{FieldDescriptor: listField, MockKind: protoreflect.GroupKind})
		elem := protoreflectextra.ListDescriptorOfField(mockField).ListElement()
		require.Equal(t, nil, elem.Default().Interface())
	})
}

func TestListDescriptorOfFieldPanic(t *testing.T) {
	t.Parallel()
	base := &protoreflectextrav1.Message{}
	field := base.ProtoReflect().Descriptor().Fields().ByName("string")
	require.Panics(t, func() {
		_ = protoreflectextra.ListDescriptorOfField(field)
	})
}

func TestListType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
		want      proto.Message
	}{
		{
			name:      "strings",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings",
			want: &protoreflectextrav1.Message{
				Strings: []string{""},
			},
		},
		{
			name:      "enums",
			base:      &protoreflectextrav1.Message{},
			fieldName: "enums",
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "messages",
			base:      &protoreflectextrav1.Message{},
			fieldName: "messages",
			want: &protoreflectextrav1.Message{
				Messages: []*protoreflectextrav1.Message{{}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			lt := protoreflectextra.NewListType(base.Type(), protoreflectextra.ListDescriptorOfField(field))

			zero := lt.Zero()
			require.False(t, zero.IsMutable())

			new := lt.New()
			require.True(t, new.IsMutable())

			new.Append(new.NewElement())
			base.Set(field, new.Value())
			RequireMessageEqual(t, test.want, base.Interface())
		})
	}
}

func TestList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
		want      proto.Message
	}{
		{
			name:      "strings",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings",
			want: &protoreflectextrav1.Message{
				Strings: []string{""},
			},
		},
		{
			name:      "enums",
			base:      &protoreflectextrav1.Message{},
			fieldName: "enums",
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "messages",
			base:      &protoreflectextrav1.Message{},
			fieldName: "messages",
			want: &protoreflectextrav1.Message{
				Messages: []*protoreflectextrav1.Message{{}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			li := protoreflectextra.NewList(base.Type(), protoreflectextra.ListDescriptorOfField(field), base.Get(field).List())
			require.False(t, li.IsMutable())

			zero := li.Zero()
			require.False(t, zero.IsMutable())

			new := li.New()
			require.True(t, new.IsMutable())

			new.Append(new.NewElement())
			base.Set(field, new.Value())
			RequireMessageEqual(t, test.want, base.Interface())
		})
	}
}

func TestListOps(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
		value     protoreflect.Value
		want      proto.Message
	}{
		{
			name:      "strings",
			fieldName: "strings",
			base:      &protoreflectextrav1.Message{},
			value:     protoreflect.ValueOf("bbb"),
			want: &protoreflectextrav1.Message{
				Strings: []string{"bbb"},
			},
		},
		{
			name:      "enums",
			fieldName: "enums",
			base:      &protoreflectextrav1.Message{},
			value:     protoreflect.ValueOf(protoreflectextrav1.Enum_ENUM_VALUE_OTHER.Number()),
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_OTHER},
			},
		},
		{
			name:      "messages",
			fieldName: "messages",
			base:      &protoreflectextrav1.Message{},
			value:     protoreflect.ValueOf((&protoreflectextrav1.Message{String_: "zzz"}).ProtoReflect()),
			want: &protoreflectextrav1.Message{
				Messages: []*protoreflectextrav1.Message{{String_: "zzz"}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			prList := base.Mutable(field).List()
			li := protoreflectextra.NewList(base.Type(), protoreflectextra.ListDescriptorOfField(field), prList)
			emptyValue := li.NewElement()

			require.Equal(t, true, li.IsMutable())

			require.Equal(t, 0, li.Len())

			li.Append(emptyValue)
			require.Equal(t, 1, li.Len())
			require.Equal(t, emptyValue.Interface(), li.Get(0).Interface())
			for i, v := range li.Range {
				require.Equal(t, 0, i)
				require.Equal(t, emptyValue.Interface(), v.Interface())
			}

			li.Set(0, test.value)
			require.Equal(t, 1, li.Len())
			require.Equal(t, test.value.Interface(), li.Get(0).Interface())
			for i, v := range li.Range {
				require.Equal(t, 0, i)
				require.Equal(t, test.value.Interface(), v.Interface())
			}

			if li.Descriptor().ListElement().Kind() == protoreflect.MessageKind {
				li.AppendMutable()
			} else {
				li.Append(emptyValue)
			}
			require.Equal(t, 2, li.Len())
			require.Equal(t, test.value.Interface(), li.Get(0).Interface())
			require.Equal(t, emptyValue.Interface(), li.Get(1).Interface())
			for i, v := range li.Range {
				require.Equal(t, 0, i)
				require.Equal(t, test.value.Interface(), v.Interface())
				break
			}

			li.Truncate(1)
			require.Equal(t, 1, li.Len())
			require.Equal(t, test.value.Interface(), li.Get(0).Interface())
			for i, v := range li.Range {
				require.Equal(t, 0, i)
				require.Equal(t, test.value.Interface(), v.Interface())
			}

			RequireMessageEqual(t, test.want, base.Interface())
		})
	}
}
