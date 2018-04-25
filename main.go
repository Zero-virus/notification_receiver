package notification2

import (
	"finciero/notifications_test/sigiriya"
	"flag"
	"fmt"
	"net/http"
	"os"

	paymentsApi "github.com/Finciero/notification_receiver-1/tree/master/cmd/server/khipu"

	"github.com/sirupsen/logrus"

	"github.com/urfave/negroni"
)

type paymentResponse struct {
	TransactionID string  `json:"transaction_id"`
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Source        string  `json:"source"`
}

/*func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, appUsage)
		for _, c := range subcmds {
			fmt.Fprintf(os.Stderr, "")
		}
	}
}*/

func AllNoti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindNoti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not")
}

func CreateNoti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not")
}

var logger *logrus.Logger

type subcmd struct {
	name        string
	description string
	run         func(args []string)
}

var subcmds = []subcmd{
	{"serve", "start web server", serveCmd},
}

func main() {
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	subcmd := flag.Arg(0)
	for _, c := range subcmds {
		if c.name == subcmd {
			c.run(flag.Args()[1:])
			return
		}
	}

	fmt.Fprintf(os.Stderr, "unknow subcmd %q\n", subcmd)
	fmt.Fprintln(os.Stderr, `Run "thief -h" for usage.`)
	os.Exit(1)
}

func serveCmd(args []string) {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)

	host := fs.String("host", "", "HTTP server host")
	port := fs.Int("port", 8092, "HTTP server port")

	sigiriyaAPIURL := fs.String("sigiriya-api-url", "http://localhost:8090", "list of allowed origins separated by comma.")
	sigiriyaToken := fs.String("auth-secret-key", "development_token", "Sigiriya authentication token")

	sigiriyaClient := sigiriya.NewClient(&sigiriya.Config{
		APIURL: *sigiriyaAPIURL,
		Token:  *sigiriyaToken,
	})
	logger.Println("Connected to sigiriya")

	ctxPayments := &paymentsApi.Context{
		KhipuClient:    khipuClient,
		SigiriyaClient: sigiriyaClient,
	}

	n := negroni.New(
		negroni.NewRecovery(),
		negronilogrus.NewMiddlewareFromLogger(logger, "Notifications API"), recovery.JSONRecovery(false),
	)

	mux := http.NewServeMux()
	mux.Handle("/notifications/khipu/", http.StripPrefix("/notifications/khipu", paymentsApi.Handle(ctxPayments)))

	n.UseHandler(mux)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	logger.Printf("Start listening on %s", addr)
	if *port == 443 {
		logger.Fatal(http.ListenAndServeTLS(addr, *certFile, *keyFile, n))
	} else {
		logger.Fatal(http.ListenAndServe(addr, n))
	}
}

func check(err error) {
	if err != nil {
		logger.fatal(err)
	}
}
