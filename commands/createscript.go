package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	. "github.com/stephan83/vultrapi/types"
	"io"
	"io/ioutil"
)

type createScript struct{ BasicCommand }

func NewCreateScript() Command {
	return &createScript{
		BasicCommand{
			Desc:        "Create a script.",
			NeedsKey:    true,
			ArgsDesc:    "[boot | pxe] name path_to_script",
			OptionsDesc: "",
		},
	}
}

func (_ *createScript) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 3 {
		err = ErrUsage{}
		return
	}

	var scriptType ScriptType
	switch args[0] {
	case "boot":
		scriptType = Boot
	case "pxe":
		scriptType = PXE
	default:
		return ErrUsage{}
	}

	name := args[1]
	path := args[2]

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	id, err := requests.PostCreateScript(c, key, scriptType, name, string(d))
	if err != nil {
		return
	}

	fmt.Fprintf(w, "SCRIPT ID:\t%s\n", id)

	return
}
