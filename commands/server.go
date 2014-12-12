package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"strconv"
	"text/tabwriter"
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
	r, err := requests.GetServers(c, key)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		err = ErrUsage{}
		return
	}

	o, ok := r[id]
	if !ok {
		err = ErrNotFound{}
		return
	}

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintf(t, "ID\t%d\n", o.Id)
	fmt.Fprintf(t, "OS\t%s\n", o.OS)
	fmt.Fprintf(t, "RAM\t%s\n", o.Disk)
	fmt.Fprintf(t, "DISK\t%s\n", o.IPV4)
	fmt.Fprintf(t, "CPUS\t%d\n", o.CPUs)
	fmt.Fprintf(t, "LOCATION\t%s\n", o.Location)
	fmt.Fprintf(t, "REGION ID\t%d\n", o.RegionId)
	fmt.Fprintf(t, "DEFAULT PASSWORD\t%s\n", o.DefaultPassword)
	fmt.Fprintf(t, "DATE CREATED\t%s\n", o.DateCreated)
	fmt.Fprintf(t, "PENDING CHARGES\t%.2f\n", o.PendingCharges)
	fmt.Fprintf(t, "STATUS\t%s\n", o.Status)
	fmt.Fprintf(t, "PRICE/MONTH\t%.2f\n", o.PricePerMonth)
	fmt.Fprintf(t, "CURRENT BANDWIDTH GB\t%.2f\n", o.CurrentBandwidthGB)
	fmt.Fprintf(t, "ALLOWED BANDWIDTH GB\t%.2f\n", o.AllowedBandwidthGB)
	fmt.Fprintf(t, "IPV4 NETMASK\t%s\n", o.IPV4Netmask)
	fmt.Fprintf(t, "IPV4 GATEWAY\t%s\n", o.IPV4Gateway)
	fmt.Fprintf(t, "POWER STATUS\t%s\n", o.PowerStatus)
	fmt.Fprintf(t, "PLAN ID\t%d\n", o.PlanId)
	fmt.Fprintf(t, "IPV6 NETWORK\t%s\n", o.IPV6Network)
	fmt.Fprintf(t, "IPV6\t%s\n", o.IPV6)
	fmt.Fprintf(t, "IPV6 NETWORK SIZE\t%d\n", o.IPV6NetworkSize)
	fmt.Fprintf(t, "LABEL\t%s\n", o.Label)
	fmt.Fprintf(t, "PRIVATE IP\t%s\n", o.PrivateIP)
	fmt.Fprintf(t, "KVM URL\t%s\n", o.KVMURL)
	fmt.Fprintf(t, "AUTO BACKUPS\t%s\n", strconv.FormatBool(o.AutoBackups))

	t.Flush()

	return
}
