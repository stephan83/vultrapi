# VULTRAPI [![Build Status](https://travis-ci.org/stephan83/vultrapi.svg?branch=master)](https://travis-ci.org/stephan83/vultrapi)

Client/Library for [Vultr API](https://vultr.com) written in go.

**WORK IN PROGRESS**

## Installation

1. Install go
2. `$ go get github.com/stephan83/vultrapi`

## Usage

		Usage: vultrapi command [options...]

		You must set env variable VULTR_API_KEY to your API key for underlined commands.

		Commands:

		  help command
		  Get help for a command.

		  listos
		  List all available operating systems.

		  listplans
		  List all available plans.

		  listregions
		  List all available regions.

		  account
		  *******
		  Get account information.

		  createserver region_id plan_id os_id
		  ************
		  Creates a server.

		  destroyserver server_id
		  *************
		  Destroys a server.

		  listservers
		  ***********
		  List all servers.

		  server server_id
		  ******
		  Get server information.

## Example

		$ vultrapi listplans -region 24
		NAME                                               | CPUS | PRICE/MONTH | ID 
		------------------------------------------------------------------------------
		768 MB RAM,15 GB SSD,1.00 TB BW                    | 1    | 5.00        | 29 
		1024 MB RAM,20 GB SSD,2.00 TB BW                   | 1    | 7.00        | 30 
		2048 MB RAM,40 GB SSD,3.00 TB BW                   | 2    | 15.00       | 3  
		4096 MB RAM,65 GB SSD,4.00 TB BW                   | 2    | 35.00       | 27 
		8192 MB RAM,120 GB SSD,5.00 TB BW                  | 4    | 70.00       | 28 
		16384 MB RAM,250 GB SSD,8.00 TB BW                 | 4    | 125.00      | 71 

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

		* snapshot                                                                [ ]
		* sshkey                                                                  [ ]
		* script                                                                  [ ]

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
