package commands

func ExampleCommandPrintUsage() {
	cd := CommandMap{
		"listregions": NewListRegions(),
		"account":     NewAccount(),
	}
	cd["help"] = NewHelp("vultrapi", cmdMap)

	cd.PrintUsage("vultrapi")
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
