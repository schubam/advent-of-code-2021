package main

import "testing"

func TestSolve(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"

	got := Solve_V1(input)
	want := 37

	if got != want {
		t.Errorf("error, got %d, want %d", got, want)
	}
}
