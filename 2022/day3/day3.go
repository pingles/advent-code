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

// dual dictionary; store items and frequency, and lookup items
// and their frequency
type dictAndFreq struct {
	d       map[byte]int
	inverse map[int]byte
}

func newDict(size int) *dictAndFreq {
	return &dictAndFreq{
		d:       make(map[byte]int, size),
		inverse: make(map[int]byte, size),
	}
}

// adds items; will only add item from items once
func (d *dictAndFreq) addUniqueItems(items []byte) {
	// use a dict to make sure we process each from items only once
	itemsDict := make(map[byte]bool, len(items))

	for _, item := range items {
		if itemsDict[item] {
			continue // already processed it from this list
		}

		d.d[item] += 1
		// store the item keyed on its frequency; don't need to worry about clashes
		// as there's only 1 item at the target frequency
		d.inverse[d.d[item]] = item
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
		dict.addUniqueItems(scanner.Bytes())
		counter += 1

		if counter == 3 {
			sumOfPriorities += priority(dict.inverse[3])
			counter = 0
			dict = newDict(dictionarySize)
		}
	}

	fmt.Println(sumOfPriorities)
}
