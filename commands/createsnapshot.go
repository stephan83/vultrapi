package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"strconv"
)

type createSnapshot struct {
	flagSet *flag.FlagSet
	desc    string
}

func NewCreateSnapshot() Command {
	o := createSnapshot{
		flagSet: flag.NewFlagSet("createsnapshot", flag.ContinueOnError),
	}

	o.flagSet.StringVar(&o.desc,
		"description", "",
		"Description")

	return &o
}

func (_ *createSnapshot) NeedsKey() bool {
	return true
}

func (_ *createSnapshot) Args() string {
	return "server_id"
}

func (_ *createSnapshot) Desc() string {
	return "Creates a snapshot."
}

func (o *createSnapshot) PrintOptions() {
	o.flagSet.SetOutput(os.Stdout)
	o.flagSet.PrintDefaults()
	o.flagSet.SetOutput(os.Stderr)
}

func (o *createSnapshot) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	sid, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}

	err = o.flagSet.Parse(args[1:])
	if err != nil {
		err = ErrUsage{}
		return
	}

	id, err := requests.PostCreateSnapshot(c, key, sid, o.desc)
	if err != nil {
		return
	}

	fmt.Printf("SNAPSHOT ID:\t%s\n", id)

	return
}
