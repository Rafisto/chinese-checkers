package game

import (
	"chinese-checkers/lib"
	"fmt"
	"slices"
)

type Game struct {
	gameID    int
	playerNum int
	players   []int
	board     Board
	turn      int
}

func NewGame(gameID, playerNum int, board Board) (*Game, error) {
	if slices.Contains([]int{2, 3, 4, 6}, playerNum) {
		game := &Game{
			gameID:    gameID,
			playerNum: playerNum,
			board:     board,
			turn:      0,
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

func (g *Game) SetBoard(board Board) error {
	if board.GetPlayerNum() != g.playerNum {
		return fmt.Errorf("expected board for %v player, got one for %v", g.playerNum, board.GetPlayerNum())
	}
	g.board = board
	return nil
}

func (g *Game) AddPlayer(playerID int) error {
	if !slices.Contains(g.players, playerID) {
		if len(g.players) < g.playerNum {
			g.players = append(g.players, playerID)
			return nil
		} else {
			return fmt.Errorf("lobby full")
		}
	} else {
		return fmt.Errorf("player is already in this game")
	}
}

func (g *Game) GetID() int {
	return g.gameID
}

func (g *Game) GetBoard() Board {
	return g.board
}

func (g *Game) GetPlayers() []int {
	return g.players
}

func (g *Game) GetPlayerNum() int {
	return g.playerNum
}

func (g *Game) GetCurrentPlayerNum() int {
	return len(g.players)
}

func (g *Game) GetTurn() int {
	return g.turn
}

func (g *Game) GetPlayerTurn() int {
	return g.players[g.turn%g.playerNum]
}

func (g *Game) nextTurn() {
	g.turn = (g.turn + 1) % g.playerNum
}

func (g *Game) Step(oldX, oldY, x, y int) bool {

	if oldY == y && lib.Abs(x-oldX) == 2 {
		return true
	}

	if lib.Abs(y-oldY) == 1 && lib.Abs(x-oldX) == 1 {
		return true
	}

	return false
}

func (g *Game) Jump(checked []Point, oldX, oldY, x, y int) bool {
	checked = append(checked, Point{oldX, oldY})
	pawns := g.board.GetPawns()
	board := g.board
	// right
	if !slices.Contains(checked, Point{oldX + 4, oldY}) {
		if pawns.Check(oldX+2, oldY) != 0 && pawns.Check(oldX+4, oldY) == 0 && board.Check(oldX+4, oldY) != -1 {
			if oldX+4 == x && oldY == y {
				return true
			}
			if g.Jump(checked, oldX+4, oldY, x, y) {
				return true
			}
		}
	}
	// left
	if !slices.Contains(checked, Point{oldX - 4, oldY}) {
		if pawns.Check(oldX-2, oldY) != 0 && pawns.Check(oldX-4, oldY) == 0 && board.Check(oldX-4, oldY) != -1 {
			if oldX-4 == x && oldY == y {
				return true
			}
			if g.Jump(checked, oldX-4, oldY, x, y) {
				return true
			}
		}
	}
	// top-left
	if !slices.Contains(checked, Point{oldX - 2, oldY - 2}) {
		if pawns.Check(oldX-1, oldY-1) != 0 && pawns.Check(oldX-2, oldY-2) == 0 && board.Check(oldX-2, oldY-2) != -1 {
			if oldX-2 == x && oldY-2 == y {
				return true
			}
			if g.Jump(checked, oldX-2, oldY-2, x, y) {
				return true
			}
		}
	}
	// bottom-left
	if !slices.Contains(checked, Point{oldX - 2, oldY + 2}) {
		if pawns.Check(oldX-1, oldY+1) != 0 && pawns.Check(oldX-2, oldY+2) == 0 && board.Check(oldX-2, oldY+2) != -1 {
			if oldX-2 == x && oldY+2 == y {
				return true
			}
			if g.Jump(checked, oldX-2, oldY+2, x, y) {
				return true
			}
		}
	}
	// top-right
	if !slices.Contains(checked, Point{oldX + 2, oldY - 2}) {
		if pawns.Check(oldX+1, oldY-1) != 0 && pawns.Check(oldX+2, oldY-2) == 0 && board.Check(oldX+2, oldY-2) != -1 {
			if oldX+2 == x && oldY-2 == y {
				return true
			}
			if g.Jump(checked, oldX+2, oldY-2, x, y) {
				return true
			}
		}
	}
	// bottom-right
	if !slices.Contains(checked, Point{oldX + 2, oldY + 2}) {
		if pawns.Check(oldX+1, oldY+1) != 0 && pawns.Check(oldX+2, oldY+2) == 0 && board.Check(oldX+2, oldY+2) != -1 {
			if oldX+2 == x && oldY+2 == y {
				return true
			}
			if g.Jump(checked, oldX+2, oldY+2, x, y) {
				return true
			}
		}
	}

	return false
}

func (g *Game) Move(playerID, oldX, oldY, x, y int) error {
	if playerID != g.players[g.turn%g.playerNum] {
		return fmt.Errorf("another player's turn")
	}

	if g.board.GetPawns().Check(oldX, oldY)-1 != g.turn {
		return fmt.Errorf("invalid pawn")
	}

	if g.board.GetPawns().Check(x, y) != 0 {
		return fmt.Errorf("space is occupied")
	}

	if g.board.GetBoard()[y][x] == -1 {
		return fmt.Errorf("invalid space")
	}

	if g.Step(oldX, oldY, x, y) {
		g.nextTurn()
		return nil
	}

	if g.Jump(nil, oldX, oldY, x, y) {
		g.nextTurn()
		return nil
	}

	return fmt.Errorf("invalid move")
}

func (g *Game) SkipTurn(playerID int) error {
	if g.players[g.turn%g.playerNum] != playerID {
		return fmt.Errorf("another player's turn")
	}
	g.nextTurn()
	return nil
}
