package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInterval(s string) (int, int) {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return start, end
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	anyOverlap := 0

	for scanner.Scan() {
		line := scanner.Text()

		intervalExpressions := strings.Split(line, ",")
		leftStart, leftEnd := parseInterval(intervalExpressions[0])
		rightStart, rightEnd := parseInterval(intervalExpressions[1])

		if max(leftStart, rightStart) <= min(leftEnd, rightEnd) {
			anyOverlap += 1
		}
	}

	fmt.Println("Overlap:", anyOverlap)
}
