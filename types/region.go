package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const regionFormat = "%-20s | %-7s | %-5s | %-30s | %s"

type Region struct {
	Id        int    `json:"DCID,string"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Continent string `json:"continent"`
	State     string `json:"state"`
}

func (o Region) String() string {
	return fmt.Sprintf(regionFormat, o.Continent, o.Country, o.State,
		o.Name, strconv.Itoa(o.Id))
}

type RegionMap map[int]Region

func (o RegionMap) MarshalJSON() ([]byte, error) {
	m := map[string]Region{}

	for k, v := range o {
		m[strconv.Itoa(k)] = v
	}

	return json.Marshal(m)
}

func (o *RegionMap) UnmarshalJSON(d []byte) error {
	*o = RegionMap{}

	if string(d) == "[]" {
		return nil
	}

	m := map[string]Region{}

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

func (o RegionMap) Array() RegionArray {
	a := []Region{}

	for _, v := range o {
		a = append(a, v)
	}

	return a
}

func (o RegionMap) String() string {
	l := []string{}

	l = append(l, fmt.Sprintf(regionFormat, "CONTINENT", "COUNTRY",
		"STATE", "NAME", "ID"))
	l = append(l, strings.Repeat("-", 78))

	a := o.Array()
	sort.Sort(a)

	for _, r := range a {
		l = append(l, r.String())
	}

	return strings.Join(l, "\n")
}

type RegionArray []Region

func (a RegionArray) Len() int {
	return len(a)
}

func (a RegionArray) Less(i, j int) bool {
	switch {
	case a[i].Continent < a[j].Continent:
		return true
	case a[i].Continent > a[j].Continent:
		return false
	default:
		switch {
		case a[i].Country < a[j].Country:
			return true
		case a[i].Country > a[j].Country:
			return false
		default:
			switch {
			case a[i].State < a[j].State:
				return true
			case a[i].State > a[j].State:
				return false
			default:
				return a[i].Name < a[j].Name
			}
		}
	}
}

func (a RegionArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
