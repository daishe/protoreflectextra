package protoreflectextra_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func DiffMessages(want, got proto.Message) string {
	return cmp.Diff(want, got, protocmp.Transform())
}

func AssertMessagesEqual(t *testing.T, want, got proto.Message, msgAndArgs ...any) bool {
	t.Helper()
	if diff := DiffMessages(want, got); diff != "" {
		msg := formatMsgAndArgs(msgAndArgs...)
		if msg == "" {
			msg = "protobuf messages mismatch"
		}
		t.Errorf("%s (-want +got):\n%s", msg, diff)
		return false
	}
	return true
}

func RequireMessageEqual(t *testing.T, want, got proto.Message, msgAndArgs ...any) {
	t.Helper()
	if !AssertMessagesEqual(t, want, got, msgAndArgs...) {
		t.FailNow()
	}
}

func formatMsgAndArgs(msgAndArgs ...any) string {
	switch len(msgAndArgs) {
	case 0:
		return ""
	case 1:
		return msgAndArgs[0].(string)
	}
	return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
}
