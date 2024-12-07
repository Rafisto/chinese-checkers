package game

import "fmt"

type Player struct {
	playerID int
	username string
	gameID   int
}

func NewPlayer(playerID int, username string, gameID int) (*Player, error) {
	if len(username) == 0 {
		return nil, fmt.Errorf("invalid username")
	}

	player := &Player{
		playerID: playerID,
		username: username,
		gameID:   gameID,
	}
	return player, nil
}

func (p *Player) GetPlayerID() int {
	return p.playerID
}

func (p *Player) GetUsername() string {
	return p.username
}

func (p *Player) GetGameID() int {
	return p.gameID
}
