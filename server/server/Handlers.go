package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"chinese-checkers/game"
)

// WriteJSON godoc
// @Summary Write a JSON response with the provided data and status code
func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": data})
}

// Response godoc
// @Summary Response message
type Response struct {
	message string
}

// CreateGameRequest godoc
// @Summary Create game request to send to the CreateGameHandler
type CreateGameRequest struct {
	PlayerNum int `json:"playerNum"`
}

// CreateGameHandler godoc
// @Summary Create a new game provided its initial parameters
// @Tags Game
// @Accept jsond
// @Produce json
// @Param playerNum body CreateGameRequest true "Initial game parameters"
// @Success 201 {object} string "Successfully created game"
// @Failure 400 {object} string "Bad request, missing fields or invalid data"
// @Router /games [post].
func (s *Server) CreateGameHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	if r.Method == http.MethodPost {
		var req CreateGameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request: %v", err))
			return
		}

		game, err := gm.CreateGame(req.PlayerNum, nil)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to create game: %v", err))
			return
		}

		WriteJSON(w, http.StatusCreated, fmt.Sprintf("Successfully created game with id: %d", game.GetID()))
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
}

// GetGameHandler godoc
// @Summary Get a game by its ID
// @Tags Game
// @Accept json
// @Produce json
// @Param id path string true "Game ID"
// @Success 200 {object} string "Scuccessfully received the desired game"
// @Failure 400 {object} string "Bad request, missing fields or invalid data"
// @Router /games/{id} [get].
func (s *Server) GetGameHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	id := r.PathValue("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid game ID")
		return
	}

	if r.Method == http.MethodGet {
		game := gm.GetGames()[id_int]
		if game == nil {
			WriteJSON(w, http.StatusNotFound, "Game not found")
			return
		}

		WriteJSON(w, http.StatusOK, game)
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
}

// GetGamesHandler godoc
// @Summary Get all currently active games
// @Tags Game
// @Accept json
// @Produce json
// @Success 200 {object} string "Successfully received all active games"
// @Failure 400 {object} string "Bad request, missing fields or invalid data"
// @Router /games [get].
func (s *Server) GetGamesHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	if r.Method == http.MethodGet {
		games := gm.GetGames()

		if len(games) == 0 {
			WriteJSON(w, http.StatusNotFound, "[]")
			return
		}

		gameIDs := []int{}

		for id := range games {
			if games[id].GetCurrentPlayerNum() != games[id].GetPlayerNum() {
				gameIDs = append(gameIDs, id)
			}
		}

		WriteJSON(w, http.StatusOK, gameIDs)
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
}

// JoinGameHandler godoc
// @Summary Join a game by its ID, provided the username
// @Tags Game
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} string "Successfully joined the game"
// @Failure 400 {object} string "Bad request, missing fields or invalid data"
// @Router /games/{game_id}/join [get].
func (s *Server) JoinGameHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	game_id := r.PathValue("game_id")

	game_id_int, err := strconv.Atoi(game_id)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid game ID")
		return
	}

	username := r.URL.Query().Get("username")

	if r.Method == http.MethodPost {
		player, err := gm.JoinGame(game_id_int, username)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, "Unable to join to the game")
			return
		}

		WriteJSON(w, http.StatusCreated, fmt.Sprintf("Successfully joined the game with player_id: %d", player.GetPlayerID()))
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
}
