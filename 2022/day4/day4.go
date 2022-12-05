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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fullyOverlap := 0

	for scanner.Scan() {
		line := scanner.Text()

		intervalExpressions := strings.Split(line, ",")
		leftStart, leftEnd := parseInterval(intervalExpressions[0])
		rightStart, rightEnd := parseInterval(intervalExpressions[1])

		if leftStart == rightStart {
			// start at the same point
			// they have to overlap
			fullyOverlap += 1
		} else if leftStart < rightStart {
			// is right fully contained in left?
			if rightEnd <= leftEnd {
				fullyOverlap += 1
			}
		} else {
			// is left fully contained in right?
			if leftEnd <= rightEnd {
				fullyOverlap += 1
			}
		}
	}

	fmt.Println("Overlap:", fullyOverlap)
}
