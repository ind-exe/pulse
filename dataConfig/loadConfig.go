package dataconfig

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/ind-exe/pulse/data"
	envvar "github.com/ind-exe/pulse/envVar"

	"github.com/miekg/dns"
)

func loadEnvVariablesToMemory() error {
	DiscordWebhook, err := envvar.GetVar("DISCORD_WEBHOOK")
	if err != nil {
		return err
	}
	data.Config.DiscordWebhook = DiscordWebhook

	TelegramBotToken, err := envvar.GetVar("TELEGRAM_BOT_TOKEN")
	if err != nil {
		return err
	}
	data.Config.TelegramBotToken = TelegramBotToken

	TelegramChatId, err := envvar.GetVar("TELEGRAM_CHAT_ID")
	if err != nil {
		return err
	}
	data.Config.TelegramChatId = TelegramChatId

	Domain, err := envvar.GetVar("DOMAIN")
	if err != nil {
		return err
	}
	data.Config.Domain = Domain
	return nil
}

func LoadConfigurations() error {
	data.DnsRecordsMu = &sync.RWMutex{}
	var baseFolder = "/etc/github.com/ind-exe/pulse/"
	// Load DNS records.
	dnsData, err := os.ReadFile(baseFolder + "dns_records.json")
	if err != nil {
		return err
	}

	var dnsRecordsFile map[string][]string
	if err := json.Unmarshal(dnsData, &dnsRecordsFile); err != nil {
		return err
	}
	for domain, rrStrs := range dnsRecordsFile {
		for _, rrStr := range rrStrs {
			rr, err := dns.NewRR(rrStr)
			if err != nil {
				log.Printf("Error parsing RR for domain %s: %v", domain, err)
				continue
			}
			data.DnsRecords[domain] = append(data.DnsRecords[domain], rr)
		}
	}

	// Load DomainMap.
	domainMapData, err := os.ReadFile(baseFolder + "domain_map.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(domainMapData, &data.DomainMap); err != nil {
		return err
	}

	// Load UrlMap.
	urlMapData, err := os.ReadFile(baseFolder + "url_map.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(urlMapData, &data.UrlMap); err != nil {
		return err
	}

	// Load UrlServeMap.
	urlServeMapData, err := os.ReadFile(baseFolder + "url_serve_map.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(urlServeMapData, &data.UrlServeMap); err != nil {
		return err
	}

	if err := loadEnvVariablesToMemory(); err != nil {
		return err
	}

	return nil
}