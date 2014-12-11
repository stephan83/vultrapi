package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleAccount() {
	c := NewTestClient(200, acc)
	NewAccount().Fexec(os.Stdout, c, []string{}, "SECRET_KEY")
	// Output:
	// BALANCE			-5.00
	// PENDING CHARGES		0.26
	// LAST PAYMENT DATE	2014-12-07 15:24:55 +0000
	// LAST PAYMENT AMOUNT	-5.00
}

var acc = []byte(`{
	"balance": -5,
	"pending_charges": 0.26,
	"last_payment_date": "2014-12-07 10:24:55",
	"last_payment_amount": "-5.00"
}`)
