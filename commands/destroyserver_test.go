package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleDestroyServer() {
	c := NewTestClient(200, []byte("ok"))
	NewDestroyServer().Exec(c, []string{"123"}, "")
	// Output:
	// OK
}
