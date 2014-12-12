package types

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type ScriptType int8

const (
	Boot ScriptType = iota
	PXE
)

func (o ScriptType) String() string {
	switch o {
	case Boot:
		return "boot"
	case PXE:
		return "pxe"
	}
	return ""
}

func (o ScriptType) MarshalJSON() ([]byte, error) {
	switch o {
	case Boot:
		return []byte(`"boot"`), nil
	case PXE:
		return []byte(`"pxe"`), nil
	}
	return nil, errors.New("Invalid script type.")
}

func (o *ScriptType) UnmarshalJSON(d []byte) error {
	switch string(d) {
	case `"boot"`:
		*o = Boot
		return nil
	case `"pxe"`:
		*o = PXE
		return nil
	}
	return errors.New("Invalid script type.")
}

type Script struct {
	Id           int        `json:"SCRIPTID,string"`
	DateCreated  Date       `json:"date_created"`
	DateModified Date       `json:"date_modified"`
	Name         string     `json:"name"`
	Type         ScriptType `json:"type"`
	Script       string     `json:"script"`
}

type ScriptMap map[int]Script

func (o ScriptMap) MarshalJSON() ([]byte, error) {
	m := map[string]Script{}

	for k, v := range o {
		m[strconv.Itoa(k)] = v
	}

	return json.Marshal(m)
}

func (o *ScriptMap) UnmarshalJSON(d []byte) error {
	*o = ScriptMap{}

	if string(d) == "[]" {
		return nil
	}

	m := map[string]Script{}

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

func (o ScriptMap) Array() ScriptArray {
	a := []Script{}

	for _, v := range o {
		a = append(a, v)
	}

	return a
}

type ScriptArray []Script

func (a ScriptArray) Len() int {
	return len(a)
}

func (a ScriptArray) Less(i, j int) bool {
	switch {
	case a[i].Type < a[j].Type:
		return true
	case a[i].Type > a[j].Type:
		return false
	default:
		switch {
		case a[i].Name < a[j].Name:
			return true
		case a[i].Name > a[j].Name:
			return false
		default:
			return time.Time(a[i].DateCreated).Before(time.Time(a[j].DateCreated))
		}
	}
}

func (a ScriptArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
