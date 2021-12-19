package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getNthBit(word, n uint16) uint16 {
	if 1<<n&word > 0 {
		//fmt.Printf("bit %d is set: %05b\n", n, value)
		return 1
	} else {
		//fmt.Printf("bit %d is not set: %05b\n", n, value)
		return 0
	}
}

func mostCommon(nums []uint16) uint16 {
	ones := 0
	for _, n := range nums {
		if n == 1 {
			ones++
		}
	}
	if ones >= len(nums)/2 {
		return 1
	}
	return 0
}

func leastCommon(nums []uint16) uint16 {
	ones := 0
	for _, n := range nums {
		if n == 1 {
			ones++
		}
	}
	if ones < len(nums)/2 {
		return 1
	}
	return 0
}

type bitcriteria func(uint16) bool

func filter(arr []uint16, col []uint16, f bitcriteria) []uint16 {
	result := []uint16{}
	for idx, i := range arr {
		if f(uint16(i)) {
			result = append(result, arr[idx])
		}
	}
	return result
}

func sliceColumn(column uint16, arr []uint16) []uint16 {
	result := []uint16{}
	for _, n := range arr {
		result = append(result, getNthBit(uint16(n), column))
	}
	//for _, n := range result {
	//fmt.Printf("bits: %015b\n", n)
	//}
	return result
}

func findOxygenGeneratorValue(arr []uint16) uint16 {
	firstColumn := sliceColumn(0, arr)
	fmt.Println(firstColumn)
	common := mostCommon(firstColumn)
	fmt.Println(common)
	filtered := filter(arr, firstColumn, func(n uint16) bool {
		return n == common
	})
	fmt.Println(filtered)

	if len(filtered) > 0 {
		return filtered[0]
	} else {
		return 0
	}
}

var length int

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	arr := []uint16{}

	for scanner.Scan() {
		line := scanner.Text()
		length = len(line)
		binary, err := strconv.ParseUint(line, 2, 16)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
		//fmt.Printf("bits: %05b\n", binary)
		arr = append(arr, uint16(binary))
	}

	counts := map[uint16][]uint16{}
	gammaRateArr := []uint16{}
	epsilonRateArr := []uint16{}

	for i := 0; i < length; i++ {
		for _, binary := range arr {
			counts[uint16(i)] = append(counts[uint16(i)], getNthBit(binary, uint16(i)))
		}
		gammaRateArr = append(gammaRateArr, mostCommon(counts[uint16(i)]))
		epsilonRateArr = append(epsilonRateArr, leastCommon(counts[uint16(i)]))
	}
	//fmt.Println(counts)
	//fmt.Println(gammaRateArr)
	//fmt.Println(epsilonRateArr)

	gammaRate := 0.0
	epsilonRate := 0.0

	for i := 0; i < length; i++ {
		//fmt.Printf("gammaRate=%f += 2^%d * %d\n", gammaRate, i, gammaRateArr[i])
		gammaRate += math.Pow(2, float64(i)) * float64(gammaRateArr[i])
		epsilonRate += math.Pow(float64(2), float64(i)) * float64(epsilonRateArr[i])
	}

	oxy := findOxygenGeneratorValue(arr)

	fmt.Println(gammaRate * epsilonRate)
	fmt.Println(oxy)
}
