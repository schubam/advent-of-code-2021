package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := 4512
	got := Solve()[0]

	if want != got {
		t.Errorf("error: want %d, got %d\n", want, got)
	}
}
