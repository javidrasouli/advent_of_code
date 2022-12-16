package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// part one
	file, err := os.Open("E:\\JavidProjects\\adventofcode\\day3\\input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var arr = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11,
		"l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20,
		"u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31,
		"F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39,
		"N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47,
		"V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}
	sumPoint := 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		part1 := text[:len(text)/2]
		part2 := text[len(text)/2:]
		samechar := []string{}
		fmt.Println("===============================================================")
		fmt.Println(strings.SplitAfter(part1, ""))
		fmt.Println(strings.SplitAfter(part2, ""))
		for _, s := range strings.SplitAfter(part1, "") {
			for _, s2 := range strings.SplitAfter(part2, "") {
				if s == s2 {
					addBefore := false
					for _, s3 := range samechar {
						if s == s3 {
							addBefore = true
						}
					}
					if !addBefore {
						samechar = append(samechar, s)
					}
				}
			}
		}
		for _, s2 := range samechar {
			sumPoint = sumPoint + arr[s2]
		}
		fmt.Println("same :", samechar)
		fmt.Println("===============================================================")
	}
	fmt.Println("sum point :", sumPoint)
	// part two
	var alfGroup []string
	for fileScanner.Scan() {
		if len(alfGroup) < 3 {
			alfGroup = append(alfGroup, fileScanner.Text())
		}
		if len(alfGroup) == 3 {
			var finded string
			for _, i := range strings.SplitAfter(alfGroup[0], "") {
				alf2 := strings.Index(alfGroup[1], i)
				alf3 := strings.Index(alfGroup[2], i)
				if alf3 != -1 && alf2 != -1 {
					finded = i
				}
			}
			sumPoint = sumPoint + arr[finded]
			alfGroup = []string{}
			finded = ""
		}
	}
	fmt.Println("sum point :", sumPoint)
}
