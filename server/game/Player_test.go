package game

import "testing"

func TestNewPlayerError(t *testing.T) {
	_, err := NewPlayer(0, "", 0)
	if err == nil {
		t.Fatalf(`NewPlayer(0, "", 0) = _, nil, want "invalid username"`)
	}
}

func TestNewPlayer(t *testing.T) {
	player, err := NewPlayer(0, "player", 0)
	if err != nil {
		t.Fatalf(`NewPlayer(0, "player", 0) = _, %v, want nil`, err)
	}

	if player.GetPlayerID() != 0 {
		t.Fatalf(`player.GetPlayerID() = %v, want 0`, player.GetPlayerID())
	}

	if player.GetUsername() != "player" {
		t.Fatalf(`player.GetUsername() = %v, want "player"`, player.GetUsername())
	}

	if player.GetGameID() != 0 {
		t.Fatalf(`player.GetGameID() = %v, want 0`, player.GetGameID())
	}

	_, err = NewPlayer(1, "gamer", 0)
	if err != nil {
		t.Fatalf(`Failure creating 2nd player: %v`, err)
	}
}
