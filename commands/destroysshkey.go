package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
)

type destroySSHKey struct{ BasicCommand }

func NewDestroySSHKey() Command {
	return &destroySSHKey{
		BasicCommand{
			Desc:        "Destroys an SSH key.",
			NeedsKey:    true,
			ArgsDesc:    "ssh_key_id",
			OptionsDesc: "",
		},
	}
}

func (_ *destroySSHKey) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
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

	fmt.Fprintln(w, "OK")

	return
}
