package commands

import(
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleDestroySSHKey() {
	c := NewTestClient(200, []byte("ok"))
	NewDestroySSHKey().Fexec(os.Stdout, c, []string{"123"}, "")
	// Output:
	// OK
}
