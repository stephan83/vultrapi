package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/errors"
	"github.com/stephan83/vultrapi/requests"
	"io"
	"strconv"
	"text/tabwriter"
	"flag"
)

type server struct{
	BasicCommandWithOptions
	Labels bool
	Fields StringSlice
}

func NewServer() Command {
	f := flag.NewFlagSet("server", flag.ContinueOnError)

	o := server{
		BasicCommandWithOptions: BasicCommandWithOptions{
			BasicCommand{
				Desc:     "Get server information.",
				NeedsKey: true,
				ArgsDesc: "server_id",
			},
			f,
		},
	}

	f.BoolVar(&o.Labels, "labels", true, "display labels")
	f.Var(&o.Fields, "field", "only display the specified field (can be set multiple times)")

	o.Initialize()

	o.OptionsDesc += "\nAvailable fields: id, os, ram, disk, ipv4, cpus,"+
	                 "location, region_id, default_password,"+
	                 "date_created, pending_charges, status,"+
	                 "price_per_month, currend_bandwidth_gb,"+
	                 "allowed_bandwidth_gb, ipv4_netmask, ipv4_gateway,"+
	                 "power_status, plan_id, ipv6_network, ipv6,"+
	                 "ipv6_network_size, label, private_ip, kvm_url,"+
	                 "auto_backups"

	return &o
}

func (o *server) Fexec(w io.Writer, c Client, args []string, key string) (err error) {
	if len(args) < 1 {
		err = ErrUsage{}
		return
	}

	o.FlagSet.SetOutput(w)

	err = o.FlagSet.Parse(args[1:])
	if err != nil {
		return ErrUsage{}
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

	e, ok := r[id]
	if !ok {
		err = ErrNotFound{}
		return
	}

	fields := map[string]bool{}

	for _, v := range o.Fields {
		fields[v] = true
	}

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	if fields["id"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "ID\t%d\n", e.Id)
		} else {
			fmt.Fprintf(t, "%d\n", e.Id)
		}
	}
	if fields["os"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "OS\t%s\n", e.OS)
		} else {
			fmt.Fprintf(t, "%s\n", e.OS)
		}
	}
	if fields["ram"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "RAM\t%d\n", e.RAM)
		} else {
			fmt.Fprintf(t, "%d\n", e.RAM)
		}
	}
	if fields["disk"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "DISK\t%s\n", e.Disk)
		} else {
			fmt.Fprintf(t, "%s\n", e.Disk)
		}
	}
	if fields["ipv4"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "IPV4\t%s\n", e.IPV4)
		} else {
			fmt.Fprintf(t, "%s\n", e.IPV4)
		}
	}
	if fields["cpus"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "CPUS\t%d\n", e.CPUs)
		} else {
			fmt.Fprintf(t, "%d\n", e.CPUs)
		}
	}
	if fields["location"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "LOCATION\t%s\n", e.Location)
		} else {
			fmt.Fprintf(t, "%s\n", e.Location)
		}
	}
	if fields["region_id"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "REGION ID\t%d\n", e.RegionId)
		} else {
			fmt.Fprintf(t, "%d\n", e.RegionId)
		}
	}
	if fields["default_password"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "DEFAULT PASSWORD\t%s\n", e.DefaultPassword)
		} else {
			fmt.Fprintf(t, "%s\n", e.DefaultPassword)
		}
	}
	if fields["date_created"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "DATE CREATED\t%s\n", e.DateCreated)
		} else {
			fmt.Fprintf(t, "%s\n", e.DateCreated)
		}
	}
	if fields["pending_charges"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "PENDING CHARGES\t%.2f\n", e.PendingCharges)
		} else {
			fmt.Fprintf(t, "%.2f\n", e.PendingCharges)
		}
	}
	if fields["status"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "STATUS\t%s\n", e.Status)
		} else {
			fmt.Fprintf(t, "%s\n", e.Status)
		}
	}
	if fields["price_per_month"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "PRICE/MONTH\t%.2f\n", e.PricePerMonth)
		} else {
			fmt.Fprintf(t, "%.2f\n", e.PricePerMonth)
		}
	}
	if fields["currend_bandwidth_gb"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "CURRENT BANDWIDTH GB\t%.2f\n", e.CurrentBandwidthGB)
		} else {
			fmt.Fprintf(t, "%.2f\n", e.CurrentBandwidthGB)
		}
	}
	if fields["allowed_bandwidth_gb"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "ALLOWED BANDWIDTH GB\t%.2f\n", e.AllowedBandwidthGB)
		} else {
			fmt.Fprintf(t, "%.2f\n", e.AllowedBandwidthGB)
		}
	}
	if fields["ipv4_netmask"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "IPV4 NETMASK\t%s\n", e.IPV4Netmask)
		} else {
			fmt.Fprintf(t, "%s\n", e.IPV4Netmask)
		}
	}
	if fields["ipv4_gateway"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "IPV4 GATEWAY\t%s\n", e.IPV4Gateway)
		} else {
			fmt.Fprintf(t, "%s\n", e.IPV4Gateway)
		}
	}
	if fields["power_status"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "POWER STATUS\t%s\n", e.PowerStatus)
		} else {
			fmt.Fprintf(t, "%s\n", e.PowerStatus)
		}
	}
	if fields["plan_id"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "PLAN ID\t%d\n", e.PlanId)
		} else {
			fmt.Fprintf(t, "%d\n", e.PlanId)
		}
	}
	if fields["ipv6_network"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "IPV6 NETWORK\t%s\n", e.IPV6Network)
		} else {
			fmt.Fprintf(t, "%s\n", e.IPV6Network)
		}
	}
	if fields["ipv6"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "IPV6\t%s\n", e.IPV6)
		} else {
			fmt.Fprintf(t, "%s\n", e.IPV6)
		}
	}
	if fields["ipv6_network_size"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "IPV6 NETWORK SIZE\t%d\n", e.IPV6NetworkSize)
		} else {
			fmt.Fprintf(t, "%d\n", e.IPV6NetworkSize)
		}
	}
	if fields["label"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "LABEL\t%s\n", e.Label)
		} else {
			fmt.Fprintf(t, "%s\n", e.Label)
		}
	}
	if fields["private_ip"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "PRIVATE IP\t%s\n", e.PrivateIP)
		} else {
			fmt.Fprintf(t, "%s\n", e.PrivateIP)
		}
	}
	if fields["kvm_url"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "KVM URL\t%s\n", e.KVMURL)
		} else {
			fmt.Fprintf(t, "%s\n", e.KVMURL)
		}
	}
	if fields["auto_backups"] || len(o.Fields) == 0 {
		if o.Labels {
			fmt.Fprintf(t, "AUTO BACKUPS\t%s\n", strconv.FormatBool(e.AutoBackups))
		} else {
			fmt.Fprintf(t, "%s\n", strconv.FormatBool(e.AutoBackups))
		}
	}

	t.Flush()

	return
}
