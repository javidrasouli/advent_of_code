package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	boxes := make(map[string][]string)
	boxes["1"] = []string{"N", "D", "M", "Q", "B", "P", "Z"}
	boxes["2"] = []string{"C", "L", "Z", "Q", "M", "D", "H", "V"}
	boxes["3"] = []string{"Q", "H", "R", "D", "V", "F", "Z", "G"}
	boxes["4"] = []string{"H", "G", "D", "N", "F"}
	boxes["5"] = []string{"N", "F", "Q"}
	boxes["6"] = []string{"D", "Q", "V", "Z", "F", "B", "T"}
	boxes["7"] = []string{"Q", "M", "T", "Z", "D", "V", "S", "H"}
	boxes["8"] = []string{"M", "G", "F", "P", "N", "Q"}
	boxes["9"] = []string{"B", "W", "R", "M"}
	// part one
	file, err := os.Open("E:\\JavidProjects\\adventofcode\\day5\\input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	defer file.Close()
	for fileScanner.Scan() {
		arrStr := strings.Split(fileScanner.Text(), " ")
		move, _ := strconv.Atoi(arrStr[1])
		fromBox := arrStr[3]
		ToBox := arrStr[5]
		outed := boxes[fromBox][len(boxes[fromBox])-move:]
		boxes[fromBox] = boxes[fromBox][:len(boxes[fromBox])-move]
		//part one
		boxes[ToBox] = append(boxes[ToBox], reverseArr(outed)...)
		// part two
		boxes[ToBox] = append(boxes[ToBox], outed...)
	}
	var strKey string
	for i := 1; i < 10; i++ {
		strKey = strKey + boxes[strconv.Itoa(i)][len(boxes[strconv.Itoa(i)])-1]
	}
	fmt.Println("key is:", strKey)
}
func reverseArr(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseArr(input[1:]), input[0])
}
