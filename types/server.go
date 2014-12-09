package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
{
  "1571183": {
    "SUBID": "1571183",
    "os": "CentOS 7 x64",
    "ram": "1024 MB",
    "disk": "Virtual 20 GB",
    "main_ip": "108.61.177.174",
    "vcpu_count": "1",
    "location": "France",
    "DCID": "24",
    "default_password": "pyetbuch!0",
    "date_created": "2014-12-09 16:18:30",
    "pending_charges": 0.02,
    "status": "active",
    "cost_per_month": "7.00",
    "current_bandwidth_gb": 0,
    "allowed_bandwidth_gb": "2000",
    "netmask_v4": "255.255.254.0",
    "gateway_v4": "108.61.176.1",
    "power_status": "running",
    "VPSPLANID": "30",
    "v6_network": "::",
    "v6_main_ip": "",
    "v6_network_size": "0",
    "label": "",
    "internal_ip": "",
    "kvm_url": "https://my.vultr.com/subs/vps/novnc/api.php?data=KFEEG6SFGRJDGQSIGRRXE6LTKZVG24JVM53WGMZVPFETCS2RFNKWGRKGKQ2VM5ZRNU3UEMRXJJ3XMK32KBHHIY3VKJGE223YKZ2GG4ZQLBLEUN2TLBWHGNKRIJKVGVSTLBHUQ4JWIFXUQSZWHBUUG4DXLFFWK33SINEXC5KBNFYGKMKZIFVWSWKSK5SEO5SFKBFWYY3VJNDTQ4CCKRIW6R2ZNVUGOYRSINBWCRSJPBGFETSCNFAUKRDHIZJUGZSYMR4HQRLCKF5FMSS2OM6Q",
    "auto_backups": "no"
  }
}
*/
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
	CostPerMonth        string  `json:"cost_per_month"`
	CurrentBandwidthGB  float64 `json:"current_bandwidth_gb"`
	AllowedBandwidthGB  string  `json:"allowed_bandwidth_gb"`
	IPV4Netmask         string  `json:"netmask_v4"`
	IPV4Gateway         string  `json:"gateway_v4"`
	PowerStatus         string  `json:"running"`
	PlanId              int     `json:"VPSPLANID,string"`
	IPV6Network         string  `json:"v6_network"`
	IVV6IP              string  `json:"v6_main_ip"`
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

func (r Server) String() string {
	return fmt.Sprintf(serverFormat, r.Location, r.Label, r.IPV4,
		r.Status, strconv.Itoa(r.Id))
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
