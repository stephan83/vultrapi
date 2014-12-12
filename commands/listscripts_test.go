package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleListScripts() {
	c := NewTestClient(200, scripts)
	NewListScripts().Fexec(os.Stdout, c, []string{}, "")
	// Output:
	// ID	NAME		DATE CREATED
	// 1689	test boot	2014-12-12 00:38:22 +0000
	// 1688	test pxe	2014-12-12 00:38:08 +0000
}

var scripts = []byte(`{
	"1688": {
		"SCRIPTID": "1688",
		"date_created": "2014-12-11 19:38:08",
		"date_modified": "2014-12-11 19:38:08",
		"name": "test pxe",
		"script": "#!ipxe\nset base-url http://beta.release.core-os.net/amd64-usr/current\nkernel ${base-url}/coreos_production_pxe.vmlinuz sshkey=\"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDIhWtVyJTCIsXQqUWC6FNtxr0I2N9V8fMgQApxyW8jPn4lc6DapwL3LGvdYWsQDMPH/WvnEd4GKVuA2WnWV6/RjGFaEZmXgEVhCtu3PH9NYJGZ0dSaBh7GjwloEhR9kYs+aAOkGKEYUzoIgYv3fmQCqcXibJE+k4qL2ag5D51sZsqiyvb+rMgTzIouuCkI6d9jhvmEk8E7FV+NAz7mi9bbPkOJJqfDDQUvy+ph5I1HUD0FSlYTfhQ557Wrm5JcZnac3rtixQBt01pCic2RjiVXDovN5K2ihxLExDLbyt7NIrqzkELSYOP9tpHijYgGyN4aBdi0KKuo5X6noqvH2KUR test\"\ninitrd ${base-url}/coreos_production_pxe_image.cpio.gz\nboot\n"
	},
	"1689": {
		"SCRIPTID": "1689",
		"date_created": "2014-12-11 19:38:22",
		"date_modified": "2014-12-11 19:38:22",
		"name": "test boot",
		"script": "#!/bin/sh\necho \"test\""
	}
}`)
