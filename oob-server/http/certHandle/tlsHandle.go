package certhandle

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/ind-exe/pulse/data"

	"github.com/caddyserver/certmagic"
)

func TLSHandler() *tls.Config{
	provider := &MyDNSProvider{}

	// Configure CertMagic for DNS-01 challenge
	certmagic.DefaultACME.DNS01Solver = &certmagic.DNS01Solver{
		DNSManager: certmagic.DNSManager{
			DNSProvider:       provider,
			PropagationTimeout: 2 * time.Minute, // Adjust if DNS is slow
		},
	}

	// Disable HTTP and TLS-ALPN challenges (only use DNS-01)
	certmagic.DefaultACME.DisableHTTPChallenge = true
	certmagic.DefaultACME.DisableTLSALPNChallenge = true

	// Required ACME fields
	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = "certIssue@" + data.Config.Domain
	certmagic.DefaultACME.CA = certmagic.LetsEncryptProductionCA // Change to Production later

	// Specify the domain(s) for the certificate
	domains := GetAllCertRequiredDomains()

	// Obtain TLS configuration from CertMagic
	tlsConfig, err := certmagic.TLS(domains)
	if err != nil {
		log.Fatal("TlsHandler | fail : Failed to get TLS config:", err)
	}

	// Explicitly enable HTTP/2
	tlsConfig.NextProtos = []string{"h2", "http/1.1"} // Explicitly add HTTP/2 and HTTP/1.1 support

	return tlsConfig

}