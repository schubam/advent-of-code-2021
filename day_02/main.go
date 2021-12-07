package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	horizontal int
	depth      int
)

func forward(amount int) {
	horizontal += amount
}

func up(amount int) {
	depth -= amount
}

func down(amount int) {
	depth += amount
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}

		command := arr[0]
		amount, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}

		switch command {
		case "forward":
			forward(amount)
		case "up":
			up(amount)
		case "down":
			down(amount)
		}
	}
	fmt.Printf("horizontal: %d, depth: %d\n", horizontal, depth)
	fmt.Printf("%d\n", horizontal*depth) }
