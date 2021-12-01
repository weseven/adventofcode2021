package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	inputfile, err := os.Open("1_input")

	if err != nil {
		log.Fatal(err)
	}

	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)

	depths := make([]int, 0)

	for scanner.Scan() {
		curDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, curDepth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	totalIncreases := 0
	for i := range depths[:len(depths)-1] {
		if depths[i+1] > depths[i] {
			totalIncreases++
		}
	}
	fmt.Printf("Total Increases problem 1: %v\n", totalIncreases)

	// Puzzle two
	totalIncreases = 0
	for i := range depths[:len(depths)-3] {
		if depths[i+3] > depths[i] {
			totalIncreases++
		}
	}
	fmt.Printf("Total Increases problem 2: %v\n", totalIncreases)
}
