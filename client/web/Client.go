package web

import (
	"fmt"
)

type Client struct {
	webSocketConnection *WebSocketConnection
	gameID              int
	username            string
}

func NewClient() *Client {
	client := &Client{
		webSocketConnection: NewWebSocketConnection(),
	}
	return client
}

func (c *Client) SetGameID(gameID int) {
	c.gameID = gameID
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c *Client) GetGameID() int {
	return c.gameID
}

func (c *Client) GetUsername() string {
	return c.username
}

func (c *Client) GetSocket() *WebSocketConnection {
	return c.webSocketConnection
}

func (c *Client) ListGames() {

}

func (c *Client) CreateGame(gameID int) (int, error) {
	return c.CreateGameHandler(gameID)
}

func (c *Client) JoinGame(gameID int) (int, error) {
	return c.JoinGameHandler(c.username, gameID)
}

func (c *Client) ChangeUsername(newUsername string) error {
	c.username = newUsername
	return nil
}

func (c *Client) SendServerMessage(message string) error {
	if c.webSocketConnection == nil {
		return fmt.Errorf("you need to join a game first")
	}

	err := c.webSocketConnection.EmitMessage(message)
	if err != nil {
		return fmt.Errorf("failure sending the message: %s", err)
	}
	return nil
}
