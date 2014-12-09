package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
)

type help struct {
	commands map[string]Command
}

func NewHelp(commands map[string]Command) Command {
	return help{commands}
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
	if len(args) < 2 {
		err = ErrUsage{}
		return
	}

	cmd, ok := h.commands[args[1]]
	if !ok {
		err = ErrUnknownCommand{}
		return
	}

	fmt.Printf("%s\n\n", cmd.Desc())
	PrintUsage(args[1], cmd)

	if cmd.NeedsKey() {
		fmt.Println("\nYou must set env variable VULTR_API_KEY to your API key.")
	}

	fmt.Println("\nOptions:")
	cmd.PrintOptions()

	return
}
