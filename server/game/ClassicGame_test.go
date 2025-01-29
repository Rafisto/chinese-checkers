package game

import (
	"testing"
)

var game, err = NewClassicGame(0, 2)

func mocknotify(i int, s string) {
	return
}

func TestGameCreation(t *testing.T) {
	game.SetNotify(mocknotify)
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

func TestAddPlayer(t *testing.T) {
	err = game.AddPlayer(0)
	if err != nil {
		t.Fatalf(`game.AddPlayer(0) = %v, want nil`, err)
	}
	err = game.AddPlayer(1)
	if err != nil {
		t.Fatalf(`game.AddPlayer(1) = %v, want nil`, err)
	}
}

func TestAddPlayerError(t *testing.T) {
	err = game.AddPlayer(2)
	if err == nil {
		t.Fatalf(`game.AddPlayer(2) = nil, want error (game full)`)
	}
}

func TestPlayerNumSetError(t *testing.T) {
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
	err = game.SetPlayerNum(3)
	if err != nil {
		t.Fatalf(`game.SetPlayerNum(3) = %v, want nil`, err)
	}
	num := game.GetPlayerNum()
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

	err = game.SetPlayerNum(2)
	if err != nil {
		t.Fatalf(`game.SetPlayerNum(2) = %v, want nil`, err)
	}
	num = game.GetPlayerNum()
	if num != 2 {
		t.Fatalf(`game.GetPlayerNum() = %v, want 2`, num)
	}
}

func TestMoveError(t *testing.T) {
	err = game.Move(0, 0, 0, 0, 0)
	if err == nil {
		t.Fatalf(`expected 'pawn doesn't exist' error, got nil`)
	}

	err = game.Move(1, 11, 3, 10, 4)
	if err == nil {
		t.Fatalf(`expected 'another player's turn' error, got nil`)
	}

	err = game.Move(0, 12, 14, 10, 12)
	if err == nil {
		t.Fatalf(`expected 'invalid pawn' error, got nil`)
	}

	err = game.Move(0, 12, 2, 11, 3)
	if err == nil {
		t.Fatalf(`expected 'space is occupied' error, got nil`)
	}

	err = game.Move(0, 11, 3, 10, 3)
	if err == nil {
		t.Fatalf(`expected 'invalid space' error, got nil`)
	}

	err = game.Move(0, 11, 3, 11, 5)
	if err == nil {
		t.Fatalf(`expected 'invalid move' error, got nil`)
	}
}

func TestMove(t *testing.T) {
	err = game.Move(0, 10, 4, 9, 5)
	if err != nil {
		t.Fatalf(`game.Move() = %v, want nil`, err)
	}

	err = game.Move(1, 11, 13, 9, 11)
	if err != nil {
		t.Fatalf(`game.Move() = %v, want nil`, err)
	}
}
