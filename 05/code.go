package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// we need to use intervals at filters and in the second part of the problem
type interval struct {
	begin int
	end   int
}

// we store the lines of the input in this structure.
// a filter consist of array/slice of filterItems
type filterItem struct {
	source      int
	destination int
	length      int
}

// Apply a filter on the `value`
func filterOne(filter []filterItem, value int) int {
	for _, item := range filter {
		if value >= item.source && value < item.source+item.length {
			return item.destination + (value - item.source)
		}
	}
	return value
}

// take the "intersection" of the filter-interval and the interval of the question
// the real intersection will be moved based on filterItem - we need to move at most once,
//
//	so we collected the moved/filtered part separated
//
// the real intersection is optional (0/1) we use a slice for it :-(
func slice(filter interval, x interval, distance int) ([]interval, []interval) {
	if filter.end < x.begin || x.end < filter.begin { // disjunct
		return []interval{x},
			[]interval{}
	}
	if filter.begin <= x.begin && x.end <= filter.end { // inner interval
		return []interval{},
			[]interval{{x.begin + distance, x.end + distance}}
	}
	if x.begin >= filter.begin { // bottom intersection
		return []interval{{filter.end + 1, x.end}},
			[]interval{{x.begin + distance, filter.end + distance}}
	}

	if filter.end < x.end { // inner filter
		return []interval{{x.begin, filter.begin - 1}, {filter.end + 1, x.end}},
			[]interval{{filter.begin + distance, filter.end + distance}}

	} else { // top intersection
		return []interval{{x.begin, filter.begin - 1}},
			[]interval{{filter.begin + distance, x.end + distance}}
	}
}

// the finction slice can produce several intervals, so we wrote this function more general
// we reuse the unfiltered/unmoved intervals for other filters, too
func filterInterval(filter []filterItem, intervals []interval) []interval {
	filtered := make([]interval, 0)
	for _, item := range filter {
		incomplete := make([]interval, 0)
		for _, spread := range intervals {
			later, done := slice(interval{item.source, item.source + item.length - 1}, spread, item.destination-item.source)
			filtered = append(filtered, done...)
			incomplete = append(incomplete, later...)
		}
		intervals = incomplete
	}
	filtered = append(filtered, intervals...)
	return filtered
}

// Apply a sequence of filters on a `value`
func filterChain(filters [][]filterItem, value int) int {
	for _, filter := range filters {
		value = filterOne(filter, value)
	}
	return value
}

// Apply a sequence of filters on an `interval`
func filterIntervalChain(filters [][]filterItem, intervals []interval) []interval {
	for _, filter := range filters {
		intervals = filterInterval(filter, intervals)
	}
	return intervals
}

// parse the integers from the first line
func processFirstLine(line string) []int {
	seedPieces := strings.Split(line[7:], " ")
	seeds := make([]int, 0, len(seedPieces))
	for _, str := range seedPieces {
		v, _ := strconv.Atoi(str)
		seeds = append(seeds, v)
	}
	return seeds
}

// interpret the whole input
func processInput(fileName string) ([]int, [][]filterItem) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	first := true
	var seeds []int
	var filters [][]filterItem = make([][]filterItem, 0, 7)
	var filter []filterItem = make([]filterItem, 0)
	for scanner.Scan() {
		str := scanner.Text()
		if first {
			// process the first line
			seeds = processFirstLine(str)
			first = false
			continue
		}
		if str == "" {
			continue
		}
		if strings.Contains(str, "map") { // description of the filter
			if len(filter) > 0 {
				filters = append(filters, filter)
				filter = make([]filterItem, 0)
			}
		} else { // typical data line
			items := strings.Split(str, " ")
			a, _ := strconv.Atoi(items[0])
			b, _ := strconv.Atoi(items[1])
			c, _ := strconv.Atoi(items[2])
			filter = append(filter, filterItem{b, a, c})
		}
	}
	filters = append(filters, filter) // the last one
	return seeds, filters
}

// first part of the problem
func part1(fileName string) {
	seeds, filters := processInput(fileName)
	output := make([]int, 0, len(seeds))
	for _, value := range seeds {
		v := filterChain(filters, value)
		output = append(output, v)
	}
	var mini int
	for i, v := range output {
		if i == 0 || v < mini {
			mini = v
		}
	}
	fmt.Println("Part 1: ", mini)
}

// check that the items in the filters are intersect each other, or not. -> no intersection
// func check3(fileName string) {
// 	_, filters := processInput(fileName)
// 	for _, filter := range filters {
// 		sort.SliceStable(filter, func(i, j int) bool {
// 			return filter[i].source < filter[j].source
// 		})
// 		for i := 1; i < len(filter); i++ {
// 			if filter[i-1].source+filter[i-1].length-1 >= filter[i].source {
// 				fmt.Println(filter[i-1], filter[i])
// 			}
// 		}
// 	}
// }

// second part of the problem
func part2(fileName string) {
	seeds, filters := processInput(fileName)

	output := make([]int, 0, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		intervals := make([]interval, 0)
		intervals = append(intervals, interval{seeds[i], seeds[i] + seeds[i+1]})
		result := filterIntervalChain(filters, intervals)
		var start int
		for i, v := range result {
			if i == 0 || v.begin < start {
				start = v.begin
			}
		}
		output = append(output, start)
	}
	var mini int
	for i, v := range output {
		if i == 0 || v < mini {
			mini = v
		}
	}
	fmt.Println("Part 2: ", mini)
}
func main() {
	//input := "test05.txt"
	input := "input05.txt"
	part1(input)
	part2(input)
	//check3(input)
}
