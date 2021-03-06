package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"sort"
	"text/tabwriter"
)

type listSSHKeys struct{ BasicCommand }

func NewListSSHKeys() Command {
	return &listSSHKeys{
		BasicCommand{
			Desc:        "List all SSH keys.",
			NeedsKey:    true,
			ArgsDesc:    "",
			OptionsDesc: "",
		},
	}
}

func (_ *listSSHKeys) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	r, err := requests.GetSSHKeys(c, key)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintln(t, "ID\tNAME\tDATE CREATED")

	for _, v := range a {
		fmt.Fprintf(t, "%s\t%s\t%s\n", v.Id, v.Name, v.DateCreated)
	}

	t.Flush()

	return
}
