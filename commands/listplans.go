package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"github.com/stephan83/vultrapi/types"
	"os"
)

type listPlans struct {
	flagSet  *flag.FlagSet
	regionId int
}

func NewListPlans() Command {
	lp := listPlans{
		flagSet: flag.NewFlagSet("listplans", flag.ContinueOnError),
	}
	lp.flagSet.IntVar(&lp.regionId, "region", 0, "limit to region id")
	return &lp
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

func (lp *listPlans) PrintOptions() {
	lp.flagSet.SetOutput(os.Stdout)
	lp.flagSet.PrintDefaults()
	lp.flagSet.SetOutput(os.Stderr)
}

func (lp *listPlans) Exec(c Client, args []string, _ string) (err error) {
	err = lp.flagSet.Parse(args)
	if err != nil {
		return ErrUsage{}
	}

	plans, err := requests.GetPlans(c)
	if err != nil {
		return
	}

	if lp.regionId > 0 {
		a, err := requests.GetRegionAvailability(c, lp.regionId)
		if err != nil {
			return err
		}

		regionPlans := types.PlanMap{}

		for _, plan := range a {
			key := plan
			regionPlans[key] = plans[key]
		}

		plans = regionPlans
	}

	fmt.Println(plans)

	return
}
