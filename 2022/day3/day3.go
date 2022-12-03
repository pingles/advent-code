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

type bag struct {
	itemsDict map[byte]int
}

func newBag(contents []byte) *bag {
	// create a dict for each bag; we'll then merge together
	b := &bag{}
	b.itemsDict = make(map[byte]int, 52)

	for _, item := range contents {
		if b.itemsDict[item] == 0 {
			b.itemsDict[item] += 1
		}
	}

	return b
}

// addContents combines contents of the bags together
func (b *bag) addContents(other ...*bag) {
	for _, otherBag := range other {
		for k, _ := range otherBag.itemsDict {
			b.itemsDict[k] += 1
		}
	}
}

// find the unique badge (where we have 3 of them)
func (b *bag) badgeType() byte {
	for k, v := range b.itemsDict {
		if v == 3 {
			return k
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	bags := make([]*bag, 0)
	sumOfPriorities := 0

	for scanner.Scan() {
		bag := newBag(scanner.Bytes())
		bags = append(bags, bag)
		counter += 1

		if counter == 3 {
			// join contents of the bags together to find the key
			bags[0].addContents(bags[1:3]...)
			sumOfPriorities += priority(bags[0].badgeType())

			bags = nil
			counter = 0
		}
	}

	fmt.Println(sumOfPriorities)
}
