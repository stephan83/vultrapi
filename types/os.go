package types

import (
	"encoding/json"
	"strconv"
)

type OS struct {
	Id      int    `json:"OSID"`
	Name    string `json:"name"`
	Arch    string `json:"arch"`
	Family  string `json:"family"`
	Windows bool   `json:"windows"`
}

type OSMap map[int]OS

func (o OSMap) MarshalJSON() ([]byte, error) {
	m := map[string]OS{}

	for k, v := range o {
		m[strconv.Itoa(k)] = v
	}

	return json.Marshal(m)
}

func (o *OSMap) UnmarshalJSON(d []byte) error {
	*o = OSMap{}

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

func (o OSMap) Array() OSArray {
	a := []OS{}

	for _, v := range o {
		a = append(a, v)
	}

	return a
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
