package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"text/tabwriter"
)

type sshKey struct{}

func NewSSHKey() Command {
	return sshKey{}
}

func (_ sshKey) NeedsKey() bool {
	return true
}

func (_ sshKey) Args() string {
	return "ssh_key_id"
}

func (_ sshKey) Desc() string {
	return "Get SSH key information."
}

func (_ sshKey) PrintOptions() {
	fmt.Println("None.")
}

func (_ sshKey) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}
	r, err := requests.GetSSHKeys(c, key)
	if err != nil {
		return
	}

	s, ok := r[args[0]]
	if !ok {
		err = ErrNotFound{}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintf(w, "ID\t%s\n", s.Id)
	fmt.Fprintf(w, "NAME\t%s\n", s.Name)
	fmt.Fprintf(w, "DATE CREATED\t%s\n", s.DateCreated)
	fmt.Fprintf(w, "KEY\t%s\n", s.Key)

	w.Flush()

	return
}
