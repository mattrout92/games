package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattrout92/games/api/games"
)

// Turn ...
func (svc *Service) Turn(w http.ResponseWriter, req *http.Request) {
	gameID := mux.Vars(req)["gameID"]
	game := svc.Store.Get(gameID)
	cardDescriptor := mux.Vars(req)["card"]

	card := games.GetCardFromDescriptor(cardDescriptor)
	card.Description = cardDescriptor

	game.Turn(card)
	json.NewEncoder(w).Encode(game)
}
