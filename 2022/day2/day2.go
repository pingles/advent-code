package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scoreRound(s string) int {
	tokens := strings.Split(s, " ")

	elf := tokens[0]
	outcome := tokens[1]

	// lose
	if outcome == "X" {
		// rock
		if elf == "A" {
			// losing hand to rock is scissors
			return 3
		}
		// paper
		if elf == "B" {
			// losing hand is rock
			return 1
		}
		// scissors
		if elf == "C" {
			// losing hand is paper
			return 2
		}
	}

	// draw
	if outcome == "Y" {
		// hand value is the same as the elf
		if elf == "A" {
			return 3 + 1
		}
		if elf == "B" {
			return 3 + 2
		}
		if elf == "C" {
			return 3 + 3
		}
	}

	// win
	if outcome == "Z" {
		if elf == "A" {
			// winning hand against rock is paper
			return 6 + 2
		}
		if elf == "B" {
			// winning aginst paper is scissors
			return 6 + 3
		}
		if elf == "C" {
			// winning against scissors is rock
			return 6 + 1
		}
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		score += scoreRound(line)
	}
	fmt.Printf("Score: %d\n", score)
}
