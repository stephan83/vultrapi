package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleDestroySnapshot() {
	c := NewTestClient(200, []byte("ok"))
	NewDestroySnapshot().Exec(c, []string{"abc"}, "")
	// Output:
	// OK
}
