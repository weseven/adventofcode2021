package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct { x,y int }

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	energyLevels := make([][]int, 0)
	numRows, numCols := 0, 0

	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		var row []int
		elRE := regexp.MustCompile(`\d`)
		enStr := elRE.FindAllString(scanner.Text(), -1)
		for _, s := range enStr {
			energylevel, _ := strconv.Atoi(s)
			row = append(row, energylevel)
		}
		energyLevels = append(energyLevels, row)
		numRows++
	}
	if numRows > 0 {
		numCols = len(energyLevels[numRows-1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one and two
	numFlashes := 0
	for n := 0; n < math.MaxInt32; n++ {
		incrementAll(&energyLevels)
		for i, row := range energyLevels {
			for j := range row {
				if energyLevels[i][j] > 9 {
					simulateFlashes(&energyLevels, Point{x: j, y: i}, numRows, numCols, &numFlashes)
				}
			}
		}
		if n==99 {
	    fmt.Printf("Number of flashes after 100 steps: %d\n", numFlashes)
		}
		if resetFlashed(&energyLevels) == (numRows * numCols) {
			fmt.Printf("Synchronized flash after %d steps.\n", n+1)
			break
		}
	}
}

func simulateFlashes(energyLevels *[][]int, pos Point, numRows int, numCols int, numFlashes *int) {
	(*energyLevels)[pos.y][pos.x]++
	if !((*energyLevels)[pos.y][pos.x] > 9) {
		return
	}
	(*energyLevels)[pos.y][pos.x] = math.MinInt32
	*numFlashes++
	for _, p := range getNeighbours(numRows, numCols, pos) {
		simulateFlashes(energyLevels, p, numRows, numCols, numFlashes)
	}
}

func incrementAll(energyLevels *[][]int) {
	for i, row := range *energyLevels {
		for j := range row {
			(*energyLevels)[i][j]++
		}
	}
}

func resetFlashed(energyLevels *[][]int) (numResets int) {
	for i, row := range *energyLevels {
		for j := range row {
			if (*energyLevels)[i][j] < 0 {
				(*energyLevels)[i][j] = 0
				numResets++
			}
		}
	}
	return numResets
}

func getNeighbours(rows int, cols int, pos Point) (neighbours []Point) {
	if pos.x > 0 {
		neighbours = append(neighbours, Point{y: pos.y, x: pos.x - 1})
		if (pos.y + 1) < rows {
			neighbours = append(neighbours, Point{y: pos.y + 1, x: pos.x - 1})
		}
		if pos.y > 0 {
			neighbours = append(neighbours, Point{y: pos.y - 1, x: pos.x - 1})
		}
	}
	if pos.y > 0 {
		neighbours = append(neighbours, Point{y: pos.y - 1, x: pos.x})
	}
	if (pos.y + 1) < rows {
		neighbours = append(neighbours, Point{y: pos.y + 1, x: pos.x})
	}
	if (pos.x + 1) < cols {
		neighbours = append(neighbours, Point{y: pos.y, x: pos.x + 1})
		if (pos.y + 1) < rows {
			neighbours = append(neighbours, Point{y: pos.y + 1, x: pos.x + 1})
		}
		if pos.y > 0 {
			neighbours = append(neighbours, Point{y: pos.y - 1, x: pos.x + 1})
		}
	}
	return neighbours
}
