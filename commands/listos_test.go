package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleListOS() {
	c := NewTestClient(200, _os)
	NewListOS().Fexec(os.Stdout, c, []string{}, "")
	// Output:
	// ID	NAME			FAMILY		ARCH	WINDOWS
	// 180	Backup			backup		x64	false
	// 163	CentOS 5 i386		centos		i386	false
	// 147	CentOS 6 i386		centos		i386	false
	// 162	CentOS 5 x64		centos		x64	false
	// 127	CentOS 6 x64		centos		x64	false
	// 167	CentOS 7 x64		centos		x64	false
	// 179	CoreOS Stable		coreos		x64	false
	// 152	Debian 7 i386 (wheezy)	debian		i386	false
	// 139	Debian 7 x64 (wheezy)	debian		x64	false
	// 140	FreeBSD 10 x64		freebsd		x64	false
	// 159	Custom			iso		x64	false
	// 164	Snapshot		snapshot	x64	false
	// 148	Ubuntu 12.04 i386	ubuntu		i386	false
	// 161	Ubuntu 14.04 i386	ubuntu		i386	false
	// 182	Ubuntu 14.10 i386	ubuntu		i386	false
	// 128	Ubuntu 12.04 x64	ubuntu		x64	false
	// 160	Ubuntu 14.04 x64	ubuntu		x64	false
	// 181	Ubuntu 14.10 x64	ubuntu		x64	false
	// 124	Windows 2012 R2 x64	windows		x64	true
}

var _os = []byte(`{
	"127": {
		"OSID": 127,
		"name": "CentOS 6 x64",
		"arch": "x64",
		"family": "centos",
		"windows": false
	},
	"147": {
		"OSID": 147,
		"name": "CentOS 6 i386",
		"arch": "i386",
		"family": "centos",
		"windows": false
	},
	"162": {
		"OSID": 162,
		"name": "CentOS 5 x64",
		"arch": "x64",
		"family": "centos",
		"windows": false
	},
	"163": {
		"OSID": 163,
		"name": "CentOS 5 i386",
		"arch": "i386",
		"family": "centos",
		"windows": false
	},
	"167": {
		"OSID": 167,
		"name": "CentOS 7 x64",
		"arch": "x64",
		"family": "centos",
		"windows": false
	},
	"160": {
		"OSID": 160,
		"name": "Ubuntu 14.04 x64",
		"arch": "x64",
		"family": "ubuntu",
		"windows": false
	},
	"161": {
		"OSID": 161,
		"name": "Ubuntu 14.04 i386",
		"arch": "i386",
		"family": "ubuntu",
		"windows": false
	},
	"128": {
		"OSID": 128,
		"name": "Ubuntu 12.04 x64",
		"arch": "x64",
		"family": "ubuntu",
		"windows": false
	},
	"148": {
		"OSID": 148,
		"name": "Ubuntu 12.04 i386",
		"arch": "i386",
		"family": "ubuntu",
		"windows": false
	},
	"181": {
		"OSID": 181,
		"name": "Ubuntu 14.10 x64",
		"arch": "x64",
		"family": "ubuntu",
		"windows": false
	},
	"182": {
		"OSID": 182,
		"name": "Ubuntu 14.10 i386",
		"arch": "i386",
		"family": "ubuntu",
		"windows": false
	},
	"139": {
		"OSID": 139,
		"name": "Debian 7 x64 (wheezy)",
		"arch": "x64",
		"family": "debian",
		"windows": false
	},
	"152": {
		"OSID": 152,
		"name": "Debian 7 i386 (wheezy)",
		"arch": "i386",
		"family": "debian",
		"windows": false
	},
	"140": {
		"OSID": 140,
		"name": "FreeBSD 10 x64",
		"arch": "x64",
		"family": "freebsd",
		"windows": false
	},
	"179": {
		"OSID": 179,
		"name": "CoreOS Stable",
		"arch": "x64",
		"family": "coreos",
		"windows": false
	},
	"124": {
		"OSID": 124,
		"name": "Windows 2012 R2 x64",
		"arch": "x64",
		"family": "windows",
		"windows": true
	},
	"159": {
		"OSID": 159,
		"name": "Custom",
		"arch": "x64",
		"family": "iso",
		"windows": false
	},
	"164": {
		"OSID": 164,
		"name": "Snapshot",
		"arch": "x64",
		"family": "snapshot",
		"windows": false
	},
	"180": {
		"OSID": 180,
		"name": "Backup",
		"arch": "x64",
		"family": "backup",
		"windows": false
	}
}`)
