package main

import (
	"testing"
)

func TestSolve_V1(t *testing.T) {
	input := "3,4,3,1,2"

	got := Solve_V1(input)
	want := 5934

	if got != want {
		t.Errorf("error, got %d, want %d", got, want)
	}
}
