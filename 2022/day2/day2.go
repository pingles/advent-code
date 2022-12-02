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

func elfHandIndex(s string) int {
	elfHands := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
	}

	return elfHands[s]
}

func handScoreFromIndex(idx int) int {
	return 1 + idx
}

func scoreRound(s string) int {
	outcomes := [3][3]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	playerHand := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	splits := strings.Split(s, " ")
	elfHand := elfHandIndex(splits[0])
	meHand := playerHand[splits[1]]

	outcomeScore := outcomes[elfHand][meHand]

	return outcomeScore + handScoreFromIndex(meHand)
}

// produces a score given a target outcome and elf hand
func scoreHandToOutcome(s string) int {
	outcomes := [3][3]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	targetOutcomes := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	splits := strings.Split(s, " ")
	elfHand := elfHandIndex(splits[0])
	target := targetOutcomes[splits[1]]

	// find the outcome given the elf's hand
	for idx, score := range outcomes[elfHand] {
		// our hand is the index we found the target outcome in
		if score == target {
			return score + handScoreFromIndex(idx)
		}
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// part 1
	score := 0
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	score += scoreRound(line)
	//}
	//fmt.Printf("Score: %d\n", score)

	// part 2
	score = 0
	for scanner.Scan() {
		line := scanner.Text()
		score += scoreHandToOutcome(line)
	}
	fmt.Printf("Score: %d\n", score)
}
