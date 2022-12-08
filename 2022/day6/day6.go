package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findMarkerIndex(bytes []byte, messageSize int) int {
	dict := make(map[byte]int)

	for i := 0; i < len(bytes); i++ {
		if i >= messageSize {
			// decrement dictionary counter; this means we don't delete
			// in-window items that share the same key as something that
			// has moved outside the window
			dict[bytes[i-messageSize]] -= 1
		}

		dict[bytes[i]] += 1

		if i >= messageSize-1 {
			// use a counter to check map frequencies of each item
			// if we're at 1 each, then we've found a unique run
			count := 0

			// TODO
			// this is unnecessary; could add and remove on each item
			// avoiding multiple loops
			for j := i; j >= i-(messageSize-1); j-- {
				count += dict[bytes[j]]
			}

			if count == messageSize {
				return i + 1
			}
		}
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		bytes := scanner.Bytes()
		messageSize, _ := strconv.Atoi(os.Args[1])
		fmt.Println("Marker", findMarkerIndex(bytes, messageSize))
	}
}
