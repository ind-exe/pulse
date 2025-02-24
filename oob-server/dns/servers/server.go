package servers

import (
	"log"
	"oob-server/oob-server/dns/routers"

	"github.com/miekg/dns"
)

const (
	DnsPort = 53
)

func InitDnsServer() {
	mux := routers.CreateGeneralRouter()
	log.Printf("HTTPListener | success : HTTP server started on port %d\n", DnsPort)
	err := dns.ListenAndServe(":53", "udp", mux)
	if err != nil{
		log.Printf("DnsServer | fail : %s\n", err)
	}

}