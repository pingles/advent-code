package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	val  int
	next *node
}

// create a linked list that will maintain sort order
// will place largest item (by val) at the front
type sortedLinkedList struct {
	start  *node
	length int
}

// insert val in order so that the start of the list
// is the largest number
func (l *sortedLinkedList) insert(val int) {
	newNode := &node{val: val}

	// empty list
	if l.length == 0 {
		l.start = newNode
		l.length++

		return
	}

	// do we need to replace the head
	if newNode.val > l.start.val {
		newNode.next = l.start
		l.start = newNode
		l.length++

		return
	}

	// need to search for where the next node has
	// a smaller value than newNode
	node := l.start
	for {
		if node.next == nil {
			// we're at the tail
			node.next = newNode
			l.length++

			return
		}

		if node.next.val < newNode.val {
			newNode.next = node.next
			node.next = newNode
			l.length++

			return
		}

		node = node.next
	}
}

var elves = &sortedLinkedList{}

func sum(stack []int) int {
	sum := 0

	for _, x := range stack {
		sum += x
	}

	return sum
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
			elves.insert(sum(inputStack))
			inputStack = nil
			continue
		}

		val, _ := strconv.Atoi(textVal)

		inputStack = append(inputStack, val)
	}
	// input has finished but process remaining stack
	if len(inputStack) > 0 {
		elves.insert(sum(inputStack))
	}

	fmt.Println("Elves:")
	fmt.Println(elves.length)

	fmt.Println("Largest:")
	fmt.Println(elves.start.val)

	fmt.Println("Largest 3:")
	fmt.Println(elves.start.val + elves.start.next.val + elves.start.next.next.val)
}
