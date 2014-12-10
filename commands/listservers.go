package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"sort"
	"text/tabwriter"
)

type listServers struct {
	flagSet  *flag.FlagSet
	regionId int
	planId   int
}

func NewListServers() Command {
	o := listServers{
		flagSet: flag.NewFlagSet("listservers", flag.ContinueOnError),
	}
	o.flagSet.IntVar(&o.regionId, "region", 0, "limit to region id")
	o.flagSet.IntVar(&o.planId, "plan", 0, "limit to plan id")
	return &o
}

func (_ *listServers) NeedsKey() bool {
	return true
}

func (_ *listServers) Args() string {
	return ""
}

func (_ *listServers) Desc() string {
	return "List all servers."
}

func (o *listServers) PrintOptions() {
	o.flagSet.SetOutput(os.Stdout)
	o.flagSet.PrintDefaults()
	o.flagSet.SetOutput(os.Stderr)
}

func (o *listServers) Exec(c Client, args []string, key string) (err error) {
	err = o.flagSet.Parse(args)
	if err != nil {
		return ErrUsage{}
	}

	r, err := requests.GetServers(c, key)
	if err != nil {
		return
	}

	if o.regionId > 0 && o.planId > 0 {
		for id, v := range r {
			if v.RegionId != o.regionId || v.PlanId != o.planId {
				delete(r, id)
			}
		}
	} else if o.regionId > 0 {
		for id, v := range r {
			if v.RegionId != o.regionId {
				delete(r, id)
			}
		}
	} else if o.planId > 0 {
		for id, v := range r {
			if v.PlanId != o.planId {
				delete(r, id)
			}
		}
	}

	a := r.Array()
	sort.Sort(a)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "ID\tLOCATION\tOS\tIPV4\tSTATUS\tLABEL")

	for _, v := range a {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\n", v.Id, v.Location,
			v.OS, v.IPV4, v.Status, v.Label)
	}

	w.Flush()

	return
}
