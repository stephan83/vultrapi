package types

type Account struct {
	Balance           float64 `json:"balance"`
	PendingCharges    float64 `json:"pending_charges"`
	LastPaymentDate   Date    `json:"last_payment_date"`
	LastPaymentAmount float64 `json:"last_payment_amount,string"`
}
