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

type Point struct {
	x int
	y int
}

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	heightmap := make([][]int, 0)

	scanner := bufio.NewScanner(inputfile)
	numRows, numCols := 0, 0
	for scanner.Scan() {
		heightsRE := regexp.MustCompile(`\d`)
		heights := heightsRE.FindAllString(scanner.Text(), -1)
		var row []int
		for _, str := range heights {
			height, _ := strconv.Atoi(str)
			row = append(row, height)
		}
		heightmap = append(heightmap, row)
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

	lowPoints := make([]Point, 0)
	// Puzzle one
	sumRisk := 0
	for i, row := range heightmap {
		for j, h := range row {
			if isLowPoint(heightmap, numRows, numCols, Point{x:j,y:i}) {
				lowPoints = append(lowPoints, Point{x:j,y:i})
				sumRisk += (h + 1)
			}
		}
	}
	fmt.Printf("Puzzle one sum risk: %d\n", sumRisk)

	// Puzzle two
	basinSizes := make([]int, len(lowPoints))
	explored := make(map[Point]bool)
	for i, lowPoint := range lowPoints {
		basinSizes[i] = exploreBasins(heightmap, lowPoint, &explored, numRows, numCols)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Printf("Puzzle 2 top 3 basinSizes product: %d\n", basinSizes[0]*basinSizes[1]*basinSizes[2])
}

func exploreBasins(heightmap [][]int, pos Point, explored *map[Point]bool, rows int, cols int) (basinSize int) {
	(*explored)[pos] = true
	basinSize += 1
	neighbourPoints := getNeighbours(rows,cols,pos)
	for _, neighbourPoint := range neighbourPoints{
		if heightmap[neighbourPoint.y][neighbourPoint.x]!=9 && !(*explored)[neighbourPoint]{
			basinSize+=exploreBasins(heightmap,neighbourPoint,explored,rows,cols)
		}
	}
	return basinSize
}

func getNeighbours(rows int, cols int, pos Point) (neighbours []Point) {
	if pos.x > 0 {
		neighbours = append(neighbours, Point{y:pos.y, x: pos.x-1})
	}
	if pos.y > 0 {
		neighbours = append(neighbours, Point{y:pos.y-1, x: pos.x})
	}
	if (pos.y + 1) < rows {
		neighbours = append(neighbours, Point{y:pos.y+1, x: pos.x})
	}
	if (pos.x + 1) < cols {
		neighbours = append(neighbours, Point{y:pos.y, x: pos.x+1})
	}
	return neighbours
}

func isLowPoint(heightmap [][]int, rows int, cols int, pos Point) bool {
	for _, neighbourPoint := range getNeighbours(rows, cols, pos) {
		if !(heightmap[pos.y][pos.x] < heightmap[neighbourPoint.y][neighbourPoint.x]) {
			return false
		}
	}
	return true
}
