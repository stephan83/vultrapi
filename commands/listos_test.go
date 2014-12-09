package commands

import . "github.com/stephan83/vultrapi/clients"

func ExampleListOS() {
	c := NewTestClient(200, _os)
	NewListOS().Exec(c, []string{"listos"}, "")
	// Output:
	// FAMILY           | ARCH | NAME                                          | ID
	// ------------------------------------------------------------------------------
	// backup           | x64  | Backup                                        | 180
	// centos           | i386 | CentOS 5 i386                                 | 163
	// centos           | i386 | CentOS 6 i386                                 | 147
	// centos           | x64  | CentOS 5 x64                                  | 162
	// centos           | x64  | CentOS 6 x64                                  | 127
	// centos           | x64  | CentOS 7 x64                                  | 167
	// coreos           | x64  | CoreOS Stable                                 | 179
	// debian           | i386 | Debian 7 i386 (wheezy)                        | 152
	// debian           | x64  | Debian 7 x64 (wheezy)                         | 139
	// freebsd          | x64  | FreeBSD 10 x64                                | 140
	// iso              | x64  | Custom                                        | 159
	// snapshot         | x64  | Snapshot                                      | 164
	// ubuntu           | i386 | Ubuntu 12.04 i386                             | 148
	// ubuntu           | i386 | Ubuntu 14.04 i386                             | 161
	// ubuntu           | i386 | Ubuntu 14.10 i386                             | 182
	// ubuntu           | x64  | Ubuntu 12.04 x64                              | 128
	// ubuntu           | x64  | Ubuntu 14.04 x64                              | 160
	// ubuntu           | x64  | Ubuntu 14.10 x64                              | 181
	// windows          | x64  | Windows 2012 R2 x64                           | 124
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
