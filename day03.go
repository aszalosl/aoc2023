package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// read the input of the problem into a string array, and use a one-char wide border
func day03_read(fileName string) []string {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var table []string
	table = append(table, "")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		table = append(table, "."+str+".")
	}
	separator := strings.Repeat(".", len(table[1]))
	table = append(table, separator)
	table[0] = separator
	//for i, v := range table {
	//	fmt.Println(i, v)
	//}
	return table
}

// There is any special character around the number?
// t - table, l - line in it, s,e starting/end positions
func surround(t []string, l int, s int, e int) bool {
	line := t[l]
	if line[s] != '.' || line[e] != '.' {
		return true
	}
	line_above := t[l-1]
	line_below := t[l+1]
	for i := s; i <= e; i++ {
		if line_above[i] != '.' || line_below[i] != '.' {
			return true
		}
	}
	return false
}

// solve the first part
func day03_part1(fileName string) {
	table := day03_read(fileName)
	sum := 0
	for lineNo := 1; lineNo < len(table)-1; lineNo++ {
		line := table[lineNo]
		position := 1
		for {
			for position < len(line)-1 && (!unicode.IsDigit(rune(line[position]))) {
				position++
			}
			if position < len(line) && unicode.IsDigit(rune(line[position])) {
				value := 0
				starting_position := position - 1
				for unicode.IsDigit(rune(line[position])) {
					value = value*10 + int(line[position]) - 48
					position++
				}
				if surround(table, lineNo, starting_position, position) {
					sum += value
				}
				value = 0
				position++
			} else {
				break
			}
		}
	}
	fmt.Println("Part 1: ", sum)
}

// parse the number from the beginning of the string
func forward(line string) int {
	var numString string
	column := 0
	for unicode.IsDigit(rune(line[column])) {
		numString = numString + string(line[column])
		column++
	}
	value, _ := strconv.Atoi(numString)
	return value
}

// parse the number from the end of the string
func backward(line string) int {
	var num string
	column := len(line) - 1
	for unicode.IsDigit(rune(line[column])) {
		num = string(line[column]) + num
		column--
	}
	value, _ := strconv.Atoi(num)
	return value
}

// parse the number at the given column - go left at first, next use the `forward` fn
func atMiddle(line string, column int) int {
	for unicode.IsDigit(rune(line[column])) {
		column--
	}
	return forward(line[column+1:])
}

// Parse the numbers around the star. If exactly two numbers are here, give back their product, otherwise zero
func numbersAtStar(table []string, lineNo int, col int) int {
	var numbers = make([]int, 0)
	line := table[lineNo]
	if line[col] != '*' {
		fmt.Println("Wrong arguments")
	}
	//fmt.Println("line: ", line)

	// same line at right
	if unicode.IsDigit(rune(line[col+1])) {
		numbers = append(numbers, forward(line[col+1:]))
	}

	// same line at left
	if unicode.IsDigit(rune(line[col-1])) {
		numbers = append(numbers, backward(line[:col]))
	}

	// previous line
	prev_line := table[lineNo-1]
	// at
	if unicode.IsDigit(rune(prev_line[col])) {
		numbers = append(numbers, atMiddle(prev_line, col))
	} else {
		// after
		if unicode.IsDigit(rune(prev_line[col+1])) {
			numbers = append(numbers, forward(prev_line[col+1:]))
		}
		// before
		if unicode.IsDigit(rune(prev_line[col-1])) {
			numbers = append(numbers, backward(prev_line[:col]))
		}
	}
	// next line
	next_line := table[lineNo+1]
	// at
	if unicode.IsDigit(rune(next_line[col])) {
		numbers = append(numbers, atMiddle(next_line, col))
	} else {
		// after
		if unicode.IsDigit(rune(next_line[col+1])) {
			numbers = append(numbers, forward(next_line[col+1:]))
		}
		// before
		if unicode.IsDigit(rune(next_line[col-1])) {
			numbers = append(numbers, backward(next_line[:col]))
		}
	}

	//fmt.Println(numbers)
	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	} else {
		return 0
	}
}
func day03_part2(fileName string) {
	table := day03_read(fileName)
	sum := 0
	for lineNo := 1; lineNo < len(table)-1; lineNo++ {
		for i, ch := range table[lineNo] {
			if ch == '*' {
				sum += numbersAtStar(table, lineNo, i)
			}
		}
	}
	fmt.Println("Part 2: ", sum)
}
