package commands

import (
	"fmt"
	"github.com/stephan83/vultrapi/requests"
)

type listOS struct{}

func NewListOS() Command {
	return listOS{}
}

func (_ listOS) NeedsKey() bool {
	return false
}

func (_ listOS) Args() string {
	return ""
}

func (_ listOS) Desc() string {
	return "List all available operating systems."
}

func (_ listOS) PrintOptions() {
	fmt.Println("None.")
}

func (_ listOS) Exec() (err error) {
	OS, err := requests.GetOS()
	if err != nil {
		return
	}

	fmt.Println(OS)

	return
}
