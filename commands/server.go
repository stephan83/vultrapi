package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"strconv"
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

	id, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}

	s, ok := sd[id]
	if !ok {
		err = ErrNotFound{}
		return
	}

	fmt.Println(s.Details())

	return
}
