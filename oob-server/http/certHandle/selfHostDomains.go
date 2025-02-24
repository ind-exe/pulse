package certhandle

import (
	"log"
	"net/url"
	"oob-server/data"
)

func domainUniqueListCreator(domainList *[]string, domain string, seen map[string]bool) {
	if _, ok := seen[domain]; !ok {
		*domainList = append(*domainList, domain)
		seen[domain] = true
	}
}

func returnDomainFromURl(rawUrl string) (string, error) {
	rawUrl = "https://" + rawUrl
	tempUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	return tempUrl.Host, nil
}

func GetAllCertRequiredDomains() []string {
	var allCertrRequiredDomains []string
	var seen = map[string]bool{}

	for key := range data.UrlMap {
		domainName, err := returnDomainFromURl(key)
		if err != nil {
			log.Printf("CertPrep | fail : could not parse the url %s", key)
			continue
		}
		domainUniqueListCreator(&allCertrRequiredDomains, domainName, seen)
	}

	for key := range data.UrlServeMap {
		domainName, err := returnDomainFromURl(key)
		if err != nil {
			log.Printf("CertPrep | fail : could not parse the url %s", key)
			continue
		}
		domainUniqueListCreator(&allCertrRequiredDomains, domainName, seen)
	}
	return allCertrRequiredDomains
	
}