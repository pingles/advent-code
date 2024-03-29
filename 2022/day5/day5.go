package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	readingData := true

	towers := make([]*list.List, 0)

	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		if readingData {
			if len(lineBytes) >= 3 {
				tokens := parseDataLineToTokens(lineBytes)

				// check we have created enough towers
				if len(towers) < len(tokens) {
					for i := len(towers); i < len(tokens); i++ {
						towers = append(towers, list.New())
					}
				}

				// add any blocks to the correct tower
				// we're reading from the top to the base of the tower
				// so need to append to the end of the stack
				for idx, token := range tokens {
					if token.t == block {
						towers[idx].PushBack(token.label)
					}
				}
			}
		} else {
			// reading operations
			executeOperation(towers, lineBytes)
		}

		if len(lineBytes) == 0 {
			// switching from input to operations
			readingData = false
		}
	}

	// what's the head
	labels := make([]string, len(towers))
	for idx, t := range towers {
		labels[idx] = t.Front().Value.(string)
	}
	fmt.Println(strings.Join(labels, ""))
}

func executeOperation(towers []*list.List, bytes []byte) {
	parts := strings.Split(string(bytes), " ")
	numberOfItems, _ := strconv.Atoi(parts[1])
	source, _ := strconv.Atoi(parts[3])
	dest, _ := strconv.Atoi(parts[5])

	// create an intermediate list so we can reverse the order
	// to insert items correctly in a batch
	itemsToMove := make([]string, numberOfItems)

	// TODO: avoid needing to use intermediate list
	for i := 0; i < numberOfItems; i++ {
		source := towers[source-1]
		val := source.Remove(source.Front())
		itemsToMove[i] = val.(string)
	}

	for i := numberOfItems - 1; i >= 0; i-- {
		towers[dest-1].PushFront(itemsToMove[i])
	}
}

const (
	block = iota
	id
	whitespace
)

type token struct {
	t     int // the type of token
	label string
}

const WhitespaceByte = 32
const OpeningParenByte = 91

func parseDataLineToTokens(bytes []byte) []*token {
	// read 3 bytes
	//  all bytes will be blank if its whitespace
	//  first byte will be [ if its a block (second will be label)
	//     second byte will be block label if its a block
	//  first and last are whitespace if its an ID block
	// repeat
	tokens := make([]*token, 0)
	i := 0

	for {
		if bytes[i+1] == WhitespaceByte {
			tokens = append(tokens, &token{t: whitespace})
		} else if bytes[i] == OpeningParenByte {
			tokens = append(tokens, &token{t: block, label: string(bytes[i+1])})
		} else {
			// should be on the ID row
			tokens = append(tokens, &token{t: id, label: string(bytes[i+1])})
		}

		// move to the next token
		i += 4

		if i >= len(bytes) {
			break
		}
	}

	return tokens
}
