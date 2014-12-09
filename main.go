package main

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/commands"
	. "github.com/stephan83/vultrapi/errors"
	"os"
)

const name = "vultrapi"

var cmdDict = CommandDict{
	"listregions":   NewListRegions(),
	"listplans":     NewListPlans(),
	"listos":        NewListOS(),
	"account":       NewAccount(),
	"createserver":  NewCreateServer(),
	"listservers":   NewListServers(),
	"server":        NewServer(),
	"destroyserver": NewDestroyServer(),
	"createsshkey":  NewCreateSSHKey(),
}

func init() {
	cmdDict["help"] = NewHelp("vultrapi", cmdDict)
}

func main() {
	c := NewVultrClient("https://api.vultr.com/v1")
	os.Exit(run(cmdDict, c, os.Args, os.Getenv("VULTR_API_KEY")))
}

func run(cd CommandDict, c Client, args []string, key string) int {
	if len(args) < 2 {
		cd.PrintUsage(name)
		return 2
	}

	switch err := cd.Exec(args[1:], c, key); err.(type) {
	case ErrUsage:
		cd.PrintCommandUsage(name, args[1])
		return 2
	case ErrUnknownCommand:
		fmt.Println(err.Error())
		return 2
	case error:
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	return 0
}
