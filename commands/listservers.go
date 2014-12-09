package commands

import (
	"fmt"
	"flag"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"os"
)

type listServers struct{
	flagSet  *flag.FlagSet
	regionId int
	planId int
}

func NewListServers() Command {
	ls := listServers{
		flagSet: flag.NewFlagSet("listservers", flag.ContinueOnError),
	}
	ls.flagSet.IntVar(&ls.regionId, "region", 0, "limit to region id")
	ls.flagSet.IntVar(&ls.planId, "plan", 0, "limit to plan id")
	return &ls
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

func (ls *listServers) PrintOptions() {
	ls.flagSet.SetOutput(os.Stdout)
	ls.flagSet.PrintDefaults()
	ls.flagSet.SetOutput(os.Stderr)
}

func (ls *listServers) Exec(c Client, args []string, key string) (err error) {
	err = ls.flagSet.Parse(args)
	if err != nil {
		return ErrUsage{}
	}

	s, err := requests.GetServers(c, key)
	if err != nil {
		return
	}

	if ls.regionId > 0 && ls.planId > 0 {
		for id, v := range s {
			if v.RegionId != ls.regionId || v.PlanId != ls.planId {
				delete(s, id)
			}
		}
	} else if ls.regionId > 0 {
		for id, v := range s {
			if v.RegionId != ls.regionId {
				delete(s, id)
			}
		}
	} else if ls.planId > 0 {
		for id, v := range s {
			if v.PlanId != ls.planId {
				delete(s, id)
			}
		}
	}

	fmt.Println(s)

	return
}
