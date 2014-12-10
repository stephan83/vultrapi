package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleSSHKey() {
	c := NewTestClient(200, keys)
	NewSSHKey().Exec(c, []string{"5487861ad6c8a"}, "API_KEY")
	// ID		5487861ad6c8a
	// NAME		test2
	// DATE CREATED	2014-12-09 23:30:34 +0000
	// KEY		****
}
