package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part one
	file, err := os.Open("E:\\JavidProjects\\adventofcode\\day4\\input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	defer file.Close()
	fully_contain := 0
	for fileScanner.Scan() {
		str := fileScanner.Text()
		elfGroup := strings.Split(str, ",")
		Elf1 := strings.Split(elfGroup[0], "-")
		Elf2 := strings.Split(elfGroup[1], "-")
		elf1Start, _ := strconv.Atoi(Elf1[0])
		elf1Finisher, _ := strconv.Atoi(Elf1[1])
		elf2Start, _ := strconv.Atoi(Elf2[0])
		elf2Finisher, _ := strconv.Atoi(Elf2[1])
		//part one
		if elf1Start <= elf2Start && elf1Finisher >= elf2Finisher {
			fully_contain = fully_contain + 1
		} else if elf2Start <= elf1Start && elf2Finisher >= elf1Finisher {
			fully_contain = fully_contain + 1
		}
		// part two
		if elf1Finisher >= elf2Start && elf1Start <= elf2Start {
			fully_contain = fully_contain + 1
		} else if elf2Finisher >= elf1Start && elf2Start <= elf1Start {
			fully_contain = fully_contain + 1
		}
	}
	fmt.Println("full count :", fully_contain)
}
