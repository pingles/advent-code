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

// get the potential outcomes given an elf hand
func elfOutcomes(elf string) [3]int {
	outcomes := [3][3]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}
	elfHand := elfHandIndex(elf)
	return outcomes[elfHand]
}

func scoreRound(elf, player string) int {
	playerHand := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	meHand := playerHand[player]
	outcomeScore := elfOutcomes(elf)[meHand]

	return outcomeScore + handScoreFromIndex(meHand)
}

// produces a score given a target outcome and elf hand
func scoreHandToOutcome(elf, target string) int {
	targetOutcomes := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	targetScore := targetOutcomes[target]

	// find the outcome given the elf's hand
	for idx, score := range elfOutcomes(elf) {
		// our hand is the index we found the target outcome in
		if score == targetScore {
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
		splits := strings.Split(line, " ")
		score += scoreHandToOutcome(splits[0], splits[1])
	}
	fmt.Printf("Score: %d\n", score)
}
