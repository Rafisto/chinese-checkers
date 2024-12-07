package server

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
}

// Send Message to a single Connection (response to a player query)
func WSendMessage(conn *websocket.Conn, msg string) {
	message := Message{Message: msg}

	if err := conn.WriteJSON(message); err != nil {
		log.Printf("[ERROR] Unable to write message: %v", err)
	}
}

// Broadcast Message to all game players (present a server-side action outcome)
func WBroadcastToGame(gameID int, msg string, s *Server) {
	if _, ok := s.GameConnections[gameID]; !ok {
		return
	}

	for _, conn := range s.GameConnections[gameID] {
		WSendMessage(conn.Conn, msg)
	}
}

// Read Message from a single connection (wait for player input)
func WReadMessage(conn *websocket.Conn) (string, error) {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Printf("[ERROR] Unable to read message: %v", err)
	}

	if messageType == websocket.TextMessage {
		log.Printf("[INFO] Received message: %s", string(p))
	}

	return string(p), err
}

// HandleWebSocket godoc
//
//	@Summary	Provided the username and game ID create a websocket connection.
//	@Tags		Game
//	@Accept		json
//	@Produce	json
//	@Param		gameID		query		int		true	"Game ID"
//	@Param		playerID	query		int		true	"Player ID"
//	@Success	200			{object}	string	"Successfully joined the game"
//	@Failure	400			{object}	string	"Bad request, missing fields or invalid data"
//	@Router		/ws [get].
func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR] Unable to upgrade connection: %v", err)
		http.Error(w, "Unable to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	gameID := r.URL.Query().Get("gameID")

	gameID_int, err := strconv.Atoi(gameID)
	if err != nil {
		conn.WriteJSON(map[string]string{"message": "Invalid game ID"})
		return
	}

	playerID := r.URL.Query().Get("playerID")

	playerID_int, err := strconv.Atoi(playerID)
	if err != nil {
		conn.WriteJSON(map[string]string{"message": "Invalid game ID"})
		return
	}

	requestedGame := s.GameManager.GetGames()[gameID_int]
	if requestedGame == nil {
		conn.WriteJSON(map[string]string{"message": "Game not found"})
		return
	}

	player := s.GameManager.GetPlayers()[playerID_int]
	if player == nil {
		conn.WriteJSON(map[string]string{"message": "Player not found"})
		return
	}

	if !slices.Contains(requestedGame.GetPlayers(), player.GetPlayerID()) {
		conn.WriteJSON(map[string]string{"message": "Player not in game"})
		return
	}

	s.RegisterNewSocketConnection(gameID_int, playerID_int, conn)

	WSendMessage(conn, "Hello!")

	for {
		msg, err := WReadMessage(conn)
		if err != nil {
			s.RemoveSocketConnection(gameID_int, playerID_int)
			log.Printf("[INFO] Player %d left the game", playerID_int)
			log.Printf("[INFO] Remaining connections: %d", len(s.GameConnections[gameID_int]))
			break
		}

		WBroadcastToGame(gameID_int, fmt.Sprintf("Player %d: %s", playerID_int, msg), s)
	}
}
