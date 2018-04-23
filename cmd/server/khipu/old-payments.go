package payments

import (
	"errors"
	"net/http"

	"context"

	"github.com/Finciero/ghipu"
	"github.com/Finciero/thief/sigiriya"
	raven "github.com/getsentry/raven-go"
)

type paymentResponse struct {
	TransactionID string  `json:"transaction_id"`
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Source        string  `json:"source"`
}

type ctxKey string

const (
	paymentKey    ctxKey = "payment"
	paymentSource        = "khipu"
)

var errAlreadyNotified = errors.New("payments: payment already notified")

// Context contains all services and variables of the applications.
type Context struct {
	KhipuClient    *ghipu.Client
	SigiriyaClient *sigiriya.Client
}

// HandlerFunc function handler signature used by sigiriya application.
type HandlerFunc func(*Context, http.ResponseWriter, *http.Request) error

// Handle creates a new bounded http.HandlerFunc with context.
func (c *Context) Handle(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			reportError(err, r)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiVersion := r.PostFormValue("api_version")
		if apiVersion != "1.3" {
			http.Error(w, "Invalid API Version", http.StatusInternalServerError)
			return
		}

		notificationToken := r.PostFormValue("notification_token")
		p, err := c.KhipuClient.PaymentStatus(notificationToken)
		if err != nil {
			reportError(err, r)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), paymentKey, p))
		if err := h(c, w, r); err != nil {
			reportError(err, r)
			if err == errAlreadyNotified {
				http.Error(w, "Invalid Payload", http.StatusBadRequest)
			} else {
				http.Error(w, "Invalid Payload", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func reportError(err error, r *http.Request) {
	m := make(map[string]string)
	m["method"] = r.Method
	m["error"] = err.Error()
	m["handler"] = "payment"
	m["error_type"] = "Unhandle API Error"
	if r.Method == "POST" {
		m["post_values"] = r.PostForm.Encode()
	}
	raven.SetHttpContext(raven.NewHttp(r))
	raven.SetTagsContext(m)
	raven.CaptureError(err, nil)
	raven.ClearContext()
}
