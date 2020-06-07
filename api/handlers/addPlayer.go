package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AddPlayer ...
func (svc *Service) AddPlayer(w http.ResponseWriter, req *http.Request) {
	gameID := mux.Vars(req)["gameID"]
	playerName := mux.Vars(req)["playerName"]

	fmt.Println("getting game")

	game := svc.Store.Get(gameID)

	err := game.AddPlayer(playerName)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	json.NewEncoder(w).Encode(game)
}
