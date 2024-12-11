package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: Handlers to communicate with the Server API

// WriteJSON godoc
//
//	@Summary	Write a JSON response with the provided data and status code
func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": data})
}

func (c *Client) ListGamesHandler() ([]byte, error) {
	// TODO change port to variable
	serverPort := 8080
	reqURL := fmt.Sprintf("http://localhost:%d/games", serverPort)
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

func (c *Client) CreateGameHandler(playerNum int) (int, error) {
	jsonBody := &CreateGameRequest{
		PlayerNum: playerNum,
	}

	body, err := json.Marshal(jsonBody)
	if err != nil {
		return -1, fmt.Errorf("failure sending the request: %s", err)
	}

	bodyReader := bytes.NewReader(body)

	serverPort := 8080
	reqURL := fmt.Sprintf("http://localhost:%d/games", serverPort)

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

	serverPort := 8080
	reqURL := fmt.Sprintf("http://localhost:%d/games/%d/join", serverPort, gameID)

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

	fmt.Println(resp.StatusCode)
	fmt.Println(resp)

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
	// TODO change port to variable
	serverPort := 8080
	reqURL := fmt.Sprintf("http://localhost:%d/games/%d", serverPort, gameID)
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
