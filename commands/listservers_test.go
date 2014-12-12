package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleListServers() {
	c := NewTestClient(200, servers)
	NewListServers().Fexec(os.Stdout, c, []string{}, "")
	// Output:
	// ID	LOCATION	OS			IPV4		STATUS	LABEL
	// 1571183	France		CentOS 7 x64		108.61.177.174	active	test
	// 1571200	France		Ubuntu 14.10 x64	0		pending	-
	// 1571201	Japan		Ubuntu 14.10 x64	0		pending	-
}

func ExampleListServersWithRegion() {
	c := NewTestClient(200, servers)
	NewListServers().Fexec(os.Stdout, c, []string{"-region", "24"}, "")
	// Output:
	// ID	LOCATION	OS			IPV4		STATUS	LABEL
	// 1571183	France		CentOS 7 x64		108.61.177.174	active	test
	// 1571200	France		Ubuntu 14.10 x64	0		pending	-
}

func ExampleListServersWithPlan() {
	c := NewTestClient(200, servers)
	NewListServers().Fexec(os.Stdout, c, []string{"-plan", "3"}, "")
	// Output:
	// ID	LOCATION	OS			IPV4	STATUS	LABEL
	// 1571200	France		Ubuntu 14.10 x64	0	pending	-
	// 1571201	Japan		Ubuntu 14.10 x64	0	pending	-
}

func ExampleListServersWithRegionAndPlan() {
	c := NewTestClient(200, servers)
	NewListServers().Fexec(os.Stdout, c, []string{"-plan", "3", "-region", "24"}, "")
	// Output:
	// ID	LOCATION	OS			IPV4	STATUS	LABEL
	// 1571200	France		Ubuntu 14.10 x64	0	pending	-
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
		"v6_main_ip": "test",
		"v6_network_size": "0",
		"label": "test",
		"internal_ip": "test",
		"kvm_url": "test",
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
		"label": "-",
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
		"label": "-",
		"internal_ip": "",
		"kvm_url": "",
		"auto_backups": "no"
	}
}`)
