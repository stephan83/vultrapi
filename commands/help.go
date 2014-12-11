package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
)

type help struct {
	Command
	name string
	cd   CommandMap
}

func NewHelp(name string, cd CommandMap) *help {
	return &help{
		Command {
			Desc: "Get help for a command.",
			NeedsKey: false,
			ArgsDesc: "command",
			OptionsDesc: "",
		},
		name,
		cd,
	}
}

func (o *help) Exec(_ Client, args []string, _ string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	cmd, ok := o.cd[args[0]]
	if !ok {
		err = ErrUnknownCommand{}
		return
	}

	fmt.Printf("%s\n\n", cmd.Desc)
	o.cd.PrintCommandUsage(o.name, args[0])

	if cmd.NeedsKey {
		fmt.Println("\nYou must set env variable VULTR_API_KEY to your API key.")
	}

	if opt := cmd.OptionsDesc; opt != "" {
		fmt.Println("\nOptions:")
		fmt.Println(cmd.OptionsDesc)
	}

	return
}
