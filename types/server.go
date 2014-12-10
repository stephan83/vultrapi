package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const serverFormat = "%-15s | %-18s | %-15s | %-10s | %s"

type Server struct {
	Id                 int     `json:"SUBID,string"`
	OS                 string  `json:"os"`
	RAM                RAM     `json:"ram"`
	Disk               string  `json:"disk"`
	IPV4               string  `json:"main_ip"`
	CPUs               int     `json:"vcpu_count,string"`
	Location           string  `json:"location"`
	RegionId           int     `json:"DCID,string"`
	DefaultPassword    string  `json:"default_password"`
	DateCreated        Date    `json:"date_created"`
	PendingCharges     float64 `json:"pending_charges"`
	Status             string  `json:"status"`
	PricePerMonth      string  `json:"cost_per_month"`
	CurrentBandwidthGB float64 `json:"current_bandwidth_gb"`
	AllowedBandwidthGB float64 `json:"allowed_bandwidth_gb,string"`
	IPV4Netmask        string  `json:"netmask_v4"`
	IPV4Gateway        string  `json:"gateway_v4"`
	PowerStatus        string  `json:"power_status"`
	PlanId             int     `json:"VPSPLANID,string"`
	IPV6Network        string  `json:"v6_network"`
	IPV6               string  `json:"v6_main_ip"`
	IPV6NetworkSize    int     `json:"v6_network_size,string"`
	Label              string  `json:"label"`
	PrivateIP          string  `json:"internal_ip"`
	KVMURL             string  `json:"kvm_url"`
	AutoBackups        bool    `json:"auto_backups,string"`
}

func (o Server) String() string {
	return fmt.Sprintf(serverFormat, o.Location, o.Label, o.IPV4,
		o.Status, strconv.Itoa(o.Id))
}

func (o Server) Details() string {
	return strings.Join([]string{
		fmt.Sprintf("%21s: %d", "ID", o.Id),
		fmt.Sprintf("%21s: %s", "OS", o.OS),
		fmt.Sprintf("%21s: %s", "RAM", strconv.Itoa(int(o.RAM))),
		fmt.Sprintf("%21s: %s", "DISK", o.Disk),
		fmt.Sprintf("%21s: %s", "IPV4", o.IPV4),
		fmt.Sprintf("%21s: %d", "CPUS", o.CPUs),
		fmt.Sprintf("%21s: %s", "LOCATION", o.Location),
		fmt.Sprintf("%21s: %d", "REGION ID", o.RegionId),
		fmt.Sprintf("%21s: %s", "DEFAULT PASSWORD", o.DefaultPassword),
		fmt.Sprintf("%21s: %s", "DATE CREATED", o.DateCreated),
		fmt.Sprintf("%21s: %.2f", "PENDING CHARGES", o.PendingCharges),
		fmt.Sprintf("%21s: %s", "STATUS", o.Status),
		fmt.Sprintf("%21s: %s", "PRICE/MONTH", o.PricePerMonth),
		fmt.Sprintf("%21s: %.2f", "CURRENT BANDWIDTH GB", o.CurrentBandwidthGB),
		fmt.Sprintf("%21s: %s", "ALLOWED BANDWIDTH GB", fmt.Sprintf("%.2f", o.AllowedBandwidthGB)),
		fmt.Sprintf("%21s: %s", "IPV4 NETMASK", o.IPV4Netmask),
		fmt.Sprintf("%21s: %s", "IPV4 GATEWAY", o.IPV4Gateway),
		fmt.Sprintf("%21s: %s", "POWER STATUS", o.PowerStatus),
		fmt.Sprintf("%21s: %d", "PLAN ID", o.PlanId),
		fmt.Sprintf("%21s: %s", "IPV6 NETWORK", o.IPV6Network),
		fmt.Sprintf("%21s: %s", "IPV6", o.IPV6),
		fmt.Sprintf("%21s: %d", "IPV6 NETWORK SIZE", o.IPV6NetworkSize),
		fmt.Sprintf("%21s: %s", "LABEL", o.Label),
		fmt.Sprintf("%21s: %s", "PRIVATE IP", o.PrivateIP),
		fmt.Sprintf("%21s: %s", "KVM URL", o.KVMURL),
		fmt.Sprintf("%21s: %s", "AUTO BACKUPS", strconv.FormatBool(o.AutoBackups)),
	}, "\n")
}

type ServerMap map[int]Server

func (o ServerMap) MarshalJSON() ([]byte, error) {
	m := map[string]Server{}

	for k, v := range o {
		m[strconv.Itoa(k)] = v
	}

	return json.Marshal(m)
}

func (o *ServerMap) UnmarshalJSON(d []byte) error {
	*o = ServerMap{}

	if string(d) == "[]" {
		return nil
	}

	m := map[string]Server{}

	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	for k, v := range m {
		i, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		(*o)[i] = v
	}

	return nil
}

func (o ServerMap) Array() ServerArray {
	a := []Server{}

	for _, r := range o {
		a = append(a, r)
	}

	return a
}

func (o ServerMap) String() string {
	l := []string{}

	l = append(l, fmt.Sprintf(serverFormat, "LOCATION", "LABEL", "IPV4",
		"STATUS", "ID"))
	l = append(l, strings.Repeat("-", 78))

	a := o.Array()
	sort.Sort(a)

	for _, r := range a {
		l = append(l, r.String())
	}

	return strings.Join(l, "\n")
}

type ServerArray []Server

func (a ServerArray) Len() int {
	return len(a)
}

func (a ServerArray) Less(i, j int) bool {
	switch {
	case a[i].Location < a[j].Location:
		return true
	case a[i].Location > a[j].Location:
		return false
	default:
		switch {
		case a[i].Status < a[j].Status:
			return true
		case a[i].Status > a[j].Status:
			return false
		default:
			switch {
			case a[i].Label < a[j].Label:
				return true
			case a[i].Label > a[j].Label:
				return false
			default:
				return a[i].IPV4 < a[j].IPV4
			}
		}
	}
}

func (ss ServerArray) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
