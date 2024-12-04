package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var listOfIncreasing [][]int
var listOfDecreasing [][]int

func main() {
	inputTxt := "./input.txt"
	file, err := os.Open(inputTxt)
	if err != nil {
		fmt.Println("Error encountered opening file: ", err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var list [][]int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var lineOfInt []int
		//logic
		lengthOfLine := len(line)
		//is it increasing
		for i := 0; i < lengthOfLine; i++ {
			//convert to int
			num, err := strconv.Atoi(line[i])
			if err != nil {
				fmt.Println("Error converting to Integer: ", err.Error())
			}
			lineOfInt = append(lineOfInt, num)

		}
		list = append(list, lineOfInt)

	}
	listOfIncreasing, listOfDecreasing = isIncreasingOrDecreasing(list)
	safeList := checkTheDifference(listOfIncreasing, listOfDecreasing)

	fmt.Println("Answer to Day 2 ", len(safeList))

}

func isIncreasingOrDecreasing(list [][]int) ([][]int, [][]int) {

	for _, value := range list {
		listLength := len(value) - 1
		if value[0] < value[listLength] {
			listOfIncreasing = append(listOfIncreasing, value)
		}
		if value[0] > value[listLength] {
			listOfDecreasing = append(listOfDecreasing, value)
		}

	}

	return listOfIncreasing, listOfDecreasing
}
func checkTheDifference(increasingList [][]int, decreasingList [][]int) [][]int {
	var safeList [][]int

	for _, value := range increasingList {
		safe := true
		errorCount := 0
		for i := 0; i < len(value)-1; i++ {
			if safe || errorCount < 2 {
				if value[i+1]-value[i] >= 1 && value[i+1]-value[i] <= 3 {
					safe = true
				} else {
					safe = false
					errorCount++
				}
			}
		}
		if safe {
			safeList = append(safeList, value)
		}

	}

	for _, value := range decreasingList {
		safe := true
		errorCount := 0
		for i := 0; i < len(value)-1; i++ {
			if safe || errorCount < 2 {
				if value[i+1]-value[i] >= -3 && value[i+1]-value[i] <= -1 {
					safe = true
				} else {
					safe = false
					errorCount++
				}
			}
		}
		if safe {

			safeList = append(safeList, value)
		}

	}

	return safeList
}
