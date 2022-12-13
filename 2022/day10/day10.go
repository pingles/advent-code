package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	instructions := make([]*instruction, 0)
	registers := &registers{x: 1}

	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}

	cycleCounter := 1
	instructionIndex := 0
	interestingCounter := 19              // used to track the interesting signals; first @ 20 then each 40 thereafter
	currentInstruction := instructions[0] // ignore that we may not have any
	sumOfInteresting := 0
	for {
		// is this an interesting cycle?
		if interestingCounter == 0 {
			val := cycleCounter * registers.x
			fmt.Printf("[%d] %d*%d = %d\n", cycleCounter, cycleCounter, registers.x, val)
			sumOfInteresting += val
			interestingCounter = 39 // reset counter
		} else {
			interestingCounter--
		}

		instructionCompleted := currentInstruction.apply(registers)
		//fmt.Printf("[%d] %v: %s\n", cycleCounter, registers, currentInstruction)

		if instructionCompleted {
			if instructionIndex == len(instructions)-1 {
				// finished the program
				break
			}

			// get the next instruction
			instructionIndex += 1
			currentInstruction = instructions[instructionIndex]
		}

		cycleCounter += 1
	}

	fmt.Println("Total:", sumOfInteresting)
}

const (
	noop = iota
	addx
)

type registers struct {
	x int
}

type instruction struct {
	op              int
	arg             int
	cyclesRemaining int // how many cycles are left to complete this instruction
}

func (i *instruction) apply(r *registers) bool {
	i.cyclesRemaining--
	if i.cyclesRemaining > 0 {
		return false
	}

	if i.op == addx {
		r.x += i.arg
	}

	return true
}

func (i *instruction) String() string {
	if i.op == noop {
		return fmt.Sprintf("([%d]noop)", i.cyclesRemaining)
	}

	return fmt.Sprintf("([%d]addx %d)", i.cyclesRemaining, i.arg)
}

func parseInstruction(text string) *instruction {
	parts := strings.Split(text, " ")

	if len(parts) == 1 {
		// noop
		return &instruction{op: noop, cyclesRemaining: 1}
	}

	arg, _ := strconv.Atoi(parts[1])
	return &instruction{op: addx, arg: arg, cyclesRemaining: 2}
}
