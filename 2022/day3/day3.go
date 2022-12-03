package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(b byte) int {
	if b < 96 {
		// uppercase characters have double
		// the priority of lower case ones
		return (int(b) % 32) + 26
	}

	return int(b) % 32
}

// dictionary for lookups. won't count more than once per addition
type uniqueDict struct {
	d map[byte]int
}

func newDict(size int) *uniqueDict {
	return &uniqueDict{d: make(map[byte]int, size)}
}

// adds items; will only add item from items once
func (d *uniqueDict) addItems(items []byte) {
	// use a dict to make sure we process each from items only once
	itemsDict := make(map[byte]bool, len(items))

	for _, item := range items {
		if itemsDict[item] {
			continue // already processed it from this list
		}

		d.d[item] += 1
		itemsDict[item] = true
	}
}

const dictionarySize = 52

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	sumOfPriorities := 0

	dict := newDict(dictionarySize) // create a dictionary for the group
	for scanner.Scan() {
		dict.addItems(scanner.Bytes())
		counter += 1

		if counter == 3 {
			// find the one item thats been counted 3 times
			for k, v := range dict.d {
				if v == 3 {
					sumOfPriorities += priority(k)
				}
			}

			counter = 0
			dict = newDict(dictionarySize)
		}
	}

	fmt.Println(sumOfPriorities)
}
