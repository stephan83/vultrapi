package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io/ioutil"
	"io"
)

type createSSHKey struct{ BasicCommand }

func NewCreateSSHKey() Command {
	return &createSSHKey{
		BasicCommand{
			Desc:        "Creates an SSH key.",
			NeedsKey:    true,
			ArgsDesc:    "name path_to_public_ssh_key",
			OptionsDesc: "",
		},
	}
}

func (_ *createSSHKey) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 2 {
		err = ErrUsage{}
		return
	}

	name := args[0]
	path := args[1]

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	id, err := requests.PostCreateSSHKey(c, key, name, string(d))
	if err != nil {
		return
	}

	fmt.Fprintf(w, "SSH KEY ID:\t%s\n", id)

	return
}
