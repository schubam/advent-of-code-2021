package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

	fmt.Println(Solve_V1(string(input)))
}
