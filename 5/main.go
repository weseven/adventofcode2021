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
	maxX := findMaxX(segments)
	maxY := findMaxY(segments)

	// var ventMap [maxX][maxY]int

	ventMap := make([][]int, 0)
	for i := 0; i <= maxY; i++ {
		row := make([]int, 0)
		for j := 0; j <= maxX; j++ {
			row = append(row, 0)
		}
		ventMap = append(ventMap, row)
	}

	for _, segm := range segments {
		if segm.x1 == segm.x2 {
			if segm.y1 > segm.y2 {
				for i := segm.y2; i <= segm.y1; i++ {
					ventMap[i][segm.x1]++
				}
			} else {
				for i := segm.y1; i <= segm.y2; i++ {
					ventMap[i][segm.x1]++
				}
			}
		} else if segm.y1 == segm.y2 {
			if segm.x1 > segm.x2 {
				for i := segm.x2; i <= segm.x1; i++ {
					ventMap[segm.y1][i]++
				}
			} else {
				for i := segm.x1; i <= segm.x2; i++ {
					ventMap[segm.y1][i]++
				}
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
	secondMap := make([][]int, 0)
	for i := 0; i <= maxY; i++ {
		row := make([]int, 0)
		for j := 0; j <= maxX; j++ {
			row = append(row, 0)
		}
		secondMap = append(secondMap, row)

	}
	for _, segm := range segments {
		if segm.x1 == segm.x2 {
			if segm.y1 > segm.y2 {
				for i := segm.y2; i <= segm.y1; i++ {
					secondMap[i][segm.x1]++
				}
			} else {
				for i := segm.y1; i <= segm.y2; i++ {
					secondMap[i][segm.x1]++
				}
			}
		} else if segm.y1 == segm.y2 {
			if segm.x1 > segm.x2 {
				for i := segm.x2; i <= segm.x1; i++ {
					secondMap[segm.y1][i]++
				}
			} else {
				for i := segm.x1; i <= segm.x2; i++ {
					secondMap[segm.y1][i]++
				}
			}
		} else if segm.x1 < segm.x2 {
			if segm.y1 < segm.y2 {
				for i := 0; i <= segm.y2-segm.y1; i++ {
					secondMap[segm.y1+i][segm.x1+i]++
				}
			} else {
				for i := 0; i <= segm.y1-segm.y2; i++ {
					secondMap[segm.y1-i][segm.x1+i]++
				}
			}
		} else if segm.x1 > segm.x2 {
			if segm.y1 < segm.y2 {
				for i := 0; i <= segm.y2-segm.y1; i++ {
					secondMap[segm.y1+i][segm.x1-i]++
				}
			} else {
				for i := 0; i <= segm.y1-segm.y2; i++ {
					secondMap[segm.y1-i][segm.x1-i]++
				}
			}
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

func findMaxX(segments []segment) int {
	maxX := 0
	for _, segm := range segments {
		if segm.x1 > maxX {
			maxX = segm.x1
		}
		if segm.x2 > maxX {
			maxX = segm.x2
		}
	}
	return maxX
}

func findMaxY(segments []segment) int {
	maxY := 0
	for _, segm := range segments {
		if segm.y1 > maxY {
			maxY = segm.y1
		}
		if segm.y2 > maxY {
			maxY = segm.y2
		}
	}
	return maxY
}
