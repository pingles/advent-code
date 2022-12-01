package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
    val int
    next *node
}

// create a linked list that will maintain sort order
// will place largest item (by val) at the front
type sortedLinkedList struct {
    start *node
    length int
}

// insert val in order so that the start of the list
// is the largest number
func (l *sortedLinkedList) insert(val int) {
    n := &node{val: val}
    if l.length == 0 {
        l.start = n
        l.length++
        return
    }

    indexNode := l.start
    for {
        if l.length == 1 {
            // if current start is smaller, put new node in front
            if l.start.val < val {
                n.next = l.start // put the old node as the next
                l.start = n
            }

            break
        }

        // there's more than 1 node, so we need to find where we fit
        // we need to find where indexNode > newNode > indexNode.next
        if indexNode.val > val {
            // we've found the "left" side of the list
            n.next = indexNode.next
            indexNode.next = n
        }
    }
}

var elves = &sortedLinkedList{}

// appends calories in sorted order. means that we
// don't need to perform a sort later.
func appendTotalCaloriesForElfsItems(stack []int) {
	calories := 0

	for _, food := range stack {
		calories += food
	}

    elves.insert(calories)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// stack to hold calories from each line of input
	// and then pop into total for each elf.
	var inputStack []int

	for scanner.Scan() {
		textVal := scanner.Text()

		// if line is empty we've seen all the items for an elf
		if textVal == "" {
            appendTotalCaloriesForElfsItems(inputStack)
            inputStack = nil
			continue
		}

		val, _ := strconv.Atoi(textVal)

        inputStack = append(inputStack, val)
	}
	// input has finished but if there's still items on the stack
	// there's one more elf
    appendTotalCaloriesForElfsItems(inputStack)

    fmt.Println("Largest:")
	fmt.Println(elves.start.val)

    fmt.Println("Largest 3:")
    fmt.Println(elves.start.val + elves.start.next.val + elves.start.next.next.val)
}