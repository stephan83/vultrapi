package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"text/tabwriter"
	"io"
)

type account struct{ BasicCommand }

func NewAccount() Command {
	return &account{
		BasicCommand{Desc: "Get account information.", NeedsKey: true},
	}
}

func (_ *account) Fexec(w io.Writer, c Client, _ []string, key string) (err error) {
	r, err := requests.GetAccount(c, key)
	if err != nil {
		return
	}

	t := tabwriter.NewWriter(w, 0, 8, 1, '\t', 0)

	fmt.Fprintf(t, "BALANCE\t%.2f\n", r.Balance)
	fmt.Fprintf(t, "PENDING CHARGES\t%.2f\n", r.PendingCharges)
	fmt.Fprintf(t, "LAST PAYMENT DATE\t%s\n", r.LastPaymentDate)
	fmt.Fprintf(t, "LAST PAYMENT AMOUNT\t%.2f\n", r.LastPaymentAmount)

	t.Flush()

	return
}
