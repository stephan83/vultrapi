package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"strconv"
)

type createSnapshot struct {
	BasicCommandWithOptions
	description    string
}

func NewCreateSnapshot() *createSnapshot {
	f := flag.NewFlagSet("createsnapshot", flag.ContinueOnError)

	o := createSnapshot{
		BasicCommandWithOptions{
			BasicCommand{
				Desc: "Creates a snapshot.",
				NeedsKey: true,
				ArgsDesc: "server_id",
			},
			f,
		},
		"",
	}

	f.StringVar(&o.description, "description", "", "Description")

	o.Initialize()

	return &o
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

	err = o.FlagSet.Parse(args[1:])
	if err != nil {
		err = ErrUsage{}
		return
	}

	id, err := requests.PostCreateSnapshot(c, key, sid, o.description)
	if err != nil {
		return
	}

	fmt.Printf("SNAPSHOT ID:\t%s\n", id)

	return
}
