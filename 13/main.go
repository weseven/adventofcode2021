package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct{ x, y int }
type FoldInstruction struct {
	direction string
	line      int
}

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	dots := make([]Point, 0)
	foldInstructions := make([]FoldInstruction, 0)
	maxX, maxY := 0, 0

	scanner := bufio.NewScanner(inputfile)
	readAllPoints := false
	pointsRE := regexp.MustCompile(`\d+`)
	instructionsRE := regexp.MustCompile(`[xy]=\d+`)
	for scanner.Scan() {
		if scanner.Text() == "" {
			readAllPoints = true
			continue
		}
		if !readAllPoints {
			pts := pointsRE.FindAllString(scanner.Text(), -1)
			x, _ := strconv.Atoi(pts[0])
			y, _ := strconv.Atoi(pts[1])
			dots = append(dots, Point{x: x, y: y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
		if readAllPoints {
			instr := strings.Split(instructionsRE.FindString(scanner.Text()), "=")
			line, _ := strconv.Atoi(instr[1])
			foldInstructions = append(foldInstructions, FoldInstruction{direction: instr[0], line: line})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	markedPoints := make([][]bool, 0)
	for rowIdx := 0; rowIdx < (maxY + 1); rowIdx++ {
		row := make([]bool, (maxX + 1))
		markedPoints = append(markedPoints, row)
	}
	for _, pt := range dots {
		markedPoints[pt.y][pt.x] = true
	}

	// Puzzle one
	/* fold := foldInstructions[0]
	markedPoints = makeFold(markedPoints,fold)
	visibleDots := 0
	for i,row:= range markedPoints{
		for j := range row {
			if markedPoints[i][j] {
				visibleDots++
			}
		}
	}
	fmt.Printf("Visible dots after the first fold: %d\n",visibleDots) */

	//Puzzle two
	for _, fold := range foldInstructions {
		markedPoints = makeFold(markedPoints, fold)
	}
	for i, row := range markedPoints {
		for j := range row {
			if markedPoints[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func makeFold(markedPoints [][]bool, instruction FoldInstruction) (newMarkedPoints [][]bool) {
	if instruction.direction == "x" {
		for _, row := range markedPoints {
			firstHalf := row[0:instruction.line]
			secondHalf := row[instruction.line+1:]
			for i := range secondHalf {
				firstHalf[i] = (firstHalf[i] || secondHalf[(len(secondHalf)-i)-1])
			}
			newMarkedPoints = append(newMarkedPoints, firstHalf)
		}
	}
	if instruction.direction == "y" {
		firstHalf := markedPoints[0:instruction.line]
		secondHalf := markedPoints[instruction.line+1:]
		for i, row := range secondHalf {
			for j := range row {
				firstHalf[i][j] = (firstHalf[i][j] || secondHalf[(len(secondHalf)-i)-1][j])
			}
		}
		newMarkedPoints = firstHalf
	}
	return newMarkedPoints
}
