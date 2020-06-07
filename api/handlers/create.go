package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Create ...
func (svc *Service) Create(w http.ResponseWriter, req *http.Request) {
	gameName := mux.Vars(req)["gameName"]

	gameID := svc.Store.Create(gameName)

	game := svc.Store.Get(gameID)

	json.NewEncoder(w).Encode(game)
}
