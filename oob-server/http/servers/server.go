package servers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	certhandle "github.com/ind-exe/pulse/oob-server/http/certHandle"
	"github.com/ind-exe/pulse/oob-server/http/routers"

	"golang.org/x/net/http2"
)

const (
	HttpPort = 80
	HttpsPort = 443
)

func InitHttpServer() {
	mux := routers.CreateGeneralRouter()
	log.Printf("HTTPListener | success : HTTP server started on port %d\n", HttpPort )
	if err := http.ListenAndServe(fmt.Sprintf(":%d", HttpPort), mux); err != nil {
		fmt.Println(err)
		log.Fatalf("HTTPListener | fail : %s\n", err)
	}
}

func InitHttpsServer() {
	tlsConfig := certhandle.TLSHandler()
	
	ln, err := tls.Listen("tcp", ":443", tlsConfig)
	if err != nil {
		log.Fatalf("TLSListener | fail : %s\n", err)
	}

	log.Printf("HTTPListener | success : HTTP server started on port %d\n", HttpsPort)
	server := &http.Server{
		Handler: routers.CreateGeneralRouter(),
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, nil)

	err = server.Serve(ln)
	if err != nil {
		log.Fatalf("TLSServe | fail : %s\n", err)
	}
}