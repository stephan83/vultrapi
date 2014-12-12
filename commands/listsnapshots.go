package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"sort"
	"text/tabwriter"
)

type listSnapshots struct{ BasicCommand }

func NewListSnapshots() Command {
	return &listSnapshots{
		BasicCommand{
			Desc:        "List all snapshots.",
			NeedsKey:    true,
			ArgsDesc:    "",
			OptionsDesc: "",
		},
	}
}

func (_ *listSnapshots) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	r, err := requests.GetSnapshots(c, key)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintln(t, "ID\tDESCRIPTION\tDATE CREATED\tSIZE\tSTATUS")

	for _, v := range a {
		fmt.Fprintf(t, "%s\t%s\t%s\t%d\t%s\n", v.Id, v.Description, v.DateCreated, v.Size, v.Status)
	}

	t.Flush()

	return
}
