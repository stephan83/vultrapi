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

func (cd CommandMap) Exec(args []string, c Client, key string) error {
	if len(args) < 1 {
		return ErrUsage{}
	}

	cmd, ok := cd[args[0]]
	if !ok {
		return ErrUnknownCommand{}
	}

	return cmd.Exec(c, args[1:], key)
}

func (cd CommandMap) PrintCommandUsage(name string, cmd string) {
	fmt.Printf("Usage: %s %s %s [options...]\n", name, cmd,
		cd[cmd].Args())
}

func (cd CommandMap) PrintUsage(name string) {
	var cmds = commandArray{}

	for name, cmd := range cd {
		cmds = append(cmds, commandWithName{cmd, name})
	}

	sort.Sort(cmds)

	fmt.Printf("Usage: %s command [options...]\n\n", name)
	fmt.Println("You must set env variable VULTR_API_KEY to your API key for underlined commands.\n")
	fmt.Println("Commands:\n")

	for i, cmd := range cmds {
		fmt.Printf("  %s", cmd.name)
		if args := cmd.Args(); args != "" {
			fmt.Printf(" %s", args)
		}
		fmt.Println()
		if cmd.NeedsKey() {
			fmt.Printf("  %s\n", strings.Repeat("*", len(cmd.name)))
		}
		desc := strings.Split(cmd.Desc(), "\n")
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

func (c commandArray) Len() int {
	return len(c)
}

func (c commandArray) Less(i, j int) bool {
	if c[i].name == "help" {
		return true
	}
	if c[j].name == "help" {
		return false
	}
	if c[i].NeedsKey() && !c[j].NeedsKey() {
		return false
	}
	if !c[i].NeedsKey() && c[j].NeedsKey() {
		return true
	}

	return c[i].name < c[j].name
}

func (c commandArray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
