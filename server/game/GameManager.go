package game

import (
	"chinese-checkers/save"
	"fmt"
)

type GameManager struct {
	nextGameID   int
	games        map[int]Game
	nextPlayerID int
	players      map[int]*Player
}

func NewGameManager() *GameManager {
	gameManager := &GameManager{
		nextGameID:   0,
		games:        make(map[int]Game),
		nextPlayerID: 0,
		players:      make(map[int]*Player),
	}
	return gameManager
}

func (gm *GameManager) CreateGame(playerNum int, gameType string) (Game, error) {
	game, err := GameTypes[gameType](gm.nextGameID, playerNum)
	if err != nil {
		return nil, err
	}
	gm.games[gm.nextGameID] = game
	gm.nextGameID += 1
	return game, nil
}

func (gm *GameManager) GetGames() map[int]Game {
	return gm.games
}

func (gm *GameManager) createPlayer(username string, gameID int) (*Player, error) {
	player, err := NewPlayer(gm.nextPlayerID, username, gameID)
	if err != nil {
		return nil, err
	}
	gm.players[gm.nextPlayerID] = player
	gm.nextPlayerID += 1
	return player, nil
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

func (gm *GameManager) SaveGame(gameID int, name string) error {
	game := gm.games[gameID]
	if game == nil {
		return fmt.Errorf("game doesn't exist")
	}

	state := save.GameState{
		Turn:      game.GetTurn(),
		Progress:  game.GetProgress(),
		Variant:   game.GetVariant(),
		PlayerNum: game.GetPlayerNum(),
		Ended:     game.GetEnded(),
		Board:     game.GetBoard().GetBoard(),
		Pawns:     game.GetBoard().GetPawns().GetPawnsMatrix(),
	}

	return save.SaveGameState(state, name)
}

func (gm *GameManager) LoadGame(name string) error {
	state, err := save.LoadGameState(name)
	if err != nil {
		return fmt.Errorf("failed to load game state: %w", err)
	}

	game, err := gm.CreateGame(state.PlayerNum, state.Variant)

	game.SetPlayerNum(state.PlayerNum)

	game.SetTurn(state.Turn)
	game.SetProgress(state.Progress)
	game.SetEnded(state.Ended)

	game.GetBoard().SetBoard(state.Board)
	game.GetBoard().GetPawns().SetPawnsMatrix(state.Pawns)

	return nil
}
