package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// first part of the problem
func part1(fileName string) {

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	id := 0  // the id of the actual game
	sum := 0 // sum of the id's

	for scanner.Scan() {

		str := scanner.Text()
		words := strings.Split(str, " ")
		id, err = strconv.Atoi(strings.TrimSuffix(words[1], ":"))
		if err != nil {
			panic(err)
		}
		possible := true
		for index := 2; index < len(words); index += 2 {
			count, err := strconv.Atoi(words[index])
			if err != nil {
				panic(err)
			}
			colour := strings.TrimSuffix(words[index+1], ",")
			colour = strings.TrimSuffix(colour, ";")
			if colour == "red" && count > 12 {
				possible = false
				break
			}
			if colour == "green" && count > 13 {
				possible = false
				break
			}
			if colour == "blue" && count > 14 {
				possible = false
				break
			}
		}

		if possible {
			sum += id
		}

	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// second part of the problem
func part2(fileName string) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0 // sum of the products

	for scanner.Scan() {
		str := scanner.Text()
		words := strings.Split(str, " ")
		reds := 0
		blues := 0
		greens := 0
		for index := 2; index < len(words); index += 2 {
			count, err := strconv.Atoi(words[index])
			if err != nil {
				panic(err)
			}
			colour := strings.TrimSuffix(words[index+1], ",")
			colour = strings.TrimSuffix(colour, ";")
			if colour == "red" && count > reds {
				reds = count
			}
			if colour == "green" && count > greens {
				greens = count
			}
			if colour == "blue" && count > blues {
				blues = count
			}
		}
		sum += reds * greens * blues

	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	input := "input02.txt"
	part1(input)
	part2(input)
}
