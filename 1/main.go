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
	prevDepth, totalIncreases := 0, -1

	for _, depth := range depths {
		if depth > prevDepth {
			totalIncreases++
		}
		prevDepth = depth
	}

	fmt.Printf("Total Increases problem 1: %v\n", totalIncreases)

	// Puzzle two
	prevDepth, totalIncreases = 0, -1
	for i := range depths[:len(depths)-2] {
		curDepth := depths[i] + depths[i+1] + depths[i+2]
		if curDepth > prevDepth {
			totalIncreases++
		}
		prevDepth = curDepth
	}

	fmt.Printf("Total Increases problem 2: %v\n", totalIncreases)
}
