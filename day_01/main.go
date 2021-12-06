package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	fmt.Println("==== Part 1 ====")

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	previous := -1
	larger := 0

	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panicf("error: %s", err.Error())
		}

		if previous == -1 {
			//fmt.Printf("%d (N/A - no previous measurement)\n", current)
			previous = current
			continue
		}

		isIncreased := current > previous
		if isIncreased {
			larger++
			//fmt.Printf("%d (increased)\n", current)
		} else {
			//fmt.Printf("%d (decreased)\n", current)
		}
		previous = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", larger)
}

func part2() {
	fmt.Println("==== Part 2 ====")

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	size := 3
	numBufs := size + 1
	bufferHolder := make(map[int][]int)
	for i := 0; i < numBufs; i++ {
		bufferHolder[i] = make([]int, 0)
	}

	var sums []int
	lineNumber := 0

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panicf("error: %s", err.Error())
		}
		lineNumber++

		for i := 0; i < numBufs; i++ {
			if lineNumber <= i {
			} else if (lineNumber-i)%(numBufs) == 0 {
				sum := 0
				for _, num := range bufferHolder[i] {
					sum += num
				}
				sums = append(sums, sum)
				bufferHolder[i] = nil
			} else if len(bufferHolder[i]) <= 3 {
				bufferHolder[i] = append(bufferHolder[i], line)
			}
		}
	}
	for i, b := range bufferHolder {
		if len(b) == 3 {
			sum := 0
			for _, num := range bufferHolder[i] {
				sum += num
			}
			sums = append(sums, sum)
			bufferHolder[i] = nil
		}
	}

	lastSum := -1
	var larger int

	for _, sum := range sums {
		if lastSum == -1 {
            //fmt.Printf("%d (N/A - no previous measurement)\n", sum)
			lastSum = sum
			continue
		}

		if sum > lastSum {
			larger++
            //fmt.Printf("%d (increased)\n", sum)
		} else if sum < lastSum {
            //fmt.Printf("%d (decreased)\n", sum)
		} else {
            //fmt.Printf("%d (no change)\n", sum)
		}

		lastSum = sum
	}
	fmt.Printf("%d\n", larger)
}

func main() {
	part1()
	fmt.Println()
	part2()
}
