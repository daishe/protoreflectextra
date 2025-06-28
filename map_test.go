package protoreflectextra_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/daishe/protoreflectextra"
	protoreflectextrav1 "github.com/daishe/protoreflectextra/internal/testtypes/protoreflectextra/v1"
)

func TestMapDescriptor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
	}{
		{
			name:      "strings_to_strings",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_strings",
		},
		{
			name:      "strings_to_enums",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_enums",
		},
		{
			name:      "strings_to_messages",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_messages",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base)
			field := base.ProtoReflect().Descriptor().Fields().ByName(test.fieldName)
			md := protoreflectextra.MapDescriptorOfField(field)
			mt := protoreflectextra.NewMapType(base.ProtoReflect().Type(), md)
			ma := protoreflectextra.GetMessageMapField(base.ProtoReflect(), field)
			RequireVirtualMapDescriptorsEqual(t, md, mt.Descriptor(), "VirtualMapDescriptor from field and VirtualMapDescriptor from map type differs")
			RequireVirtualMapDescriptorsEqual(t, md, ma.Descriptor(), "VirtualMapDescriptor from field and VirtualMapDescriptor from map differs")
			RequireVirtualMapDescriptorsEqual(t, md, ma.Type().Descriptor(), "VirtualMapDescriptor from field and VirtualMapDescriptor from map type from map differs")
		})
	}
}

func RequireVirtualMapDescriptorsEqual(t *testing.T, want, got protoreflectextra.VirtualMapDescriptor, msgAndArgs ...any) {
	require.Equal(t, want.Parent(), got.Parent(), msgAndArgs...)
	require.Equal(t, want.MapKey(), got.MapKey(), msgAndArgs...)
	require.Equal(t, want.MapValue(), got.MapValue(), msgAndArgs...)
}

func TestMapDescriptorOfFieldPanic(t *testing.T) {
	t.Parallel()
	base := &protoreflectextrav1.Message{}
	field := base.ProtoReflect().Descriptor().Fields().ByName("string")
	require.Panics(t, func() {
		_ = protoreflectextra.MapDescriptorOfField(field)
	})
}

func TestMapType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
		want      proto.Message
	}{
		{
			name:      "strings_to_strings",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_strings",
			want: &protoreflectextrav1.Message{
				StringsToStrings: map[string]string{"": ""},
			},
		},
		{
			name:      "strings_to_enums",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_enums",
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "strings_to_messages",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_messages",
			want: &protoreflectextrav1.Message{
				StringsToMessages: map[string]*protoreflectextrav1.Message{"": {}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			mt := protoreflectextra.NewMapType(base.Type(), protoreflectextra.MapDescriptorOfField(field))

			zero := mt.Zero()
			require.False(t, zero.IsMutable())

			new := mt.New()
			require.True(t, new.IsMutable())

			new.Set(new.NewKey(), new.NewValue())
			base.Set(field, new.Value())
			RequireMessageEqual(t, test.want, base.Interface())
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
		want      proto.Message
	}{
		{
			name:      "strings_to_strings",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_strings",
			want: &protoreflectextrav1.Message{
				StringsToStrings: map[string]string{"": ""},
			},
		},
		{
			name:      "strings_to_enums",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_enums",
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "strings_to_messages",
			base:      &protoreflectextrav1.Message{},
			fieldName: "strings_to_messages",
			want: &protoreflectextrav1.Message{
				StringsToMessages: map[string]*protoreflectextrav1.Message{"": {}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			ma := protoreflectextra.NewMap(base.Type(), protoreflectextra.MapDescriptorOfField(field), base.Get(field).Map())
			require.False(t, ma.IsMutable())

			zero := ma.Zero()
			require.False(t, zero.IsMutable())

			new := ma.New()
			require.True(t, new.IsMutable())

			new.Set(new.NewKey(), new.NewValue())
			base.Set(field, new.Value())
			RequireMessageEqual(t, test.want, base.Interface())
		})
	}
}

func TestMapOps(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
		key       protoreflect.MapKey
		value     protoreflect.Value
		want      proto.Message
	}{
		{
			name:      "strings_to_strings",
			fieldName: "strings_to_strings",
			base:      &protoreflectextrav1.Message{},
			key:       protoreflect.ValueOf("aaa").MapKey(),
			value:     protoreflect.ValueOf("bbb"),
			want: &protoreflectextrav1.Message{
				StringsToStrings: map[string]string{"aaa": "bbb"},
			},
		},
		{
			name:      "strings_to_enums",
			fieldName: "strings_to_enums",
			base:      &protoreflectextrav1.Message{},
			key:       protoreflect.ValueOf("aaa").MapKey(),
			value:     protoreflect.ValueOf(protoreflectextrav1.Enum_ENUM_VALUE_OTHER.Number()),
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"aaa": protoreflectextrav1.Enum_ENUM_VALUE_OTHER},
			},
		},
		{
			name:      "strings_to_messages",
			fieldName: "strings_to_messages",
			base:      &protoreflectextrav1.Message{},
			key:       protoreflect.ValueOf("aaa").MapKey(),
			value:     protoreflect.ValueOf((&protoreflectextrav1.Message{String_: "zzz"}).ProtoReflect()),
			want: &protoreflectextrav1.Message{
				StringsToMessages: map[string]*protoreflectextrav1.Message{"aaa": {String_: "zzz"}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			prMap := base.Mutable(field).Map()
			ma := protoreflectextra.NewMap(base.Type(), protoreflectextra.MapDescriptorOfField(field), prMap)
			emptyKey, emptyValue := ma.NewKey(), ma.NewValue()

			require.Equal(t, field.MapKey().Default().Interface(), emptyKey.Interface())
			require.Equal(t, prMap.NewValue().Interface(), emptyValue.Interface())
			require.Equal(t, true, ma.IsMutable())

			require.Equal(t, 0, ma.Len())
			require.Equal(t, false, ma.Has(test.key))
			require.Equal(t, false, ma.Get(test.key).IsValid())
			require.Equal(t, false, ma.Has(emptyKey))
			require.Equal(t, false, ma.Get(emptyKey).IsValid())

			ma.Set(test.key, test.value)
			require.Equal(t, 1, ma.Len())
			require.Equal(t, true, ma.Has(test.key))
			require.Equal(t, test.value.Interface(), ma.Get(test.key).Interface())
			require.Equal(t, false, ma.Has(emptyKey))
			require.Equal(t, false, ma.Get(emptyKey).IsValid())
			for k, v := range ma.Range {
				require.Equal(t, test.key.Interface(), k.Interface())
				require.Equal(t, test.value.Interface(), v.Interface())
			}

			if ma.Descriptor().MapValue().Kind() == protoreflect.MessageKind {
				ma.Mutable(emptyKey)
			} else {
				ma.Set(emptyKey, emptyValue)
			}
			require.Equal(t, 2, ma.Len())
			require.Equal(t, true, ma.Has(test.key))
			require.Equal(t, test.value.Interface(), ma.Get(test.key).Interface())
			require.Equal(t, true, ma.Has(emptyKey))
			require.Equal(t, emptyValue.Interface(), ma.Get(emptyKey).Interface())

			ma.Clear(test.key)
			require.Equal(t, 1, ma.Len())
			require.Equal(t, false, ma.Has(test.key))
			require.Equal(t, false, ma.Get(test.key).IsValid())
			require.Equal(t, true, ma.Has(emptyKey))
			require.Equal(t, emptyValue.Interface(), ma.Get(emptyKey).Interface())
			for k, v := range ma.Range {
				require.Equal(t, emptyKey.Interface(), k.Interface())
				require.Equal(t, emptyValue.Interface(), v.Interface())
			}

			ma.Clear(emptyKey)
			require.Equal(t, 0, ma.Len())
			require.Equal(t, false, ma.Has(test.key))
			require.Equal(t, false, ma.Get(test.key).IsValid())
			require.Equal(t, false, ma.Has(emptyKey))
			require.Equal(t, false, ma.Get(emptyKey).IsValid())

			ma.Set(test.key, test.value)
			RequireMessageEqual(t, test.want, base.Interface())
		})
	}
}
