package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
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

func (_ account) Exec(c Client, _ []string, key string) (err error) {
	a, err := requests.GetAccount(c, key)

	if err != nil {
		return
	}

	fmt.Println(a)

	return
}
