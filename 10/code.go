package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type direction int

const (
	Nowhere direction = iota
	North
	East
	South
	West
	Inner // for debugging only
)

type fromDir struct {
	up direction
	dn direction
	rh direction
	lf direction
}

var nextDir = map[byte]fromDir{
	'|': {North, South, Nowhere, Nowhere},
	'-': {Nowhere, Nowhere, East, West},
	'L': {Nowhere, East, Nowhere, North},
	'J': {Nowhere, West, North, Nowhere},
	'7': {West, Nowhere, South, Nowhere},
	'F': {East, Nowhere, Nowhere, South},
	'.': {Nowhere, Nowhere, Nowhere, Nowhere},
	'I': {Nowhere, Nowhere, Nowhere, Nowhere}, // for examples of part 2
	'O': {Nowhere, Nowhere, Nowhere, Nowhere}, //
}

func drawTheLine(desertMap []string, line int, col int, dir direction) int {
	// create a matrix same size as desertMap
	otherMap := make([][]direction, len(desertMap))
	for i := range desertMap {
		otherMap[i] = make([]direction, len(desertMap[0]))
	}
	otherMap[line][col] = dir

	// mark the path
pathsearch: // label for breaking out from the outer loop
	for {
		switch dir {
		case North:
			if line == 0 { // step off the board
				dir = Nowhere
				break pathsearch
			} else {
				line--
				otherMap[line][col] = North
				if desertMap[line][col] == 'S' {
					break pathsearch
				} else {
					dir = nextDir[desertMap[line][col]].up
				}
			}
		case East:
			if col == len(desertMap[0])-1 {
				dir = Nowhere // step off the board
				break pathsearch
			} else {
				col++
				if desertMap[line][col] == 'S' {
					break pathsearch
				} else {
					dir = nextDir[desertMap[line][col]].rh
				}
				if dir == North || dir == South {
					otherMap[line][col] = dir
				} else {
					otherMap[line][col] = East
				}
			}
		case South:
			if line == len(desertMap)-1 {
				dir = Nowhere // step off the board
				break pathsearch
			} else {
				line++
				otherMap[line][col] = South
				if desertMap[line][col] == 'S' {
					break pathsearch
				} else {
					dir = nextDir[desertMap[line][col]].dn
				}
			}
		case West:
			if col == 0 {
				dir = Nowhere // step off the board
				break pathsearch
			} else {
				col--
				if desertMap[line][col] == 'S' {
					break pathsearch
				} else {
					dir = nextDir[desertMap[line][col]].lf
				}
				if dir == North || dir == South {
					otherMap[line][col] = dir
				} else {
					otherMap[line][col] = West
				}
			}
		case Nowhere:
			break pathsearch
		}

	}
	// if there is no loop, return 0
	if dir == Nowhere {
		return 0
	}

	// otherwise count the inner positions
	counter := 0
	for i := 1; i < len(desertMap)-1; i++ { // the first and last line cannot contain inner position
		status := false
		last := Nowhere
		for j := 0; j < len(desertMap[0]); j++ {
			if (otherMap[i][j] == North || otherMap[i][j] == South) && otherMap[i][j] != last {
				status = !status
				last = otherMap[i][j]
			}
			if otherMap[i][j] == Nowhere && status {
				if status {
					counter++
					otherMap[i][j] = Inner
				}
			}
		}
	}
	// for debugging
	// for i := 1; i < len(desertMap)-1; i++ {
	// 	for j := 0; j < len(desertMap[0]); j++ {
	// 		fmt.Printf("%d", otherMap[i][j])
	// 	}
	// 	fmt.Println()
	// }
	return counter
}

// start from S, go direction `dir` return back the length of the loop, or 0
func followTheLine(desertMap []string, line int, col int, dir direction) int {
	counter := 0
	for {
		//fmt.Println(line, col, dir, desertMap[line][col])
		switch dir {
		case North:
			if line == 0 {
				return 0 // step off the board
			} else {
				counter++
				line--
				if desertMap[line][col] == 'S' {
					return counter
				} else {
					dir = nextDir[desertMap[line][col]].up
				}
			}
		case East:
			if col == len(desertMap[0])-1 {
				return 0 // step off the board
			} else {
				col++
				counter++
				if desertMap[line][col] == 'S' {
					return counter
				} else {
					dir = nextDir[desertMap[line][col]].rh
				}
			}
		case South:
			if line == len(desertMap)-1 {
				return 0 // step off the board
			} else {
				counter++
				line++
				if desertMap[line][col] == 'S' {
					return counter
				} else {
					dir = nextDir[desertMap[line][col]].dn
				}
			}
		case West:
			if col == 0 {
				return 0 // step off the board
			} else {
				col--
				counter++
				if desertMap[line][col] == 'S' {
					return counter
				} else {
					dir = nextDir[desertMap[line][col]].lf
				}
			}
		case Nowhere:
			return 0
		}
	}
}
func part1(fileName string) int {
	// load the map
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	desertMap := make([]string, 0)
	var startLine int
	counter := 0
	for scanner.Scan() {
		str := scanner.Text()
		if strings.Contains(str, "S") {
			startLine = counter
		}
		desertMap = append(desertMap, str)
		counter++
	}
	startColumn := strings.Index(desertMap[startLine], "S")
	var length int
	length = followTheLine(desertMap, startLine, startColumn, North)
	if length > 0 {
		return length / 2
	} else {
		length = followTheLine(desertMap, startLine, startColumn, East)
		if length > 0 {
			return length / 2
		} else {
			length = followTheLine(desertMap, startLine, startColumn, South)
			if length > 0 {
				return length / 2
			} else {
				return 0
			}
		}
	}
}
func part2(fileName string) int {
	// load the map
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	desertMap := make([]string, 0)
	var startLine int
	counter := 0
	for scanner.Scan() {
		str := scanner.Text()
		if strings.Contains(str, "S") {
			startLine = counter
		}
		desertMap = append(desertMap, str)
		counter++
	}
	startColumn := strings.Index(desertMap[startLine], "S")
	var length int
	length = drawTheLine(desertMap, startLine, startColumn, North)
	if length > 0 {
		return length
	} else {
		length = drawTheLine(desertMap, startLine, startColumn, East)
		if length > 0 {
			return length
		} else {
			length = drawTheLine(desertMap, startLine, startColumn, South)
			if length > 0 {
				return length
			} else {
				return 0
			}
		}
	}
}
func main() {
	input := "input10.txt"
	//fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
