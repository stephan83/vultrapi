package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleListRegions() {
	c := NewTestClient(200, regions)
	NewListRegions().Exec(c, []string{}, "")
	// Output:
	// ID	NAME		CONTINENT	COUNTRY	STATE
	// 39	Miami				US	FL
	// 25	Tokyo		Asia		JP	-
	// 19	Australia	Australia	AU	-
	// 9	Frankfurt	Europe		DE	-
	// 24	France		Europe		FR	-
	// 8	London		Europe		GB	-
	// 7	Amsterdam	Europe		NL	-
	// 5	Los Angeles	North America	US	CA
	// 12	Silicon Valley	North America	US	CA
	// 6	Atlanta		North America	US	GA
	// 2	Chicago		North America	US	IL
	// 1	New Jersey	North America	US	NJ
	// 3	Dallas		North America	US	TX
	// 4	Seattle		North America	US	WA
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
		"state": "-"
	},
	"25": {
		"DCID": "25",
		"name": "Tokyo",
		"country": "JP",
		"continent": "Asia",
		"state": "-"
	},
	"8": {
		"DCID": "8",
		"name": "London",
		"country": "GB",
		"continent": "Europe",
		"state": "-"
	},
	"24": {
		"DCID": "24",
		"name": "France",
		"country": "FR",
		"continent": "Europe",
		"state": "-"
	},
	"9": {
		"DCID": "9",
		"name": "Frankfurt",
		"country": "DE",
		"continent": "Europe",
		"state": "-"
	},
	"19": {
		"DCID": "19",
		"name": "Australia",
		"country": "AU",
		"continent": "Australia",
		"state": "-"
	}
}`)
