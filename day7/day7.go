package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("E:\\JavidProjects\\adventofcode\\day7\\input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	defer file.Close()
	var DirPath []string
	FileSystem := make(map[string]int)
	UsedSpace := 0
	// part one
	for fileScanner.Scan() {
		commandSplit := strings.Split(fileScanner.Text(), " ")
		switch commandSplit[0] {
		case "$":
			switch commandSplit[1] {
			case "cd":
				if commandSplit[2] != ".." {
					DirPath = append(DirPath, commandSplit[2])
				} else {
					DirPath = DirPath[:len(DirPath)-1]
				}
			case "ls":
				fmt.Println("i want to list this dir")
			}
		case "dir":
			fmt.Println("dir is :", commandSplit[1])
		default:
			size, _ := strconv.Atoi(commandSplit[0])
			UsedSpace = UsedSpace + size
			FileSystem[strings.Join(DirPath[1:], "/")] = FileSystem[strings.Join(DirPath[1:], "/")] + size
		}
	}
	TotalValueDir := make(map[string]int)
	for s, i := range FileSystem {
		dirSplit := strings.Split(s, "/")
		lenDir := len(dirSplit)
		for dirNum := 0; dirNum < lenDir; dirNum++ {
			TotalValueDir[strings.Join(dirSplit, "/")] = TotalValueDir[strings.Join(dirSplit, "/")] + i
			dirSplit = dirSplit[:len(dirSplit)-1]
		}
	}
	Total := 0
	for _, i2 := range TotalValueDir {
		if i2 < 100000 {
			Total = Total + i2
		}
	}
	// part two
	TotalSpace := 70000000
	NeededSpace := 30000000
	NeededForUpdate := 0
	//for _, i2 := range FileSystem {
	//	UsedSpace = UsedSpace + i2
	//}
	var DirCanDel []string
	NeededForUpdate = NeededSpace - (TotalSpace - UsedSpace)
	for l, v := range TotalValueDir {
		if v >= NeededForUpdate {
			DirCanDel = append(DirCanDel, l)
		}
	}
	smallestDirName := ""
	smallestDirSize := 0
	for _, Dir := range DirCanDel {
		if smallestDirSize == 0 || smallestDirSize > TotalValueDir[Dir] {
			smallestDirName = Dir
			smallestDirSize = TotalValueDir[Dir]
		}
	}
	fmt.Println("dir smallest is :", smallestDirName)
}
