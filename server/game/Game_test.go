package game

import "testing"

var game, err = NewGame(0, 6, nil)

func TestGameCreation(t *testing.T) {
	if err != nil {
		t.Fatalf(`NewGame(0, 6, nil) = _, %v, want _, nil`, err)
	}
}

func TestGameIDGet(t *testing.T) {
	id := game.GetID()
	if id != 0 {
		t.Fatalf(`game.GetID() = %v, want 0`, id)
	}
}

func TestBoardGetNil(t *testing.T) {
	board := game.GetBoard()
	if board != nil {
		t.Fatalf(`game.GetBoard() = %v, want nil`, board)
	}
}

func TestBoardSetGet(t *testing.T) {
	board, err := NewClassicBoard(6)
	if err != nil {
		t.Fatalf(`Board creation failed, err: %v`, err)
	}
	err = game.SetBoard(board)
	if err != nil {
		t.Fatalf(`game.SetBoard(board) = %v, want nil`, err)
	}
	newBoard := game.GetBoard()
	if board != newBoard {
		t.Fatalf(`game.GetBoard() = %v, want %v`, newBoard, board)
	}
}

func TestBoardSetInvalid(t *testing.T) {
	board, err := NewClassicBoard(2)
	if err != nil {
		t.Fatalf(`Board creation failed, err: %v`, err)
	}
	err = game.SetBoard(board)
	if err == nil {
		t.Fatalf(`game.SetBoard(board) = nil, expected error`)
	}
}

func TestPlayerNumSetInvalid(t *testing.T) {
	err = game.SetPlayerNum(0)
	if err == nil {
		t.Fatalf(`game.SetPlayerNum(0) = nil, expected error`)
	}
	err = game.SetPlayerNum(1)
	if err == nil {
		t.Fatalf(`game.SetPlayerNum(1) = nil, expected error`)
	}
	err = game.SetPlayerNum(5)
	if err == nil {
		t.Fatalf(`game.SetPlayerNum(5) = nil, expected error`)
	}
	err = game.SetPlayerNum(7)
	if err == nil {
		t.Fatalf(`game.SetPlayerNum(7) = nil, expected error`)
	}
}

func TestPlayerNumSetGet(t *testing.T) {
	err = game.SetPlayerNum(2)
	if err != nil {
		t.Fatalf(`game.SetPlayerNum(2) = %v, want nil`, err)
	}
	num := game.GetPlayerNum()
	if num != 2 {
		t.Fatalf(`game.GetPlayerNum() = %v, want 2`, num)
	}

	err = game.SetPlayerNum(3)
	if err != nil {
		t.Fatalf(`game.SetPlayerNum(3) = %v, want nil`, err)
	}
	num = game.GetPlayerNum()
	if num != 3 {
		t.Fatalf(`game.GetPlayerNum() = %v, want 3`, num)
	}

	err = game.SetPlayerNum(4)
	if err != nil {
		t.Fatalf(`game.SetPlayerNum(4) = %v, want nil`, err)
	}
	num = game.GetPlayerNum()
	if num != 4 {
		t.Fatalf(`game.GetPlayerNum() = %v, want 4`, num)
	}

	err = game.SetPlayerNum(6)
	if err != nil {
		t.Fatalf(`game.SetPlayerNum(6) = %v, want nil`, err)
	}
	num = game.GetPlayerNum()
	if num != 6 {
		t.Fatalf(`game.GetPlayerNum() = %v, want 6`, num)
	}
}
