package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func render(m map[Point]int) {
	var maxX, maxY, minX, minY int
	for p := range m {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	fmt.Printf("width: %d, height: %d\n", width, height)

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			if math.Mod(float64(i), float64(width)) == 0 {
				fmt.Printf("\n")
			}
			value, ok := m[Point{i, j}]
			if !ok {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", value)
			}
		}
	}
	fmt.Printf("\n")
}

type Point struct {
	x, y int
}

type Line struct {
	a, b Point
}

func (l Line) Coordinates(isV1 bool) []Point {
	var coords []Point
	ax := float64(l.a.x)
	bx := float64(l.b.x)
	ay := float64(l.a.y)
	by := float64(l.b.y)
	deltaX := math.Abs(ax - bx)
	deltaY := math.Abs(ay - by)

	if l.a.x == l.b.x || l.a.y == l.b.y {
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
	} else {
		if !isV1 {
			// diagonal
			var offset float64
			for offset = 0; offset <= deltaY; offset++ {
				signX := (ax - bx) / deltaX * -1
				signY := (ay - by) / deltaY * -1
				coords = append(coords, Point{int(ax + signX*offset), int(ay + signY*offset)})
			}
		}
	}

	return coords
}

func Solve(input string) (int, int) {
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

	overlap_V1 := make(map[Point]int)
	for _, line := range lines {
		for _, point := range line.Coordinates(true) {
			overlap_V1[point] += 1
		}
	}
	//fmt.Printf("overlap: %v\n", overlap_V1)
	//render(overlap_V1)

	overlap_V2 := make(map[Point]int)
	for _, line := range lines {
		for _, point := range line.Coordinates(false) {
			overlap_V2[point] += 1
		}
	}
	//fmt.Printf("overlap: %v\n", overlap_V2)
	render(overlap_V2)

	var count_V1, count_V2 int
	for _, value := range overlap_V1 {
		if value > 1 {
			count_V1++
		}
	}
	for _, value := range overlap_V2 {
		if value > 1 {
			count_V2++
		}
	}
	return count_V1, count_V2
}

func main() {
	input, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	fmt.Println(Solve(string(input)))
}
