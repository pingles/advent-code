package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// we can model the scores in a matrix
//            Paper
//        Rock     Scissors
// e.g.     A   B   C
//        X D   W   L
//        Y L   D   W
//        Z W   L   D
//
// "Our" hand is in the columns, so we find the cell for the result
// in the extension, we're interested in finding the losing outcome
// given the opposition, so we read from the row across to the losing
// one and read up to the top

const (
	rock     = 0
	paper    = 1
	scissors = 2
)

func scoreRound(s string) int {
	outcomes := [3][3]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	elfHands := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
	}

	playerHand := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	splits := strings.Split(s, " ")
	elfHand := elfHands[splits[0]]
	meHand := playerHand[splits[1]]

	outcomeScore := outcomes[elfHand][meHand]
	handScore := 1 + meHand
	
	return outcomeScore + handScore
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
