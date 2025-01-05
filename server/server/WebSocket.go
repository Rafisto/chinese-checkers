package server

import (
	"encoding/json"
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

	for {
		msg, err := WReadMessage(conn)
		if err != nil {
			s.RemoveSocketConnection(gameID_int, playerID_int)
			log.Printf("[INFO] Player %d left the game", playerID_int)
			log.Printf("[INFO] Remaining connections: %d", len(s.GameConnections[gameID_int]))
			break
		}

		s.HandlePlayerMessage(conn, gameID_int, playerID_int, msg)

		// WBroadcastToGame(gameID_int, fmt.Sprintf("Player %d: %s", playerID_int, msg), s)
		// WBroadcastToGame(gameID_int, msg, s)
	}
}

// HandlePlayerMessage godoc
// @Summary	Handle player move
// @Description	When player sends a websocket meessage with a defined type and an action it is be parsed, logged and processed further
func (s *Server) HandlePlayerMessage(conn *websocket.Conn, gameID, playerID int, msg string) {
	var playerRequest WSRequest
	if err := json.Unmarshal([]byte(msg), &playerRequest); err != nil {
		log.Printf("[ERROR] Unable to parse player message: %v", err)
		return
	}

	// get state
	if playerRequest.Type == "player" && playerRequest.Action == "state" {
		log.Printf("[INFO] Game %d Player %d requests the game state", gameID, playerID)

		// return whole game state:
		// - current player color
		// - whose turn is it

		requestedGame, ok := s.GameManager.GetGames()[gameID]
		if !ok {
			log.Printf("[ERROR] Game %d does not exist", gameID)
			return
		}

		gameState := map[string]interface{}{
			"type":    "server",
			"action":  "state",
			"color":   slices.Index(requestedGame.GetPlayers(), playerID),
			"players": requestedGame.GetPlayers(),
			"current": requestedGame.GetPlayerTurn(),
			"turn":    requestedGame.GetTurn(),
		}

		gameStateJSON, err := json.Marshal(gameState)
		if err != nil {
			log.Printf("[ERROR] Unable to marshal game state: %v", err)
			return
		}

		log.Printf("[INFO] Sending game state to Game %d Player %d", gameID, playerID)
		WSendMessage(conn, string(gameStateJSON))
	}

	// get board
	if playerRequest.Type == "player" && playerRequest.Action == "board" {
		log.Printf("[INFO] Game %d Player %d requests the board state", gameID, playerID)

		requestedGame, ok := s.GameManager.GetGames()[gameID]
		if !ok {
			log.Printf("[ERROR] Game %d does not exist", gameID)
			return
		}

		board := requestedGame.GetBoard().GetBoard()
		boardJSON, err := json.Marshal(map[string]interface{}{"type": "server", "board": board})
		if err != nil {
			log.Printf("[ERROR] Unable to marshal board: %v", err)
			return
		}

		log.Printf("[INFO] Sending board to Game %d Player %d", gameID, playerID)

		WSendMessage(conn, string(boardJSON))
	}

	// get board state (all pawns)
	if playerRequest.Type == "player" && playerRequest.Action == "pawns" {
		log.Printf("[INFO] Game %d Player %d requests the board pawns", gameID, playerID)

		requestedGame, ok := s.GameManager.GetGames()[gameID]
		if !ok {
			log.Printf("[ERROR] Game %d does not exist", gameID)
			return
		}

		pawns := requestedGame.GetBoard().GetPawns().GetPawnsMatrix()
		pawnsJSON, err := json.Marshal(map[string]interface{}{"type": "server", "pawns": pawns})
		if err != nil {
			log.Printf("[ERROR] Unable to marshal pawns: %v", err)
			return
		}

		log.Printf("[INFO] Sending pawns to Game %d Player %d", gameID, playerID)

		WSendMessage(conn, string(pawnsJSON))
		return
	}

	// move a pawn
	if playerRequest.Type == "player" && playerRequest.Action == "move" {
		if playerRequest.Start.Row == 0 && playerRequest.Start.Col == 0 && playerRequest.End.Row == 0 && playerRequest.End.Col == 0 {
			log.Printf("[INFO] Game %d Player %d requests a skip", gameID, playerID)
			s.GameManager.GetGames()[gameID].SkipTurn(playerID)

			log.Printf("[INFO] Sending game state to Game %d (all players)", gameID)
			errorMessage, _ := json.Marshal(map[string]string{"message": "Skipped Turn"})
			WBroadcastToGame(gameID, string(errorMessage), s)
		}

		log.Printf("[INFO] Game %d Player %d requests a move from (%d, %d) to (%d, %d)", gameID, playerRequest.PlayerID, playerRequest.Start.Col, playerRequest.Start.Row, playerRequest.End.Col, playerRequest.End.Row)
		err := s.GameManager.GetGames()[gameID].Move(playerID, playerRequest.Start.Col, playerRequest.Start.Row, playerRequest.End.Col, playerRequest.End.Row)

		if err != nil {
			log.Printf("[ERROR] Unable to move: %v", err)
			errorMessage, _ := json.Marshal(map[string]string{"message": "Invalid Move"})
			WBroadcastToGame(gameID, string(errorMessage), s)
		} else {
			playerRequest.Type = "server"
			response, err := json.Marshal(playerRequest)
			if err != nil {
				log.Printf("[ERROR] Unable to marshal player request: %v", err)
				return
			}

			WBroadcastToGame(gameID, string(response), s)
		}
		return
	}
}
