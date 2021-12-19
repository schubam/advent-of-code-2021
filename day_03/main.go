package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type calculation map[int]int

func CalculateBits(words []string) []calculation {
	l := len(words[0])
	result := []calculation{}

	for pos := 0; pos < l; pos++ {
		c := calculation{0: 0, 1: 0}
		for _, w := range words {
            //fmt.Printf("pos: %d, w: %s, w[pos]: %s, c: %v\n", pos, w, string(w[pos]), c)
			if string(w[pos]) == "1" {
				c[1]++
			} else {
				c[0]++
			}
		}
		result = append(result, c)
	}
	return result
}

type appendOne func(a, b int) bool

func forEach(cs []calculation, f appendOne) int {
	var gamma string

	for _, c := range cs {
		if f(c[0], c[1]) {
			gamma = gamma + "1"
		} else {
			gamma = gamma + "0"
		}
	}
	//fmt.Printf("gamma word: %s\n", gamma)
	i, err := strconv.ParseInt(gamma, 2, 32)
	if err != nil {
		fmt.Println("error: ", err)
	}
	//fmt.Printf("gamma number: %d\n", i)
	return int(i)
}

func epsilonRate(cs []calculation) int {
	return forEach(cs, func(a, b int) bool {
		return a > b
	})
}

func gammaRate(cs []calculation) int {
	return forEach(cs, func(a, b int) bool {
		return a < b
	})
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
    //fmt.Println(lines)
	c := CalculateBits(lines)
	return gammaRate(c) * epsilonRate(c)
}

func main() {
    data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
    content := strings.TrimSpace(string(data))
	fmt.Println(Solve(content))
}
