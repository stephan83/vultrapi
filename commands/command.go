package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"sort"
	"strings"
)

type Command interface {
	Desc() string
	NeedsKey() bool
	Args() string
	PrintOptions()
	Exec(client Client, args []string, key string) error
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
	fmt.Printf("Usage: %s %s %s [options...]\n", name, cmd,
		o[cmd].Args())
}

func (o CommandMap) PrintUsage(name string) {
	var cmds = commandArray{}

	for name, cmd := range o {
		cmds = append(cmds, commandWithName{cmd, name})
	}

	sort.Sort(cmds)

	fmt.Printf("Usage: %s command [arguments...] [options...]\n\n", name)
	fmt.Println("You must set env variable VULTR_API_KEY to your API key for commands prefixed with *.\n")
	fmt.Println("Commands:\n")

	for i, c := range cmds {
		if c.NeedsKey() {
			fmt.Printf("* %s", c.name)
		} else {
			fmt.Printf("  %s", c.name)
		}
		if args := c.Args(); args != "" {
			fmt.Printf(" %s", args)
		}
		fmt.Println()
		desc := strings.Split(c.Desc(), "\n")
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
	if a[i].NeedsKey() && !a[j].NeedsKey() {
		return false
	}
	if !a[i].NeedsKey() && a[j].NeedsKey() {
		return true
	}

	return a[i].name < a[j].name
}

func (a commandArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
