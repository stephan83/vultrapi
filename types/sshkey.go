package types

import (
	"fmt"
	"sort"
	"strings"
	"encoding/json"
)

type SSHKey struct {
	Id          string `json:"SSHKEYID"`
	Name        string `json:"name"`
	DateCreated string `json:"date_created"`
	Key         string `json:"ssh_key"`
}

const (
	keyFormat = "%-39s | %-19s | %s"
)

func (s SSHKey) String() string {
	return fmt.Sprintf(keyFormat, s.Name, s.DateCreated, s.Id)
}

func (s SSHKey) Details() string {
	return strings.Join([]string{
		fmt.Sprintf("%12s: %s", "ID", s.Id),
		fmt.Sprintf("%12s: %s", "NAME", s.Name),
		fmt.Sprintf("%12s: %s", "DATE CREATED", s.DateCreated),
		fmt.Sprintf("%12s: %s", "KEY", s.Key),
	}, "\n")
}

type SSHKeyDict map[string]SSHKey

func (d *SSHKeyDict) UnmarshalJSON(data []byte) error {
	*d = SSHKeyDict{}
	
	if string(data) == "[]" {
		return nil
	}

	m := map[string]SSHKey{}

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	for s, v := range m {
		(*d)[s] = v
	}

	return nil
}

func (sd SSHKeyDict) Slice() SSHKeySlice {
	keys := []SSHKey{}

	for _, s := range sd {
		keys = append(keys, s)
	}

	return keys
}

func (sd SSHKeyDict) String() string {
	lines := []string{}

	lines = append(lines, fmt.Sprintf(keyFormat, "NAME",
		"DATE CREATED", "ID"))

	lines = append(lines, strings.Repeat("-", 78))

	ss := sd.Slice()
	sort.Sort(ss)

	for _, s := range ss {
		lines = append(lines, s.String())
	}

	return strings.Join(lines, "\n")
}

type SSHKeySlice []SSHKey

func (ss SSHKeySlice) Len() int {
	return len(ss)
}

func (ss SSHKeySlice) Less(i, j int) bool {
	switch {
	case ss[i].Name < ss[j].Name:
		return true
	case ss[i].Name > ss[j].Name:
		return false
	default:
		switch {
		case ss[i].DateCreated < ss[j].DateCreated:
			return true
		case ss[i].DateCreated > ss[j].DateCreated:
			return false
		default:
			return ss[i].Id < ss[j].Id
		}
	}
}

func (ss SSHKeySlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
