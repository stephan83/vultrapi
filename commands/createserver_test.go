package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"testing"
)

func ExampleCreateServer() {
	c := NewTestClient(200, []byte(`{"SUBID": "123456"}`))
	NewCreateServer().Exec(c, []string{"1", "2", "3"}, "SECRET_KEY")
	// Output: SERVER ID: 123456
}

func TestCreateServerNotEnoughArgs(t *testing.T) {
	err := NewCreateServer().Exec(nil, []string{"1", "2"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}

func TestCreateServerInvalidFirstArg(t *testing.T) {
	err := NewCreateServer().Exec(nil, []string{"a", "2", "3"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}

func TestCreateServerInvalidSecondArg(t *testing.T) {
	err := NewCreateServer().Exec(nil, []string{"1", "b", "3"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}

func TestCreateServerInvalidThirdArg(t *testing.T) {
	err := NewCreateServer().Exec(nil, []string{"1", "2", "false"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}

func TestCreateServerInvalidFlag(t *testing.T) {
	err := NewCreateServer().Exec(nil, []string{"1", "2", "3", "-enabel_private_network"}, "SECRET_KEY")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}
