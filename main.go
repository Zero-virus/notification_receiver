package notification2

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type paymentResponse struct {
	TransactionID string  `json:"transaction_id"`
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Source        string  `json:"source"`
}

func AllNoti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindNoti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not")
}

func CreateNoti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notification", AllNoti).Methods("GET")
	r.HandleFunc("/notification", CreateNoti).Methods("POST")
	r.HandleFunc("/notification/{id}", FindNoti).Methods("GET")
}
