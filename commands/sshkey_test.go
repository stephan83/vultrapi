package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleSSHKey() {
	c := NewTestClient(200, keys)
	NewSSHKey().Fexec(os.Stdout, c, []string{"5487861ad6c8a"}, "API_KEY")
	// Output:
	// ID		5487861ad6c8a
	// NAME		test2
	// DATE CREATED	2014-12-09 23:30:34 +0000
	// KEY		****
}

func ExampleSSHKeySpecificField() {
	c := NewTestClient(200, keys)
	NewSSHKey().Fexec(os.Stdout, c, []string{"5487861ad6c8a", "-field", "name", "-field", "key"}, "API_KEY")
	// Output:
	// NAME	test2
	// KEY	****
}
