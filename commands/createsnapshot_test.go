package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"testing"
)

func ExampleCreateSnapshot() {
	c := NewTestClient(200, []byte(`{"SNAPSHOTID": "123456"}`))
	NewCreateSnapshot().Exec(c, []string{"1234"}, "SECRET_KEY")
	// Output: SNAPSHOT ID:	123456
}

func TestCreateSnapshotNotEnoughArgs(t *testing.T) {
	err := NewCreateSnapshot().Exec(nil, []string{}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}
