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

	crabPositions := make([]int, 0)
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		crabsRE := regexp.MustCompile(`\d+`)
		crabsStr := crabsRE.FindAllString(scanner.Text(), -1)
		for _, str := range crabsStr {
			crabPos, _ := strconv.Atoi(str)
			crabPositions = append(crabPositions, crabPos)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	minPos, maxPos := minMax(crabPositions)
	fuelConsumptions := make(map[int]int)
	minFuel, bestPos := math.MaxInt32, 0
	for i := minPos; i <= maxPos; i++ {
		if _, ok := fuelConsumptions[i]; !ok {
			fuelConsumptions[i] = computeDeltaSum(crabPositions, i)
		}
	}
	for pos, fuel := range fuelConsumptions {
		if fuel < minFuel {
			bestPos = pos
			minFuel = fuel
		}
	}
	fmt.Printf("Minfuel: %d, BestPos: %d\n", minFuel, bestPos)

	// Puzzle two
	twoFuelConsumptions := make(map[int]int)
	twoMinFuel, twoBestPos := math.MaxInt32, 0
	for i := minPos; i <= maxPos; i++ {
		if _, ok := twoFuelConsumptions[i]; !ok {
			twoFuelConsumptions[i] = computeTwoConsumption(crabPositions, i)
		}
	}
	for pos, fuel := range twoFuelConsumptions {
		if fuel < twoMinFuel {
			twoBestPos = pos
			twoMinFuel = fuel
		}
	}
	fmt.Printf("twoMinfuel: %d, twoBestPos: %d\n", twoMinFuel, twoBestPos)
}

func minMax(nums []int) (min int, max int) {
	max = 0
	min = math.MaxInt32
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return min, max
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func computeDeltaSum(nums []int, point int) (deltaSum int) {
	for _, n := range nums {
		deltaSum += abs(n - point)
	}
	return deltaSum
}

func computeTwoConsumption(nums []int, point int) (consumption int) {
	for _, n := range nums {
		steps := abs(n - point)
		consumption += (steps * (steps + 1)) / 2
	}
	return consumption
}
