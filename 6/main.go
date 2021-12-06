package main

import (
	"bufio"
	"fmt"
	"log"
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

	fishes := make([]int, 0)
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		fishesRE := regexp.MustCompile(`\d+`)
		fishesStr := fishesRE.FindAllString(scanner.Text(), -1)
		for _, str := range fishesStr {
			fish, _ := strconv.Atoi(str)
			fishes = append(fishes, fish)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	// fmt.Println(fishes)
	// days := 80
	// for i:=0;i<days;i++{
	// 	passDay(&fishes)
	// }
	// fmt.Printf("Number of fishes after 80 days: %d", len(fishes))

	// Puzzle two
	fishesCounts := make([]int, 9)
	for _, f := range fishes {
		fishesCounts[f]++
	}

	days := 256
	for i := 0; i < days; i++ {
		spawningFishes := fishesCounts[0]
		for t := 0; t < 8; t++ {
			fishesCounts[t] = fishesCounts[t+1]
		}
		fishesCounts[8] = spawningFishes
		fishesCounts[6] += spawningFishes
	}
	var sumFishes int
	for i := range fishesCounts {
		sumFishes += fishesCounts[i]
	}

	fmt.Printf("Fishes after %d days: %d", days, sumFishes)

}

// func passDay(fishes *[]int){
// 	for i := range *fishes{
// 		(*fishes)[i]=(*fishes)[i]-1
// 		if (*fishes)[i] == -1 {
// 			(*fishes)[i] = 6
// 			(*fishes) = append((*fishes), 8)
// 		}
// 	}
// }
