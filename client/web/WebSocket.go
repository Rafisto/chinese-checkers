package web

import (
	"fmt"

	"github.com/gorilla/websocket"
)

const (
	WEBSOCKET_URL = "ws://localhost:8080/ws"
)

type WebSocketConnection struct {
	Conn     *websocket.Conn
	GameID   int
	PlayerID int
}

func NewWebSocketConnection() *WebSocketConnection {
	return &WebSocketConnection{
		Conn:     nil,
		GameID:   -1,
		PlayerID: -1,
	}
}

func (wc *WebSocketConnection) EstablishConnection(gameID, playerID int) error {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("%s?gameID=%d&playerID=%d", WEBSOCKET_URL, gameID, playerID), nil)
	if err != nil {
		fmt.Println("Unable to establish connection:", err)
		return err
	}

	wc.Conn = conn

	return nil
}

func (wc *WebSocketConnection) CloseConnection() {
	wc.Conn.Close()
}

func (wc *WebSocketConnection) EmitMessage(message string) error {
	err := wc.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Println("Unable to emit message:", err)
		return err
	}
	return nil
}

func (wc *WebSocketConnection) ReceiveMessage() (string, error) {
	_, msg, err := wc.Conn.ReadMessage()
	if err != nil {
		return "", err
	}

	return string(msg), nil
}
