package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	a, b Point
}

func (l Line) coordinates() []Point {
	var coords []Point
	if l.a.x == l.b.x || l.a.y == l.b.y {
		ax := float64(l.a.x)
		bx := float64(l.b.x)
		ay := float64(l.a.y)
		by := float64(l.b.y)
		deltaX := math.Abs(ax - bx)
		deltaY := math.Abs(ay - by)

		// 0 -> 5
		// 0,1,2,3,4,5 | six elements
		//
		if deltaX > 0 {
			for offset := 0; offset <= int(deltaX); offset++ {
				coords = append(coords, Point{int(math.Min(ax, bx)) + offset, l.a.y})
			}
		}

		if deltaY > 0 {
			for offset := 0; offset <= int(deltaY); offset++ {
				coords = append(coords, Point{l.a.x, int(math.Min(ay, by)) + offset})
			}
		}
	}

	return coords
}

func Solve(input string) int {
	regex := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	var lines []Line
	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			//fmt.Printf("line: %s\n", line)
			matches := regex.FindStringSubmatch(line)
			//fmt.Printf("splits: %v\n", matches)

			x1, _ := strconv.Atoi(matches[1])
			y1, _ := strconv.Atoi(matches[2])
			x2, _ := strconv.Atoi(matches[3])
			y2, _ := strconv.Atoi(matches[4])
			l := Line{Point{x1, y1}, Point{x2, y2}}
			//fmt.Printf("l: %v\n", l)
			lines = append(lines, l)
		}
	}

	overlap := make(map[Point]int)
	for _, line := range lines {
		for _, point := range line.coordinates() {
			overlap[point] += 1
		}
	}
	//fmt.Printf("overlap: %v\n", overlap)

	var count int
	for _, value := range overlap {
		if value > 1 {
			count++
		}
	}
	return count
}

func main() {
	input, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	fmt.Println(Solve(string(input)))
}
