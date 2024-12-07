package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"chinese-checkers/game"
)

type Server struct {
	GameManager *game.GameManager
}

func NewServer() *Server {
	return &Server{
		GameManager: game.NewGameManager(),
	}
}

func (s *Server) createHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateGameHandler(w, r, s.GameManager)
		} else if r.Method == http.MethodGet {
			GetGamesHandler(w, r, s.GameManager)
		} else {
			WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	mux.HandleFunc("/games/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetGameHandler(w, r, s.GameManager)
	})

	mux.HandleFunc("/games/{game_id}/join", func(w http.ResponseWriter, r *http.Request) {
		JoinGameHandler(w, r, s.GameManager)
	})
}

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
		Handler:      mux,
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
