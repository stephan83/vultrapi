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
	flagSet  *flag.FlagSet
	regionId int
}

func NewListPlans() Command {
	o := listPlans{
		flagSet: flag.NewFlagSet("listplans", flag.ContinueOnError),
	}
	o.flagSet.IntVar(&o.regionId, "region", 0, "limit to region id")
	return &o
}

func (_ *listPlans) Args() string {
	return ""
}

func (_ *listPlans) Desc() string {
	return "List all available plans."
}

func (_ *listPlans) NeedsKey() bool {
	return false
}

func (o *listPlans) PrintOptions() {
	o.flagSet.SetOutput(os.Stdout)
	o.flagSet.PrintDefaults()
	o.flagSet.SetOutput(os.Stderr)
}

func (o *listPlans) Exec(c Client, args []string, _ string) (err error) {
	err = o.flagSet.Parse(args)
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
