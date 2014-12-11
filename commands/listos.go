package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"sort"
	"strconv"
	"text/tabwriter"
	"io"
)

type listOS struct{ BasicCommand }

func NewListOS() Command {
	return &listOS{
		BasicCommand{
			Desc:        "List all available operating systems.",
			NeedsKey:    false,
			ArgsDesc:    "",
			OptionsDesc: "",
		},
	}
}

func (_ *listOS) Fexec(w io.Writer, c Client, _ []string, _ string) (err error) {
	r, err := requests.GetOS(c)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintln(t, "ID\tNAME\tFAMILY\tARCH\tWINDOWS")

	for _, v := range a {
		fmt.Fprintf(t, "%d\t%s\t%s\t%s\t%s\n", v.Id, v.Name, v.Family,
			v.Arch, strconv.FormatBool(v.Windows))
	}

	t.Flush()

	return
}
