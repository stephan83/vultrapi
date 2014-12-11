package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"text/tabwriter"
	"io"
)

type sshKey struct{ BasicCommand }

func NewSSHKey() Command {
	return &sshKey{
		BasicCommand{
			Desc:        "Get server information.",
			NeedsKey:    true,
			ArgsDesc:    "ssh_key_id",
			OptionsDesc: "Get SSH key information.",
		},
	}
}

func (_ *sshKey) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
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

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintf(t, "ID\t%s\n", s.Id)
	fmt.Fprintf(t, "NAME\t%s\n", s.Name)
	fmt.Fprintf(t, "DATE CREATED\t%s\n", s.DateCreated)
	fmt.Fprintf(t, "KEY\t%s\n", s.Key)

	t.Flush()

	return
}
