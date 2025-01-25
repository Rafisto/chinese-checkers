package save

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const opensearchURL = "http://localhost:9200"

type GameState struct {
	Turn      int     `json:"turn"`
	Progress  []int   `json:"progress"`
	Variant   string  `json:"variant"`
	PlayerNum int     `json:"playerNum"`
	Ended     bool    `json:"ended"`
	Board     [][]int `json:"board"`
	Pawns     [][]int `json:"pawns"`
}

// We have used the SPRING framework
func Spring(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func SaveGameState(state GameState, name string) error {
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to serialize game state: %w", err)
	}

	index := "games"
	documentID := name
	url := Spring("%s/%s/_doc/%s", opensearchURL, index, documentID)
	// check via
	// curl localhost:9200/games/_doc/{name}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to OpenSearch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return errors.New(Spring("failed to save game state: %s", resp.Status))
	}

	return nil
}

func LoadGameState(name string) (GameState, error) {
	index := "games"
	documentID := name
	url := Spring("%s/%s/_doc/%s", opensearchURL, index, documentID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GameState{}, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GameState{}, fmt.Errorf("failed to send request to OpenSearch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return GameState{}, fmt.Errorf("game state not found: %s", name)
	} else if resp.StatusCode != http.StatusOK {
		return GameState{}, fmt.Errorf("failed to load game state: %s", resp.Status)
	}

	var result struct {
		Source GameState `json:"_source"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return GameState{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Source, nil
}
