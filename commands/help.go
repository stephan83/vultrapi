package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/errors"
	"os"
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

func (h help) Exec() (err error) {
	if len(os.Args) < 3 {
		err = ErrUsage{}
		return
	}

	cmd, ok := h.commands[os.Args[2]]
	if !ok {
		err = ErrUnknownCommand{}
		return
	}

	fmt.Printf("\033[1m%s\033[0m\n\n", cmd.Desc())
	PrintUsage(os.Args[2], cmd)

	if cmd.NeedsKey() {
		fmt.Println("\nYou must set env variable \033[1mVULTR_API_KEY\033[0m to your API key.")
	}

	fmt.Println("\n\033[1mOptions:\033[0m")
	cmd.PrintOptions()

	return
}
