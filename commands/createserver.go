package commands

import (
	"flag"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"strconv"
)

type createServer struct {
	CommandWithOptions
	options requests.CreateServerOptions
}

func NewCreateServer() *createServer {
	f := flag.NewFlagSet("createserver", flag.ContinueOnError)

	o := createServer{
		CommandWithOptions{
			Command{
				Desc: "Creates a server.",
				NeedsKey: true,
				ArgsDesc: "region_id plan_id os_id",
			},
			f,
		},
		requests.CreateServerOptions{},
	}

	f.StringVar(&o.options.IPXEChainURL, "ipxe_chain_url", "", "IPXE chain url")
	f.IntVar(&o.options.ISOId, "iso_id", 0, "ISO ID")
	f.IntVar(&o.options.ScriptId, "script_id", 0, "Script ID")
	f.IntVar(&o.options.SnapshotId, "snapshot_id", 0, "Snapshot ID")
	f.BoolVar(&o.options.EnableIPV6, "enable_ipv6", false, "Enable IPV6")
	f.BoolVar(&o.options.EnablePrivateNetwork, "enable_private_network", false, "Enable private network")
	f.StringVar(&o.options.Label, "label", "", "Label")
	f.StringVar(&o.options.SSHKeyId, "ssh_key_id", "", "SSH key ID")
	f.BoolVar(&o.options.EnableAutoBackups, "enable_auto_backups", false, "Enable auto auto backups")

	o.Initialize()

	return &o
}

func (o *createServer) Exec(c Client, args []string, key string) (err error) {
	if len(args) < 3 {
		err = ErrUsage{}
		return
	}

	region, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}
	plan, err := strconv.Atoi(args[1])
	if err != nil {
		err = ErrUsage{}
		return
	}
	os, err := strconv.Atoi(args[2])
	if err != nil {
		err = ErrUsage{}
		return
	}

	err = o.FlagSet.Parse(args[3:])
	if err != nil {
		err = ErrUsage{}
		return
	}

	id, err := requests.PostCreateServer(c, key, region, plan, os, o.options)
	if err != nil {
		return
	}

	fmt.Printf("SERVER ID:\t%d\n", id)

	return
}
