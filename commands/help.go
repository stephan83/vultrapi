package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
)

type help struct {
	name string
	cd CommandDict
}

func NewHelp(name string, cd CommandDict) Command {
	return help{name, cd}
}

func (_ help) Desc() string {
	return "Get help for a command."
}

func (_ help) NeedsKey() bool {
	return false
}

func (_ help) Args() string {
	return "command"
}

func (_ help) PrintOptions() {
	fmt.Println("None.")
}

func (h help) Exec(_ Client, args []string, _ string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	cmd, ok := h.cd[args[0]]
	if !ok {
		err = ErrUnknownCommand{}
		return
	}

	fmt.Printf("%s\n\n", cmd.Desc())
	h.cd.PrintCommandUsage(h.name, args[0])

	if cmd.NeedsKey() {
		fmt.Println("\nYou must set env variable VULTR_API_KEY to your API key.")
	}

	fmt.Println("\nOptions:")
	cmd.PrintOptions()

	return
}
