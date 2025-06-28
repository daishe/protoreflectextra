package protoreflectextra_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/daishe/protoreflectextra"
	protoreflectextrav1 "github.com/daishe/protoreflectextra/internal/testtypes/protoreflectextra/v1"
)

func TestValueAccessors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		base      proto.Message
		fieldName protoreflect.Name
		mutate    func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor)
		want      proto.Message
	}{
		{
			name:      "NewMessageListField/strings",
			fieldName: "strings",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.NewMessageListField(base, field)
				li.Append(li.NewElement())
				base.Set(field, li.Value())
			},
			want: &protoreflectextrav1.Message{
				Strings: []string{""},
			},
		},
		{
			name:      "NewMessageListField/enums",
			fieldName: "enums",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.NewMessageListField(base, field)
				li.Append(li.NewElement())
				base.Set(field, li.Value())
			},
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "NewMessageListField/messages",
			fieldName: "messages",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.NewMessageListField(base, field)
				li.Append(li.NewElement())
				base.Set(field, li.Value())
			},
			want: &protoreflectextrav1.Message{
				Messages: []*protoreflectextrav1.Message{{}},
			},
		},
		{
			name:      "GetMessageListField/read-only/strings",
			fieldName: "strings",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.GetMessageListField(base, field)
				require.False(t, li.IsMutable())
			},
			want: &protoreflectextrav1.Message{},
		},
		{
			name:      "GetMessageListField/read-only/enums",
			fieldName: "enums",
			base: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.GetMessageListField(base, field)
				require.False(t, li.IsMutable())
			},
			want: &protoreflectextrav1.Message{},
		},
		{
			name:      "GetMessageListField/read-only/messages",
			fieldName: "messages",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.GetMessageListField(base, field)
				require.False(t, li.IsMutable())
			},
			want: &protoreflectextrav1.Message{},
		},
		{
			name:      "GetMessageListField/strings",
			fieldName: "strings",
			base: &protoreflectextrav1.Message{
				Strings: []string{"aaa"},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.GetMessageListField(base, field)
				li.Append(li.NewElement())
			},
			want: &protoreflectextrav1.Message{
				Strings: []string{"aaa", ""},
			},
		},
		{
			name:      "GetMessageListField/enums",
			fieldName: "enums",
			base: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_OTHER},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.GetMessageListField(base, field)
				li.Append(li.NewElement())
			},
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_OTHER, protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "GetMessageListField/messages",
			fieldName: "messages",
			base: &protoreflectextrav1.Message{
				Messages: []*protoreflectextrav1.Message{{String_: "zzz"}},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.GetMessageListField(base, field)
				li.Append(li.NewElement())
			},
			want: &protoreflectextrav1.Message{
				Messages: []*protoreflectextrav1.Message{{String_: "zzz"}, {}},
			},
		},
		{
			name:      "MutableMessageListField/strings",
			fieldName: "strings",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.MutableMessageListField(base, field)
				li.Append(li.NewElement())
			},
			want: &protoreflectextrav1.Message{
				Strings: []string{""},
			},
		},
		{
			name:      "MutableMessageListField/enums",
			fieldName: "enums",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.MutableMessageListField(base, field)
				li.Append(li.NewElement())
			},
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "MutableMessageListField/messages",
			fieldName: "messages",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.MutableMessageListField(base, field)
				li.Append(li.NewElement())
			},
			want: &protoreflectextrav1.Message{
				Messages: []*protoreflectextrav1.Message{{}},
			},
		},

		{
			name:      "NewMessageMapField/strings",
			fieldName: "strings_to_strings",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.NewMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
				base.Set(field, ma.Value())
			},
			want: &protoreflectextrav1.Message{
				StringsToStrings: map[string]string{"": ""},
			},
		},
		{
			name:      "NewMessageMapField/enums",
			fieldName: "strings_to_enums",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.NewMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
				base.Set(field, ma.Value())
			},
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "NewMessageMapField/messages",
			fieldName: "strings_to_messages",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.NewMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
				base.Set(field, ma.Value())
			},
			want: &protoreflectextrav1.Message{
				StringsToMessages: map[string]*protoreflectextrav1.Message{"": {}},
			},
		},
		{
			name:      "GetMessageMapField/read-only/strings",
			fieldName: "strings_to_strings",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.GetMessageMapField(base, field)
				require.False(t, ma.IsMutable())
			},
			want: &protoreflectextrav1.Message{},
		},
		{
			name:      "GetMessageMapField/read-only/enums",
			fieldName: "strings_to_enums",
			base: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.GetMessageMapField(base, field)
				require.False(t, ma.IsMutable())
			},
			want: &protoreflectextrav1.Message{},
		},
		{
			name:      "GetMessageMapField/read-only/messages",
			fieldName: "strings_to_messages",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.GetMessageMapField(base, field)
				require.False(t, ma.IsMutable())
			},
			want: &protoreflectextrav1.Message{},
		},
		{
			name:      "GetMessageMapField/strings",
			fieldName: "strings_to_strings",
			base: &protoreflectextrav1.Message{
				StringsToStrings: map[string]string{"aaa": "bbb"},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.GetMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
			},
			want: &protoreflectextrav1.Message{
				StringsToStrings: map[string]string{"aaa": "bbb", "": ""},
			},
		},
		{
			name:      "GetMessageMapField/enums",
			fieldName: "strings_to_enums",
			base: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"aaa": protoreflectextrav1.Enum_ENUM_VALUE_OTHER},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.GetMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
			},
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"aaa": protoreflectextrav1.Enum_ENUM_VALUE_OTHER, "": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "GetMessageMapField/messages",
			fieldName: "strings_to_messages",
			base: &protoreflectextrav1.Message{
				StringsToMessages: map[string]*protoreflectextrav1.Message{"aaa": {String_: "zzz"}},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.GetMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
			},
			want: &protoreflectextrav1.Message{
				StringsToMessages: map[string]*protoreflectextrav1.Message{"aaa": {String_: "zzz"}, "": {}},
			},
		},
		{
			name:      "MutableMessageMapField/strings",
			fieldName: "strings_to_strings",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.MutableMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
			},
			want: &protoreflectextrav1.Message{
				StringsToStrings: map[string]string{"": ""},
			},
		},
		{
			name:      "MutableMessageMapField/enums",
			fieldName: "strings_to_enums",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.MutableMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
			},
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "MutableMessageMapField/messages",
			fieldName: "strings_to_messages",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.MutableMessageMapField(base, field)
				ma.Set(ma.NewKey(), ma.NewValue())
			},
			want: &protoreflectextrav1.Message{
				StringsToMessages: map[string]*protoreflectextrav1.Message{"": {}},
			},
		},

		{
			name:      "NewMessageEnumField/enum",
			fieldName: "enum_option",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				e := protoreflectextra.NewMessageEnumField(base, field)
				base.Set(field, protoreflect.ValueOfEnum(e.Number()+1))
			},
			want: &protoreflectextrav1.Message{
				Oneof: &protoreflectextrav1.Message_EnumOption{
					EnumOption: protoreflectextrav1.Enum_ENUM_VALUE_OTHER,
				},
			},
		},
		{
			name:      "GetMessageEnumField/read-only/enum",
			fieldName: "enum_option",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				e := protoreflectextra.GetMessageEnumField(base, field)
				require.Equal(t, field.Default().Enum(), e.Number())
			},
			want: &protoreflectextrav1.Message{},
		},
		{
			name:      "GetMessageEnumField/enum",
			fieldName: "enum_option",
			base: &protoreflectextrav1.Message{
				Oneof: &protoreflectextrav1.Message_EnumOption{
					EnumOption: protoreflectextrav1.Enum_ENUM_VALUE_OTHER,
				},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				e := protoreflectextra.GetMessageEnumField(base, field)
				require.Equal(t, field.Default().Enum()+1, e.Number())
			},
			want: &protoreflectextrav1.Message{
				Oneof: &protoreflectextrav1.Message_EnumOption{
					EnumOption: protoreflectextrav1.Enum_ENUM_VALUE_OTHER,
				},
			},
		},

		{
			name:      "NewListEnumElement/enums",
			fieldName: "enums",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.MutableMessageListField(base, field)
				e := protoreflectextra.NewListEnumElement(li)
				li.Append(protoreflect.ValueOfEnum(e.Number()))
			},
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "GetListEnumElement/enums",
			fieldName: "enums",
			base: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				li := protoreflectextra.GetMessageListField(base, field)
				e := protoreflectextra.GetListEnumElement(li, 0)
				require.Equal(t, li.Descriptor().ListElement().Default().Enum(), e.Number())
			},
			want: &protoreflectextrav1.Message{
				Enums: []protoreflectextrav1.Enum{protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},

		{
			name:      "NewMapEnumValue/strings_to_enums",
			fieldName: "strings_to_enums",
			base:      &protoreflectextrav1.Message{},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.MutableMessageMapField(base, field)
				e := protoreflectextra.NewMapEnumValue(ma)
				ma.Set(ma.NewKey(), protoreflect.ValueOfEnum(e.Number()))
			},
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
		{
			name:      "GetMapEnumValue/strings_to_enums",
			fieldName: "strings_to_enums",
			base: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
			mutate: func(t *testing.T, base protoreflect.Message, field protoreflect.FieldDescriptor) {
				ma := protoreflectextra.GetMessageMapField(base, field)
				e := protoreflectextra.GetMapEnumValue(ma, ma.NewKey())
				require.Equal(t, ma.Descriptor().MapValue().Default().Enum(), e.Number())
			},
			want: &protoreflectextrav1.Message{
				StringsToEnums: map[string]protoreflectextrav1.Enum{"": protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := proto.Clone(test.base).ProtoReflect()
			field := base.Descriptor().Fields().ByName(test.fieldName)
			test.mutate(t, base, field)
			RequireMessageEqual(t, test.want, base.Interface())
		})
	}
}
