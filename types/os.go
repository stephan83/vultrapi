package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"encoding/json"
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

type OSDict map[int]OS

func (o OSDict) MarshalJSON() ([]byte, error) {
	m := map[string]OS{}

	for i, v := range o {
		m[strconv.Itoa(i)] = v
	}

	return json.Marshal(m)
}

func (o *OSDict) UnmarshalJSON(d []byte) error {
	*o = OSDict{}

	if string(d) == "[]" {
		return nil
	}

	m := map[string]OS{}

	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	for s, v := range m {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		(*o)[i] = v
	}

	return nil
}

func (o OSDict) Slice() OSSlice {
	a := []OS{}

	for _, v := range o {
		a = append(a, v)
	}

	return a
}

func (o OSDict) String() string {
	l := []string{}

	l = append(l, fmt.Sprintf(osFormat, "FAMILY", "ARCH", "NAME", "ID"))
	l = append(l, strings.Repeat("-", 78))

	a := o.Slice()
	sort.Sort(a)

	for _, v := range a {
		l = append(l, v.String())
	}

	return strings.Join(l, "\n")
}

type OSSlice []OS

func (o OSSlice) Len() int {
	return len(o)
}

func (o OSSlice) Less(i, j int) bool {
	switch {
	case o[i].Family < o[j].Family:
		return true
	case o[i].Family > o[j].Family:
		return false
	default:
		switch {
		case o[i].Arch < o[j].Arch:
			return true
		case o[i].Arch > o[j].Arch:
			return false
		default:
			return o[i].Name < o[j].Name
		}
	}
}

func (o OSSlice) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
