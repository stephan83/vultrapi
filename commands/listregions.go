package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"sort"
	"text/tabwriter"
)

type listRegions struct{ BasicCommand }

func NewListRegions() Command {
	return &listRegions{
		BasicCommand{
			Desc:        "List all available regions.",
			NeedsKey:    false,
			ArgsDesc:    "",
			OptionsDesc: "",
		},
	}
}

func (_ *listRegions) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	r, err := requests.GetRegions(c)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintln(t, "ID\tNAME\tCONTINENT\tCOUNTRY\tSTATE")

	for _, v := range a {
		fmt.Fprintf(t, "%d\t%s\t%s\t%s\t%s\n", v.Id, v.Name, v.Continent, v.Country, v.State)
	}

	t.Flush()

	return
}
