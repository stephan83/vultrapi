package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleListRegions() {
	c := NewTestClient(200, regions)
	NewListRegions().Exec(c, []string{"listregions"}, "")
	// Output:
	// CONTINENT            | COUNTRY | STATE | NAME                           | ID
	// ------------------------------------------------------------------------------
	//                      | US      | FL    | Miami                          | 39
	// Asia                 | JP      |       | Tokyo                          | 25
	// Australia            | AU      |       | Australia                      | 19
	// Europe               | DE      |       | Frankfurt                      | 9
	// Europe               | FR      |       | France                         | 24
	// Europe               | GB      |       | London                         | 8
	// Europe               | NL      |       | Amsterdam                      | 7
	// North America        | US      | CA    | Los Angeles                    | 5
	// North America        | US      | CA    | Silicon Valley                 | 12
	// North America        | US      | GA    | Atlanta                        | 6
	// North America        | US      | IL    | Chicago                        | 2
	// North America        | US      | NJ    | New Jersey                     | 1
	// North America        | US      | TX    | Dallas                         | 3
	// North America        | US      | WA    | Seattle                        | 4
}

var regions = []byte(`{
	"6": {
		"DCID": "6",
		"name": "Atlanta",
		"country": "US",
		"continent": "North America",
		"state": "GA"
	},
	"2": {
		"DCID": "2",
		"name": "Chicago",
		"country": "US",
		"continent": "North America",
		"state": "IL"
	},
	"3": {
		"DCID": "3",
		"name": "Dallas",
		"country": "US",
		"continent": "North America",
		"state": "TX"
	},
	"5": {
		"DCID": "5",
		"name": "Los Angeles",
		"country": "US",
		"continent": "North America",
		"state": "CA"
	},
	"39": {
		"DCID": "39",
		"name": "Miami",
		"country": "US",
		"continent": "",
		"state": "FL"
	},
	"1": {
		"DCID": "1",
		"name": "New Jersey",
		"country": "US",
		"continent": "North America",
		"state": "NJ"
	},
	"4": {
		"DCID": "4",
		"name": "Seattle",
		"country": "US",
		"continent": "North America",
		"state": "WA"
	},
	"12": {
		"DCID": "12",
		"name": "Silicon Valley",
		"country": "US",
		"continent": "North America",
		"state": "CA"
	},
	"7": {
		"DCID": "7",
		"name": "Amsterdam",
		"country": "NL",
		"continent": "Europe",
		"state": ""
	},
	"25": {
		"DCID": "25",
		"name": "Tokyo",
		"country": "JP",
		"continent": "Asia",
		"state": ""
	},
	"8": {
		"DCID": "8",
		"name": "London",
		"country": "GB",
		"continent": "Europe",
		"state": ""
	},
	"24": {
		"DCID": "24",
		"name": "France",
		"country": "FR",
		"continent": "Europe",
		"state": ""
	},
	"9": {
		"DCID": "9",
		"name": "Frankfurt",
		"country": "DE",
		"continent": "Europe",
		"state": ""
	},
	"19": {
		"DCID": "19",
		"name": "Australia",
		"country": "AU",
		"continent": "Australia",
		"state": ""
	}
}`)
