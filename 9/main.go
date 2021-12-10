package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	heightmap := make([]int, 0)

	scanner := bufio.NewScanner(inputfile)
	numRows, numCols := 0, 0
	for scanner.Scan() {
		heightsRE := regexp.MustCompile(`\d`)
		heights := heightsRE.FindAllString(scanner.Text(), -1)
		for _, str := range heights {
			height, _ := strconv.Atoi(str)
			heightmap = append(heightmap, height)
		}
		if numCols == 0 {
			numCols = len(heights)
		}
		numRows++
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lowPoints := make([]int, 0)
	// Puzzle one
	sumRisk := 0
	for i, h := range heightmap {
		if isLowPoint(heightmap, numRows, numCols, i) {
			lowPoints = append(lowPoints, i)
			sumRisk += (h + 1)
		}
	}

	fmt.Printf("Puzzle one sum risk: %d\n", sumRisk)

	// Puzzle two
	basinSizes := make([]int, len(lowPoints))
	explored := make([]bool, len(heightmap))
	for i, lowPoint := range lowPoints {
		basinSizes[i] = exploreBasins(heightmap, lowPoint, &explored, numRows, numCols)
	}
	// fmt.Printf("BasinSizes: %v\n", basinSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Printf("Puzzle 2 top 3 basinSizes product: %d\n",basinSizes[0]*basinSizes[1]*basinSizes[2])

}

func exploreBasins(heightmap []int, pos int, explored *[]bool, rows int, cols int) (basinSize int) {
	(*explored)[pos] = true
	if (pos - cols) >= 0 {
		if heightmap[pos-cols] != 9 && !(*explored)[pos-cols] {
			//explore top
			basinSize += exploreBasins(heightmap, pos-cols, explored, rows, cols)
		}
	}
	if (pos/cols) < (rows-1) && !(*explored)[pos+cols] {
		if heightmap[pos+cols] != 9 {
			//explore bottom
			basinSize += exploreBasins(heightmap, pos+cols, explored, rows, cols)
		}
	}
	if (pos%cols) != 0 && !(*explored)[pos-1] {
		if heightmap[pos-1] != 9 {
			//explore left
			basinSize += exploreBasins(heightmap, pos-1, explored, rows, cols)
		}
	}
	if (pos%cols) != (cols-1) && !(*explored)[pos+1] {
		if heightmap[pos+1] != 9 {
			//explore right
			basinSize += exploreBasins(heightmap, pos+1, explored, rows, cols)
		}
	}
	basinSize += 1

	return basinSize
}

func isLowPoint(heightmap []int, rows int, cols int, pos int) bool {
	//angles
	if pos == 0 {
		return ((heightmap[pos] < heightmap[pos+1]) && (heightmap[pos] < heightmap[pos+cols]))
	}
	if pos == (cols - 1) {
		return ((heightmap[pos] < heightmap[pos-1]) && (heightmap[pos] < heightmap[pos+cols]))
	}
	if pos == (rows-1)*(cols) {
		return ((heightmap[pos] < heightmap[pos+1]) && (heightmap[pos] < heightmap[pos-cols]))
	}
	if pos == ((rows * cols) - 1) {
		return ((heightmap[pos] < heightmap[pos-1]) && (heightmap[pos] < heightmap[pos-cols]))
	}
	if pos%cols == 0 {
		//first column
		return ((heightmap[pos] < heightmap[pos-cols]) && (heightmap[pos] < heightmap[pos+1]) && (heightmap[pos] < heightmap[pos+cols]))
	}
	if pos/cols == 0 {
		//first row
		return ((heightmap[pos] < heightmap[pos-1]) && (heightmap[pos] < heightmap[pos+cols]) && (heightmap[pos] < heightmap[pos+1]))
	}
	if pos%cols == (cols - 1) {
		//last column
		return ((heightmap[pos] < heightmap[pos-cols]) && (heightmap[pos] < heightmap[pos-1]) && (heightmap[pos] < heightmap[pos+cols]))
	}
	if pos/cols == (rows - 1) {
		//last row
		return ((heightmap[pos] < heightmap[pos-1]) && (heightmap[pos] < heightmap[pos-cols]) && (heightmap[pos] < heightmap[pos+1]))
	}

	return ((heightmap[pos] < heightmap[pos-1]) && (heightmap[pos] < heightmap[pos-cols]) && (heightmap[pos] < heightmap[pos+1]) && (heightmap[pos] < heightmap[pos+cols]))

}
