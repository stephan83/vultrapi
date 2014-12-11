package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"sort"
	"strings"
	"flag"
	"os"
	"bytes"
)

type Command interface {
	GetDesc() string
	GetNeedsKey() bool
	GetArgsDesc() string
	GetOptionsDesc() string
	Exec(c Client, args []string, key string) error
}

type BasicCommand struct {
	Desc string
	NeedsKey bool
	ArgsDesc string
	OptionsDesc string
}

func (o *BasicCommand) GetDesc() string {
	return o.Desc
}

func (o *BasicCommand) GetNeedsKey() bool {
	return o.NeedsKey
}

func (o *BasicCommand) GetArgsDesc() string {
	return o.ArgsDesc
}

func (o *BasicCommand) GetOptionsDesc() string {
	return o.OptionsDesc
}

func (_ *BasicCommand) Exec(c Client, args []string, key string) error {
	fmt.Println("Not implemented.")
	return nil
}

type BasicCommandWithOptions struct {
	BasicCommand
	FlagSet *flag.FlagSet
}

func (o *BasicCommandWithOptions) Initialize() {
	var buffer bytes.Buffer
	o.FlagSet.SetOutput(&buffer)
	o.FlagSet.PrintDefaults()
	o.FlagSet.SetOutput(os.Stderr)
	o.OptionsDesc = buffer.String()
}

type CommandMap map[string]Command

func (o CommandMap) Exec(args []string, c Client, key string) error {
	if len(args) < 1 {
		return ErrUsage{}
	}

	cmd, ok := o[args[0]]
	if !ok {
		return ErrUnknownCommand{}
	}

	return cmd.Exec(c, args[1:], key)
}

func (o CommandMap) PrintCommandUsage(name string, cmd string) {
	fmt.Printf("Usage: %s %s %s [options...]\n", name, cmd, o[cmd].GetArgsDesc())
}

func (o CommandMap) PrintUsage(name string) {
	var cmds = commandArray{}

	for name, cmd := range o {
		cmds = append(cmds, commandWithName{cmd, name})
	}

	sort.Sort(cmds)

	fmt.Printf("Usage: %s command [arguments...] [options...]\n\n", name)
	fmt.Print("You must set env variable VULTR_API_KEY to your API key for commands prefixed with *.\n\n")
	fmt.Print("Commands:\n\n")

	for i, c := range cmds {
		if c.GetNeedsKey() {
			fmt.Printf("* %s", c.name)
		} else {
			fmt.Printf("  %s", c.name)
		}
		if args := c.GetArgsDesc(); args != "" {
			fmt.Printf(" %s", args)
		}
		fmt.Println()
		desc := strings.Split(c.GetDesc(), "\n")
		for _, line := range desc {
			fmt.Printf("  %s\n", line)
		}
		if i+1 < len(cmds) {
			fmt.Println("")
		}
	}
}

type commandWithName struct {
	Command
	name string
}

type commandArray []commandWithName

func (a commandArray) Len() int {
	return len(a)
}

func (a commandArray) Less(i, j int) bool {
	if a[i].name == "help" {
		return true
	}
	if a[j].name == "help" {
		return false
	}
	if a[i].GetNeedsKey() && !a[j].GetNeedsKey() {
		return false
	}
	if !a[i].GetNeedsKey() && a[j].GetNeedsKey() {
		return true
	}

	return a[i].name < a[j].name
}

func (a commandArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
