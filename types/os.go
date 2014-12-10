package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"encoding/json"
)

const osFormat = "%-16s | %-4s | %-45s | %s"

type OS struct {
	Id      int    `json:"OSID"`
	Name    string `json:"name"`
	Arch    string `json:"arch"`
	Family  string `json:"family"`
	Windows bool   `json:"windows"`
}

func (o OS) String() string {
	return fmt.Sprintf(osFormat, o.Family, o.Arch, o.Name,
		strconv.Itoa(o.Id))
}

type OSDict map[int]OS

func (o OSDict) MarshalJSON() ([]byte, error) {
	m := map[string]OS{}

	for k, v := range o {
		m[strconv.Itoa(k)] = v
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

	for k, v := range m {
		i, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		(*o)[i] = v
	}

	return nil
}

func (o OSDict) Array() OSArray {
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

	a := o.Array()
	sort.Sort(a)

	for _, v := range a {
		l = append(l, v.String())
	}

	return strings.Join(l, "\n")
}

type OSArray []OS

func (a OSArray) Len() int {
	return len(a)
}

func (a OSArray) Less(i, j int) bool {
	switch {
	case a[i].Family < a[j].Family:
		return true
	case a[i].Family > a[j].Family:
		return false
	default:
		switch {
		case a[i].Arch < a[j].Arch:
			return true
		case a[i].Arch > a[j].Arch:
			return false
		default:
			return a[i].Name < a[j].Name
		}
	}
}

func (a OSArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
