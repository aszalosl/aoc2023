package main

import (
	"testing"
)

type testCase struct {
	fileName string
	expected int
}

func TestPart1(t *testing.T) {
	testCases := []testCase{{"test10a.txt", 4}, {"test10b.txt", 8}}
	for _, test := range testCases {
		result := part1(test.fileName)
		if result != test.expected {
			t.Errorf("Output %d not equal with the expected %d for ", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []testCase{
		{"test10a.txt", 1},
		{"test10b.txt", 1},
		{"test10c.txt", 4},
		{"test10d.txt", 4},
		{"test10e.txt", 8},
		{"test10f.txt", 10}}
	for _, test := range testCases {
		result := part2(test.fileName)
		if result != test.expected {
			t.Errorf("Output %d not equal with the expected %d for ", result, test.expected)
		}
	}
}
