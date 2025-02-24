package routers

import (
	"net/http"

	"github.com/ind-exe/pulse/oob-server/http/handlers"
)

func CreateGeneralRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Decider)
	return mux
}