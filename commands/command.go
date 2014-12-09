package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
)

type Command interface {
	Desc() string
	NeedsKey() bool
	Args() string
	PrintOptions()
	Exec(client Client, args []string, key string) error
}

func PrintUsage(name string, cmd Command) {
	fmt.Printf("Usage: vultrapi %s %s [options...]\n",
		name, cmd.Args())
}
