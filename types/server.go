package types

import (
	"encoding/json"
	"strconv"
)

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
	PricePerMonth      float64 `json:"cost_per_month,string"`
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
