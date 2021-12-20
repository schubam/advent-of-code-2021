package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type calculation map[int]int

func (c calculation) winner() string {
	if c[1] >= c[0] {
		return "1"
	} else {
		return "0"
	}
}

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

type comparer func(a, b int) bool

func eachLine(cs []calculation, f comparer) int {
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
	return eachLine(cs, func(a, b int) bool {
		return a > b
	})
}

func gammaRate(cs []calculation) int {
	return eachLine(cs, func(a, b int) bool {
		return a < b
	})
}

func DrillDown(lines []string) string {
	result := lines

    wordLen := len(lines[0])
    i := 0
    for {
        if (i >= wordLen) {
            break
        }
        cs := CalculateBits(result)

        var r2 []string
        c := cs[i]

		for _, l := range result {
			if string(l[i]) == c.winner() {
                r2 = append(r2, l)
			}
        }
        fmt.Printf("i: %d, c: %v, winner: %s, results: %v, r2: %v\n", i, c, c.winner(), result, r2)

        result = r2
		if len(result) == 1 {
			break
		}
        i++
    }

	return result[0]
}

type result struct {
	part1 int
	part2 int
}

func Solve(input string) result {
	lines := strings.Split(input, "\n")
	//fmt.Println(lines)
	c := CalculateBits(lines)
	r := result{}
	r.part1 = gammaRate(c) * epsilonRate(c)
	return r
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	content := strings.TrimSpace(string(data))
	fmt.Println(Solve(content))
}
