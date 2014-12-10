package main

import "github.com/stephan83/vultrapi/commands"

func ExampleRunNoCommand() {
	cd := commands.CommandMap{
		"listregions": commands.NewListRegions(),
		"account":     commands.NewAccount(),
	}
	cd["help"] = commands.NewHelp("vultrapi", cmdMap)

	run(cd, nil, []string{"vultrapi"}, "")
	// Output:
	// Usage: vultrapi command [options...]
	//
	// You must set env variable VULTR_API_KEY to your API key for commands prefixed with *.
	//
	// Commands:
	//
	//   help command
	//   Get help for a command.
	//
	//   listregions
	//   List all available regions.
	//
	// * account
	//   Get account information.
}

func ExampleRunUnknownCommand() {
	cd := commands.CommandMap{
		"listregions": commands.NewListRegions(),
		"account":     commands.NewAccount(),
	}
	cd["help"] = commands.NewHelp("vultrapi", cmdMap)

	run(cd, nil, []string{"vultrapi", "accountn"}, "")
	// Output:
	// Unknown command.
}

func ExampleRunWrongUsage() {
	cd := commands.CommandMap{
		"listregions": commands.NewListRegions(),
		"account":     commands.NewAccount(),
	}
	cd["help"] = commands.NewHelp("vultrapi", cmdMap)

	run(cd, nil, []string{"vultrapi", "help"}, "")
	// Output:
	// Usage: vultrapi help command [options...]
}
