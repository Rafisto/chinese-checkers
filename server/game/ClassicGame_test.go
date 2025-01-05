package game

import "testing"

var game, err = NewClassicGame(0, 6)

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
