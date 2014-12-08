package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Plan struct {
	Id            int    `json:"VPSPLANID,string"`
	Name          string `json:"name"`
	CPUs          int    `json:"vcpu_count,string"`
	RAM           int    `json:"ram,string"`
	Disk          int    `json:"disk,string"`
	Bandwidth     string `json:"bandwidth"`
	PricePerMonth string `json:"price_per_month"`
	Windows       bool   `json:"windows"`
}

const (
	planFormat = "%-50s | %-4s | %-11s | %-3s"
)

func (r Plan) String() string {
	return fmt.Sprintf(planFormat, r.Name, strconv.Itoa(r.CPUs),
		r.PricePerMonth, strconv.Itoa(r.Id))
}

type PlanDict map[string]Plan

func (pd PlanDict) Slice() PlanSlice {
	plans := []Plan{}

	for _, p := range pd {
		plans = append(plans, p)
	}

	return plans
}

func (pd PlanDict) String() string {
	lines := []string{}

	lines = append(lines, fmt.Sprintf(planFormat, "NAME",
		"CPUS", "PRICE/MONTH", "ID"))

	lines = append(lines, strings.Repeat("-", 78))

	ps := pd.Slice()
	sort.Sort(ps)

	for _, r := range ps {
		lines = append(lines, r.String())
	}

	return strings.Join(lines, "\n")
}

type PlanSlice []Plan

func (ps PlanSlice) Len() int {
	return len(ps)
}

func (ps PlanSlice) Less(i, j int) bool {
	switch {
	case ps[i].CPUs < ps[j].CPUs:
		return true
	case ps[i].CPUs > ps[j].CPUs:
		return false
	default:
		switch {
		case ps[i].RAM < ps[j].RAM:
			return true
		case ps[i].RAM > ps[j].RAM:
			return false
		default:
			switch {
			case ps[i].Disk < ps[j].Disk:
				return true
			case ps[i].Disk > ps[j].Disk:
				return false
			default:
				return ps[i].Name < ps[j].Name
			}
		}
	}
}

func (ps PlanSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
