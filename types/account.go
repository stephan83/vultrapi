package types

import (
	"fmt"
)

type Account struct {
	Balance           int     `json:"balance"`
	PendingCharges    float64 `json:"pending_charges"`
	LastPaymentDate   string  `json:"last_payment_date"`
	LastPaymentAmount string  `json:"last_payment_amount"`
}

func (a Account) String() string {
	return fmt.Sprintf("%20s %d\n%20s %.2f\n%20s %s\n%20s %s",
		"BALANCE:", a.Balance,
		"PENDING CHARGES:", a.PendingCharges,
		"LAST PAYMENT DATE:", a.LastPaymentDate,
		"LAST PAYMENT AMOUNT:", a.LastPaymentAmount)
}
