package commands

import (
	"flag"
	"fmt"
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
	cs := createServer{
		flagSet: flag.NewFlagSet("createserver", flag.ExitOnError),
	}

	cs.flagSet.StringVar(&cs.options.IPXEChainURL,
		"ipxe_chain_url", "",
		"IPXE chain url")
	cs.flagSet.IntVar(&cs.options.ISOId,
		"iso_id", 0,
		"ISO ID")
	cs.flagSet.IntVar(&cs.options.ScriptId,
		"script_id", 0,
		"Script ID")
	cs.flagSet.IntVar(&cs.options.SnapshotId,
		"snapshot_id", 0,
		"Snapshot ID")
	cs.flagSet.BoolVar(&cs.options.EnableIPV6,
		"enable_ipv6", false,
		"Enable IPV6")
	cs.flagSet.BoolVar(&cs.options.EnablePrivateNetwork,
		"enable_private_network", false,
		"Enable private network")
	cs.flagSet.StringVar(&cs.options.Label,
		"label", "",
		"Label")
	cs.flagSet.IntVar(&cs.options.SSHKeyId,
		"ssh_key_id", 0,
		"SSH key ID")
	cs.flagSet.BoolVar(&cs.options.EnableAutoBackups,
		"enable_auto_backups", false,
		"Enable auto auto backups")

	return &cs
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

func (cs *createServer) PrintOptions() {
	cs.flagSet.PrintDefaults()
}

func (s *createServer) Exec() (err error) {
	if len(os.Args) < 5 {
		err = ErrUsage{}
		return
	}

	regionId, err := strconv.Atoi(os.Args[2])
	if err != nil {
		err = ErrUsage{}
		return
	}
	planId, err := strconv.Atoi(os.Args[3])
	if err != nil {
		err = ErrUsage{}
		return
	}
	OSId, err := strconv.Atoi(os.Args[4])
	if err != nil {
		err = ErrUsage{}
		return
	}

	err = s.flagSet.Parse(os.Args[5:])
	if err != nil {
		return
	}

	id, err := requests.PostCreateServer(os.Getenv("VULTR_API_KEY"),
		regionId, planId, OSId, s.options)

	if err != nil {
		return
	}

	fmt.Printf("SERVER ID: %d\n", id)

	return
}
