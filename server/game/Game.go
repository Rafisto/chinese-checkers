package game

import (
	"fmt"
	"slices"
)

type Game struct {
	gameID    int
	playerNum int
	players   []string
	board     Board
}

func NewGame(gameID, playerNum int, board Board) (*Game, error) {
	if slices.Contains([]int{2, 3, 4, 6}, playerNum) {
		game := &Game{
			gameID:    gameID,
			playerNum: playerNum,
			board:     board,
		}
		return game, nil
	} else {
		return nil, fmt.Errorf("invalid playerNum, allowed: 2, 3, 4, 6")
	}
}

func (g *Game) SetPlayerNum(playerNum int) error {
	allowed := []int{2, 3, 4, 6}
	if slices.Contains(allowed, playerNum) {
		if playerNum >= len(g.players) {
			g.playerNum = playerNum
			return nil
		} else {
			return fmt.Errorf("can't change number of players, more players already in lobby")
		}
	} else {
		return fmt.Errorf("invalid playerNum, allowed: 2, 3, 4, 6")
	}
}

func (g *Game) SetBoard(board Board) {
	g.board = board
}

func (g *Game) AddPlayer(username string) error {
	if g.playerNum != 0 {
		if !slices.Contains(g.players, username) {
			if len(g.players) < g.playerNum {
				g.players = append(g.players, username)
				return nil
			} else {
				return fmt.Errorf("lobby full")
			}
		} else {
			return fmt.Errorf("username already in use")
		}
	} else {
		return fmt.Errorf("critical error: Lobby not fully initialized")
	}
}

func (g *Game) GetID() int {
	return g.gameID
}

func (g *Game) GetBoard() Board {
	return g.board
}

func (g *Game) GetPlayers() []string {
	return g.players
}

func (g *Game) GetPlayerNum() int {
	return g.playerNum
}

func (g *Game) GetCurrentPlayerNum() int {
	return len(g.players)
}
