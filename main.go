package notification2

type paymentResponse struct {
	TransactionID string  `json:"transaction_id"`
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Source        string  `json:"source"`
}
