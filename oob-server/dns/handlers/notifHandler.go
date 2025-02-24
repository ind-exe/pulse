package handlers

import (
	"oob-server/data"
	"oob-server/models"
	"strconv"
	"sync"
	"time"

	"github.com/miekg/dns"
)

func NotifHandler(chechDomain string, r *dns.Msg, w dns.ResponseWriter, wg *sync.WaitGroup) {
	if value, ok := data.DomainMap[chechDomain]; ok {
		LogData := models.DnsModel{
			Type: strconv.Itoa(int(r.Question[0].Qtype)),
			IP: w.RemoteAddr().String(),
			Timestamp: time.Now(),
			Question: r.Question[0].String(),
		}
		wg.Done()
		defer w.Close()
		value.Decider(models.Notifier(&LogData))
	} else {
		wg.Done()
	}
}
