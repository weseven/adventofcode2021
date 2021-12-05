package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoBoard struct {
	numbers []int
	marks   []bool
}

const bingoRows int = 5
const bingoCols int = 5

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)

	scanner.Scan()
	extractedNumStr := strings.Split(scanner.Text(), ",")
	extractedNumbers := make([]int, 0)
	for _, str := range extractedNumStr {
		num, _ := strconv.Atoi(str)
		extractedNumbers = append(extractedNumbers, num)
	}
	scanner.Scan()
	bingoBoards := make([]bingoBoard, 0)

	for scanner.Scan() {
		//empty row
		scanner.Text()
		board := bingoBoard{make([]int, 0), make([]bool, bingoCols*bingoRows)}
		for i := 0; i < bingoRows; i++ {
			rowStr := strings.Fields(scanner.Text())
			for _, strNum := range rowStr {
				num, _ := strconv.Atoi(strNum)
				board.numbers = append(board.numbers, num)
			}
			scanner.Scan()
		}
		bingoBoards = append(bingoBoards, board)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Puzzle 1
	var firstBoardScore int

firstextraction:
	for _, num := range extractedNumbers {
		for _, b := range bingoBoards {
			if marked, pos := markNumberInBoard(b, num); marked {
				if checkWin(b, pos) {
					firstBoardScore = computeScore(b, pos)
					break firstextraction
				}
			}
		}
	}
	fmt.Printf("First board win score: %d\n", firstBoardScore)

	//Puzzle 2
	hasBoardWon := make([]bool, len(bingoBoards))
	var lastBoardScore int

	//reset board marks...
	for _, b := range bingoBoards {
		b.marks = make([]bool, len(b.numbers))
	}

secondextraction:
	for _, num := range extractedNumbers {
		for i, b := range bingoBoards {
			if !hasBoardWon[i] {
				if marked, pos := markNumberInBoard(b, num); marked {
					if checkWin(b, pos) {
						hasBoardWon[i] = true
						if checkAllTrue(hasBoardWon) {
							lastBoardScore = computeScore(b, pos)
							break secondextraction
						}
					}
				}
			}
		}
	}
	fmt.Printf("Last board win score: %d\n", lastBoardScore)
}

func checkAllTrue(marks []bool) bool {
	for _, mark := range marks {
		if !mark {
			return false
		}
	}
	return true
}

func checkWin(board bingoBoard, lastMarkPos int) bool {
	//check row
	rowMarks := board.marks[((lastMarkPos / bingoCols) * bingoCols):((lastMarkPos/bingoCols)*bingoCols + bingoCols)]
	if checkAllTrue(rowMarks) {
		return true
	}
	//check col
	colMarks := make([]bool, bingoRows)
	for i := 0; i < bingoRows; i++ {
		colMarks[i] = board.marks[((lastMarkPos % bingoCols) + i*bingoCols)]
	}
	if checkAllTrue(colMarks) {
		return true
	}
	return false
}

// return number position and true if a number has been marked on this board
func markNumberInBoard(board bingoBoard, num int) (marked bool, pos int) {
	for pos, n := range board.numbers {
		if num == n {
			board.marks[pos] = true
			return true, pos
		}
	}
	return false, -1
}

func computeScore(board bingoBoard, lastMarkPos int) (score int) {
	score = 0
	for i := range board.marks {
		if !board.marks[i] {
			score += board.numbers[i]
		}
	}
	return score * board.numbers[lastMarkPos]
}
