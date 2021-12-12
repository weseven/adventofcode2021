package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	connections := make(map[string][]string)
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		conns := strings.Split(scanner.Text(), "-")
		connections[conns[0]] = append(connections[conns[0]], conns[1])
		connections[conns[1]] = append(connections[conns[1]], conns[0])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	paths := make(map[string]interface{})
	exploreOne(connections, "start", "", &paths)

	fmt.Printf("Total paths puzzle one: %d\n", len(paths))

	// Puzzle two
	pathsTwo := make(map[string]interface{})
	exploreTwo(connections, "start", "", &pathsTwo, false)

	fmt.Printf("Total paths puzzle two: %d\n", len(pathsTwo))
}

func exploreOne(connections map[string][]string, currentNode string, currentPath string, paths *map[string]interface{}) {
	currentPath += "-" + currentNode
	if currentNode == "end" {
		(*paths)[currentPath] = nil
		return
	}
	for _, nextNode := range connections[currentNode] {
		if nextNode == "start" {
			continue
		}
		if isBig(nextNode) || !hasBeenExploredOnce(nextNode, currentPath) {
			exploreOne(connections, nextNode, currentPath, paths)
		}
	}
}

func exploreTwo(connections map[string][]string, currentNode string, currentPath string, paths *map[string]interface{}, smallCaveTwice bool) {
	currentPath += "-" + currentNode
	if !isBig(currentNode) && hasBeenExploredTwice(currentNode, currentPath) {
		smallCaveTwice = true
	}
	if currentNode == "end" {
		(*paths)[currentPath] = nil
		return
	}
	for _, nextNode := range connections[currentNode] {
		if nextNode == "start" {
			continue
		}
		if isBig(nextNode) || !hasBeenExploredOnce(nextNode, currentPath) || (!smallCaveTwice && !hasBeenExploredTwice(nextNode, currentPath)) {
			exploreTwo(connections, nextNode, currentPath, paths, smallCaveTwice)
		}
	}
}

func hasBeenExploredOnce(cave string, path string) bool {
	return strings.Contains(path, "-"+cave)
}

func isBig(cave string) bool {
	return cave == strings.ToUpper(cave)
}

func hasBeenExploredTwice(cave string, path string) bool {
	return (strings.Count(path, "-"+cave) > 1)
}
