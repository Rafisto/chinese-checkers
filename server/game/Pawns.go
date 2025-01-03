package game

type Point struct {
	x int
	y int
}

type Pawns interface {
	PrintPawns()
	Check(x, y int) int
	GetPawns() map[Point]int
}
