package routers

import (
	"net/http"
	"oob-server/oob-server/http/handlers"
)

func CreateGeneralRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Decider)
	return mux
}