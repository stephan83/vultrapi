package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleDestroyScript() {
	c := NewTestClient(200, []byte("ok"))
	NewDestroyScript().Fexec(os.Stdout, c, []string{"123"}, "")
	// Output:
	// OK
}
