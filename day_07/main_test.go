package main

import "testing"

func TestSolveV1(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"

	got := Solve(input, V1)
	want := 37

	if got != want {
		t.Errorf("error, got %d, want %d", got, want)
	}
}

func TestSolveV2(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"

	got := Solve(input, V2)
	want := 168

	if got != want {
		t.Errorf("error, got %d, want %d", got, want)
	}
}
