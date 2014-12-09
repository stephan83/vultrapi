package main

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
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
	c := NewVultrClient("https://api.vultr.com/v1")
	run(c, os.Args, os.Getenv("VULTR_API_KEY"))
}

func run(c Client, args []string, key string) {
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd, ok := cmdDict[args[1]]
	if !ok {
		printUsage()
		os.Exit(1)
	}

	if err := cmd.Exec(c, args[1:], key); err != nil {
		if _, ok = err.(ErrUsage); ok {
			commands.PrintUsage(args[1], cmd)
		} else {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: vultrapi command [options...]\n")
	fmt.Println("You must set env variable VULTR_API_KEY to your API key for underlined commands.\n")
	fmt.Println("Commands:\n")

	for _, cmd := range cmds {
		fmt.Printf("  %s %s\n", cmd.name, cmd.Args())
		if cmd.NeedsKey() {
			fmt.Printf("  %s\n", strings.Repeat("*", len(cmd.name)))
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
