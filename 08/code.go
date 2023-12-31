package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type dir struct {
	left  string
	right string
}

func readInput(fileName string) (string, map[string]dir) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	first := true
	var direction string
	graph := make(map[string]dir, 0)
	for scanner.Scan() {
		str := scanner.Text()
		if first {
			direction = str
			first = false
			continue
		}
		if str == "" {
			continue
		}
		pieces := strings.Split(str, " ")
		// the labels of the nodes consist of exactly 3 characters
		graph[pieces[0]] = dir{pieces[2][1:4], pieces[3][:3]}
	}
	return direction, graph
}

func part1(fileName string) {
	direction, graph := readInput(fileName)
	node := "AAA"
	counter := 0
	for {
		if node == "ZZZ" {
			fmt.Println("Part 1: ", counter)
			break
		}
		if direction[counter%len(direction)] == 'L' {
			node = graph[node].left
		} else {
			node = graph[node].right
		}
		counter++
	}
}
func cycle(node string, direction string, graph map[string]dir) int {
	counter := 0
	for {
		if node[2:] == "Z" {
			return counter
		}
		if direction[counter%len(direction)] == 'L' {
			node = graph[node].left
		} else {
			node = graph[node].right
		}
		counter++
	}
}

// from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func LCM(a, b int) int {
	result := a * b / GCD(a, b)
	return result
}

func part2(fileName string) {
	direction, graph := readInput(fileName)
	numbers := make([]int, 0)
	for k := range graph {
		if k[2:] == "A" {
			numbers = append(numbers, cycle(k, direction, graph))
		}
	}
	result := numbers[0]
	for _, n := range numbers[1:] {
		result = LCM(result, n)
	}
	fmt.Println("Part 2: ", result)
}
func main() {
	//input := "test08a.txt"
	//input := "test08b.txt"
	//input := "test08c.txt"
	input := "input08.txt"
	//d, g := readInput(input)
	//fmt.Println(d, g)
	//part1(input)
	part2(input)
}
