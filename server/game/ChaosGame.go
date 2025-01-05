package game

import (
	"chinese-checkers/lib"
	"fmt"
	"slices"
)

type ChaosGame struct {
	gameID    int
	playerNum int
	players   []int
	board     Board
	turn      int
	progress  []int
	ended     bool
}

func NewChaosGame(gameID, playerNum int) (Game, error) {
	if slices.Contains([]int{2, 3, 4, 6}, playerNum) {
		board, err := NewChaosBoard(playerNum)
		if err != nil {
			return nil, fmt.Errorf("error creating board: %v", err)
		}

		progress := make([]int, playerNum)
		for i := 0; i < playerNum; i++ {
			progress[i] = 0
		}

		game := &ChaosGame{
			gameID:    gameID,
			playerNum: playerNum,
			board:     board,
			turn:      0,
			progress:  progress,
			ended:     false,
		}
		return game, nil
	} else {
		return nil, fmt.Errorf("invalid playerNum, allowed: 2, 3, 4, 6")
	}
}

func (g *ChaosGame) SetPlayerNum(playerNum int) error {
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

func (g *ChaosGame) AddPlayer(playerID int) error {
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

func (g *ChaosGame) GetID() int {
	return g.gameID
}

func (g *ChaosGame) GetBoard() Board {
	return g.board
}

func (g *ChaosGame) GetPlayers() []int {
	return g.players
}

func (g *ChaosGame) GetPlayerNum() int {
	return g.playerNum
}

func (g *ChaosGame) GetCurrentPlayerNum() int {
	return len(g.players)
}

func (g *ChaosGame) GetTurn() int {
	return g.turn
}

func (g *ChaosGame) GetPlayerTurn() int {
	return g.players[g.turn%g.playerNum]
}

func (g *ChaosGame) GetProgress() []int {
	return g.progress
}

func (g *ChaosGame) nextTurn() {
	g.turn = (g.turn + 1) % g.playerNum
}

func (g *ChaosGame) stepCheck(oldX, oldY, x, y int) bool {

	if oldY == y && lib.Abs(x-oldX) == 2 {
		return true
	}

	if lib.Abs(y-oldY) == 1 && lib.Abs(x-oldX) == 1 {
		return true
	}

	return false
}

func (g *ChaosGame) jumpCheck(checked []Point, oldX, oldY, x, y int) bool {
	checked = append(checked, Point{oldX, oldY})
	pawns := g.board.GetPawns()
	board := g.board
	// right
	if !slices.Contains(checked, Point{oldX + 4, oldY}) {
		if pawns.Check(oldX+2, oldY) != 0 && pawns.Check(oldX+4, oldY) == 0 && board.Check(oldX+4, oldY) != -1 {
			if oldX+4 == x && oldY == y {
				return true
			}
			if g.jumpCheck(checked, oldX+4, oldY, x, y) {
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
			if g.jumpCheck(checked, oldX-4, oldY, x, y) {
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
			if g.jumpCheck(checked, oldX-2, oldY-2, x, y) {
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
			if g.jumpCheck(checked, oldX-2, oldY+2, x, y) {
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
			if g.jumpCheck(checked, oldX+2, oldY-2, x, y) {
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
			if g.jumpCheck(checked, oldX+2, oldY+2, x, y) {
				return true
			}
		}
	}

	return false
}

func (g *ChaosGame) validMove(oldX, oldY, x, y int) {
	pawn := g.board.GetPawns().Check(oldX, oldY)
	currentSquare := g.board.Check(oldX, oldY)
	newSquare := g.board.Check(x, y)

	if newSquare == pawn && currentSquare != pawn {
		g.progress[pawn-1] += 1
		fmt.Println(g.progress[pawn-1])
		if g.progress[pawn-1] == 10 {
			g.ended = true
		}
	}

	g.board.GetPawns().Move(oldX, oldY, x, y)
	g.nextTurn()
}

func (g *ChaosGame) Move(playerID, oldX, oldY, x, y int) error {
	if g.ended {
		return fmt.Errorf("game has ended")
	}

	if playerID != g.players[g.turn%g.playerNum] {
		return fmt.Errorf("another player's turn")
	}

	pawn := g.board.GetPawns().Check(oldX, oldY)
	currentSquare := g.board.Check(oldX, oldY)
	newSquare := g.board.Check(x, y)

	if pawn-1 != g.turn {
		if g.playerNum != 3 {
			return fmt.Errorf("invalid pawn")
		}
		if pawn-1 != 2*g.turn {
			return fmt.Errorf("invalid pawn")
		}
	}

	if g.board.GetPawns().Check(x, y) != 0 {
		return fmt.Errorf("space is occupied")
	}

	if newSquare == -1 {
		return fmt.Errorf("invalid space")
	}

	if currentSquare == pawn {
		if newSquare != currentSquare {
			return fmt.Errorf("cannot escape home destination")
		}
	}

	if g.stepCheck(oldX, oldY, x, y) {
		g.validMove(oldX, oldY, x, y)
		return nil
	}

	if g.jumpCheck(nil, oldX, oldY, x, y) {
		g.validMove(oldX, oldY, x, y)
		return nil
	}

	return fmt.Errorf("invalid move")
}

func (g *ChaosGame) SkipTurn(playerID int) error {
	if g.players[g.turn%g.playerNum] != playerID {
		return fmt.Errorf("another player's turn")
	}
	g.nextTurn()
	return nil
}
