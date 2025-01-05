package game

import (
	"fmt"
	"math/rand"
)

type ChaosPawns struct {
	pawns map[Point]int
}

func NewChaosPawns(playerNum int) (*ChaosPawns, error) {
	points := make(map[Point]int)

	for i := 0; i < 17; i++ {
		for j := 0; j < 25; j++ {
			points[Point{x: j, y: i}] = 0
		}
	}

	chaosPawns := &ChaosPawns{
		pawns: points,
	}

	centerArr := [][2]int{{8, 4}, {10, 4}, {12, 4}, {14, 4}, {16, 4}, {7, 5}, {9, 5}, {11, 5}, {13, 5}, {15, 5}, {17, 5}, {6, 6}, {8, 6}, {10, 6}, {12, 6}, {14, 6}, {16, 6}, {18, 6}, {5, 7}, {7, 7}, {9, 7}, {11, 7}, {13, 7}, {15, 7}, {17, 7}, {19, 7}, {4, 8}, {6, 8}, {8, 8}, {10, 8}, {12, 8}, {14, 8}, {16, 8}, {18, 8}, {20, 8}, {5, 9}, {7, 9}, {9, 9}, {11, 9}, {13, 9}, {15, 9}, {17, 9}, {19, 9}, {6, 10}, {8, 10}, {10, 10}, {12, 10}, {14, 10}, {16, 10}, {18, 10}, {7, 11}, {9, 11}, {11, 11}, {13, 11}, {15, 11}, {17, 11}, {8, 12}, {10, 12}, {12, 12}, {14, 12}, {16, 12}}

	switch playerNum {
	case 2:
		max := len(centerArr)
		for i := 1; i < 3; i++ {
			for j := 0; j < 10; j++ {
				random := rand.Intn(max)
				point := centerArr[random]

				centerArr, temp := centerArr[:random], centerArr[random+1:]
				centerArr = append(centerArr, temp...)

				if centerArr == nil {
					return nil, fmt.Errorf("failure setting the pawns")
				}

				chaosPawns.pawns[Point{x: point[0], y: point[1]}] = i
				max--
			}
		}
	case 3:
		max := len(centerArr)
		for i := 1; i < 6; i += 2 {
			for j := 0; j < 10; j++ {
				random := rand.Intn(max)
				point := centerArr[random]

				centerArr, temp := centerArr[:random], centerArr[random+1:]
				centerArr = append(centerArr, temp...)

				if centerArr == nil {
					return nil, fmt.Errorf("failure setting the pawns")
				}

				chaosPawns.pawns[Point{x: point[0], y: point[1]}] = i
				max--
			}
		}
	case 4:
		max := len(centerArr)
		for i := 1; i < 5; i++ {
			for j := 0; j < 10; j++ {
				random := rand.Intn(max)
				point := centerArr[random]

				centerArr, temp := centerArr[:random], centerArr[random+1:]
				centerArr = append(centerArr, temp...)

				if centerArr == nil {
					return nil, fmt.Errorf("failure setting the pawns")
				}

				chaosPawns.pawns[Point{x: point[0], y: point[1]}] = i
				max--
			}
		}
	case 6:
		max := len(centerArr)
		for i := 1; i < 7; i++ {
			for j := 0; j < 10; j++ {
				random := rand.Intn(max)
				point := centerArr[random]

				centerArr, temp := centerArr[:random], centerArr[random+1:]
				centerArr = append(centerArr, temp...)

				if centerArr == nil {
					return nil, fmt.Errorf("failure setting the pawns")
				}

				chaosPawns.pawns[Point{x: point[0], y: point[1]}] = i
				max--
			}
		}
	default:
		return nil, fmt.Errorf("invalid number of players")
	}

	return chaosPawns, nil
}

func (p *ChaosPawns) PrintPawns() {
	for i := 0; i < 17; i++ {
		for j := 0; j < 25; j++ {
			print(p.pawns[Point{x: j, y: i}])
			print(" ")
		}
		print("\n")
	}
}

func (p *ChaosPawns) Check(x, y int) int {
	if x < 0 || y < 0 || x > 24 || y > 16 {
		return -1
	}
	return p.pawns[Point{x: x, y: y}]
}

func (p *ChaosPawns) Move(oldX, oldY, x, y int) {
	old := Point{x: oldX, y: oldY}
	new := Point{x: x, y: y}
	pawn := p.pawns[old]
	p.pawns[old] = 0
	p.pawns[new] = pawn
}

func (p *ChaosPawns) GetPawns() map[Point]int {
	return p.pawns
}

func (p *ChaosPawns) GetPawnsMatrix() [][]int {
	pawnsArr := make([][]int, 17)
	for i := 0; i < 17; i++ {
		pawnsArr[i] = make([]int, 25)
		for j := 0; j < 25; j++ {
			pawnsArr[i][j] = p.pawns[Point{x: j, y: i}]
		}
	}
	return pawnsArr
}
