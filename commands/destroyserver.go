package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"strconv"
	"io"
)

type destroyServer struct{ BasicCommand }

func NewDestroyServer() Command {
	return &destroyServer{
		BasicCommand{
			Desc:        "Destroy a server.",
			NeedsKey:    true,
			ArgsDesc:    "server_id",
			OptionsDesc: "",
		},
	}
}

func (_ *destroyServer) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}

	err = requests.PostDestroyServer(c, key, id)
	if err != nil {
		return
	}

	fmt.Fprintln(w, "OK")

	return
}
