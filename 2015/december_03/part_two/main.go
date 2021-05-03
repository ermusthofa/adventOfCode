package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var onlyOnce sync.Once

type coordinate struct {
	x int32
	y int32
}

type delivery struct {
	coordinate
	trackingLocation []coordinate
}

func main() {

	defer measureTime("1st.go")()

	clues, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	santa := newDelivery()
	roboSanta := newDelivery()

	for i, c := range clues {
		if modZero(i) {
			roboSanta.route(c)
		} else {
			santa.route(c)
		}
	}

	aggregatedCoordinate := findAggregatedCoordinate(santa.trackingLocation, roboSanta.trackingLocation)
	uniqueCoordinate := findUniqueCoordinate(aggregatedCoordinate)

	fmt.Printf("Number of Houses that received a gift are %d\n", len(uniqueCoordinate))

}

func newDelivery() *delivery {
	return &delivery{
		coordinate: coordinate{
			x: 0,
			y: 0,
		},
		trackingLocation: []coordinate{
			{
				x: 0,
				y: 0,
			},
		},
	}
}

func (d *delivery) route(c byte) {
	switch string(c) {
	case "^":
		d.goNorth()
	case ">":
		d.goEast()
	case "v":
		d.goSouth()
	case "<":
		d.goWest()
	default:
	}
}

func findAggregatedCoordinate(x, y []coordinate) []coordinate {
	return append(x, y...)
}

func findUniqueCoordinate(c []coordinate) []coordinate {
	keys := make(map[coordinate]bool)
	list := []coordinate{}
	for _, entry := range c {
		v := keys[entry]
		if !v {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (d *delivery) goNorth() {
	d.coordinate.y++
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func (d *delivery) goSouth() {
	d.coordinate.y--
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func (d *delivery) goEast() {
	d.coordinate.x++
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func (d *delivery) goWest() {
	d.coordinate.x--
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func modZero(i int) bool {
	return i%2 == 0
}

func measureTime(appName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s application is %v \n", appName, time.Since(start))
	}
}
