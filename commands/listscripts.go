package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"sort"
	"text/tabwriter"
)

type listScripts struct{ BasicCommand }

func NewListScripts() Command {
	return &listScripts{
		BasicCommand{
			Desc:        "List all scripts.",
			NeedsKey:    true,
			ArgsDesc:    "",
			OptionsDesc: "",
		},
	}
}

func (_ *listScripts) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	r, err := requests.GetScripts(c, key)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintln(t, "ID\tNAME\tDATE CREATED")

	for _, v := range a {
		fmt.Fprintf(t, "%d\t%s\t%s\n", v.Id, v.Name, v.DateCreated)
	}

	t.Flush()

	return
}
