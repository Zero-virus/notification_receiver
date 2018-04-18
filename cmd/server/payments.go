package payments

import (
	"net/http"

	"github.com/Finciero/ghipu"
	"github.com/Finciero/thief/sigiriya"
)

type paymentResponse struct {
	TransactionID string  `json:"transaction_id"`
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Source        string  `json:"source"`
}

type Context struct {
	KhipuClient    *ghipu.Client
	SigiriyaClient *sigiriya.CLient
}

type HandlerFunc func(*Context, http.ResponseWriter, *http.Request) error
