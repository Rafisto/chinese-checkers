package game

import (
	"chinese-checkers/lib"
	"encoding/json"
	"fmt"
	"log"
	"slices"
)

type ClassicGame struct {
	gameID    int
	playerNum int
	players   []int
	board     *ClassicBoard
	turn      int
	progress  []int
	ended     bool
	bots      map[int]*Bot
	notify    func(int, string)
}

func NewClassicGame(gameID, playerNum int) (Game, error) {
	if slices.Contains([]int{2, 3, 4, 6}, playerNum) {
		board, err := NewClassicBoard(playerNum)
		if err != nil {
			return nil, fmt.Errorf("error creating board: %v", err)
		}

		progress := make([]int, playerNum)
		for i := 0; i < playerNum; i++ {
			progress[i] = 0
		}

		game := &ClassicGame{
			gameID:    gameID,
			playerNum: playerNum,
			board:     board,
			turn:      0,
			progress:  progress,
			ended:     false,
			bots:      make(map[int]*Bot),
		}
		return game, nil
	} else {
		return nil, fmt.Errorf("invalid playerNum, allowed: 2, 3, 4, 6")
	}
}

func (g *ClassicGame) SetPlayerNum(playerNum int) error {
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

func (g *ClassicGame) AddPlayer(playerID int) error {
	if slices.Contains(g.players, playerID) {
		return fmt.Errorf("player is already in this game")
	}

	if len(g.players) >= g.playerNum {
		return fmt.Errorf("lobby full")
	}

	g.players = append(g.players, playerID)

	if len(g.players) == g.playerNum {
		log.Printf("[BOT] (GameID=%d) Notify of full lobby", g.gameID)

		msg := map[string]interface{}{
			"message": "Skipped Turn",
		}

		jsonData, _ := json.Marshal(msg)
		g.notify(g.gameID, string(jsonData))

		if _, ok := g.bots[g.players[g.turn%g.playerNum]]; ok {
			err := g.botMove()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *ClassicGame) GetID() int {
	return g.gameID
}

func (g *ClassicGame) GetVariant() string {
	return "classic"
}

func (g *ClassicGame) GetBoard() Board {
	return g.board
}

func (g *ClassicGame) GetPlayers() []int {
	return g.players
}

func (g *ClassicGame) GetPlayerNum() int {
	return g.playerNum
}

func (g *ClassicGame) GetCurrentPlayerNum() int {
	return len(g.players)
}

func (g *ClassicGame) GetTurn() int {
	return g.turn
}

func (g *ClassicGame) SetTurn(turn int) {
	g.turn = turn
}

func (g *ClassicGame) GetPlayerTurn() int {
	if len(g.players) != g.playerNum {
		return -1
	}
	return g.players[g.turn%g.playerNum]
}

func (g *ClassicGame) GetProgress() []int {
	return g.progress
}

func (g *ClassicGame) SetProgress(progress []int) {
	g.progress = progress
}

func (g *ClassicGame) GetEnded() bool {
	return g.ended
}

func (g *ClassicGame) SetEnded(ended bool) {
	g.ended = ended
}

func (g *ClassicGame) nextTurn() {
	g.turn = (g.turn + 1) % g.playerNum
	if _, ok := g.bots[g.players[g.turn%g.playerNum]]; ok {
		go (func() {
			g.botMove()
		})()
	}
}

func (g *ClassicGame) stepCheck(oldX, oldY, x, y int) bool {

	if oldY == y && lib.Abs(x-oldX) == 2 {
		return true
	}

	if lib.Abs(y-oldY) == 1 && lib.Abs(x-oldX) == 1 {
		return true
	}

	return false
}

func (g *ClassicGame) jumpCheck(checked []Point, oldX, oldY, x, y int) bool {
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

func (g *ClassicGame) validMove(oldX, oldY, x, y int) {
	pawn := g.board.GetPawns().Check(oldX, oldY)
	pawnDestination := pawn + (pawn%2)*2 - 1
	currentSquare := g.board.Check(oldX, oldY)
	newSquare := g.board.Check(x, y)

	if newSquare == pawnDestination && currentSquare != pawnDestination {
		g.progress[pawn-1] += 1
		fmt.Println(g.progress[pawn-1])
		if g.playerNum == 2 {
			if g.progress[pawn-1] == 15 {
				g.ended = true
			}
		} else {
			if g.progress[pawn-1] == 10 {
				g.ended = true
			}
		}
	}

	g.board.GetPawns().Move(oldX, oldY, x, y)
	if !g.ended {
		g.nextTurn()
	}
}

func (g *ClassicGame) Move(playerID, oldX, oldY, x, y int) error {
	if g.ended {
		return fmt.Errorf("game has ended")
	}

	if playerID != g.players[g.turn%g.playerNum] {
		return fmt.Errorf("another player's turn")
	}

	pawn := g.board.GetPawns().Check(oldX, oldY)

	if pawn == 0 {
		return fmt.Errorf("pawn doesn't exist")
	}

	currentSquare := g.board.Check(oldX, oldY)
	newSquare := g.board.Check(x, y)

	if pawn-1 != g.turn && g.playerNum != 3 {
		return fmt.Errorf("invalid pawn")
	}

	if pawn-1 != 2*g.turn && g.playerNum == 3 {
		return fmt.Errorf("invalid pawn")
	}

	if g.board.GetPawns().Check(x, y) != 0 {
		return fmt.Errorf("space is occupied")
	}

	if newSquare == -1 {
		return fmt.Errorf("invalid space")
	}

	if currentSquare == pawn+(pawn%2)*2-1 {
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

func (g *ClassicGame) SkipTurn(playerID int) error {
	if g.players[g.turn%g.playerNum] != playerID {
		return fmt.Errorf("another player's turn")
	}
	g.nextTurn()
	return nil
}

func (g *ClassicGame) botMove() error {
	log.Printf("bot should move")
	turn := g.players[g.turn%g.playerNum]
	bot, ok := g.bots[turn]
	if !ok {
		return fmt.Errorf("it's not a bot's turn")
	}

	bot.UpdateBoard(g.board)

	x, y, newx, newy := bot.Move()

	if x == 0 && y == 0 && newx == 0 && newy == 0 {
		err := g.SkipTurn(bot.GetBotID())
		return err
	}

	err := g.Move(bot.GetBotID(), x, y, newx, newy)
	if err != nil {
		g.SkipTurn(bot.GetBotID())
	}

	move := MoveToJSON(bot.GetBotID(), x, y, newx, newy)
	log.Printf("[BOT] (GameID=%d) Notify of move (%d,%d)->(%d,%d)", g.gameID, x, y, newx, newy)
	g.notify(g.gameID, move)

	return err
}

func (g *ClassicGame) AddBot(botID int) error {
	var color int

	err := g.AddPlayer(botID)
	log.Printf("%v: %v", botID, err)
	if err != nil {
		return err
	}

	for i := 0; i < len(g.players); i++ {
		if g.players[i] == botID {
			color = i + 1
		}
	}

	if g.playerNum == 3 {
		color = 2*color - 1
	}

	bot := &Bot{
		botID: botID,
		color: color,
		board: g.board,
	}

	g.bots[botID] = bot

	if len(g.players) == g.playerNum {
		_, ok := g.bots[g.players[g.turn%g.playerNum]]
		if ok {
			err := g.botMove()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *ClassicGame) GetNotify() func(int, string) {
	return g.notify
}

func (g *ClassicGame) SetNotify(notify func(int, string)) {
	g.notify = notify
}
