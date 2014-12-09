package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"strconv"
)

type destroyServer struct{}

func NewDestroyServer() Command {
	return destroyServer{}
}

func (_ destroyServer) NeedsKey() bool {
	return true
}

func (_ destroyServer) Args() string {
	return "server_id"
}

func (_ destroyServer) Desc() string {
	return "Destroys a server."
}

func (_ destroyServer) PrintOptions() {
	fmt.Println("None.")
}

func (_ destroyServer) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	serverId, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}

	err = requests.PostDestroyServer(c, key, serverId)

	if err != nil {
		return
	}

	fmt.Println("OK")

	return
}
