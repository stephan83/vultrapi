package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
)

type listRegions struct{}

func NewListRegions() Command {
	return listRegions{}
}

func (_ listRegions) Desc() string {
	return "List all available regions."
}

func (_ listRegions) Args() string {
	return ""
}

func (_ listRegions) NeedsKey() bool {
	return false
}

func (_ listRegions) PrintOptions() {
	fmt.Println("None.")
}

func (_ listRegions) Exec(c Client, args []string, key string) (err error) {
	regions, err := requests.GetRegions(c)
	if err != nil {
		return
	}

	fmt.Println(regions)

	return
}
