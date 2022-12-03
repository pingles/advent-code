package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(b byte) int {
	// uppercase?
	if b < 96 {
		return (int(b) % 32) + 26
	}

	return int(b) % 32
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sumOfPriorities := 0
	for scanner.Scan() {
		line := scanner.Text()

		// break contents into 2 compartments
		left := line[:len(line)/2]
		right := line[len(line)/2:]

		// use a dictionary to hold counts of each character
		dict := make(map[byte]int, 52)

		// store values for all of one compartment
		for i := 0; i < len(right); i++ {
			dict[right[i]] = 1
		}

		for i := 0; i < len(left); i++ {
			if dict[left[i]] == 1 {
				// found our item
				sumOfPriorities += priority(left[i])
				break
			}
		}
	}

	fmt.Println(sumOfPriorities)
}
