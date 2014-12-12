# VULTRAPI [![Build Status](https://travis-ci.org/stephan83/vultrapi.svg?branch=master)](https://travis-ci.org/stephan83/vultrapi)

Client/library for [Vultr API](https://vultr.com) written in go.

Covers most of the API functionality.

## Installation

1. Install go
2. `$ go get github.com/stephan83/vultrapi`
3. Type `vultrapi` to display all available commands

## Examples

### List all regions

		$ vultrapi listregions
		ID  NAME            CONTINENT      COUNTRY  STATE
		39  Miami                          US       FL
		25  Tokyo           Asia           JP
		19  Australia       Australia      AU
		9   Frankfurt       Europe         DE
		24  France          Europe         FR
		8   London          Europe         GB
		7   Amsterdam       Europe         NL
		5   Los Angeles     North America  US       CA
		12  Silicon Valley  North America  US       CA
		6   Atlanta         North America  US       GA
		2   Chicago         North America  US       IL
		1   New Jersey      North America  US       NJ
		3   Dallas          North America  US       TX
		4   Seattle         North America  US       WA

### List plans available in a specific region

		$ vultrapi listplans -region 7
		ID  NAME                                CPUS  PRICE/MONTH
		29  768 MB RAM,15 GB SSD,1.00 TB BW     1     5.00
		30  1024 MB RAM,20 GB SSD,2.00 TB BW    1     7.00
		3   2048 MB RAM,40 GB SSD,3.00 TB BW    2     15.00
		27  4096 MB RAM,65 GB SSD,4.00 TB BW    2     35.00
		28  8192 MB RAM,120 GB SSD,5.00 TB BW   4     70.00
		71  16384 MB RAM,250 GB SSD,8.00 TB BW  4     125.00

### List all operation systems

		$ vultrapi listos
		ID   NAME                    FAMILY    ARCH  WINDOWS
		180  Backup                  backup    x64   false
		163  CentOS 5 i386           centos    i386  false
		147  CentOS 6 i386           centos    i386  false
		162  CentOS 5 x64            centos    x64   false
		127  CentOS 6 x64            centos    x64   false
		167  CentOS 7 x64            centos    x64   false
		179  CoreOS Stable           coreos    x64   false
		152  Debian 7 i386 (wheezy)  debian    i386  false
		139  Debian 7 x64 (wheezy)   debian    x64   false
		140  FreeBSD 10 x64          freebsd   x64   false
		159  Custom                  iso       x64   false
		164  Snapsho                 snapshot  x64   false
		148  Ubuntu 12.04 i386       ubuntu    i386  false
		161  Ubuntu 14.04 i386       ubuntu    i386  false
		182  Ubuntu 14.10 i386       ubuntu    i386  false
		128  Ubuntu 12.04 x64        ubuntu    x64   false
		160  Ubuntu 14.04 x64        ubuntu    x64   false
		181  Ubuntu 14.10 x64        ubuntu    x64   false
		124  Windows 2012 R2 x64     windows   x64   true

### Get help for a command

		$ vultrapi help createserver
		Create a server.

		Usage: vultrapi createserver region_id plan_id os_id [options...]

		You must set env variable VULTR_API_KEY to your API key.

		Options:
		  -enable_auto_backups=false: Enable auto backups
		  -enable_ipv6=false: Enable IPV6
		  -enable_private_network=false: Enable private network
		  -ipxe_chain_url="": IPXE chain url
		  -iso_id=0: ISO ID
		  -label="": Label
		  -script_id=0: Script ID
		  -snapshot_id=0: Snapshot ID
		  -ssh_key_id="": SSH key ID

### Create a Ubuntu server in Amsterdam

		$ VULTR_API_KEY="My API key" vultrapi createserver 7 30 160
		$ SERVER ID:  123456

### Destroy a server

		$ VULTR_API_KEY="My API key" vultrapi destroyserver 123456
		OK

**Many more commands are available**

## Usage

		Usage: vultrapi command [arguments...] [options...]

		You must set env variable VULTR_API_KEY to your API key for commands prefixed with *.

		Commands:

		  help command
		  Get help for a command.

		  listos
		  List all available operating systems.

		  listplans
		  List all available plans.

		  listregions
		  List all available regions.

		* account
		  Get account information.

		* createscript [boot | pxe] name path_to_script
		  Create a script.

		* createserver region_id plan_id os_id
		  Create a server.

		* createsnapshot server_id
		  Create a snapshot.

		* createsshkey name path_to_public_ssh_key
		  Create an SSH key.

		* destroyscript script_id
		  Destroy a script.

		* destroyserver server_id
		  Destroy a server.

		* destroysnapshot snapshot_id
		  Destroy a snapshot.

		* destroysshkey ssh_key_id
		  Destroy an SSH key.

		* listscripts
		  List all scripts.

		* listservers
		  List all servers.

		* listsnapshots
		  List all snapshots.

		* listsshkeys
		  List all SSH keys.

		* script script_id
		  Get script information.

		* server server_id
		  Get server information.

		* sshkey ssh_key_id
		  Get server information.

## Progress

		* public                                                                  [x]
			* listregions                                                         [x]
			* listplans                                                           [x]
			* listos                                                              [x]

		* account                                                                 [x]

		* server                                                                  [ ]
			* createserver                                                        [x]
			* listservers                                                         [x]
			* server                                                              [x]
			* destroyserver                                                       [x]
			* other server commands                                               [ ]

		* sshkey                                                                  [x]
		* snapshot                                                                [x]
		* script                                                                  [x]

		* extra
			* display only specified fields                                       [ ]
			* higher test coverage                                                [ ]

## License

The MIT License (MIT)

		Copyright (c) 2014 Stephan Florquin

		Permission is hereby granted, free of charge, to any person obtaining a copy
		of this software and associated documentation files (the "Software"), to deal
		in the Software without restriction, including without limitation the rights
		to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
		copies of the Software, and to permit persons to whom the Software is
		furnished to do so, subject to the following conditions:

		The above copyright notice and this permission notice shall be included in
		all copies or substantial portions of the Software.

		THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
		IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
		FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
		AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
		LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
		OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
		THE SOFTWARE.

**USE AT YOUR OWN RISK**
