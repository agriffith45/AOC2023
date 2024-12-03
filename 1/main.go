package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "./input.txt"
	f, err := os.Open(input)
	msg := "Encountered error Reading file: "
	checkError(msg, err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines = make([]string, 0)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		lines = append(lines, line...)
	}
	evens, odds := splitIntoOddsAndEvens(lines)
	answer := doTheMath(evens, odds)
	fmt.Println("the answer is: ", answer)

}

func splitIntoOddsAndEvens(arg1 []string) (evens []int, odds []int) {
	for index, value := range arg1 {
		i, err := strconv.Atoi(value)
		message := "error encountered converting string to integer"
		checkError(message, err)
		if index%2 == 0 {
			evens = append(evens, i)
		} else {
			odds = append(odds, i)
		}
	}
	return evens, odds
}

func doTheMath(evens []int, odds []int) int {
	sort.Ints(evens)
	sort.Ints(odds)

	if len(evens) != len(odds) {
		message := "The lists don't match"
		err := fmt.Errorf("The length of each list is not equal")
		checkError(message, err)

	}
	var b []int
	for i := 0; i < len(evens); i++ {
		a := math.Abs(float64(evens[i]) - float64(odds[i]))
		b = append(b, int(a))
	}
	sum := 0
	for _, num := range b {
		sum += num
	}

	return sum

}
func checkError(message string, errorM error) {
	if errorM != nil {
		fmt.Sprintf("%s, %s", message, errorM)
	}
}
