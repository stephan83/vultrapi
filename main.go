package main

import (
	"fmt"
	"github.com/stephan83/vultrapi/commands"
	. "github.com/stephan83/vultrapi/errors"
	"os"
	"sort"
	"strings"
)

var cmdDict = map[string]commands.Command{
	"listregions":  commands.NewListRegions(),
	"listplans":    commands.NewListPlans(),
	"listos":       commands.NewListOS(),
	"account":      commands.NewAccount(),
	"createserver": commands.NewCreateServer(),
}

var cmds = cmdArray{}

func init() {
	cmdDict["help"] = commands.NewHelp(cmdDict)

	for name, cmd := range cmdDict {
		cmds = append(cmds, cmdWithName{cmd, name})
	}

	sort.Sort(cmds)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd, ok := cmdDict[os.Args[1]]
	if !ok {
		printUsage()
		os.Exit(1)
	}

	if err := cmd.Exec(); err != nil {
		if _, ok = err.(ErrUsage); ok {
			commands.PrintUsage(os.Args[1], cmd)
		} else {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("\033[1mUsage:\033[0m vultrapi command [options...]\n")
	fmt.Println("You must set env variable \033[1mVULTR_API_KEY\033[0m to your API key for commands in bold.\n")
	fmt.Println("\033[1mCommands:\033[0m\n")

	for _, cmd := range cmds {
		if cmd.NeedsKey() {
			fmt.Printf("  \033[1m%s\033[0m %s\n", cmd.name,
				cmd.Args())
		} else {
			fmt.Printf("  %s %s\n", cmd.name, cmd.Args())
		}
		desc := strings.Split(cmd.Desc(), "\n")
		for _, line := range desc {
			fmt.Printf("  %s\n\n", line)
		}
	}
}

type cmdWithName struct {
	commands.Command
	name string
}

type cmdArray []cmdWithName

func (c cmdArray) Len() int {
	return len(c)
}

func (c cmdArray) Less(i, j int) bool {
	if c[i].name == "help" {
		return true
	}

	return c[i].name < c[j].name
}

func (c cmdArray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
