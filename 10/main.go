package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	patterns := make([]string, 0)

	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	openings := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	closings := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	puzzleTwoScores := make([]int, 0)

	// Puzzle one and two
	illegalCharPoints := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	completeCharPoints := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	puzzleOneScore := 0
	for _, line := range patterns {
		illegalCharFound := false
		stack := make([]rune, 0)

		for _, char := range line {
			if _, found := openings[char]; found {
				//push
				stack = append(stack, char)
			} else {
				if stack[len(stack)-1] != closings[char] {
					illegalCharFound = true
					puzzleOneScore += illegalCharPoints[char]
					break
				}
				//pop
				stack = stack[0 : len(stack)-1]
			}
		}
		if !illegalCharFound {
			puzzleTwoScore := 0
			for i := len(stack) - 1; i >= 0; i-- {
				puzzleTwoScore = 5*puzzleTwoScore + completeCharPoints[stack[i]]
			}
			puzzleTwoScores = append(puzzleTwoScores, puzzleTwoScore)
		}
	}
	sort.Ints(puzzleTwoScores)

	fmt.Printf("Puzzle one score: %d\n", puzzleOneScore)
	fmt.Printf("Puzzle two score: %d\n", puzzleTwoScores[len(puzzleTwoScores)/2])
}
