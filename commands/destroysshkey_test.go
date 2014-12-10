package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleDestroySSHKey() {
	c := NewTestClient(200, []byte("ok"))
	NewDestroySSHKey().Exec(c, []string{"123"}, "")
	// Output:
	// OK
}
