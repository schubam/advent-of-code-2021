package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getNthBit(value, n uint16) int {
	if 1<<n&value > 0 {
		//fmt.Printf("bit %d is set: %05b\n", n, value)
		return 1
	} else {
		//fmt.Printf("bit %d is not set: %05b\n", n, value)
		return 0
	}
}

func mostCommon(nums []int) int {
	ones := 0
	for _, n := range nums {
		if n == 1 {
			ones++
		}
	}
	if ones > len(nums)/2 {
		return 1
	}
	return 0
}

func leastCommon(nums []int) int {
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

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	arr := []uint16{}

	for scanner.Scan() {
		binary, err := strconv.ParseInt(scanner.Text(), 2, 16)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
		//fmt.Printf("bits: %05b\n", binary)
		arr = append(arr, uint16(binary))
	}

	counts := map[int][]int{}
	gammaRateArr := []int{}
	epsilonRateArr := []int{}

	for i := 0; i < 12; i++ {
		for _, binary := range arr {
			counts[i] = append(counts[i], getNthBit(binary, uint16(i)))
		}
		gammaRateArr = append(gammaRateArr, mostCommon(counts[i]))
		epsilonRateArr = append(epsilonRateArr, leastCommon(counts[i]))
	}
	//fmt.Println(counts)
	//fmt.Println(gammaRateArr)
	//fmt.Println(epsilonRateArr)

	gammaRate := 0.0
	epsilonRate := 0.0

	for i := 0; i < 12; i++ {
		//fmt.Printf("gammaRate=%f += 2^%d * %d\n", gammaRate, i, gammaRateArr[i])
		gammaRate += math.Pow(2, float64(i)) * float64(gammaRateArr[i])
		epsilonRate += math.Pow(float64(2), float64(i)) * float64(epsilonRateArr[i])
	}

	fmt.Println(gammaRate*epsilonRate)
}
