package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

type listOS struct{BasicCommand}

func NewListOS() *listOS {
	return &listOS{
		BasicCommand {
			Desc: "List all available operating systems.",
			NeedsKey: false,
			ArgsDesc: "",
			OptionsDesc: "",
		},
	}
}

func (_ *listOS) Exec(c Client, _ []string, _ string) (err error) {
	r, err := requests.GetOS(c)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "ID\tNAME\tFAMILY\tARCH\tWINDOWS")

	for _, v := range a {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", v.Id, v.Name, v.Family,
			v.Arch, strconv.FormatBool(v.Windows))
	}

	w.Flush()

	return
}
