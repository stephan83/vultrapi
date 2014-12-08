package types

type Server struct {
	Id                  int     `json:"SUBID, string"`
	OS                  string  `json:"os"`
	RAM                 string  `json:"ram"`
	Disk                string  `json:"disk"`
	PublicIP            string  `json:"main_ip"`
	CPUs                int     `json:"vcpu_count, string"`
	Location            string  `json:"location"`
	RegionId            string  `json:"DCID"`
	DefaultPassword     string  `json:"default_password"`
	DateCreated         string  `json:"date_created"`
	PendingCharges      string  `json:"pending_charges"`
	Status              string  `json:"status"`
	CostPerMonth        string  `json:"cost_per_month"`
	CurrentBandwidthGB  float64 `json:"current_bandwidth_gb"`
	AllowedBandwidthGB  string  `json:"allowed_bandwidth_gb"`
	PublicNetmask       string  `json:"netmask_v4"`
	PublicGateway       string  `json:"gateway_v4"`
	PowerStatus         string  `json:"running"`
	PlanId              int     `json:"VPSPLANID, string"`
	V6PublicNetwork     string  `json:"v6_network"`
	V6PublicIP          string  `json:"v6_main_ip"`
	V6PublicNetworkSize string  `json:"v6_network_size"`
	PrivateIP           string  `json:"internal_ip"`
	KVMURL              string  `json:"kvm_url"`
	AutoBackups         bool    `json:"auto_backups, string"`
}
