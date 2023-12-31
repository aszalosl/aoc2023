package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// one item of the input
type handBid struct {
	hand string
	bid  int
}
type Poker int8

const Strength1 = "AKQJT98765432"
const Strength2 = "AKQT98765432J"

const (
	FiveOfAKind Poker = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePar
	HighCard
)

// determine the value of the hand
func classify1(hand string) Poker {
	var repeat = make(map[byte]int, 0)
	for i := 0; i < 5; i++ {
		repeat[hand[i]] += 1
	}
	var s = ""
	for _, v := range repeat {
		s += strconv.Itoa(v)
	}
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	s = string(bs)

	return classifyType(s)
}

// Jack is the Joker
func classify2(hand string) Poker {
	var repeat = make(map[byte]int, 0)
	for i := 0; i < 5; i++ {
		repeat[hand[i]] += 1
	}
	if repeat['J'] > 0 { // we have Joker, find the best replacement
		maxk := byte('X')
		maxv := 0

		for k, v := range repeat {
			if k != 'J' && v > maxv {
				maxk = k
				maxv = v
			}
		}
		repeat[byte(maxk)] += repeat['J']
		delete(repeat, 'J')
	}

	var s = ""
	for k, v := range repeat {
		if k != 'J' {
			s += strconv.Itoa(v)
		}
	}
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	s = string(bs)

	return classifyType(s)
}
func classifyType(s string) Poker {
	switch s {
	case "5":
		return FiveOfAKind
	case "14":
		return FourOfAKind
	case "23":
		return FullHouse
	case "113":
		return ThreeOfAKind
	case "122":
		return TwoPair
	case "1112":
		return OnePar
	default:
		return HighCard
	}
}

// utility function for ordering
func less1(a, b handBid) bool {
	va := classify1(a.hand)
	vb := classify1(b.hand)

	if va > vb {
		return true
	}
	if vb > va {
		return false
	}
	for i := 0; i < 5; i++ {
		sa := strings.Index(Strength1, string(a.hand[i]))
		sb := strings.Index(Strength1, string(b.hand[i]))
		if sa > sb {
			return true
		}
		if sa < sb {
			return false
		}
	}
	return false
}
func less2(a, b handBid) bool {
	va := classify2(a.hand)
	vb := classify2(b.hand)

	if va > vb {
		return true
	}
	if vb > va {
		return false
	}
	for i := 0; i < 5; i++ {
		sa := strings.Index(Strength2, string(a.hand[i]))
		sb := strings.Index(Strength2, string(b.hand[i]))
		if sa > sb {
			return true
		}
		if sa < sb {
			return false
		}
	}
	return false
}

// load the whole file
func readInput(fileName string) []handBid {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	cards := make([]handBid, 0)
	for scanner.Scan() {
		str := scanner.Text()
		pieces := strings.Split(str, " ")
		n, _ := strconv.Atoi(pieces[1])
		cards = append(cards, handBid{pieces[0], n})
	}
	return cards
}

func part1(fileName string) {
	cs := readInput(fileName)
	sort.SliceStable(cs, func(i, j int) bool {
		return less1(cs[i], cs[j])
	})
	sum := 0
	for i, card := range cs {
		sum += (i + 1) * card.bid
	}
	fmt.Println("Part 1: ", sum)
}
func part2(fileName string) {
	cs := readInput(fileName)
	sort.SliceStable(cs, func(i, j int) bool {
		return less2(cs[i], cs[j])
	})
	sum := 0
	for i, card := range cs {
		sum += (i + 1) * card.bid
	}
	fmt.Println("Part 2: ", sum)
}

func main() {
	//input := "test07.txt"
	input := "input07.txt"
	part1(input)
	part2(input)

}
