// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: protoreflectextra/v1/types.proto

package protoreflectextrav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum int32

const (
	Enum_ENUM_VALUE_UNSPECIFIED Enum = 0
	Enum_ENUM_VALUE_OTHER       Enum = 1
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ENUM_VALUE_UNSPECIFIED",
		1: "ENUM_VALUE_OTHER",
	}
	Enum_value = map[string]int32{
		"ENUM_VALUE_UNSPECIFIED": 0,
		"ENUM_VALUE_OTHER":       1,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_protoreflectextra_v1_types_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_protoreflectextra_v1_types_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_protoreflectextra_v1_types_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	String_           string                 `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	Strings           []string               `protobuf:"bytes,2,rep,name=strings,proto3" json:"strings,omitempty"`
	StringsToStrings  map[string]string      `protobuf:"bytes,3,rep,name=strings_to_strings,json=stringsToStrings,proto3" json:"strings_to_strings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Enum              Enum                   `protobuf:"varint,21,opt,name=enum,proto3,enum=protoreflectextra.v1.Enum" json:"enum,omitempty"`
	Enums             []Enum                 `protobuf:"varint,22,rep,packed,name=enums,proto3,enum=protoreflectextra.v1.Enum" json:"enums,omitempty"`
	StringsToEnums    map[string]Enum        `protobuf:"bytes,23,rep,name=strings_to_enums,json=stringsToEnums,proto3" json:"strings_to_enums,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=protoreflectextra.v1.Enum"`
	Message           *Message               `protobuf:"bytes,31,opt,name=message,proto3" json:"message,omitempty"`
	Messages          []*Message             `protobuf:"bytes,32,rep,name=messages,proto3" json:"messages,omitempty"`
	StringsToMessages map[string]*Message    `protobuf:"bytes,33,rep,name=strings_to_messages,json=stringsToMessages,proto3" json:"strings_to_messages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Oneof:
	//
	//	*Message_StringOption
	//	*Message_EnumOption
	//	*Message_MessageOption
	Oneof         isMessage_Oneof `protobuf_oneof:"oneof"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_protoreflectextra_v1_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protoreflectextra_v1_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_protoreflectextra_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Message) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

func (x *Message) GetStringsToStrings() map[string]string {
	if x != nil {
		return x.StringsToStrings
	}
	return nil
}

func (x *Message) GetEnum() Enum {
	if x != nil {
		return x.Enum
	}
	return Enum_ENUM_VALUE_UNSPECIFIED
}

func (x *Message) GetEnums() []Enum {
	if x != nil {
		return x.Enums
	}
	return nil
}

func (x *Message) GetStringsToEnums() map[string]Enum {
	if x != nil {
		return x.StringsToEnums
	}
	return nil
}

func (x *Message) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Message) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Message) GetStringsToMessages() map[string]*Message {
	if x != nil {
		return x.StringsToMessages
	}
	return nil
}

func (x *Message) GetOneof() isMessage_Oneof {
	if x != nil {
		return x.Oneof
	}
	return nil
}

func (x *Message) GetStringOption() string {
	if x != nil {
		if x, ok := x.Oneof.(*Message_StringOption); ok {
			return x.StringOption
		}
	}
	return ""
}

func (x *Message) GetEnumOption() Enum {
	if x != nil {
		if x, ok := x.Oneof.(*Message_EnumOption); ok {
			return x.EnumOption
		}
	}
	return Enum_ENUM_VALUE_UNSPECIFIED
}

func (x *Message) GetMessageOption() *Message {
	if x != nil {
		if x, ok := x.Oneof.(*Message_MessageOption); ok {
			return x.MessageOption
		}
	}
	return nil
}

type isMessage_Oneof interface {
	isMessage_Oneof()
}

type Message_StringOption struct {
	StringOption string `protobuf:"bytes,41,opt,name=string_option,json=stringOption,proto3,oneof"`
}

type Message_EnumOption struct {
	EnumOption Enum `protobuf:"varint,42,opt,name=enum_option,json=enumOption,proto3,enum=protoreflectextra.v1.Enum,oneof"`
}

type Message_MessageOption struct {
	MessageOption *Message `protobuf:"bytes,43,opt,name=message_option,json=messageOption,proto3,oneof"`
}

func (*Message_StringOption) isMessage_Oneof() {}

func (*Message_EnumOption) isMessage_Oneof() {}

func (*Message_MessageOption) isMessage_Oneof() {}

var File_protoreflectextra_v1_types_proto protoreflect.FileDescriptor

const file_protoreflectextra_v1_types_proto_rawDesc = "" +
	"\n" +
	" protoreflectextra/v1/types.proto\x12\x14protoreflectextra.v1\"\xf7\a\n" +
	"\aMessage\x12\x16\n" +
	"\x06string\x18\x01 \x01(\tR\x06string\x12\x18\n" +
	"\astrings\x18\x02 \x03(\tR\astrings\x12a\n" +
	"\x12strings_to_strings\x18\x03 \x03(\v23.protoreflectextra.v1.Message.StringsToStringsEntryR\x10stringsToStrings\x12.\n" +
	"\x04enum\x18\x15 \x01(\x0e2\x1a.protoreflectextra.v1.EnumR\x04enum\x120\n" +
	"\x05enums\x18\x16 \x03(\x0e2\x1a.protoreflectextra.v1.EnumR\x05enums\x12[\n" +
	"\x10strings_to_enums\x18\x17 \x03(\v21.protoreflectextra.v1.Message.StringsToEnumsEntryR\x0estringsToEnums\x127\n" +
	"\amessage\x18\x1f \x01(\v2\x1d.protoreflectextra.v1.MessageR\amessage\x129\n" +
	"\bmessages\x18  \x03(\v2\x1d.protoreflectextra.v1.MessageR\bmessages\x12d\n" +
	"\x13strings_to_messages\x18! \x03(\v24.protoreflectextra.v1.Message.StringsToMessagesEntryR\x11stringsToMessages\x12%\n" +
	"\rstring_option\x18) \x01(\tH\x00R\fstringOption\x12=\n" +
	"\venum_option\x18* \x01(\x0e2\x1a.protoreflectextra.v1.EnumH\x00R\n" +
	"enumOption\x12F\n" +
	"\x0emessage_option\x18+ \x01(\v2\x1d.protoreflectextra.v1.MessageH\x00R\rmessageOption\x1aC\n" +
	"\x15StringsToStringsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a]\n" +
	"\x13StringsToEnumsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x120\n" +
	"\x05value\x18\x02 \x01(\x0e2\x1a.protoreflectextra.v1.EnumR\x05value:\x028\x01\x1ac\n" +
	"\x16StringsToMessagesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x123\n" +
	"\x05value\x18\x02 \x01(\v2\x1d.protoreflectextra.v1.MessageR\x05value:\x028\x01B\a\n" +
	"\x05oneof*8\n" +
	"\x04Enum\x12\x1a\n" +
	"\x16ENUM_VALUE_UNSPECIFIED\x10\x00\x12\x14\n" +
	"\x10ENUM_VALUE_OTHER\x10\x01B\xf8\x01\n" +
	"\x18com.protoreflectextra.v1B\n" +
	"TypesProtoP\x01Z_github.com/daishe/protoreflectextra/internal/testtypes/protoreflectextra/v1;protoreflectextrav1\xa2\x02\x03PXX\xaa\x02\x14Protoreflectextra.V1\xca\x02\x14Protoreflectextra\\V1\xe2\x02 Protoreflectextra\\V1\\GPBMetadata\xea\x02\x15Protoreflectextra::V1b\x06proto3"

var (
	file_protoreflectextra_v1_types_proto_rawDescOnce sync.Once
	file_protoreflectextra_v1_types_proto_rawDescData []byte
)

func file_protoreflectextra_v1_types_proto_rawDescGZIP() []byte {
	file_protoreflectextra_v1_types_proto_rawDescOnce.Do(func() {
		file_protoreflectextra_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protoreflectextra_v1_types_proto_rawDesc), len(file_protoreflectextra_v1_types_proto_rawDesc)))
	})
	return file_protoreflectextra_v1_types_proto_rawDescData
}

var file_protoreflectextra_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protoreflectextra_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protoreflectextra_v1_types_proto_goTypes = []any{
	(Enum)(0),       // 0: protoreflectextra.v1.Enum
	(*Message)(nil), // 1: protoreflectextra.v1.Message
	nil,             // 2: protoreflectextra.v1.Message.StringsToStringsEntry
	nil,             // 3: protoreflectextra.v1.Message.StringsToEnumsEntry
	nil,             // 4: protoreflectextra.v1.Message.StringsToMessagesEntry
}
var file_protoreflectextra_v1_types_proto_depIdxs = []int32{
	2,  // 0: protoreflectextra.v1.Message.strings_to_strings:type_name -> protoreflectextra.v1.Message.StringsToStringsEntry
	0,  // 1: protoreflectextra.v1.Message.enum:type_name -> protoreflectextra.v1.Enum
	0,  // 2: protoreflectextra.v1.Message.enums:type_name -> protoreflectextra.v1.Enum
	3,  // 3: protoreflectextra.v1.Message.strings_to_enums:type_name -> protoreflectextra.v1.Message.StringsToEnumsEntry
	1,  // 4: protoreflectextra.v1.Message.message:type_name -> protoreflectextra.v1.Message
	1,  // 5: protoreflectextra.v1.Message.messages:type_name -> protoreflectextra.v1.Message
	4,  // 6: protoreflectextra.v1.Message.strings_to_messages:type_name -> protoreflectextra.v1.Message.StringsToMessagesEntry
	0,  // 7: protoreflectextra.v1.Message.enum_option:type_name -> protoreflectextra.v1.Enum
	1,  // 8: protoreflectextra.v1.Message.message_option:type_name -> protoreflectextra.v1.Message
	0,  // 9: protoreflectextra.v1.Message.StringsToEnumsEntry.value:type_name -> protoreflectextra.v1.Enum
	1,  // 10: protoreflectextra.v1.Message.StringsToMessagesEntry.value:type_name -> protoreflectextra.v1.Message
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_protoreflectextra_v1_types_proto_init() }
func file_protoreflectextra_v1_types_proto_init() {
	if File_protoreflectextra_v1_types_proto != nil {
		return
	}
	file_protoreflectextra_v1_types_proto_msgTypes[0].OneofWrappers = []any{
		(*Message_StringOption)(nil),
		(*Message_EnumOption)(nil),
		(*Message_MessageOption)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protoreflectextra_v1_types_proto_rawDesc), len(file_protoreflectextra_v1_types_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoreflectextra_v1_types_proto_goTypes,
		DependencyIndexes: file_protoreflectextra_v1_types_proto_depIdxs,
		EnumInfos:         file_protoreflectextra_v1_types_proto_enumTypes,
		MessageInfos:      file_protoreflectextra_v1_types_proto_msgTypes,
	}.Build()
	File_protoreflectextra_v1_types_proto = out.File
	file_protoreflectextra_v1_types_proto_goTypes = nil
	file_protoreflectextra_v1_types_proto_depIdxs = nil
}
