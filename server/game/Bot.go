package game

import (
	"math/rand"
	"slices"
)

type Bot struct {
	botID int
	color int
	moves [][2]Point
	board Board
}

func NewBot(botID, color int, board Board) *Bot {
	bot := &Bot{
		botID: botID,
		color: color,
		moves: nil,
		board: board,
	}

	return bot
}

func (b *Bot) GetBotID() int {
	return b.botID
}

func (b *Bot) jumpCheck(point1 Point, x, y int) {
	pawns := b.board.GetPawns()

	point2 := Point{x: x, y: y}
	if slices.Contains(b.moves, [2]Point{point1, point2}) {
		return
	}

	b.moves = append(b.moves, [2]Point{point1, point2})

	if pawns.Check(x-2, y) != 0 && pawns.Check(x-4, y) == 0 && b.board.Check(x-4, y) >= 0 {
		b.jumpCheck(point1, x-4, y)
	}

	if pawns.Check(x+2, y) != 0 && pawns.Check(x+4, y) == 0 && b.board.Check(x+4, y) >= 0 {
		b.jumpCheck(point1, x+4, y)
	}

	if pawns.Check(x-1, y-1) != 0 && pawns.Check(x-2, y-2) == 0 && b.board.Check(x-2, y-2) >= 0 {
		b.jumpCheck(point1, x-2, y-2)
	}

	if pawns.Check(x-1, y+1) != 0 && pawns.Check(x-2, y+2) == 0 && b.board.Check(x-2, y+2) >= 0 {
		b.jumpCheck(point1, x-2, y+2)
	}

	if pawns.Check(x+1, y-1) != 0 && pawns.Check(x+2, y-2) == 0 && b.board.Check(x+2, y-2) >= 0 {
		b.jumpCheck(point1, x+2, y-2)
	}

	if pawns.Check(x+1, y+1) != 0 && pawns.Check(x+2, y+2) == 0 && b.board.Check(x+2, y+2) >= 0 {
		b.jumpCheck(point1, x+2, y+2)
	}
}

func (b *Bot) CalculateMoves() error {
	b.moves = nil

	pawns := b.board.GetPawns()

	for y := 0; y < len(pawns.GetPawnsMatrix()); y++ {
		for x := 0; x < len(pawns.GetPawnsMatrix()[y]); x++ {
			pawn := pawns.Check(x, y)
			if pawn != b.color {
				continue
			}

			point1 := Point{x: x, y: y}

			if pawns.Check(x-2, y) == 0 && b.board.Check(x-2, y) >= 0 {
				point2 := Point{x: x - 2, y: y}
				b.moves = append(b.moves, [2]Point{point1, point2})
			}

			if pawns.Check(x+2, y) == 0 && b.board.Check(x+2, y) >= 0 {
				point2 := Point{x: x + 2, y: y}
				b.moves = append(b.moves, [2]Point{point1, point2})
			}

			if pawns.Check(x-1, y-1) == 0 && b.board.Check(x-1, y-1) >= 0 {
				point2 := Point{x: x - 1, y: y - 1}
				b.moves = append(b.moves, [2]Point{point1, point2})
			}

			if pawns.Check(x-1, y+1) == 0 && b.board.Check(x-1, y+1) >= 0 {
				point2 := Point{x: x - 1, y: y + 1}
				b.moves = append(b.moves, [2]Point{point1, point2})
			}

			if pawns.Check(x+1, y-1) == 0 && b.board.Check(x+1, y-1) >= 0 {
				point2 := Point{x: x + 1, y: y - 1}
				b.moves = append(b.moves, [2]Point{point1, point2})
			}

			if pawns.Check(x+1, y+1) == 0 && b.board.Check(x+1, y+1) >= 0 {
				point2 := Point{x: x + 1, y: y + 1}
				b.moves = append(b.moves, [2]Point{point1, point2})
			}

			if pawns.Check(x-2, y) != 0 && pawns.Check(x-4, y) == 0 && b.board.Check(x-4, y) >= 0 {
				b.jumpCheck(point1, x-4, y)
			}

			if pawns.Check(x+2, y) != 0 && pawns.Check(x+4, y) == 0 && b.board.Check(x+4, y) >= 0 {
				b.jumpCheck(point1, x+4, y)
			}

			if pawns.Check(x-1, y-1) != 0 && pawns.Check(x-2, y-2) == 0 && b.board.Check(x-2, y-2) >= 0 {
				b.jumpCheck(point1, x-2, y-2)
			}

			if pawns.Check(x-1, y+1) != 0 && pawns.Check(x-2, y+2) == 0 && b.board.Check(x-2, y+2) >= 0 {
				b.jumpCheck(point1, x-2, y+2)
			}

			if pawns.Check(x+1, y-1) != 0 && pawns.Check(x+2, y-2) == 0 && b.board.Check(x+2, y-2) >= 0 {
				b.jumpCheck(point1, x+2, y-2)
			}

			if pawns.Check(x+1, y+1) != 0 && pawns.Check(x+2, y+2) == 0 && b.board.Check(x+2, y+2) >= 0 {
				b.jumpCheck(point1, x+2, y+2)
			}
		}
	}

	return nil
}

func (b *Bot) Move() (x, y, newx, newy int) {
	b.CalculateMoves()

	length := len(b.moves)

	if length == 0 {
		return 0, 0, 0, 0
	}

	move := b.moves[rand.Intn(length)]

	return move[0].x, move[0].y, move[1].x, move[1].y
}

func (b *Bot) UpdateBoard(board Board) {
	b.board = board
}
