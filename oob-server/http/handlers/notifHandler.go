package handlers

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ind-exe/pulse/data"
	"github.com/ind-exe/pulse/models"
)

func NotifHandler(checkUrl string, r *http.Request, wg *sync.WaitGroup) {
	if value, ok := data.UrlMap[checkUrl]; ok {
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("NotifHanlder | fail : could not read the request body", err)
			return
		}
		LogData := models.UrlModel{
			Timestamp: time.Now(),
			IP: r.RemoteAddr,
			Port: r.URL.Port(),
			Method: r.Method,
			HostName: r.Host,
			Path: r.URL.Path,
			Headers: r.Header,
			Body: string(requestBody),
		}
		wg.Done()
		defer r.Body.Close()
		value.Decider(models.Notifier(&LogData))
	} else {
		wg.Done()
	}
}

