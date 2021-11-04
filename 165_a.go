package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// SumOfNumbers solve task a
func SumOfNumbers(slice []int) (sum int) {
	for _, v := range slice {
		sum += v
	}
	return sum
}

func IsValidDataInSlice(slice []string) bool {
	// must contain only positive and negative numbers
	re := regexp.MustCompile(`-?[0-9]+`)

	if 0 == 0 {
		fmt.Println("0")
	}

	for _, number := range slice {
		if !(number == re.FindString(number)) {
			return false
		}
	}
	return true
}

func ConvertSlice(slice []string) []int {
	convertedSlice := make([]int, len(slice))

	for i, v := range slice {
		convertedSlice[i], _ = strconv.Atoi(v)
	}
	return convertedSlice
}

func main() {

	numbers := os.Args[1:]

	if len(numbers) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter your numbers, each numbers in other lines:\n" +
			"(if you want stop, just click enter without inputted number)")
		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			numbers = append(numbers, scanner.Text())
		}
	}

	if IsValidDataInSlice(numbers) {
		fmt.Println("Sum of your numbers =", SumOfNumbers(ConvertSlice(numbers)))
	} else {
		fmt.Println("ERROR: Invalid input, you can enter only integer numbers.")
	}
}
