package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BingoGame struct {
	last     int
	sequence []int
	boards   []Board
	winners  []int
}

func (game *BingoGame) addBoard(b Board) {
	game.boards = append(game.boards, b)
}

func (game *BingoGame) draw() error {
	if game.last >= len(game.sequence) {
		return errors.New("sequence finished, nothing more to draw")
	}
	number := game.sequence[game.last]
	for _, board := range game.boards {
		if !board.hasWon() {
			board.markNumber(number)
			if board.hasWon() {
				game.winners = append(game.winners, board.score(number))
			}
		}
	}
	game.last++
	return nil
}

func (game *BingoGame) play() []int {
	for {
		err := game.draw()
		if err != nil {
			//fmt.Printf("game over: %s\n", err)
			return game.winners
		}
	}
}

type Board struct {
	numbers  []int
	selected map[int]bool // indices of `numbers`
}

func (b *Board) score(e int) int {
	var sum int
	for i := 0; i < 25; i++ {
		if _, ok := b.selected[i]; !ok {
			sum += b.numbers[i]
		}
	}

	//fmt.Printf("score; e: %d, sum: %d, score: %d\n", e, sum, sum*e)
	//fmt.Printf("board: %v\n", b)
	return sum * e
}

func (b *Board) markNumber(drawn int) {
	for i, n := range b.numbers {
		if n == drawn {
			b.selected[i] = true
		}
	}
}

func (b *Board) hasWon() bool {
	for offset := 0; offset <= 20; offset += 5 {
		counter := 0
		for i := range []int{0, 1, 2, 3, 4} {
			val := i + offset
			if _, ok := b.selected[val]; ok {
				counter++
			}
			if counter == 5 {
				return true
			}
		}
	}

	for offset := 0; offset <= 5; offset += 1 {
		counter := 0
		for _, i := range []int{0, 5, 10, 15, 20} {
			val := i + offset
			if _, ok := b.selected[val]; ok {
				counter++
			}
			if counter == 5 {
				return true
			}
		}
	}

	return false
}

func NewBoard(s string) *Board {
	regex := regexp.MustCompile(`\s+`)
	var numbers []int
	splits := regex.Split(s, -1)
	for _, str := range splits {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Sprintf("error: can't convert string to number '%s'", str))
		}
		numbers = append(numbers, n)
	}

	if len(numbers) != 25 {
		panic(fmt.Sprintf("error: board input not 25 (5x5), but %d", len(numbers)))
	}

	b := &Board{numbers: numbers}
	b.selected = make(map[int]bool)
	return b
}

func Solve() []int {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("error: failed to open file")
	}

	game := &BingoGame{}
	var b *Board

	reader := bufio.NewReader(file)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	//fmt.Printf("'%s'\n", text)

	var boardNums string
	for _, n := range strings.Split(text, ",") {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("error: can't convert string to number, %s\n", n)
			break
		}
		game.sequence = append(game.sequence, number)
	}

	for {
		text, err = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		//fmt.Printf("text: '%s'\n", text)

		switch text {
		case "":
			if len(boardNums) > 0 {
				b = NewBoard(boardNums)
				boardNums = ""
				//fmt.Printf("board: %v\n", *b)
				game.addBoard(*b)
			}
		default:
			boardNums += " " + text
			boardNums = strings.TrimSpace(boardNums)
		}

		if err != nil {
			//fmt.Printf("error: %s\n", err)
			break
		}
	}

	//fmt.Printf("game: %v\n", game)
	return game.play()
}
