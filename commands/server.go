package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
)

type server struct{}

func NewServer() Command {
	return server{}
}

func (_ server) NeedsKey() bool {
	return true
}

func (_ server) Args() string {
	return "server_id"
}

func (_ server) Desc() string {
	return "Get server information."
}

func (_ server) PrintOptions() {
	fmt.Println("None.")
}

func (_ server) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}
	sd, err := requests.GetServers(c, key)
	if err != nil {
		return
	}

	s, ok := sd[args[0]]
	if !ok {
		err = ErrNotFound{}
		return
	}

	fmt.Println(s.Details())

	return
}
