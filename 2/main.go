package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	movement string
	amount   int
}

func main() {

	inputfile, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)

	instructions := make([]instruction, 0)

	for scanner.Scan() {
		// instr := strings.Split(scanner.Text(), " ")
		instr := strings.Fields(scanner.Text())
		amt, err := strconv.Atoi(instr[1])
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, instruction{movement: instr[0], amount: amt})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Puzzle one
	hPos, vPos := 0, 0

	for _, instruction := range instructions {
		switch instruction.movement {
		case "forward":
			hPos += instruction.amount
		case "up":
			vPos -= instruction.amount
		case "down":
			vPos += instruction.amount
		default:
			log.Fatalf("instruction not recognized: %v", instruction.movement)
		}
	}
	fmt.Printf("Puzzle1: hPos x vPos: %v\n", hPos*vPos)

	// Puzzle two
	hPos, vPos, aim := 0, 0, 0

	for _, instruction := range instructions {
		switch instruction.movement {
		case "forward":
			hPos += instruction.amount
			vPos += instruction.amount * aim
		case "up":
			aim -= instruction.amount
		case "down":
			aim += instruction.amount
		default:
			log.Fatalf("instruction not recognized: %v", instruction.movement)
		}
	}
	fmt.Printf("Puzzle2: hPos x vPos: %v\n", hPos*vPos)

}
