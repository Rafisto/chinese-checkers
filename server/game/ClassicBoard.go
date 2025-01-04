package game

type ClassicBoard struct {
	playerNum int
	board     [][]int
	pawns     *ClassicPawns
}

func NewClassicBoard(playerNum int) (*ClassicBoard, error) {
	classicPawns, err := NewClassicPawns(playerNum)

	if err != nil {
		return nil, err
	}

	board := [][]int{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, 1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{6, -1, 6, -1, 6, -1, 6, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 3, -1, 3, -1, 3, -1, 3},
		{-1, 6, -1, 6, -1, 6, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 3, -1, 3, -1, 3, -1},
		{-1, -1, 6, -1, 6, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 3, -1, 3, -1, -1},
		{-1, -1, -1, 6, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 3, -1, -1, -1},
		{-1, -1, -1, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, -1, -1, -1}, // center
		{-1, -1, -1, 4, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 5, -1, -1, -1},
		{-1, -1, 4, -1, 4, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 5, -1, 5, -1, -1},
		{-1, 4, -1, 4, -1, 4, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 5, -1, 5, -1, 5, -1},
		{4, -1, 4, -1, 4, -1, 4, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 5, -1, 5, -1, 5, -1, 5},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, 2, -1, 2, -1, 2, -1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 2, -1, 2, -1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 2, -1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}

	if playerNum == 2 {
		board[4][8] = 1
		board[4][10] = 1
		board[4][12] = 1
		board[4][14] = 1
		board[4][16] = 1
		board[12][8] = 2
		board[12][10] = 2
		board[12][12] = 2
		board[12][14] = 2
		board[12][16] = 2
	}

	if playerNum == 3 {
		board[4][18] = 4
		board[4][20] = 4
		board[4][22] = 4
		board[4][24] = 4
		board[5][19] = 4
		board[5][21] = 4
		board[5][23] = 4
		board[6][20] = 4
		board[6][22] = 4
		board[7][21] = 4
		board[12][0] = 3
		board[12][2] = 3
		board[12][4] = 3
		board[12][6] = 3
		board[11][1] = 3
		board[11][3] = 3
		board[11][5] = 3
		board[10][2] = 3
		board[10][4] = 3
		board[9][3] = 3
	}

	classicBoard := &ClassicBoard{
		playerNum: playerNum,
		board:     board,
		pawns:     classicPawns,
	}
	return classicBoard, nil
}

func (b *ClassicBoard) Check(x, y int) int {
	if x < 0 || y < 0 || x > 24 || y > 16 {
		return -1
	}
	return b.board[y][x]
}

func (b *ClassicBoard) PrintBoard() {
	for i := 0; i < 17; i++ {
		for j := 0; j < 25; j++ {
			if b.board[i][j] == -1 {
				print("- ")
			} else {
				print(b.board[i][j])
				print(" ")
			}
		}
		print("\n")
	}
}
func (b *ClassicBoard) GetPlayerNum() int {
	return b.playerNum
}

func (b *ClassicBoard) GetBoard() [][]int {
	return b.board
}

func (b *ClassicBoard) GetPawns() Pawns {
	return b.pawns
}
