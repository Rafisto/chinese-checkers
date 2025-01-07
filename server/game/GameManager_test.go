package game

import "testing"

var gm = NewGameManager()

func TestCreateGame(t *testing.T) {
	_, err := gm.CreateGame(2, "classic")
	if err != nil {
		t.Fatalf(`gm.CreateGame(2, "classic") = _, %v, want nil`, err)
	}

	_, err = gm.CreateGame(3, "classic")
	if err != nil {
		t.Fatalf(`gm.CreateGame(3, "classic") = _, %v, want nil`, err)
	}

	_, err = gm.CreateGame(4, "classic")
	if err != nil {
		t.Fatalf(`gm.CreateGame(4, "classic") = _, %v, want nil`, err)
	}

	_, err = gm.CreateGame(6, "classic")
	if err != nil {
		t.Fatalf(`gm.CreateGame(6, "classic") = _, %v, want nil`, err)
	}

	_, err = gm.CreateGame(2, "chaos")
	if err != nil {
		t.Fatalf(`gm.CreateGame(2, "chaos") = _, %v, want nil`, err)
	}

	_, err = gm.CreateGame(3, "chaos")
	if err != nil {
		t.Fatalf(`gm.CreateGame(3, "chaos") = _, %v, want nil`, err)
	}

	_, err = gm.CreateGame(4, "chaos")
	if err != nil {
		t.Fatalf(`gm.CreateGame(4, "chaos") = _, %v, want nil`, err)
	}

	_, err = gm.CreateGame(6, "chaos")
	if err != nil {
		t.Fatalf(`gm.CreateGame(6, "chaos") = _, %v, want nil`, err)
	}
}

func TestCreateGameError(t *testing.T) {
	_, err := gm.CreateGame(0, "classic")
	if err == nil {
		t.Fatalf(`gm.CreateGame(0, "classic") = nil, want error`)
	}

	_, err = gm.CreateGame(1, "classic")
	if err == nil {
		t.Fatalf(`gm.CreateGame(1, "classic") = nil, want error`)
	}

	_, err = gm.CreateGame(5, "classic")
	if err == nil {
		t.Fatalf(`gm.CreateGame(5, "classic") = nil, want error`)
	}

	_, err = gm.CreateGame(100, "classic")
	if err == nil {
		t.Fatalf(`gm.CreateGame(100, "classic") = nil, want error`)
	}

	_, err = gm.CreateGame(0, "chaos")
	if err == nil {
		t.Fatalf(`gm.CreateGame(0, "chaos") = nil, want error`)
	}

	_, err = gm.CreateGame(1, "chaos")
	if err == nil {
		t.Fatalf(`gm.CreateGame(1, "chaos") = nil, want error`)
	}

	_, err = gm.CreateGame(5, "chaos")
	if err == nil {
		t.Fatalf(`gm.CreateGame(5, "chaos") = nil, want error`)
	}

	_, err = gm.CreateGame(100, "chaos")
	if err == nil {
		t.Fatalf(`gm.CreateGame(100, "chaos") = nil, want error`)
	}
}

func TestJoinGame(t *testing.T) {
	_, err := gm.JoinGame(0, "player")
	if err != nil {
		t.Fatalf(`gm.JoinGame(0, "player") = _, %v, want nil`, err)
	}

	_, err = gm.JoinGame(0, "player2")
	if err != nil {
		t.Fatalf(`gm.JoinGame(0, "player2") = _, %v, want nil`, err)
	}
}

func TestJoinGameError(t *testing.T) {
	_, err := gm.JoinGame(-1, "player")
	if err == nil {
		t.Fatalf(`gm.JoinGame(-1, "player") = _, nil, want error`)
	}

	_, err = gm.JoinGame(1000, "player")
	if err == nil {
		t.Fatalf(`gm.JoinGame(1000, "player") = _, nil, want error`)
	}

	_, err = gm.JoinGame(1, "")
	if err == nil {
		t.Fatalf(`gm.JoinGame(1, "") = _, nil, want error`)
	}

	_, err = gm.JoinGame(0, "gamer")
	if err == nil {
		t.Fatalf(`gm.JoinGame(0, "gamer") = _, nil, want error`)
	}
}
