package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type directory struct {
	childDirectories map[string]*directory
	sizeOfFiles      int
	parent           *directory
}

func newDirectory() *directory {
	return &directory{
		childDirectories: make(map[string]*directory),
	}
}

func (d *directory) addChildDirectory(name string) *directory {
	dir := newDirectory()
	dir.parent = d
	d.childDirectories[name] = dir
	return dir
}

func (d *directory) addFile(size int) {
	d.sizeOfFiles += size
}

// changes to sub directory
func (d *directory) cd(s string) *directory {
	if s == ".." {
		return d.parent
	}
	return d.childDirectories[s]
}

func (d *directory) size() int {
	sum := d.sizeOfFiles
	for _, v := range d.childDirectories {
		sum += v.size()
	}
	return sum
}

// interprets the output of a command
type interpretCommandOutput func(parts []string)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	currentDirectory := newDirectory()
	rootDirectory := currentDirectory.addChildDirectory("/") // create the root

	// push the currentDirectory command here; this will interpret the output
	var commandOutputHandler interpretCommandOutput

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if parts[0] == "$" {
			switch parts[1] {
			case "cd":
				currentDirectory = currentDirectory.cd(parts[2])
				commandOutputHandler = nil // no output
			case "ls":
				commandOutputHandler = func(lineParts []string) {
					// add items to the current directory
					if lineParts[0] == "dir" {
						// add a child directory
						currentDirectory.addChildDirectory(lineParts[1])
					} else {
						// add a file and size
						size, _ := strconv.Atoi(lineParts[0])
						currentDirectory.addFile(size)
					}
				}
			}
		} else {
			if commandOutputHandler != nil {
				commandOutputHandler(parts)
			}
		}
	}

	findPartA(rootDirectory)
	findPartB(rootDirectory)
}

func findPartA(root *directory) {
	directories := make([]*directory, 0)
	findChildren(&directories, root, func(dir *directory) bool {
		return dir.size() <= 100000
	})
	sum := 0
	for _, v := range directories {
		sum += v.size()
	}
	fmt.Println("<= 100000", sum)
}

func findPartB(root *directory) {
	unusedSpace := 70000000 - root.size()
	installNeeds := 30000000
	remainingNeeded := installNeeds - unusedSpace

	directories := make([]*directory, 0)
	findChildren(&directories, root, func(dir *directory) bool {
		return dir.size() >= remainingNeeded
	})
	sort.Slice(directories, func(i, j int) bool {
		return directories[i].size() < directories[j].size()
	})
	fmt.Println("Smallest needed to free", remainingNeeded, directories[0].size())
}

// recursively find directories matching the filter
func findChildren(directories *[]*directory, currentDirectory *directory, filter func(dir *directory) bool) {
	if filter(currentDirectory) {
		*directories = append(*directories, currentDirectory)
	}
	for _, child := range currentDirectory.childDirectories {
		findChildren(directories, child, filter)
	}
}
