package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"sort"
	"text/tabwriter"
	"io"
)

type listServers struct {
	BasicCommandWithOptions
	regionId int
	planId   int
}

func NewListServers() Command {
	f := flag.NewFlagSet("listservers", flag.ContinueOnError)

	o := listServers{
		BasicCommandWithOptions{
			BasicCommand{
				Desc:     "List all servers.",
				NeedsKey: true,
				ArgsDesc: "",
			},
			f,
		},
		0,
		0,
	}

	f.IntVar(&o.regionId, "region", 0, "limit to region id")
	f.IntVar(&o.planId, "plan", 0, "limit to plan id")

	o.Initialize()

	return &o
}

func (o *listServers) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	o.FlagSet.SetOutput(w)

	err = o.FlagSet.Parse(args)
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

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintln(t, "ID\tLOCATION\tOS\tIPV4\tSTATUS\tLABEL")

	for _, v := range a {
		fmt.Fprintf(t, "%d\t%s\t%s\t%s\t%s\t%s\n", v.Id, v.Location, v.OS, v.IPV4, v.Status, v.Label)
	}

	t.Flush()

	return
}
