package commands

import(
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleListSnapshots() {
	c := NewTestClient(200, snapshots)
	NewListSnapshots().Fexec(os.Stdout, c, []string{}, "")
	// ID		DESCRIPTION		DATE CREATED			SIZE		STATUS
	// 5488d7f7e521e				2014-12-10 23:32:07 +0000	21474836480	complete
	// 5488da65095a6				2014-12-10 23:42:29 +0000	21474836480	complete
	// 5488da448aa0b	Testing snapshots.	2014-12-10 23:41:56 +0000	21474836480	complete
}

var snapshots = []byte(`{
	"5488d7f7e521e": {
		"SNAPSHOTID": "5488d7f7e521e",
		"date_created": "2014-12-10 18:32:07",
		"description": "",
		"size": "21474836480",
		"status": "complete"
	},
	"5488da448aa0b": {
		"SNAPSHOTID": "5488da448aa0b",
		"date_created": "2014-12-10 18:41:56",
		"description": "Testing snapshots.",
		"size": "21474836480",
		"status": "complete"
	},
	"5488da65095a6": {
		"SNAPSHOTID": "5488da65095a6",
		"date_created": "2014-12-10 18:42:29",
		"description": "",
		"size": "21474836480",
		"status": "complete"
	}
}`)
