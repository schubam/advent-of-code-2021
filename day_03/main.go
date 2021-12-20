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

func (c calculation) loser() string {
	if c[0] <= c[1] {
		return "0"
	} else {
		return "1"
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

type WinnerFunc func(pos int, c calculation, collection []string) []string



func DrillDown(lines []string, f WinnerFunc) string {
	result := lines

    wordLen := len(lines[0])
    i := 0
    for {
        if (i >= wordLen) {
            break
        }
        cs := CalculateBits(result)

        c := cs[i]

        result = f(i, c, result)
        fmt.Printf("i: %d, c: %v, results: %v\n", i, c, result)

		if len(result) == 1 {
			break
		}
        i++
    }

	return result[0]
}

func Oxy(pos int, c calculation, collection []string) []string{
        var result []string
		for _, l := range collection {
			if string(l[pos]) == c.winner() {
                result = append(result, l)
			}
        }
        return result
    }

func Co2(pos int, c calculation, collection []string) []string{
        var result []string
		for _, l := range collection {
			if string(l[pos]) == c.loser(){
                result = append(result, l)
			}
        }
        return result
    }

func oxygenGeneratorRating(lines []string) int {
    result := DrillDown(lines, Oxy)
	i, err := strconv.ParseInt(result, 2, 32)
	if err != nil {
		fmt.Println("error: ", err)
	}
    return int(i)
}

func co2Rating(lines []string) int {
    result := DrillDown(lines, Co2)
	i, err := strconv.ParseInt(result, 2, 32)
	if err != nil {
		fmt.Println("error: ", err)
	}
    fmt.Println(i)
    return int(i)
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
    r.part2 = oxygenGeneratorRating(lines) * co2Rating(lines)
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
