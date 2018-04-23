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


func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, appUsage)
		for -, c := range subcmds {
			fmt.Fprintf(os.Stderr, "")
		}
	}
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
	flag.Parse()

	if flag.
}

func serveCmd(args []string) {

	sigiriyaAPIURL := fs.String("sigiriya-api-url", "http://localhost:8090", "list of allowed origins separated by comma.")
	sigiriyaToken := fs.String("auth-secret-key", "development_token", "Sigiriya authentication token")

	sigiriyaClient := sigiriya.NewClient(&sigiriya.Config{
		APIURL: *sigiriyaAPIURL,
		Token:	*sigiriyaToken,
	})
	logger.Println("Connected to sigiriya")

	ctxPayments := &paymentsApi.Context{
		KhipuCLient:	khipuCLient,
		SigiriyaClient:	sigiriyaClient,
	}

	n := negroni.New(
		negroni.NewRecovery(),
		negronilogrus.NewMiddlewareFromLogger(logger, "Notifications API"), recovery.JSONRecovery(false),
	)


	mux := http.NewServeMux()
	mux.Handle("/notifications/khipu/",http.StripPrefix("/notifications/khipu",paymentsApi.Handle(ctxPayments)))

	n.UseHandler(mux)

	addr :=
}