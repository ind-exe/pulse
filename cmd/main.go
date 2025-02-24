package main

import (
	"sync"

	dataconfig "github.com/ind-exe/pulse/dataConfig"
	envvar "github.com/ind-exe/pulse/envVar"
	dnsserver "github.com/ind-exe/pulse/oob-server/dns/servers"
	httpserver "github.com/ind-exe/pulse/oob-server/http/servers"
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