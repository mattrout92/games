package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start ...
func (svc *Service) Start(w http.ResponseWriter, req *http.Request) {
	gameID := mux.Vars(req)["gameID"]

	game := svc.Store.Get(gameID)

	err := game.Start()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(game)
}
