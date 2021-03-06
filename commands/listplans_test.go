package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleListPlans() {
	c := NewTestClient(200, plans)
	NewListPlans().Fexec(os.Stdout, c, []string{}, "")
	// Output:
	// ID	NAME						CPUS	PRICE/MONTH
	// 11	512 MB RAM,160 GB SATA,1.00 TB BW		1	5.00
	// 31	768 MB RAM,15 GB SSD,0.20 TB BW			1	5.00
	// 29	768 MB RAM,15 GB SSD,1.00 TB BW			1	5.00
	// 32	1024 MB RAM,20 GB SSD,0.40 TB BW		1	7.00
	// 30	1024 MB RAM,20 GB SSD,2.00 TB BW		1	7.00
	// 12	1024 MB RAM,320 GB SATA,2.00 TB BW		1	8.00
	// 62	1024 MB RAM,320 GB SATAPERF,1.00 TB BW, 10GigE	1	15.00
	// 13	2048 MB RAM,640 GB SATA,3.00 TB BW		1	15.00
	// 63	2048 MB RAM,640 GB SATAPERF,2.00 TB BW, 10GigE	1	25.00
	// 8	2048 MB RAM,40 GB SSD,0.60 TB BW		2	15.00
	// 3	2048 MB RAM,40 GB SSD,3.00 TB BW		2	15.00
	// 33	4096 MB RAM,65 GB SSD,0.80 TB BW		2	35.00
	// 27	4096 MB RAM,65 GB SSD,4.00 TB BW		2	35.00
	// 64	4096 MB RAM,1280 GB SATAPERF,3.00 TB BW, 10GigE	2	50.00
	// 34	8192 MB RAM,120 GB SSD,1.00 TB BW		4	70.00
	// 28	8192 MB RAM,120 GB SSD,5.00 TB BW		4	70.00
	// 68	16384 MB RAM,250 GB SSD,1.60 TB BW		4	125.00
	// 71	16384 MB RAM,250 GB SSD,8.00 TB BW		4	125.00
	// 78	16384 MB RAM,250 GB SSD,6.00 TB BW		8	149.95
	// 79	32768 MB RAM,400 GB SSD,7.00 TB BW		12	299.95
	// 80	49152 MB RAM,600 GB SSD,8.00 TB BW		16	429.95
	// 81	65536 MB RAM,800 GB SSD,9.00 TB BW		24	599.95
}

func ExampleListPlansWithRegion() {
	c := NewTestMultiClient(map[string]Client{
		"^/regions/availability.*": NewTestClient(200, availability),
		"^/plans/list.*":           NewTestClient(200, plans),
	})
	NewListPlans().Fexec(os.Stdout, c, []string{"-region", "24"}, "")
	// Output:
	// ID	NAME					CPUS	PRICE/MONTH
	// 29	768 MB RAM,15 GB SSD,1.00 TB BW		1	5.00
	// 30	1024 MB RAM,20 GB SSD,2.00 TB BW	1	7.00
	// 3	2048 MB RAM,40 GB SSD,3.00 TB BW	2	15.00
	// 27	4096 MB RAM,65 GB SSD,4.00 TB BW	2	35.00
	// 28	8192 MB RAM,120 GB SSD,5.00 TB BW	4	70.00
	// 71	16384 MB RAM,250 GB SSD,8.00 TB BW	4	125.00
}

var plans = []byte(`{
	"31": {
		"VPSPLANID": "31",
		"name": "768 MB RAM,15 GB SSD,0.20 TB BW",
		"vcpu_count": "1",
		"ram": "768",
		"disk": "15",
		"bandwidth": "0.20",
		"bandwidth_gb": "204.8",
		"price_per_month": "5.00",
		"windows": false
	},
	"29": {
		"VPSPLANID": "29",
		"name": "768 MB RAM,15 GB SSD,1.00 TB BW",
		"vcpu_count": "1",
		"ram": "768",
		"disk": "15",
		"bandwidth": "1.00",
		"bandwidth_gb": "1024",
		"price_per_month": "5.00",
		"windows": false
	},
	"32": {
		"VPSPLANID": "32",
		"name": "1024 MB RAM,20 GB SSD,0.40 TB BW",
		"vcpu_count": "1",
		"ram": "1024",
		"disk": "20",
		"bandwidth": "0.40",
		"bandwidth_gb": "409.6",
		"price_per_month": "7.00",
		"windows": false
	},
	"30": {
		"VPSPLANID": "30",
		"name": "1024 MB RAM,20 GB SSD,2.00 TB BW",
		"vcpu_count": "1",
		"ram": "1024",
		"disk": "20",
		"bandwidth": "2.00",
		"bandwidth_gb": "2048",
		"price_per_month": "7.00",
		"windows": false
	},
	"3": {
		"VPSPLANID": "3",
		"name": "2048 MB RAM,40 GB SSD,3.00 TB BW",
		"vcpu_count": "2",
		"ram": "2048",
		"disk": "40",
		"bandwidth": "3.00",
		"bandwidth_gb": "3072",
		"price_per_month": "15.00",
		"windows": false
	},
	"8": {
		"VPSPLANID": "8",
		"name": "2048 MB RAM,40 GB SSD,0.60 TB BW",
		"vcpu_count": "2",
		"ram": "2048",
		"disk": "40",
		"bandwidth": "0.60",
		"bandwidth_gb": "614.4",
		"price_per_month": "15.00",
		"windows": false
	},
	"33": {
		"VPSPLANID": "33",
		"name": "4096 MB RAM,65 GB SSD,0.80 TB BW",
		"vcpu_count": "2",
		"ram": "4096",
		"disk": "65",
		"bandwidth": "0.80",
		"bandwidth_gb": "819.2",
		"price_per_month": "35.00",
		"windows": false
	},
	"27": {
		"VPSPLANID": "27",
		"name": "4096 MB RAM,65 GB SSD,4.00 TB BW",
		"vcpu_count": "2",
		"ram": "4096",
		"disk": "65",
		"bandwidth": "4.00",
		"bandwidth_gb": "4096",
		"price_per_month": "35.00",
		"windows": false
	},
	"28": {
		"VPSPLANID": "28",
		"name": "8192 MB RAM,120 GB SSD,5.00 TB BW",
		"vcpu_count": "4",
		"ram": "8192",
		"disk": "120",
		"bandwidth": "5.00",
		"bandwidth_gb": "5120",
		"price_per_month": "70.00",
		"windows": false
	},
	"34": {
		"VPSPLANID": "34",
		"name": "8192 MB RAM,120 GB SSD,1.00 TB BW",
		"vcpu_count": "4",
		"ram": "8192",
		"disk": "120",
		"bandwidth": "1.00",
		"bandwidth_gb": "1024",
		"price_per_month": "70.00",
		"windows": false
	},
	"11": {
		"VPSPLANID": "11",
		"name": "512 MB RAM,160 GB SATA,1.00 TB BW",
		"vcpu_count": "1",
		"ram": "512",
		"disk": "160",
		"bandwidth": "1.00",
		"bandwidth_gb": "1024",
		"price_per_month": "5.00",
		"windows": false
	},
	"78": {
		"VPSPLANID": "78",
		"name": "16384 MB RAM,250 GB SSD,6.00 TB BW",
		"vcpu_count": "8",
		"ram": "16384",
		"disk": "250",
		"bandwidth": "6.00",
		"bandwidth_gb": "6144",
		"price_per_month": "149.95",
		"windows": false
	},
	"71": {
		"VPSPLANID": "71",
		"name": "16384 MB RAM,250 GB SSD,8.00 TB BW",
		"vcpu_count": "4",
		"ram": "16384",
		"disk": "250",
		"bandwidth": "8.00",
		"bandwidth_gb": "8192",
		"price_per_month": "125.00",
		"windows": false
	},
	"68": {
		"VPSPLANID": "68",
		"name": "16384 MB RAM,250 GB SSD,1.60 TB BW",
		"vcpu_count": "4",
		"ram": "16384",
		"disk": "250",
		"bandwidth": "1.60",
		"bandwidth_gb": "1638.4",
		"price_per_month": "125.00",
		"windows": false
	},
	"12": {
		"VPSPLANID": "12",
		"name": "1024 MB RAM,320 GB SATA,2.00 TB BW",
		"vcpu_count": "1",
		"ram": "1024",
		"disk": "320",
		"bandwidth": "2.00",
		"bandwidth_gb": "2048",
		"price_per_month": "8.00",
		"windows": false
	},
	"62": {
		"VPSPLANID": "62",
		"name": "1024 MB RAM,320 GB SATAPERF,1.00 TB BW, 10GigE",
		"vcpu_count": "1",
		"ram": "1024",
		"disk": "320",
		"bandwidth": "1.00",
		"bandwidth_gb": "1024",
		"price_per_month": "15.00",
		"windows": false
	},
	"79": {
		"VPSPLANID": "79",
		"name": "32768 MB RAM,400 GB SSD,7.00 TB BW",
		"vcpu_count": "12",
		"ram": "32768",
		"disk": "400",
		"bandwidth": "7.00",
		"bandwidth_gb": "7168",
		"price_per_month": "299.95",
		"windows": false
	},
	"80": {
		"VPSPLANID": "80",
		"name": "49152 MB RAM,600 GB SSD,8.00 TB BW",
		"vcpu_count": "16",
		"ram": "49152",
		"disk": "600",
		"bandwidth": "8.00",
		"bandwidth_gb": "8192",
		"price_per_month": "429.95",
		"windows": false
	},
	"13": {
		"VPSPLANID": "13",
		"name": "2048 MB RAM,640 GB SATA,3.00 TB BW",
		"vcpu_count": "1",
		"ram": "2048",
		"disk": "640",
		"bandwidth": "3.00",
		"bandwidth_gb": "3072",
		"price_per_month": "15.00",
		"windows": false
	},
	"63": {
		"VPSPLANID": "63",
		"name": "2048 MB RAM,640 GB SATAPERF,2.00 TB BW, 10GigE",
		"vcpu_count": "1",
		"ram": "2048",
		"disk": "640",
		"bandwidth": "2.00",
		"bandwidth_gb": "2048",
		"price_per_month": "25.00",
		"windows": false
	},
	"81": {
		"VPSPLANID": "81",
		"name": "65536 MB RAM,800 GB SSD,9.00 TB BW",
		"vcpu_count": "24",
		"ram": "65536",
		"disk": "800",
		"bandwidth": "9.00",
		"bandwidth_gb": "9216",
		"price_per_month": "599.95",
		"windows": false
	},
	"64": {
		"VPSPLANID": "64",
		"name": "4096 MB RAM,1280 GB SATAPERF,3.00 TB BW, 10GigE",
		"vcpu_count": "2",
		"ram": "4096",
		"disk": "1280",
		"bandwidth": "3.00",
		"bandwidth_gb": "3072",
		"price_per_month": "50.00",
		"windows": false
	}
}`)

var availability = []byte("[29,30,3,27,28,71]")
