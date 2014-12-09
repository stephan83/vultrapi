package commands

func ExampleCommandPrintUsage() {
	cd := CommandDict{
		"listregions":  NewListRegions(),
		"account":      NewAccount(),
	}
	cd["help"] = NewHelp("vultrapi", cmdDict)

	cd.PrintUsage("vultrapi")
	// Output:
	// Usage: vultrapi command [options...]
	// 
	// You must set env variable VULTR_API_KEY to your API key for underlined commands.
	// 
	// Commands:
	// 
	//   help command
	//   Get help for a command.
	// 
	//   listregions
	//   List all available regions.
	// 
	//   account
	//   *******
	//   Get account information.
}
