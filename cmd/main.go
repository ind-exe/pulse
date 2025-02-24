package main

import (
	dataconfig "oob-server/dataConfig"
	envvar "oob-server/envVar"
	dnsserver "oob-server/oob-server/dns/servers"
	httpserver "oob-server/oob-server/http/servers"
	"sync"
)

func main() {
	envvar.LoadConfig()
	err := dataconfig.LoadConfigurations()
	if err != nil {
		panic(err)
	}

	go dnsserver.InitDnsServer()
	go httpserver.InitHttpServer()
	go httpserver.InitHttpsServer()
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}