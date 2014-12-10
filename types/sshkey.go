package types

import (
	"encoding/json"
	"time"
)

type SSHKey struct {
	Id          string `json:"SSHKEYID"`
	Name        string `json:"name"`
	DateCreated Date   `json:"date_created"`
	Key         string `json:"ssh_key"`
}

type SSHKeyMap map[string]SSHKey

func (o *SSHKeyMap) UnmarshalJSON(d []byte) error {
	*o = SSHKeyMap{}

	if string(d) == "[]" {
		return nil
	}

	m := map[string]SSHKey{}

	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	for k, v := range m {
		(*o)[k] = v
	}

	return nil
}

func (o SSHKeyMap) Array() SSHKeyArray {
	keys := []SSHKey{}

	for _, s := range o {
		keys = append(keys, s)
	}

	return keys
}

type SSHKeyArray []SSHKey

func (a SSHKeyArray) Len() int {
	return len(a)
}

func (a SSHKeyArray) Less(i, j int) bool {
	switch {
	case a[i].Name < a[j].Name:
		return true
	case a[i].Name > a[j].Name:
		return false
	default:
		switch {
		case time.Time(a[i].DateCreated).Before(time.Time(a[j].DateCreated)):
			return true
		case time.Time(a[i].DateCreated).Before(time.Time(a[j].DateCreated)):
			return false
		default:
			return a[i].Id < a[j].Id
		}
	}
}

func (a SSHKeyArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
