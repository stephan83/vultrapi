package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"strconv"
)

type destroyScript struct{ BasicCommand }

func NewDestroyScript() Command {
	return &destroyScript{
		BasicCommand{
			Desc:        "Destroy a script.",
			NeedsKey:    true,
			ArgsDesc:    "script_id",
			OptionsDesc: "",
		},
	}
}

func (_ *destroyScript) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}

	err = requests.PostDestroyScript(c, key, id)
	if err != nil {
		return
	}

	fmt.Fprintln(w, "OK")

	return
}
