package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
)

type destroySnapshot struct{BasicCommand}

func NewDestroySnapshot() *destroySnapshot {
	return &destroySnapshot{
		BasicCommand {
			Desc: "Destroys a snapshot.",
			NeedsKey: true,
			ArgsDesc: "snapshot_id",
			OptionsDesc: "",
		},
	}
}

func (_ *destroySnapshot) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	if err != nil {
		err = ErrUsage{}
		return
	}

	err = requests.PostDestroySnapshot(c, key, args[0])
	if err != nil {
		return
	}

	fmt.Println("OK")

	return
}
