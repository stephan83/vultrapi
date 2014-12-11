package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
)

type destroySnapshot struct{}

func NewDestroySnapshot() Command {
	return destroySnapshot{}
}

func (_ destroySnapshot) NeedsKey() bool {
	return true
}

func (_ destroySnapshot) Args() string {
	return "snapshot_id"
}

func (_ destroySnapshot) Desc() string {
	return "Destroys a snapshot."
}

func (_ destroySnapshot) PrintOptions() {
	fmt.Println("None.")
}

func (_ destroySnapshot) Exec(c Client, args []string, key string) (err error) {
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
