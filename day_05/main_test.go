package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `0,9 -> 5,9
    8,0 -> 0,8
    9,4 -> 3,4
    2,2 -> 2,1
    7,0 -> 7,4
    6,4 -> 2,0
    0,9 -> 2,9
    3,4 -> 1,4
    0,0 -> 8,8
    5,5 -> 8,2`

	got1, got2 := Solve(input)
	want1 := 5
	want2 := 12

	if got1 != want1 {
		t.Errorf("error, got %d, want %d", got1, want1)
	}

	if got2 != want2 {
		t.Errorf("error, got %d, want %d", got2, want2)
	}
}

func TestCoordinates(t *testing.T) {
	line := Line{Point{5, 5}, Point{8, 2}}
	got := line.Coordinates(false)
	want := []Point{{5, 5}, {6, 4}, {7, 3}, {8, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error, got %v, want %v", got, want)
	}
}
