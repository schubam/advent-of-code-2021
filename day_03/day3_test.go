package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	want := 198
	got := Solve(input).part1
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCalculate(t *testing.T) {
	input := strings.Split(strings.TrimSpace(`
0
1
0`), "\n")
    got := CalculateBits(input)
    want := map[int]int{0:2, 1:1}
	if got[0][0] != want[0] || got[0][1] != want[1] {
		t.Errorf("got %d want %d", got, want)
	}
}

func stringSlicesEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestDrillDown(t *testing.T) {
	input := strings.Split(strings.TrimSpace(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`), "\n")

    got := DrillDown(input, Oxy)
    want := "10111"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
