package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"sort"
	"text/tabwriter"
)

type listSSHKeys struct{}

func NewListSSHKeys() Command {
	return listSSHKeys{}
}

func (_ listSSHKeys) NeedsKey() bool {
	return true
}

func (_ listSSHKeys) Args() string {
	return ""
}

func (_ listSSHKeys) Desc() string {
	return "List all SSH keys."
}

func (_ listSSHKeys) PrintOptions() {
	fmt.Println("None.")
}

func (o listSSHKeys) Exec(c Client, args []string, key string) (err error) {
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
