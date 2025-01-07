package game

import "testing"

func TestChaosCreation(t *testing.T) {
	// 2 players
	game, err := NewChaosGame(0, 2)
	if err != nil {
		t.Fatalf(`NewChaosGame(0, 2) = _, %v, want nil`, err)
	}

	pawns := game.GetBoard().GetPawns().GetPawnsMatrix()
	pawnsCount2 := [2]int{0, 0}

	for i := 0; i < len(pawns); i++ {
		for j := 0; j < len(pawns[i]); j++ {
			if pawns[i][j] != 0 {
				pawnsCount2[pawns[i][j]-1] += 1
			}
		}
	}

	for i := 0; i < len(pawnsCount2); i++ {
		if pawnsCount2[i] != 10 {
			t.Fatalf("Not enough pawns created, want 10, got %v", pawnsCount2[i])
		}
	}

	// 3 players
	game, err = NewChaosGame(0, 3)
	if err != nil {
		t.Fatalf(`NewChaosGame(0, 3) = _, %v, want nil`, err)
	}

	pawns = game.GetBoard().GetPawns().GetPawnsMatrix()
	pawnsCount3 := [3]int{0, 0, 0}

	for i := 0; i < len(pawns); i++ {
		for j := 0; j < len(pawns[i]); j++ {
			if pawns[i][j] != 0 {
				pawnsCount3[(pawns[i][j]-1)/2] += 1
			}
		}
	}

	for i := 0; i < len(pawnsCount3); i++ {
		if pawnsCount3[i] != 10 {
			t.Fatalf("Not enough pawns created, want 10, got %v", pawnsCount3[i])
		}
	}

	// 4 players
	game, err = NewChaosGame(0, 4)
	if err != nil {
		t.Fatalf(`NewChaosGame(0, 4) = _, %v, want nil`, err)
	}

	pawns = game.GetBoard().GetPawns().GetPawnsMatrix()
	pawnsCount4 := [4]int{0, 0, 0, 0}

	for i := 0; i < len(pawns); i++ {
		for j := 0; j < len(pawns[i]); j++ {
			if pawns[i][j] != 0 {
				pawnsCount4[pawns[i][j]-1] += 1
			}
		}
	}

	for i := 0; i < len(pawnsCount4); i++ {
		if pawnsCount4[i] != 10 {
			t.Fatalf("Not enough pawns created, want 10, got %v", pawnsCount4[i])
		}
	}

	// 6 players
	game, err = NewChaosGame(0, 6)
	if err != nil {
		t.Fatalf(`NewChaosGame(0, 6) = _, %v, want nil`, err)
	}

	pawns = game.GetBoard().GetPawns().GetPawnsMatrix()
	pawnsCount6 := [6]int{0, 0, 0, 0, 0, 0}

	for i := 0; i < len(pawns); i++ {
		for j := 0; j < len(pawns[i]); j++ {
			if pawns[i][j] != 0 {
				pawnsCount6[pawns[i][j]-1] += 1
			}
		}
	}

	for i := 0; i < len(pawnsCount6); i++ {
		if pawnsCount6[i] != 10 {
			t.Fatalf("Not enough pawns created, want 10, got %v", pawnsCount6[i])
		}
	}
}
