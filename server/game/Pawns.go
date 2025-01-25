package game

type Pawns interface {
	Move(oldX, oldY, x, y int)
	PrintPawns()
	Check(x, y int) int
	GetPawns() map[Point]int
	GetPawnsMatrix() [][]int
	SetPawnsMatrix(pawns [][]int)
}
