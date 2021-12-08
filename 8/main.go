package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	signalPatterns := make([][]string, 0)
	signalOutputs := make([][]string, 0)
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		signalsRE := regexp.MustCompile(`\w+`)
		signals := signalsRE.FindAllString(scanner.Text(), -1)
		signalPatterns = append(signalPatterns, signals[0:10])
		signalOutputs = append(signalOutputs, signals[10:14])
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	sumSignals := 0
	for _, outputs := range signalOutputs {
		sumSignals += puzzleOne(outputs)
	}

	fmt.Printf("Number of known outputs: %d\n", sumSignals)

	//puzzleTwo
	sum :=0
	for i := range signalOutputs {
		sum += decodeLines(signalPatterns[i],signalOutputs[i])
	}
	fmt.Printf("Puzzle 2 sum outputs: %d\n",sum)
}

func puzzleOne(outputs []string) (count int) {
	for _, str := range outputs {
		if len(str) == 2 || len(str) == 4 || len(str) == 7 || len(str) == 3 {
			count++
		}
	}
	return count
}

func charsInCommon(str1 string, str2 string) (commonChars int) {
	for _, rune1 := range str1 {
		for _, rune2 := range str2 {
			if rune1 == rune2 {
				commonChars++
			}
		}
	}
	return commonChars
}

func isEqualDisp(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	} else {
		if charsInCommon(str1, str2) == len(str1) {
			return true
		}
	}
	return false
}

func containsStrAt(strArray []string, str string) int {
	for idx, s := range strArray{
		if isEqualDisp(s,str){
			return idx
		}
	}
	return -1
}

func decodeLines(patterns []string, outputs []string) (num int) {
	decoded := make([]string, 10)
	tobedecoded := make([]string, 0)
	for _, str := range patterns {
		switch len(str) {
		case 4:
			decoded[4] = str
		case 2:
			decoded[1] = str
		case 3:
			decoded[7] = str
		case 7:
			decoded[8] = str
		default:
			tobedecoded = append(tobedecoded, str)
		}
	}
	for _, str := range tobedecoded {
		if len(str) == 5 {
			if charsInCommon(str, decoded[7]) == 3 {
				decoded[3] = str
			} else {
				if charsInCommon(str, decoded[4]) == 3 {
					decoded[5] = str
				} else {
					decoded[2] = str
				}
			}
		}
		if len(str) == 6 {
			if charsInCommon(str, decoded[4]) == 4 {
				decoded[9] = str
			} else {
				if charsInCommon(str, decoded[7]) == 3 {
					decoded[0] = str
				} else {
					decoded[6] = str
				}
			}
		}
	}
	num = 0
	for _, out := range outputs {
		num = 10*num + containsStrAt(decoded,out)
	}
	return num
}
