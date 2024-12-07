package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *Server) HandleWS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR] Unable to upgrade connection: %v", err)
		return
	}

	var msg map[string]string
	err = conn.ReadJSON(&msg)
	if err != nil {
		log.Printf("[ERROR] Unable to read message: %v", err)
		return
	}

	conn.WriteJSON(map[string]string{"message": "Hello!"})
	return
}
