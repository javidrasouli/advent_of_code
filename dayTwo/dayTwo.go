package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// part one
	file, err := os.Open("E:\\JavidProjects\\adventofcode\\dayTwo\\input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	scoreMap := make(map[string]int)
	scoreMap["A"] = 1
	scoreMap["B"] = 2
	scoreMap["C"] = 3
	LosserMap := make(map[string]string)
	LosserMap["A"] = "B"
	LosserMap["B"] = "C"
	LosserMap["C"] = "A"
	sameMap := make(map[string]string)
	sameMap["A"] = "X"
	sameMap["B"] = "Y"
	sameMap["C"] = "Z"
	totalPoint := 0
	for fileScanner.Scan() {
		round := fileScanner.Text()
		roundArr := strings.Split(round, " ")
		totalPoint = totalPoint + scoreMap[roundArr[1]]
		if roundArr[1] == LosserMap[roundArr[0]] {
			totalPoint = totalPoint + 6
		}
		if roundArr[1] == sameMap[roundArr[0]] {
			totalPoint = totalPoint + 3
		}
	}
	println("total point : ", totalPoint)
	// part two
	scoreMap = make(map[string]int)
	scoreMap["A"] = 1
	scoreMap["B"] = 2
	scoreMap["C"] = 3
	LosserMap = make(map[string]string)
	LosserMap["A"] = "B"
	LosserMap["B"] = "C"
	LosserMap["C"] = "A"
	WinnerMap := make(map[string]string)
	WinnerMap["A"] = "C"
	WinnerMap["B"] = "A"
	WinnerMap["C"] = "B"
	ActionRound := make(map[string]string)
	ActionRound["X"] = "Loss"
	ActionRound["Y"] = "Draw"
	ActionRound["Z"] = "Win"
	totalPoint = 0
	for fileScanner.Scan() {
		round := fileScanner.Text()
		roundArr := strings.Split(round, " ")
		switch ActionRound[roundArr[1]] {
		case "Loss":
			totalPoint = totalPoint + scoreMap[WinnerMap[roundArr[0]]]
		case "Draw":
			totalPoint = totalPoint + scoreMap[roundArr[0]]
			totalPoint = totalPoint + 3
		case "Win":
			totalPoint = totalPoint + scoreMap[LosserMap[roundArr[0]]]
			totalPoint = totalPoint + 6
		}
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
}
