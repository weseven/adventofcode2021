package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
)

func main() {

	inputfile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	values := make([]int, 0)
	rows, cols := 0, 0
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		for _, rune := range runes {
			n, _ := strconv.Atoi(string(rune))
			values = append(values, n)
		}
		rows++
		cols = len(runes)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	graph := dijkstra.NewGraph()
	for i := range values {
		graph.AddVertex(i)
	}

	for i := range values {
		for _, neighbour := range getNeighbours(rows, cols, i) {
			graph.AddArc(i, neighbour, int64(values[neighbour]))
		}
	}

	bestPath, err := graph.Shortest(0, (rows*cols)-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Puzzle one: %d\n", bestPath.Distance)

	// Puzzle two
	fullvalues := make([]int, rows*cols*25)
	for i := 0; i < 5; i++ {
		for rowidx := 0; rowidx < rows; rowidx++ {
			for colidx := 0; colidx < cols; colidx++ {
				for j := 0; j < 5; j++ {
					fullvalues[colidx+(j*cols)+rowidx*cols*5+i*cols*5*rows] = ((values[colidx+rowidx*cols] + j + i - 1) % 9) + 1
				}
			}
		}
	}
	graph2 := dijkstra.NewGraph()
	for i := range fullvalues {
		graph2.AddVertex(i)
	}
	for i := range fullvalues {
		for _, neighbour := range getNeighbours(rows*5, cols*5, i) {
			graph2.AddArc(i, neighbour, int64(fullvalues[neighbour]))
		}
	}
	bestPath2, err := graph2.Shortest(0, len(fullvalues)-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Puzzle two: %d\n", bestPath2.Distance)
}

func getNeighbours(rows int, cols int, pos int) (neighbours []int) {
	if pos%cols != 0 {
		neighbours = append(neighbours, pos-1)
	}
	if pos >= cols {
		neighbours = append(neighbours, pos-cols)
	}
	if pos%cols != (cols - 1) {
		neighbours = append(neighbours, pos+1)
	}
	if (pos / cols) < (rows - 1) {
		neighbours = append(neighbours, pos+cols)
	}
	return neighbours
}
