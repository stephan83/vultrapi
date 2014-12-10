package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

const keyFormat = "%-33s | %-25s | %s"

type SSHKey struct {
	Id          string `json:"SSHKEYID"`
	Name        string `json:"name"`
	DateCreated Date   `json:"date_created"`
	Key         string `json:"ssh_key"`
}

func (o SSHKey) String() string {
	return fmt.Sprintf(keyFormat, o.Name, o.DateCreated, o.Id)
}

func (o SSHKey) Details() string {
	return strings.Join([]string{
		fmt.Sprintf("%12s: %s", "ID", o.Id),
		fmt.Sprintf("%12s: %s", "NAME", o.Name),
		fmt.Sprintf("%12s: %s", "DATE CREATED", o.DateCreated),
		fmt.Sprintf("%12s: %s", "KEY", o.Key),
	}, "\n")
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

func (o SSHKeyMap) String() string {
	l := []string{}

	l = append(l, fmt.Sprintf(keyFormat, "NAME", "DATE CREATED", "ID"))
	l = append(l, strings.Repeat("-", 78))

	a := o.Array()
	sort.Sort(a)

	for _, s := range a {
		l = append(l, s.String())
	}

	return strings.Join(l, "\n")
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
