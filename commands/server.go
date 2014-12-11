package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"strconv"
	"text/tabwriter"
	"io"
)

type server struct{ BasicCommand }

func NewServer() Command {
	return &server{
		BasicCommand{
			Desc:        "Get server information.",
			NeedsKey:    true,
			ArgsDesc:    "server_id",
			OptionsDesc: "",
		},
	}
}

func (_ *server) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
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

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintf(t, "ID\t%d\n", s.Id)
	fmt.Fprintf(t, "OS\t%s\n", s.OS)
	fmt.Fprintf(t, "RAM\t%s\n", s.Disk)
	fmt.Fprintf(t, "DISK\t%s\n", s.IPV4)
	fmt.Fprintf(t, "CPUS\t%d\n", s.CPUs)
	fmt.Fprintf(t, "LOCATION\t%s\n", s.Location)
	fmt.Fprintf(t, "REGION ID\t%d\n", s.RegionId)
	fmt.Fprintf(t, "DEFAULT PASSWORD\t%s\n", s.DefaultPassword)
	fmt.Fprintf(t, "DATE CREATED\t%s\n", s.DateCreated)
	fmt.Fprintf(t, "PENDING CHARGES\t%.2f\n", s.PendingCharges)
	fmt.Fprintf(t, "STATUS\t%s\n", s.Status)
	fmt.Fprintf(t, "PRICE/MONTH\t%.2f\n", s.PricePerMonth)
	fmt.Fprintf(t, "CURRENT BANDWIDTH GB\t%.2f\n", s.CurrentBandwidthGB)
	fmt.Fprintf(t, "ALLOWED BANDWIDTH GB\t%.2f\n", s.AllowedBandwidthGB)
	fmt.Fprintf(t, "IPV4 NETMASK\t%s\n", s.IPV4Netmask)
	fmt.Fprintf(t, "IPV4 GATEWAY\t%s\n", s.IPV4Gateway)
	fmt.Fprintf(t, "POWER STATUS\t%s\n", s.PowerStatus)
	fmt.Fprintf(t, "PLAN ID\t%d\n", s.PlanId)
	fmt.Fprintf(t, "IPV6 NETWORK\t%s\n", s.IPV6Network)
	fmt.Fprintf(t, "IPV6\t%s\n", s.IPV6)
	fmt.Fprintf(t, "IPV6 NETWORK SIZE\t%d\n", s.IPV6NetworkSize)
	fmt.Fprintf(t, "LABEL\t%s\n", s.Label)
	fmt.Fprintf(t, "PRIVATE IP\t%s\n", s.PrivateIP)
	fmt.Fprintf(t, "KVM URL\t%s\n", s.KVMURL)
	fmt.Fprintf(t, "AUTO BACKUPS\t%s\n", strconv.FormatBool(s.AutoBackups))

	t.Flush()

	return
}
