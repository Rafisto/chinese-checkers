package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"chinese-checkers/game"
)

// WriteJSON godoc
//
//	@Summary Write a JSON response with the provided data and status code
func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// Response godoc
//
//	@Summary	Response message
type Response struct {
	Message string `json:"message"`
}

// WriteJSONMessage godoc
//
// @Summary Write a Json message response in the form of {"message": message} provided a message
func WriteJSONMessage(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{Message: message})
}

// ErrorResponse godoc
//
// @Summary	Error message
type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteJSONError godoc
//
// @Summary Write a JSON error response in the form of {"error": error} provided the error
func WriteJSONError(w http.ResponseWriter, code int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err})
}

// CreateGameRequest godoc
//
//	@Summary	Create game request to send to the CreateGameHandler
type CreateGameRequest struct {
	PlayerNum int `json:"playerNum"`
}

// CreateGameHandler godoc
//
//	@Summary	Create a new game provided its initial parameters
//	@Tags		Game
//	@Accept		json
//	@Produce	json
//	@Param		playerNum	body		CreateGameRequest	true	"Initial game parameters"
//	@Success	201			{object}	Response				"Successfully created game"
//	@Failure	400			{object}	ErrorResponse				"Bad request, missing fields or invalid data"
//	@Router		/games [post].
func (s *Server) CreateGameHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	if r.Method == http.MethodPost {
		var req CreateGameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			WriteJSONError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request: %v", err))
			return
		}

		// TODO: add game type to request
		game, err := gm.CreateGame(req.PlayerNum, 0)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Failed to create game: %v", err))
			return
		}

		WriteJSON(w, http.StatusCreated, map[string]interface{}{"message": "Successfully created game", "id": game.GetID()})
		return
	}

	WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

type GameResponse struct {
	ID             int `json:"id"`
	CurrentPlayers int `json:"currentPlayers"`
	MaxPlayers     int `json:"maxPlayers"`
}

// GetGameHandler godoc
//
//	@Summary	Get a game by its ID
//	@Tags		Game
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Game ID"
//	@Success	200	{object}	GameResponse	"Scuccessfully received the desired game"
//	@Failure	400	{object}	ErrorResponse	"Bad request, missing fields or invalid data"
//	@Router		/games/{id} [get].
func (s *Server) GetGameHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	id := r.PathValue("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, "GameID must be an integer.")
		return
	}

	if r.Method == http.MethodGet {
		game := gm.GetGames()[id_int]
		if game == nil {
			WriteJSONError(w, http.StatusNotFound, fmt.Sprintf("Game with GameID=%d not found", id_int))
			return
		}

		WriteJSON(w, http.StatusOK, GameResponse{
			ID:             game.GetID(),
			CurrentPlayers: game.GetCurrentPlayerNum(),
			MaxPlayers:     game.GetPlayerNum(),
		})

		return
	}

	WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

type GamesResponse []GameResponse

// GetGamesHandler godoc
//
//	@Summary	Get all currently active games
//	@Tags		Game
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	GamesResponse	"Successfully received all active games"
//	@Failure	400	{object}	ErrorResponse	"Bad request, missing fields or invalid data"
//	@Router		/games [get].
func (s *Server) GetGamesHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	if r.Method == http.MethodGet {
		games := gm.GetGames()

		if len(games) == 0 {
			WriteJSON(w, http.StatusOK, []interface{}{})
			return
		}

		gameList := make(GamesResponse, 0)
		for _, game := range games {
			gameList = append(gameList, struct {
				ID             int `json:"id"`
				CurrentPlayers int `json:"currentPlayers"`
				MaxPlayers     int `json:"maxPlayers"`
			}{
				ID:             game.GetID(),
				CurrentPlayers: game.GetCurrentPlayerNum(),
				MaxPlayers:     game.GetPlayerNum(),
			})
		}

		WriteJSON(w, http.StatusOK, gameList)
		return
	}

	WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

type JoinGameRequest struct {
	Username string `json:"username"`
}

type JoinGameResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// JoinGameHandler godoc
//
//	@Summary	Join a game by its ID, provided the username
//	@Tags		Game
//	@Accept		json
//	@Produce	json
//	@Param		game_id		path		string	true	"Game ID"
//	@Param		username	body		JoinGameRequest	true	"Player username"
//	@Success	200			{object}	JoinGameResponse	"Successfully joined the game"
//	@Failure	400			{object}	ErrorResponse	"Bad request, missing fields or invalid data"
//	@Router		/games/{game_id}/join [post].
func (s *Server) JoinGameHandler(w http.ResponseWriter, r *http.Request, gm *game.GameManager) {
	game_id := r.PathValue("game_id")

	game_id_int, err := strconv.Atoi(game_id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, "GameID must be an integer.")
		return
	}

	var req JoinGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSONError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request: %v", err))
		return
	}

	username := req.Username

	if r.Method == http.MethodPost {
		player, err := gm.JoinGame(game_id_int, username)
		if err != nil {
			WriteJSONError(w, http.StatusBadRequest, "Unable to join to the game")
			return
		}

		WriteJSON(w, http.StatusCreated, map[string]interface{}{"message": "Successfully joined the game", "id": player.GetPlayerID()})
		return
	}

	WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
}
