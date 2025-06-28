package protoreflectextra_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/daishe/protoreflectextra"
	protoreflectextrav1 "github.com/daishe/protoreflectextra/internal/testtypes/protoreflectextra/v1"
)

func TestEnum(t *testing.T) {
	t.Parallel()
	base := protoreflectextrav1.Enum_ENUM_VALUE_UNSPECIFIED
	n := base.Number()
	ed := base.Descriptor()
	et := protoreflectextra.NewEnumType(ed)
	e := protoreflectextra.NewEnum(ed, n)

	require.Equal(t, ed, et.Descriptor())
	require.Equal(t, ed, et.New(n).Descriptor())
	require.Equal(t, n, et.New(n).Number())
	require.Equal(t, ed, e.Descriptor())
	require.Equal(t, n, e.Number())
	require.Equal(t, ed, e.Type().Descriptor())
}
