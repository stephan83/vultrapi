package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"strconv"
	"text/tabwriter"
	"flag"
)

type script struct{
	BasicCommandWithOptions
	Labels bool
	Fields StringSlice
}

func NewScript() Command {
	f := flag.NewFlagSet("script", flag.ContinueOnError)

	o := script{
		BasicCommandWithOptions: BasicCommandWithOptions{
			BasicCommand{
				Desc:     "Get script information.",
				NeedsKey: true,
				ArgsDesc: "script_id",
			},
			f,
		},
	}

	f.BoolVar(&o.Labels, "labels", true, "display labels")
	f.Var(&o.Fields, "field", "only display the specified field (can be set multiple times)")

	o.Initialize()

	o.OptionsDesc += "\nAvailable fields: id, name, date_created, date_modified, script"

	return &o
}

func (o *script) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	o.FlagSet.SetOutput(w)

	err = o.FlagSet.Parse(args[1:])
	if err != nil {
		return ErrUsage{}
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

	e, ok := r[id]
	if !ok {
		err = ErrNotFound{}
		return
	}

	fields := map[string]bool{}
	for _, v := range o.Fields {
		fields[v] = true
	}

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	if fields["id"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "ID\t%d\n", e.Id)
		} else {
			fmt.Fprintf(t, "%d\n", e.Id)
		}
	}
	if fields["name"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "NAME\t%s\n", e.Name)
		} else {
			fmt.Fprintf(t, "%s\n", e.Name)
		}
	}
	if fields["date_created"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "DATE CREATED\t%s\n", e.DateCreated)
		} else {
			fmt.Fprintf(t, "%s\n", e.DateCreated)
		}
	}
	if fields["date_modified"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "DATE MODIFIED\t%s\n", e.DateModified)
		} else {
			fmt.Fprintf(t, "%s\n", e.DateModified)
		}
	}
	if fields["script"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "SCRIPT\t%s\n", e.Script)
		} else {
			fmt.Fprintf(t, "%s\n", e.Script)
		}
	}

	t.Flush()

	return
}
