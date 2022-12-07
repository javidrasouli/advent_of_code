package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open("E:\\JavidProjects\\adventofcode\\dayOne\\input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	alfCalories := 0
	numberOfAlf := 1
	BiggestCalories := 0
	BestAlf := 0
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			Calorie, _ := strconv.Atoi(fileScanner.Text())
			alfCalories = alfCalories + Calorie
		} else {
			if BiggestCalories < alfCalories {
				BiggestCalories = alfCalories
				BestAlf = numberOfAlf
			}
			alfCalories = 0
			numberOfAlf = numberOfAlf + 1
		}
	}
	fmt.Println("the best of alf is number :", BestAlf)
	fmt.Println("the biggest calories was :", BiggestCalories)
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	//====================================== part two ===============================================
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err = os.Open("E:\\JavidProjects\\adventofcode\\dayOne\\input.txt")
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		return
	}
	var caloriesTotal []int
	calories := 0
	fileScanner = bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			Calorie, _ := strconv.Atoi(fileScanner.Text())
			calories = calories + Calorie
		} else {
			caloriesTotal = append(caloriesTotal, calories)
			calories = 0
		}
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	sort.Ints(caloriesTotal)
	total := caloriesTotal[len(caloriesTotal)-1] + caloriesTotal[len(caloriesTotal)-2] + caloriesTotal[len(caloriesTotal)-3]
	fmt.Println("total of three best alf calories :", total)
}
