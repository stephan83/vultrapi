package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleDestroyServer() {
	c := NewTestClient(200, []byte("ok"))
	NewDestroyServer().Fexec(os.Stdout, c, []string{"123"}, "")
	// Output:
	// OK
}
