package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Websocket ...
func (svc *Service) Websocket(w http.ResponseWriter, req *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	gameID := mux.Vars(req)["gameID"]

	game := svc.Store.Get(gameID)

	err = conn.WriteJSON(game)
	if err != nil {
		log.Println(err)
	}

	msgs := game.Messages()

	for {
		select {
		case <-msgs:
			err := conn.WriteJSON(game)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
