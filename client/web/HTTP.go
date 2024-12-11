package web

import (
	"bytes"
	"chinese-checkers-client/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ListGamesResponse struct {
	GameID         int `json:"id"`
	CurrentPlayers int `json:"currentPlayers"`
	MaxPlayers     int `json:"maxPlayers"`
}

type CreateGameResponse struct {
	GameID  int    `json:"id"`
	Message string `json:"message"`
}

type CreateGameRequest struct {
	PlayerNum int `json:"playerNum"`
}

type JoinGameResponse struct {
	PlayerID int `json:"id"`
}

type JoinGameRequest struct {
	Username string `json:"username"`
	GameID   int    `json:"game_ud"`
}

type ShowGameRequest struct {
	GameID int `json:"id"`
}

func (c *Client) ListGamesHandler() ([]*ListGamesResponse, error) {
	reqURL := config.GetConfig().GetURL() + "/games"
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	var body []*ListGamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to decode request: %s", err)
	}
	// err = json.Unmarshal(resp.Body, games)
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	return nil, err
	// }

	return body, nil
}

func (c *Client) CreateGameHandler(playerNum int) (int, error) {
	jsonBody := &CreateGameRequest{
		PlayerNum: playerNum,
	}

	body, err := json.Marshal(jsonBody)
	if err != nil {
		return -1, fmt.Errorf("failure sending the request: %s", err)
	}

	bodyReader := bytes.NewReader(body)

	reqURL := config.GetConfig().GetURL() + "/games"

	req, err := http.NewRequest(http.MethodPost, reqURL, bodyReader)

	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	if resp.StatusCode != 201 {
		return -1, fmt.Errorf("invalid response")
	}

	var respStruct CreateGameResponse

	if err := json.NewDecoder(resp.Body).Decode(&respStruct); err != nil {
		return -1, fmt.Errorf("failed to decode request: %s", err)
	}

	return respStruct.GameID, nil
}

func (c *Client) JoinGameHandler(username string, gameID int) (int, error) {
	jsonBody := &JoinGameRequest{
		Username: username,
		GameID:   gameID,
	}

	body, err := json.Marshal(jsonBody)

	if err != nil {
		return -1, fmt.Errorf("failure sending the request: %s", err)
	}

	bodyReader := bytes.NewReader(body)

	reqURL := config.GetConfig().GetURL() + fmt.Sprintf("/games/%d/join", gameID)

	req, err := http.NewRequest(http.MethodPost, reqURL, bodyReader)

	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	if resp.StatusCode != 201 {
		return -1, fmt.Errorf("invalid response")
	}

	var respStruct JoinGameResponse

	if err := json.NewDecoder(resp.Body).Decode(&respStruct); err != nil {
		return -1, fmt.Errorf("failed to decode request: %s", err)
	}

	return respStruct.PlayerID, nil
}

func (c *Client) ShowGamesHandler(gameID int) ([]byte, error) {
	reqURL := config.GetConfig().GetURL() + fmt.Sprintf("/games/%d", gameID)
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
