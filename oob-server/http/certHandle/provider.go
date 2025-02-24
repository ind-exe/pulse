package certhandle

import (
	"context"
	"fmt"
	"log"

	"github.com/ind-exe/pulse/data"

	"github.com/libdns/libdns"
	"github.com/miekg/dns"
)

type MyDNSProvider struct{}

func (c *MyDNSProvider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	for _, record := range records {
		fmt.Printf("Adding TXT record: %s -> %s\n", record.Name, record.Value)
		fqdn := dns.Fqdn(record.Name) + zone
		txtRecord, err := dns.NewRR(fmt.Sprintf("%s 300 IN TXT %s", fqdn, record.Value))
		fmt.Println(txtRecord.String())
		if err != nil {
			log.Printf("AcmeRecord | fail : %s\n", err)
		}
		data.DnsRecordsMu.Lock()
		data.DnsRecords[fqdn] = []dns.RR{txtRecord}
		data.DnsRecordsMu.Unlock()
	}
	return records, nil
}

func (c *MyDNSProvider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	for _, record := range records {
		fqdn := dns.Fqdn(record.Name) + zone
		fmt.Printf("Removing TXT record: %s -> %s\n", record.Name, record.Value)
		data.DnsRecordsMu.Lock()
		delete(data.DnsRecords, fqdn)
		data.DnsRecordsMu.Unlock()
	}
	return records, nil
}

