package routers

import (
	"github.com/ind-exe/pulse/oob-server/dns/handlers"

	"github.com/miekg/dns"
)

func CreateGeneralRouter() dns.Handler {
	mux := dns.NewServeMux()
	mux.HandleFunc(".", handlers.Decider)
	return mux
}