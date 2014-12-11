package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"io"
)

type help struct {
	BasicCommand
	name string
	cd   CommandMap
}

func NewHelp(name string, cd CommandMap) Command {
	return &help{
		BasicCommand{
			Desc:        "Get help for a command.",
			NeedsKey:    false,
			ArgsDesc:    "command",
			OptionsDesc: "",
		},
		name,
		cd,
	}
}

func (o *help) Fexec(w io.Writer, _ Client, args []string, _ string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	cmd, ok := o.cd[args[0]]
	if !ok {
		err = ErrUnknownCommand{}
		return
	}

	fmt.Fprintf(w, "%s\n\n", cmd.GetDesc())
	o.cd.FprintCommandUsage(w, o.name, args[0])

	if cmd.GetNeedsKey() {
		fmt.Fprintln(w, "\nYou must set env variable VULTR_API_KEY to your API key.")
	}

	if opt := cmd.GetOptionsDesc(); opt != "" {
		fmt.Fprintln(w, "\nOptions:")
		fmt.Fprintln(w, cmd.GetOptionsDesc())
	}

	return
}
