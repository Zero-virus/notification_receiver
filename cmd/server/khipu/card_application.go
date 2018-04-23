package payments

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Finciero/ghipu"
)

func handleApplicationPayment(ctx *Context, w http.ResponseWriter, r *http.Request) error {
	payment := r.Context().Value(paymentKey).(*ghipu.PaymentResponse)

	if payment.Status != "done" {
		return nil
	}

	b, err := json.Marshal(&paymentResponse{
		TransactionID: payment.TransactionID,
		PaymentID:     payment.PaymentID,
		Amount:        payment.Amount,
		Status:        payment.Status,
		Source:        paymentSource,
	})
	if err != nil {
		return err
	}

	if _, err := ctx.SigiriyaClient.Post("/payments/card-applications", bytes.NewBuffer(b)); err != nil {
		return err
	}

	return nil
}
