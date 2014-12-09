package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Server struct {
	Id                  int     `json:"SUBID,string"`
	OS                  string  `json:"os"`
	RAM                 string  `json:"ram"`
	Disk                string  `json:"disk"`
	IPV4	            string  `json:"main_ip"`
	CPUs                int     `json:"vcpu_count,string"`
	Location            string  `json:"location"`
	RegionId            int     `json:"DCID,string"`
	DefaultPassword     string  `json:"default_password"`
	DateCreated         string  `json:"date_created"`
	PendingCharges      float64 `json:"pending_charges"`
	Status              string  `json:"status"`
	PricePerMonth       string  `json:"cost_per_month"`
	CurrentBandwidthGB  float64 `json:"current_bandwidth_gb"`
	AllowedBandwidthGB  string  `json:"allowed_bandwidth_gb"`
	IPV4Netmask         string  `json:"netmask_v4"`
	IPV4Gateway         string  `json:"gateway_v4"`
	PowerStatus         string  `json:"power_status"`
	PlanId              int     `json:"VPSPLANID,string"`
	IPV6Network         string  `json:"v6_network"`
	IPV6                string  `json:"v6_main_ip"`
	IPV6NetworkSize     int  `json:"v6_network_size,string"`
	Label               string  `json:"label"`
	PrivateIP           string  `json:"internal_ip"`
	KVMURL              string  `json:"kvm_url"`
	AutoBackups         bool    `json:"auto_backups,string"`
}

const (
	// Location | Label | IP | Status | Id
	serverFormat = "%-15s | %-18s | %-15s | %-10s | %s"
)

func (s Server) String() string {
	return fmt.Sprintf(serverFormat, s.Location, s.Label, s.IPV4,
		s.Status, strconv.Itoa(s.Id))
}

func (s Server) Details() string {
	return strings.Join([]string{
		fmt.Sprintf("%21s: %d", "ID", s.Id),
		fmt.Sprintf("%21s: %s", "OS", s.OS),
		fmt.Sprintf("%21s: %s", "RAM", s.RAM),
		fmt.Sprintf("%21s: %s", "DISK", s.Disk),
		fmt.Sprintf("%21s: %s", "IPV4", s.IPV4),
		fmt.Sprintf("%21s: %d", "CPUS", s.CPUs),
		fmt.Sprintf("%21s: %s", "LOCATION", s.Location),
		fmt.Sprintf("%21s: %d", "REGION ID", s.RegionId),
		fmt.Sprintf("%21s: %s", "DEFAULT PASSWORD", s.DefaultPassword),
		fmt.Sprintf("%21s: %s", "DATE CREATED", s.DateCreated),
		fmt.Sprintf("%21s: %.2f", "PENDING CHARGES", s.PendingCharges),
		fmt.Sprintf("%21s: %s", "STATUS", s.Status),
		fmt.Sprintf("%21s: %s", "PRICE/MONTH", s.PricePerMonth),
		fmt.Sprintf("%21s: %.2f", "CURRENT BANDWIDTH GB", s.CurrentBandwidthGB),
		fmt.Sprintf("%21s: %s", "ALLOWED BANDWIDTH GB", s.AllowedBandwidthGB),
		fmt.Sprintf("%21s: %s", "IPV4 NETMASK", s.IPV4Netmask),
		fmt.Sprintf("%21s: %s", "IPV4 GATEWAY", s.IPV4Gateway),
		fmt.Sprintf("%21s: %s", "POWER STATUS", s.PowerStatus),
		fmt.Sprintf("%21s: %d", "PLAN ID", s.PlanId),
		fmt.Sprintf("%21s: %s", "IPV6 NETWORK", s.IPV6Network),
		fmt.Sprintf("%21s: %s", "IPV6", s.IPV6),
		fmt.Sprintf("%21s: %d", "IPV6 NETWORK SIZE", s.IPV6NetworkSize),
		fmt.Sprintf("%21s: %s", "LABEL", s.Label),
		fmt.Sprintf("%21s: %s", "PRIVATE IP", s.PrivateIP),
		fmt.Sprintf("%21s: %s", "KVM URL", s.KVMURL),
		fmt.Sprintf("%21s: %s", "AUTO BACKUPS", strconv.FormatBool(s.AutoBackups)),
	}, "\n")
}

type ServerDict map[string]Server

func (sd ServerDict) Slice() ServerSlice {
	servers := []Server{}

	for _, r := range sd {
		servers = append(servers, r)
	}

	return servers
}

func (sd ServerDict) String() string {
	lines := []string{}

	lines = append(lines, fmt.Sprintf(serverFormat, "LOCATION",
		"LABEL", "IPV4", "STATUS", "ID"))

	lines = append(lines, strings.Repeat("-", 78))

	ss := sd.Slice()
	sort.Sort(ss)

	for _, r := range ss {
		lines = append(lines, r.String())
	}

	return strings.Join(lines, "\n")
}

type ServerSlice []Server

func (ss ServerSlice) Len() int {
	return len(ss)
}

func (ss ServerSlice) Less(i, j int) bool {
	switch {
	case ss[i].Location < ss[j].Location:
		return true
	case ss[i].Location > ss[j].Location:
		return false
	default:
		switch {
		case ss[i].Status < ss[j].Status:
			return true
		case ss[i].Status > ss[j].Status:
			return false
		default:
			switch {
			case ss[i].Label < ss[j].Label:
				return true
			case ss[i].Label > ss[j].Label:
				return false
			default:
				return ss[i].IPV4 < ss[j].IPV4
			}
		}
	}
}

func (ss ServerSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
