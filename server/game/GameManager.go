package game

import (
	"fmt"
)

type GameManager struct {
	nextGameID   int
	games        map[int]*Game
	nextPlayerID int
	players      map[int]*Player
}

func NewGameManager() *GameManager {
	gameManager := &GameManager{
		nextGameID:   0,
		games:        make(map[int]*Game),
		nextPlayerID: 0,
		players:      make(map[int]*Player),
	}
	return gameManager
}

func (gm *GameManager) CreateGame(playerNum int, board Board) (*Game, error) {
	game, err := NewGame(gm.nextGameID, playerNum, board)
	if err == nil {
		gm.games[gm.nextGameID] = game
		gm.nextGameID += 1
		return game, nil
	} else {
		return nil, err
	}
}

func (gm *GameManager) GetGames() map[int]*Game {
	return gm.games
}

func (gm *GameManager) createPlayer(username string, gameID int) (*Player, error) {
	player, err := NewPlayer(gm.nextPlayerID, username, gameID)
	if err == nil {
		gm.players[gm.nextPlayerID] = player
		gm.nextPlayerID += 1
		return player, nil
	} else {
		return nil, err
	}
}

func (gm *GameManager) GetPlayers() map[int]*Player {
	return gm.players
}

func (gm *GameManager) JoinGame(gameID int, username string) (*Player, error) {
	game := gm.games[gameID]
	if game == nil {
		return nil, fmt.Errorf("game doesn't exist")
	}

	if game.GetCurrentPlayerNum() == game.GetPlayerNum() {
		return nil, fmt.Errorf("game full")
	}

	player, err := gm.createPlayer(username, gameID)
	if err != nil {
		return nil, err
	}

	err = game.AddPlayer(player.GetPlayerID())
	if err != nil {
		return nil, err
	}

	return player, nil
}
