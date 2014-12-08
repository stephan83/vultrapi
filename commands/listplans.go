package commands

import (
	"flag"
	"fmt"
	"github.com/stephan83/vultrapi/requests"
	"github.com/stephan83/vultrapi/types"
	"os"
	"strconv"
)

type listPlans struct {
	flagSet  *flag.FlagSet
	regionId int
}

func NewListPlans() Command {
	lp := listPlans{
		flagSet: flag.NewFlagSet("listplans", flag.ExitOnError),
	}
	lp.flagSet.IntVar(&lp.regionId, "region", 0, "limit to region id")
	return &lp
}

func (_ *listPlans) Args() string {
	return ""
}

func (_ *listPlans) Desc() string {
	return "List available plans."
}

func (_ *listPlans) NeedsKey() bool {
	return false
}

func (lp *listPlans) PrintOptions() {
	lp.flagSet.PrintDefaults()
}

func (lp *listPlans) Exec() (err error) {
	lp.flagSet.Parse(os.Args[2:])

	plans, err := requests.GetPlans()
	if err != nil {
		return
	}

	if lp.regionId > 0 {
		avaibility, err := requests.GetRegionAvailability(lp.regionId)
		if err != nil {
			return err
		}

		regionPlans := types.PlanDict{}

		for _, plan := range avaibility {
			key := strconv.Itoa(plan)
			regionPlans[key] = plans[key]
		}

		plans = regionPlans
	}

	fmt.Println(plans)

	return
}
