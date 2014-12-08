package commands

import (
	"fmt"
	"github.com/stephan83/vultrapi/requests"
	"os"
)

type account struct{}

func NewAccount() Command {
	return account{}
}

func (_ account) NeedsKey() bool {
	return true
}

func (_ account) Args() string {
	return ""
}

func (_ account) Desc() string {
	return "Get account information."
}

func (_ account) PrintOptions() {
	fmt.Println("None.")
}

func (_ account) Exec() (err error) {
	OS, err := requests.GetAccount(os.Getenv("VULTR_API_KEY"))

	if err != nil {
		return
	}

	fmt.Println(OS)

	return
}
