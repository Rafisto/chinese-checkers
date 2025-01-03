package game

type Board interface {
	PrintBoard()
	Check(x, y int) int
	GetBoard() [][]int
	GetPawns() Pawns
}
