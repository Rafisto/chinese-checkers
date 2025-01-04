package game

import "fmt"

type ClassicPawns struct {
	pawns map[Point]int
}

func NewClassicPawns(playerNum int) (*ClassicPawns, error) {
	points := make(map[Point]int)

	for i := 0; i < 17; i++ {
		for j := 0; j < 25; j++ {
			points[Point{x: j, y: i}] = 0
		}
	}

	classicPawns := &ClassicPawns{
		pawns: points,
	}

	switch playerNum {
	case 2:
		classicPawns.fill1()
		classicPawns.fill2()
		classicPawns.fill2Players()
	case 3:
		classicPawns.fill1()
		classicPawns.fill4(2)
		classicPawns.fill5(3)
	case 4:
		classicPawns.fill1()
		classicPawns.fill2()
		classicPawns.fill3()
		classicPawns.fill4(4)
	case 6:
		classicPawns.fill1()
		classicPawns.fill2()
		classicPawns.fill3()
		classicPawns.fill4(4)
		classicPawns.fill5(5)
		classicPawns.fill6()
	default:
		return nil, fmt.Errorf("invalid number of players")
	}

	return classicPawns, nil
}

func (p *ClassicPawns) PrintPawns() {
	for i := 0; i < 17; i++ {
		for j := 0; j < 25; j++ {
			print(p.pawns[Point{x: j, y: i}])
			print(" ")
		}
		print("\n")
	}
}

func (p *ClassicPawns) Check(x, y int) int {
	if x < 0 || y < 0 || x > 24 || y > 16 {
		return -1
	}
	return p.pawns[Point{x: x, y: y}]
}

func (p *ClassicPawns) GetPawns() map[Point]int {
	return p.pawns
}

func (p *ClassicPawns) GetPawnsMatrix() [][]int {
	pawnsArr := make([][]int, 17)
	for i := 0; i < 17; i++ {
		pawnsArr[i] = make([]int, 25)
		for j := 0; j < 25; j++ {
			pawnsArr[i][j] = p.pawns[Point{x: j, y: i}]
		}
	}
	return pawnsArr
}

func (p *ClassicPawns) fill1() {
	p.pawns[Point{x: 12, y: 0}] = 1
	p.pawns[Point{x: 11, y: 1}] = 1
	p.pawns[Point{x: 13, y: 1}] = 1
	p.pawns[Point{x: 10, y: 2}] = 1
	p.pawns[Point{x: 12, y: 2}] = 1
	p.pawns[Point{x: 14, y: 2}] = 1
	p.pawns[Point{x: 9, y: 3}] = 1
	p.pawns[Point{x: 11, y: 3}] = 1
	p.pawns[Point{x: 13, y: 3}] = 1
	p.pawns[Point{x: 15, y: 3}] = 1
}

func (p *ClassicPawns) fill2() {
	p.pawns[Point{x: 12, y: 16}] = 2
	p.pawns[Point{x: 11, y: 15}] = 2
	p.pawns[Point{x: 13, y: 15}] = 2
	p.pawns[Point{x: 10, y: 14}] = 2
	p.pawns[Point{x: 12, y: 14}] = 2
	p.pawns[Point{x: 14, y: 14}] = 2
	p.pawns[Point{x: 9, y: 13}] = 2
	p.pawns[Point{x: 11, y: 13}] = 2
	p.pawns[Point{x: 13, y: 13}] = 2
	p.pawns[Point{x: 15, y: 13}] = 2
}

func (p *ClassicPawns) fill3() {
	p.pawns[Point{x: 18, y: 4}] = 3
	p.pawns[Point{x: 20, y: 4}] = 3
	p.pawns[Point{x: 22, y: 4}] = 3
	p.pawns[Point{x: 24, y: 4}] = 3
	p.pawns[Point{x: 19, y: 5}] = 3
	p.pawns[Point{x: 21, y: 5}] = 3
	p.pawns[Point{x: 23, y: 5}] = 3
	p.pawns[Point{x: 20, y: 6}] = 3
	p.pawns[Point{x: 22, y: 6}] = 3
	p.pawns[Point{x: 21, y: 7}] = 3
}

func (p *ClassicPawns) fill4(n int) {
	p.pawns[Point{x: 3, y: 9}] = n
	p.pawns[Point{x: 2, y: 10}] = n
	p.pawns[Point{x: 4, y: 10}] = n
	p.pawns[Point{x: 1, y: 11}] = n
	p.pawns[Point{x: 3, y: 11}] = n
	p.pawns[Point{x: 5, y: 11}] = n
	p.pawns[Point{x: 0, y: 12}] = n
	p.pawns[Point{x: 2, y: 12}] = n
	p.pawns[Point{x: 4, y: 12}] = n
	p.pawns[Point{x: 6, y: 12}] = n
}

func (p *ClassicPawns) fill5(n int) {
	p.pawns[Point{x: 21, y: 9}] = n
	p.pawns[Point{x: 20, y: 10}] = n
	p.pawns[Point{x: 22, y: 10}] = n
	p.pawns[Point{x: 19, y: 11}] = n
	p.pawns[Point{x: 21, y: 11}] = n
	p.pawns[Point{x: 23, y: 11}] = n
	p.pawns[Point{x: 18, y: 12}] = n
	p.pawns[Point{x: 20, y: 12}] = n
	p.pawns[Point{x: 22, y: 12}] = n
	p.pawns[Point{x: 24, y: 12}] = n
}

func (p *ClassicPawns) fill6() {
	p.pawns[Point{x: 0, y: 4}] = 6
	p.pawns[Point{x: 2, y: 4}] = 6
	p.pawns[Point{x: 4, y: 4}] = 6
	p.pawns[Point{x: 6, y: 4}] = 6
	p.pawns[Point{x: 1, y: 5}] = 6
	p.pawns[Point{x: 3, y: 5}] = 6
	p.pawns[Point{x: 5, y: 5}] = 6
	p.pawns[Point{x: 2, y: 6}] = 6
	p.pawns[Point{x: 4, y: 6}] = 6
	p.pawns[Point{x: 3, y: 7}] = 6
}

func (p *ClassicPawns) fill2Players() {
	p.pawns[Point{x: 8, y: 4}] = 1
	p.pawns[Point{x: 10, y: 4}] = 1
	p.pawns[Point{x: 12, y: 4}] = 1
	p.pawns[Point{x: 14, y: 4}] = 1
	p.pawns[Point{x: 16, y: 4}] = 1
	p.pawns[Point{x: 8, y: 12}] = 2
	p.pawns[Point{x: 10, y: 12}] = 2
	p.pawns[Point{x: 12, y: 12}] = 2
	p.pawns[Point{x: 14, y: 12}] = 2
	p.pawns[Point{x: 16, y: 12}] = 2
}
