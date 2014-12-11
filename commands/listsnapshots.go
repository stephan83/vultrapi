package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"sort"
	"text/tabwriter"
)

type listSnapshots struct{}

func NewListSnapshots() Command {
	return listSnapshots{}
}

func (_ listSnapshots) Desc() string {
	return "List all snapshots."
}

func (_ listSnapshots) Args() string {
	return ""
}

func (_ listSnapshots) NeedsKey() bool {
	return true
}

func (_ listSnapshots) PrintOptions() {
	fmt.Println("None.")
}

func (_ listSnapshots) Exec(c Client, args []string, key string) (err error) {
	r, err := requests.GetSnapshots(c, key)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "ID\tDESCRIPTION\tDATE CREATED\tSIZE\tSTATUS")

	for _, v := range a {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\n", v.Id, v.Description,
			v.DateCreated, v.Size, v.Status)
	}

	w.Flush()

	return
}