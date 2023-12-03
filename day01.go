package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// first part of the problem
func day01_part1(fileName string) {

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	first := 0
	last := 0
	sum := 0

	for scanner.Scan() {

		str := scanner.Text()
		for index := 0; index < len(str); index++ {
			ch := str[index]
			if '0' <= ch && ch <= '9' {
				first = int(ch) - int('0')
				break
			}
		}
		for index := len(str) - 1; index >= 0; index-- {
			ch := str[index]
			if '0' <= ch && ch <= '9' {
				last = int(ch) - int('0')
				break
			}
		}
		sum += 10*first + last

	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// there is a "digit" at index 'index' in the string 'str'? Which one?
func check01(str string, index int) int {
	l := len(str)
	if l >= index+3 && str[index:index+3] == "one" {
		return 1
	}
	if l >= index+3 && str[index:index+3] == "two" {
		return 2
	}
	if l >= index+5 && str[index:index+5] == "three" {
		return 3
	}
	if l >= index+4 && str[index:index+4] == "four" {
		return 4
	}
	if l >= index+4 && str[index:index+4] == "five" {
		return 5
	}
	if l >= index+3 && str[index:index+3] == "six" {
		return 6
	}
	if l >= index+5 && str[index:index+5] == "seven" {
		return 7
	}
	if l >= index+5 && str[index:index+5] == "eight" {
		return 8
	}
	if l >= index+4 && str[index:index+4] == "nine" {
		return 9
	}
	return 0
}

// second part of the problem
func day01_part2(fileName string) {

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	first := 0
	last := 0
	sum := 0

	for scanner.Scan() {

		str := scanner.Text()
		for index := 0; index < len(str); index++ {
			chk := check01(str, index)
			if chk > 0 {
				first = chk
				break
			}
			ch := str[index]
			if '0' <= ch && ch <= '9' {
				first = int(ch) - int('0')
				break
			}
		}
		for index := len(str) - 1; index >= 0; index-- {
			chk := check01(str, index)
			if chk > 0 {
				last = chk
				break
			}
			ch := str[index]
			if '0' <= ch && ch <= '9' {
				last = int(ch) - int('0')
				break
			}
		}
		sum += 10*first + last

	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
