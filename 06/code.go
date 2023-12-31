package main

import (
	"fmt"
	"math"
)

type race struct {
	time     int
	distance int
}

func calculate(r race) int {
	counter := 0
	for i := 1; i < r.time; i++ {
		if i*(r.time-i) > r.distance {
			counter++
		}
	}
	return counter
}

// brute force
func part1() {
	// test data -> 288
	// fmt.Println(calculate(race{7, 9}) * calculate(race{15, 40}) * calculate(race{30, 200}))

	// real data
	fmt.Println(calculate(race{48, 261}) * calculate(race{93, 1192}) *
		calculate(race{84, 1019}) * calculate(race{66, 1063}))
}

func part2() {
	//test data
	//t := -71530.0
	//d := 940200.0

	//real data
	t := -48938466.0
	d := 261119210191063.0
	discr := math.Sqrt(t*t - 4*d)
	x1 := int(math.Ceil((-t - discr) / 2.0))
	x2 := int(math.Floor((-t + discr) / 2.0))
	spread := x2 - x1 + 1
	fmt.Printf("%d: %d-%d\n", spread, x1, x2)
}
func main() {
	//part1()
	part2()

}
