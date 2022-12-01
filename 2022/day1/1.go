package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// list containing sum of calories carried by
// each elf
var elfCalories []int

func appendTotalCaloriesForElfsItems(stack []int) {
	calories := 0

	for _, food := range stack {
		calories += food
	}

	elfCalories = append(elfCalories, calories)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// stack to hold calories from each line of input
	// and then pop into total for each elf.
	var itemCaloriesStack []int

	for scanner.Scan() {
		textVal := scanner.Text()

		// if line is empty we've seen all the items for an elf
		if textVal == "" {
			appendTotalCaloriesForElfsItems(itemCaloriesStack)
			itemCaloriesStack = nil
			continue
		}

		val, _ := strconv.Atoi(textVal)

		itemCaloriesStack = append(itemCaloriesStack, val)
	}
	// input has finished but if there's still items on the stack
	// there's one more elf
	appendTotalCaloriesForElfsItems(itemCaloriesStack)

	// sort the slice by most calories
	sort.Slice(elfCalories, func(i, j int) bool { return elfCalories[i] > elfCalories[j] })

	// largest of top 3
	fmt.Println(elfCalories[0] + elfCalories[1] + elfCalories[2])
}
