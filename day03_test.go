package main

import "testing"

type starTest struct {
	table    []string
	column   int
	expected int
}

var testCases = []starTest{
	//starTest{[]string{".......", ".12*23..", "......."}, 3, 276},
	//starTest{[]string{".......", "...*234.", "......."}, 3, 0},
	//starTest{[]string{".......", ".12*....", "......."}, 3, 0},
	starTest{[]string{"...23..", "...*....", "..12..."}, 3, 276},
	starTest{[]string{".23....", "...*....", "..12..."}, 3, 276},
}

func TestNumbersAtStar(t *testing.T) {
	for _, test := range testCases {
		result := numbersAtStar(test.table, 1, test.column)
		if result != test.expected {
			t.Errorf("Output %d not equal with the expected %d", result, test.expected)
		}
	}
}
