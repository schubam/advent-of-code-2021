package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
			fmt.Printf("%d (N/A - no previous measurement)\n", current)
			previous = current
			continue
		}

		isIncreased := current > previous
		if isIncreased {
			larger++
			fmt.Printf("%d (increased)\n", current)
		} else {
			fmt.Printf("%d (decreased)\n", current)
		}
		previous = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", larger)
}
