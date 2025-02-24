package handlers

import (
	"log"
	"strings"
	"sync"

	"github.com/ind-exe/pulse/data"

	"github.com/miekg/dns"
)

func Decider(w dns.ResponseWriter, r *dns.Msg) {
	response := new(dns.Msg)
	response.SetReply(r)
	response.Authoritative = true
	
	wg := sync.WaitGroup{}

	for _, question := range r.Question {
		wg.Add(1)
		domain := strings.ToLower(question.Name)

		go NotifHandler(domain, r, w, &wg)

		wg.Wait()

		if records, found := data.DnsRecords[domain]; found {
			for _, rr := range records {
				if rr.Header().Rrtype == question.Qtype {
					response.Answer = append(response.Answer, rr)
				}
			}
		}
	}

	if err := w.WriteMsg(response); err != nil {
		log.Printf("DnsHandler | fail : Failed to write message: %v", err)
	}
}