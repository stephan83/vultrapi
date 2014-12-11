package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
)

type destroySSHKey struct{Command}

func NewDestroySSHKey() *destroySSHKey {
	return &destroySSHKey{
		Command {
			Desc: "Destroys an SSH key.",
			NeedsKey: true,
			ArgsDesc: "ssh_key_id",
			OptionsDesc: "",
		},
	}
}

func (_ *destroySSHKey) Exec(c Client, args []string, key string) (err error) {
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
