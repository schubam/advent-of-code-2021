package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type distanceFunc func(int, int) int

func Solve(input string, f distanceFunc) int {
	sequence := strings.Split(strings.TrimSpace(input), ",")
	seq := []int{}
	for _, str := range sequence {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		seq = append(seq, num)
	}

	distance := make(map[int]int)
	var acc int
	min, max := findMinAndMax(seq)
	//fmt.Printf("min: %d, max: %d\n", min, max)

	for pick := min; pick <= max; pick++ {
		for _, cursor := range seq {
			d := f(pick, cursor)
			acc += int(d)
		}
		distance[pick] = acc
		acc = 0
	}
	//fmt.Printf("distance: %v\n", distance)

	low := math.Inf(1)
	for _, v := range distance {
		fl := float64(v)
		if fl < low {
			low = fl
		}
	}

	return int(low)
}

func findMinAndMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func V1(pick, cursor int) int {
	return int(math.Abs(float64(pick - cursor)))
}

func V2(pick, cursor int) int {
	var dist int
	times := int(math.Abs(float64(pick - cursor)))
	for i := 1; i <= times; i++ {
		dist += i
	}
	return dist
}

func main() {
	input, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	fmt.Println(Solve(string(input), V1))
	fmt.Println(Solve(string(input), V2))
}
