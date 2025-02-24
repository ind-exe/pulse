package data

import (
	"sync"

	"github.com/miekg/dns"
)

var (
	DnsRecords = map[string][]dns.RR{}
	DnsRecordsMu *sync.RWMutex
)
