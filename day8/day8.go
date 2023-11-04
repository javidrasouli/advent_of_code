package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// part one
	file, err := os.Open("/home/javid/project/mine/fun/advent_of_code/day8/input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	defer file.Close()
	var trees []string
	for fileScanner.Scan() {
		trees = append(trees, fileScanner.Text())
	}

	count := findTree(trees)

	fmt.Println("count is :", count)
}

func findTree(trees []string) int {
	count := int(0)

	for i, tree := range trees {
		if i == 0 || i == len(trees)-1 {
			count = count + len(tree)
			continue
		}

		for i2 := 0; i2 < len(tree); i2++ {
			if i2 == 0 || i2 == len(tree)-1 {
				count++
				continue
			}

			t := tree[i2]

			isVisible := false

			for up := i - 1; up >= 0; up-- {
				if trees[up][i2] < t {
					isVisible = true
				} else {
					isVisible = false
					break
				}
			}

			if isVisible {
				count++
				continue
			}

			for down := i + 1; down <= len(trees)-1; down++ {
				if trees[down][i2] < t {
					isVisible = true
				} else {
					isVisible = false
					break
				}
			}

			if isVisible {
				count++
				continue
			}

			for left := i2 - 1; left >= 0; left-- {
				if tree[left] < t {
					isVisible = true
				} else {
					isVisible = false
					break
				}
			}

			if isVisible {
				count++
				continue
			}

			for right := i2 + 1; right <= len(tree)-1; right++ {
				if tree[right] < t {
					isVisible = true
				} else {
					isVisible = false
					break
				}
			}

			if isVisible {
				count++
			}
		}

	}

	return count
}
