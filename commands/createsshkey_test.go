package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"path"
	"runtime"
	"testing"
)

func getCurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func ExampleCreateSSHKey() {
	keyPath := path.Join(getCurrentDir(), "..", "test_rsa.pub")
	c := NewTestClient(200, []byte(`{"SSHKEYID": "123456"}`))
	NewCreateSSHKey().Exec(c, []string{"test", keyPath}, "SECRET_KEY")
	// Output: SSH KEY ID:	123456
}

func TestCreateSSHKeyNotEnoughArgs(t *testing.T) {
	err := NewCreateSSHKey().Exec(nil, []string{"test"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}
