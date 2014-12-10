package types

import "fmt"

type Account struct {
	Balance           int     `json:"balance"`
	PendingCharges    float64 `json:"pending_charges"`
	LastPaymentDate   string  `json:"last_payment_date"`
	LastPaymentAmount string  `json:"last_payment_amount"`
}

func (o Account) String() string {
	return fmt.Sprintf("%19s: %d\n%19s: %.2f\n%19s: %s\n%19s: %s",
		"BALANCE", o.Balance,
		"PENDING CHARGES", o.PendingCharges,
		"LAST PAYMENT DATE", o.LastPaymentDate,
		"LAST PAYMENT AMOUNT", o.LastPaymentAmount)
}
