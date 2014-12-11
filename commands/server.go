package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"strconv"
	"text/tabwriter"
)

type server struct{Command}

func NewServer() *server {
	return &server{
		Command {
			Desc: "Get server information.",
			NeedsKey: true,
			ArgsDesc: "server_id",
			OptionsDesc: "",
		},
	}
}

func (_ *server) Exec(c Client, args []string, key string) (err error) {
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

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintf(w, "ID\t%d\n", s.Id)
	fmt.Fprintf(w, "OS\t%s\n", s.OS)
	fmt.Fprintf(w, "RAM\t%s\n", s.Disk)
	fmt.Fprintf(w, "DISK\t%s\n", s.IPV4)
	fmt.Fprintf(w, "CPUS\t%d\n", s.CPUs)
	fmt.Fprintf(w, "LOCATION\t%s\n", s.Location)
	fmt.Fprintf(w, "REGION ID\t%d\n", s.RegionId)
	fmt.Fprintf(w, "DEFAULT PASSWORD\t%s\n", s.DefaultPassword)
	fmt.Fprintf(w, "DATE CREATED\t%s\n", s.DateCreated)
	fmt.Fprintf(w, "PENDING CHARGES\t%.2f\n", s.PendingCharges)
	fmt.Fprintf(w, "STATUS\t%s\n", s.Status)
	fmt.Fprintf(w, "PRICE/MONTH\t%.2f\n", s.PricePerMonth)
	fmt.Fprintf(w, "CURRENT BANDWIDTH GB\t%.2f\n", s.CurrentBandwidthGB)
	fmt.Fprintf(w, "ALLOWED BANDWIDTH GB\t%.2f\n", s.AllowedBandwidthGB)
	fmt.Fprintf(w, "IPV4 NETMASK\t%s\n", s.IPV4Netmask)
	fmt.Fprintf(w, "IPV4 GATEWAY\t%s\n", s.IPV4Gateway)
	fmt.Fprintf(w, "POWER STATUS\t%s\n", s.PowerStatus)
	fmt.Fprintf(w, "PLAN ID\t%d\n", s.PlanId)
	fmt.Fprintf(w, "IPV6 NETWORK\t%s\n", s.IPV6Network)
	fmt.Fprintf(w, "IPV6\t%s\n", s.IPV6)
	fmt.Fprintf(w, "IPV6 NETWORK SIZE\t%d\n", s.IPV6NetworkSize)
	fmt.Fprintf(w, "LABEL\t%s\n", s.Label)
	fmt.Fprintf(w, "PRIVATE IP\t%s\n", s.PrivateIP)
	fmt.Fprintf(w, "KVM URL\t%s\n", s.KVMURL)
	fmt.Fprintf(w, "AUTO BACKUPS\t%s\n", strconv.FormatBool(s.AutoBackups))

	w.Flush()

	return
}
