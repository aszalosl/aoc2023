package main

import (
	"sort"
	"testing"
)

var filter1 = []filterItem{{98, 50, 2}, {50, 52, 48}}                        // seed-to-soil
var filter2 = []filterItem{{15, 0, 37}, {52, 37, 2}, {0, 39, 15}}            // soil-to-fertilizer
var filter3 = []filterItem{{53, 49, 8}, {11, 0, 42}, {0, 42, 7}, {7, 57, 4}} // fertilizer-to-water
var filter4 = []filterItem{{18, 88, 7}, {25, 18, 70}}                        // water-to-light
var filter5 = []filterItem{{77, 45, 23}, {45, 81, 19}, {64, 68, 13}}         // light-to-temperature
var filter6 = []filterItem{{69, 0, 1}, {0, 1, 69}}                           // temperature-to-humidity
var filter7 = []filterItem{{56, 60, 37}, {93, 56, 4}}                        // humidity-to-location

var filters = [][]filterItem{filter1, filter2, filter3, filter4, filter5, filter6, filter7}

type testCase struct {
	input    int
	expected int
}

func different(a, b []interval) bool {
	if len(a) != len(b) {
		return true
	}
	sort.SliceStable(a, func(i, j int) bool { return a[i].begin < a[j].begin })
	sort.SliceStable(b, func(i, j int) bool { return b[i].begin < b[j].begin })
	for i := 0; i < len(a); i++ {
		if a[i].begin != b[i].begin || a[i].end != b[i].end {
			return true
		}
	}
	return false // just permutation of each other
}
func TestFilterOne(t *testing.T) {
	var filterCases = []testCase{{79, 81}, {14, 14}, {55, 57}, {13, 13}}
	for _, test := range filterCases {
		result := filterOne(filter1, test.input)
		if result != test.expected {
			t.Errorf("Output %d not equal with the expected %d", result, test.expected)
		}
	}
}

func TestFilterChain(t *testing.T) {
	var filterCases = []testCase{{79, 82}, {14, 43}, {55, 86}, {13, 35}}
	for _, test := range filterCases {
		result := filterChain(filters, test.input)
		if result != test.expected {
			t.Errorf("Output %d not equal with the expected %d", result, test.expected)
		}
	}
}

// func TestIntervalIntersection(t *testing.T) {
// 	var testCases = [][]int{{1, 2, 3, 4, 0, 0}, {5, 6, 2, 3, 0, 0}, {1, 3, 2, 4, 2, 3}, {2, 4, 1, 3, 2, 3}, {2, 3, 1, 4, 2, 3}, {1, 4, 2, 3, 2, 3}}
// 	for _, test := range testCases {
// 		result := intervalIntersection(interval{test[0], test[1]}, interval{test[2], test[3]})
// 		if result.begin != test[4] || result.end != test[5] {
// 			t.Errorf("Output (%d,%d) not equal with the expected (%d, %d)", result.begin, result.end, test[4], test[5])
// 		}
// 	}
// }

func TestFilterInterval1(t *testing.T) {
	result := filterInterval(filter1, []interval{{79, 92}, {55, 67}})
	expected := []interval{{81, 94}, {57, 69}}
	if different(result, expected) {
		t.Errorf("Filter1: %v != %v ", result, expected)
	}
}
func TestFilterInterval2(t *testing.T) {
	result := filterInterval(filter2, []interval{{81, 94}, {57, 69}})
	expected := []interval{{81, 94}, {57, 69}}
	if different(result, expected) {
		t.Errorf("Filter1: %v != %v ", result, expected)
	}
}
func TestFilterInterval3(t *testing.T) {
	result := filterInterval(filter3, []interval{{81, 94}, {57, 69}})
	expected := []interval{{81, 94}, {53, 56}, {61, 69}}
	if different(result, expected) {
		t.Errorf("Filter1: %v != %v ", result, expected)
	}
}
func TestFilterInterval4(t *testing.T) {
	result := filterInterval(filter4, []interval{{81, 94}, {53, 56}, {61, 69}})
	expected := []interval{{74, 87}, {46, 49}, {54, 62}}
	if different(result, expected) {
		t.Errorf("Filter1: %v != %v ", result, expected)
	}
}
func TestFilterInterval5(t *testing.T) {
	result := filterInterval(filter5, []interval{{74, 87}, {46, 49}, {54, 62}})
	expected := []interval{{78, 80}, {45, 55}, {82, 85}, {90, 98}}
	if different(result, expected) {
		t.Errorf("Filter1: %v != %v ", result, expected)
	}
}
func TestFilterInterval6(t *testing.T) {
	result := filterInterval(filter6, []interval{{78, 80}, {45, 55}, {82, 85}, {90, 98}})
	expected := []interval{{78, 80}, {46, 56}, {82, 85}, {90, 98}}
	if different(result, expected) {
		t.Errorf("Filter1: %v != %v ", result, expected)
	}
}
func TestFilterInterval7(t *testing.T) {
	result := filterInterval(filter7, []interval{{78, 80}, {46, 56}, {82, 85}, {90, 98}})
	expected := []interval{{82, 84}, {46, 55}, {60, 60}, {86, 89}, {94, 96}, {56, 59}, {97, 98}}
	if different(result, expected) {
		t.Errorf("Filter1: %v != %v ", result, expected)
	}
}
func TestFilterIntervalChain(t *testing.T) {
	result := filterIntervalChain(filters, []interval{{79, 92}, {55, 67}})
	expected := []interval{{82, 84}, {46, 55}, {60, 60}, {86, 89}, {94, 96}, {56, 59}, {97, 98}}
	if different(result, expected) {
		t.Errorf("Output %d not equal with the expected %d", result, expected)
	}
}
