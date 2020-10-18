package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// NextRound ...
func (svc *Service) NextRound(w http.ResponseWriter, req *http.Request) {
	gameID := mux.Vars(req)["gameID"]
	game := svc.Store.Get(gameID)

	game.NextRound()

	json.NewEncoder(w).Encode(game)
}
