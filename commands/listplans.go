package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"github.com/stephan83/vultrapi/types"
	"os"
	"sort"
	"text/tabwriter"
)

type listPlans struct {
	BasicCommandWithOptions
	regionId int
}

func NewListPlans() Command {
	f := flag.NewFlagSet("listplans", flag.ContinueOnError)

	o := listPlans{
		BasicCommandWithOptions{
			BasicCommand{
				Desc:     "List all available plans.",
				NeedsKey: false,
				ArgsDesc: "",
			},
			f,
		},
		0,
	}

	f.IntVar(&o.regionId, "region", 0, "limit to region id")

	o.Initialize()

	return &o
}

func (o *listPlans) Exec(c Client, args []string, _ string) (err error) {
	err = o.FlagSet.Parse(args)
	if err != nil {
		return ErrUsage{}
	}

	r, err := requests.GetPlans(c)
	if err != nil {
		return
	}

	if o.regionId > 0 {
		a, err := requests.GetRegionAvailability(c, o.regionId)
		if err != nil {
			return err
		}

		filtered := types.PlanMap{}

		for _, plan := range a {
			key := plan
			filtered[key] = r[key]
		}

		r = filtered
	}

	a := r.Array()
	sort.Sort(a)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "ID\tNAME\tCPUS\tPRICE/MONTH")

	for _, v := range a {
		fmt.Fprintf(w, "%d\t%s\t%d\t%.2f\n", v.Id, v.Name, v.CPUs,
			v.PricePerMonth)
	}

	w.Flush()

	return
}
