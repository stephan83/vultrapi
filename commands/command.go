package commands

import "fmt"

type Command interface {
	Desc() string
	NeedsKey() bool
	Args() string
	PrintOptions()
	Exec() error
}

func PrintUsage(name string, cmd Command) {
	fmt.Printf("\033[1mUsage:\033[0m vultrapi %s %s [options...]\n",
		name, cmd.Args())
}
