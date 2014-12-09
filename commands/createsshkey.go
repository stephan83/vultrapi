package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io/ioutil"
)

type createSSHKey struct {}

func NewCreateSSHKey() Command {
	return createSSHKey{}
}

func (_ createSSHKey) NeedsKey() bool {
	return true
}

func (_ createSSHKey) Args() string {
	return "name path_to_public_ssh_key"
}

func (_ createSSHKey) Desc() string {
	return "Creates an SSH key."
}

func (cs createSSHKey) PrintOptions() {
	fmt.Println("None.")
}

func (s createSSHKey) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 2 {
		err = ErrUsage{}
		return
	}

	name := args[0]
	keyPath := args[1]

	sshkey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	id, err := requests.PostCreateSSHKey(c, key, name, string(sshkey))
	if err != nil {
		return
	}

	fmt.Printf("SSH KEY ID: %s\n", id)

	return
}
