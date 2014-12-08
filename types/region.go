package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Region struct {
	Id        int    `json:"DCID,string"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Continent string `json:"continent"`
	State     string `json:"state"`
}

const (
	regionFormat = "%-20s | %-7s | %-5s | %-30s | %-4s"
)

func (r Region) String() string {
	return fmt.Sprintf(regionFormat, r.Continent, r.Country, r.State,
		r.Name, strconv.Itoa(r.Id))
}

type RegionDict map[string]Region

func (rd RegionDict) Slice() RegionSlice {
	regions := []Region{}

	for _, r := range rd {
		regions = append(regions, r)
	}

	return regions
}

func (rd RegionDict) String() string {
	lines := []string{}

	lines = append(lines, fmt.Sprintf(regionFormat, "CONTINENT",
		"COUNTRY", "STATE", "NAME", "ID"))

	lines = append(lines, strings.Repeat("-", 78))

	rs := rd.Slice()
	sort.Sort(rs)

	for _, r := range rs {
		lines = append(lines, r.String())
	}

	return strings.Join(lines, "\n")
}

type RegionSlice []Region

func (rs RegionSlice) Len() int {
	return len(rs)
}

func (rs RegionSlice) Less(i, j int) bool {
	switch {
	case rs[i].Continent < rs[j].Continent:
		return true
	case rs[i].Continent > rs[j].Continent:
		return false
	default:
		switch {
		case rs[i].Country < rs[j].Country:
			return true
		case rs[i].Country > rs[j].Country:
			return false
		default:
			switch {
			case rs[i].State < rs[j].State:
				return true
			case rs[i].State > rs[j].State:
				return false
			default:
				return rs[i].Name < rs[j].Name
			}
		}
	}
}

func (rs RegionSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
