package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	scanner.Scan()
	polymer := scanner.Text()
	pairingRules := make(map[string]rune)
	scanner.Scan()

	for scanner.Scan() {
		pairLine := strings.Split(scanner.Text(), " -> ")
		pairingRules[pairLine[0]] = []rune(pairLine[1])[0]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elementCount := make(map[rune]int)
	for _, elem := range []rune(polymer) {
		if _, found := elementCount[elem]; !found {
			elementCount[elem] = strings.Count(polymer, string(elem))
		}
	}
	matchingPairsCount := make(map[string]int)
	for pair := range pairingRules {
		matchingPairsCount[pair] = countOverlapping(polymer, pair)
	}

	//Puzzle one and two
	for i := 0; i < 40; i++ {
		matchingPairsCount = pairingStep(matchingPairsCount, &elementCount, pairingRules)
	}
	max, min := findMaxMin(elementCount)
	fmt.Printf("max - min = %d\n", max-min)
}

func pairingStep(matchingPairsCount map[string]int, elemCount *map[rune]int, pairingRules map[string]rune) (newMatchingPairsCount map[string]int) {
	newMatchingPairsCount = make(map[string]int)
	for rule, elem := range pairingRules {
		matches := matchingPairsCount[rule]
		(*elemCount)[elem] += matches
		newMatchingPairsCount[rule[0:1]+string(elem)] += matches
		newMatchingPairsCount[string(elem)+rule[1:]] += matches
	}
	return newMatchingPairsCount
}

func findMaxMin(elemCount map[rune]int) (max int, min int) {
	min = math.MaxInt64
	for _, value := range elemCount {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return max, min
}

// FFS
func countOverlapping(str string, substr string) (count int) {
	for i := range str {
		if strings.HasPrefix(str[i:], substr) {
			count++
		}
	}
	return count
}
