package handlers

import (
	"net/http"
	"oob-server/data"
	"sync"
)

func Decider(w http.ResponseWriter, r *http.Request) {
	checkString := r.Host + r.URL.Path
	wg := sync.WaitGroup{}
	wg.Add(1)
	go NotifHandler(checkString, r, &wg)
	wg.Wait()
	if value, ok := data.UrlServeMap[checkString]; ok {
		FileServeHandler(w, r, value)
		return
	}
	
	http.NotFound(w,r)
}