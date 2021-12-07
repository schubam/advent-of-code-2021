package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ------- Part 1
type Part1 struct {
	horizontal int
	depth      int
}

func (p *Part1) result() int {
	return p.horizontal * p.depth
}

func (p *Part1) forward(amount int) {
	p.horizontal += amount
}

func (p *Part1) up(amount int) {
	p.depth -= amount
}

func (p *Part1) down(amount int) {
	p.depth += amount
}

// ------- Part 2
type Part2 struct {
	horizontal int
	depth      int
	aim        int
}

func (p *Part2) result() int {
	return p.horizontal * p.depth
}

func (p *Part2) forward(amount int) {
	p.horizontal += amount
	p.depth += p.aim * amount
}

func (p *Part2) up(amount int) {
	p.aim -= amount
}

func (p *Part2) down(amount int) {
	p.aim += amount
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	p1 := Part1{}
	p2 := Part2{}

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
			p1.forward(amount)
			p2.forward(amount)
		case "up":
			p1.up(amount)
			p2.up(amount)
		case "down":
			p1.down(amount)
			p2.down(amount)
		}
	}
	fmt.Println("=== Part 1 ===")
	fmt.Printf("horizontal: %d, depth: %d\n", p1.horizontal, p1.depth)
	fmt.Printf("%d\n", p1.result())
	fmt.Println("")
	fmt.Println("=== Part 2 ===")
	fmt.Printf("horizontal: %d, depth: %d\n", p2.horizontal, p2.depth)
	fmt.Printf("%d\n", p2.result())
}
