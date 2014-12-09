package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type OS struct {
	Id      int    `json:"OSID"`
	Name    string `json:"name"`
	Arch    string `json:"arch"`
	Family  string `json:"family"`
	Windows bool   `json:"windows"`
}

const (
	osFormat = "%-16s | %-4s | %-45s | %s"
)

func (o OS) String() string {
	return fmt.Sprintf(osFormat, o.Family, o.Arch, o.Name,
		strconv.Itoa(o.Id))
}

type OSDict map[string]OS

func (od OSDict) Slice() OSSlice {
	OSs := []OS{}

	for _, o := range od {
		OSs = append(OSs, o)
	}

	return OSs
}

func (od OSDict) String() string {
	lines := []string{}

	lines = append(lines, fmt.Sprintf(osFormat, "FAMILY",
		"ARCH", "NAME", "ID"))

	lines = append(lines, strings.Repeat("-", 78))

	os := od.Slice()
	sort.Sort(os)

	for _, o := range os {
		lines = append(lines, o.String())
	}

	return strings.Join(lines, "\n")
}

type OSSlice []OS

func (os OSSlice) Len() int {
	return len(os)
}

func (os OSSlice) Less(i, j int) bool {
	switch {
	case os[i].Family < os[j].Family:
		return true
	case os[i].Family > os[j].Family:
		return false
	default:
		switch {
		case os[i].Arch < os[j].Arch:
			return true
		case os[i].Arch > os[j].Arch:
			return false
		default:
			return os[i].Name < os[j].Name
		}
	}
}

func (os OSSlice) Swap(i, j int) {
	os[i], os[j] = os[j], os[i]
}
