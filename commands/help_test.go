package commands

import (
	. "github.com/stephan83/vultrapi/errors"
	"testing"
)

func ExampleHelpListRegions() {
	NewHelp("vultrapi", cmdDict).Exec(nil, []string{"help", "listregions"}, "")
	// Output:
	// List all available regions.
	//
	// Usage: vultrapi listregions  [options...]
	//
	// Options:
	// None.
}

func ExampleHelpListPlans() {
	NewHelp("vultrapi", cmdDict).Exec(nil, []string{"help", "listplans"}, "")
	// Output:
	// List available plans.
	//
	// Usage: vultrapi listplans  [options...]
	//
	// Options:
	//   -region=0: limit to region id
}

func ExampleHelpCreateServer() {
	NewHelp("vultrapi", cmdDict).Exec(nil, []string{"help", "createserver"}, "")
	// Output:
	// Creates a server.
	//
	// Usage: vultrapi createserver region_id plan_id os_id [options...]
	//
	// You must set env variable VULTR_API_KEY to your API key.
	//
	// Options:
	//   -enable_auto_backups=false: Enable auto auto backups
	//   -enable_ipv6=false: Enable IPV6
	//   -enable_private_network=false: Enable private network
	//   -ipxe_chain_url="": IPXE chain url
	//   -iso_id=0: ISO ID
	//   -label="": Label
	//   -script_id=0: Script ID
	//   -snapshot_id=0: Snapshot ID
	//   -ssh_key_id=0: SSH key ID
}

func TestHelpNotEnoughArgs(t *testing.T) {
	err := NewHelp("vultrapi", cmdDict).Exec(nil, []string{"help"}, "")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUsage); !ok {
		t.Error("Error is not ErrUsage.")
	}
}

func TestHelpUnknownCommand(t *testing.T) {
	err := NewHelp("vultrapi", cmdDict).Exec(nil, []string{"help", "listplasn"}, "")
	if err == nil {
		t.Error("No error returned.")
	}
	if _, ok := err.(ErrUnknownCommand); !ok {
		t.Error("Error is not ErrUnknownCommand.")
	}
}

var cmdDict = map[string]Command{
	"listregions":  NewListRegions(),
	"listplans":    NewListPlans(),
	"createserver": NewCreateServer(),
}
