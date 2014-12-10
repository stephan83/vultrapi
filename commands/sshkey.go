package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
)

type sshKey struct{}

func NewSSHKey() Command {
	return sshKey{}
}

func (_ sshKey) NeedsKey() bool {
	return true
}

func (_ sshKey) Args() string {
	return "ssh_key_id"
}

func (_ sshKey) Desc() string {
	return "Get SSH key information."
}

func (_ sshKey) PrintOptions() {
	fmt.Println("None.")
}

func (_ sshKey) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}
	sd, err := requests.GetSSHKeys(c, key)
	if err != nil {
		return
	}

	s, ok := sd[args[0]]
	if !ok {
		err = ErrNotFound{}
		return
	}

	fmt.Println(s.Details())

	return
}
