package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"encoding/json"
)

const planFormat = "%-50s | %-4s | %-11s | %s"

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

func (o Plan) String() string {
	return fmt.Sprintf(planFormat, o.Name, strconv.Itoa(o.CPUs),
		o.PricePerMonth, strconv.Itoa(o.Id))
}

type PlanDict map[int]Plan

func (o PlanDict) MarshalJSON() ([]byte, error) {
	m := map[string]Plan{}

	for k, v := range o {
		m[strconv.Itoa(k)] = v
	}

	return json.Marshal(m)
}

func (o *PlanDict) UnmarshalJSON(d []byte) error {
	*o = PlanDict{}

	if string(d) == "[]" {
		return nil
	}

	m := map[string]Plan{}

	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	for k, v := range m {
		i, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		(*o)[i] = v
	}

	return nil
}

func (o PlanDict) Array() PlanArray {
	a := []Plan{}

	for _, v := range o {
		a = append(a, v)
	}

	return a
}

func (o PlanDict) String() string {
	l := []string{}

	l = append(l, fmt.Sprintf(planFormat, "NAME", "CPUS", "PRICE/MONTH",
		"ID"))
	l = append(l, strings.Repeat("-", 78))

	a := o.Array()
	sort.Sort(a)

	for _, r := range a {
		l = append(l, r.String())
	}

	return strings.Join(l, "\n")
}

type PlanArray []Plan

func (a PlanArray) Len() int {
	return len(a)
}

func (a PlanArray) Less(i, j int) bool {
	switch {
	case a[i].CPUs < a[j].CPUs:
		return true
	case a[i].CPUs > a[j].CPUs:
		return false
	default:
		switch {
		case a[i].RAM < a[j].RAM:
			return true
		case a[i].RAM > a[j].RAM:
			return false
		default:
			switch {
			case a[i].Disk < a[j].Disk:
				return true
			case a[i].Disk > a[j].Disk:
				return false
			default:
				return a[i].Name < a[j].Name
			}
		}
	}
}

func (a PlanArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
