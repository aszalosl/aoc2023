package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func processLine(s string, firstPart bool) int {
	colon := strings.Index(s, ":")
	bar := strings.Index(s, "|")
	winning := strings.Trim(s[colon+1:bar-1], " ")
	first := strings.Split(winning, " ")
	// store the winning numbers in a map
	var firstMap = make(map[string]int)
	for _, value := range first {
		if value != "" {
			firstMap[value] = 1
		}
	}
	counter := 0
	myNumbers := strings.Trim(s[bar+1:], " ")
	second := strings.Split(myNumbers, " ")
	// if my some number is in the map, increment the counter, otherwise add zero to it
	for _, value := range second {
		counter += firstMap[value]
	}
	// return back the corresponding point
	if counter == 0 {
		return 0
	} else {
		if firstPart {
			return int(math.Pow(2, float64(counter)-1.0))
		} else {
			return counter
		}
	}
}

// first part
func part1(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0
	// process all the lines
	for scanner.Scan() {
		str := scanner.Text()
		sum += processLine(str, true)

	}
	fmt.Println("Part 1: ", sum)
}
func part2(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	var sum uint64 = 0
	counter := 1
	var multiplicity = make(map[int]uint64)
	// process all the lines
	for scanner.Scan() {
		str := scanner.Text()
		multiplicity[counter] += 1
		copies := processLine(str, false)
		if copies > 0 {
			for i := 1; i <= copies; i++ {
				multiplicity[counter+i] += multiplicity[counter]
			}
		}
		counter++
	}

	for _, value := range multiplicity {
		sum += value
	}
	fmt.Println("Part 2: ", sum)
}
func main() {
	//input := "test04.txt"
	input := "input04.txt"
	part1(input)
	part2(input)
}
