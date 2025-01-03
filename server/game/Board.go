package game

type Board interface {
	PrintBoard()
	Check(x, y int) int
	GetPlayerNum() int
	GetBoard() [][]int
	GetPawns() Pawns
}
