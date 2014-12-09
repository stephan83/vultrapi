package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleListServers() {
	c := NewTestClient(200, servers)
	NewListServers().Exec(c, []string{}, "")
	// Output:
	// LOCATION        | LABEL              | IPV4            | STATUS     | ID
	// ------------------------------------------------------------------------------
	// France          |                    | 108.61.177.174  | active     | 1571183
	// France          |                    | 0               | pending    | 1571200
	// Japan           |                    | 0               | pending    | 1571201
}

func ExampleListServersWithRegion() {
	c := NewTestClient(200, servers)
	NewListServers().Exec(c, []string{"-region", "24"}, "")
	// Output:
	// LOCATION        | LABEL              | IPV4            | STATUS     | ID
	// ------------------------------------------------------------------------------
	// France          |                    | 108.61.177.174  | active     | 1571183
	// France          |                    | 0               | pending    | 1571200
}

func ExampleListServersWithPlan() {
	c := NewTestClient(200, servers)
	NewListServers().Exec(c, []string{"-plan", "3"}, "")
	// Output:
	// LOCATION        | LABEL              | IPV4            | STATUS     | ID
	// ------------------------------------------------------------------------------
	// France          |                    | 0               | pending    | 1571200
	// Japan           |                    | 0               | pending    | 1571201
}

func ExampleListServersWithRegionAndPlan() {
	c := NewTestClient(200, servers)
	NewListServers().Exec(c, []string{"-plan", "3", "-region", "24"}, "")
	// Output:
	// LOCATION        | LABEL              | IPV4            | STATUS     | ID
	// ------------------------------------------------------------------------------
	// France          |                    | 0               | pending    | 1571200
}

var servers = []byte(`{
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
	},
	"1571200": {
		"SUBID": "1571200",
		"os": "Ubuntu 14.10 x64",
		"ram": "2048 MB",
		"disk": "Virtual 40 GB",
		"main_ip": "0",
		"vcpu_count": "2",
		"location": "France",
		"DCID": "24",
		"default_password": "rmovcexi",
		"date_created": "2014-12-09 16:45:36",
		"pending_charges": 0,
		"status": "pending",
		"cost_per_month": "15.00",
		"current_bandwidth_gb": 0,
		"allowed_bandwidth_gb": "3000",
		"netmask_v4": "0.0.0.0",
		"gateway_v4": "0.0.0.0",
		"power_status": "running",
		"VPSPLANID": "3",
		"v6_network": "::",
		"v6_main_ip": "",
		"v6_network_size": "0",
		"label": "",
		"internal_ip": "",
		"kvm_url": "",
		"auto_backups": "no"
	},
	"1571201": {
		"SUBID": "1571201",
		"os": "Ubuntu 14.10 x64",
		"ram": "2048 MB",
		"disk": "Virtual 40 GB",
		"main_ip": "0",
		"vcpu_count": "2",
		"location": "Japan",
		"DCID": "231",
		"default_password": "rmovcexi",
		"date_created": "2014-12-09 16:45:36",
		"pending_charges": 0,
		"status": "pending",
		"cost_per_month": "15.00",
		"current_bandwidth_gb": 0,
		"allowed_bandwidth_gb": "3000",
		"netmask_v4": "0.0.0.0",
		"gateway_v4": "0.0.0.0",
		"power_status": "running",
		"VPSPLANID": "3",
		"v6_network": "::",
		"v6_main_ip": "",
		"v6_network_size": "0",
		"label": "",
		"internal_ip": "",
		"kvm_url": "",
		"auto_backups": "no"
	}
}`)
