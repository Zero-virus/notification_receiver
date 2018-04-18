package payments

import (
	"net/http"

	"github.com/Finciero/ghipu"
	"github.com/Finciero/thief/sigiriya"
)

type paymentResponse struct {
	TransactionID string  `bson:"transactionID" json:"transaction_id"`
	PaymentID     string  `bson:"paymentID" json:"payment_id"`
	Amount        float64 `bson:"amount" json:"amount"`
	Status        string  `bson:"status" json:"status"`
	Source        string  `bson:"source" json:"source"`
}

type ctxKey string

const (
	paymentKey	ctxKey	= "payment"
	paymentSource		= "khipu"
)

var errAlreadyNotified = errors.New("payments: payment already notified")

type Context struct {
	KhipuClient    *ghipu.Client
	SigiriyaClient *sigiriya.CLient
}

type HandlerFunc func(*Context, http.ResponseWriter, *http.Request) error

func (c  *Context) Handle(h HandlerFunc) http.HandlerFunc {
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

}
)

