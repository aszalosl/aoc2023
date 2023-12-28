// Not a deterministic solution.  It selects two nodes randomly and tries to arrange the nodes.
// It it is successful, we get the answer, otherwise try again.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Checks the number of edges crossing the zero
func cross(g map[string]string, p map[string]float64) int {
	counter := 0
	for key, value := range g {
		neighbours := strings.Split(value, ":")
		for _, n := range neighbours {
			if p[key]*p[n] < 0.0 {
				//fmt.Println(key, n)
				counter++
			}
		}
	}
	return counter / 2
}

// Moves the nodes of the graph based on their neighbours' positions.
// g - the graph, p - 1D position of nodes, f/l the two fixed nodes
func move(g map[string]string, p map[string]float64, f string, l string) {
	for key, value := range g {
		if key != f && key != l {
			neighbours := strings.Split(value, ":")
			sum := 0.0
			for _, n := range neighbours {
				sum += p[n]
			}
			p[key] = sum / float64(len(neighbours))
		}
	}
}

// https://adventofcode.com/2023/day/25
func part1(fileName string) {

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var graph = map[string]string{}
	// read the input
	for scanner.Scan() {
		str := scanner.Text()
		//fmt.Println("\n", "input line: ", str)
		words := strings.Split(str, " ")
		id := strings.TrimSuffix(words[0], ":")
		for index := 1; index < len(words); index += 1 {
			graph[id] += words[index] + ":"
			graph[words[index]] += id + ":"
		}
		//for key, value := range graph {
		//	fmt.Println(key, value)
		//}
	}

	//
	var keys = make([]string, len(graph))
	counter := 0
	var position = map[string]float64{}
	for key := range graph {
		keys[counter] = key
		counter++
		position[key] = 0.0
	}
	var first = keys[0]
	var last = keys[1+rand.Intn(len(graph)-1)]
	position[first] = -100
	position[last] = 100
	for i := 1; i <= 10; i++ {
		move(graph, position, first, last)
	}

	// calculate the number of nodes on both sides
	positive, negative := 0, 0
	for _, value := range position {
		if value < 0.0 {
			negative++
		} else {
			positive++
		}
	}
	if cross(graph, position) == 3 {
		fmt.Println(positive * negative)
	} else {
		fmt.Println("Run me again.")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	input := "input25.txt"
	part1(input)
}
