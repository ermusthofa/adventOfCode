package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	clues, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	floor, pos := countFloors(clues)
	fmt.Printf("Santa should be at basement (floor %d) at position %d\n", floor, pos)
}

func countFloors(f []uint8) (int32, int32) {
	var floor, pos int32
	pos = 1
	for _, c := range f {
		if goUp(c) {
			floor++
		}
		if goDown(c) {
			floor--
		}
		if isBasement(floor) {
			break
		}
		pos++
	}
	return floor, pos
}

func isBasement(floor int32) bool {
	return true == (floor == -1)
}

func goUp(sign byte) bool {
	return string(sign) == "("
}

func goDown(sign byte) bool {
	return string(sign) == ")"
}
