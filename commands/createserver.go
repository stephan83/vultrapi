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

type createServer struct {
	flagSet *flag.FlagSet
	options requests.CreateServerOptions
}

func NewCreateServer() Command {
	o := createServer{
		flagSet: flag.NewFlagSet("createserver", flag.ContinueOnError),
	}

	o.flagSet.StringVar(&o.options.IPXEChainURL,
		"ipxe_chain_url", "",
		"IPXE chain url")
	o.flagSet.IntVar(&o.options.ISOId,
		"iso_id", 0,
		"ISO ID")
	o.flagSet.IntVar(&o.options.ScriptId,
		"script_id", 0,
		"Script ID")
	o.flagSet.IntVar(&o.options.SnapshotId,
		"snapshot_id", 0,
		"Snapshot ID")
	o.flagSet.BoolVar(&o.options.EnableIPV6,
		"enable_ipv6", false,
		"Enable IPV6")
	o.flagSet.BoolVar(&o.options.EnablePrivateNetwork,
		"enable_private_network", false,
		"Enable private network")
	o.flagSet.StringVar(&o.options.Label,
		"label", "",
		"Label")
	o.flagSet.StringVar(&o.options.SSHKeyId,
		"ssh_key_id", "",
		"SSH key ID")
	o.flagSet.BoolVar(&o.options.EnableAutoBackups,
		"enable_auto_backups", false,
		"Enable auto auto backups")

	return &o
}

func (_ *createServer) NeedsKey() bool {
	return true
}

func (_ *createServer) Args() string {
	return "region_id plan_id os_id"
}

func (_ *createServer) Desc() string {
	return "Creates a server."
}

func (o *createServer) PrintOptions() {
	o.flagSet.SetOutput(os.Stdout)
	o.flagSet.PrintDefaults()
	o.flagSet.SetOutput(os.Stderr)
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

	err = o.flagSet.Parse(args[3:])
	if err != nil {
		err = ErrUsage{}
		return
	}

	id, err := requests.PostCreateServer(c, key, region, plan, os,
		o.options)

	if err != nil {
		return
	}

	fmt.Printf("SERVER ID:\t%d\n", id)

	return
}
