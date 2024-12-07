package game

type GameManager struct {
	nextID int
	games  map[int]*Game
}

func NewGameManager() *GameManager {
	gameManager := &GameManager{
		nextID: 0,
		games:  make(map[int]*Game),
	}
	return gameManager
}

func (gm *GameManager) CreateGame(playerNum int, board Board) (*Game, error) {
	game, err := NewGame(gm.nextID, playerNum, board)
	if err == nil {
		gm.games[gm.nextID] = game
		gm.nextID += 1
		return game, nil
	} else {
		return nil, err
	}
}

func (gm *GameManager) GetGames() map[int]*Game {
	return gm.games
}

func (gm *GameManager) JoinGame(gameID int, username string) error {
	err := gm.games[gameID].AddPlayer(username)
	if err == nil {
		return nil
	} else {
		return err
	}
}
