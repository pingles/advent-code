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

	// parse grid of trees
	grid := make([][]int, 0)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "")
		ints := make([]int, len(items))

		for idx, v := range items {
			val, _ := strconv.Atoi(v)
			ints[idx] = val
		}
		grid = append(grid, ints)
	}
	cols := len(grid[0])
	rows := len(grid)

	// initialise grid to track tree visibility
	visible := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visible[i] = make([]bool, cols)
		for j := 0; j < cols; j++ {
			visible[i][j] = false
		}
	}

	// check left and right visibility on each row
	// starting at the origin
	for rowIndex := 1; rowIndex < rows-1; rowIndex++ {
		leftTallest := grid[rowIndex][0]
		rightTallest := grid[rowIndex][cols-1]

		// move a cursor from left and right
		for j := 1; j < cols-1; j++ {
			// check left side
			leftIdx := j
			leftVal := grid[rowIndex][leftIdx]
			if leftVal > leftTallest {
				visible[rowIndex][leftIdx] = true
				leftTallest = leftVal
			}

			rightIdx := cols - 1 - j
			rightVal := grid[rowIndex][rightIdx]
			if rightVal > rightTallest {
				visible[rowIndex][rightIdx] = true
				rightTallest = rightVal
			}
		}
	}

	// check top and bottom visibility for each column
	// starting at the origin
	for colIndex := 1; colIndex < cols-1; colIndex++ {
		topTallest := grid[0][colIndex]
		bottomTallest := grid[rows-1][colIndex]

		// move a cursor from top and bottom
		for j := 1; j < rows-1; j++ {
			topIdx := j
			topVal := grid[topIdx][colIndex]
			if topVal > topTallest {
				visible[topIdx][colIndex] = true
				topTallest = topVal
			}

			bottomIdx := rows - 1 - j
			bottomVal := grid[bottomIdx][colIndex]
			if bottomVal > bottomTallest {
				visible[bottomIdx][colIndex] = true
				bottomTallest = bottomVal
			}
		}
	}

	// count number of visible trees
	innerCount := 0
	for _, row := range visible {
		for _, isVisible := range row {
			if isVisible {
				innerCount += 1
			}
		}
	}

	// all trees around the outside are visible
	exterior := (rows * 2) + (cols * 2) - 4
	fmt.Println("rows:", rows)
	fmt.Println("cols", cols)
	fmt.Println("inner visible", innerCount)
	fmt.Println("exterior", exterior)
	fmt.Println("total", exterior+innerCount)
}
