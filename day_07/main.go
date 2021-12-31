package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func Solve_V1(input string) int {
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
	for _, pick := range seq {
		for _, cursor := range seq {
			d := math.Abs(float64(pick - cursor))
			acc += int(d)
		}
		distance[pick] = acc
		acc = 0
	}
	fmt.Printf("distance: %v\n", distance)

	min := math.Inf(1)
	for _, v := range distance {
		fl := float64(v)
		if fl < min {
			min = fl
		}
	}

	return int(min)
}

func main() {
	input, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	fmt.Println(Solve_V1(string(input)))
}
