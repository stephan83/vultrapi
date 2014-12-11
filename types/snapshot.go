package types

import (
	"encoding/json"
	"time"
)

type Snapshot struct {
	Id          string `json:"SNAPSHOTID"`
	DateCreated Date   `json:"date_created"`
	Description string `json:"description"`
	Size        int64  `json:"size,string"`
	Status      string `json:"status"`
}

type SnapshotMap map[string]Snapshot

func (o *SnapshotMap) UnmarshalJSON(d []byte) error {
	*o = SnapshotMap{}

	if string(d) == "[]" {
		return nil
	}

	m := map[string]Snapshot{}

	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	for k, v := range m {
		(*o)[k] = v
	}

	return nil
}

func (o SnapshotMap) Array() SnapshotArray {
	a := []Snapshot{}

	for _, v := range o {
		a = append(a, v)
	}

	return a
}

type SnapshotArray []Snapshot

func (a SnapshotArray) Len() int {
	return len(a)
}

func (a SnapshotArray) Less(i, j int) bool {
	switch {
	case a[i].Status < a[j].Status:
		return true
	case a[i].Status > a[j].Status:
		return false
	default:
		switch {
		case a[i].Description < a[j].Description:
			return true
		case a[i].Description > a[j].Description:
			return false
		default:
			switch {
			case time.Time(a[i].DateCreated).Before(time.Time(a[j].DateCreated)):
				return true
			case time.Time(a[i].DateCreated).Before(time.Time(a[j].DateCreated)):
				return false
			default:
				return a[i].Size < a[j].Size
			}
		}
	}
}

func (a SnapshotArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
