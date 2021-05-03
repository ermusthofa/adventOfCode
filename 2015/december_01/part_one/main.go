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

	floor := countFloors(clues)
	fmt.Printf("Santa should go to floor %d\n", floor)
}

func countFloors(f []uint8) int32 {
	var floor int32
	for _, c := range f {
		if goUp(c) {
			floor++
		}
		if goDown(c) {
			floor--
		}
	}
	return floor
}

func goUp(sign byte) bool {
	return string(sign) == "("
}

func goDown(sign byte) bool {
	return string(sign) == ")"
}
