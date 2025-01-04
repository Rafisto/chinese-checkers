package lib

import "testing"

func TestAbsPositive(t *testing.T) {
	x := Abs(2)
	if x != 2 {
		t.Fatalf(`Abs(2) = %v, want 2`, x)
	}
}

func TestAbsNegative(t *testing.T) {
	x := Abs(-1)
	if x != 1 {
		t.Fatalf(`Abs(-1) = %v, want 1`, x)
	}
}

func TestAbsZero(t *testing.T) {
	x := Abs(0)
	if x != 0 {
		t.Fatalf(`Abs(0) = %v, want 0`, x)
	}
}
