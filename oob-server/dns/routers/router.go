package routers

import (
	"oob-server/oob-server/dns/handlers"

	"github.com/miekg/dns"
)

func CreateGeneralRouter() dns.Handler {
	mux := dns.NewServeMux()
	mux.HandleFunc(".", handlers.Decider)
	return mux
}