package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattrout92/games/api/handlers"
	"github.com/mattrout92/games/api/store"
)

func main() {
	r := mux.NewRouter()

	svc := handlers.Service{
		Store: &store.Memory{},
	}

	r.Methods("OPTIONS", "POST").Path("/games/{gameName}").HandlerFunc(svc.Create)
	r.Methods("OPTIONS", "POST").Path("/games/{gameName}/instances/{gameID}").HandlerFunc(svc.Start)
	r.Methods("OPTIONS", "POST").Path("/games/{gameName}/instances/{gameID}/players/{playerName}").HandlerFunc(svc.AddPlayer)

	r.Methods("OPTIONS", "GET").Path("/games/{gameName}/instances/{gameID}").HandlerFunc(svc.Websocket)

	log.Println("starting games api")

	http.ListenAndServe(":8080", Middleware(r))
}

// Middleware performs CORS checks
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// This needs to be congigurable when we do security for real
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if req.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400") // 1 day
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Cache-Control", "no-cache")
		h.ServeHTTP(w, req)
	})
}
