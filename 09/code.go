package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	sumOfSums := 0
	for scanner.Scan() {
		//convert the next line into a numberic array
		str := scanner.Text()
		numStrings := strings.Split(str, " ")
		numbers := make([]int, 0, len(numStrings))
		for _, num := range numStrings {
			value, _ := strconv.Atoi(num)
			numbers = append(numbers, value)
		}

		// calculate the differences again and again
		levels := make([][]int, 0)
		levels = append(levels, numbers)
		counter := 1
		for {
			level := make([]int, 0, len(numbers)-counter)
			for i := 0; i < len(numbers)-counter; i++ {
				lastSequence := levels[len(levels)-1]
				level = append(level, lastSequence[i+1]-lastSequence[i])
			}
			allzero := true
			for i := 0; i < len(numbers)-counter; i++ {
				if level[i] != 0 {
					allzero = false
				}
			}
			if allzero {
				break
			}
			levels = append(levels, level)
			counter++
		}

		// sum the last numbers
		sum := 0
		for _, line := range levels {
			sum += line[len(line)-1]
		}
		sumOfSums += sum
	}
	fmt.Println("Part 1:", sumOfSums)
}

func part2(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	sumOfSums := 0
	for scanner.Scan() {
		//convert the next line into a numberic array
		str := scanner.Text()
		numStrings := strings.Split(str, " ")
		numbers := make([]int, 0, len(numStrings))
		for _, num := range numStrings {
			value, _ := strconv.Atoi(num)
			numbers = append([]int{value}, numbers...) // append in reverse direction, and do the same
		}

		// calculate the differences again and again
		levels := make([][]int, 0)
		levels = append(levels, numbers)
		counter := 1
		for {
			level := make([]int, 0, len(numbers)-counter)
			for i := 0; i < len(numbers)-counter; i++ {
				lastSequence := levels[len(levels)-1]
				level = append(level, lastSequence[i+1]-lastSequence[i])
			}
			allzero := true
			for i := 0; i < len(numbers)-counter; i++ {
				if level[i] != 0 {
					allzero = false
				}
			}
			if allzero {
				break
			}
			levels = append(levels, level)
			counter++
		}

		// sum the last numbers
		sum := 0
		for _, line := range levels {
			sum += line[len(line)-1]
		}
		sumOfSums += sum
	}
	fmt.Println("Part 2:", sumOfSums)
}

func main() {
	//input := "test09.txt"
	input := "input09.txt"
	//part1(input)
	part2(input)
}
