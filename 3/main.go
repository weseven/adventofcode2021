package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	binParams := make([][]byte, 0)

	for scanner.Scan() {
		// instr := strings.Split(scanner.Text(), " ")
		// instr := strings.Fields(scanner.Text())
		binParams = append(binParams, []byte(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	var mostCommonBits, leastCommonBits string

	for col := range binParams[0] {
		most, least := mostAndLeastCommonBits(binParams, col)
		mostCommonBits += string(most)
		leastCommonBits += string(least)
	}

	gamma, _ := strconv.ParseInt(mostCommonBits, 2, 0)
	epsilon, _ := strconv.ParseInt(leastCommonBits, 2, 0)

	fmt.Printf("gamma*epsilon: %d\n", gamma*epsilon)

	// Puzzle two
	oxyNums := make([][]byte, len(binParams))
	co2Nums := make([][]byte, len(binParams))

	copy(oxyNums, binParams)
	copy(co2Nums, binParams)

	for col := range oxyNums[0] {
		mostOxy, _ := mostAndLeastCommonBits(oxyNums,col)
		oxyNums = filterBits(oxyNums, mostOxy, col)
		if len(oxyNums) == 1 {
			break
		}
	}
	for col := range co2Nums[0] {
		_, leastCO2 := mostAndLeastCommonBits(co2Nums,col)
		co2Nums = filterBits(co2Nums, leastCO2, col)
		if len(co2Nums) == 1 {
			break
		}
	}

	oxy, _ := strconv.ParseInt(string(oxyNums[0]), 2, 0)
	co2, _ := strconv.ParseInt(string(co2Nums[0]), 2, 0)

	fmt.Printf("oxy*co2: %d\n", oxy*co2)
}

// find most and least common bits for a given column of a matrix of bytes
func mostAndLeastCommonBits(matrix [][]byte, column int)(most byte, least byte){
	numOnes := 0
		for row := range matrix {
			if matrix[row][column] == '1' {
				numOnes++
			}
		}
		if numOnes >= ((len(matrix)/2)+(len(matrix)%2)) {
		// if numOnes >= (len(matrix)/2) {
			return '1','0'
		} else {
			return '0','1'
		}
}

// filter a list of []bytes by eliminating all elements that do not have the given byte at the given position
func filterBits(list [][]byte, filter byte, position int) (filteredList [][]byte) {
	for _, row := range list {
		if row[position] == filter {
			filteredList = append(filteredList, row)
		}
	}
	return filteredList
}
