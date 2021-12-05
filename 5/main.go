package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type segment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	segments := make([]segment, 0)

	for scanner.Scan() {
		segmRE := regexp.MustCompile(`\d+`)
		coords := segmRE.FindAllString(scanner.Text(), -1)
		segm := segment{}
		segm.x1, _ = strconv.Atoi(coords[0])
		segm.y1, _ = strconv.Atoi(coords[1])
		segm.x2, _ = strconv.Atoi(coords[2])
		segm.y2, _ = strconv.Atoi(coords[3])
		segments = append(segments, segm)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	maxX, maxY := findMaxXY(segments)

	ventMap := make([][]int, maxY+1)
	for i := 0; i <= maxY; i++ {
		ventMap[i] = make([]int, maxX+1)
	}

	for _, segm := range segments {
		if segm.x1 == segm.x2 || segm.y1 == segm.y2 {
			for i := 0; i <= max(abs(segm.x2-segm.x1), abs(segm.y2-segm.y1)); i++ {
				ventMap[segm.x1+(i*sign(segm.x2-segm.x1))][segm.y1+(i*sign(segm.y2-segm.y1))]++
			}
		}
	}

	puzzleOnePoints := 0
	for _, row := range ventMap {
		for _, col := range row {
			if col > 1 {
				puzzleOnePoints++
			}
		}
	}

	fmt.Printf("Puzzle one points: %d\n", puzzleOnePoints)

	//Puzzle 2
	secondMap := make([][]int, maxY+1)
	for i := 0; i <= maxY; i++ {
		secondMap[i] = make([]int, maxX+1)
	}

	for _, segm := range segments {
		for i := 0; i <= max(abs(segm.x2-segm.x1), abs(segm.y2-segm.y1)); i++ {
			secondMap[segm.x1+(i*sign(segm.x2-segm.x1))][segm.y1+(i*sign(segm.y2-segm.y1))]++
		}
	}

	puzzleTwoPoints := 0
	for _, row := range secondMap {
		for _, col := range row {
			if col > 1 {
				puzzleTwoPoints++
			}
		}
	}
	fmt.Printf("Puzzle two points: %d\n", puzzleTwoPoints)
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x > 0 {
		return 1
	} else {
		return -1
	}
}

func max(x int, y int) int {
	if y > x {
		return y
	}
	return x
}

func findMaxXY(segments []segment) (maxX int, maxY int) {
	for _, segm := range segments {
		if segm.y1 > maxY {
			maxY = segm.y1
		}
		if segm.y2 > maxY {
			maxY = segm.y2
		}
		if segm.x1 > maxX {
			maxX = segm.x1
		}
		if segm.x2 > maxX {
			maxX = segm.x2
		}
	}
	return maxX, maxY
}
