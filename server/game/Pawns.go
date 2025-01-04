package game

type Point struct {
	x int
	y int
}

type Pawns interface {
	Move(oldX, oldY, x, y int)
	PrintPawns()
	Check(x, y int) int
	GetPawns() map[Point]int
	GetPawnsMatrix() [][]int
}
