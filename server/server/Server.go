package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "chinese-checkers/docs"
	"chinese-checkers/game"
)

// GameConnection godoc
// @Summary Store a single websocket connection to a game
type GameConnection struct {
	GameID   int
	PlayerID int
	Conn     *websocket.Conn
}

// Server godoc
// @Summary Store game manager and all websocket connections by gameID
type Server struct {
	GameManager     *game.GameManager
	GameConnections map[int][]GameConnection
}

// NewServer godoc
// @Summary Create a new HTTP/WebSocket Server
func NewServer() *Server {
	server := Server{
		GameManager:     game.NewGameManager(),
		GameConnections: make(map[int][]GameConnection),
	}

	server.GameManager.RegisterNotify(func(i int, s string) {
		WBroadcastToGame(i, s, &server)
	})

	return &server
}

// createHandlers godoc
// @Summary Handle all game creation and join endpoints. Handle websocket endpoint
func (s *Server) createHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.CreateGameHandler(w, r, s.GameManager)
		} else if r.Method == http.MethodGet {
			s.GetGamesHandler(w, r, s.GameManager)
		} else {
			WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	mux.HandleFunc("/games/{id}", func(w http.ResponseWriter, r *http.Request) {
		s.GetGameHandler(w, r, s.GameManager)
	})

	mux.HandleFunc("/games/{game_id}/join", func(w http.ResponseWriter, r *http.Request) {
		s.JoinGameHandler(w, r, s.GameManager)
	})

	mux.HandleFunc("/games/{game_id}/bot", func(w http.ResponseWriter, r *http.Request) {
		s.AddBotHandler(w, r, s.GameManager)
	})

	mux.HandleFunc("/games/{game_id}/save", func(w http.ResponseWriter, r *http.Request) {
		s.SaveGameHandler(w, r, s.GameManager)
	})

	mux.HandleFunc("/load/{name}", func(w http.ResponseWriter, r *http.Request) {
		s.LoadGameHandler(w, r, s.GameManager)
	})

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.HandleWebSocket(w, r)
	})

	// mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(swaggerFiles.Handler)))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}

// RegisterNewSocketConnection godoc
// @Summary Add a websocket connection to the server map (by game)
func (s *Server) RegisterNewSocketConnection(gameID, playerID int, conn *websocket.Conn) {
	if _, ok := s.GameConnections[gameID]; !ok {
		s.GameConnections[gameID] = make([]GameConnection, 0)
	}

	s.GameConnections[gameID] = append(s.GameConnections[gameID], GameConnection{
		GameID:   gameID,
		PlayerID: playerID,
		Conn:     conn,
	})
}

// RemoveSocketConnection godoc
// @Summary Remvoe a websocket connection from the server map (by game)
func (s *Server) RemoveSocketConnection(gameID, playerID int) {
	if _, ok := s.GameConnections[gameID]; !ok {
		return
	}

	for i, conn := range s.GameConnections[gameID] {
		if conn.PlayerID == playerID {
			s.GameConnections[gameID] = append(s.GameConnections[gameID][:i], s.GameConnections[gameID][i+1:]...)
			return
		}
	}
}

// RunServer godoc
// @Summary Create handlers and bind HTTP server instance to the desired port
func (s *Server) RunServer(port int) {
	if port > 65535 || port < 1024 {
		log.Printf("[ERROR] Invalid port number: %d", port)
		return
	}

	addr := fmt.Sprintf(":%d", port)

	mux := http.NewServeMux()

	s.createHandlers(mux)

	srv := &http.Server{
		Addr:         addr,
		Handler:      corsMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("[INFO] Server running on port %d", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("[ERROR] Unable to start server: %v", err)
		panic(err)
	}
}
