package types

import "fmt"

type Account struct {
	Balance           float64 `json:"balance"`
	PendingCharges    float64 `json:"pending_charges"`
	LastPaymentDate   Date    `json:"last_payment_date"`
	LastPaymentAmount float64 `json:"last_payment_amount,string"`
}

func (o Account) String() string {
	return fmt.Sprintf("%19s: %.2f\n%19s: %.2f\n%19s: %s\n%19s: %.2f",
		"BALANCE", o.Balance,
		"PENDING CHARGES", o.PendingCharges,
		"LAST PAYMENT DATE", o.LastPaymentDate,
		"LAST PAYMENT AMOUNT", o.LastPaymentAmount)
}
