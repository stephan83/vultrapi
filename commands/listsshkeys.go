package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
)

type listSSHKeys struct {}

func NewListSSHKeys() Command {
	return listSSHKeys{}
}

func (_ listSSHKeys) NeedsKey() bool {
	return true
}

func (_ listSSHKeys) Args() string {
	return ""
}

func (_ listSSHKeys) Desc() string {
	return "List all SSH keys."
}

func (ls listSSHKeys) PrintOptions() {
	fmt.Println("None.")
}

func (ls listSSHKeys) Exec(c Client, args []string, key string) (err error) {
	s, err := requests.GetSSHKeys(c, key)
	if err != nil {
		return
	}

	fmt.Println(s)

	return
}
