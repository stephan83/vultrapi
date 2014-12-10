package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
)

type destroySSHKey struct{}

func NewDestroySSHKey() Command {
	return destroySSHKey{}
}

func (_ destroySSHKey) NeedsKey() bool {
	return true
}

func (_ destroySSHKey) Args() string {
	return "ssh_key_id"
}

func (_ destroySSHKey) Desc() string {
	return "Destroys an SSH key."
}

func (_ destroySSHKey) PrintOptions() {
	fmt.Println("None.")
}

func (_ destroySSHKey) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	if err != nil {
		err = ErrUsage{}
		return
	}

	err = requests.PostDestroySSHKey(c, key, args[0])
	if err != nil {
		return
	}

	fmt.Println("OK")

	return
}
