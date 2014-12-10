package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"text/tabwriter"
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
	r, err := requests.GetAccount(c, key)

	if err != nil {
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintf(w, "BALANCE\t%.2f\n", r.Balance)
	fmt.Fprintf(w, "PENDING CHARGES\t%.2f\n", r.PendingCharges)
	fmt.Fprintf(w, "LAST PAYMENT DATE\t%s\n", r.LastPaymentDate)
	fmt.Fprintf(w, "LAST PAYMENT AMOUNT\t%.2f\n", r.LastPaymentAmount)

	w.Flush()

	return
}
