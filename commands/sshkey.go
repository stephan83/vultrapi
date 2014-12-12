package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"text/tabwriter"
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

	o, ok := r[args[0]]
	if !ok {
		err = ErrNotFound{}
		return
	}

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintf(t, "ID\t%s\n", o.Id)
	fmt.Fprintf(t, "NAME\t%s\n", o.Name)
	fmt.Fprintf(t, "DATE CREATED\t%s\n", o.DateCreated)
	fmt.Fprintf(t, "KEY\t%s\n", o.Key)

	t.Flush()

	return
}
