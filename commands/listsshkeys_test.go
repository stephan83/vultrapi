package commands

import(
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleListSSHKeys() {
	c := NewTestClient(200, keys)
	NewListSSHKeys().Fexec(os.Stdout, c, []string{}, "")
	// Output:
	// ID		NAME	DATE CREATED
	// 548785eca773d	test1	2014-12-09 23:29:48 +0000
	// 5487861ad6c8a	test2	2014-12-09 23:30:34 +0000
}

var keys = []byte(`{
	"548785eca773d": {
		"SSHKEYID": "548785eca773d",
		"date_created": "2014-12-09 18:29:48",
		"name": "test1",
		"ssh_key": "****"
	},
	"5487861ad6c8a": {
		"SSHKEYID": "5487861ad6c8a",
		"date_created": "2014-12-09 18:30:34",
		"name": "test2",
		"ssh_key": "****"
	}
}`)
