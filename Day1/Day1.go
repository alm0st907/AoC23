package Day1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// remap strings because of edge cases like oneight, twone, etc.
var numberMap = map[string]string{
	"zero":  "zero0zero",
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

// simple forward search to get the first and last digit, but use first digit as both if only one
func Part1(line string) (combinedDigits int) {
	foundFirst := false
	var first, last int

	for _, char := range line {
		if unicode.IsDigit(char) {
			if !foundFirst {
				first = int(char - '0')
				foundFirst = true
			}
			last = int(char - '0')
		}
	}

	return first*10 + last
}

// use string replacement to leverage the work of initial solution
func Part2(line string) (combinedDigits int) {

	for word, digit := range numberMap {

		if strings.Contains(line, word) {
			line = strings.ReplaceAll(line, word, digit)
		}
	}
	combinedDigits = Part1(line)

	return combinedDigits
}

func RunDayOnePartOne(filepath string) (sum int) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	sum_part1 := 0

	for scanner.Scan() {
		line := scanner.Text()

		combinedDigits := Part1(line)
		sum_part1 += combinedDigits

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	return sum_part1
}

func RunDayOnePartTwo(filepath string) (sum int) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	sum_part2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		combinedDigits := Part2(line)
		sum_part2 += combinedDigits

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	return sum_part2
}

func RunDay1(filepath string) {

	sum_part1 := RunDayOnePartOne(filepath)
	sum_part2 := RunDayOnePartTwo(filepath)

	fmt.Printf("Sum of Part1: %d\n", sum_part1)
	fmt.Printf("Sum of Part2: %d\n", sum_part2)
}
