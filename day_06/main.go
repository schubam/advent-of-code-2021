package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Solve_V2(s string) int {
	sequence := strings.Split(strings.TrimSpace(s), ",")

	population := make(map[int]int, 9)
	for _, str := range sequence {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		population[num] += 1
	}

	for tick := 1; tick <= 256; tick++ {
		zeros := population[0]
		for i := 1; i < 9; i++ {
			population[i-1] = population[i]
		}
		population[8] = zeros
		population[6] += zeros
	}

	var num int
	for _, v := range population {
		num += v
	}
	return num
}

func Solve_V1(s string) int {
	sequence := strings.Split(strings.TrimSpace(s), ",")

	population := []int{}
	for _, str := range sequence {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		population = append(population, num)
	}

	//fmt.Printf("Initial state: %v\n", population)

	var toAdd int
	for tick := 1; tick <= 80; tick++ {
		fmt.Printf("\rtick: %d", tick)
		for i, fish := range population {
			switch fish {
			case 1, 2, 3, 4, 5, 6, 7, 8:
				population[i] = fish - 1
			case 0:
				population[i] = 6
				toAdd++
			}
		}
		for i := 0; i < toAdd; i++ {
			population = append(population, 8)
		}
		toAdd = 0
		//fmt.Printf("After %02d days: %v\n", tick, population)
	}

	return len(population)
}

func main() {
	input, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	fmt.Println(Solve_V2(string(input)))
}
