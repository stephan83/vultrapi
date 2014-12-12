package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"strconv"
	"text/tabwriter"
)

type script struct{ BasicCommand }

func NewScript() Command {
	return &script{
		BasicCommand{
			Desc:        "Get script information.",
			NeedsKey:    true,
			ArgsDesc:    "script_id",
			OptionsDesc: "",
		},
	}
}

func (_ *script) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}
	r, err := requests.GetScripts(c, key)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}

	o, ok := r[id]
	if !ok {
		err = ErrNotFound{}
		return
	}

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintf(t, "ID\t%d\n", o.Id)
	fmt.Fprintf(t, "NAME\t%s\n", o.Name)
	fmt.Fprintf(t, "DATE CREATED\t%s\n", o.DateCreated)
	fmt.Fprintf(t, "DATE MODIFIED\t%s\n", o.DateModified)
	fmt.Fprintf(t, "SCRIPT\t%s\n", o.Script)

	t.Flush()

	return
}
