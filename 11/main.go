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
		numCols = len(row)
		numRows++
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
					simulateFlashes(&energyLevels, j, i, numRows, numCols, &numFlashes)
				}
			}
		}
		if resetFlashed(&energyLevels) == (numRows*numCols) {
			fmt.Printf("Synchronized flash after %d steps.\n",n+1)
			break
		}
	}

	// fmt.Printf("Number of flashes after 100 steps: %d\n", numFlashes)

}

func incrementAll(energyLevels *[][]int) {
	for i, row := range *energyLevels {
		for j := range row {
			(*energyLevels)[i][j]++
		}
	}
}

func resetFlashed(energyLevels *[][]int) (numResets int){
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

func simulateFlashes(energyLevels *[][]int, posx int, posy int, numRows int, numCols int, numFlashes *int) {
	(*energyLevels)[posy][posx]++
	if !((*energyLevels)[posy][posx] > 9) {
		return
	}
	(*energyLevels)[posy][posx] = math.MinInt32
	*numFlashes++
	if posx > 0 {
		simulateFlashes(energyLevels, posx-1, posy, numRows, numCols, numFlashes)
		if (posy + 1) < numRows {
			simulateFlashes(energyLevels, posx-1, posy+1, numRows, numCols, numFlashes)
		}
		if posy > 0 {
			simulateFlashes(energyLevels, posx-1, posy-1, numRows, numCols, numFlashes)
		}
	}
	if posy > 0 {
		simulateFlashes(energyLevels, posx, posy-1, numRows, numCols, numFlashes)
	}
	if (posy + 1) < numRows {
		simulateFlashes(energyLevels, posx, posy+1, numRows, numCols, numFlashes)
	}
	if (posx + 1) < numCols {
		simulateFlashes(energyLevels, posx+1, posy, numRows, numCols, numFlashes)
		if (posy + 1) < numRows {
			simulateFlashes(energyLevels, posx+1, posy+1, numRows, numCols, numFlashes)
		}
		if posy > 0 {
			simulateFlashes(energyLevels, posx+1, posy-1, numRows, numCols, numFlashes)
		}
	}
}
