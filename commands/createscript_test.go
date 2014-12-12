package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"os"
	"path"
	"testing"
)

func ExampleCreateScriptPxe() {
	keyPath := path.Join(getCurrentDir(), "..", "test_script.pxe")
	c := NewTestClient(200, []byte(`{"SCRIPTID": "123456"}`))
	NewCreateScript().Fexec(os.Stdout, c, []string{"pxe", "test", keyPath}, "SECRET_KEY")
	// Output: SCRIPT ID:	123456
}

func ExampleCreateScriptBoot() {
	keyPath := path.Join(getCurrentDir(), "..", "test_script.sh")
	c := NewTestClient(200, []byte(`{"SCRIPTID": "123456"}`))
	NewCreateScript().Fexec(os.Stdout, c, []string{"boot", "test", keyPath}, "SECRET_KEY")
	// Output: SCRIPT ID:	123456
}

func TestCreateScriptInvalidType(t *testing.T) {
	err := NewCreateScript().Fexec(os.Stdout, nil, []string{"go", "test", "path"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}

func TestCreateScriptNotEnoughArgs(t *testing.T) {
	err := NewCreateScript().Fexec(os.Stdout, nil, []string{"pxe", "test"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}
