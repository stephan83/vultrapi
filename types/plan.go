package types

import (
	"encoding/json"
	"strconv"
)

type Plan struct {
	Id            int     `json:"VPSPLANID,string"`
	Name          string  `json:"name"`
	CPUs          int     `json:"vcpu_count,string"`
	RAM           int     `json:"ram,string"`
	Disk          int     `json:"disk,string"`
	Bandwidth     float64 `json:"bandwidth,string"`
	PricePerMonth float64 `json:"price_per_month,string"`
	Windows       bool    `json:"windows"`
}

type PlanMap map[int]Plan

func (o PlanMap) MarshalJSON() ([]byte, error) {
	m := map[string]Plan{}

	for k, v := range o {
		m[strconv.Itoa(k)] = v
	}

	return json.Marshal(m)
}

func (o *PlanMap) UnmarshalJSON(d []byte) error {
	*o = PlanMap{}

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

func (o PlanMap) Array() PlanArray {
	a := []Plan{}

	for _, v := range o {
		a = append(a, v)
	}

	return a
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
