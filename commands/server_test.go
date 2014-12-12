package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleServer() {
	c := NewTestClient(200, servers)
	NewServer().Fexec(os.Stdout, c, []string{"1571183"}, "API_KEY")
	// Output:
	// ID			1571183
	// OS			CentOS 7 x64
	// RAM			Virtual 20 GB
	// DISK			108.61.177.174
	// CPUS			1
	// LOCATION		France
	// REGION ID		24
	// DEFAULT PASSWORD	pyetbuch!0
	// DATE CREATED		2014-12-09 21:18:30 +0000
	// PENDING CHARGES		0.02
	// STATUS			active
	// PRICE/MONTH		7.00
	// CURRENT BANDWIDTH GB	0.00
	// ALLOWED BANDWIDTH GB	2000.00
	// IPV4 NETMASK		255.255.254.0
	// IPV4 GATEWAY		108.61.176.1
	// POWER STATUS		running
	// PLAN ID			30
	// IPV6 NETWORK		::
	// IPV6			test
	// IPV6 NETWORK SIZE	0
	// LABEL			test
	// PRIVATE IP		test
	// KVM URL			test
	// AUTO BACKUPS		false
}
