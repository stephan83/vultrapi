package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"text/tabwriter"
)

type sshKey struct {
	BasicCommandWithOptions
	Labels bool
	Fields StringSlice
}

func NewSSHKey() Command {
	f := flag.NewFlagSet("sshkey", flag.ContinueOnError)

	o := sshKey{
		BasicCommandWithOptions: BasicCommandWithOptions{
			BasicCommand{
				Desc:     "Get SSH key information.",
				NeedsKey: true,
				ArgsDesc: "ssh_key_id",
			},
			f,
		},
	}

	f.BoolVar(&o.Labels, "labels", true, "display labels")
	f.Var(&o.Fields, "field", "only display the specified field (can be set multiple times)")

	o.Initialize()

	o.OptionsDesc += "\nAvailable fields: id, name, date_created, key"

	return &o
}

func (o *sshKey) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	o.FlagSet.SetOutput(w)

	err = o.FlagSet.Parse(args[1:])
	if err != nil {
		return ErrUsage{}
	}

	r, err := requests.GetSSHKeys(c, key)
	if err != nil {
		return
	}

	e, ok := r[args[0]]
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
			fmt.Fprintf(t, "ID\t%s\n", e.Id)
		} else {
			fmt.Fprintf(t, "%s\n", e.Id)
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
	if fields["key"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "KEY\t%s\n", e.Key)
		} else {
			fmt.Fprintf(t, "%s\n", e.Key)
		}
	}

	t.Flush()

	return
}
