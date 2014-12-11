package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"os"
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

func (_ *listSSHKeys) Exec(c Client, args []string, key string) (err error) {
	r, err := requests.GetSSHKeys(c, key)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "ID\tNAME\tDATE CREATED")

	for _, v := range a {
		fmt.Fprintf(w, "%s\t%s\t%s\n", v.Id, v.Name, v.DateCreated)
	}

	w.Flush()

	return
}
