package commands

import(
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleDestroySnapshot() {
	c := NewTestClient(200, []byte("ok"))
	NewDestroySnapshot().Fexec(os.Stdout, c, []string{"abc"}, "")
	// Output:
	// OK
}
