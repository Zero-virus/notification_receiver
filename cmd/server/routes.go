package payments

import (
	"net/http"

	"github.com/Finciero/thief/cmd/server/router"
)

func Handle(c *Context) http.Handler {
	r := router.New()

	r.POST("/card-applications", c.Handle(handleApplicationPayment))
	r.POST("/card-deposits", c.Handle(handleDepositPayment))
	r.POST("/renew-applications", c.Handle(handleRenewApplicationPayment))

	return r
}
