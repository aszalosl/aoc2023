package main

import (
	"testing"
)

type handType struct {
	hand  string
	value Poker
}

func TestClassify1(t *testing.T) {
	testCases := []handType{
		{"32T3K", OnePar},
		{"T55J5", ThreeOfAKind},
		{"KK677", TwoPair},
		{"KTJJT", TwoPair},
		{"QQQJA", ThreeOfAKind}}

	for _, test := range testCases {
		result := classify1(test.hand)
		if result != test.value {
			t.Errorf("Output %v not equal with the expected %v", result, test.value)
		}
	}
}
func TestClassify2(t *testing.T) {
	testCases := []handType{
		{"32T3K", OnePar},
		{"T55J5", FourOfAKind},
		{"KK677", TwoPair},
		{"KTJJT", FourOfAKind},
		{"QQQJA", FourOfAKind}}

	for _, test := range testCases {
		result := classify2(test.hand)
		if result != test.value {
			t.Errorf("Output %v not equal with the expected %v at %v", result, test.value, test.hand)
		}
	}
}

func TestLess(t *testing.T) {
	cards := []handBid{{"32T3K", 765}, {"T55J5", 684}, {"KK677", 28}, {"KTJJT", 220}, {"QQQJA", 483}}
	for _, pair := range [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {3, 2}, {1, 4}, {2, 1}, {3, 1}} {
		a, b := cards[pair[0]], cards[pair[1]]
		//t.Errorf("%v < %v\n", a, b)
		if !less1(a, b) {
			t.Errorf("%v (%v) is not weaker than %v (%v)", a, classify1(a.hand), b, classify1(b.hand))
		}
	}
	for _, pair := range [][]int{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {2, 3}, {4, 1}, {1, 2}, {1, 3}} {
		a, b := cards[pair[0]], cards[pair[1]]
		//t.Errorf("%v < %v\n", a, b)
		if less1(a, b) {
			t.Errorf("%v (%v) is not stronger than %v (%v)", a, classify1(a.hand), b, classify1(b.hand))
		}
	}

}
